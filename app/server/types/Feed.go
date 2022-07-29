package types

import (
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

type Feed struct {
	// Structure related
	AccountID uint `gorm:"index"`

	// Feed Metadata
	Platform    string    `json:"platform" gorm:"-"` // For table splitting
	CollectedAt time.Time `json:"collected_at"`

	// Raw feed
	commonTypes.RawFeed
}

type OnChainData struct {
	IPFSUri     string `json:"ipfs_uri,omitempty"`
	Transaction string `json:"transaction,omitempty"`
}
