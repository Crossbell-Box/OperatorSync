package inits

import (
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DB() error {
	var err error
	var gormConfig gorm.Config

	// Set config
	if config.Config.DevelopmentMode {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.Logger = logger.Default.LogMode(logger.Error)
	}

	// Connect to database
	global.DB, err = gorm.Open(postgres.Open(config.Config.DBConnString), &gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run migrations
	err = mig()
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func mig() error {
	err := global.DB.AutoMigrate(
		&models.Account{},
	)
	if err != nil {
		return err
	}

	for platformID, _ := range commonConsts.SUPPORTED_PLATFORM {
		feedWithPlatform := models.Feed{
			Feed: types.Feed{
				Platform: platformID, // Cannot be initialized with platform specified
			},
		}

		if err = global.DB.
			Scopes(models.FeedTable(feedWithPlatform)).
			AutoMigrate(&feedWithPlatform); err != nil {
			return err
		}
	}
	return nil
}
