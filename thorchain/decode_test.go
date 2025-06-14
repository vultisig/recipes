package thorchain

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func TestParseThorchainTransaction(t *testing.T) {
	// Test with invalid hex data (not valid protobuf) - should now return error
	txHex := "0x1234567890abcdef1234567890abcdef1234567890abcdef"

	_, err := ParseThorchainTransaction(txHex)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode protobuf transaction")
}

func TestParseThorchainTransactionWithoutPrefix(t *testing.T) {
	// Test without 0x prefix (still invalid protobuf) - should return error
	txHex := "1234567890abcdef1234567890abcdef1234567890abcdef"

	_, err := ParseThorchainTransaction(txHex)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode protobuf transaction")
}

func TestParseThorchainTransactionTooShort(t *testing.T) {
	// Test with transaction data that's too short
	txHex := "0x12345678" // Only 4 bytes

	_, err := ParseThorchainTransaction(txHex)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "transaction too short")
}

func TestParseThorchainTransactionInvalidHex(t *testing.T) {
	// Test with invalid hex data
	txHex := "invalid_hex_data"

	_, err := ParseThorchainTransaction(txHex)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode hex transaction")
}

func TestParsedThorchainTransactionInterface(t *testing.T) {
	parsed := &ParsedThorchainTransaction{
		txHash:     "test_hash",
		rawData:    []byte("test_data"),
		fromAddr:   "thor1test_from",
		toAddr:     "thor1test_to",
		amount:     big.NewInt(1000000000), // 10 RUNE (8 decimals)
		denom:      "rune",
		memo:       "test memo",
		msgType:    "MsgSend",
		gasPrice:   big.NewInt(100),
		gasLimit:   200000,
		nonce:      0,
		sequence:   1,
		accountNum: 123,
	}

	// Test all interface methods
	assert.Equal(t, "thorchain", parsed.ChainIdentifier())
	assert.Equal(t, "test_hash", parsed.Hash())
	assert.Equal(t, "thor1test_from", parsed.From())
	assert.Equal(t, "thor1test_to", parsed.To())
	assert.Equal(t, big.NewInt(1000000000), parsed.Value())
	assert.Equal(t, []byte("test memo"), parsed.Data())
	assert.Equal(t, uint64(1), parsed.Nonce()) // Now returns sequence
	assert.Equal(t, big.NewInt(100), parsed.GasPrice())
	assert.Equal(t, uint64(200000), parsed.GasLimit())

	// Test Thorchain-specific methods
	assert.Equal(t, "test memo", parsed.GetMemo())
	assert.Equal(t, "rune", parsed.GetDenom())
	assert.Equal(t, "MsgSend", parsed.GetMsgType())
	assert.Equal(t, uint64(1), parsed.GetSequence())
	assert.Equal(t, uint64(123), parsed.GetAccountNumber())
}

func TestValidateThorchainAddress(t *testing.T) {
	// Test valid addresses with proper bech32 checksums
	validAddresses := []string{
		"thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2", // Valid Thorchain address
		"thor13m8mqtcv7c2srnpr4efucwkrdu29tq3ejgu52r", // Another valid Thorchain address
	}

	for _, addr := range validAddresses {
		err := ValidateThorchainAddress(addr)
		assert.NoError(t, err, "Address %s should be valid", addr)
	}

	// Test invalid addresses that should definitely fail
	invalidAddresses := []string{
		"eth0x1234567890abcdef", // Wrong prefix
		"thor123",               // Too short
		"cosmos1234567890abcdefghijklmnopqrstuvwxyz123", // Wrong prefix
		"", // Empty
		"thor1234567890abcdefghijklmnopqrstuvwxyz12345", // Too long (45 chars)
		"thor123456789@abcdefghijklmnopqrstuvwxyz123",   // Invalid character (@)
		"thor1qpyxw8nhed4afrxjgwru5vrtaz3mr3hskr6tkmw",  // Invalid checksum
	}

	for _, addr := range invalidAddresses {
		err := ValidateThorchainAddress(addr)
		assert.Error(t, err, "Address %s should be invalid", addr)
	}

	// Test that empty address fails
	err := ValidateThorchainAddress("")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "address cannot be empty")

	// Test that wrong prefix fails
	err = ValidateThorchainAddress("cosmos1234567890abcdefghijklmnopqrstuvwxyz123")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "address must start with 'thor'")
}

func TestValidateMemo(t *testing.T) {
	// Test valid memos
	validMemos := []string{
		"",                   // Empty memo is valid
		"SWAP:BTC.BTC",       // Valid operation memo
		"ADD:BTC.BTC:thor1x", // Valid liquidity memo
		"short memo",         // Short memo
	}

	for _, memo := range validMemos {
		err := ValidateMemo(memo)
		assert.NoError(t, err, "Memo %s should be valid", memo)
	}

	// Test invalid memo (too long)
	longMemo := string(make([]byte, MaxMemoLength+1))
	for i := range longMemo {
		longMemo = longMemo[:i] + "a" + longMemo[i+1:]
	}
	err := ValidateMemo(longMemo)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "memo too long")
}

func TestValidateAmount(t *testing.T) {
	// Test valid amounts
	validAmount := big.NewInt(100000000) // 1 RUNE
	err := ValidateAmount(validAmount, "rune")
	assert.NoError(t, err)

	// Test valid TCY amount
	validTCYAmount := big.NewInt(100000000) // 1 TCY
	err = ValidateAmount(validTCYAmount, "tcy")
	assert.NoError(t, err)

	// Test invalid amounts
	invalidAmounts := []*big.Int{
		nil,            // Nil amount
		big.NewInt(0),  // Zero amount
		big.NewInt(-1), // Negative amount
	}

	for _, amount := range invalidAmounts {
		err := ValidateAmount(amount, "rune")
		assert.Error(t, err)
	}

	// Test excessive amount for TCY (more than reasonable limit)
	excessiveTCY := new(big.Int).Mul(big.NewInt(600000000), big.NewInt(100000000)) // 600M TCY (above limit)
	err = ValidateAmount(excessiveTCY, "tcy")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "exceeds maximum allowed")
}

func TestComputeThorchainTxHash(t *testing.T) {
	// Test hash computation
	txBytes1 := []byte("test transaction 1")
	txBytes2 := []byte("test transaction 2")

	hash1 := computeThorchainTxHash(txBytes1)
	hash2 := computeThorchainTxHash(txBytes2)

	// Hashes should be different for different inputs
	assert.NotEqual(t, hash1, hash2)

	// Hash should be deterministic
	hash1Again := computeThorchainTxHash(txBytes1)
	assert.Equal(t, hash1, hash1Again)

	// Hash should be uppercase hex (like Cosmos)
	assert.Equal(t, 64, len(hash1)) // SHA256 = 32 bytes = 64 hex chars
	assert.Regexp(t, "^[A-F0-9]+$", hash1)
}

func TestRealCosmosTransactionParsing(t *testing.T) {
	// Create a real Cosmos SDK transaction for testing
	interfaceRegistry := types.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	// Use real valid Thorchain addresses for proper testing
	fromAddr := "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2"
	toAddr := "thor13m8mqtcv7c2srnpr4efucwkrdu29tq3ejgu52r"
	amount := sdktypes.NewInt64Coin("rune", 100000000) // 1 RUNE

	msgSend := &banktypes.MsgSend{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		Amount:      []sdktypes.Coin{amount},
	}

	// Create message Any
	msgAny, err := types.NewAnyWithValue(msgSend)
	require.NoError(t, err)

	// Create transaction body
	txBody := &tx.TxBody{
		Messages: []*types.Any{msgAny},
		Memo:     "test transfer",
	}

	// Create auth info with gas and fee
	fee := &tx.Fee{
		Amount:   []sdktypes.Coin{sdktypes.NewInt64Coin("rune", 2000000)}, // Fee amount (2M to get integer gas price)
		GasLimit: 200000,
	}

	authInfo := &tx.AuthInfo{
		Fee: fee,
		SignerInfos: []*tx.SignerInfo{
			{
				Sequence: 5,
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
	require.NoError(t, err)

	// Convert to hex string
	txHex := "0x" + hex.EncodeToString(txBytes)

	// Parse the transaction using our implementation
	decodedTx, err := ParseThorchainTransaction(txHex)
	require.NoError(t, err)

	// Verify extracted values
	assert.Equal(t, "thorchain", decodedTx.ChainIdentifier())
	assert.NotEmpty(t, decodedTx.Hash())

	// Test that actual values were extracted
	assert.Equal(t, fromAddr, decodedTx.From())
	assert.Equal(t, toAddr, decodedTx.To())
	assert.Equal(t, big.NewInt(100000000), decodedTx.Value())
	assert.Equal(t, uint64(200000), decodedTx.GasLimit())
	assert.Equal(t, uint64(5), decodedTx.Nonce()) // sequence
	assert.Equal(t, []byte("test transfer"), decodedTx.Data())

	// Test Thorchain-specific methods
	parsed := decodedTx.(*ParsedThorchainTransaction)
	assert.Equal(t, "test transfer", parsed.GetMemo())
	assert.Equal(t, "rune", parsed.GetDenom())
	assert.Equal(t, "MsgSend", parsed.GetMsgType())
	assert.Equal(t, uint64(5), parsed.GetSequence())

	// Test gas price calculation (fee amount / gas limit)
	expectedGasPrice := big.NewInt(10) // 2000000 / 200000 = 10
	assert.Equal(t, expectedGasPrice, decodedTx.GasPrice())
}

func TestCosmosTransactionParsingWithMultipleCoins(t *testing.T) {
	// Test with transaction containing multiple coins
	interfaceRegistry := types.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	// Use real valid Thorchain addresses for proper testing
	fromAddr := "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2"
	toAddr := "thor13m8mqtcv7c2srnpr4efucwkrdu29tq3ejgu52r"

	// Multiple coins with TCY as first coin
	coins := []sdktypes.Coin{
		sdktypes.NewInt64Coin("tcy", 50000000),   // 0.5 TCY (first coin)
		sdktypes.NewInt64Coin("rune", 100000000), // 1 RUNE
	}

	msgSend := &banktypes.MsgSend{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		Amount:      coins,
	}

	msgAny, err := types.NewAnyWithValue(msgSend)
	require.NoError(t, err)

	txBody := &tx.TxBody{
		Messages: []*types.Any{msgAny},
		Memo:     "multi-coin transfer",
	}

	authInfo := &tx.AuthInfo{
		Fee: &tx.Fee{
			Amount:   []sdktypes.Coin{sdktypes.NewInt64Coin("rune", 2000)},
			GasLimit: 250000,
		},
		SignerInfos: []*tx.SignerInfo{{Sequence: 10}},
	}

	cosmosTx := &tx.Tx{
		Body:     txBody,
		AuthInfo: authInfo,
	}

	txBytes, err := cdc.Marshal(cosmosTx)
	require.NoError(t, err)

	txHex := hex.EncodeToString(txBytes)

	// Parse the transaction
	decodedTx, err := ParseThorchainTransaction(txHex)
	require.NoError(t, err)

	// Should extract first coin (TCY)
	assert.Equal(t, big.NewInt(50000000), decodedTx.Value())

	parsed := decodedTx.(*ParsedThorchainTransaction)
	assert.Equal(t, "tcy", parsed.GetDenom()) // First coin denomination
	assert.Equal(t, "multi-coin transfer", parsed.GetMemo())
	assert.Equal(t, uint64(10), parsed.GetSequence())
}

func TestInvalidProtobufTransaction(t *testing.T) {
	// Test with invalid protobuf data (should return an error now)
	invalidHex := "deadbeefcafebabe1234567890abcdef"

	_, err := ParseThorchainTransaction(invalidHex)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode protobuf transaction")
}
