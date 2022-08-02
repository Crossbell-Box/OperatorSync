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

	// Concurrency control
	ConcurrencyStateful  int
	ConcurrencyStateless int
	ConcurrencyDirect    int

	// Crossbell chain related
	CrossbellChainID   int64
	CrossbellJsonRPC   string
	EthereumPrivateKey string

	commonConfig.Config
}

var Config workerConfig
