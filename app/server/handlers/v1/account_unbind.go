package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func UnbindAccount(ctx *gin.Context) {
	// Only accept when fails validate

	reqCharacterID := ctx.Param("character")
	reqPlatform := ctx.Param("platform")
	reqUsername := ctx.Param("username")

	if _, ok := commonConsts.SUPPORTED_PLATFORM[reqPlatform]; !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Platform not supported",
		})
		return
	}

	global.Logger.Debugf("Account #%s (%s@%s) unbind request received.", reqCharacterID, reqUsername, reqPlatform)

	// Check if accounts already exists
	var account models.Account
	if err := global.DB.First(
		&account,
		"crossbell_character_id = ? AND platform = ? AND username = ?",
		reqCharacterID, reqPlatform, reqUsername,
	).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// No exist
		global.Logger.Debugf("Account #%s (%s@%s) not exist", reqCharacterID, reqUsername, reqPlatform)

		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Account not exist (might already unbind)",
			"result":  true,
		})
	} else if err != nil {
		// Failed
		global.Logger.Errorf("Account #%s (%s@%s) failed to retrieve data from database with error: %s", reqUsername, reqPlatform, reqCharacterID, err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Failed to retrieve data from database.",
		})
	} else {

		if ok, err := utils.ValidateAccount(reqCharacterID, reqPlatform, reqUsername); err != nil {
			global.Logger.Errorf("Account #%s (%s@%s) failed to finish validate request with error: %s", reqUsername, reqPlatform, reqCharacterID, err.Error())
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      false,
				"message": "Failed to finish account validate process.",
			})
		} else if !ok {
			global.Logger.Debugf("Account #%s (%s@%s) unbind.", reqCharacterID, reqUsername, reqPlatform)
			// Update database
			global.DB.Delete(&account)
			// Clear cache
			listAccountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:list", reqCharacterID)
			commonGlobal.Redis.Del(context.Background(), listAccountsCacheKey)
			// Response
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Account unbind",
				"result":  true,
			})
		} else {
			global.Logger.Debugf("Account #%s (%s@%s) validate information still exists.", reqCharacterID, reqUsername, reqPlatform)
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Account validate information still exists, please remove it.",
				"result":  false,
			})
		}
	}
}
