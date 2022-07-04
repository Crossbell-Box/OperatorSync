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

func UploadURLToIPFS(targetUrl string) (string, error) {
	// Get filename
	reqUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}
	filename := path.Base(reqUrl.Path)

	// Retrieve data
	body, err := HttpRequest(targetUrl, false)
	if err != nil {
		global.Logger.Error("Failed to retrieve data from: ", targetUrl)
		return "", err
	}

	// Prepare
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return "", err
	}

	_, err = fileWriter.Write(body)
	if err != nil {
		global.Logger.Error("Failed to copy buffer data")
	}

	// Upload to proxy
	var resp response
	contentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close() // Ignore error
	ipfsReq, _ := http.NewRequest("POST", config.Config.IPFSEndpoint, bodyBuffer)
	ipfsReq.Header.Set("Content-Type", contentType)
	ipfsRes, _ := (&http.Client{}).Do(ipfsReq)
	err = json.NewDecoder(ipfsRes.Body).Decode(&resp)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return "", err
	}

	// Return URL
	if resp.Status == "ok" {
		return resp.URL, nil
	} else {
		return "", fmt.Errorf(resp.Error)
	}

}
