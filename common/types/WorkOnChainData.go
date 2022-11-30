package types

type CheckOnChainDataRequest struct {
	CrossbellCharacterID string `json:"crossbell_character_id"`
}

type CheckOnChainDataResponse struct {
	IsSucceeded       bool      `json:"is_succeeded"`
	Message           string    `json:"message"`
	IsOperatorValid   bool      `json:"is_operator_valid"`
	ConnectedAccounts []Account `json:"connected_accounts"`
}
