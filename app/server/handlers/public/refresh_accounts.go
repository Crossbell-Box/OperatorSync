package public

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
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
	var dbAccounts []models.Account

	// Get accounts from database
	global.DB.
		Model(&models.Account{}).
		Order("created_at").
		Find(&dbAccounts)

	// Get accounts from CrossBell
	csbAccounts, err := utils.GetAccountsFromCrossbell(reqCharacter)
	if err != nil {
		// Internal server error
		global.Logger.Error("Failed to get accounts from CSB for (", reqCharacter, ") : ", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "Internal server error: " + err.Error(),
		})
		return
	}

	// Prepare accounts
	var accounts []models.Account

	dbAccountsShouldKeep := make([]bool, len(dbAccounts))
	csbAccountsAlreadyExist := make([]bool, len(csbAccounts))

	/// Hint: Golang would initialize memory automatically, so there's no need to manually set all member to `false`

	// Check accounts
	for csbI, accountFromCSB := range csbAccounts {
		accounts = append(accounts, models.Account{
			Account: accountFromCSB,
		})
		for dbI, accountFromDB := range dbAccounts {
			if accountFromDB.Platform == accountFromCSB.Platform && accountFromDB.Username == accountFromCSB.Username {
				csbAccountsAlreadyExist[csbI] = true
				dbAccountsShouldKeep[dbI] = true
			}
		}
	}

	// Delete non-exist accounts
	for dbI, shouldKeep := range dbAccountsShouldKeep {
		if !shouldKeep {
			global.DB.Delete(&dbAccounts[dbI])
		}
	}

	// Insert new accounts
	for csbI, alreadyExist := range csbAccountsAlreadyExist {
		if !alreadyExist {
			global.DB.Create(&csbAccounts[csbI])
		}
	}

	// Cache key
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
