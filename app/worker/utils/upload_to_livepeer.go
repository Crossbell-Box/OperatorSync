package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"io"
	"net/http"
)

type LivePeerUploadURLRequest struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type LivePeerUploadURLResponse struct {
	Asset struct {
		ID         string `json:"id"`
		PlaybackID string `json:"playbackId"`
		UserID     string `json:"userId"`
		CreatedAt  uint   `json:"createdAt"`
		Status     string `json:"status"`
		Name       string `json:"name"`
	} `json:"asset"`
	Task struct {
		ID            string `json:"id"`
		CreatedAt     uint   `json:"createdAt"`
		Type          string `json:"type"`
		OutputAssetID string `json:"outputAssetId"`
		UserID        string `json:"userId"`
		Params        struct {
			Import struct {
				URL string `json:"url"`
			} `json:"import"`
		} `json:"params"`
		Status struct {
			Phase     string `json:"phase"`
			UpdatedAt uint   `json:"updatedAt"`
		} `json:"status"`
	} `json:"task"`
}

func UploadURIToLivePeer(url string, name string) (string, string, error) {
	// Upload docs https://docs.livepeer.studio/guides/on-demand#upload-with-url

	reqBody := LivePeerUploadURLRequest{
		URL:  url,
		Name: name,
	}
	reqBodyBytes, err := json.Marshal(&reqBody)
	if err != nil {
		global.Logger.Errorf("Failed to parse media upload request for URI %s with error: %s", url, err.Error())
		return "", "", err
	}
	req, err := http.NewRequest("POST", "https://livepeer.studio/api/asset/import", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		global.Logger.Errorf("Failed to crete request with error: %s", err.Error())
		return "", "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Config.LivePeerAPIKey))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		global.Logger.Errorf("Failed to do http request with error: %s", err.Error())
		return "", "", err
	}
	var resBody LivePeerUploadURLResponse
	resBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		global.Logger.Errorf("Failed to read bytes from response body with error: %s", err.Error())
		return "", "", err
	}
	err = json.Unmarshal(resBodyBytes, &resBody)
	if err != nil {
		global.Logger.Errorf("Failed to parse http response result with error: %s", err.Error())
		return "", "", err
	}

	// Playback docs https://docs.livepeer.studio/guides/#playback-a-livestream
	return fmt.Sprintf("https://livepeercdn.com/hls/%s/index.m3u8", resBody.Asset.PlaybackID), string(resBodyBytes), nil
}
