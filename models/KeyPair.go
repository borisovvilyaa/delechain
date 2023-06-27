package blockchain

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type KeyPair struct {
	PrivateKey string
	PublicKey  string
}

// generatePrivateKey generates a random private key.
// @return string: The generated private key.
func generatePrivateKey() string {
	// Generate a random byte array
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// Hash the random byte array using SHA-256
	hash := sha256.Sum256(randomBytes)

	// Private key is the hex-encoded hash
	privateKey := hex.EncodeToString(hash[:])
	return privateKey
}

// generatePublicKey generates a public key based on the provided private key.
// @param privateKey: The private key used to generate the public key.
// @return string: The generated public key.
func generatePublicKey(privateKey string) string {
	// Convert the private key from hex to bytes
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		panic(err)
	}

	// Hash the private key bytes using SHA-256
	hash := sha256.Sum256(privateKeyBytes)

	// Public key is the hex-encoded hash
	publicKey := hex.EncodeToString(hash[:])
	return publicKey
}

// GetKeyPair generates a key pair consisting of a private key and a public key.
// @return KeyPair: The generated key pair.
func GetKeyPair() KeyPair {
	// Generate private key
	privateKey := generatePrivateKey()

	// Generate public key based on private key
	publicKey := generatePublicKey(privateKey)

	keys := KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
	return keys
}
