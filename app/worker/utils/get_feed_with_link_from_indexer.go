package utils

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"net/http"
)

type IndexerResponse struct {
	List []struct {
		//CharacterID     uint   `json:"characterId"`
		//NoteID          uint   `json:"noteId"`
		Uri             string `json:"uri"`
		TransactionHash string `json:"transactionHash"`
		// Ignore other fields for now
	} `json:"list"`
	//Count uint `json:"count"`
	// Ignore cursor
}

func GetFeedWithLinkFromIndexer(link string) (string, string, error) {
	// Refer to https://indexer.crossbell.io/v1/notes?externalUrls=https%3A%2F%2Ftwitter.com%2FNyaRSS3%2Fstatus%2F1534713583270998018&includeDeleted=false&includeEmptyMetadata=false&limit=1

	// Prepare request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/notes", config.Config.CrossbellIndexer), nil)
	if err != nil {
		global.Logger.Errorf("Failed to prepare request for indexer query %s with error: %s", link, err.Error())
		return "", "", err
	}

	// Prepare query
	q := req.URL.Query()
	q.Add("externalUrls", link)
	q.Add("includeDeleted", "false")
	q.Add("includeEmptyMetadata", "false")
	q.Add("limit", "1")
	req.URL.RawQuery = q.Encode()

	// Execute request
	rawRes, err := (&http.Client{}).Do(req)
	if err != nil {
		global.Logger.Errorf("Failed to execute request for indexer query %s with error: %s", link, err.Error())
		return "", "", err
	}

	// Parse response
	var res IndexerResponse
	err = json.NewDecoder(rawRes.Body).Decode(&res)
	if err != nil {
		global.Logger.Errorf("Failed to parse request for indexer query %s with error: %s", link, err.Error())
		return "", "", err
	}

	// Check response
	if len(res.List) == 0 {
		// Nothing
		return "", "", fmt.Errorf("no maching found")
	} else {
		return res.List[0].Uri, res.List[0].TransactionHash, nil
	}
}
