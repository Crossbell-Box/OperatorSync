package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func RefreshAccounts(ctx *gin.Context) {
	// Parse request params
	reqCharacter := ctx.Param("character")
	if reqCharacter == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Please specify character",
		})
		return
	}

	refreshAccountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:refresh", reqCharacter)

	// Check if is in cool-down time
	getCacheCtx, cancelGetCacheCtx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelGetCacheCtx()
	if exist, err := global.Redis.Exists(getCacheCtx, refreshAccountsCacheKey).Result(); err != nil {
		global.Logger.Error("Unable to check accounts cache")
	} else if exist > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Please try again later",
		})
		return
	}

	// Check if is in database
	var character models.Character
	if err := global.DB.First(&character, "crossbell_character = ?", reqCharacter).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// No, so it's time to insert a new one
		character.CrossbellCharacter = reqCharacter
		character.MediaUsage = 0
		global.DB.Create(&character)
	}

	accounts, err := utils.RefreshAccountForCharacter(reqCharacter)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": fmt.Sprintf("Internal server error: %s", err.Error()),
		})
		return
	}

	// Respond
	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Accounts flushed",
		"result":  accounts,
	})

	setCacheCtx := context.Background()
	global.Redis.Set(setCacheCtx, refreshAccountsCacheKey, time.Now(), consts.ACCOUNT_REFRESH_COOLDOWN)

}
