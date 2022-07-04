package public

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/gin-gonic/gin"
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

	// Use redis cache
	refreshAccountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:refresh", reqCharacter)

	// Get data from cache
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

	// Start flush account work
	var accounts []models.Account

	// TODO: Finish this

	// Reset cache key
	listAccountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:list", reqCharacter)

	// Set cache
	setCacheCtx := context.Background()
	if accountsBytes, err := json.Marshal(&accounts); err != nil {
		// WTF?
		global.Logger.Error("Failed to parse accounts")
	} else {
		global.Redis.Set(setCacheCtx, listAccountsCacheKey, accountsBytes, consts.ACCOUNT_LIST_CACHE_EXPIRE)
	}
	global.Redis.Set(setCacheCtx, refreshAccountsCacheKey, time.Now(), consts.ACCOUNT_REFRESH_COOLDOWN)

	// Respond
	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Accounts flushed",
		"result":  accounts,
	})

}
