package bingx

import (
	"fmt"

	"github.com/yvv4git/api-go-bingx/utils"
)

// SwapBalancePayload - specific url payload for swap balance operation.
type SwapBalancePayload struct {
	timestamp int64
	signature string
}

// NewSwapBalancePayload - used for create instance of SwapBalancePayload.
func NewSwapBalancePayload(timestamp int64, secret string) *SwapBalancePayload {
	timestampStr := fmt.Sprintf("timestamp=%d", timestamp)

	return &SwapBalancePayload{
		timestamp: timestamp,
		signature: utils.Sign(secret, timestampStr),
	}
}

// Create - used for create payload.
func (s *SwapBalancePayload) Create() string {
	return fmt.Sprintf("timestamp=%d&signature=%s", s.timestamp, s.signature)
}
