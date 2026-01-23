package tron

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

// hexToBase58 converts a hex-encoded TRON address to base58check format.
// TRON addresses are 21 bytes: 1 byte version (0x41) + 20 bytes address.
func hexToBase58(hexAddr string) string {
	data, err := hex.DecodeString(hexAddr)
	if err != nil || len(data) != 21 {
		return hexAddr // Return as-is if invalid
	}
	firstHash := sha256.Sum256(data)
	secondHash := sha256.Sum256(firstHash[:])
	checksum := secondHash[:4]
	result := make([]byte, len(data)+4)
	copy(result, data)
	copy(result[len(data):], checksum)
	return base58.Encode(result)
}

func TestNewTron(t *testing.T) {
	tron := NewTron()
	assert.NotNil(t, tron)
	assert.NotNil(t, tron.chain)
}

func TestTron_Supports(t *testing.T) {
	tron := NewTron()

	tests := []struct {
		name     string
		chain    common.Chain
		expected bool
	}{
		{
			name:     "supports Tron",
			chain:    common.Tron,
			expected: true,
		},
		{
			name:     "does not support Bitcoin",
			chain:    common.Bitcoin,
			expected: false,
		},
		{
			name:     "does not support Ethereum",
			chain:    common.Ethereum,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tron.Supports(tt.chain)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTron_Evaluate_DenyRule(t *testing.T) {
	tron := NewTron()
	rule := &types.Rule{
		Effect: types.Effect_EFFECT_DENY,
	}

	err := tron.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only allow rules supported")
}

func TestTron_Evaluate_InvalidResource(t *testing.T) {
	tron := NewTron()
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "invalid-resource",
	}

	err := tron.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse rule resource")
}

func TestTron_Evaluate_InvalidTransactionData(t *testing.T) {
	tron := NewTron()
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "tron.trx.transfer",
	}

	err := tron.Evaluate(rule, []byte("invalid-tx-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse TRON transaction")
}

func TestTron_Evaluate_EmptyTransaction(t *testing.T) {
	tron := NewTron()
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "tron.trx.transfer",
	}

	err := tron.Evaluate(rule, []byte{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "empty transaction data")
}

// buildTronTransferTx builds a valid TRON TransferContract transaction in protobuf format.
// This follows the TRON protobuf schema for raw_data.
func buildTronTransferTx(ownerAddr, toAddr []byte, amount int64) []byte {
	// Build TransferContract value (inner message)
	// Field 1: owner_address (bytes), Field 2: to_address (bytes), Field 3: amount (varint)
	transferValue := []byte{}
	// owner_address: field 1, wire type 2 (length-delimited)
	transferValue = append(transferValue, 0x0a) // (1 << 3) | 2
	transferValue = append(transferValue, byte(len(ownerAddr)))
	transferValue = append(transferValue, ownerAddr...)
	// to_address: field 2, wire type 2 (length-delimited)
	transferValue = append(transferValue, 0x12) // (2 << 3) | 2
	transferValue = append(transferValue, byte(len(toAddr)))
	transferValue = append(transferValue, toAddr...)
	// amount: field 3, wire type 0 (varint)
	transferValue = append(transferValue, 0x18) // (3 << 3) | 0
	transferValue = append(transferValue, encodeVarint(uint64(amount))...)

	// Build parameter Any message
	// Field 1: type_url (string), Field 2: value (bytes)
	typeUrl := "type.googleapis.com/protocol.TransferContract"
	parameter := []byte{}
	// type_url: field 1, wire type 2
	parameter = append(parameter, 0x0a) // (1 << 3) | 2
	parameter = append(parameter, encodeVarint(uint64(len(typeUrl)))...)
	parameter = append(parameter, []byte(typeUrl)...)
	// value: field 2, wire type 2
	parameter = append(parameter, 0x12) // (2 << 3) | 2
	parameter = append(parameter, encodeVarint(uint64(len(transferValue)))...)
	parameter = append(parameter, transferValue...)

	// Build contract message
	// Field 1: type (enum=1 for TransferContract), Field 2: parameter (Any)
	contract := []byte{}
	// type: field 1, wire type 0 (varint)
	contract = append(contract, 0x08) // (1 << 3) | 0
	contract = append(contract, 0x01) // TransferContract = 1
	// parameter: field 2, wire type 2
	contract = append(contract, 0x12) // (2 << 3) | 2
	contract = append(contract, encodeVarint(uint64(len(parameter)))...)
	contract = append(contract, parameter...)

	// Build raw_data message
	// Field 1: ref_block_bytes, Field 4: ref_block_hash, Field 8: expiration,
	// Field 11: contract, Field 14: timestamp
	rawData := []byte{}

	// ref_block_bytes: field 1, wire type 2
	refBlockBytes := []byte{0x12, 0x34}
	rawData = append(rawData, 0x0a) // (1 << 3) | 2
	rawData = append(rawData, byte(len(refBlockBytes)))
	rawData = append(rawData, refBlockBytes...)

	// ref_block_hash: field 4, wire type 2
	refBlockHash := []byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11}
	rawData = append(rawData, 0x22) // (4 << 3) | 2
	rawData = append(rawData, byte(len(refBlockHash)))
	rawData = append(rawData, refBlockHash...)

	// expiration: field 8, wire type 0
	expiration := uint64(1700000000000) // Some future timestamp in ms
	rawData = append(rawData, 0x40)     // (8 << 3) | 0
	rawData = append(rawData, encodeVarint(expiration)...)

	// contract: field 11, wire type 2
	rawData = append(rawData, 0x5a) // (11 << 3) | 2
	rawData = append(rawData, encodeVarint(uint64(len(contract)))...)
	rawData = append(rawData, contract...)

	// timestamp: field 14, wire type 0
	timestamp := uint64(1699999990000) // Current timestamp in ms
	rawData = append(rawData, 0x70)    // (14 << 3) | 0
	rawData = append(rawData, encodeVarint(timestamp)...)

	return rawData
}

// encodeVarint encodes a uint64 as a protobuf varint
func encodeVarint(v uint64) []byte {
	var buf []byte
	for v >= 0x80 {
		buf = append(buf, byte(v)|0x80)
		v >>= 7
	}
	buf = append(buf, byte(v))
	return buf
}

func TestTron_Evaluate_Success_ValidTransfer(t *testing.T) {
	tron := NewTron()

	// Build a valid TRON transfer transaction
	// TRON addresses are 21 bytes (0x41 prefix + 20 bytes)
	ownerAddr, _ := hex.DecodeString("41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	toAddr, _ := hex.DecodeString("41b614f803b6fd780986a42c78ec9c7f77e6ded13d")
	amount := int64(1000000) // 1 TRX = 1,000,000 SUN

	txBytes := buildTronTransferTx(ownerAddr, toAddr, amount)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "tron.trx.transfer",
	}

	err := tron.Evaluate(rule, txBytes)
	require.NoError(t, err)
}

func TestTron_Evaluate_Success_WithTargetAddress(t *testing.T) {
	tron := NewTron()

	ownerAddr, _ := hex.DecodeString("41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	toAddr, _ := hex.DecodeString("41b614f803b6fd780986a42c78ec9c7f77e6ded13d")
	amount := int64(5000000) // 5 TRX

	txBytes := buildTronTransferTx(ownerAddr, toAddr, amount)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "tron.trx.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: hexToBase58(hex.EncodeToString(toAddr)),
			},
		},
	}

	err := tron.Evaluate(rule, txBytes)
	require.NoError(t, err)
}

func TestTron_Evaluate_Failure_TargetMismatch(t *testing.T) {
	tron := NewTron()

	ownerAddr, _ := hex.DecodeString("41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	toAddr, _ := hex.DecodeString("41b614f803b6fd780986a42c78ec9c7f77e6ded13d")
	differentAddr := hexToBase58("41c614f803b6fd780986a42c78ec9c7f77e6ded13e")
	amount := int64(5000000)

	txBytes := buildTronTransferTx(ownerAddr, toAddr, amount)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "tron.trx.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: differentAddr,
			},
		},
	}

	err := tron.Evaluate(rule, txBytes)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "target address mismatch")
}

func TestTron_Evaluate_Success_WithAmountConstraint(t *testing.T) {
	tron := NewTron()

	ownerAddr, _ := hex.DecodeString("41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	toAddr, _ := hex.DecodeString("41b614f803b6fd780986a42c78ec9c7f77e6ded13d")
	amount := int64(1000000) // 1 TRX

	txBytes := buildTronTransferTx(ownerAddr, toAddr, amount)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "tron.trx.transfer",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	err := tron.Evaluate(rule, txBytes)
	require.NoError(t, err)
}

func TestTron_Evaluate_Failure_AmountConstraintViolation(t *testing.T) {
	tron := NewTron()

	ownerAddr, _ := hex.DecodeString("41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	toAddr, _ := hex.DecodeString("41b614f803b6fd780986a42c78ec9c7f77e6ded13d")
	actualAmount := int64(5000000) // 5 TRX
	expectedAmount := "1000000"    // 1 TRX

	txBytes := buildTronTransferTx(ownerAddr, toAddr, actualAmount)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "tron.trx.transfer",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: expectedAmount,
					},
				},
			},
		},
	}

	err := tron.Evaluate(rule, txBytes)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to validate parameter constraints")
}

func TestTron_Evaluate_Success_WithMaxAmountConstraint(t *testing.T) {
	tron := NewTron()

	ownerAddr, _ := hex.DecodeString("41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	toAddr, _ := hex.DecodeString("41b614f803b6fd780986a42c78ec9c7f77e6ded13d")
	amount := int64(5000000) // 5 TRX

	txBytes := buildTronTransferTx(ownerAddr, toAddr, amount)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "tron.trx.transfer",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "10000000", // Max 10 TRX
					},
				},
			},
		},
	}

	err := tron.Evaluate(rule, txBytes)
	require.NoError(t, err)
}

func TestTron_Evaluate_Failure_MaxAmountExceeded(t *testing.T) {
	tron := NewTron()

	ownerAddr, _ := hex.DecodeString("41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	toAddr, _ := hex.DecodeString("41b614f803b6fd780986a42c78ec9c7f77e6ded13d")
	amount := int64(50000000) // 50 TRX

	txBytes := buildTronTransferTx(ownerAddr, toAddr, amount)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "tron.trx.transfer",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "10000000", // Max 10 TRX
					},
				},
			},
		},
	}

	err := tron.Evaluate(rule, txBytes)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to validate parameter constraints")
}

func TestTron_Evaluate_Success_WithRecipientConstraint(t *testing.T) {
	tron := NewTron()

	ownerAddr, _ := hex.DecodeString("41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	toAddr, _ := hex.DecodeString("41b614f803b6fd780986a42c78ec9c7f77e6ded13d")
	amount := int64(1000000)

	txBytes := buildTronTransferTx(ownerAddr, toAddr, amount)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "tron.trx.transfer",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: hexToBase58(hex.EncodeToString(toAddr)),
					},
				},
			},
		},
	}

	err := tron.Evaluate(rule, txBytes)
	require.NoError(t, err)
}

func TestTron_Evaluate_Failure_RecipientConstraintViolation(t *testing.T) {
	tron := NewTron()

	ownerAddr, _ := hex.DecodeString("41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	toAddr, _ := hex.DecodeString("41b614f803b6fd780986a42c78ec9c7f77e6ded13d")
	differentRecipient := hexToBase58("41c614f803b6fd780986a42c78ec9c7f77e6ded13e")
	amount := int64(1000000)

	txBytes := buildTronTransferTx(ownerAddr, toAddr, amount)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "tron.trx.transfer",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: differentRecipient,
					},
				},
			},
		},
	}

	err := tron.Evaluate(rule, txBytes)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to validate parameter constraints")
}
