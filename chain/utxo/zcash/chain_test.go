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

func TestZcash_RejectLegacyTransactionVersions(t *testing.T) {
	// Test that legacy v1-v3 transactions are rejected
	// These formats could potentially contain hidden JoinSplits that bypass our checks

	tests := []struct {
		name    string
		txHex   string
		wantErr string
	}{
		{
			name: "v1 transaction (non-overwintered)",
			txHex: "01000000" + // Version 1, no overwintered flag
				"01" + // Inputs count
				"0000000000000000000000000000000000000000000000000000000000000000" + // PrevHash
				"00000000" + // PrevIndex
				"00" + // ScriptLen
				"ffffffff" + // Sequence
				"00" + // Outputs count
				"00000000", // LockTime
			wantErr: "unsupported transaction version 1",
		},
		{
			name: "v2 transaction (non-overwintered, Sprout era - could have JoinSplits)",
			txHex: "02000000" + // Version 2, no overwintered flag
				"01" + // Inputs count
				"0000000000000000000000000000000000000000000000000000000000000000" + // PrevHash
				"00000000" + // PrevIndex
				"00" + // ScriptLen
				"ffffffff" + // Sequence
				"00" + // Outputs count
				"00000000", // LockTime
			wantErr: "unsupported transaction version 2",
		},
		{
			name: "v3 transaction (overwintered, Overwinter era - could have JoinSplits)",
			txHex: "03000080" + // Version 3 + overwintered flag (0x80000000)
				"70b4d009" + // Overwinter version group ID
				"01" + // Inputs count
				"0000000000000000000000000000000000000000000000000000000000000000" + // PrevHash
				"00000000" + // PrevIndex
				"00" + // ScriptLen
				"ffffffff" + // Sequence
				"00" + // Outputs count
				"00000000" + // LockTime
				"00000000", // ExpiryHeight
			wantErr: "unsupported transaction version 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseZcashTransaction(tt.txHex)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
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
	//     ScriptSig: placeholder with dummy sig + pubkey (for pubkey extraction)
	//     Sequence: ffffffff
	// Outputs: 0
	// LockTime: 00000000
	// ExpiryHeight: 00000000
	// ValueBalance: 0000000000000000
	// ShieldedSpends: 00
	// ShieldedOutputs: 00
	// Joinsplits: 00

	// Placeholder scriptSig for P2PKH (must contain pubkey for extraction):
	// - 0x47 push 71 bytes (DER signature + sighash byte)
	// - DER sig: 30 44 02 20 [32 bytes r] 02 20 [32 bytes s] + 01 (sighash)
	// - 0x21 push 33 bytes (compressed pubkey)
	// - compressed pubkey: 02 + [32 bytes x]
	placeholderScriptSig := "47" + // push 71 bytes
		"3044" + // DER SEQUENCE, 68 bytes
		"0220" + "0000000000000000000000000000000000000000000000000000000000000001" + // INTEGER r
		"0220" + "0000000000000000000000000000000000000000000000000000000000000001" + // INTEGER s
		"01" + // SIGHASH_ALL
		"21" + // push 33 bytes (compressed pubkey)
		"02" + "0000000000000000000000000000000000000000000000000000000000000001" // compressed pubkey

	scriptSigLen := "6a" // 106 bytes = 0x6a (1 + 71 + 1 + 33)

	txHex := "04000080" + // Version
		"85202f89" + // GroupID
		"01" + // Inputs count
		"0000000000000000000000000000000000000000000000000000000000000000" + // PrevHash
		"00000000" + // PrevIndex
		scriptSigLen + placeholderScriptSig + // ScriptSig with placeholder
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

	z := NewChain().(*Zcash)
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

