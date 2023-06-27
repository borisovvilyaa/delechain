package blockchain

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type Transaction struct {
	Version  int32
	Inputs   []TxInput
	Outputs  []TxOutput
	LockTime uint32
}

type TxInput struct {
	PreviousOutput OutPoint
	ScriptSig      []byte
	Sequence       uint32
}

type OutPoint struct {
	Hash  []byte
	Index uint32
}

type TxOutput struct {
	Value        int64
	ScriptPubKey []byte
}

// CreateTransaction creates a new transaction with the specified parameters.
// @param version: The version of the transaction.
// @param lockTime: The lock time of the transaction.
// @param inputHash: The hash of the previous output being spent.
// @param inputIndex: The index of the previous output being spent.
// @param inputScriptSig: The script signature for the input.
// @param inputSequence: The sequence number of the input.
// @param outputValue: The value of the output.
// @param outputScriptPubKey: The script public key of the output.
// @return Transaction: The created transaction.
func CreateTransaction(version int32, lockTime uint32, inputHash string, inputIndex uint32, inputScriptSig string, inputSequence uint32, outputValue int64, outputScriptPubKey string) Transaction {
	// Create a new transaction
	transaction := Transaction{
		Version:  version,
		LockTime: lockTime,
	}

	// Create transaction inputs
	input := TxInput{
		PreviousOutput: OutPoint{
			Hash:  hexToBytes(inputHash),
			Index: inputIndex,
		},
		ScriptSig: hexToBytes(inputScriptSig),
		Sequence:  inputSequence,
	}

	// Create transaction outputs
	output := TxOutput{
		Value:        outputValue,
		ScriptPubKey: hexToBytes(outputScriptPubKey),
	}

	// Add inputs and outputs to the transaction
	transaction.Inputs = append(transaction.Inputs, input)
	transaction.Outputs = append(transaction.Outputs, output)

	return transaction
}

// hexToBytes is a utility function to convert a string to a byte slice.
func hexToBytes(s string) []byte {
	bytes, _ := hex.DecodeString(s)
	return bytes
}

// PrintTransactionTable prints the information about a transaction in a table format.
// @param transaction: The transaction to be printed.
func PrintTransactionTable(transaction Transaction) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})

	table.Append([]string{"Version", fmt.Sprintf("%d", transaction.Version)})
	table.Append([]string{"LockTime", fmt.Sprintf("%d", transaction.LockTime)})

	// Print information about transaction inputs
	for i, input := range transaction.Inputs {
		table.Append([]string{"Input", fmt.Sprintf("Input %d", i+1)})
		table.Append([]string{"- Previous Output Hash", hex.EncodeToString(input.PreviousOutput.Hash)})
		table.Append([]string{"- Previous Output Index", fmt.Sprintf("%d", input.PreviousOutput.Index)})
		table.Append([]string{"- ScriptSig", hex.EncodeToString(input.ScriptSig)})
		table.Append([]string{"- Sequence", fmt.Sprintf("%d", input.Sequence)})
	}

	// Print information about transaction outputs
	for i, output := range transaction.Outputs {
		table.Append([]string{"Output", fmt.Sprintf("Output %d", i+1)})
		table.Append([]string{"- Value", fmt.Sprintf("%d", output.Value)})
		table.Append([]string{"- ScriptPubKey", hex.EncodeToString(output.ScriptPubKey)})
	}

	table.Render()
}
