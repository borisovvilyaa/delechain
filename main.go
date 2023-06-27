package main

import (
	blockchain "blockchain/models"
	"fmt"
)

func main() {
	keys := blockchain.GetKeyPair()

	// Вывод приватного и публичного ключей
	fmt.Println("Приватный ключ:", keys.PrivateKey)
	fmt.Println("Публичный ключ:", keys.PublicKey)
}
