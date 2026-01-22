package cosmos

import (
	"bytes"
	"crypto/sha256"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	rsdk "github.com/vultisig/recipes/sdk"
)

// Test constants
const (
	testChainID       = "cosmoshub-4"
	testAccountNumber = uint64(12345)
	testSequence      = uint64(1)
)

// createTestCodec creates a codec for testing
func createTestCodec() (codec.Codec, codectypes.InterfaceRegistry) {
	ir := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(ir)
	banktypes.RegisterInterfaces(ir)
	return codec.NewProtoCodec(ir), ir
}

// createTestTxRawAndSignDoc creates matching TxRaw and SignDoc for testing
func createTestTxRawAndSignDoc(t *testing.T) ([]byte, []byte) {
	cdc, _ := createTestCodec()

	// Create a minimal body
	body := &tx.TxBody{
		Messages: []*codectypes.Any{},
		Memo:     "test transaction",
	}
	bodyBytes, err := cdc.Marshal(body)
	require.NoError(t, err)

	// Create auth info
	authInfo := &tx.AuthInfo{
		Fee: &tx.Fee{
			Amount:   cosmostypes.NewCoins(),
			GasLimit: 200000,
		},
	}
	authInfoBytes, err := cdc.Marshal(authInfo)
	require.NoError(t, err)

	// Create TxRaw
	txRaw := &tx.TxRaw{
		BodyBytes:     bodyBytes,
		AuthInfoBytes: authInfoBytes,
		Signatures:    [][]byte{},
	}
	txBytes, err := cdc.Marshal(txRaw)
	require.NoError(t, err)

	// Create matching SignDoc
	signDoc := &tx.SignDoc{
		BodyBytes:     bodyBytes,
		AuthInfoBytes: authInfoBytes,
		ChainId:       testChainID,
		AccountNumber: testAccountNumber,
	}
	signBytes, err := cdc.Marshal(signDoc)
	require.NoError(t, err)

	return txBytes, signBytes
}

// createMismatchedTxRawAndSignDoc creates TxRaw and SignDoc with different body bytes
func createMismatchedTxRawAndSignDoc(t *testing.T) ([]byte, []byte) {
	cdc, _ := createTestCodec()

	// Create body for TxRaw
	body1 := &tx.TxBody{
		Messages: []*codectypes.Any{},
		Memo:     "first transaction",
	}
	bodyBytes1, err := cdc.Marshal(body1)
	require.NoError(t, err)

	// Create different body for SignDoc
	body2 := &tx.TxBody{
		Messages: []*codectypes.Any{},
		Memo:     "different transaction", // Different memo
	}
	bodyBytes2, err := cdc.Marshal(body2)
	require.NoError(t, err)

	// Create auth info
	authInfo := &tx.AuthInfo{
		Fee: &tx.Fee{
			Amount:   cosmostypes.NewCoins(),
			GasLimit: 200000,
		},
	}
	authInfoBytes, err := cdc.Marshal(authInfo)
	require.NoError(t, err)

	// Create TxRaw with first body
	txRaw := &tx.TxRaw{
		BodyBytes:     bodyBytes1,
		AuthInfoBytes: authInfoBytes,
		Signatures:    [][]byte{},
	}
	txBytes, err := cdc.Marshal(txRaw)
	require.NoError(t, err)

	// Create SignDoc with different body
	signDoc := &tx.SignDoc{
		BodyBytes:     bodyBytes2, // Different from txRaw
		AuthInfoBytes: authInfoBytes,
		ChainId:       testChainID,
		AccountNumber: testAccountNumber,
	}
	signBytes, err := cdc.Marshal(signDoc)
	require.NoError(t, err)

	return txBytes, signBytes
}

func TestDeriveSigningHashes_ValidTransactionWithMatchingSignBytes(t *testing.T) {
	sdk := NewSDK(nil)
	txBytes, signBytes := createTestTxRawAndSignDoc(t)

	hashes, err := sdk.DeriveSigningHashes(txBytes, rsdk.DeriveOptions{
		SignBytes: signBytes,
	})
	require.NoError(t, err)

	// Should return exactly 1 hash
	require.Len(t, hashes, 1)

	// Hash should be 32 bytes (SHA256)
	assert.Len(t, hashes[0].Hash, 32)

	// Message should be 32 bytes (same as hash for Cosmos)
	assert.Len(t, hashes[0].Message, 32)

	// For Cosmos, Message and Hash should be identical
	assert.True(t, bytes.Equal(hashes[0].Hash, hashes[0].Message))

	// Verify hash is SHA256 of signBytes
	expectedHash := sha256.Sum256(signBytes)
	assert.True(t, bytes.Equal(hashes[0].Hash, expectedHash[:]))
}

func TestDeriveSigningHashes_MissingSignBytes(t *testing.T) {
	sdk := NewSDK(nil)
	txBytes, _ := createTestTxRawAndSignDoc(t)

	// Call without SignBytes
	_, err := sdk.DeriveSigningHashes(txBytes, rsdk.DeriveOptions{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "signBytes required")
}

func TestDeriveSigningHashes_EmptySignBytes(t *testing.T) {
	sdk := NewSDK(nil)
	txBytes, _ := createTestTxRawAndSignDoc(t)

	// Call with empty SignBytes
	_, err := sdk.DeriveSigningHashes(txBytes, rsdk.DeriveOptions{
		SignBytes: []byte{},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "signBytes required")
}

func TestDeriveSigningHashes_BodyMismatch(t *testing.T) {
	sdk := NewSDK(nil)
	txBytes, signBytes := createMismatchedTxRawAndSignDoc(t)

	_, err := sdk.DeriveSigningHashes(txBytes, rsdk.DeriveOptions{
		SignBytes: signBytes,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "body does not match")
}

func TestDeriveSigningHashes_InvalidTxBytes(t *testing.T) {
	sdk := NewSDK(nil)
	_, signBytes := createTestTxRawAndSignDoc(t)

	testCases := []struct {
		name  string
		input []byte
	}{
		{"random_bytes", []byte{0x01, 0x02, 0x03, 0x04}},
		{"empty_bytes", []byte{}},
		{"invalid_protobuf", []byte{0xFF, 0xFF, 0xFF}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := sdk.DeriveSigningHashes(tc.input, rsdk.DeriveOptions{
				SignBytes: signBytes,
			})
			require.Error(t, err)
		})
	}
}

func TestDeriveSigningHashes_InvalidSignBytes(t *testing.T) {
	sdk := NewSDK(nil)
	txBytes, _ := createTestTxRawAndSignDoc(t)

	testCases := []struct {
		name  string
		input []byte
	}{
		{"random_bytes", []byte{0x01, 0x02, 0x03, 0x04}},
		{"invalid_protobuf", []byte{0xFF, 0xFF, 0xFF}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := sdk.DeriveSigningHashes(txBytes, rsdk.DeriveOptions{
				SignBytes: tc.input,
			})
			require.Error(t, err)
		})
	}
}

func TestDeriveSigningHashes_ConsistentResults(t *testing.T) {
	sdk := NewSDK(nil)
	txBytes, signBytes := createTestTxRawAndSignDoc(t)

	opts := rsdk.DeriveOptions{SignBytes: signBytes}

	// Call multiple times
	hashes1, err := sdk.DeriveSigningHashes(txBytes, opts)
	require.NoError(t, err)

	hashes2, err := sdk.DeriveSigningHashes(txBytes, opts)
	require.NoError(t, err)

	// Results should be identical
	require.Len(t, hashes1, 1)
	require.Len(t, hashes2, 1)
	assert.True(t, bytes.Equal(hashes1[0].Hash, hashes2[0].Hash))
	assert.True(t, bytes.Equal(hashes1[0].Message, hashes2[0].Message))
}

func TestDeriveSigningHashes_DifferentSignBytesProduceDifferentHashes(t *testing.T) {
	sdk := NewSDK(nil)
	cdc, _ := createTestCodec()

	// Create first set
	txBytes1, signBytes1 := createTestTxRawAndSignDoc(t)

	// Create second set with different chain ID
	body := &tx.TxBody{
		Messages: []*codectypes.Any{},
		Memo:     "test transaction",
	}
	bodyBytes, err := cdc.Marshal(body)
	require.NoError(t, err)

	authInfo := &tx.AuthInfo{
		Fee: &tx.Fee{
			Amount:   cosmostypes.NewCoins(),
			GasLimit: 200000,
		},
	}
	authInfoBytes, err := cdc.Marshal(authInfo)
	require.NoError(t, err)

	txRaw2 := &tx.TxRaw{
		BodyBytes:     bodyBytes,
		AuthInfoBytes: authInfoBytes,
		Signatures:    [][]byte{},
	}
	txBytes2, err := cdc.Marshal(txRaw2)
	require.NoError(t, err)

	signDoc2 := &tx.SignDoc{
		BodyBytes:     bodyBytes,
		AuthInfoBytes: authInfoBytes,
		ChainId:       "osmosis-1", // Different chain ID
		AccountNumber: testAccountNumber,
	}
	signBytes2, err := cdc.Marshal(signDoc2)
	require.NoError(t, err)

	// Derive hashes for both
	hashes1, err := sdk.DeriveSigningHashes(txBytes1, rsdk.DeriveOptions{SignBytes: signBytes1})
	require.NoError(t, err)

	hashes2, err := sdk.DeriveSigningHashes(txBytes2, rsdk.DeriveOptions{SignBytes: signBytes2})
	require.NoError(t, err)

	// Hashes should be different (different chain IDs in SignDoc)
	assert.False(t, bytes.Equal(hashes1[0].Hash, hashes2[0].Hash))
}
