package wallet

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

func AuthenticateRequest(body []byte, digest string) bool {
	key := []byte("your-secret-key") // Replace with your secret key
	hash := hmac.New(sha1.New, key)
	hash.Write(body)
	expectedDigest := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return hmac.Equal([]byte(digest), []byte(expectedDigest))
}
