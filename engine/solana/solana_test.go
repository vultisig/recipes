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
	engine, err := NewSolana("sol")
	require.NoError(t, err)

	txBytes := buildMockSystemTransferTx(fromKey.PublicKey(), toKey.PublicKey(), lamports)

	// Test valid transfer
	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.sol.transfer",
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

func TestEvaluate_MultipleInstructions(t *testing.T) {
	const lamports = uint64(1000000)
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()
	engine, err := NewSolana("sol")
	require.NoError(t, err)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.sol.transfer",
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
	engine, err := NewSolana("sol")
	require.NoError(t, err)

	// Create mock SPL token accounts
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
