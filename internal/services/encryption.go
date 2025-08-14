package services

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// EncryptionService handles data hashing and verification
type EncryptionService struct{}

// NewEncryptionService creates a new encryption service
func NewEncryptionService(secretKey string) *EncryptionService {
	return &EncryptionService{}
}

// Hash hashes data using Argon2
func (es *EncryptionService) Hash(data string) (string, error) {
	// Generate random salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Hash data using Argon2
	hash := argon2.IDKey([]byte(data), salt, 1, 64*1024, 4, 32)

	// Encode salt and hash
	saltEncoded := base64.StdEncoding.EncodeToString(salt)
	hashEncoded := base64.StdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s$%s", saltEncoded, hashEncoded), nil
}

// Verify verifies data against its hash
func (es *EncryptionService) Verify(data, hashedData string) (bool, error) {
	// Parse salt and hash
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

	// Hash provided data with same salt
	providedHash := argon2.IDKey([]byte(data), salt, 1, 64*1024, 4, 32)

	// Compare hashes
	return subtle.ConstantTimeCompare(hash, providedHash) == 1, nil
}
