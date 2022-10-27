package v1

import (
	"errors"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func CheckAccountOnChainStatus(ctx *gin.Context) {

	reqCharacterID := ctx.Param("character")
	reqPlatform := ctx.Param("platform")
	reqUsername := ctx.Param("username")

	if _, ok := commonConsts.SUPPORTED_PLATFORM[reqPlatform]; !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Platform not supported",
			"result":  nil,
		})
		return
	}

	global.Logger.Debugf("Account #%s (%s@%s) force sync request received.", reqCharacterID, reqUsername, reqPlatform)

	// Check if account exists
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
			"message": "Account not exist",
			"result":  nil,
		})
	} else if err != nil {
		// Failed
		global.Logger.Errorf("Account #%s (%s@%s) failed to retrieve data from database with error: %s", reqUsername, reqPlatform, reqCharacterID, err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Failed to retrieve data from database.",
			"result":  nil,
		})
	} else {
		// Check if anything stuck
		var (
			pendingNotes int64
			succeedNotes int64
		)
		err = global.DB.Scopes(models.FeedTable(models.Feed{
			Feed: types.Feed{
				Platform: account.Platform,
			},
		})).Where("account_id = ? AND transaction = ?", account.ID, "").Count(&pendingNotes).Error
		if err != nil {
			global.Logger.Errorf("Account #%s (%s@%s) failed to count pending notes from database with error: %s", reqUsername, reqPlatform, reqCharacterID, err.Error())
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      false,
				"message": "Failed to count pending notes from database.",
				"result":  nil,
			})
			return
		}
		err = global.DB.Scopes(models.FeedTable(models.Feed{
			Feed: types.Feed{
				Platform: account.Platform,
			},
		})).Where("account_id = ? AND transaction <> ?", account.ID, "").Count(&succeedNotes).Error
		if err != nil {
			global.Logger.Errorf("Account #%s (%s@%s) failed to count succeeded notes from database with error: %s", reqUsername, reqPlatform, reqCharacterID, err.Error())
			ctx.JSON(http.StatusOK, gin.H{
				"ok":      false,
				"message": "Failed to count succeeded notes from database.",
				"result":  nil,
			})
			return
		}

		// Fix number
		account.NotesCount = uint(succeedNotes)
		if pendingNotes > 0 && !account.IsOnChainPaused {
			account.IsOnChainPaused = true
			account.OnChainPausedAt = time.Now()
			account.OnChainPauseMessage = "Unknown error, user request manual check"
		}

		// Save account
		global.DB.Save(&account)

		// Response
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Account pending to sync",
			"result":  account,
		})
	}
}
