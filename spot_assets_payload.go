package bingx

import (
	"fmt"

	"github.com/yvv4git/api-go-bingx/utils"
)

// SpotAssetsPayload - specific url payload for swap balance operation.
type SpotAssetsPayload struct {
	timestamp int64
	signature string
}

// NewSpotAssetsPayload - used for create instance of SpotAssetsPayload.
func NewSpotAssetsPayload(timestamp int64, secret string) *SpotAssetsPayload {
	timestampStr := fmt.Sprintf("timestamp=%d", timestamp)

	return &SpotAssetsPayload{
		timestamp: timestamp,
		signature: utils.Sign(secret, timestampStr),
	}
}

// Create - used for create payload.
func (s *SpotAssetsPayload) Create() string {
	return fmt.Sprintf("timestamp=%d&signature=%s", s.timestamp, s.signature)
}
