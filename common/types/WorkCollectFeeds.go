package types

import (
	"time"
)

type WorkDispatched struct {
	DispatchAt time.Time `json:"dispatch_at"`
	AccountID  uint      `json:"account_id"`
	Platform   string    `json:"platform"`
	Username   string    `json:"username"`    // The unique identifier on platform
	DropBefore time.Time `json:"drop_before"` // Ignore feeds before last updated time
	DropAfter  time.Time `json:"drop_after"`  // Ignore feeds after next updated time
}

type WorkSucceeded struct {
	WorkDispatched

	AcceptedAt  time.Time     `json:"accepted_at"`
	SucceededAt time.Time     `json:"succeeded_at"`
	NewInterval time.Duration `json:"new_interval"`

	Feeds []RawFeed `json:"feeds"`
}

type WorkFailed struct {
	WorkDispatched

	AcceptAt    time.Time `json:"accept_at"`
	ErrorAt     time.Time `json:"error_at"`
	ErrorCode   uint      `json:"error_code"`
	ErrorReason string    `json:"error_reason"`
}
