package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/jobs"
)

func Jobs() error {

	// Start listening to FeedCollect works response
	if err := jobs.FeedCollectStartRetrieveWork(); err != nil {
		return err
	}

	if config.Config.IsMainServer {
		// Start dispatch flush works
		jobs.FeedCollectStartDispatchWork()
		jobs.ResumePausedAccounts()
	}

	return nil

}
