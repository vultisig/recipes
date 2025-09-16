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
	assert.Contains(t, err.Error(), "expected xrpl protocol")
}

func TestXRPL_Evaluate_UnsupportedFunction(t *testing.T) {
	xrpl := NewXRPL()
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "xrp.xrpl.swap",
	}

	err := xrpl.Evaluate(rule, []byte("any-data"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only 'send' function supported")
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

	target := &types.Target{
		TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
		Target: &types.Target_Address{
			Address: "rRecipient456",
		},
	}

	err := xrpl.validateTarget(target, tx)
	assert.NoError(t, err)
}

func TestXRPL_ValidateTarget_Mismatch(t *testing.T) {
	xrpl := NewXRPL()
	tx := &XRPLTransaction{Destination: "rRecipient456"}

	target := &types.Target{
		TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
		Target: &types.Target_Address{
			Address: "rDifferentRecipient789",
		},
	}

	err := xrpl.validateTarget(target, tx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "target address mismatch")
}

func TestXRPL_ValidateParameterConstraints_Success(t *testing.T) {
	xrpl := NewXRPL()
	tx := &XRPLTransaction{
		Destination: "rRecipient456",
		Amount:      "1000000000", // 1000 XRP
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

	err := xrpl.validateParameterConstraints(constraints, tx)
	assert.NoError(t, err)
}

func TestXRPL_ValidateParameterConstraints_Failure(t *testing.T) {
	xrpl := NewXRPL()
	tx := &XRPLTransaction{
		Destination: "rRecipient456",
		Amount:      "3000000000", // 3000 XRP - exceeds max
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

	err := xrpl.validateParameterConstraints(constraints, tx)
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
