package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
)

func RefreshAccountForCharacter(character string) ([]models.Account, error) {
	// Start flush account work
	var dbAccounts []models.Account

	// Get accounts from database
	global.DB.
		Model(&models.Account{}).
		Order("created_at").
		Find(&dbAccounts)

	// Get accounts from CrossBell
	csbAccounts, err := GetAccountsFromCrossbell(character)
	if err != nil {
		// Internal server error
		global.Logger.Error("Failed to get accounts from CSB for (", character, ") : ", err.Error())
		return nil, err
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
	listAccountsCacheKey := fmt.Sprintf("%s:%s:%s", consts.CACHE_PREFIX, "accounts:list", character)

	// Set cache
	setCacheCtx := context.Background()
	if accountsBytes, err := json.Marshal(&accounts); err != nil {
		// WTF?
		global.Logger.Error("Failed to parse accounts")
	} else {
		global.Redis.Set(setCacheCtx, listAccountsCacheKey, accountsBytes, consts.ACCOUNT_LIST_CACHE_EXPIRE)
	}

	return accounts, nil
}
