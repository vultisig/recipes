package solana

import (
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
		Resource: "solana.token.transfer",
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
		Resource: "solana.token.transfer",
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
		Resource: "solana.token.transfer",
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
		Resource: "solana.token.transfer",
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
		Resource: "solana.token.transfer",
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
		Resource: "solana.token.transfer",
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
		Resource: "solana.token.transfer",
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
		Resource: "solana.token.transfer",
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
