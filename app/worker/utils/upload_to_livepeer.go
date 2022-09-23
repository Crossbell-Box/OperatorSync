package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"net/http"
	"time"
)

type LivePeerUploadURLRequest struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type LivePeerAssetUploadStatus struct {
	Phase        string `json:"phase"` // "waiting" / "ready" / "failed"
	UpdatedAt    uint   `json:"updatedAt"`
	ErrorMessage string `json:"errorMessage"`
}

type LivePeerUploadURLResponse struct {
	Asset struct {
		ID         string                    `json:"id"`
		PlaybackID string                    `json:"playbackId"`
		UserID     string                    `json:"userId"`
		CreatedAt  uint                      `json:"createdAt"`
		Status     LivePeerAssetUploadStatus `json:"status"`
		Name       string                    `json:"name"`
	} `json:"asset"`
	Task struct {
		ID string `json:"id"`
	} `json:"task"`
}

type LivePeerAssetStatusResponse struct {
	ID   string `json:"id"`
	Hash []struct {
		Hash      string `json:"hash"`
		Algorithm string `json:"algorithm"`
	} `json:"hash"`
	Name      string                    `json:"name"`
	Size      uint                      `json:"size"`
	Status    LivePeerAssetUploadStatus `json:"status"`
	UserID    string                    `json:"userId"`
	CreatedAt uint                      `json:"createdAt"`
	//UpdatedAt uint   `json:"updatedAt"`
	VideoSpec struct {
		Format   string  `json:"format"`
		Duration float32 `json:"duration"`
	} `json:"videoSpec"`
	PlaybackID  string `json:"playbackId"`
	PlaybackUrl string `json:"playbackUrl"`
	DownloadUrl string `json:"downloadUrl"`
}

type VideoAdditionalProps struct {
	// Required by frontend
	Height uint `json:"height"`
	Width  uint `json:"width"`

	// Additional props
	Work   LivePeerUploadURLResponse   `json:"work"`
	Status LivePeerAssetStatusResponse `json:"status"`
}

func UploadURIToLivePeer(url string, name string) (string, uint, string, string, error) {
	// Docs https://docs.livepeer.studio/guides/on-demand#upload-with-url

	global.Logger.Debugf("Uploading file %s to LivePeer with url %s", name, url)

	reqBody := LivePeerUploadURLRequest{
		URL:  url,
		Name: name,
	}
	reqBodyBytes, err := json.Marshal(&reqBody)
	if err != nil {
		global.Logger.Errorf("Failed to parse media upload request for URI %s with error: %s", url, err.Error())
		return "", 0, "", "", err
	}
	req, err := http.NewRequest("POST", "https://livepeer.studio/api/asset/import", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		global.Logger.Errorf("Failed to crete request with error: %s", err.Error())
		return "", 0, "", "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Config.LivePeerAPIKey))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		global.Logger.Errorf("Failed to do http request with error: %s", err.Error())
		return "", 0, "", "", err
	}
	var uploadTask LivePeerUploadURLResponse
	if err := json.NewDecoder(res.Body).Decode(&uploadTask); err != nil {
		global.Logger.Errorf("Failed to parse http response result with error: %s", err.Error())
		return "", 0, "", "", err
	}

	var uploadTaskStatus *LivePeerAssetStatusResponse

	// Wait till asset upload succeed
	// TODO: Use a check query to optimise

	for {
		// Let's sleep for 30 seconds
		global.Logger.Debugf("Waiting for 30 seconds till next check (asset id: %s)...", uploadTask.Asset.ID)
		time.Sleep(30 * time.Second)
		global.Logger.Debugf("Start check upload task status (asset id: %s)", uploadTask.Asset.ID)
		uploadTaskStatus, err = CheckUploadStatus(uploadTask.Asset.ID)
		if err != nil {
			// Failed to get data
			global.Logger.Errorf("Failed to check video (%s) upload status with error: %s", uploadTask.Asset.ID, err.Error())
			continue // Try again later
		}
		global.Logger.Debugf("Video upload (asset id: %s) task tatus: %s", uploadTask.Asset.ID, uploadTaskStatus.Status.Phase)

		if uploadTaskStatus.Status.Phase == "ready" {
			// Ready to publish
			global.Logger.Debugf("Video is ready (asset id: %s)", uploadTask.Asset.ID)
			break
		} else if uploadTaskStatus.Status.Phase == "failed" {
			// Upload failed, need to report back
			global.Logger.Errorf("Failed to upload video %s %v to LivePeer with error: %s", url, uploadTask, uploadTaskStatus.Status.ErrorMessage)
			return "", 0, "", "", fmt.Errorf(uploadTaskStatus.Status.ErrorMessage)
		}
	}

	additionalProps := VideoAdditionalProps{
		Work:   uploadTask,
		Status: *uploadTaskStatus,
	}

	additionalPropsBytes, err := json.Marshal(&additionalProps)
	if err != nil {
		global.Logger.Errorf("Failed to parse additional props into bytes with error: %s", err.Error())
	}

	return uploadTaskStatus.PlaybackUrl, uploadTaskStatus.Size, fmt.Sprintf("video/%s", uploadTaskStatus.VideoSpec.Format), string(additionalPropsBytes), nil
}

func CheckUploadStatus(assetID string) (*LivePeerAssetStatusResponse, error) {
	// Docs https://docs.livepeer.studio/guides/on-demand#verify-that-the-status-is-ready

	req, err := http.NewRequest("GET", fmt.Sprintf("https://livepeer.studio/api/asset/%s/", assetID), nil)
	if err != nil {
		global.Logger.Errorf("Failed to crete request with error: %s", err.Error())
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Config.LivePeerAPIKey))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		global.Logger.Errorf("Failed to do http request with error: %s", err.Error())
		return nil, err
	}

	var uploadStatus LivePeerAssetStatusResponse
	if err := json.NewDecoder(res.Body).Decode(&uploadStatus); err != nil {
		global.Logger.Errorf("Failed to parse http response result with error: %s", err.Error())
		return nil, err
	}

	return &uploadStatus, nil

}
