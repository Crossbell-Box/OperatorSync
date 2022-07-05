package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"net/url"
	"os"
	"strconv"
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

	if concurrencyStatefulStr, exist := os.LookupEnv("CONCURRENCY_CONTROL_STATEFUL"); !exist {
		config.Config.ConcurrencyStateful = 10 // Default
	} else if config.Config.ConcurrencyStateful, err = strconv.Atoi(concurrencyStatefulStr); err != nil || config.Config.ConcurrencyStateful <= 0 {
		global.Logger.Error("Invalid stateful concurrency control settings, using default value")
		config.Config.ConcurrencyStateful = 10 // Default
	}
	if concurrencyStatelessStr, exist := os.LookupEnv("CONCURRENCY_CONTROL_STATELESS"); !exist {
		config.Config.ConcurrencyStateless = 50 // Default
	} else if config.Config.ConcurrencyStateless, err = strconv.Atoi(concurrencyStatelessStr); err != nil || config.Config.ConcurrencyStateless <= 0 {
		global.Logger.Error("Invalid stateless concurrency control settings, using default value")
		config.Config.ConcurrencyStateless = 50 // Default
	}
	if concurrencyDirectStr, exist := os.LookupEnv("CONCURRENCY_CONTROL_DIRECT"); !exist {
		config.Config.ConcurrencyDirect = 100 // Default
	} else if config.Config.ConcurrencyDirect, err = strconv.Atoi(concurrencyDirectStr); err != nil || config.Config.ConcurrencyDirect <= 0 {
		global.Logger.Error("Invalid direct concurrency control settings, using default value")
		config.Config.ConcurrencyDirect = 100 // Default
	}

	config.Config.DevelopmentMode = !strings.Contains(strings.ToLower(os.Getenv("MODE")), "prod")

	return nil

}
