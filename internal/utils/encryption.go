package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Hash hashes data using Argon2 and returns salt$hash
func Hash(data string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(data), salt, 1, 64*1024, 4, 32)
	saltEncoded := base64.StdEncoding.EncodeToString(salt)
	hashEncoded := base64.StdEncoding.EncodeToString(hash)
	return fmt.Sprintf("%s$%s", saltEncoded, hashEncoded), nil
}

// Verify checks if data matches the hashedData (salt$hash)
func Verify(data, hashedData string) (bool, error) {
	parts := strings.Split(hashedData, "$")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid hash format")
	}
	salt, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}
	hash, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}
	providedHash := argon2.IDKey([]byte(data), salt, 1, 64*1024, 4, 32)
	if subtle.ConstantTimeCompare(providedHash, hash) == 1 {
		return true, nil
	}
	return false, nil
}
