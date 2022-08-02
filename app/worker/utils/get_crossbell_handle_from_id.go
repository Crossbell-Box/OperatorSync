package utils

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"net/http"
)

type crossbellIndexerCharacterRes struct {
	//CharacterID uint `json:"characterId"`
	Handle string `json:"handle"`
	// Ignore others
}

func GetCrossbellHandleFromID(characterID string) (string, error) {
	// With character ID https://indexer.crossbell.io/v1/characters/19
	// With handle https://indexer.crossbell.io/v1/handles/candinya/character

	reqUrl := fmt.Sprintf("https://indexer.crossbell.io/v1/characters/%s", characterID)
	rawRes, err := (&http.Client{}).Get(reqUrl)
	if err != nil {
		global.Logger.Error("Failed to get character info for ", characterID, " : ", err.Error())
		return "", err
	}

	var crossbellIndexerResponse crossbellIndexerCharacterRes
	if err = json.NewDecoder(rawRes.Body).Decode(&crossbellIndexerResponse); err != nil {
		global.Logger.Error("Failed to parse response data: ", err.Error())
		return "", err
	}

	return crossbellIndexerResponse.Handle, nil
}
