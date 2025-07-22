package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestThorchainIntegration(t *testing.T) {
	// Test that Thorchain can be retrieved from the registry
	thorchainFromRegistry, err := GetChain("thorchain")
	require.NoError(t, err)

	assert.Equal(t, "thorchain", thorchainFromRegistry.ID())
	assert.Equal(t, "Thorchain", thorchainFromRegistry.Name())

	// Test that supported protocols are accessible
	protocols := thorchainFromRegistry.SupportedProtocols()
	assert.Contains(t, protocols, "rune")
	assert.Contains(t, protocols, "tcy")

	// Test that protocols can be retrieved
	runeProtocol, err := thorchainFromRegistry.GetProtocol("rune")
	require.NoError(t, err)
	assert.Equal(t, "rune", runeProtocol.ID())

	tcyProtocol, err := thorchainFromRegistry.GetProtocol("tcy")
	require.NoError(t, err)
	assert.Equal(t, "tcy", tcyProtocol.ID())
}

func TestThorchainInRegistryList(t *testing.T) {
	// Test that Thorchain appears in the list of all chains
	allChains := DefaultRegistry.List()

	var foundThorchain bool
	for _, c := range allChains {
		if c.ID() == "thorchain" {
			foundThorchain = true
			break
		}
	}

	assert.True(t, foundThorchain, "Thorchain should be in the registry")

	// Verify we have at least 3 chains (bitcoin, ethereum, thorchain)
	assert.GreaterOrEqual(t, len(allChains), 3, "Should have at least 3 chains registered")
}

func TestThorchainTransactionParsing(t *testing.T) {
	// Test that the chain can parse transactions
	thorchainFromRegistry, err := GetChain("thorchain")
	require.NoError(t, err)

	// Test parsing with invalid protobuf data (should return error now)
	txHex := "0x1234567890abcdef1234567890abcdef"
	_, err = thorchainFromRegistry.ParseTransaction(txHex)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode protobuf transaction")
}

func TestThorchainProtocolFunctionality(t *testing.T) {
	// Get Thorchain from registry and test protocol functionality
	thorchainFromRegistry, err := GetChain("thorchain")
	require.NoError(t, err)

	// Test RUNE protocol functionality
	runeProtocol, err := thorchainFromRegistry.GetProtocol("rune")
	require.NoError(t, err)

	// Test that the protocol has the expected functions
	functions := runeProtocol.Functions()
	assert.Len(t, functions, 1)

	transferFunc, err := runeProtocol.GetFunction("transfer")
	require.NoError(t, err)
	assert.Equal(t, "transfer", transferFunc.ID)
	assert.Equal(t, "Transfer RUNE", transferFunc.Name)

	// Verify parameters
	params := transferFunc.Parameters
	assert.Len(t, params, 3)

	paramNames := make(map[string]bool)
	for _, param := range params {
		paramNames[param.Name] = true
	}

	assert.True(t, paramNames["recipient"], "Should have recipient parameter")
	assert.True(t, paramNames["amount"], "Should have amount parameter")
	assert.True(t, paramNames["memo"], "Should have memo parameter")
}
