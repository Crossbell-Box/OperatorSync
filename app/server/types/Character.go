package types

type Character struct {
	// Structure related
	CrossbellCharacterID string `json:"crossbell_character_id" gorm:"column:crossbell_character_id"`

	//// Accounts related
	//MediaUsage MediaTypeRecordArray `gorm:"type:text" json:"media_usage"`
}
