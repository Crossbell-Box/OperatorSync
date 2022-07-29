package types

type ValidateRequest struct {
	Platform             string `json:"platform"`
	Username             string `json:"username"` // The unique identifier on platform
	CrossbellCharacterID string `json:"crossbell_character_id"`
}
