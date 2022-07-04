package types

import (
	"github.com/lib/pq"
	"time"
)

// RawFeed : Minimal content of feed (just for work response)
type RawFeed struct {
	Language    string         `json:"language"`
	Title       string         `json:"title"`
	Text        string         `json:"text"`
	PublishedAt time.Time      `json:"published_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Media       pq.StringArray `json:"media"` // Videos / Audios / Images, upload to IPFS
}

type WorkSucceeded struct {
	WorkDispatched

	AcceptedAt  time.Time `json:"accepted_at"`
	SucceededAt time.Time `json:"succeeded_at"`

	Feeds []RawFeed `json:"feeds"`
}
