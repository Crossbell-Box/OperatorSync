package main

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/inits"
	"log"
)

func main() {
	// Initialize everything
	log.Println("Initializing...")

	// Initialize config
	if err := inits.Config(); err != nil {
		log.Fatalln("Failed to load config: ", err)
	}

	// Initialize logger
	if err := inits.Logger(); err != nil {
		log.Fatalln("Failed to load logger: ", err)
	}

	global.Logger.Info("Logger initialized, switch to here.")

	// Initialize MQ
	if err := inits.MQ(); err != nil {
		global.Logger.Fatal("Failed to load MQ: ", err.Error())
	}

	// Initialize jobs
	if err := inits.Jobs(); err != nil {
		global.Logger.Fatal("Failed to start jobs: ", err.Error())
	}

	global.Logger.Info("Worker started!")

	select {} // Keep process running

}
