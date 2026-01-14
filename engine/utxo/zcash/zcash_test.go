package zcash

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/chain/utxo/zcash"
)

// Sample Zcash v4 transparent transaction (2 inputs, 2 outputs)
// This is a synthetic test transaction for testing purposes
func createTestZcashTransaction(t *testing.T, outputs []*zcash.ZcashOutput) []byte {
	tx := &zcash.ZcashTransaction{
		Version:        4,
		VersionGroupID: 0x892F2085, // Sapling version group
		Inputs: []*zcash.ZcashInput{
			{
				SignatureScript: []byte{},
				Sequence:        0xffffffff,
			},
		},
		Outputs:      outputs,
		LockTime:     0,
		ExpiryHeight: 1000000,
		ValueBalance: 0,
	}

	// Serialize the transaction
	txBytes, err := serializeTestZcashTx(tx)
	assert.NoError(t, err)
	return txBytes
}

func serializeTestZcashTx(tx *zcash.ZcashTransaction) ([]byte, error) {
	// Simple serialization for test purposes
	// This creates a valid v4 Zcash transaction structure
	var data []byte

	// Header with overwintered flag (version | 0x80000000)
	header := uint32(tx.Version) | 0x80000000
	data = append(data, byte(header), byte(header>>8), byte(header>>16), byte(header>>24))

	// Version group ID
	data = append(data, byte(tx.VersionGroupID), byte(tx.VersionGroupID>>8), byte(tx.VersionGroupID>>16), byte(tx.VersionGroupID>>24))

	// Input count (varint)
	data = append(data, byte(len(tx.Inputs)))

	// Inputs
	for _, input := range tx.Inputs {
		// Previous outpoint hash (32 bytes zeros)
		data = append(data, make([]byte, 32)...)
		// Previous outpoint index (4 bytes)
		data = append(data, 0, 0, 0, 0)
		// Script length (varint)
		data = append(data, byte(len(input.SignatureScript)))
		// Script
		data = append(data, input.SignatureScript...)
		// Sequence
		data = append(data, byte(input.Sequence), byte(input.Sequence>>8), byte(input.Sequence>>16), byte(input.Sequence>>24))
	}

	// Output count (varint)
	data = append(data, byte(len(tx.Outputs)))

	// Outputs
	for _, output := range tx.Outputs {
		// Value (8 bytes little endian)
		v := output.Value
		data = append(data, byte(v), byte(v>>8), byte(v>>16), byte(v>>24), byte(v>>32), byte(v>>40), byte(v>>48), byte(v>>56))
		// Script length (varint)
		data = append(data, byte(len(output.PkScript)))
		// Script
		data = append(data, output.PkScript...)
	}

	// Lock time (4 bytes)
	data = append(data, byte(tx.LockTime), byte(tx.LockTime>>8), byte(tx.LockTime>>16), byte(tx.LockTime>>24))

	// Expiry height (4 bytes)
	data = append(data, byte(tx.ExpiryHeight), byte(tx.ExpiryHeight>>8), byte(tx.ExpiryHeight>>16), byte(tx.ExpiryHeight>>24))

	// Value balance (8 bytes)
	data = append(data, 0, 0, 0, 0, 0, 0, 0, 0)

	// Shielded spends count (0)
	data = append(data, 0)

	// Shielded outputs count (0)
	data = append(data, 0)

	// Joinsplits count (0)
	data = append(data, 0)

	return data, nil
}

// Create a P2PKH script for a given pubkey hash
func createP2PKHScript(pubKeyHash []byte) []byte {
	// OP_DUP OP_HASH160 <20 bytes> OP_EQUALVERIFY OP_CHECKSIG
	script := []byte{0x76, 0xa9, 0x14}
	script = append(script, pubKeyHash...)
	script = append(script, 0x88, 0xac)
	return script
}

// Create an OP_RETURN script with data
func createOpReturnScript(data []byte) []byte {
	// OP_RETURN <data_len> <data>
	script := []byte{0x6a, byte(len(data))}
	script = append(script, data...)
	return script
}

func newFixed(index int, address, value string) []*types.ParameterConstraint {
	return []*types.ParameterConstraint{{
		ParameterName: fmt.Sprintf("output_address_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
			Value: &types.Constraint_FixedValue{
				FixedValue: address,
			},
			Required: true,
		},
	}, {
		ParameterName: fmt.Sprintf("output_value_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
			Value: &types.Constraint_FixedValue{
				FixedValue: value,
			},
			Required: true,
		},
	}}
}

func newMax(index int, address, maxValue string) []*types.ParameterConstraint {
	return []*types.ParameterConstraint{{
		ParameterName: fmt.Sprintf("output_address_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
			Value: &types.Constraint_FixedValue{
				FixedValue: address,
			},
			Required: true,
		},
	}, {
		ParameterName: fmt.Sprintf("output_value_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
			Value: &types.Constraint_MaxValue{
				MaxValue: maxValue,
			},
			Required: true,
		},
	}}
}

func newMin(index int, address, minValue string) []*types.ParameterConstraint {
	return []*types.ParameterConstraint{{
		ParameterName: fmt.Sprintf("output_address_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
			Value: &types.Constraint_FixedValue{
				FixedValue: address,
			},
			Required: true,
		},
	}, {
		ParameterName: fmt.Sprintf("output_value_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
			Value: &types.Constraint_MinValue{
				MinValue: minValue,
			},
			Required: true,
		},
	}}
}

func newDataRegexp(index int, dataPattern string) []*types.ParameterConstraint {
	return []*types.ParameterConstraint{{
		ParameterName: fmt.Sprintf("output_data_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_REGEXP,
			Value: &types.Constraint_RegexpValue{
				RegexpValue: dataPattern,
			},
			Required: true,
		},
	}}
}

func TestZcash_Evaluate_Fixed(t *testing.T) {
	// Create a test pubkey hash (20 bytes)
	pubKeyHash, _ := hex.DecodeString("89abcdefabcdef0123456789abcdef0123456789")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    1000000, // 0.01 ZEC in zatoshis
			PkScript: createP2PKHScript(pubKeyHash),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	// Get the expected address for this pubkey hash
	expectedAddr, err := zcash.ExtractZcashAddress(outputs[0].PkScript, zcash.ZcashMainNetParams)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newFixed(0, expectedAddr, "1000000")...)

	err = NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}

func TestZcash_Evaluate_MaxConstraints(t *testing.T) {
	pubKeyHash, _ := hex.DecodeString("89abcdefabcdef0123456789abcdef0123456789")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    1000000,
			PkScript: createP2PKHScript(pubKeyHash),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	expectedAddr, err := zcash.ExtractZcashAddress(outputs[0].PkScript, zcash.ZcashMainNetParams)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newMax(0, expectedAddr, "2000000")...) // max 2M, actual 1M

	err = NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}

func TestZcash_Evaluate_MaxConstraints_ShouldFail(t *testing.T) {
	pubKeyHash, _ := hex.DecodeString("89abcdefabcdef0123456789abcdef0123456789")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    1000000,
			PkScript: createP2PKHScript(pubKeyHash),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	expectedAddr, err := zcash.ExtractZcashAddress(outputs[0].PkScript, zcash.ZcashMainNetParams)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newMax(0, expectedAddr, "500000")...) // max 500k, actual 1M - should fail

	err = NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to compare max values")
}

func TestZcash_Evaluate_MinConstraints(t *testing.T) {
	pubKeyHash, _ := hex.DecodeString("89abcdefabcdef0123456789abcdef0123456789")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    1000000,
			PkScript: createP2PKHScript(pubKeyHash),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	expectedAddr, err := zcash.ExtractZcashAddress(outputs[0].PkScript, zcash.ZcashMainNetParams)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newMin(0, expectedAddr, "500000")...) // min 500k, actual 1M

	err = NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}

func TestZcash_Evaluate_MinConstraints_ShouldFail(t *testing.T) {
	pubKeyHash, _ := hex.DecodeString("89abcdefabcdef0123456789abcdef0123456789")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    1000000,
			PkScript: createP2PKHScript(pubKeyHash),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	expectedAddr, err := zcash.ExtractZcashAddress(outputs[0].PkScript, zcash.ZcashMainNetParams)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newMin(0, expectedAddr, "2000000")...) // min 2M, actual 1M - should fail

	err = NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to compare min values")
}

func TestZcash_Evaluate_WrongAddress_ShouldFail(t *testing.T) {
	pubKeyHash, _ := hex.DecodeString("89abcdefabcdef0123456789abcdef0123456789")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    1000000,
			PkScript: createP2PKHScript(pubKeyHash),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	var params []*types.ParameterConstraint
	params = append(params, newFixed(0, "t1WrongAddress", "1000000")...)

	err := NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to compare fixed values")
}

func TestZcash_Evaluate_WrongValue_ShouldFail(t *testing.T) {
	pubKeyHash, _ := hex.DecodeString("89abcdefabcdef0123456789abcdef0123456789")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    1000000,
			PkScript: createP2PKHScript(pubKeyHash),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	expectedAddr, err := zcash.ExtractZcashAddress(outputs[0].PkScript, zcash.ZcashMainNetParams)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newFixed(0, expectedAddr, "999999")...) // wrong value

	err = NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to compare fixed values")
}

func TestZcash_Evaluate_MismatchedOutputCounts_ShouldFail(t *testing.T) {
	pubKeyHash, _ := hex.DecodeString("89abcdefabcdef0123456789abcdef0123456789")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    1000000,
			PkScript: createP2PKHScript(pubKeyHash),
		},
		{
			Value:    500000,
			PkScript: createP2PKHScript(pubKeyHash),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	expectedAddr, err := zcash.ExtractZcashAddress(outputs[0].PkScript, zcash.ZcashMainNetParams)
	assert.NoError(t, err)

	// Only one output constraint instead of two
	var params []*types.ParameterConstraint
	params = append(params, newFixed(0, expectedAddr, "1000000")...)

	err = NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "output count mismatch")
}

func TestZcash_Evaluate_DataRegexpConstraints(t *testing.T) {
	// Create a transaction with an OP_RETURN output containing MayaChain swap memo
	memoData := []byte("=:ETH.ETH:0x1234567890abcdef1234567890abcdef12345678")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    0,
			PkScript: createOpReturnScript(memoData),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	var params []*types.ParameterConstraint
	params = append(params, newDataRegexp(0, "^=:ETH\\.ETH:0x[a-fA-F0-9]+$")...)

	err := NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}

func TestZcash_Evaluate_DataRegexpConstraints_ShouldFail(t *testing.T) {
	memoData := []byte("=:ETH.ETH:0x1234567890abcdef1234567890abcdef12345678")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    0,
			PkScript: createOpReturnScript(memoData),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	var params []*types.ParameterConstraint
	// Pattern that doesn't match the memo
	params = append(params, newDataRegexp(0, "^=:BTC\\.BTC:bc1[a-z0-9]+$")...)

	err := NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "regexp value constraint failed")
}

func TestZcash_Evaluate_InvalidTxBytes_ShouldFail(t *testing.T) {
	invalidTxBytes := []byte("invalid transaction data")

	pubKeyHash, _ := hex.DecodeString("89abcdefabcdef0123456789abcdef0123456789")
	expectedAddr, _ := zcash.ExtractZcashAddress(createP2PKHScript(pubKeyHash), zcash.ZcashMainNetParams)

	var params []*types.ParameterConstraint
	params = append(params, newFixed(0, expectedAddr, "1000000")...)

	err := NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, invalidTxBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse zcash transaction")
}

func TestZcash_Evaluate_DenyEffect_ShouldFail(t *testing.T) {
	pubKeyHash, _ := hex.DecodeString("89abcdefabcdef0123456789abcdef0123456789")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    1000000,
			PkScript: createP2PKHScript(pubKeyHash),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	expectedAddr, err := zcash.ExtractZcashAddress(outputs[0].PkScript, zcash.ZcashMainNetParams)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newFixed(0, expectedAddr, "1000000")...)

	err = NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_DENY, // DENY effect should fail
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only allow rules supported")
}

func TestZcash_Evaluate_WithTarget_ShouldFail(t *testing.T) {
	pubKeyHash, _ := hex.DecodeString("89abcdefabcdef0123456789abcdef0123456789")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    1000000,
			PkScript: createP2PKHScript(pubKeyHash),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	expectedAddr, err := zcash.ExtractZcashAddress(outputs[0].PkScript, zcash.ZcashMainNetParams)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newFixed(0, expectedAddr, "1000000")...)

	err = NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
		},
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "target type must be unspecified for Zcash")
}

// Test a MayaChain swap transaction structure:
// Output 0: Send to MayaChain vault
// Output 1: OP_RETURN with swap memo
func TestZcash_Evaluate_MayaChainSwapStructure(t *testing.T) {
	// Vault pubkey hash (in real scenario, resolved from MayaChain API)
	vaultPubKeyHash, _ := hex.DecodeString("1234567890abcdef1234567890abcdef12345678")

	// Swap memo
	memoData := []byte("=:ETH.ETH:0xRecipientAddress:0")

	outputs := []*zcash.ZcashOutput{
		{
			Value:    10000000, // 0.1 ZEC
			PkScript: createP2PKHScript(vaultPubKeyHash),
		},
		{
			Value:    0,
			PkScript: createOpReturnScript(memoData),
		},
	}

	txBytes := createTestZcashTransaction(t, outputs)

	vaultAddr, err := zcash.ExtractZcashAddress(outputs[0].PkScript, zcash.ZcashMainNetParams)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	// Output 0: vault address with max amount
	params = append(params, newMax(0, vaultAddr, "100000000")...) // max 1 ZEC
	// Output 1: OP_RETURN with swap memo pattern
	params = append(params, newDataRegexp(1, "^=:ETH\\.ETH:0x[a-zA-Z0-9]+:.*")...)

	err = NewZcash().Evaluate(&types.Rule{
		Resource:             "zcash.zec.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}
