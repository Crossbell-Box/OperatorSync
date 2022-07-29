package types

import "time"

type Account struct {
	// Structure related
	CrossbellCharacterID string `json:"crossbell_character_id" gorm:"index;column:crossbell_character_id"`

	// Metadata
	Platform string `json:"platform" gorm:"index"` // Platform ID
	Username string `json:"username" gorm:"index"`

	// Update related fields
	LastUpdated    time.Time     `json:"-"`              // Only update when new data found
	UpdateInterval time.Duration `json:"-"`              // UpdateInterval = [min, (NextUpdate - LastUpdated), max]
	NextUpdate     time.Time     `json:"-" gorm:"index"` // NextUpdate += UpdateInterval

	FeedsCount uint `json:"feeds_count"` // Recorded feeds
	NotesCount uint `json:"notes_count"` // On-Chain notes
	MediaUsage uint `json:"media_usage"`
}
