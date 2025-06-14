package thorchain

import (
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

func TestThorchainWithNetwork(t *testing.T) {
	// Test with custom network
	testnetChain := NewThorchainWithNetwork("testnet")
	assert.Equal(t, "thorchain", testnetChain.ID())
	assert.Equal(t, "Thorchain", testnetChain.Name())

	// Test that it still supports the same protocols
	protocols := testnetChain.SupportedProtocols()
	assert.Contains(t, protocols, "rune")
	assert.Contains(t, protocols, "tcy")
}
