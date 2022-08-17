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

type LivePeerAssetStatusResponse struct {
	ID   string `json:"id"`
	Hash []struct {
		Hash      string `json:"hash"`
		Algorithm string `json:"algorithm"`
	} `json:"hash"`
	Name      string `json:"name"`
	Size      uint   `json:"size"`
	Status    string `json:"status"` // "ready"
	UserID    string `json:"userId"`
	CreatedAt uint   `json:"createdAt"`
	UpdatedAt uint   `json:"updatedAt"`
	VideoSpec struct {
		Format string `json:"format"`
		Tracks []struct {
			Type        string  `json:"type"` // "video" / "audio“
			Codec       string  `json:"codec"`
			Bitrate     uint    `json:"bitrate"`
			Duration    float32 `json:"duration"`
			FPS         float32 `json:"fps,omitempty"`         // Video only
			Width       uint    `json:"width,omitempty"`       // Video only
			Height      uint    `json:"height,omitempty"`      // Video only
			PixelFormat string  `json:"pixelFormat,omitempty"` // Video only
			Channels    uint    `json:"channels,omitempty"`    // Audio only
		} `json:"tracks"`
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
	var uploadWork LivePeerUploadURLResponse
	if err := json.NewDecoder(res.Body).Decode(&uploadWork); err != nil {
		global.Logger.Errorf("Failed to parse http response result with error: %s", err.Error())
		return "", 0, "", "", err
	}

	var uploadStatus *LivePeerAssetStatusResponse

	// Wait till asset upload succeed
	// TODO: Use a check query to optimise

	for {
		// Let's sleep for 30 seconds
		time.Sleep(30 * time.Second)
		uploadStatus, err = CheckUploadStatus(uploadWork.Asset.ID)
		if err != nil {
			// Failed to get data
			global.Logger.Errorf("Failed to check video (%s) upload status with error: %s", uploadWork.Asset.ID, err.Error())
			continue // Try again later
		}

		if uploadStatus.Status == "ready" {
			// Ready to publish
			break
		}
	}

	additionalProps := VideoAdditionalProps{
		Work:   uploadWork,
		Status: *uploadStatus,
	}

	// Find the video track and fill the width / height props
	for _, track := range uploadStatus.VideoSpec.Tracks {
		if track.Type == "video" {
			additionalProps.Height = track.Height
			additionalProps.Width = track.Width
			break
		}
	}

	additionalPropsBytes, err := json.Marshal(&additionalProps)
	if err != nil {
		global.Logger.Errorf("Failed to parse additional props into bytes with error: %s", err.Error())
	}

	return uploadStatus.PlaybackUrl, uploadStatus.Size, fmt.Sprintf("video/%s", uploadStatus.VideoSpec.Format), string(additionalPropsBytes), nil
}

func CheckUploadStatus(assetID string) (*LivePeerAssetStatusResponse, error) {
	// Docs https://docs.livepeer.studio/guides/on-demand#verify-that-the-status-is-ready

	req, err := http.NewRequest("GET", fmt.Sprintf("https://livepeer.studio/api/asset/%s/", assetID), nil)
	if err != nil {
		global.Logger.Errorf("Failed to crete request with error: %s", err.Error())
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Config.LivePeerAPIKey))

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
