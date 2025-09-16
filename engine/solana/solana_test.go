package solana

import (
	"strings"
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
				Address: toKey.PublicKey().String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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
				Address: toKey.PublicKey().String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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
	assert.Contains(t, err.Error(), "failed to compare fixed values")
}

func TestEvaluate_SOLTransfer_InvalidRecipient(t *testing.T) {
	const lamports = uint64(1000000)
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()
	wrongRecipient := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	txBytes := buildMockSystemTransferTx(fromKey.PublicKey(), toKey.PublicKey(), lamports)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.system.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: wrongRecipient.PublicKey().String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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
	assert.Contains(t, err.Error(), "tx target is wrong")
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
				Address: toKey.PublicKey().String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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

	invalidRule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.system.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: toKey.PublicKey().String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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

	err = engine.Evaluate(invalidRule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to compare min values")
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
				Address: toKey.PublicKey().String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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

	invalidRule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.system.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: toKey.PublicKey().String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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

	err = engine.Evaluate(invalidRule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to compare max values")
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
				Address: toKey.PublicKey().String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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
	authorityKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	sourceTokenAccount := solana.NewWallet().PublicKey()
	destinationTokenAccount := solana.NewWallet().PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityKey.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: destinationTokenAccount.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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
	authorityKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	sourceTokenAccount := solana.NewWallet().PublicKey()
	destinationTokenAccount := solana.NewWallet().PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityKey.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: destinationTokenAccount.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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
	assert.Contains(t, err.Error(), "failed to compare fixed values")
}

func TestEvaluate_SPLTokenTransfer_InvalidRecipient(t *testing.T) {
	const tokenAmount = uint64(500000)
	authorityKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	sourceTokenAccount := solana.NewWallet().PublicKey()
	destinationTokenAccount := solana.NewWallet().PublicKey()
	wrongDestination := solana.NewWallet().PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityKey.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: wrongDestination.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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
	assert.Contains(t, err.Error(), "tx target is wrong")
}

func TestEvaluate_SPLTokenTransfer_MinAmount(t *testing.T) {
	const tokenAmount = uint64(500000)
	authorityKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	sourceTokenAccount := solana.NewWallet().PublicKey()
	destinationTokenAccount := solana.NewWallet().PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityKey.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: destinationTokenAccount.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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

	invalidRule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: destinationTokenAccount.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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

	err = engine.Evaluate(invalidRule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to compare min values")
}

func TestEvaluate_SPLTokenTransfer_MaxAmount(t *testing.T) {
	const tokenAmount = uint64(500000)
	authorityKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	sourceTokenAccount := solana.NewWallet().PublicKey()
	destinationTokenAccount := solana.NewWallet().PublicKey()

	txBytes := buildMockSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityKey.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: destinationTokenAccount.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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

	invalidRule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: destinationTokenAccount.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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

	err = engine.Evaluate(invalidRule, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to compare max values")
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
	authorityKey := solana.NewWallet()
	engine, err := NewSolana()
	require.NoError(t, err)

	sourceTokenAccount := solana.NewWallet().PublicKey()
	destinationTokenAccount := solana.NewWallet().PublicKey()

	txBytes := buildMockFakeSPLTokenTransferTx(
		sourceTokenAccount,
		destinationTokenAccount,
		authorityKey.PublicKey(),
		tokenAmount,
	)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.spl_token.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: destinationTokenAccount.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
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
	assert.True(t,
		strings.Contains(err.Error(), "instruction is not calling official SPL Token program") ||
			strings.Contains(err.Error(), "tx target is wrong"),
		"Expected security-related error, got: %s", err.Error(),
	)
}
