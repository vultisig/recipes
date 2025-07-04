package thorchain

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/recipes/types"
)

func TestThorchainChain(t *testing.T) {
	chain := NewThorchain()

	// Test basic chain properties
	assert.Equal(t, "thorchain", chain.ID())
	assert.Equal(t, "Thorchain", chain.Name())
	assert.Contains(t, chain.Description(), "Thorchain")

	// Test supported protocols
	protocols := chain.SupportedProtocols()
	assert.Contains(t, protocols, "rune")
	assert.Contains(t, protocols, "tcy")
	assert.Len(t, protocols, 2)
}

func TestThorchainProtocols(t *testing.T) {
	chain := NewThorchain()

	// Test RUNE protocol
	runeProtocol, err := chain.GetProtocol("rune")
	require.NoError(t, err)
	assert.Equal(t, "rune", runeProtocol.ID())
	assert.Equal(t, "RUNE", runeProtocol.Name())
	assert.Equal(t, "thorchain", runeProtocol.ChainID())

	// Test TCY protocol
	tcyProtocol, err := chain.GetProtocol("tcy")
	require.NoError(t, err)
	assert.Equal(t, "tcy", tcyProtocol.ID())
	assert.Equal(t, "TCY", tcyProtocol.Name())
	assert.Equal(t, "thorchain", tcyProtocol.ChainID())

	// Test invalid protocol
	_, err = chain.GetProtocol("invalid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found or not supported")
}

func TestThorchainProtocolFunctions(t *testing.T) {
	chain := NewThorchain()

	// Test RUNE protocol functions
	runeProtocol, err := chain.GetProtocol("rune")
	require.NoError(t, err)

	functions := runeProtocol.Functions()
	assert.Len(t, functions, 1) // Only transfer function now

	// Find transfer function
	var transferFunc *types.Function
	for _, fn := range functions {
		if fn.ID == "transfer" {
			transferFunc = fn
			break
		}
	}
	require.NotNil(t, transferFunc, "transfer function should exist")

	assert.Equal(t, "transfer", transferFunc.ID)
	assert.Equal(t, "Transfer RUNE", transferFunc.Name)
	assert.Contains(t, transferFunc.Description, "Transfer")

	// Test function parameters for transfer
	params := transferFunc.Parameters
	assert.Len(t, params, 3)

	// Check parameter names
	paramNames := make([]string, len(params))
	for i, param := range params {
		paramNames[i] = param.Name
	}
	assert.Contains(t, paramNames, "recipient")
	assert.Contains(t, paramNames, "amount")
	assert.Contains(t, paramNames, "memo")

	// Test GetFunction
	getFunc, err := runeProtocol.GetFunction("transfer")
	require.NoError(t, err)
	assert.Equal(t, "transfer", getFunc.ID)

	// Test invalid function
	_, err = runeProtocol.GetFunction("invalid")
	assert.Error(t, err)
}

func TestProtocolDenominationMatching(t *testing.T) {
	runeProtocol := NewRUNE()
	tcyProtocol := NewTCY()

	// Create mock transactions with different denominations
	runeTx := &ParsedThorchainTransaction{
		txHash:   "rune_tx_hash",
		fromAddr: "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
		toAddr:   "thor13m8mqtcv7c2srnpr4efucwkrdu29tq3ejgu52r",
		amount:   big.NewInt(100000000), // 1 RUNE
		denom:    "rune",
		memo:     "test rune transfer",
		msgType:  "MsgSend",
	}

	tcyTx := &ParsedThorchainTransaction{
		txHash:   "tcy_tx_hash",
		fromAddr: "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
		toAddr:   "thor13m8mqtcv7c2srnpr4efucwkrdu29tq3ejgu52r",
		amount:   big.NewInt(50000000), // 0.5 TCY
		denom:    "tcy",
		memo:     "test tcy transfer",
		msgType:  "MsgSend",
	}

	// Create policy matcher
	matcher := &types.PolicyFunctionMatcher{
		FunctionID:  "transfer",
		Constraints: []*types.ParameterConstraint{},
	}

	// Test RUNE protocol with RUNE transaction (should match)
	matched, params, err := runeProtocol.MatchFunctionCall(runeTx, matcher)
	assert.NoError(t, err)
	assert.True(t, matched, "RUNE protocol should match RUNE transaction")
	assert.Equal(t, "rune", params["denom"])

	// Test RUNE protocol with TCY transaction (should NOT match)
	matched, _, err = runeProtocol.MatchFunctionCall(tcyTx, matcher)
	assert.NoError(t, err)
	assert.False(t, matched, "RUNE protocol should NOT match TCY transaction")

	// Test TCY protocol with TCY transaction (should match)
	matched, params, err = tcyProtocol.MatchFunctionCall(tcyTx, matcher)
	assert.NoError(t, err)
	assert.True(t, matched, "TCY protocol should match TCY transaction")
	assert.Equal(t, "tcy", params["denom"])

	// Test TCY protocol with RUNE transaction (should NOT match)
	matched, _, err = tcyProtocol.MatchFunctionCall(runeTx, matcher)
	assert.NoError(t, err)
	assert.False(t, matched, "TCY protocol should NOT match RUNE transaction")
}

func TestConstraintHandlingWithBigInt(t *testing.T) {
	runeProtocol := NewRUNE()

	// Create transaction with specific amount
	testAmount := big.NewInt(100000000) // 1 RUNE
	runeTx := &ParsedThorchainTransaction{
		txHash:   "test_hash",
		fromAddr: "thor1jkndhpfauwmtn2uk7ytmruu0yywz5e66mq94e2",
		toAddr:   "thor13m8mqtcv7c2srnpr4efucwkrdu29tq3ejgu52r",
		amount:   testAmount,
		denom:    "rune",
		memo:     "test transfer",
		msgType:  "MsgSend",
	}

	// Test FIXED constraint with matching *big.Int value
	fixedConstraint := &types.ParameterConstraint{
		ParameterName: "amount",
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
			Value: &types.Constraint_FixedValue{
				FixedValue: "100000000", // Should match testAmount.String()
			},
		},
	}

	matcher := &types.PolicyFunctionMatcher{
		FunctionID:  "transfer",
		Constraints: []*types.ParameterConstraint{fixedConstraint},
	}

	// Should match because amount (100000000) equals fixed value
	matched, _, err := runeProtocol.MatchFunctionCall(runeTx, matcher)
	assert.NoError(t, err)
	assert.True(t, matched, "Should match when *big.Int amount equals fixed constraint")

	// Test FIXED constraint with non-matching *big.Int value
	nonMatchingConstraint := &types.ParameterConstraint{
		ParameterName: "amount",
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
			Value: &types.Constraint_FixedValue{
				FixedValue: "200000000", // Different from testAmount
			},
		},
	}

	matcherNonMatching := &types.PolicyFunctionMatcher{
		FunctionID:  "transfer",
		Constraints: []*types.ParameterConstraint{nonMatchingConstraint},
	}

	// Should NOT match because amount (100000000) != fixed value (200000000)
	matched, _, err = runeProtocol.MatchFunctionCall(runeTx, matcherNonMatching)
	assert.NoError(t, err)
	assert.False(t, matched, "Should NOT match when *big.Int amount differs from fixed constraint")

	// Test MAX constraint with *big.Int
	maxConstraint := &types.ParameterConstraint{
		ParameterName: "amount",
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
			Value: &types.Constraint_MaxValue{
				MaxValue: "200000000", // Higher than testAmount
			},
		},
	}

	matcherMax := &types.PolicyFunctionMatcher{
		FunctionID:  "transfer",
		Constraints: []*types.ParameterConstraint{maxConstraint},
	}

	// Should match because amount (100000000) < max value (200000000)
	matched, _, err = runeProtocol.MatchFunctionCall(runeTx, matcherMax)
	assert.NoError(t, err)
	assert.True(t, matched, "Should match when *big.Int amount is less than max constraint")
}
