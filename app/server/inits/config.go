package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"log"
	"os"
	"strings"
)

func Config() error {
	// Get from environment variables

	var exist bool // If non-exist use default

	if config.Config.DBConnString, exist = os.LookupEnv("DATABASE_CONNECTION_STRING"); !exist {
		config.Config.DBConnString = "postgres://postgres:postgres@localhost:5432/postgres"
	}
	if config.Config.RedisConnString, exist = os.LookupEnv("REDIS_CONNECTION_STRING"); !exist {
		config.Config.RedisConnString = "redis://localhost:6379/0"
	}
	if config.Config.MQConnString, exist = os.LookupEnv("MQ_CONNECTION_STRING"); !exist {
		config.Config.MQConnString = "nats://localhost:4222"
	}
	config.Config.IsMainServer = strings.Contains(strings.ToLower(os.Getenv("MAIN_SERVER")), "t")

	config.Config.DevelopmentMode = !strings.Contains(strings.ToLower(os.Getenv("MODE")), "prod")

	if config.Config.DevelopmentMode {
		log.Println("Configurations: ", config.Config)
	}

	return nil
}
