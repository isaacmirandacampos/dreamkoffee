package password_hashing

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

func Encrypt(password string) (string, error) {
	const (
		time      = 1
		memory    = 64 * 1024
		threads   = 4
		keyLength = 32
	)
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)

	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	return fmt.Sprintf("%s.%s", encodedSalt, encodedHash), nil
}

func generateSalt() ([]byte, error) {
	const saltLength = 16
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func ComparePassword(password, encoded string) (bool, error) {
	const (
		time      = 1
		memory    = 64 * 1024
		threads   = 4
		keyLength = 32
	)
	parts := strings.Split(encoded, ".")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid hash format")
	}

	encodedSalt := parts[0]
	encodedHashPart := parts[1]

	// Decode the salt and hash
	salt, err := base64.RawStdEncoding.DecodeString(encodedSalt)
	if err != nil {
		return false, err
	}
	hash, err := base64.RawStdEncoding.DecodeString(encodedHashPart)
	if err != nil {
		return false, err
	}

	inputHash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)

	if subtle.ConstantTimeCompare(inputHash, hash) == 1 {
		return true, nil
	}
	return false, nil
}
