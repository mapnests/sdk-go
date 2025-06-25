package sdk

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

const (
	TOKEN_EXPIRY_SECONDS = 30
)

func generateToken(apiKey string) (string, error) {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	expires := time.Now().Unix() + TOKEN_EXPIRY_SECONDS

	input := fmt.Sprintf("%d:%d:%s", expires, randInt.Int64(), apiKey)
	hash := sha256.Sum256([]byte(input))
	sign := hex.EncodeToString(hash[:])

	jsonPayload := fmt.Sprintf(`{"random":%d,"expires":%d,"sign":"%s"}`,
		randInt.Int64(), expires, sign)

	return base64.StdEncoding.EncodeToString([]byte(jsonPayload)), nil
}
