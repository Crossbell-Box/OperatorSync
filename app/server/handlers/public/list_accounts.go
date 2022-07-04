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

func ListAccounts(ctx *gin.Context) {
	// Parse request params
	reqCharacter := ctx.Param("character")
	if reqCharacter == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Please specify character",
		})
		return
	}

	// Get requests
	var accounts []models.Account

	// Use redis cache
	cacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts", reqCharacter)

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
		Order("created_at").
		Find(&accounts)

	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Records found",
		"result":  accounts,
	})

	// Refresh accounts by retrieving from CrossBell chain (Character settings)

	// TODO: finish this

	// Set cache
	setCacheCtx := context.Background()
	if accountsBytes, err := json.Marshal(&accounts); err != nil {
		// WTF?
		global.Logger.Error("Failed to parse accounts")
	} else {
		global.Redis.Set(setCacheCtx, cacheKey, accountsBytes, consts.ACCOUNT_CACHE_EXPIRE)
	}

}
