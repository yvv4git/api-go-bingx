package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Sign - used for data signing.
func Sign(secret, payload string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))

	return hex.EncodeToString(h.Sum(nil))
}
