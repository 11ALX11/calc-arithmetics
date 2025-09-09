package app

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

const GCM_nonce_size = 12 // GCM nonce size is 12 bytes

// Same as Encrypt, but use key file path instead of a key string
func EncryptFileKey(ciphertext, keyPath string) (string, error) {
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return "", err
	}

	return Encrypt(ciphertext, string(key))
}

// Encrypt encrypts a plaintext string using AES and returns the ciphertext
func Encrypt(plaintext, key string) (string, error) {
	// Convert the key to a byte array
	keyBytes := []byte(key)

	// Create a new AES cipher
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Generate a new nonce
	nonce := make([]byte, GCM_nonce_size)
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Create a GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Encrypt the plaintext
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// Encode to base64 to return as a string
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
