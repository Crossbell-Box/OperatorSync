package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/jobs"
	"time"
)

func Jobs() error {

	// Start listening to FeedCollect works response
	if err := jobs.FeedCollectStartRetrieveWork(); err != nil {
		return err
	}

	if config.Config.IsMainServer {
		// Start dispatch flush works
		config.Status.Jobs.FeedCollectLastRun = time.Now()
		jobs.FeedCollectStartDispatchWork()
		config.Status.Jobs.ResumePausedAccountsLastRun = time.Now()
		jobs.ResumePausedAccounts()
	}

	return nil

}
