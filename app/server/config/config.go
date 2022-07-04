package config

type serverConfig struct {
	DBConnString    string // Postgres Database
	RedisConnString string // Redis
	MQConnString    string // NATS MQ
	DevelopmentMode bool
}

var Config serverConfig
