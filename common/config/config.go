package config

type Config struct {
	MQConnString    string // RabbitMQ
	RedisConnString string // Redis
	WorkerRPCPort   string // Internal RPC port, default 22915
	DevelopmentMode bool
}
