package testdata

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// TxParams holds parameters for generating a Thorchain transaction
type TxParams struct {
	FromAddr  string
	ToAddr    string
	Amount    int64
	Denom     string
	Memo      string
	GasLimit  uint64
	FeeAmount int64
	Sequence  uint64
}

// TestCase represents a test scenario with expected result
type TestCase struct {
	Name        string
	Params      TxParams
	ShouldPass  bool
	Description string
}

// GetThorchainTestCases returns predefined test cases for payroll.json policy
func GetThorchainTestCases() []TestCase {
	return []TestCase{
		{
			Name:        "PASS - Exact limit 100 RUNE",
			ShouldPass:  true,
			Description: "Transfer exactly 100 RUNE to allowed recipient",
			Params: TxParams{
				FromAddr:  "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
				ToAddr:    "thor1c3xl04vgw7f7q9xrmqc4llr6s9l8p7qmgqwzsl",
				Amount:    10000000000, // 100 RUNE
				Denom:     "rune",
				Memo:      "Max allowed transfer",
				GasLimit:  200000,
				FeeAmount: 2000000,
				Sequence:  1,
			},
		},
		{
			Name:        "PASS - Under limit 50 RUNE",
			ShouldPass:  true,
			Description: "Transfer 50 RUNE to allowed recipient (under limit)",
			Params: TxParams{
				FromAddr:  "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
				ToAddr:    "thor1c3xl04vgw7f7q9xrmqc4llr6s9l8p7qmgqwzsl",
				Amount:    5000000000, // 50 RUNE
				Denom:     "rune",
				Memo:      "Partial payment",
				GasLimit:  200000,
				FeeAmount: 2000000,
				Sequence:  2,
			},
		},
		{
			Name:        "FAIL - Amount exceeds 150 RUNE",
			ShouldPass:  false,
			Description: "Transfer 150 RUNE exceeds policy limit of 100 RUNE",
			Params: TxParams{
				FromAddr:  "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
				ToAddr:    "thor1c3xl04vgw7f7q9xrmqc4llr6s9l8p7qmgqwzsl",
				Amount:    15000000000, // 150 RUNE
				Denom:     "rune",
				Memo:      "Exceeds limit",
				GasLimit:  200000,
				FeeAmount: 2000000,
				Sequence:  3,
			},
		},
		{
			Name:        "FAIL - Wrong recipient",
			ShouldPass:  false,
			Description: "Transfer to unauthorized recipient address",
			Params: TxParams{
				FromAddr:  "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
				ToAddr:    "thor1wrongaddress000000000000000000000000",
				Amount:    5000000000, // 50 RUNE
				Denom:     "rune",
				Memo:      "Wrong recipient",
				GasLimit:  200000,
				FeeAmount: 2000000,
				Sequence:  4,
			},
		},
		{
			Name:        "FAIL - Wrong denomination",
			ShouldPass:  false,
			Description: "Transfer TCY instead of RUNE (wrong protocol)",
			Params: TxParams{
				FromAddr:  "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
				ToAddr:    "thor1c3xl04vgw7f7q9xrmqc4llr6s9l8p7qmgqwzsl",
				Amount:    5000000000, // 50 TCY
				Denom:     "tcy",
				Memo:      "Wrong token",
				GasLimit:  200000,
				FeeAmount: 2000000,
				Sequence:  5,
			},
		},
	}
}

// TxHex functions for engine_test.go (following Uniswap pattern)

// ValidRuneTransferMaxTxHex returns hex for valid 100 RUNE transfer (should pass)
func ValidRuneTransferMaxTxHex() string {
	params := TxParams{
		FromAddr:  "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
		ToAddr:    "thor1c3xl04vgw7f7q9xrmqc4llr6s9l8p7qmgqwzsl",
		Amount:    10000000000, // 100 RUNE
		Denom:     "rune",
		Memo:      "Max allowed transfer",
		GasLimit:  200000,
		FeeAmount: 2000000,
		Sequence:  1,
	}

	txHex, err := GenerateThorchainTx(params)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate valid RUNE transfer tx: %v", err))
	}
	return txHex
}

// ValidRuneTransferUnderLimitTxHex returns hex for valid 50 RUNE transfer (should pass)
func ValidRuneTransferUnderLimitTxHex() string {
	params := TxParams{
		FromAddr:  "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
		ToAddr:    "thor1c3xl04vgw7f7q9xrmqc4llr6s9l8p7qmgqwzsl",
		Amount:    5000000000, // 50 RUNE
		Denom:     "rune",
		Memo:      "Partial payment",
		GasLimit:  200000,
		FeeAmount: 2000000,
		Sequence:  2,
	}

	txHex, err := GenerateThorchainTx(params)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate valid under-limit RUNE transfer tx: %v", err))
	}
	return txHex
}

// ExceedAmountRuneTransferTxHex returns hex for 150 RUNE transfer (should fail - exceeds limit)
func ExceedAmountRuneTransferTxHex() string {
	params := TxParams{
		FromAddr:  "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
		ToAddr:    "thor1c3xl04vgw7f7q9xrmqc4llr6s9l8p7qmgqwzsl",
		Amount:    15000000000, // 150 RUNE
		Denom:     "rune",
		Memo:      "Exceeds limit",
		GasLimit:  200000,
		FeeAmount: 2000000,
		Sequence:  3,
	}

	txHex, err := GenerateThorchainTx(params)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate exceed-amount RUNE transfer tx: %v", err))
	}
	return txHex
}

// InvalidRecipientRuneTransferTxHex returns hex for wrong recipient (should fail)
func InvalidRecipientRuneTransferTxHex() string {
	params := TxParams{
		FromAddr:  "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
		ToAddr:    "thor1wrongaddress000000000000000000000000",
		Amount:    5000000000, // 50 RUNE
		Denom:     "rune",
		Memo:      "Wrong recipient",
		GasLimit:  200000,
		FeeAmount: 2000000,
		Sequence:  4,
	}

	txHex, err := GenerateThorchainTx(params)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate invalid recipient RUNE transfer tx: %v", err))
	}
	return txHex
}

// InvalidDenomTCYTransferTxHex returns hex for TCY transfer (should fail - wrong protocol)
func InvalidDenomTCYTransferTxHex() string {
	params := TxParams{
		FromAddr:  "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
		ToAddr:    "thor1c3xl04vgw7f7q9xrmqc4llr6s9l8p7qmgqwzsl",
		Amount:    5000000000, // 50 TCY
		Denom:     "tcy",
		Memo:      "Wrong token",
		GasLimit:  200000,
		FeeAmount: 2000000,
		Sequence:  5,
	}

	txHex, err := GenerateThorchainTx(params)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate invalid denom TCY transfer tx: %v", err))
	}
	return txHex
}

// GenerateThorchainTx creates a Thorchain transaction hex from parameters
func GenerateThorchainTx(params TxParams) (string, error) {
	// Create codec
	interfaceRegistry := types.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	// Create coin for amount
	amount := sdktypes.NewInt64Coin(params.Denom, params.Amount)

	// Create MsgSend
	msgSend := &banktypes.MsgSend{
		FromAddress: params.FromAddr,
		ToAddress:   params.ToAddr,
		Amount:      []sdktypes.Coin{amount},
	}

	// Create message Any
	msgAny, err := types.NewAnyWithValue(msgSend)
	if err != nil {
		return "", fmt.Errorf("failed to create message Any: %w", err)
	}

	// Create transaction body
	txBody := &tx.TxBody{
		Messages: []*types.Any{msgAny},
		Memo:     params.Memo,
	}

	// Create fee
	fee := &tx.Fee{
		Amount:   []sdktypes.Coin{sdktypes.NewInt64Coin(params.Denom, params.FeeAmount)},
		GasLimit: params.GasLimit,
	}

	// Create auth info
	authInfo := &tx.AuthInfo{
		Fee: fee,
		SignerInfos: []*tx.SignerInfo{
			{
				Sequence: params.Sequence,
			},
		},
	}

	// Create the complete transaction
	cosmosTx := &tx.Tx{
		Body:     txBody,
		AuthInfo: authInfo,
	}

	// Encode the transaction to bytes
	txBytes, err := cdc.Marshal(cosmosTx)
	if err != nil {
		return "", fmt.Errorf("failed to marshal transaction: %w", err)
	}

	// Convert to hex string with 0x prefix
	txHex := "0x" + hex.EncodeToString(txBytes)
	return txHex, nil
}

// PrintTestCase prints details of a test case (for debugging)
func PrintTestCase(testCase TestCase) {
	fmt.Printf("🔹 %s\n", testCase.Name)
	fmt.Printf("Expected: %s\n", map[bool]string{true: "PASS", false: "FAIL"}[testCase.ShouldPass])
	fmt.Printf("From: %s\n", testCase.Params.FromAddr)
	fmt.Printf("To: %s\n", testCase.Params.ToAddr)
	fmt.Printf("Amount: %s %s (%.8f %s)\n",
		formatAmount(testCase.Params.Amount), testCase.Params.Denom,
		float64(testCase.Params.Amount)/100000000, strings.ToUpper(testCase.Params.Denom))
	fmt.Printf("Memo: %s\n", testCase.Params.Memo)
	fmt.Printf("Resource: thorchain.%s.transfer\n", testCase.Params.Denom)
	fmt.Printf("Description: %s\n", testCase.Description)
	fmt.Println("---")
}

// formatAmount converts int64 to string
func formatAmount(amount int64) string {
	return strconv.FormatInt(amount, 10)
}
