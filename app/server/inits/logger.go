package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"go.uber.org/zap"
)

func Logger() error {
	var err error

	var logger *zap.Logger

	// Prepare logger
	if config.Config.DevelopmentMode {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	// Flush logs
	defer logger.Sync() // Unable to handle errors here

	// Sugar it
	global.Logger = logger.Sugar()

	return nil
}
