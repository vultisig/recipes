package gaia

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

func TestNewGaia(t *testing.T) {
	gaia := NewGaia()
	assert.NotNil(t, gaia)
	assert.NotNil(t, gaia.cdc)
}

func TestGaia_Supports(t *testing.T) {
	gaia := NewGaia()

	tests := []struct {
		name     string
		chain    common.Chain
		expected bool
	}{
		{
			name:     "supports GaiaChain",
			chain:    common.GaiaChain,
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
		{
			name:     "does not support THORChain",
			chain:    common.THORChain,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gaia.Supports(tt.chain)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGaia_Evaluate_DenyRule(t *testing.T) {
	gaia := NewGaia()
	rule := &vtypes.Rule{
		Effect: vtypes.Effect_EFFECT_DENY,
	}

	err := gaia.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only allow rules supported")
}

func TestGaia_Evaluate_InvalidResource(t *testing.T) {
	gaia := NewGaia()
	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "invalid-resource",
	}

	err := gaia.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse rule resource")
}

func TestGaia_Evaluate_InvalidTransactionData(t *testing.T) {
	gaia := NewGaia()
	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "cosmos.atom.transfer",
	}

	err := gaia.Evaluate(rule, []byte("invalid-tx-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse Cosmos transaction")
}

func TestGaia_Evaluate_Success_WithTargetValidation(t *testing.T) {
	gaia := NewGaia()

	expectedRecipient := "cosmos1recipient123"
	txBytes := createValidMsgSendTransaction(t, "cosmos1from123", expectedRecipient, "1000000", "uatom", "target test")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "cosmos.atom.transfer",
		Target: &vtypes.Target{
			TargetType: vtypes.TargetType_TARGET_TYPE_ADDRESS,
			Target: &vtypes.Target_Address{
				Address: expectedRecipient,
			},
		},
	}

	err := gaia.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestGaia_Evaluate_Failure_TargetMismatch(t *testing.T) {
	gaia := NewGaia()

	actualRecipient := "cosmos1recipient123"
	expectedRecipient := "cosmos1different456"
	txBytes := createValidMsgSendTransaction(t, "cosmos1from123", actualRecipient, "1000000", "uatom", "target mismatch test")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "cosmos.atom.transfer",
		Target: &vtypes.Target{
			TargetType: vtypes.TargetType_TARGET_TYPE_ADDRESS,
			Target: &vtypes.Target_Address{
				Address: expectedRecipient,
			},
		},
	}

	err := gaia.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "target address mismatch")
}

func TestGaia_Evaluate_Success_WithParameterConstraints(t *testing.T) {
	gaia := NewGaia()

	expectedRecipient := "cosmos1recipient123"
	expectedAmount := "1000000"
	expectedDenom := "uatom"
	txBytes := createValidMsgSendTransaction(t, "cosmos1from123", expectedRecipient, expectedAmount, expectedDenom, "test memo")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "cosmos.atom.transfer",
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

	err := gaia.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestGaia_Evaluate_Failure_ParameterConstraintViolation(t *testing.T) {
	gaia := NewGaia()

	actualRecipient := "cosmos1recipient123"
	expectedRecipient := "cosmos1different456"
	txBytes := createValidMsgSendTransaction(t, "cosmos1from123", actualRecipient, "1000000", "uatom", "constraint test")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "cosmos.atom.transfer",
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

	err := gaia.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to validate parameter constraints")
}

func TestGaia_Evaluate_Success_SwapViaTHORChain(t *testing.T) {
	gaia := NewGaia()

	// Simulate a swap via THORChain - ATOM -> BTC
	// The transaction is a normal MsgSend to THORChain's GAIA vault with a swap memo
	swapMemo := "=:BTC.BTC:bc1qz6erfztfn4ge32fh9nlrdl89h0ymurz36dcetg"
	txBytes := createValidMsgSendTransaction(t, "cosmos1from123", "cosmos1thorchainvault", "10000000", "uatom", swapMemo)

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "cosmos.atom.transfer",
		Target: &vtypes.Target{
			TargetType: vtypes.TargetType_TARGET_TYPE_ADDRESS,
			Target: &vtypes.Target_Address{
				Address: "cosmos1thorchainvault",
			},
		},
		ParameterConstraints: []*vtypes.ParameterConstraint{
			{
				ParameterName: "memo",
				Constraint: &vtypes.Constraint{
					Type: vtypes.ConstraintType_CONSTRAINT_TYPE_REGEXP,
					Value: &vtypes.Constraint_RegexpValue{
						RegexpValue: "^=:BTC\\.BTC:bc1[a-zA-Z0-9]+$",
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &vtypes.Constraint{
					Type: vtypes.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &vtypes.Constraint_MaxValue{
						MaxValue: "100000000", // 100 ATOM max
					},
				},
			},
		},
	}

	err := gaia.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestGaia_Evaluate_Failure_SwapAmountTooHigh(t *testing.T) {
	gaia := NewGaia()

	swapMemo := "=:BTC.BTC:bc1qz6erfztfn4ge32fh9nlrdl89h0ymurz36dcetg"
	txBytes := createValidMsgSendTransaction(t, "cosmos1from123", "cosmos1thorchainvault", "200000000", "uatom", swapMemo)

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "cosmos.atom.transfer",
		Target: &vtypes.Target{
			TargetType: vtypes.TargetType_TARGET_TYPE_ADDRESS,
			Target: &vtypes.Target_Address{
				Address: "cosmos1thorchainvault",
			},
		},
		ParameterConstraints: []*vtypes.ParameterConstraint{
			{
				ParameterName: "amount",
				Constraint: &vtypes.Constraint{
					Type: vtypes.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &vtypes.Constraint_MaxValue{
						MaxValue: "100000000", // 100 ATOM max
					},
				},
			},
		},
	}

	err := gaia.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to validate parameter constraints")
}

// Helper function to create valid MsgSend transactions for testing
func createValidMsgSendTransaction(t *testing.T, fromAddr, toAddr, amount, denom, memo string) []byte {
	engine := NewGaia()
	cdc := engine.cdc

	amountInt, ok := math.NewIntFromString(amount)
	require.True(t, ok, "Invalid amount string")

	msgSend := &banktypes.MsgSend{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		Amount:      sdk.NewCoins(sdk.NewCoin(denom, amountInt)),
	}

	msgAny, err := types.NewAnyWithValue(msgSend)
	require.NoError(t, err)

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

	protoBytes, err := cdc.Marshal(txData)
	require.NoError(t, err)

	return protoBytes
}

func TestGaia_parseTransaction(t *testing.T) {
	gaia := NewGaia()

	tests := []struct {
		name        string
		txBytes     []byte
		shouldError bool
		errorMsg    string
		setup       func() []byte
	}{
		{
			name: "valid protobuf transaction with MsgSend",
			setup: func() []byte {
				interfaceRegistry := types.NewInterfaceRegistry()
				banktypes.RegisterInterfaces(interfaceRegistry)
				cdc := codec.NewProtoCodec(interfaceRegistry)

				msgSend := &banktypes.MsgSend{
					FromAddress: "cosmos1from123456789",
					ToAddress:   "cosmos1to123456789",
					Amount:      sdk.NewCoins(sdk.NewCoin("uatom", math.NewInt(75000000))),
				}

				msgAny, err := types.NewAnyWithValue(msgSend)
				if err != nil {
					panic(err)
				}

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
						},
					},
				}

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
			shouldError: true,
			errorMsg:    "empty transaction data",
		},
		{
			name:        "invalid protobuf",
			txBytes:     []byte(`{"invalid": json}`),
			shouldError: true,
			errorMsg:    "failed to unmarshal protobuf transaction",
		},
		{
			name: "transaction size too large",
			setup: func() []byte {
				return make([]byte, 32*1024+1) // 32 KB + 1 byte
			},
			shouldError: true,
			errorMsg:    "transaction too large",
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

			result, err := gaia.parseTransaction(txBytes)

			if tt.shouldError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

