package utils

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"net/http"
	"strings"
	"time"
)

type connectedAccounts struct {
	URI string `json:"uri"`
	// Ignore others
}

type crossbellCharacter struct {
	ConnectedAccounts []connectedAccounts `json:"connected_accounts"`
	// Ignore others
}

type crossbellIndexerCharacterRes struct {
	Metadata struct {
		Content crossbellCharacter `json:"content"`
		// Ignore others
	} `json:"metadata"`
	// Ignore others
}

func GetAccountsFromCrossbell(character string) ([]types.Account, error) {
	// With character ID https://indexer.crossbell.io/v1/characters/19
	// With handle https://indexer.crossbell.io/v1/handles/candinya/character

	reqUrl := fmt.Sprintf("https://indexer.crossbell.io/v1/characters/%s", character)
	rawRes, err := (&http.Client{}).Get(reqUrl)
	if err != nil {
		global.Logger.Error("Failed to get character info for ", character, " : ", err.Error())
		return nil, err
	}

	var res crossbellIndexerCharacterRes
	if err = json.NewDecoder(rawRes.Body).Decode(&res); err != nil {
		global.Logger.Error("Failed to parse response data: ", err.Error())
		return nil, err
	}

	var accounts []types.Account
	for _, accountDetail := range res.Metadata.Content.ConnectedAccounts {
		// csb://account:Username@platform (latter @)
		usernamePplatformStr := strings.ReplaceAll(accountDetail.URI, "csb://account:", "")
		usernamePplatformSlice := strings.Split(usernamePplatformStr, "@")
		platform := usernamePplatformSlice[len(usernamePplatformSlice)-1]
		if _, ok := commonConsts.SUPPORTED_PLATFORM[platform]; !ok {
			// Platform not supported yet, skip
			continue
		}
		username := strings.Join(usernamePplatformSlice[:len(usernamePplatformSlice)-1], "@")
		accounts = append(accounts, types.Account{
			CrossbellCharacter: character,
			Platform:           platform,
			Username:           username,
			LastUpdated:        time.Now(),
			UpdateInterval:     0,
			NextUpdate:         time.Now(),
		})
	}

	return accounts, nil
}
