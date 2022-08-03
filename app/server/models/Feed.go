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

type FeedsArray []Feed

func (feeds FeedsArray) Len() int {
	return len(feeds)
}
func (feeds FeedsArray) Swap(i, j int) {
	feeds[i], feeds[j] = feeds[j], feeds[i]
}
func (feeds FeedsArray) Less(i, j int) bool {
	return feeds[i].PublishedAt.Before(feeds[j].PublishedAt)
}
