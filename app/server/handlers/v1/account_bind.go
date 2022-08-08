package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/jobs"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func BindAccount(ctx *gin.Context) {
	// Only accept when pass validate

	reqCharacterID := ctx.Param("character") // ID
	reqPlatform := ctx.Param("platform")
	reqUsername := ctx.Param("username")

	if _, ok := commonConsts.SUPPORTED_PLATFORM[reqPlatform]; !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Platform not supported",
		})
		return
	}

	// User selectable start time
	nowTimeStr := time.Now().Format(time.RFC3339)
	startFromStr := ctx.DefaultQuery("from", nowTimeStr)
	startFromTime, err := time.Parse(time.RFC3339, startFromStr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": fmt.Sprintf("Failed to parse time with error (%s), please use RFC3339 format (%s)", err.Error(), nowTimeStr),
		})
		return
	}

	// Check character -> create if not exist
	if err := global.DB.First(&types.Character{}, "crossbell_character_id = ?", reqCharacterID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debugf("Character %s doesn't exist, creating...", reqCharacterID)
		global.DB.Create(&types.Character{
			CrossbellCharacterID: reqCharacterID,
		})
	}

	global.Logger.Debugf("Account #%s (%s@%s) bind request received.", reqCharacterID, reqUsername, reqPlatform)

	var account models.Account

	// Check if accounts already exists
	if err := global.DB.First(
		&account,
		"crossbell_character_id = ? AND platform = ? AND username = ?",
		reqCharacterID, reqPlatform, reqUsername,
	).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Not exist

		// Check if already bind by others
		if err := global.DB.First(
			&account,
			"platform = ? AND username = ?",
			reqPlatform, reqUsername,
		).Error; !errors.Is(err, gorm.ErrRecordNotFound) && account.CrossbellCharacterID != reqCharacterID {
			// Already bind but not this one
			global.Logger.Debugf("Account (%s@%s) has already been occupied by #%s", reqUsername, reqPlatform, reqCharacterID)
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": fmt.Sprintf("Account (%s@%s) has already been occupied by #%s, please unbind it first.", reqUsername, reqPlatform, reqCharacterID),
				"result":  false,
			})
			return
		}

		// Check if already bind account on this platform
		if err := global.DB.First(
			&account,
			"crossbell_character_id = ? AND platform = ?",
			reqCharacterID, reqPlatform,
		).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			// Already bind but not this one
			global.Logger.Debugf("Account #%s already has an account (%s) on platform %s", reqCharacterID, reqUsername, reqPlatform)
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": fmt.Sprintf("Account #%s already has an account (%s) on platform %s", reqCharacterID, reqUsername, reqPlatform),
				"result":  false,
			})
			return
		}

		// Nope? So it's empty and available to bind.
		global.Logger.Debugf("Account #%s (%s@%s) not exist, start validating...", reqCharacterID, reqUsername, reqPlatform)

		if ok, err := jobs.ValidateAccount(reqCharacterID, reqPlatform, reqUsername); err != nil {
			global.Logger.Errorf("Account #%s (%s@%s) failed to finish account validate process with error: %s", reqUsername, reqPlatform, reqCharacterID, err.Error())
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      false,
				"message": "Failed to finish account validate process.",
			})
		} else if ok {
			global.Logger.Debugf("Account #%s (%s@%s) added.", reqCharacterID, reqUsername, reqPlatform)
			// Update database
			account.Account = types.Account{
				CrossbellCharacterID: reqCharacterID,
				Platform:             reqPlatform,
				Username:             reqUsername,
				LastUpdated:          startFromTime,
				UpdateInterval:       commonConsts.SUPPORTED_PLATFORM[reqPlatform].MinRefreshGap,
				NextUpdate:           time.Now(),
				FeedsCount:           0,
				NotesCount:           0,
				MediaUsage:           nil,
			}
			account.OnChainStatusManageForAccount = types.OnChainStatusManageForAccount{
				IsOnChainPaused: false,
			}
			global.DB.Create(&account)
			// Clear cache
			listAccountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:list", reqCharacterID)
			global.Redis.Del(context.Background(), listAccountsCacheKey)
			// Response
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Account added",
				"result":  true,
			})
		} else {
			global.Logger.Debugf("Account #%s (%s@%s) no necessary validate info found.", reqCharacterID, reqUsername, reqPlatform)
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Necessary validate info not found",
				"result":  false,
			})
		}

	} else if err != nil {
		// Failed
		global.Logger.Errorf("Account #%s (%s@%s) failed to retrieve data from database with error: %s", reqUsername, reqPlatform, reqCharacterID, err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Failed to retrieve data from database.",
		})
	} else {
		// Already exists
		global.Logger.Debugf("Account #%s (%s@%s) record already exists.", reqCharacterID, reqUsername, reqPlatform)

		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Account already added",
			"result":  true,
		})
	}
}
