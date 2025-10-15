package solana

import (
	"encoding/base64"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/recipes/types"
)

func buildMockSystemTransferTx(from, to solana.PublicKey, amount uint64) []byte {
	transferInst := system.NewTransferInstruction(amount, from, to).Build()
	tx, err := solana.NewTransaction(
		[]solana.Instruction{transferInst},
		solana.Hash{},
		solana.TransactionPayer(from),
	)
	if err != nil {
		panic(err)
	}
	data, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return data
}

func buildMockMultiInstructionTx(from, to solana.PublicKey, amount uint64) []byte {
	transferInst1 := system.NewTransferInstruction(amount, from, to).Build()
	transferInst2 := system.NewTransferInstruction(amount, from, to).Build()
	tx, err := solana.NewTransaction(
		[]solana.Instruction{transferInst1, transferInst2},
		solana.Hash{},
		solana.TransactionPayer(from),
	)
	if err != nil {
		panic(err)
	}
	data, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return data
}

func buildMockSPLTokenTransferTx(source, destination, authority solana.PublicKey, amount uint64) []byte {
	transferInst := token.NewTransferInstruction(amount, source, destination, authority, nil).Build()
	tx, err := solana.NewTransaction(
		[]solana.Instruction{transferInst},
		solana.Hash{},
		solana.TransactionPayer(authority),
	)
	if err != nil {
		panic(err)
	}
	data, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return data
}

func TestEvaluate_SOLTransfer(t *testing.T) {
	const lamports = uint64(1000000)
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	txBytes := buildMockSystemTransferTx(fromKey.PublicKey(), toKey.PublicKey(), lamports)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.system.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_from",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_to",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_lamports",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
					Required: true,
				},
			},
		},
	}
	err = engine.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestEvaluate_SOLTransfer_InvalidAmount(t *testing.T) {
	const lamports = uint64(1000000)
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	txBytes := buildMockSystemTransferTx(fromKey.PublicKey(), toKey.PublicKey(), lamports)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.system.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_from",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_to",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_lamports",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "2000000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to assert: name=arg_lamports")
}

func TestEvaluate_SOLTransfer_InvalidRecipient(t *testing.T) {
	const lamports = uint64(1000000)
	from := solana.NewWallet()
	toTx := solana.NewWallet()
	wrongToRule := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	txBytes := buildMockSystemTransferTx(from.PublicKey(), toTx.PublicKey(), lamports)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.system.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_from",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: from.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_to",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: wrongToRule.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_lamports",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to assert: name=account_to")
}

func TestEvaluate_SOLTransfer_MinAmount(t *testing.T) {
	const lamports = uint64(1000000)
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	txBytes := buildMockSystemTransferTx(fromKey.PublicKey(), toKey.PublicKey(), lamports)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.system.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_from",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_to",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_lamports",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
					Value: &types.Constraint_MinValue{
						MinValue: "500000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestEvaluate_SOLTransfer_MinAmount_Invalid(t *testing.T) {
	const lamports = uint64(1000000)
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	txBytes := buildMockSystemTransferTx(fromKey.PublicKey(), toKey.PublicKey(), lamports)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.system.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_from",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_to",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_lamports",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
					Value: &types.Constraint_MinValue{
						MinValue: "2000000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(
		t,
		err.Error(),
		"failed to assert: name=arg_lamports: failed to compare min values: expected=2000000, actual=1000000",
	)
}

func TestEvaluate_SOLTransfer_MaxAmount(t *testing.T) {
	const lamports = uint64(1000000)
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	txBytes := buildMockSystemTransferTx(fromKey.PublicKey(), toKey.PublicKey(), lamports)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.system.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_from",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_to",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_lamports",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "2000000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestEvaluate_SOLTransfer_MaxAmount_Invalid(t *testing.T) {
	const lamports = uint64(1000000)
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	txBytes := buildMockSystemTransferTx(fromKey.PublicKey(), toKey.PublicKey(), lamports)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.system.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_from",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_to",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_lamports",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "500000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(
		t,
		err.Error(),
		"failed to assert: name=arg_lamports: failed to compare max values: expected=500000, actual=1000000",
	)
}

func TestEvaluate_MultipleInstructions(t *testing.T) {
	const lamports = uint64(1000000)
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.system.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_from",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_to",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toKey.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_lamports",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
					Required: true,
				},
			},
		},
	}

	multiTxBytes := buildMockMultiInstructionTx(fromKey.PublicKey(), toKey.PublicKey(), lamports)
	err = engine.Evaluate(rule, multiTxBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only single instruction transactions are allowed")
}

func TestEvaluate_SPLTokenTransfer(t *testing.T) {
	const tokenAmount = uint64(500000)
	authorityWallet := solana.NewWallet()
	sourceWallet := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	destinationWallet := solana.NewWallet()
	sourceTokenAccount := sourceWallet.PublicKey()
	destinationTokenAccount := destinationWallet.PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityWallet.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: token.ProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_source",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: sourceTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_destination",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: destinationTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_authority",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: authorityWallet.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "500000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestEvaluate_SPLTokenTransfer_InvalidAmount(t *testing.T) {
	const tokenAmount = uint64(500000)
	authorityWallet := solana.NewWallet()
	sourceWallet := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	destinationWallet := solana.NewWallet()
	sourceTokenAccount := sourceWallet.PublicKey()
	destinationTokenAccount := destinationWallet.PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityWallet.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: token.ProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_source",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: sourceTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_destination",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: destinationTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_authority",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: authorityWallet.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(
		t,
		err.Error(),
		"failed to assert: name=arg_amount: failed to compare fixed values: expected=1000000, actual=500000",
	)
}

func TestEvaluate_SPLTokenTransfer_InvalidRecipient(t *testing.T) {
	const tokenAmount = uint64(500000)
	authorityWallet := solana.NewWallet()
	sourceWallet := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	destinationWallet := solana.NewWallet()
	wrongDestinationWallet := solana.NewWallet()
	sourceTokenAccount := sourceWallet.PublicKey()
	destinationTokenAccount := destinationWallet.PublicKey()
	wrongDestination := wrongDestinationWallet.PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityWallet.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: token.ProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_source",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: sourceTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_destination",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: wrongDestination.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_authority",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: authorityWallet.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "500000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to assert accounts: failed to assert: name=account_destination")
}

func TestEvaluate_SPLTokenTransfer_MinAmount(t *testing.T) {
	const tokenAmount = uint64(500000)
	authorityWallet := solana.NewWallet()
	sourceWallet := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	destinationWallet := solana.NewWallet()
	sourceTokenAccount := sourceWallet.PublicKey()
	destinationTokenAccount := destinationWallet.PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityWallet.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: token.ProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_authority",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: authorityWallet.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_source",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: sourceTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_destination",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: destinationTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
					Value: &types.Constraint_MinValue{
						MinValue: "300000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestEvaluate_SPLTokenTransfer_MinAmount_Invalid(t *testing.T) {
	const tokenAmount = uint64(500000)
	authorityWallet := solana.NewWallet()
	sourceWallet := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	destinationWallet := solana.NewWallet()
	sourceTokenAccount := sourceWallet.PublicKey()
	destinationTokenAccount := destinationWallet.PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityWallet.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: token.ProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_authority",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: authorityWallet.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_source",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: sourceTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_destination",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: destinationTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
					Value: &types.Constraint_MinValue{
						MinValue: "800000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(
		t,
		err.Error(),
		"failed to assert: name=arg_amount: failed to compare min values: expected=800000, actual=500000",
	)
}

func TestEvaluate_SPLTokenTransfer_MaxAmount(t *testing.T) {
	const tokenAmount = uint64(500000)
	authorityWallet := solana.NewWallet()
	sourceWallet := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	destinationWallet := solana.NewWallet()
	sourceTokenAccount := sourceWallet.PublicKey()
	destinationTokenAccount := destinationWallet.PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityWallet.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: token.ProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_authority",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: authorityWallet.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_source",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: sourceTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_destination",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: destinationTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "800000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}

func TestEvaluate_SPLTokenTransfer_MaxAmount_Invalid(t *testing.T) {
	const tokenAmount = uint64(500000)
	authorityWallet := solana.NewWallet()
	sourceWallet := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	destinationWallet := solana.NewWallet()
	sourceTokenAccount := sourceWallet.PublicKey()
	destinationTokenAccount := destinationWallet.PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityWallet.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: token.ProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_authority",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: authorityWallet.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_source",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: sourceTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_destination",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: destinationTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "300000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(
		t,
		err.Error(),
		"failed to assert: name=arg_amount: failed to compare max values: expected=300000, actual=500000",
	)
}

func buildMockFakeSPLTokenTransferTx(source, destination, authority solana.PublicKey, amount uint64) []byte {
	fakeProgramID := solana.NewWallet().PublicKey()

	instructionData := make([]byte, 9)
	instructionData[0] = 3
	for i := 0; i < 8; i++ {
		instructionData[1+i] = byte(amount >> (i * 8))
	}

	instruction := solana.NewInstruction(
		fakeProgramID,
		solana.AccountMetaSlice{
			{PublicKey: source, IsWritable: true, IsSigner: false},
			{PublicKey: destination, IsWritable: true, IsSigner: false},
			{PublicKey: authority, IsWritable: false, IsSigner: true},
		},
		instructionData,
	)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{instruction},
		solana.Hash{},
		solana.TransactionPayer(authority),
	)
	if err != nil {
		panic(err)
	}

	data, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return data
}

func TestEvaluate_SPLTokenTransfer_InvalidProgram(t *testing.T) {
	const tokenAmount = uint64(500000)
	authorityWallet := solana.NewWallet()
	sourceWallet := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	destinationWallet := solana.NewWallet()
	sourceTokenAccount := sourceWallet.PublicKey()
	destinationTokenAccount := destinationWallet.PublicKey()

	txBytes := buildMockFakeSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityWallet.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: token.ProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_authority",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: authorityWallet.PublicKey().String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_source",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: sourceTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_destination",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: destinationTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "500000",
					},
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to assert target: tx target is wrong")
}

func TestEvaluate_JupiterSharedAccountsRoute(t *testing.T) {
	tx := "CgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKCQgTBi8dL5WH7sWHTBp28kz8oNCmbKHfJXQoeBxJllkgszhxM32R3yvnX/HBzojB9fApPOfuRWL27j4EX6B/qtMR3WZ/y1OFUjaxGSYPYSYxL38DI5jSwryQ6RMhtMx+2NckxFh1gLBtjJtNTvcrXvXuh/PFZ/fM4IPgPkuzzyhB/G27Pda2pnGc3R0WktKvNfpBJorRv4iVoUOTn784IlhxGfXyb4o/nn+6d0cPFBPmUKZ4EVZLKuWY5mIcburZi74iieJ/mkkOagUIRksszWMNCgHFsDMiMCfHN+J1rwoXyWiagAv/TIc2iJbCD8FAc+vxy1qjdf6B/k29yCuk37deeOvseGzjAy+w0P7kS+cz0i3FdAlPQMKRmD4EAqGenMprv67g5bFDc7KxpFz5i1fj1470rV3a6ExmnSw6yPuHi/M6az5ztylrtKOlPWSGFPJNVv+Lb9+NyZHORsoMDZu6UAbd9uHXZaGT2cvhRs7reawctIXtX1s3kTqM9YV+/wCpBpuIV/6rgYT7aH9jRhjANdrEOdwa6ztVmKDwAAAAAAHG+nrzvtutOj1l82qryXQxsbvkwtL24OR8pgIDRS9dYQR51VvyMcBu7nTFbs5oFQf9sbLeo/SOUQKxzaJWvBOPtD/6J/XX9kp0wJsfKVh53ksJqzbfyd1RSzIap7OM5ei1w1W367Ykl8/1heeE1Ct6pgMZQ89eFMSv0TWee6UaMEcn0nz5UKgy0QJ34xepN6SZQQ1LggwZ6QPHCYVaRRN9zNbBTO0ZD5TB0ZEBaayT6GzFN/sZJShMvBo8S2hacJKEB66guxxmMx9J7GcDW1ZYuFjwk5echKsWPvICBacgSQEOGQsHCgABAgMMDQ4ODw4QEQQSBQYHAQIICQskwSCbM0HWnIECAQAAADBkAAGA8PoCAAAAAGSslgAAAAAAZAAA"

	engine, err := NewSolana()
	require.NoError(t, err)

	txBytes, err := base64.StdEncoding.DecodeString(tx)
	require.NoError(t, err)

	jupiterV6ProgramID := solana.MustPublicKeyFromBase58("JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4")
	jupiterEvent := solana.MustPublicKeyFromBase58("D8cy77BBepLMngZx6ZukaTff5hCt1HrWyKk3Hnd9oitf")
	usdcMint := solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")

	programAuthority := solana.MustPublicKeyFromBase58("BQ72nSv9f3PRyRKCBnHLVrerrv37CYTHm5h3s9VSGQDV")
	userTransferAuthority := solana.MustPublicKeyFromBase58("4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM")
	sourceTokenAccount := solana.MustPublicKeyFromBase58("R97cgCoxcqrUaaW7wg8drNBiLkicyQQHmct7pY8tdMR")
	programSourceTokenAccount := solana.MustPublicKeyFromBase58("8ctcHN52LY21FEipCjr1MVWtoZa1irJQTPyAaTj72h7S")
	programDestinationTokenAccount := solana.MustPublicKeyFromBase58("7u7cD7NxcZEuzRCBaYo8uVpotRdqZwez47vvuwzCov43")
	destinationTokenAccount := solana.MustPublicKeyFromBase58("EDT9FrASLP4gRKFnka5h4vgVBHzrmTZxgxmGa4G4vect")

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.jupiter_aggregatorv6.sharedAccountsRoute",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: jupiterV6ProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "account_tokenProgram",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: solana.TokenProgramID.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_userTransferAuthority",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: userTransferAuthority.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_programAuthority",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: programAuthority.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_sourceTokenAccount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: sourceTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_programSourceTokenAccount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: programSourceTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_programDestinationTokenAccount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: programDestinationTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_destinationTokenAccount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: destinationTokenAccount.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_sourceMint",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: solana.SolMint.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_destinationMint",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: usdcMint.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_platformFeeAccount",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
			{
				ParameterName: "account_token2022Program",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
			{
				ParameterName: "account_eventAuthority",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: jupiterEvent.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "account_program",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: jupiterV6ProgramID.String(),
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_id",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
			{
				ParameterName: "arg_routePlan",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
			{
				ParameterName: "arg_slippageBps",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
			{
				ParameterName: "arg_platformFeeBps",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
			{
				ParameterName: "arg_inAmount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "50000000",
					},
					Required: true,
				},
			},
			{
				ParameterName: "arg_quotedOutAmount",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
		},
	}

	err = engine.Evaluate(rule, txBytes)
	assert.NoError(t, err)
}
