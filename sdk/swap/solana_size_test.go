package swap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolanaMaxTxSizeConstant(t *testing.T) {
	assert.Equal(t, 1232, solanaMaxTxSize, "Solana max tx size should be 1232 bytes")
}

func TestSolanaRetryMaxRouteLengthConstant(t *testing.T) {
	assert.Equal(t, 2, solanaRetryMaxRouteLength, "Retry route length should be 2 hops")
}

func TestRelayQuoteRequest_SolanaFields(t *testing.T) {
	maxRoute := 3
	useShared := true

	req := relayQuoteRequest{
		User:              "test",
		OriginChainID:     792703809, // Solana
		DestinationChainID: 792703809,
		MaxRouteLength:    &maxRoute,
		UseSharedAccounts: &useShared,
	}

	assert.NotNil(t, req.MaxRouteLength)
	assert.Equal(t, 3, *req.MaxRouteLength)
	assert.NotNil(t, req.UseSharedAccounts)
	assert.True(t, *req.UseSharedAccounts)
}

func TestRelayQuoteRequest_NonSolanaOmitsFields(t *testing.T) {
	req := relayQuoteRequest{
		User:              "test",
		OriginChainID:     1, // Ethereum
		DestinationChainID: 1,
	}

	assert.Nil(t, req.MaxRouteLength, "Non-Solana requests should not set MaxRouteLength")
	assert.Nil(t, req.UseSharedAccounts, "Non-Solana requests should not set UseSharedAccounts")
}
