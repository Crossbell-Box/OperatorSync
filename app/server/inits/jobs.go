package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/config"
	"github.com/Crossbell-Box/OperatorSync/app/server/jobs"
)

func Jobs() error {

	// Start listening to response
	if err := jobs.ReceiveSucceededWork(); err != nil {
		return err
	}
	if err := jobs.ReceiveFailedWork(); err != nil {
		return err
	}

	if !config.Config.DisableFlushWorkDispatch {
		// Start dispatch flush works
		jobs.StartDispatchFeedCollectWork()
	}

	return nil

}
