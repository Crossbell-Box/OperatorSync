package config

import (
	commonConfig "github.com/Crossbell-Box/OperatorSync/common/config"
	"net/url"
)

type workerConfig struct {
	RSSHubEndpointStateful  string
	RSSHubEndpointStateless string

	ProxyURL *url.URL

	IPFSEndpoint string

	commonConfig.Config
}

var Config workerConfig
