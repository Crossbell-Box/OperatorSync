package models

import (
	"database/sql/driver"
	"encoding/json"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

type FeedRecord struct {
	Platform string `json:"platform"`
	ID       uint   `json:"id"`
}

type FeedRecordArray []FeedRecord

func (fa *FeedRecordArray) Scan(src interface{}) error {
	return json.Unmarshal([]byte(src.(string)), fa)
}

func (fa FeedRecordArray) Value() (driver.Value, error) {
	val, err := json.Marshal(&fa)
	return string(val), err
}

type Media struct {
	// Database related fields
	ID        uint      `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `gorm:"index" json:"-"`

	// Structure related
	CrossbellCharacterID string          `gorm:"index" json:"crossbell_character_id"`
	RelatedFeeds         FeedRecordArray `gorm:"type:text" json:"related_feeds"`

	commonTypes.Media
}
