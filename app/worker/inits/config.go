package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/consts"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"log"
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
		log.Println("Http proxy URL not set, skip proxy")
		config.Config.ProxyURL = nil
	} else if config.Config.ProxyURL, err = url.Parse(rawProxyURL); err != nil {
		log.Println("Invalid http proxy URL setting, skip proxy")
		config.Config.ProxyURL = nil
	}

	if config.Config.IPFSEndpoint, exist = os.LookupEnv("IPFS_ENDPOINT"); !exist {
		return fmt.Errorf("please specify endpoint URI for your IPFS-Upload-Relay instance (https://github.com/NaturalSelectionLabs/IPFS-Upload-Relay)")
	}

	if config.Config.RedisConnString, exist = os.LookupEnv("REDIS_CONNECTION_STRING"); !exist {
		config.Config.RedisConnString = commonConsts.CONFIG_DEFAULT_REDIS_CONNECTION_STRING
	}
	if config.Config.MQConnString, exist = os.LookupEnv("MQ_CONNECTION_STRING"); !exist {
		config.Config.MQConnString = commonConsts.CONFIG_DEFAULT_MQ_CONNECTION_STRING
	}
	if config.Config.WorkerRPCPort, exist = os.LookupEnv("WORKER_RPC_PORT"); !exist {
		config.Config.WorkerRPCPort = commonConsts.CONFIG_DEFAULT_WORKER_RPC_PORT
	}

	if concurrencyStatefulStr, exist := os.LookupEnv("CONCURRENCY_CONTROL_STATEFUL"); !exist {
		config.Config.ConcurrencyStateful = consts.CONFIG_DEFAULT_CONCURRENCY_CONTROL_STATEFUL // Default
	} else if config.Config.ConcurrencyStateful, err = strconv.Atoi(concurrencyStatefulStr); err != nil || config.Config.ConcurrencyStateful <= 0 {
		log.Println("Invalid stateful concurrency control settings, using default value")
		config.Config.ConcurrencyStateful = consts.CONFIG_DEFAULT_CONCURRENCY_CONTROL_STATEFUL // Default
	}
	if concurrencyStatelessStr, exist := os.LookupEnv("CONCURRENCY_CONTROL_STATELESS"); !exist {
		config.Config.ConcurrencyStateless = consts.CONFIG_DEFAULT_CONCURRENCY_CONTROL_STATELESS // Default
	} else if config.Config.ConcurrencyStateless, err = strconv.Atoi(concurrencyStatelessStr); err != nil || config.Config.ConcurrencyStateless <= 0 {
		log.Println("Invalid stateless concurrency control settings, using default value")
		config.Config.ConcurrencyStateless = consts.CONFIG_DEFAULT_CONCURRENCY_CONTROL_STATELESS // Default
	}
	if concurrencyDirectStr, exist := os.LookupEnv("CONCURRENCY_CONTROL_DIRECT"); !exist {
		config.Config.ConcurrencyDirect = consts.CONFIG_DEFAULT_CONCURRENCY_CONTROL_DIRECT // Default
	} else if config.Config.ConcurrencyDirect, err = strconv.Atoi(concurrencyDirectStr); err != nil || config.Config.ConcurrencyDirect <= 0 {
		log.Println("Invalid direct concurrency control settings, using default value")
		config.Config.ConcurrencyDirect = consts.CONFIG_DEFAULT_CONCURRENCY_CONTROL_DIRECT // Default
	}

	config.Config.DevelopmentMode = !strings.Contains(strings.ToLower(os.Getenv("MODE")), "prod")

	if crossbellChainIDStr, exist := os.LookupEnv("CROSSBELL_CHAIN_ID"); !exist {
		config.Config.CrossbellChainID = consts.CONFIG_DEFAULT_CROSSBELL_CHAIN_ID // Default
	} else if config.Config.CrossbellChainID, err = strconv.ParseInt(crossbellChainIDStr, 10, 64); err != nil || config.Config.ConcurrencyDirect <= 0 {
		log.Println("Invalid Crossbell chain ID settings, using default value (3737)")
		config.Config.CrossbellChainID = consts.CONFIG_DEFAULT_CROSSBELL_CHAIN_ID // Default
	}

	if config.Config.CrossbellJsonRPC, exist = os.LookupEnv("CROSSBELL_JSON_RPC"); !exist {
		log.Println("No Crossbell JSON RPC endpoint found, use default")
		config.Config.CrossbellJsonRPC = consts.CONFIG_DEFAULT_CROSSBELL_JSON_RPC // Default
	}

	if config.Config.CrossbellIndexer, exist = os.LookupEnv("CROSSBELL_INDEXER"); !exist {
		log.Println("No Crossbell Indexer endpoint found, use default")
		config.Config.CrossbellIndexer = consts.CONFIG_DEFAULT_CROSSBELL_INDEXER // Default
	}

	if config.Config.CrossbellContractAddress, exist = os.LookupEnv("CROSSBELL_CONTRACT_ADDRESS"); !exist {
		log.Println("No Crossbell contract address found, use default")
		config.Config.CrossbellContractAddress = consts.CONFIG_DEFAULT_CROSSBELL_CONTRACT_ADDRESS // Default
	}

	if config.Config.EthereumPrivateKey, exist = os.LookupEnv("ETHEREUM_PRIVATE_KEY"); !exist {
		return fmt.Errorf("please specify ethereum private key for on-chain requests")
	}

	// Print configurations
	if config.Config.DevelopmentMode {
		log.Println("Configurations: ", config.Config)
	}

	return nil

}
