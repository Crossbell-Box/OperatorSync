package config

import commonConfig "github.com/Crossbell-Box/OperatorSync/common/config"

type serverConfig struct {
	DBConnString    string // Postgres Database
	RedisConnString string // Redis

	WorkerRPCEndpoint string // Worker

	IsMainServer bool // Works on cluster mode

	commonConfig.Config
}

var Config serverConfig
