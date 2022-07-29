package v1

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

func ListAccounts(ctx *gin.Context) {
	// Parse request params
	reqCharacterID := ctx.Param("character")

	// Get requests
	var accounts []models.Account

	// Use redis cache
	cacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:list", reqCharacterID)

	// Get data from cache
	getCacheCtx, cancelGetCacheCtx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelGetCacheCtx()
	if exist, err := global.Redis.Exists(getCacheCtx, cacheKey).Result(); err != nil {
		global.Logger.Error("Unable to check accounts cache")
	} else if exist > 0 {
		if accountsBytes, err := global.Redis.Get(getCacheCtx, cacheKey).Bytes(); err != nil {
			global.Logger.Error("Unable to get accounts cache")
		} else if err = json.Unmarshal(accountsBytes, &accounts); err != nil {
			global.Logger.Error("Unable to parse accounts cache")
			// Clear invalid cache
			global.Redis.Del(getCacheCtx, cacheKey)
		} else {
			// Successfully parsed

			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Records found",
				"result":  accounts,
			})

			return
		}
	}

	// Get data from database
	global.DB.
		Model(&models.Account{}).
		Where("crossbell_character_id = ?", reqCharacterID).
		Order("created_at").
		Find(&accounts)

	// Set cache
	setCacheCtx := context.Background()
	if accountsBytes, err := json.Marshal(&accounts); err != nil {
		// WTF?
		global.Logger.Error("Failed to parse accounts")
	} else {
		global.Redis.Set(setCacheCtx, cacheKey, accountsBytes, consts.ACCOUNT_LIST_CACHE_EXPIRE)
	}

	// Respond
	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Records found",
		"result":  accounts,
	})

}
