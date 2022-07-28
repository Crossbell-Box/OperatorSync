package inits

import "github.com/Crossbell-Box/OperatorSync/app/worker/jobs"

func Jobs() error {
	if err := jobs.StartProcessFeeds(); err != nil {
		return err
	}
	if err := jobs.StartValidateAccounts(); err != nil {
		return err
	}

	return nil
}
