package dogecoin

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

// Dogecoin uses D... prefix for P2PKH addresses on mainnet (prefix byte 0x1e)
// This matches the format used by THORChain for DOGE vault addresses.

func TestDogecoin_Supports(t *testing.T) {
	doge := NewDogecoin()

	assert.True(t, doge.Supports(common.Dogecoin))
	assert.False(t, doge.Supports(common.Bitcoin))
	assert.False(t, doge.Supports(common.Litecoin))
	assert.False(t, doge.Supports(common.BitcoinCash))
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

// createP2PKHScript creates a P2PKH script for a given 20-byte pubkey hash.
// Format: OP_DUP OP_HASH160 <20 bytes> OP_EQUALVERIFY OP_CHECKSIG
func createP2PKHScript(pubKeyHash []byte) []byte {
	script := []byte{0x76, 0xa9, 0x14} // OP_DUP OP_HASH160 PUSH20
	script = append(script, pubKeyHash...)
	script = append(script, 0x88, 0xac) // OP_EQUALVERIFY OP_CHECKSIG
	return script
}

// createTestTransaction creates a test Dogecoin transaction with P2PKH outputs.
func createTestTransaction(t *testing.T, outputs []struct {
	pubKeyHash []byte
	value      int64
}) []byte {
	tx := &wire.MsgTx{
		Version: 1,
		TxIn: []*wire.TxIn{
			{
				PreviousOutPoint: wire.OutPoint{
					Hash:  [32]byte{},
					Index: 0,
				},
				SignatureScript: []byte{},
				Sequence:        0xffffffff,
			},
		},
		TxOut:    make([]*wire.TxOut, 0, len(outputs)),
		LockTime: 0,
	}

	for _, out := range outputs {
		tx.TxOut = append(tx.TxOut, &wire.TxOut{
			Value:    out.value,
			PkScript: createP2PKHScript(out.pubKeyHash),
		})
	}

	var buf bytes.Buffer
	err := tx.BtcEncode(&buf, wire.ProtocolVersion, wire.BaseEncoding)
	require.NoError(t, err)

	return buf.Bytes()
}

// Dogecoin P2PKH addresses use prefix 0x1e (D prefix)
// Test pubkey hashes that produce valid D-addresses
var testPubKeyHash1 = []byte{
	0x77, 0xbf, 0xf2, 0x0c, 0x60, 0xe5, 0x22, 0xdf,
	0xaa, 0x91, 0x3e, 0xb0, 0x12, 0x3f, 0x6b, 0x3f,
	0xf7, 0x60, 0xd0, 0xb0,
}

var testPubKeyHash2 = []byte{
	0x88, 0xce, 0xe3, 0x1d, 0x71, 0xf6, 0x33, 0xe0,
	0xbb, 0xa2, 0x4e, 0xc1, 0x23, 0x4e, 0x7c, 0x4e,
	0xe8, 0x71, 0xe1, 0xc1,
}

func TestDogecoin_Evaluate_Fixed(t *testing.T) {
	txBytes := createTestTransaction(t, []struct {
		pubKeyHash []byte
		value      int64
	}{
		{testPubKeyHash1, 100000000000}, // 1000 DOGE in koinu
		{testPubKeyHash2, 50000000000},  // 500 DOGE
	})

	var params []*types.ParameterConstraint
	params = append(params, newFixed(0, "DG4GthBCBJQwRtGGcQFuSy7EznJsJxU53W", "100000000000")...)
	params = append(params, newFixed(1, "DHcUH7uKpNCHeiquAXtcYmCPmVX2JNE1qp", "50000000000")...)

	err := NewDogecoin().Evaluate(&types.Rule{
		Resource:             "dogecoin.doge.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}

func TestDogecoin_Evaluate_MaxConstraints(t *testing.T) {
	txBytes := createTestTransaction(t, []struct {
		pubKeyHash []byte
		value      int64
	}{
		{testPubKeyHash1, 100000000000}, // 1000 DOGE
	})

	var params []*types.ParameterConstraint
	params = append(params, newMax(0, "DG4GthBCBJQwRtGGcQFuSy7EznJsJxU53W", "200000000000")...) // max 2000 DOGE

	err := NewDogecoin().Evaluate(&types.Rule{
		Resource:             "dogecoin.doge.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}

func TestDogecoin_Evaluate_MaxConstraints_ShouldFail(t *testing.T) {
	txBytes := createTestTransaction(t, []struct {
		pubKeyHash []byte
		value      int64
	}{
		{testPubKeyHash1, 100000000000}, // 1000 DOGE
	})

	var params []*types.ParameterConstraint
	params = append(params, newMax(0, "DG4GthBCBJQwRtGGcQFuSy7EznJsJxU53W", "50000000000")...) // max 500 DOGE - should fail

	err := NewDogecoin().Evaluate(&types.Rule{
		Resource:             "dogecoin.doge.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "max value constraint failed")
}

func TestDogecoin_Evaluate_MinConstraints(t *testing.T) {
	txBytes := createTestTransaction(t, []struct {
		pubKeyHash []byte
		value      int64
	}{
		{testPubKeyHash1, 100000000000}, // 1000 DOGE
	})

	var params []*types.ParameterConstraint
	params = append(params, newMin(0, "DG4GthBCBJQwRtGGcQFuSy7EznJsJxU53W", "50000000000")...) // min 500 DOGE

	err := NewDogecoin().Evaluate(&types.Rule{
		Resource:             "dogecoin.doge.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}

func TestDogecoin_Evaluate_MinConstraints_ShouldFail(t *testing.T) {
	txBytes := createTestTransaction(t, []struct {
		pubKeyHash []byte
		value      int64
	}{
		{testPubKeyHash1, 100000000000}, // 1000 DOGE
	})

	var params []*types.ParameterConstraint
	params = append(params, newMin(0, "DG4GthBCBJQwRtGGcQFuSy7EznJsJxU53W", "200000000000")...) // min 2000 DOGE - should fail

	err := NewDogecoin().Evaluate(&types.Rule{
		Resource:             "dogecoin.doge.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "min value constraint failed")
}

func TestDogecoin_Evaluate_WrongAddress_ShouldFail(t *testing.T) {
	txBytes := createTestTransaction(t, []struct {
		pubKeyHash []byte
		value      int64
	}{
		{testPubKeyHash1, 100000000000},
	})

	var params []*types.ParameterConstraint
	params = append(params, newFixed(0, "DWrongAddressXXXXXXXXXXXXXXXXX", "100000000000")...)

	err := NewDogecoin().Evaluate(&types.Rule{
		Resource:             "dogecoin.doge.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fixed value constraint failed")
}

func TestDogecoin_Evaluate_MismatchedOutputCounts_ShouldFail(t *testing.T) {
	txBytes := createTestTransaction(t, []struct {
		pubKeyHash []byte
		value      int64
	}{
		{testPubKeyHash1, 100000000000},
		{testPubKeyHash2, 50000000000},
	})

	// Only one output constraint instead of two
	var params []*types.ParameterConstraint
	params = append(params, newFixed(0, "DG4GthBCBJQwRtGGcQFuSy7EznJsJxU53W", "100000000000")...)

	err := NewDogecoin().Evaluate(&types.Rule{
		Resource:             "dogecoin.doge.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "output count mismatch")
}

func TestDogecoin_Evaluate_InvalidTxBytes_ShouldFail(t *testing.T) {
	invalidTxBytes := []byte("invalid transaction data")

	var params []*types.ParameterConstraint
	params = append(params, newFixed(0, "DG4GthBCBJQwRtGGcQFuSy7EznJsJxU53W", "100000000000")...)

	err := NewDogecoin().Evaluate(&types.Rule{
		Resource:             "dogecoin.doge.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, invalidTxBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse dogecoin transaction")
}

// Test OP_RETURN data constraints (for THORChain swap memos)
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

func createOpReturnTransaction(t *testing.T, data string) []byte {
	pkScript, err := txscript.NullDataScript([]byte(data))
	assert.NoError(t, err)

	tx := &wire.MsgTx{
		Version: 1,
		TxIn: []*wire.TxIn{
			{
				PreviousOutPoint: wire.OutPoint{
					Hash:  [32]byte{},
					Index: 0,
				},
				SignatureScript: []byte{},
				Sequence:        0xffffffff,
			},
		},
		TxOut: []*wire.TxOut{
			{
				Value:    0,
				PkScript: pkScript,
			},
		},
		LockTime: 0,
	}

	var buf bytes.Buffer
	err = tx.BtcEncode(&buf, wire.ProtocolVersion, wire.BaseEncoding)
	assert.NoError(t, err)

	return buf.Bytes()
}

func TestDogecoin_Evaluate_DataRegexpConstraints(t *testing.T) {
	// THORChain swap memo format
	txBytes := createOpReturnTransaction(t, "=:e:0x86d526d6624AbC0178cF7296cD538Ecc080A95F1:0/1/0")

	var params []*types.ParameterConstraint
	params = append(params, newDataRegexp(
		0,
		"^=:e:0x86d526d6624AbC0178cF7296cD538Ecc080A95F1:.",
	)...)

	err := NewDogecoin().Evaluate(&types.Rule{
		Resource:             "dogecoin.doge.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}

func TestDogecoin_Evaluate_DataRegexpConstraints_ShouldFail(t *testing.T) {
	txBytes := createOpReturnTransaction(t, "=:e:0x86d526d6624AbC0178cF7296cD538Ecc080A95F1:0/1/0")

	var params []*types.ParameterConstraint
	params = append(params, newDataRegexp(
		0,
		"^=:e:0x86d526d6624AbC0178cF7296cD538Ecc08088888:.",
	)...)

	err := NewDogecoin().Evaluate(&types.Rule{
		Resource:             "dogecoin.doge.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "regexp value constraint failed")
}
