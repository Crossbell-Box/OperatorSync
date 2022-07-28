package types

type ValidateRequest struct {
	Platform           string `json:"platform"`
	Username           string `json:"username"` // The unique identifier on platform
	CrossbellCharacter string `json:"crossbell_character"`
}
