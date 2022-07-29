package types

type ValidateRequest struct {
	Platform             string `json:"platform"`
	Username             string `json:"username"` // The unique identifier on platform
	CrossbellCharacterID string `json:"crossbell_character_id"`
}

type ValidateResponse struct {
	IsSucceeded            bool   `json:"is_succeeded"`
	Code                   uint   `json:"code"`
	Message                string `json:"message"`
	IsValidateStringExists bool   `json:"is_valid"`
}
