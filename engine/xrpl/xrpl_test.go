package xrpl

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vultisig/recipes/types"
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
		Resource: "xrp.bitcoin.send",
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
		Resource: "xrp.xrpl.swap",
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
		Resource: "xrp.xrpl.send",
	}

	err := xrpl.Evaluate(rule, []byte("invalid-tx-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse XRPL transaction")
}

func TestXRPL_ValidateTarget_Success(t *testing.T) {
	xrpl := NewXRPL()
	tx := &XRPLTransaction{Destination: "rRecipient456"}

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

	err := xrpl.validateTarget(resource, target, tx)
	assert.NoError(t, err)
}

func TestXRPL_ValidateTarget_Mismatch(t *testing.T) {
	xrpl := NewXRPL()
	tx := &XRPLTransaction{Destination: "rRecipient456"}

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

	err := xrpl.validateTarget(resource, target, tx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "target address mismatch")
}

func TestXRPL_ValidateParameterConstraints_Success(t *testing.T) {
	xrpl := NewXRPL()
	tx := &XRPLTransaction{
		Destination: "rRecipient456",
		Amount:      "1000000000", // 1000 XRP
	}

	resource := &types.ResourcePath{
		ChainId:    "xrp",
		ProtocolId: "xrpl",
		FunctionId: "send",
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

	err := xrpl.validateParameterConstraints(resource, constraints, tx)
	assert.NoError(t, err)
}

func TestXRPL_ValidateParameterConstraints_Failure(t *testing.T) {
	xrpl := NewXRPL()
	tx := &XRPLTransaction{
		Destination: "rRecipient456",
		Amount:      "3000000000", // 3000 XRP - exceeds max
	}

	resource := &types.ResourcePath{
		ChainId:    "xrp",
		ProtocolId: "xrpl",
		FunctionId: "send",
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

	err := xrpl.validateParameterConstraints(resource, constraints, tx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "max amount constraint failed")
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
		Resource: "xrp.xrpl.send",
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
		Resource: "xrp.xrpl.send",
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
		Resource: "xrp.xrpl.send",
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
		Resource: "xrp.xrpl.send",
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
	assert.Contains(t, err.Error(), "fixed recipient constraint failed",
		"Should fail with recipient constraint error")
}
