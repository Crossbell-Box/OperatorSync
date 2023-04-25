package v1

import (
	"errors"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func ForceSyncAccount(ctx *gin.Context) {

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
		// Check if is just bind
		if account.LastUpdated.After(time.Unix(0, 0)) {
			// Already synced, check if eligible to sync now
			earliestNextUpdate := account.LastUpdated.Add(commonConsts.SUPPORTED_PLATFORM[reqPlatform].MinRefreshGap)
			if earliestNextUpdate.Before(time.Now()) {
				account.NextUpdate = time.Now()
			} else {
				account.NextUpdate = earliestNextUpdate
			}

			// Save account
			global.DB.Save(&account)
		} else {
			// Yes, it's just bind, prevent 1970-1-1
			account.LastUpdated = account.NextUpdate
		}

		// Parse accounts update interval to seconds
		account.UpdateInterval = time.Duration(account.UpdateInterval.Seconds())

		// Response
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Account pending to sync",
			"result":  account,
		})
	}
}
