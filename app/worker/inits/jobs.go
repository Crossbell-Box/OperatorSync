package inits

import (
	"github.com/Crossbell-Box/OperatorSync/app/worker/mq/jobs"
)

func Jobs() error {
	if err := jobs.FeedCollectStartProcess(); err != nil {
		return err
	}

	return nil
}
