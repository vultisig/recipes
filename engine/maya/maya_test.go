package maya

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	vtypes "github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

func TestNewMaya(t *testing.T) {
	maya := NewMaya()
	assert.NotNil(t, maya)
	assert.NotNil(t, maya.cdc)
}

func TestMaya_Supports(t *testing.T) {
	maya := NewMaya()

	tests := []struct {
		name     string
		chain    common.Chain
		expected bool
	}{
		{
			name:     "supports MAYAChain",
			chain:    common.MayaChain,
			expected: true,
		},
		{
			name:     "does not support Bitcoin",
			chain:    common.Bitcoin,
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
			result := maya.Supports(tt.chain)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaya_Evaluate_DenyRule(t *testing.T) {
	maya := NewMaya()
	rule := &vtypes.Rule{
		Effect: vtypes.Effect_EFFECT_DENY,
	}

	err := maya.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only allow rules supported")
}

func TestMaya_Evaluate_InvalidResource(t *testing.T) {
	maya := NewMaya()
	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "invalid-resource",
	}

	err := maya.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse rule resource")
}

func TestMaya_Evaluate_InvalidTransactionData(t *testing.T) {
	maya := NewMaya()
	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "mayachain.send.cacao",
	}

	err := maya.Evaluate(rule, []byte("invalid-tx-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse MAYAChain transaction")
}

func TestMaya_Evaluate_Success_WithTargetValidation(t *testing.T) {
	maya := NewMaya()

	expectedRecipient := "maya1recipient123"
	txBytes := createValidMsgSendTransaction(t, "maya1from123", expectedRecipient, "1000000", "cacao", "target test")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "mayachain.send.cacao",
		Target: &vtypes.Target{
			TargetType: vtypes.TargetType_TARGET_TYPE_ADDRESS,
			Target: &vtypes.Target_Address{
				Address: expectedRecipient,
			},
		},
	}

	err := maya.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestMaya_Evaluate_Failure_TargetMismatch(t *testing.T) {
	maya := NewMaya()

	actualRecipient := "maya1recipient123"
	expectedRecipient := "maya1different456"
	txBytes := createValidMsgSendTransaction(t, "maya1from123", actualRecipient, "1000000", "cacao", "target mismatch test")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "mayachain.send.cacao",
		Target: &vtypes.Target{
			TargetType: vtypes.TargetType_TARGET_TYPE_ADDRESS,
			Target: &vtypes.Target_Address{
				Address: expectedRecipient,
			},
		},
	}

	err := maya.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "target address mismatch")
}

func TestMaya_Evaluate_Success_WithParameterConstraints(t *testing.T) {
	maya := NewMaya()

	expectedRecipient := "maya1recipient123"
	expectedAmount := "1000000"
	expectedDenom := "cacao"
	txBytes := createValidMsgSendTransaction(t, "maya1from123", expectedRecipient, expectedAmount, expectedDenom, "test memo")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "mayachain.send.cacao",
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

	err := maya.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestMaya_Evaluate_Failure_ParameterConstraintViolation(t *testing.T) {
	maya := NewMaya()

	actualRecipient := "maya1recipient123"
	expectedRecipient := "maya1different456"
	txBytes := createValidMsgSendTransaction(t, "maya1from123", actualRecipient, "1000000", "cacao", "constraint test")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "mayachain.send.cacao",
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

	err := maya.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to validate parameter constraints")
}

func TestMaya_Evaluate_Success_MayachainSwap(t *testing.T) {
	maya := NewMaya()

	expectedAmount := "75000000"
	expectedSymbol := "cacao"
	expectedMemo := "=:ETH.ETH:0x1234567890123456789012345678901234567890"
	txBytes := createValidMsgDepositTransaction(t, "maya1signer123", expectedAmount, expectedSymbol, expectedMemo)

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "mayachain.mayachain_swap",
		Target: &vtypes.Target{
			TargetType: vtypes.TargetType_TARGET_TYPE_UNSPECIFIED,
		},
		ParameterConstraints: []*vtypes.ParameterConstraint{
			{
				ParameterName: "amount",
				Constraint: &vtypes.Constraint{
					Type: vtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &vtypes.Constraint_FixedValue{
						FixedValue: expectedAmount,
					},
				},
			},
			{
				ParameterName: "memo",
				Constraint: &vtypes.Constraint{
					Type: vtypes.ConstraintType_CONSTRAINT_TYPE_REGEXP,
					Value: &vtypes.Constraint_RegexpValue{
						RegexpValue: "^=:ETH\\.ETH:0x[a-fA-F0-9]{40}$",
					},
				},
			},
			{
				ParameterName: "from_asset",
				Constraint: &vtypes.Constraint{
					Type: vtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &vtypes.Constraint_FixedValue{
						FixedValue: "CACAO",
					},
				},
			},
		},
	}

	err := maya.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestMaya_Evaluate_Failure_MayachainSwap_ParameterMismatch(t *testing.T) {
	maya := NewMaya()

	actualAmount := "75000000"
	expectedAmount := "100000000"
	txBytes := createValidMsgDepositTransaction(t, "maya1signer123", actualAmount, "cacao", "=:ETH.ETH:0x1234567890123456789012345678901234567890")

	rule := &vtypes.Rule{
		Effect:   vtypes.Effect_EFFECT_ALLOW,
		Resource: "mayachain.mayachain_swap.swap",
		Target: &vtypes.Target{
			TargetType: vtypes.TargetType_TARGET_TYPE_UNSPECIFIED,
		},
		ParameterConstraints: []*vtypes.ParameterConstraint{
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

	err := maya.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to validate parameter constraints")
}

// Helper function to create valid MsgDeposit transactions for testing mayachain_swap
func createValidMsgDepositTransaction(t *testing.T, signer, amount, symbol, memo string) []byte {
	engine := NewMaya()
	cdc := engine.cdc

	signerBytes := []byte(signer)

	pbCoins := make([]*vtypes.Coin, 1)
	pbCoins[0] = &vtypes.Coin{
		Asset: &vtypes.Asset{
			Chain:  "MAYA",
			Symbol: symbol,
			Ticker: symbol,
		},
		Amount:   amount,
		Decimals: 10, // MAYAChain uses 10 decimals
	}

	msgDeposit := &vtypes.MsgDeposit{
		Coins:  pbCoins,
		Memo:   memo,
		Signer: signerBytes,
	}

	msgAny, err := types.NewAnyWithValue(msgDeposit)
	require.NoError(t, err)

	txData := &tx.Tx{
		Body: &tx.TxBody{
			Messages: []*types.Any{msgAny},
			Memo:     "",
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

// Helper function to create valid MsgSend transactions for testing
func createValidMsgSendTransaction(t *testing.T, fromAddr, toAddr, amount, denom, memo string) []byte {
	engine := NewMaya()
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

func TestMaya_parseTransaction(t *testing.T) {
	maya := NewMaya()

	tests := []struct {
		name        string
		txBytes     []byte
		shouldError bool
		errorMsg    string
		setup       func() []byte
	}{
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
				return make([]byte, 32*1024+1)
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

			result, err := maya.parseTransaction(txBytes)

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

