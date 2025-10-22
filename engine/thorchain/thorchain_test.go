package thorchain

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	vtypes "github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

func TestNewThorchain(t *testing.T) {
	thorchain := NewThorchain()
	assert.NotNil(t, thorchain)
	assert.NotNil(t, thorchain.cdc)
}

func TestThorchain_Supports(t *testing.T) {
	thorchain := NewThorchain()

	tests := []struct {
		name     string
		chain    common.Chain
		expected bool
	}{
		{
			name:     "supports THORChain",
			chain:    common.THORChain,
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
			result := thorchain.Supports(tt.chain)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestThorchain_Evaluate_DenyRule(t *testing.T) {
	thorchain := NewThorchain()
	rule := &vtypes.Rule{
		Effect: vtypes.Effect_EFFECT_DENY,
	}

	err := thorchain.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only allow rules supported")
}

func TestThorchain_Evaluate_InvalidResource(t *testing.T) {
	thorchain := NewThorchain()
	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "invalid-resource",
	}

	err := thorchain.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse rule resource")
}

func TestThorchain_Evaluate_InvalidTransactionData(t *testing.T) {
	thorchain := NewThorchain()
	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "thorchain.send",
	}

	err := thorchain.Evaluate(rule, []byte("invalid-tx-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse Thorchain transaction")
}

func TestThorchain_Evaluate_Success_WithTargetValidation(t *testing.T) {
	thorchain := NewThorchain()

	expectedRecipient := "thor1recipient123"
	txBytes := createValidMsgSendTransaction(t, "thor1from123", expectedRecipient, "1000000", "rune", "target test")

	// Create a rule with target validation
	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "thorchain.send",
		Target: &vtypes.Target{
			TargetType: vtypes.TargetType_TARGET_TYPE_ADDRESS,
			Target: &vtypes.Target_Address{
				Address: expectedRecipient,
			},
		},
	}

	err := thorchain.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestThorchain_Evaluate_Failure_TargetMismatch(t *testing.T) {
	thorchain := NewThorchain()

	actualRecipient := "thor1recipient123"
	expectedRecipient := "thor1different456"
	txBytes := createValidMsgSendTransaction(t, "thor1from123", actualRecipient, "1000000", "rune", "target mismatch test")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "thorchain.send",
		Target: &vtypes.Target{
			TargetType: vtypes.TargetType_TARGET_TYPE_ADDRESS,
			Target: &vtypes.Target_Address{
				Address: expectedRecipient,
			},
		},
	}

	err := thorchain.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "target address mismatch")
}

func TestThorchain_Evaluate_Success_WithParameterConstraints(t *testing.T) {
	thorchain := NewThorchain()

	expectedRecipient := "thor1recipient123"
	expectedAmount := "1000000"
	expectedDenom := "rune"
	txBytes := createValidMsgSendTransaction(t, "thor1from123", expectedRecipient, expectedAmount, expectedDenom, "test memo")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "thorchain.send",
		ParameterConstraints: []*vtypes.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &vtypes.Constraint{
					Type: vtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &vtypes.Constraint_FixedValue{
						FixedValue: expectedRecipient,
					},
				},
			},
			{
				ParameterName: "denom",
				Constraint: &vtypes.Constraint{
					Type: vtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &vtypes.Constraint_FixedValue{
						FixedValue: expectedDenom,
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &vtypes.Constraint{
					Type: vtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &vtypes.Constraint_FixedValue{
						FixedValue: expectedAmount,
					},
				},
			},
		},
	}

	err := thorchain.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestThorchain_Evaluate_Failure_ParameterConstraintViolation(t *testing.T) {
	thorchain := NewThorchain()

	actualRecipient := "thor1recipient123"
	expectedRecipient := "thor1different456"
	txBytes := createValidMsgSendTransaction(t, "thor1from123", actualRecipient, "1000000", "rune", "constraint test")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "thorchain.send",
		ParameterConstraints: []*vtypes.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &vtypes.Constraint{
					Type: vtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &vtypes.Constraint_FixedValue{
						FixedValue: expectedRecipient,
					},
				},
			},
		},
	}

	err := thorchain.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to validate parameter constraints")
}

// Helper function to create valid MsgSend transactions for testing
func createValidMsgSendTransaction(t *testing.T, fromAddr, toAddr, amount, denom, memo string) []byte {
	// Use the exact same codec setup as NewThorchain()
	engine := NewThorchain()
	cdc := engine.cdc

	// Create MsgSend
	amountInt, ok := math.NewIntFromString(amount)
	require.True(t, ok, "Invalid amount string")

	msgSend := &banktypes.MsgSend{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		Amount:      sdk.NewCoins(sdk.NewCoin(denom, amountInt)),
	}

	// Pack the message into Any
	msgAny, err := types.NewAnyWithValue(msgSend)
	require.NoError(t, err)

	// Create transaction
	txData := &tx.Tx{
		Body: &tx.TxBody{
			Messages: []*types.Any{msgAny},
			Memo:     memo,
		},
		AuthInfo: &tx.AuthInfo{
			Fee: &tx.Fee{
				Amount:   sdk.NewCoins(),
				GasLimit: 200000,
			},
		},
	}

	// Marshal to protobuf
	protoBytes, err := cdc.Marshal(txData)
	require.NoError(t, err)

	return protoBytes
}

func TestThorchain_parseTransaction(t *testing.T) {
	thorchain := NewThorchain()

	tests := []struct {
		name        string
		txBytes     []byte
		shouldError bool
		errorMsg    string
		setup       func() []byte
	}{
		{
			name: "protobuf transaction with MsgSend",
			setup: func() []byte {
				// Create a protobuf-encoded transaction
				interfaceRegistry := types.NewInterfaceRegistry()
				banktypes.RegisterInterfaces(interfaceRegistry)
				cdc := codec.NewProtoCodec(interfaceRegistry)

				// Create MsgSend
				msgSend := &banktypes.MsgSend{
					FromAddress: "thor1from123456789",
					ToAddress:   "thor1to123456789",
					Amount:      sdk.NewCoins(sdk.NewCoin("rune", math.NewInt(75000000))),
				}

				// Pack the message into Any
				msgAny, err := types.NewAnyWithValue(msgSend)
				if err != nil {
					panic(err)
				}

				// Create transaction
				txData := &tx.Tx{
					Body: &tx.TxBody{
						Messages:                    []*types.Any{msgAny},
						Memo:                        "protobuf test",
						TimeoutHeight:               0,
						ExtensionOptions:            []*types.Any{},
						NonCriticalExtensionOptions: []*types.Any{},
					},
					AuthInfo: &tx.AuthInfo{
						SignerInfos: []*tx.SignerInfo{},
						Fee: &tx.Fee{
							Amount:   sdk.NewCoins(),
							GasLimit: 200000,
							Payer:    "",
							Granter:  "",
						},
					},
				}

				// Marshal to protobuf
				protoBytes, err := cdc.Marshal(txData)
				if err != nil {
					panic(err)
				}
				return protoBytes
			},
			shouldError: false,
		},
		{
			name:        "empty transaction data",
			txBytes:     []byte{},
			shouldError: false, // Empty bytes actually parse successfully as empty protobuf
		},
		{
			name:        "invalid JSON",
			txBytes:     []byte(`{"invalid": json}`),
			shouldError: true,
			errorMsg:    "failed to parse transaction as protobuf, JSON, or hex",
		},
		{
			name:        "invalid hex string",
			txBytes:     []byte("0xZZZZ"),
			shouldError: true,
			errorMsg:    "failed to parse transaction as protobuf, JSON, or hex",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var txBytes []byte
			if tt.setup != nil {
				txBytes = tt.setup()
			} else {
				txBytes = tt.txBytes
			}

			result, err := thorchain.parseTransaction(txBytes)

			if tt.shouldError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				// Note: Body and AuthInfo might be nil for empty transactions
			}
		})
	}
}

func TestThorchain_parseTransaction_Protobuf(t *testing.T) {
	thorchain := NewThorchain()

	// Create a proper protobuf transaction that we know will work
	interfaceRegistry := types.NewInterfaceRegistry()
	banktypes.RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	// Create MsgSend
	msgSend := &banktypes.MsgSend{
		FromAddress: "thor1from123456789",
		ToAddress:   "thor1to123456789",
		Amount:      sdk.NewCoins(sdk.NewCoin("rune", math.NewInt(75000000))),
	}

	// Pack the message into Any
	msgAny, err := types.NewAnyWithValue(msgSend)
	require.NoError(t, err)

	// Create transaction
	txData := &tx.Tx{
		Body: &tx.TxBody{
			Messages:                    []*types.Any{msgAny},
			Memo:                        "test memo",
			TimeoutHeight:               100,
			ExtensionOptions:            []*types.Any{},
			NonCriticalExtensionOptions: []*types.Any{},
		},
		AuthInfo: &tx.AuthInfo{
			SignerInfos: []*tx.SignerInfo{},
			Fee: &tx.Fee{
				Amount:   sdk.NewCoins(),
				GasLimit: 200000,
				Payer:    "",
				Granter:  "",
			},
		},
	}

	// Marshal to protobuf
	protoBytes, err := cdc.Marshal(txData)
	require.NoError(t, err)

	// Test parsing
	result, err := thorchain.parseTransaction(protoBytes)
	require.NoError(t, err)
	require.NotNil(t, result)

	// Validate the parsed transaction structure
	assert.NotNil(t, result.Body)
	assert.Len(t, result.Body.Messages, 1)
	assert.Equal(t, "test memo", result.Body.Memo)
	assert.Equal(t, uint64(100), result.Body.TimeoutHeight)

	assert.NotNil(t, result.AuthInfo)
	assert.NotNil(t, result.AuthInfo.Fee)
	assert.Equal(t, uint64(200000), result.AuthInfo.Fee.GasLimit)
}
