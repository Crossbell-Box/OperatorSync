package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/jobs"
)

func Jobs() error {

	// Start listening to FeedCollect works response
	if err := jobs.FeedCollectStartReceiveSucceededWork(); err != nil {
		return err
	}
	if err := jobs.FeedCollectStartReceiveFailedWork(); err != nil {
		return err
	}

	if config.Config.IsMainServer {
		// Start dispatch flush works
		jobs.FeedCollectStartDispatchWork()
	}

	return nil

}
