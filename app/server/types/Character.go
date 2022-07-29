package types

import "time"

type Character struct {
	// Structure related
	CrossbellCharacterID string `json:"crossbell_character_id" gorm:"primarykey;column:crossbell_character_id"`

	// Accounts related
	AccountLastUpdatedAt time.Time `json:"account_last_updated_at" gorm:"index"`
	MediaUsage           uint      `json:"media_usage"`
}
