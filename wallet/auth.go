package wallet

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"os"
)

func AuthenticateRequest(body []byte, digest string) bool {
	secretKey := os.Getenv("SECRET_KEY")
	key := []byte(secretKey)
	hash := hmac.New(sha1.New, key)
	hash.Write(body)
	expectedDigest := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return hmac.Equal([]byte(digest), []byte(expectedDigest))
}
