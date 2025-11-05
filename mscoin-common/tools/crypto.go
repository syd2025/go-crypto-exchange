package tools

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func ComputeHmacSha256(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return hex.EncodeToString([]byte(sha))
}
