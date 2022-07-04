package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"net/url"
	"os"
	"strings"
)

func Config() error {
	var exist bool
	var err error

	if config.Config.RSSHubEndpointStateful, exist = os.LookupEnv("RSSHUB_STATEFUL"); !exist {
		return fmt.Errorf("please specify endpoint URI for stateful RSSHub (https://rsshub.app)")
	}
	if config.Config.RSSHubEndpointStateless, exist = os.LookupEnv("RSSHUB_STATELESS"); !exist {
		return fmt.Errorf("please specify endpoint URI for stateless RSSHub (https://rsshub.app)")
	}

	if rawProxyURL, exist := os.LookupEnv("PROXY_URL"); !exist {
		return fmt.Errorf("please specify http proxy URL")
	} else if config.Config.ProxyURL, err = url.Parse(rawProxyURL); err != nil {
		return fmt.Errorf("http proxy URL is invalid")
	}

	if config.Config.IPFSEndpoint, exist = os.LookupEnv("IPFS_ENDPOINT"); !exist {
		config.Config.IPFSEndpoint = "https://ipfs-relay.crossbell.io/upload"
	}
	if config.Config.MQConnString, exist = os.LookupEnv("MQ_CONNECTION_STRING"); !exist {
		config.Config.MQConnString = "nats://localhost:4222"
	}

	config.Config.DevelopmentMode = !strings.Contains(strings.ToLower(os.Getenv("MODE")), "prod")

	return nil

}
