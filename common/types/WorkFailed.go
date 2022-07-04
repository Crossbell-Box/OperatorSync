package types

import "time"

type WorkFailed struct {
	WorkDispatched

	AcceptAt    time.Time `json:"accept_at"`
	ErrorAt     time.Time `json:"error_at"`
	ErrorCode   uint      `json:"error_code"`
	ErrorReason string    `json:"error_reason"`
}
