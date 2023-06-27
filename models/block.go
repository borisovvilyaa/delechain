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

func CreateTransaction(version int32, lockTime uint32, inputHash string, inputIndex uint32, inputScriptSig string, inputSequence uint32, outputValue int64, outputScriptPubKey string) Transaction {
	// Создаем новую транзакцию
	transaction := Transaction{
		Version:  version,
		LockTime: lockTime,
	}

	// Создаем входы транзакции
	input := TxInput{
		PreviousOutput: OutPoint{
			Hash:  hexToBytes(inputHash),
			Index: inputIndex,
		},
		ScriptSig: hexToBytes(inputScriptSig),
		Sequence:  inputSequence,
	}

	// Создаем выходы транзакции
	output := TxOutput{
		Value:        outputValue,
		ScriptPubKey: hexToBytes(outputScriptPubKey),
	}

	// Добавляем входы и выходы к транзакции
	transaction.Inputs = append(transaction.Inputs, input)
	transaction.Outputs = append(transaction.Outputs, output)

	return transaction
}

// Вспомогательная функция для преобразования строки в байтовый срез
func hexToBytes(s string) []byte {
	bytes, _ := hex.DecodeString(s)
	return bytes
}

func PrintTransactionTable(transaction Transaction) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})

	table.Append([]string{"Version", fmt.Sprintf("%d", transaction.Version)})
	table.Append([]string{"LockTime", fmt.Sprintf("%d", transaction.LockTime)})

	// Выводим информацию о входах транзакции
	for i, input := range transaction.Inputs {
		table.Append([]string{"Input", fmt.Sprintf("Input %d", i+1)})
		table.Append([]string{"- Previous Output Hash", hex.EncodeToString(input.PreviousOutput.Hash)})
		table.Append([]string{"- Previous Output Index", fmt.Sprintf("%d", input.PreviousOutput.Index)})
		table.Append([]string{"- ScriptSig", hex.EncodeToString(input.ScriptSig)})
		table.Append([]string{"- Sequence", fmt.Sprintf("%d", input.Sequence)})
	}

	// Выводим информацию о выходах транзакции
	for i, output := range transaction.Outputs {
		table.Append([]string{"Output", fmt.Sprintf("Output %d", i+1)})
		table.Append([]string{"- Value", fmt.Sprintf("%d", output.Value)})
		table.Append([]string{"- ScriptPubKey", hex.EncodeToString(output.ScriptPubKey)})
	}

	table.Render()
}
