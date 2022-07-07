package types

import "time"

type WorkDispatched struct {
	DispatchAt  time.Time `json:"dispatch_at"`
	AccountID   uint      `json:"account_id"`
	VerifyKey   string    `json:"verify_key"`
	Platform    string    `json:"platform"`
	CollectLink string    `json:"collect_link"` // HTTP(S) link to the RSS endpoint
	DropBefore  time.Time `json:"drop_before"`  // Ignore feeds before last updated time
	DropAfter   time.Time `json:"drop_after"`   // Ignore feeds after next updated time
}
