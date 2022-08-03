package models

import (
	"database/sql/driver"
	"encoding/json"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"time"
)

type MediaFeedRecord struct {
	Platform string `json:"platform"`
	ID       uint   `json:"id"`
}

type MediaFeedRecordArray []MediaFeedRecord

func (fa *MediaFeedRecordArray) Scan(src interface{}) error {
	return json.Unmarshal([]byte(src.(string)), fa)
}

func (fa MediaFeedRecordArray) Value() (driver.Value, error) {
	val, err := json.Marshal(&fa)
	return string(val), err
}

type Media struct {
	// Database related fields
	ID        uint      `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `gorm:"index" json:"-"`

	// Structure related
	CrossbellCharacterID string               `gorm:"index" json:"crossbell_character_id"`
	RelatedFeeds         MediaFeedRecordArray `gorm:"type:text" json:"related_feeds"`

	commonTypes.Media
}
