package swap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelayQuoteRequest_SolanaUseSharedAccounts(t *testing.T) {
	useShared := true

	req := relayQuoteRequest{
		User:              "test",
		OriginChainID:     792703809, // Solana
		DestinationChainID: 792703809,
		UseSharedAccounts: &useShared,
	}

	assert.NotNil(t, req.UseSharedAccounts)
	assert.True(t, *req.UseSharedAccounts)
}

func TestRelayQuoteRequest_NonSolanaOmitsSharedAccounts(t *testing.T) {
	req := relayQuoteRequest{
		User:              "test",
		OriginChainID:     1, // Ethereum
		DestinationChainID: 1,
	}

	assert.Nil(t, req.UseSharedAccounts, "Non-Solana requests should not set UseSharedAccounts")
}
