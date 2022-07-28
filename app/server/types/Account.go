package types

import "time"

type Account struct {
	// Structure related
	CrossbellCharacter string `json:"crossbell_character" gorm:"index;column:crossbell_character"`

	// Metadata
	Platform string `json:"platform" gorm:"index"` // Platform ID
	Username string `json:"username" gorm:"index"`

	// Update related fields
	LastUpdated    time.Time     `json:"-"`              // Only update when new data found
	UpdateInterval time.Duration `json:"-"`              // UpdateInterval = [min, (NextUpdate - LastUpdated), max]
	NextUpdate     time.Time     `json:"-" gorm:"index"` // NextUpdate += UpdateInterval
}
