package cryptoutils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword will take a string as an input and
// returns the hashed string to securily store in the
// db
func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum(nil))
}
