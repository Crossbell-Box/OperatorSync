package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func UnbindAccount(ctx *gin.Context) {
	// Only accept when fails validate

	reqCharacter := ctx.Param("character")
	reqPlatform := ctx.Param("platform")
	reqUsername := ctx.Param("username")

	if _, ok := commonConsts.SUPPORTED_PLATFORM[reqPlatform]; !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Platform not supported",
		})
		return
	}

	global.Logger.Debug("Account ", reqCharacter, reqPlatform, reqUsername, " bind request received.")

	// Check if accounts already exists
	var account types.Account
	if err := global.DB.First(
		&account,
		"crossbell_character = ? AND platform = ? AND username = ?",
		reqCharacter, reqPlatform, reqUsername,
	).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Yes
		global.Logger.Debug("Account ", reqCharacter, reqPlatform, reqUsername, " not exist")

		// No exist
		global.Logger.Debug("Account ", reqCharacter, reqPlatform, reqUsername, " record already exists")
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Account not exist (might already unbind)",
			"result":  true,
		})

	} else if err != nil {
		// Failed
		global.Logger.Error("Account ", reqCharacter, reqPlatform, reqUsername, "failed to retrieve data from database with error: ", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Failed to retrieve data from database.",
		})
	} else {

		if ok, err := utils.ValidateAccount(reqCharacter, reqPlatform, reqUsername); err != nil {
			global.Logger.Error("Account ", reqCharacter, reqPlatform, reqUsername, " failed to finish account validate process with error: ", err.Error())
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      false,
				"message": "Failed to finish account validate process.",
			})
		} else if !ok {
			global.Logger.Debug("Account ", reqCharacter, reqPlatform, reqUsername, " unbind")
			// Update database
			global.DB.Delete(&account)
			// Clear cache
			listAccountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:list", reqCharacter)
			global.Redis.Del(context.Background(), listAccountsCacheKey)
			// Response
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Account unbind",
				"result":  true,
			})
		} else {
			global.Logger.Debug("Account ", reqCharacter, reqPlatform, reqUsername, " validate information still exists")
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Account validate information still exists, please remove it.",
				"result":  false,
			})
		}
	}
}
