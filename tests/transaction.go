package main

import blockchain "blockchain/models"

func PrintTransaction() {
	// Creating a sample transaction
	transaction := blockchain.CreateTransaction(1, 1234, "0123456789abcdef", 0, "abcdef0123456789", 1, 1000, "0123456789abcdef")

	// Printing transaction information
	blockchain.PrintTransactionTable(transaction)
}
