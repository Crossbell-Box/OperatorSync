package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	"gorm.io/gorm"
	"time"
)

type BindAccountProps struct {
	// Basic information
	CrossbellCharacterID string
	Platform             string
	Username             string

	// Advanced options
	StartFrom time.Time
	IsInherit bool
}

func AccountBind(props *BindAccountProps) (bool, string, error) {
	if _, ok := commonConsts.SUPPORTED_PLATFORM[props.Platform]; !ok {
		return false, "Platform not supported", fmt.Errorf("platform not supported")
	}

	// Check character -> create if not exist
	if err := global.DB.First(&models.Character{}, "crossbell_character_id = ?", props.CrossbellCharacterID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debugf("Character %s doesn't exist, creating...", props.CrossbellCharacterID)
		if _, err := CreateOrRecoverCharacter(props.CrossbellCharacterID); err != nil {
			return false, "Failed to activate character", err
		}
	}

	global.Logger.Debugf("Account #%s (%s@%s) bind request received.", props.CrossbellCharacterID, props.Username, props.Platform)

	var account models.Account

	// Check if accounts already exists
	if err := global.DB.First(
		&account,
		"crossbell_character_id = ? AND platform = ? AND username = ?",
		props.CrossbellCharacterID, props.Platform, props.Username,
	).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Not exist

		// Check if already bind by others
		if err := global.DB.First(
			&account,
			"platform = ? AND username = ?",
			props.Platform, props.Username,
		).Error; !errors.Is(err, gorm.ErrRecordNotFound) && account.CrossbellCharacterID != props.CrossbellCharacterID {
			// Already bind but not this one
			global.Logger.Debugf("Account (%s@%s) has already been occupied by #%s", props.Username, props.Platform, props.CrossbellCharacterID)
			return false, fmt.Sprintf("Account (%s@%s) has already been occupied by #%s, please unbind it first.", props.Username, props.Platform, props.CrossbellCharacterID), nil
		}

		// Check if already bind account on this platform
		if err := global.DB.First(
			&account,
			"crossbell_character_id = ? AND platform = ?",
			props.CrossbellCharacterID, props.Platform,
		).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			// Already bind but not this one
			global.Logger.Debugf("Account #%s already has an account (%s) on platform %s", props.CrossbellCharacterID, props.Username, props.Platform)
			return false, fmt.Sprintf("Account #%s already has an account (%s) on platform %s", props.CrossbellCharacterID, props.Username, props.Platform), nil
		}

		// Nope? So it's empty and available to bind.
		global.Logger.Debugf("Account #%s (%s@%s) not exist, start validating...", props.CrossbellCharacterID, props.Username, props.Platform)

		// Check if operator is set
		if ok, err := CheckOperator(props.CrossbellCharacterID); err != nil {
			global.Logger.Errorf("Failed to check operator for %s", props.CrossbellCharacterID)
			// Just ignore
		} else if !ok {
			// Oops
			return false, "Operator not set properly", nil
		}

		if ok, err := ValidateAccount(props.CrossbellCharacterID, props.Platform, props.Username); err != nil {
			global.Logger.Errorf("Account #%s (%s@%s) failed to finish account validate process with error: %s", props.Username, props.Platform, props.CrossbellCharacterID, err.Error())
			return false, "Failed to finish account validate process.", err
		} else if ok {
			global.Logger.Debugf("Account #%s (%s@%s) added.", props.CrossbellCharacterID, props.Username, props.Platform)
			// Update database
			possibleDeletedAccount := models.Account{
				ID: 0,
			}
			if props.IsInherit {
				// Check if is abandoned
				global.DB.Unscoped().Where(
					"platform = ? AND username = ?",
					props.Platform, props.Username,
				).Order("deleted_at DESC").First(
					&possibleDeletedAccount,
				)
			}
			if possibleDeletedAccount.ID > 0 {
				// Inherit
				account.Account = possibleDeletedAccount.Account
				account.OnChainStatusManageForAccount = possibleDeletedAccount.OnChainStatusManageForAccount
				// But new
				account.Account.CrossbellCharacterID = props.CrossbellCharacterID
			} else {
				// Create new
				account.Account = types.Account{
					CrossbellCharacterID: props.CrossbellCharacterID,
					Platform:             props.Platform,
					Username:             props.Username,
					LastUpdated:          props.StartFrom,
					UpdateInterval:       commonConsts.SUPPORTED_PLATFORM[props.Platform].MinRefreshGap,
					NextUpdate:           time.Now(),
					FeedsCount:           0,
					NotesCount:           0,
					MediaUsage:           nil,
				}
				account.OnChainStatusManageForAccount = types.OnChainStatusManageForAccount{
					IsOnChainPaused: false,
				}
			}
			global.DB.Create(&account)
			// Clear cache
			listAccountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:list", props.CrossbellCharacterID)
			commonGlobal.Redis.Del(context.Background(), listAccountsCacheKey)
			// Response
			return true, "Account added", nil
		} else {
			global.Logger.Debugf("Account #%s (%s@%s) no necessary validate info found.", props.CrossbellCharacterID, props.Username, props.Platform)
			return false, "Necessary validate info not found", nil
		}

	} else if err != nil {
		// Failed
		global.Logger.Errorf("Account #%s (%s@%s) failed to retrieve data from database with error: %s", props.Username, props.Platform, props.CrossbellCharacterID, err.Error())
		return false, "Failed to retrieve data from database.", err
	} else {
		// Already exists
		global.Logger.Debugf("Account #%s (%s@%s) record already exists.", props.CrossbellCharacterID, props.Username, props.Platform)
		return true, "Account already added", nil
	}

}
