package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"image"
	_ "image/gif"  // Add GIF support
	_ "image/jpeg" // Add JPEG support
	_ "image/png"  // Add PNG support
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
)

type response struct {
	Status  string `json:"status"`
	CID     string `json:"cid"`
	URL     string `json:"url"`
	Web2URL string `json:"web2url"`
	Error   string `json:"error"`
}

func UploadURLToIPFS(targetUrl string, withProxy bool) (string, string, uint, string, string, error) {
	// Get filename
	global.Logger.Debug("Uploading file ", targetUrl, " to IPFS...")

	if targetUrl == "" {
		return "", "", 0, "", "", fmt.Errorf("empty uri")
	}

	reqUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", "", 0, "", "", err
	}
	filename := path.Base(reqUrl.Path)

	// Retrieve data
	body, err := HttpRequest(targetUrl, withProxy)
	if err != nil {
		global.Logger.Error("Failed to retrieve data from: ", targetUrl)
		return "", "", 0, "", "", err
	}

	bodyBytes := body[:]

	// Detect content-type
	contentType := http.DetectContentType(bodyBytes)

	// Upload to IPFS
	global.Logger.Debug("File download successfully, uploading...")

	ipfsUri, fileSize, err := UploadBytesToIPFS(bodyBytes, filename)
	if err != nil {
		global.Logger.Errorf("Failed to upload data to IPFS with error: %s", err.Error())
		return "", "", 0, "", "", err
	}

	// Attach additional props
	additionalProps := make(map[string]string)

	if strings.Contains(contentType, "image/") {
		// Is image
		imgConfig, format, err := image.DecodeConfig(bytes.NewReader(bodyBytes))
		if err != nil {
			// Unable to handle this
			global.Logger.Errorf("Failed to decode image with error: %s", err.Error())
		} else {
			additionalProps["format"] = format
			additionalProps["width"] = strconv.Itoa(imgConfig.Width)
			additionalProps["height"] = strconv.Itoa(imgConfig.Height)
		}

	}

	additionalPropsBytes, err := json.Marshal(&additionalProps)
	if err != nil {
		global.Logger.Errorf("Failed to parse additional props with error: %s", err.Error())
		additionalPropsBytes = nil
	}

	// Return response
	return filename, ipfsUri, fileSize, contentType, string(additionalPropsBytes), nil

}

func UploadBytesToIPFS(data []byte, filename string) (string, uint, error) {
	// Prepare
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return "", 0, err
	}

	fileSize, err := fileWriter.Write(data)
	if err != nil {
		global.Logger.Error("Failed to copy buffer data")
	}

	// Upload to proxy
	var resp response
	formContentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close() // Ignore error
	ipfsReq, err := http.NewRequest("POST", config.Config.IPFSEndpoint, bodyBuffer)
	if err != nil {
		global.Logger.Error("Failed to initialize request: ", err.Error())
		return "", 0, err
	}
	ipfsReq.Header.Set("Content-Type", formContentType)
	ipfsRes, err := (&http.Client{}).Do(ipfsReq)
	if err != nil {
		global.Logger.Error("Failed to do request: ", err.Error())
		return "", 0, err
	}
	err = json.NewDecoder(ipfsRes.Body).Decode(&resp)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return "", 0, err
	}

	// Return URL
	if resp.Status == "ok" {
		global.Logger.Debug("File (Size: ", fileSize, ") upload succeeded: ", resp.URL)
		return resp.URL, uint(fileSize), nil
	} else {
		return "", 0, fmt.Errorf(resp.Error)
	}
}
