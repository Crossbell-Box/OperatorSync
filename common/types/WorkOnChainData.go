package types

type CheckOnChainDataRequest struct {
	CrossbellCharacterID string `json:"crossbell_character_id"`
	Platform             string `json:"platform"`
	Account              string `json:"account"`
}

type CheckOnChainDataResponse struct {
	IsSucceeded        bool   `json:"is_succeeded"`
	Message            string `json:"message"`
	IsOperatorValid    bool   `json:"is_operator_valid"`
	IsAccountConnected bool   `json:"is_platform_set"`
}
