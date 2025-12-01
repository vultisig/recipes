package zcash

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vultisig/mobile-tss-lib/tss"
)

func TestZcash_ChainMethods(t *testing.T) {
	z := NewChain()

	assert.Equal(t, "zcash", z.ID())
	assert.Equal(t, "Zcash", z.Name())
	assert.Equal(t, []string{"zec"}, z.SupportedProtocols())
	assert.Contains(t, z.Description(), "privacy-focused")

	// Test GetProtocol
	p, err := z.GetProtocol("zec")
	assert.NoError(t, err)
	assert.Equal(t, "zec", p.ID())

	_, err = z.GetProtocol("invalid")
	assert.Error(t, err)
}

func TestZcash_ComputeTxHash(t *testing.T) {
	// Create a simple transaction to test ComputeTxHash
	// We'll manually construct a v4 transaction bytes
	// Header: 04000080 (v4 + overwintered)
	// GroupId: 85202F89
	// Inputs: 1
	//   Input 0:
	//     PrevHash: [32 bytes 00]
	//     PrevIndex: 00000000
	//     ScriptLen: 00
	//     Sequence: ffffffff
	// Outputs: 0
	// LockTime: 00000000
	// ExpiryHeight: 00000000
	// ValueBalance: 0000000000000000
	// ShieldedSpends: 00
	// ShieldedOutputs: 00
	// Joinsplits: 00

	txHex := "04000080" + // Version
		"85202f89" + // GroupID
		"01" + // Inputs count
		"0000000000000000000000000000000000000000000000000000000000000000" + // PrevHash
		"00000000" + // PrevIndex
		"00" + // ScriptLen
		"ffffffff" + // Sequence
		"00" + // Outputs count
		"00000000" + // LockTime
		"00000000" + // ExpiryHeight
		"0000000000000000" + // ValueBalance
		"00" + // Spends
		"00" + // Outputs
		"00" // Joinsplits

	txBytes, err := hex.DecodeString(txHex)
	assert.NoError(t, err)

	// Create dummy signatures
	sigs := []tss.KeysignResponse{
		{
			R: "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20",
			S: "201f1e1d1c1b1a191817161514131211100f0e0d0c0b0a090807060504030201",
		},
	}

	z := NewChain()
	hash, err := z.ComputeTxHash(txBytes, sigs)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	// To verify correct insertion, we could try to verify the scriptSig length in the resulting tx
	// but we don't have easy access to the intermediate signed bytes here, just the hash.
	// However, if ComputeTxHash didn't error, it means it successfully deserialized, applied sigs, and reserialized.

	// Test mismatch count
	_, err = z.ComputeTxHash(txBytes, []tss.KeysignResponse{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "input count")
}

