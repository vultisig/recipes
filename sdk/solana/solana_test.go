package solana

import (
	"encoding/hex"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/mobile-tss-lib/tss"
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
