package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/consts"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"log"
	"os"
	"strings"
)

func Config() error {
	// Get from environment variables

	var exist bool // If non-exist use default

	if config.Config.DBConnString, exist = os.LookupEnv("DATABASE_CONNECTION_STRING"); !exist {
		config.Config.DBConnString = consts.CONFIG_DEFAULT_DATABASE_CONNECTION_STRING
	}
	if config.Config.RedisConnString, exist = os.LookupEnv("REDIS_CONNECTION_STRING"); !exist {
		config.Config.RedisConnString = commonConsts.CONFIG_DEFAULT_REDIS_CONNECTION_STRING
	}
	if config.Config.MQConnString, exist = os.LookupEnv("MQ_CONNECTION_STRING"); !exist {
		config.Config.MQConnString = commonConsts.CONFIG_DEFAULT_MQ_CONNECTION_STRING
	}
	if config.Config.WorkerRPCEndpoint, exist = os.LookupEnv("WORKER_RPC_ENDPOINT"); !exist {
		config.Config.WorkerRPCEndpoint = consts.CONFIG_DEFAULT_WORKER_RPC_ENDPOINT
	}
	if config.Config.WorkerRPCPort, exist = os.LookupEnv("WORKER_RPC_PORT"); !exist {
		config.Config.WorkerRPCPort = commonConsts.CONFIG_DEFAULT_WORKER_RPC_PORT
	}
	config.Config.IsMainServer = strings.Contains(strings.ToLower(os.Getenv("MAIN_SERVER")), "t")

	if config.Config.IsMainServer {
		// Check jobs webhook
		if config.Config.HeartBeatWebhooks.FeedCollect, exist = os.LookupEnv("HEARTBEAT_WEBHOOK_FEED_COLLECT");
			!exist || !utils.ValidateUri(config.Config.HeartBeatWebhooks.FeedCollect) {
			config.Config.HeartBeatWebhooks.FeedCollect = "" // Nope
		}
		if config.Config.HeartBeatWebhooks.AccountResume, exist = os.LookupEnv("HEARTBEAT_WEBHOOK_ACCOUNT_RESUME");
			!exist || !utils.ValidateUri(config.Config.HeartBeatWebhooks.AccountResume) {
			config.Config.HeartBeatWebhooks.AccountResume = "" // Nope
		}
	}

	config.Config.DevelopmentMode = !strings.Contains(strings.ToLower(os.Getenv("MODE")), "prod")

	if config.Config.DevelopmentMode {
		log.Println("Configurations: ", config.Config)
	}

	return nil
}
