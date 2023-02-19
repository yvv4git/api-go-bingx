package bingx

import (
	"fmt"

	"github.com/yvv4git/api-go-bingx/utils"
)

// SpotDepthPayload - specific url payload for swap balance operation.
type SpotDepthPayload struct {
	payload   string
	signature string
}

// NewSpotDepthPayload - used for create instance of SpotDepthPayload.
func NewSpotDepthPayload(timestamp int64, limit int, symbol, secret string) *SpotDepthPayload {
	payload := fmt.Sprintf("symbol=%s&limit=%d&timestamp=%d", symbol, limit, timestamp)

	return &SpotDepthPayload{
		payload:   payload,
		signature: utils.Sign(secret, payload),
	}
}

// Create - used for create payload.
func (s *SpotDepthPayload) Create() string {
	return fmt.Sprintf("%s&signature=%s", s.payload, s.signature)
}
