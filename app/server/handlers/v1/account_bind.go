package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
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

	// Check character -> create if not exist
	if err := global.DB.First(&types.Character{}, "crossbell_character_id = ?", reqCharacterID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Debugf("Character %s doesn't exist, creating...", reqCharacterID)
		global.DB.Create(&types.Character{
			CrossbellCharacterID: reqCharacterID,
			AccountLastUpdatedAt: time.Now(),
			MediaUsage:           0,
		})
	}

	global.Logger.Debugf("Account #%s (%s@%s) bind request received.", reqCharacterID, reqUsername, reqPlatform)

	// Check if accounts already exists
	var account models.Account

	if err := global.DB.First(
		&account,
		"platform = ? AND username = ?",
		reqPlatform, reqUsername,
	).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Not exist
		global.Logger.Debugf("Account #%s (%s@%s) not exist, start validating...", reqCharacterID, reqUsername, reqPlatform)

		if ok, err := utils.ValidateAccount(reqCharacterID, reqPlatform, reqUsername); err != nil {
			global.Logger.Error("Account ", reqCharacterID, "- (", reqPlatform, "-", reqUsername, ") failed to finish account validate process with error: ", err.Error())
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
				LastUpdated:          time.Now(),
				UpdateInterval:       0,
				NextUpdate:           time.Now(),
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
	} else if account.CrossbellCharacterID != reqCharacterID {
		// Already bind but not this one
		global.Logger.Errorf("Account (%s@%s) has already been occupied by #%s", reqUsername, reqPlatform, reqCharacterID)
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": fmt.Sprintf("Account (%s@%s) has already been occupied by #%s, please unbind it first.", reqUsername, reqPlatform, reqCharacterID),
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
