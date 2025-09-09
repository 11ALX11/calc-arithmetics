package app

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

// Same as Decrypt, but use key file path instead of a key string
func DecryptFileKey(ciphertext, keyPath string) (string, error) {
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return "", err
	}

	return Decrypt(ciphertext, string(key))
}

// Decrypt decrypts a ciphertext string using AES and returns the plaintext
func Decrypt(ciphertext, key string) (string, error) {
	// Convert the key to a byte array
	keyBytes := []byte(key)

	// Decode the base64 ciphertext
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// Check if string doesn't contain nonce (probably its empty or malformed)
	if len(data) < GCM_nonce_size {
		return "", fmt.Errorf("decoded base64 string doesn't contain nonce. Possible reason is empty or malformed encrypted message")
	}

	// Create a new AES cipher
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Create a GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Split the nonce and the ciphertext
	nonce, ciphertextPure := data[:GCM_nonce_size], data[GCM_nonce_size:]

	// Decrypt the ciphertext
	plaintext, err := gcm.Open(nil, nonce, ciphertextPure, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
