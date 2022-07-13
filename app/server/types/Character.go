package types

import "time"

type Character struct {
	// Structure related
	CrossbellCharacter string `json:"crossbell_character" gorm:"primarykey"`

	// Accounts related
	AccountLastUpdatedAt time.Time `json:"account_last_updated_at" gorm:"index"`
	MediaUsage           uint      `json:"media_usage"`
}
