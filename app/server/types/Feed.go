package types

import (
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

type Feed struct {
	// Structure related
	AccountID uint `gorm:"index"`

	// Feed Metadata
	Platform    string    `json:"platform"`
	CollectedAt time.Time `json:"collected_at"`

	// Raw feed
	commonTypes.RawFeed
}
