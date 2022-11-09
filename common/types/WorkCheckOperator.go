package types

type CheckOperatorRequest struct {
	CrossbellCharacterID string `json:"crossbell_character_id"`
}

type CheckOperatorResponse struct {
	IsSucceeded     bool   `json:"is_succeeded"`
	Message         string `json:"message"`
	IsOperatorValid bool   `json:"is_operator_valid"`
}
