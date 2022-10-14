package main

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/inits"
	commonInits "github.com/Crossbell-Box/OperatorSync/common/inits"
	"log"
)

func main() {
	// Initialize everything
	log.Println("Initializing...")

	// Initialize config
	if err := inits.Config(); err != nil {
		log.Fatalln("Failed to load config:", err)
	}

	// Initialize logger
	if err := inits.Logger(); err != nil {
		log.Fatalln("Failed to load logger:", err)
	}

	global.Logger.Info("Logger initialized, switch to here.")

	// Initialize database
	if err := inits.DB(); err != nil {
		global.Logger.Fatal("Failed to load database: ", err.Error())
	}

	// Initialize redis
	if err := commonInits.Redis(config.Config.RedisConnString); err != nil {
		global.Logger.Fatal("Failed to load redis: ", err.Error())
	}

	// Initialize MQ
	if err := commonInits.MQ(config.Config.MQConnString, global.Logger); err != nil {
		global.Logger.Fatal("Failed to load MQ: ", err.Error())
	}

	// Initialize RPC Client
	if err := inits.RPC(); err != nil {
		global.Logger.Fatal("Failed to connect to worker RPC: ", err.Error())
	}

	// Initialize metrics
	if err := inits.Metrics(); err != nil {
		global.Logger.Fatal("Failed to load metrics: ", err.Error())
	}

	// Initialize jobs
	if err := inits.Jobs(); err != nil {
		global.Logger.Fatal("Failed to start jobs: ", err.Error())
	}

	// Initializing router
	engine := inits.Routes()

	global.Logger.Info("Initialization complete.")

	// Start
	global.Logger.Info("Service starting...")
	if err := engine.Run(); err != nil {
		global.Logger.Fatal("Failed to start service:", err.Error())
	}

}
