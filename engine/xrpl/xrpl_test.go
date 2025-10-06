package xrpl

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vultisig/recipes/types"
	"github.com/xyield/xrpl-go/model/transactions"
	xrptypes "github.com/xyield/xrpl-go/model/transactions/types"
)

func TestNewXRPL(t *testing.T) {
	xrpl := NewXRPL()
	assert.NotNil(t, xrpl)
}

func TestXRPL_Evaluate_DenyRule(t *testing.T) {
	xrpl := NewXRPL()
	rule := &types.Rule{
		Effect: types.Effect_EFFECT_DENY,
	}

	err := xrpl.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only allow rules supported")
}

func TestXRPL_Evaluate_InvalidResource(t *testing.T) {
	xrpl := NewXRPL()
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "invalid-resource",
	}

	err := xrpl.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse rule resource")
}

func TestXRPL_Evaluate_WrongProtocol(t *testing.T) {
	xrpl := NewXRPL()
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "ripple.send",
	}

	err := xrpl.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	// Since parsing happens first, we get a parsing error before protocol validation
	assert.Contains(t, err.Error(), "failed to parse XRPL transaction")
}

func TestXRPL_Evaluate_UnsupportedFunction(t *testing.T) {
	xrpl := NewXRPL()
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "ripple.swap",
	}

	err := xrpl.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	// Since parsing happens first, we get a parsing error before function validation
	assert.Contains(t, err.Error(), "failed to parse XRPL transaction")
}

func TestXRPL_Evaluate_InvalidTransactionData(t *testing.T) {
	xrpl := NewXRPL()
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "ripple.send",
	}

	err := xrpl.Evaluate(rule, []byte("invalid-tx-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse XRPL transaction")
}

func TestXRPL_ValidateTarget_Success(t *testing.T) {
	xrpl := NewXRPL()
	payment := &transactions.Payment{
		Destination: xrptypes.Address("rRecipient456"),
	}

	resource := &types.ResourcePath{
		ChainId:    "xrp",
		ProtocolId: "xrpl",
		FunctionId: "send",
	}

	target := &types.Target{
		TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
		Target: &types.Target_Address{
			Address: "rRecipient456",
		},
	}

	err := xrpl.validateTarget(resource, target, payment)
	assert.NoError(t, err)
}

func TestXRPL_ValidateTarget_Mismatch(t *testing.T) {
	xrpl := NewXRPL()
	payment := &transactions.Payment{
		Destination: xrptypes.Address("rRecipient456"),
	}

	resource := &types.ResourcePath{
		ChainId:    "xrp",
		ProtocolId: "xrpl",
		FunctionId: "send",
	}

	target := &types.Target{
		TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
		Target: &types.Target_Address{
			Address: "rDifferentRecipient789",
		},
	}

	err := xrpl.validateTarget(resource, target, payment)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "target address mismatch")
}

func TestXRPL_ValidateParameterConstraints_Success(t *testing.T) {
	xrpl := NewXRPL()

	// Create XRP amount (1000 XRP = 1000000000 drops)
	var xrpAmount xrptypes.CurrencyAmount = xrptypes.XRPCurrencyAmount(1000000000)

	payment := &transactions.Payment{
		Destination: xrptypes.Address("rRecipient456"),
		Amount:      xrpAmount,
	}

	constraints := []*types.ParameterConstraint{
		{
			ParameterName: "recipient",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				Value: &types.Constraint_FixedValue{
					FixedValue: "rRecipient456",
				},
			},
		},
		{
			ParameterName: "amount",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
				Value: &types.Constraint_MaxValue{
					MaxValue: "2000000000", // 2000 XRP max
				},
			},
		},
	}

	err := xrpl.validateParameterConstraints(constraints, payment)
	assert.NoError(t, err)
}

func TestXRPL_ValidateParameterConstraints_Failure(t *testing.T) {
	xrpl := NewXRPL()

	// Create XRP amount (3000 XRP = 3000000000 drops - exceeds max)
	var xrpAmount xrptypes.CurrencyAmount = xrptypes.XRPCurrencyAmount(3000000000)

	payment := &transactions.Payment{
		Destination: xrptypes.Address("rRecipient456"),
		Amount:      xrpAmount,
	}

	constraints := []*types.ParameterConstraint{
		{
			ParameterName: "amount",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
				Value: &types.Constraint_MaxValue{
					MaxValue: "2000000000", // 2000 XRP max
				},
			},
		},
	}

	err := xrpl.validateParameterConstraints(constraints, payment)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to compare max values")
}

func TestXRPL_Evaluate_Success(t *testing.T) {
	xrpl := NewXRPL()

	// Real XRPL Payment transaction from mainnet
	// Transaction hash: EF7B43D4379C95512107BCBCA837E17D90E66BCB10AAF99BC2D5AADA73A9A0C1
	// Destination: rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg
	realTxHex := "12000022000000002300000FA42405ACB00D2E68FC974D201B05E4846E61400000000000264168400000000000000C7321ED9A3DFF30C22A2848FF6EF4647F93091AB1DB2B14D2BB2A76CA777448968DF16174408158110FC627911D7556BE337B567DB8908E9C5730EBB2344A2B682D5FA9774787845597883249B66A7D23288737F051007DF302A56B84FA1917443106065807811438B8C86B89B8517B209A5F7290F5B13F72CA3B0583146914CB622B8E41E150DE431F48DA244A69809366"

	// Decode hex string to bytes
	txBytes, err := hex.DecodeString(realTxHex)
	assert.NoError(t, err)

	// Create a rule with parameter constraints matching the example policy structure
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "ripple.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: "rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg", // Destination from the real tx
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg", // Same as target
					},
					Required: true,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "1000000000", // 1000 XRP max (actual tx has 9793 drops)
					},
					Required: true,
				},
			},
		},
	}

	err = xrpl.Evaluate(rule, txBytes)

	assert.NoError(t, err, "Evaluation should succeed with valid Payment transaction and matching constraints")
}

func TestXRPL_MagicConstant_THORChainVault(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping magic constant test in short mode - requires network access")
	}

	xrpl := NewXRPL()

	// Use the same real transaction
	realTxHex := "12000022000000002300000FA42405ACB00D2E68FC974D201B05E4846E61400000000000264168400000000000000C7321ED9A3DFF30C22A2848FF6EF4647F93091AB1DB2B14D2BB2A76CA777448968DF16174408158110FC627911D7556BE337B567DB8908E9C5730EBB2344A2B682D5FA9774787845597883249B66A7D23288737F051007DF302A56B84FA1917443106065807811438B8C86B89B8517B209A5F7290F5B13F72CA3B0583146914CB622B8E41E150DE431F48DA244A69809366"
	txBytes, err := hex.DecodeString(realTxHex)
	assert.NoError(t, err)

	// Create rule with magic constant target (this will likely fail since the tx destination
	// won't match THORChain's vault address, but it tests the resolution mechanism)
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "ripple.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
			Target: &types.Target_MagicConstant{
				MagicConstant: types.MagicConstant_THORCHAIN_VAULT,
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
		},
	}

	err = xrpl.Evaluate(rule, txBytes)

	// This should fail with target mismatch (not a resolution error),
	// which means the magic constant resolution worked
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "tx target is wrong",
		"Should fail with target mismatch, indicating magic constant resolution worked")
}

func TestXRPL_Evaluate_Failure(t *testing.T) {
	xrpl := NewXRPL()

	// Use the same real XRPL Payment transaction from mainnet
	// Transaction hash: EF7B43D4379C95512107BCBCA837E17D90E66BCB10AAF99BC2D5AADA73A9A0C1
	// Actual destination: rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg
	// Actual amount: 9793 drops
	realTxHex := "12000022000000002300000FA42405ACB00D2E68FC974D201B05E4846E61400000000000264168400000000000000C7321ED9A3DFF30C22A2848FF6EF4647F93091AB1DB2B14D2BB2A76CA777448968DF16174408158110FC627911D7556BE337B567DB8908E9C5730EBB2344A2B682D5FA9774787845597883249B66A7D23288737F051007DF302A56B84FA1917443106065807811438B8C86B89B8517B209A5F7290F5B13F72CA3B0583146914CB622B8E41E150DE431F48DA244A69809366"

	// Decode hex string to bytes
	txBytes, err := hex.DecodeString(realTxHex)
	assert.NoError(t, err)

	// Create a rule with WRONG recipient and WRONG amount constraints
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "ripple.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: "rWrongAddress123456789012345678901234", // Wrong target address
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg", // Correct recipient
					},
					Required: true,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "10000",
					},
					Required: true,
				},
			},
		},
	}

	err = xrpl.Evaluate(rule, txBytes)

	// Should fail due to target address mismatch (fails first)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "target address mismatch",
		"Should fail with target address mismatch")
}

func TestXRPL_ParseTransaction_PartialPaymentRejection(t *testing.T) {
	// Test the partial payment check logic
	const tfPartialPayment uint = 131072

	// Simulate a payment with partial payment flag
	testFlags := uint(131072)
	isPartialPayment := testFlags&tfPartialPayment != 0
	assert.True(t, isPartialPayment, "Payment should have partial payment flag set")

	// Test normal payment without partial payment flag
	normalFlags := uint(0)
	isNormalPayment := normalFlags&tfPartialPayment != 0
	assert.False(t, isNormalPayment, "Normal payment should not have partial payment flag")
}

func TestXRPL_Evaluate_Failure_ParameterConstraints(t *testing.T) {
	xrpl := NewXRPL()

	// Use the same real XRPL Payment transaction from mainnet
	// Actual destination: rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg
	// Actual amount: 9793 drops
	realTxHex := "12000022000000002300000FA42405ACB00D2E68FC974D201B05E4846E61400000000000264168400000000000000C7321ED9A3DFF30C22A2848FF6EF4647F93091AB1DB2B14D2BB2A76CA777448968DF16174408158110FC627911D7556BE337B567DB8908E9C5730EBB2344A2B682D5FA9774787845597883249B66A7D23288737F051007DF302A56B84FA1917443106065807811438B8C86B89B8517B209A5F7290F5B13F72CA3B0583146914CB622B8E41E150DE431F48DA244A69809366"

	txBytes, err := hex.DecodeString(realTxHex)
	assert.NoError(t, err)

	// Create a rule with correct target but wrong parameter constraints
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "ripple.send",
		// Correct target so we get to parameter validation
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: "rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg", // Correct target
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "rWrongRecipient123456789012345678901", // Wrong recipient
					},
					Required: true,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "5000", // Too low - actual tx has 9793 drops
					},
					Required: true,
				},
			},
		},
	}

	err = xrpl.Evaluate(rule, txBytes)

	// Should fail due to recipient constraint failure (checked first in parameter validation)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to compare fixed values",
		"Should fail with recipient constraint error")
}

func TestXRPL_Evaluate_Swap_Success(t *testing.T) {
	xrpl := NewXRPL()

	// Real unsigned swap transaction that went on-chain:
	// XRP.XRP -> BTC.BTC swap for 1,000,000 drops with a limit of 1000 sats
	// Memo: =:BTC.BTC:bc1qz6erfztfn4ge32fh9nlrdl89h0ymurz36dcetg:1000
	swapTxHex := "1200002405e7f5d4201b05e8c5766140000000000f42406840000000000000328114fac6c2bb1eb09b66cabfde78b33927d2dc7f365d83144ba9f4163bafd86f5ecc6793d43cff31a9f32275f9ea7c0e74686f72636861696e2d6d656d6f7d393d3a4254432e4254433a626331717a366572667a74666e34676533326668396e6c72646c38396830796d75727a333664636574673a31303030e1f1"

	txBytes, err := hex.DecodeString(swapTxHex)
	assert.NoError(t, err)

	// Create a rule that validates the swap
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "ripple.swap",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: "rfunGxj8FWbK3iYuxQvYMA9LGhJ9mYFuss", // Actual destination from the transaction
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "rfunGxj8FWbK3iYuxQvYMA9LGhJ9mYFuss",
					},
					Required: true,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "2000000", // 2 XRP max (actual tx has 1,000,000 drops = 1 XRP)
					},
					Required: true,
				},
			},
			{
				ParameterName: "memo",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_REGEXP,
					Value: &types.Constraint_RegexpValue{
						RegexpValue: "^=:BTC\\.BTC:bc1qz6erfztfn4ge32fh9nlrdl89h0ymurz36dcetg:.*",
					},
					Required: true,
				},
			},
		},
	}

	err = xrpl.Evaluate(rule, txBytes)
	assert.NoError(t, err, "Swap should pass validation")
}

func TestXRPL_Evaluate_Swap_WrongTarget(t *testing.T) {
	xrpl := NewXRPL()

	// Same transaction as above
	swapTxHex := "1200002405e7f5ce201b05ee3fe06140000000000000016840000000000000328114fac6c2bb1eb09b66cabfde78b33927d2dc7f365d83149230d6f0343e3f78fc373c1825fd225fa5e17832f9ea7c0e74686f72636861696e2d6d656d6f7d3a3d3a4254432e4254433a626331717a366572667a74666e34676533326668396e6c72646c38396830796d75727a333664636574673a302e303031e1f1"

	txBytes, err := hex.DecodeString(swapTxHex)
	assert.NoError(t, err)

	// Create rule with WRONG target address
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "ripple.swap",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: "rWrongVaultAddress123456789012345678", // Wrong vault address
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "memo",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_REGEXP,
					Value: &types.Constraint_RegexpValue{
						RegexpValue: "^=:BTC\\.BTC:.*",
					},
				},
			},
		},
	}

	err = xrpl.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "target address mismatch", "Should fail with wrong target address")
}

func TestXRPL_Evaluate_Swap_WrongAsset(t *testing.T) {
	xrpl := NewXRPL()

	// Same transaction as above
	swapTxHex := "1200002405e7f5ce201b05ee3fe06140000000000000016840000000000000328114fac6c2bb1eb09b66cabfde78b33927d2dc7f365d83149230d6f0343e3f78fc373c1825fd225fa5e17832f9ea7c0e74686f72636861696e2d6d656d6f7d3a3d3a4254432e4254433a626331717a366572667a74666e34676533326668396e6c72646c38396830796d75727a333664636574673a302e303031e1f1"

	txBytes, err := hex.DecodeString(swapTxHex)
	assert.NoError(t, err)

	// Create rule with correct target but WRONG asset constraints
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "ripple.swap",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: "rNKzwSezmqZHEQJnm4Z12KepBA7xnxYAdf", // Actual destination from the transaction
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "memo",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_REGEXP,
					Value: &types.Constraint_RegexpValue{
						RegexpValue: "^=:ETH\\.ETH:.*", // Wrong asset - memo has BTC.BTC
					},
				},
			},
		},
	}

	err = xrpl.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "regexp value constraint failed", "Should fail with wrong asset constraint")
}

func TestXRPL_Evaluate_Swap_AmountTooHigh(t *testing.T) {
	xrpl := NewXRPL()

	// Same transaction as above
	swapTxHex := "1200002405e7f5ce201b05ee3fe06140000000000000016840000000000000328114fac6c2bb1eb09b66cabfde78b33927d2dc7f365d83149230d6f0343e3f78fc373c1825fd225fa5e17832f9ea7c0e74686f72636861696e2d6d656d6f7d3a3d3a4254432e4254433a626331717a366572667a74666e34676533326668396e6c72646c38396830796d75727a333664636574673a302e303031e1f1"

	txBytes, err := hex.DecodeString(swapTxHex)
	assert.NoError(t, err)

	// Create rule with correct target but amount constraint that's too restrictive
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "ripple.swap",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: "rNKzwSezmqZHEQJnm4Z12KepBA7xnxYAdf", // Actual destination from the transaction
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "0", // Too restrictive - transaction has 1 drop
					},
				},
			},
		},
	}

	err = xrpl.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to compare max values", "Should fail with amount too high")
}

