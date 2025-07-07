package arbitrum

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vultisig/mobile-tss-lib/tss"
)

func TestArbitrum_BasicChainProperties(t *testing.T) {
	arb := NewArbitrum()

	// Test chain properties
	require.Equal(t, "arbitrum", arb.ID())
	require.Equal(t, "Arbitrum One", arb.Name())
	require.Contains(t, arb.Description(), "Arbitrum One")
	require.Contains(t, arb.Description(), "Optimistic Rollup")
}

func TestArbitrum_SupportedProtocols(t *testing.T) {
	arb := NewArbitrum()
	protocols := arb.SupportedProtocols()

	// Should at least support ETH
	require.Contains(t, protocols, "eth")
	require.Greater(t, len(protocols), 0)
}

func TestArbitrum_ParseTransaction(t *testing.T) {
	arb := NewArbitrum()

	// Use the same transaction hex as Ethereum test but expect Arbitrum chain identifier
	txHex := "0x00ec80872386f26fc10000830f424094b0b00000000000000000000000000000000000018806f05b59d3b2000080"

	tx, err := arb.ParseTransaction(txHex)
	require.NoError(t, err)
	require.NotNil(t, tx)

	// Should return "42161" as chain identifier, not "ethereum"
	require.Equal(t, "42161", tx.ChainIdentifier())
	require.NotEmpty(t, tx.Hash())
}

func TestArbitrum_ComputeTxHash(t *testing.T) {
	// Use same test pattern as ethereum/ethereum_test.go but with Arbitrum chain
	arb := NewArbitrum()

	// Use a basic transaction hex for testing
	txHex := "0x00ec80872386f26fc10000830f424094b0b00000000000000000000000000000000000018806f05b59d3b2000080"

	txHash, err := arb.ComputeTxHash(txHex, []tss.KeysignResponse{{
		R:          "d55e81731a80a10a66475fb52021b03b9173359a3220c10db76739b674355f7a",
		S:          "6ebf679597e97da64d048e28fe418b2ca0ef08c2a0583b97d11703dc11cd727b",
		RecoveryID: "01",
	}})

	require.NoError(t, err)
	require.NotEmpty(t, txHash)
	require.True(t, len(txHash) > 0)
}
