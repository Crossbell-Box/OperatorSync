package types

type Character struct {
	// Structure related
	CrossbellCharacterID string `json:"crossbell_character_id" gorm:"primarykey;column:crossbell_character_id"`

	// Accounts related
	MediaUsage uint `json:"media_usage"`
}
