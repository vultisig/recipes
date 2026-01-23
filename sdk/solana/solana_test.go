package solana

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/mobile-tss-lib/tss"
	rsdk "github.com/vultisig/recipes/sdk"
)

// createTestUnsignedTransaction
// https://solscan.io/tx/pmfbrM7c3mdmTN3RbJFRBkyCYfdEEShTtPf9tEuPC4s1q3YEWEC6yzjK6XLUWJCpSx3qUjw8nCCbpAdbGU3Y8ip
const (
	realMainnetRValue = "012931f32f618f1b67345f990de57283556792891cdbf47ac35566a25992ada8"
	realMainnetSValue = "585c7143c9c5ffe37cbfd09399ec51e6c0179cd5a5e71ce95394e192d0d8f93f"
)

func createTestUnsignedTransaction(t *testing.T) []byte {
	fromPubKey := solana.MustPublicKeyFromBase58("GJvewfRjqTUPtx6WsBSUnaFbdgXwgXnWfpDyLm65T4YA")
	toPubKey := solana.MustPublicKeyFromBase58("DttWaMuVvTiduZRnguLF7jNxTgiMBZ1hyAumKUiL2KRL")
	computeBudgetProgram := solana.MustPublicKeyFromBase58("ComputeBudget111111111111111111111111111111")

	recentBlockhash := solana.MustHashFromBase58("DM5gm4dykyivZ8itzpzEn2H4Vd5sPQerajcw9XL6qt7g")

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			solana.NewInstruction(
				computeBudgetProgram,
				solana.AccountMetaSlice{},
				[]byte{0x03, 0x13, 0x88, 0x00, 0x00},
			),
			solana.NewInstruction(
				solana.SystemProgramID,
				solana.AccountMetaSlice{
					{PublicKey: fromPubKey, IsSigner: true, IsWritable: true},
					{PublicKey: toPubKey, IsSigner: false, IsWritable: true},
				},
				[]byte{0x02, 0x00, 0x00, 0x00, 0x53, 0x16, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			),
		},
		recentBlockhash,
		solana.TransactionPayer(fromPubKey),
	)
	require.NoError(t, err)

	fullTx := solana.Transaction{
		Signatures: make([]solana.Signature, tx.Message.Header.NumRequiredSignatures),
		Message:    tx.Message,
	}

	fullTxBytes, err := fullTx.MarshalBinary()
	require.NoError(t, err)

	return fullTxBytes
}

func TestSDK_Sign_WithValidSignature(t *testing.T) {
	sdk := &SDK{}
	unsignedTx := createTestUnsignedTransaction(t)

	tx, err := solana.TransactionFromBytes(unsignedTx)
	require.NoError(t, err)

	messageBytes, err := tx.Message.MarshalBinary()
	require.NoError(t, err)

	derivedKey := sdk.deriveKeyFromMessage(messageBytes)

	signatures := map[string]tss.KeysignResponse{
		derivedKey: {
			R: realMainnetRValue,
			S: realMainnetSValue,
		},
	}

	signedTxBytes, err := sdk.Sign(unsignedTx, signatures)
	require.NoError(t, err)
	assert.NotEmpty(t, signedTxBytes)

	signedTx, err := solana.TransactionFromBytes(signedTxBytes)
	require.NoError(t, err)
	assert.NotEmpty(t, signedTx.Signatures)

	expectedSig := make([]byte, 64)
	rBytes, _ := hex.DecodeString(realMainnetRValue)
	sBytes, _ := hex.DecodeString(realMainnetSValue)
	copy(expectedSig[0:32], rBytes)
	copy(expectedSig[32:64], sBytes)

	assert.Equal(t, expectedSig, signedTx.Signatures[0][:])
}

func TestDeriveSigningHashes_ValidTransaction(t *testing.T) {
	sdk := &SDK{}
	unsignedTx := createTestUnsignedTransaction(t)

	hashes, err := sdk.DeriveSigningHashes(unsignedTx, rsdk.DeriveOptions{})
	require.NoError(t, err)

	// Should return exactly 1 hash
	require.Len(t, hashes, 1)

	// Hash should be 32 bytes (SHA256 of message bytes)
	assert.Len(t, hashes[0].Hash, 32)

	// Message should contain the serialized message bytes
	assert.NotEmpty(t, hashes[0].Message)

	// Hash should be SHA256 of Message
	expectedHash := sha256.Sum256(hashes[0].Message)
	assert.True(t, bytes.Equal(hashes[0].Hash, expectedHash[:]))
}

func TestDeriveSigningHashes_MessageMatchesManualExtraction(t *testing.T) {
	sdk := &SDK{}
	unsignedTx := createTestUnsignedTransaction(t)

	// Get hashes using DeriveSigningHashes
	hashes, err := sdk.DeriveSigningHashes(unsignedTx, rsdk.DeriveOptions{})
	require.NoError(t, err)

	// Manually extract message bytes for comparison
	tx, err := solana.TransactionFromBytes(unsignedTx)
	require.NoError(t, err)
	messageBytes, err := tx.Message.MarshalBinary()
	require.NoError(t, err)

	// The Message field should match the serialized message
	assert.True(t, bytes.Equal(hashes[0].Message, messageBytes))
}

func TestDeriveSigningHashes_InvalidTransaction(t *testing.T) {
	sdk := &SDK{}

	testCases := []struct {
		name  string
		input []byte
	}{
		{"random_bytes", []byte{0x01, 0x02, 0x03, 0x04}},
		{"empty_bytes", []byte{}},
		{"truncated_tx", []byte{0x01, 0x00, 0x00}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := sdk.DeriveSigningHashes(tc.input, rsdk.DeriveOptions{})
			require.Error(t, err)
		})
	}
}

func TestDeriveSigningHashes_ConsistentResults(t *testing.T) {
	sdk := &SDK{}
	unsignedTx := createTestUnsignedTransaction(t)

	// Call multiple times
	hashes1, err := sdk.DeriveSigningHashes(unsignedTx, rsdk.DeriveOptions{})
	require.NoError(t, err)

	hashes2, err := sdk.DeriveSigningHashes(unsignedTx, rsdk.DeriveOptions{})
	require.NoError(t, err)

	// Results should be identical
	require.Len(t, hashes1, 1)
	require.Len(t, hashes2, 1)
	assert.True(t, bytes.Equal(hashes1[0].Hash, hashes2[0].Hash))
	assert.True(t, bytes.Equal(hashes1[0].Message, hashes2[0].Message))
}

func TestDeriveSigningHashes_DifferentTransactions_DifferentHashes(t *testing.T) {
	sdk := &SDK{}

	// Create first transaction
	tx1Bytes := createTestUnsignedTransaction(t)

	// Create second transaction with different parameters
	fromPubKey := solana.MustPublicKeyFromBase58("GJvewfRjqTUPtx6WsBSUnaFbdgXwgXnWfpDyLm65T4YA")
	toPubKey := solana.MustPublicKeyFromBase58("DttWaMuVvTiduZRnguLF7jNxTgiMBZ1hyAumKUiL2KRL")

	// Use a different blockhash
	differentBlockhash := solana.MustHashFromBase58("4vJ9JU1bJJE96FWSJKvHsmmFADCg4gpZQff4P3bkLKi")

	tx2, err := solana.NewTransaction(
		[]solana.Instruction{
			solana.NewInstruction(
				solana.SystemProgramID,
				solana.AccountMetaSlice{
					{PublicKey: fromPubKey, IsSigner: true, IsWritable: true},
					{PublicKey: toPubKey, IsSigner: false, IsWritable: true},
				},
				[]byte{0x02, 0x00, 0x00, 0x00, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // Different amount
			),
		},
		differentBlockhash,
		solana.TransactionPayer(fromPubKey),
	)
	require.NoError(t, err)

	fullTx2 := solana.Transaction{
		Signatures: make([]solana.Signature, tx2.Message.Header.NumRequiredSignatures),
		Message:    tx2.Message,
	}
	tx2Bytes, err := fullTx2.MarshalBinary()
	require.NoError(t, err)

	// Derive hashes for both
	hashes1, err := sdk.DeriveSigningHashes(tx1Bytes, rsdk.DeriveOptions{})
	require.NoError(t, err)

	hashes2, err := sdk.DeriveSigningHashes(tx2Bytes, rsdk.DeriveOptions{})
	require.NoError(t, err)

	// Hashes should be different
	assert.False(t, bytes.Equal(hashes1[0].Hash, hashes2[0].Hash))
	assert.False(t, bytes.Equal(hashes1[0].Message, hashes2[0].Message))
}
