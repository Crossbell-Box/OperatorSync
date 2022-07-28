package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
)

type response struct {
	Status  string `json:"status"`
	CID     string `json:"cid"`
	URL     string `json:"url"`
	Web2URL string `json:"web2url"`
	Error   string `json:"error"`
}

func UploadURLToIPFS(targetUrl string) (string, uint, string, error) {
	// Get filename
	global.Logger.Debug("Uploading file ", targetUrl, " to IPFS...")

	reqUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", 0, "", err
	}
	filename := path.Base(reqUrl.Path)

	// Retrieve data
	body, err := HttpRequest(targetUrl, false)
	if err != nil {
		global.Logger.Error("Failed to retrieve data from: ", targetUrl)
		return "", 0, "", err
	}

	global.Logger.Debug("File download successfully, uploading...")

	// Prepare
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return "", 0, "", err
	}

	bodyBytes := body[:]
	fileSize, err := fileWriter.Write(bodyBytes)
	if err != nil {
		global.Logger.Error("Failed to copy buffer data")
	}

	contentType := http.DetectContentType(bodyBytes)

	// Upload to proxy
	var resp response
	formContentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close() // Ignore error
	ipfsReq, err := http.NewRequest("POST", config.Config.IPFSEndpoint, bodyBuffer)
	if err != nil {
		global.Logger.Error("Failed to initialize request: ", err.Error())
		return "", 0, "", err
	}
	ipfsReq.Header.Set("Content-Type", formContentType)
	ipfsRes, err := (&http.Client{}).Do(ipfsReq)
	if err != nil {
		global.Logger.Error("Failed to do request: ", err.Error())
		return "", 0, "", err
	}
	err = json.NewDecoder(ipfsRes.Body).Decode(&resp)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return "", 0, "", err
	}

	// Return URL
	if resp.Status == "ok" {
		global.Logger.Debug("File (Size: ", fileSize, ") upload succeeded: ", resp.URL)
		return resp.URL, uint(fileSize), contentType, nil
	} else {
		return "", 0, "", fmt.Errorf(resp.Error)
	}

}
