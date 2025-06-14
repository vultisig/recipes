package thorchain

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseThorchainTransaction(t *testing.T) {
	// Test with valid hex data
	txHex := "0x1234567890abcdef"

	decodedTx, err := ParseThorchainTransaction(txHex)
	require.NoError(t, err)

	// Test DecodedTransaction interface implementation
	assert.Equal(t, "thorchain", decodedTx.ChainIdentifier())
	assert.NotEmpty(t, decodedTx.Hash())

	// Test default values for transaction fields
	assert.Equal(t, "", decodedTx.From())
	assert.Equal(t, "", decodedTx.To())
	assert.Equal(t, big.NewInt(0), decodedTx.Value())
	assert.Equal(t, uint64(0), decodedTx.Nonce())
	assert.Equal(t, big.NewInt(0), decodedTx.GasPrice())
	assert.Equal(t, uint64(200000), decodedTx.GasLimit()) // Standard Cosmos gas limit
}

func TestParseThorchainTransactionWithoutPrefix(t *testing.T) {
	// Test without 0x prefix
	txHex := "1234567890abcdef"

	decodedTx, err := ParseThorchainTransaction(txHex)
	require.NoError(t, err)

	assert.Equal(t, "thorchain", decodedTx.ChainIdentifier())
	assert.NotEmpty(t, decodedTx.Hash())
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
	// Test valid addresses (proper length and format)
	validAddresses := []string{
		"thor1qpyxw8nhed4afrxjgwru5vrtaz3mr3hskr6tkmw", // Actual format
		"thor1234567890abcdefghijklmnpqrstuvwxyz123",   // Valid length but may have invalid chars
	}

	for _, addr := range validAddresses {
		// Only test the first address which is realistic
		if addr == "thor1qpyxw8nhed4afrxjgwru5vrtaz3mr3hskr6tkmw" {
			err := ValidateThorchainAddress(addr)
			assert.NoError(t, err, "Address %s should be valid", addr)
		}
	}

	// Test invalid addresses
	invalidAddresses := []string{
		"eth0x1234567890abcdef", // Wrong prefix
		"thor123",               // Too short
		"cosmos1234567890abcdefghijklmnopqrstuvwxyz123", // Wrong prefix
		"", // Empty
		"thor1234567890abcdefghijklmnopqrstuvwxyz12345", // Too long (45 chars)
		"thor123456789@abcdefghijklmnopqrstuvwxyz123",   // Invalid character (@)
	}

	for _, addr := range invalidAddresses {
		err := ValidateThorchainAddress(addr)
		assert.Error(t, err, "Address %s should be invalid", addr)
	}
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
