package types

import "time"

type AccountMetrics struct {
	Total int64 `json:"total"`
	Valid int64 `json:"valid"` // With at least one feed
}

type PlatformMetrics struct {
	Feed    int64          `json:"feed"`
	Account AccountMetrics `json:"account"`
}

type Metrics struct {
	GeneratedAt time.Time `json:"generated_at"`
	//Address     struct {
	//	Total int64 `json:"total"` // How many unique addresses for characters
	//} `json:"address"`
	Character struct {
		Total int64 `json:"total"`
		Valid int64 `json:"valid"` // With at least one account
	} `json:"character"`
	Account  AccountMetrics             `json:"account"`
	Platform map[string]PlatformMetrics `json:"platform"`
}
