package config

import "time"

type serverStatus struct {
	Jobs struct {
		FeedCollectLastRun          time.Time
		ResumePausedAccountsLastRun time.Time
	}
}

var Status serverStatus
