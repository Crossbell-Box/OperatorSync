package utils

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"io/ioutil"
	"net/http"
)

func HttpRequest(url string, withProxy bool) ([]byte, error) {

	var tr http.Transport
	if withProxy && config.Config.ProxyURL != nil {
		tr.Proxy = http.ProxyURL(config.Config.ProxyURL)
	}

	client := &http.Client{
		Transport: &tr,
	}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	_ = res.Body.Close() // Ignore error

	return body, nil
}
