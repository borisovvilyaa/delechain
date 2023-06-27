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
