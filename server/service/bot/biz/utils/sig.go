package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func VerifySignature(body []byte, signature string, key string) bool {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write(body)
	expectedMAC := mac.Sum(nil)
	expectedSig := hex.EncodeToString(expectedMAC)
	return hmac.Equal([]byte(expectedSig), []byte(signature))
}
