package types

import "time"

type WorkDispatched struct {
	DispatchAt  time.Time `json:"dispatch_at"`
	AccountID   uint      `json:"account_id"`
	Platform    string    `json:"platform"`
	CollectLink string    `json:"collect_link"` // HTTP(S) link to the RSS endpoint
	DropBefore  time.Time `json:"drop_before"`  // Ignore feeds before last updated time
	Deadline    time.Time `yaml:"ddl"`          // If works cannot be accepted than this time, it fails
}
