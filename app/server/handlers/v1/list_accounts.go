package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	commonGlobal "github.com/Crossbell-Box/OperatorSync/common/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ListAccounts(ctx *gin.Context) {
	// Parse request params
	reqCharacterID := ctx.Param("character")

	// Get requests
	var accounts []models.Account
	var accountsForResponse []types.AccountWithAdditionalPropsForListResponse

	// Use redis cache
	cacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:list", reqCharacterID)

	// Get data from cache
	getCacheCtx, cancelGetCacheCtx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelGetCacheCtx()
	if exist, err := commonGlobal.Redis.Exists(getCacheCtx, cacheKey).Result(); err != nil {
		global.Logger.Error("Unable to check accounts cache")
	} else if exist > 0 {
		if accountsBytes, err := commonGlobal.Redis.Get(getCacheCtx, cacheKey).Bytes(); err != nil {
			global.Logger.Error("Unable to get accounts cache")
		} else if err = json.Unmarshal(accountsBytes, &accountsForResponse); err != nil {
			global.Logger.Error("Unable to parse accounts cache")
			// Clear invalid cache
			commonGlobal.Redis.Del(getCacheCtx, cacheKey)
		} else {
			// Successfully parsed

			ctx.JSON(http.StatusOK, gin.H{
				"ok":      true,
				"message": "Records found",
				"result":  accountsForResponse,
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

	// Accounts on crossbell chain
	_, accountsOnChain, _ := utils.CheckOnChainData(reqCharacterID)

	// Process account for result check
	for _, rawAccount := range accounts {
		// Fill status
		accountForResponse := types.AccountWithAdditionalPropsForListResponse{
			Account:                       rawAccount.Account,
			OnChainStatusManageForAccount: rawAccount.OnChainStatusManageForAccount,
			IsMetadataCorrect:             utils.IsInConnectedAccounts(rawAccount.Platform, rawAccount.Username, accountsOnChain),
		}

		// Parse accounts update interval to seconds
		accountForResponse.UpdateInterval = time.Duration(accountForResponse.UpdateInterval.Seconds())
		if !accountForResponse.LastUpdated.After(time.Unix(0, 0)) {
			// Prevent 1970-1-1
			accountForResponse.LastUpdated = accountForResponse.NextUpdate
		}

		accountsForResponse = append(accountsForResponse, accountForResponse)
	}

	// Respond
	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Records found",
		"result":  accountsForResponse,
	})

	// Set cache
	/*
		setCacheCtx := context.Background()
		if accountsBytes, err := json.Marshal(&accountsForResponse); err != nil {
			// WTF?
			global.Logger.Error("Failed to parse accounts")
		} else {
			commonGlobal.Redis.Set(setCacheCtx, cacheKey, accountsBytes, consts.ACCOUNT_LIST_CACHE_EXPIRE)
		}

	*/

}
