package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	"gorm.io/gorm"
)

type UnbindAccountProps struct {
	// Basic information
	CrossbellCharacterID string
	Platform             string
	Username             string
}

func AccountUnbind(props *UnbindAccountProps) (bool, string, error) {
	if _, ok := commonConsts.SUPPORTED_PLATFORM[props.Platform]; !ok {
		return false, "Platform not supported", fmt.Errorf("platform not supported")
	}

	global.Logger.Debugf("Account #%s (%s@%s) unbind request received.", props.CrossbellCharacterID, props.Username, props.Platform)

	// Check if accounts already exists
	var account models.Account
	if err := global.DB.First(
		&account,
		"crossbell_character_id = ? AND platform = ? AND username = ?",
		props.CrossbellCharacterID, props.Platform, props.Username,
	).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// No exist
		global.Logger.Debugf("Account #%s (%s@%s) not exist", props.CrossbellCharacterID, props.Username, props.Platform)
		return true, "Account not exist (might already unbind)", nil
	} else if err != nil {
		// Failed
		global.Logger.Errorf("Account #%s (%s@%s) failed to retrieve data from database with error: %s", props.Username, props.Platform, props.CrossbellCharacterID, err.Error())
		return false, "Failed to retrieve data from database.", err
	} else {

		//// Check if connected account is not removed
		//if _, isAccountConnected, err := CheckOnChainData(props.CrossbellCharacterID, props.Platform, props.Username); err != nil {
		//	global.Logger.Errorf("Failed to check operator for %s", props.CrossbellCharacterID)
		//	// Just ignore
		//} else if isAccountConnected {
		//	// Oops
		//	return false, "Account not disconnected", nil
		//}

		if ok, err := ValidateAccount(props.CrossbellCharacterID, props.Platform, props.Username); err != nil {
			global.Logger.Errorf("Account #%s (%s@%s) failed to finish validate request with error: %s", props.Username, props.Platform, props.CrossbellCharacterID, err.Error())
			return false, "Failed to finish account validate process.", err
		} else if !ok {
			global.Logger.Debugf("Account #%s (%s@%s) unbind.", props.CrossbellCharacterID, props.Username, props.Platform)
			// Update database
			global.DB.Delete(&account)
			// Clear cache
			listAccountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:list", props.CrossbellCharacterID)
			commonGlobal.Redis.Del(context.Background(), listAccountsCacheKey)
			// Response
			return true, "Account unbind", nil
		} else {
			global.Logger.Debugf("Account #%s (%s@%s) validate information still exists.", props.CrossbellCharacterID, props.Username, props.Platform)
			return false, "Account validate information still exists, please remove it.", nil
		}
	}

}
