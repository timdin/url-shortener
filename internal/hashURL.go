package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"url-shortener/constants"
)

func HashURL(url string) string {
	// Calculate the SHA-256 hash of the input string
	hash := sha256.Sum256([]byte(url))

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash[:])

	// Truncate the hash to predefined length in constants
	truncatedHash := hashString[:constants.SHORT_URL_LENGTH]
	return truncatedHash
}
