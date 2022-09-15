package config

type Config struct {
	MQConnString    string // RabbitMQ
	WorkerRPCPort   string // Internal RPC port, default 2915
	DevelopmentMode bool
}
