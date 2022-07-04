package types

import (
	"time"
)

type WorkSucceeded struct {
	WorkDispatched

	AcceptedAt  time.Time `json:"accepted_at"`
	SucceededAt time.Time `json:"succeeded_at"`

	Feeds []RawFeed `json:"feeds"`
}
