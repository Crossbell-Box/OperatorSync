package indexer

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/types"
	"net/http"
)

type characterMetadataIndexerResponse struct {
	Metadata types.CharacterMetadata `json:"metadata"`
	// Ignore other fields
}

func GetCharacterMetadataFromIndexer(characterIdStr string) (*types.CharacterMetadata, error) {
	// Prepare request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/characters/%s", config.Config.CrossbellIndexer, characterIdStr), nil)
	if err != nil {
		global.Logger.Errorf("Failed to prepare request for get character data (%s) from indexer with error: %s", characterIdStr, err.Error())
		return nil, err
	}

	// Execute request
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		global.Logger.Errorf("Failed to execute request for get character data (%s) from indexer with error: %s", characterIdStr, err.Error())
		return nil, err
	}

	// Parse response
	var resBody characterMetadataIndexerResponse
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		global.Logger.Errorf("Failed to parse response for get character data (%s) from indexer with error: %s", characterIdStr, err.Error())
		return nil, err
	}

	// Return
	return &resBody.Metadata, nil

}
