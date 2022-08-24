package types

import (
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/lib/pq"
	"time"
)

type Feed struct {
	// Structure related
	AccountID uint `json:"-" gorm:"index"`

	// Feed Metadata
	Platform    string    `json:"platform" gorm:"-"` // For table splitting
	CollectedAt time.Time `json:"collected_at"`

	// Related Media
	MediaIPFSUris pq.StringArray `json:"media_ipfs_uris" gorm:"type:text[];column:media_ipfs_uris"`

	// Raw feed
	commonTypes.RawFeed
}

type OnChainData struct {
	IPFSUri     string `json:"ipfs_uri" gorm:"column:ipfs_uri"`
	Transaction string `json:"transaction" gorm:"index"` // Query for unpublished feeds
}
