package models

import (
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	"gorm.io/gorm"
	"time"
)

type Feed struct {
	// Database related fields
	ID        uint      `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `gorm:"index" json:"-"`

	types.Feed
	types.OnChainData
}

const feedPrefix = "feed_"

func FeedTable(f Feed) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Table(feedPrefix + f.Platform)
	}
}
