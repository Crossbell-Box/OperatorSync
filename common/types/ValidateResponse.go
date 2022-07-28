package types

type ValidateResponse struct {
	IsSucceeded            bool   `json:"is_succeeded"`
	Code                   uint   `json:"code"`
	Message                string `json:"message"`
	IsValidateStringExists bool   `json:"is_valid"`
}
