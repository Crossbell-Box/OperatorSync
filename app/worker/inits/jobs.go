package inits

import "github.com/Crossbell-Box/OperatorSync/app/worker/jobs"

func Jobs() error {
	if err := jobs.FeedCollectStartProcess(); err != nil {
		return err
	}
	if err := jobs.AccountValidateStartProcess(); err != nil {
		return err
	}

	if err := jobs.OnChainFeedsStartProcess(); err != nil {
		return err
	}

	return nil
}
