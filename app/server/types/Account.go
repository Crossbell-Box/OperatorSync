package types

import "time"

type Account struct {
	// Structure related
	CrossbellCharacterID string `json:"crossbell_character_id" gorm:"index;column:crossbell_character_id"`

	// Metadata
	Platform string `json:"platform" gorm:"index"` // Platform ID
	Username string `json:"username" gorm:"index"`

	// Update related fields
	LastUpdated    time.Time     `json:"last_updated"`             // Only update when new data found
	UpdateInterval time.Duration `json:"update_interval"`          // UpdateInterval = [min, (NextUpdate - LastUpdated), max]
	NextUpdate     time.Time     `json:"next_update" gorm:"index"` // NextUpdate += UpdateInterval

	FeedsCount uint                 `json:"feeds_count"` // Recorded feeds
	NotesCount uint                 `json:"notes_count"` // On-Chain notes
	MediaUsage MediaTypeRecordArray `gorm:"type:text" json:"media_usage"`
}

type OnChainStatusManageForAccount struct {
	IsOnChainPaused     bool      `json:"is_onchain_paused" gorm:"index;column:is_onchain_paused"`
	OnChainPausedAt     time.Time `json:"onChain_paused_at"`
	OnChainPauseMessage string    `json:"onChain_pause_message"`
}
