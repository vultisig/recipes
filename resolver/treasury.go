package resolver

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

type TreasuryResolver struct {
	treasuryConfig map[string]map[string]string
}

// NewTreasuryResolver takes a config param to load configs from external sources
func NewTreasuryResolver(config map[string]map[string]string) Resolver {
	return &TreasuryResolver{treasuryConfig: config}
}

// NewDefaultTreasuryResolver returns a resolver preloaded with default treasury addresses.
func NewDefaultTreasuryResolver() Resolver {
	return NewTreasuryResolver(defaultTreasuryConfig())
}

func defaultTreasuryConfig() map[string]map[string]string {
	return map[string]map[string]string{
		"ethereum": {
			"eth":     "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			"usdc":    "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			"dai":     "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			"weth":    "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			"default": "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
		},
		"bitcoin": {
			"btc":     "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
			"default": "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
		},
		"arbitrum": {
			"eth":     "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			"usdc":    "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			"default": "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
		},
	}
}

func (r *TreasuryResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_MAGIC_CONSTANT_VULTISIG_TREASURY
}

func (r *TreasuryResolver) Resolve(constant types.MagicConstant, chainID, assetID string) (string, error) {
	if !r.Supports(constant) {
		return "", fmt.Errorf("TreasuryResolver does not support type: %v", constant)
	}
	chainAddresses, exists := r.treasuryConfig[chainID]
	if !exists {
		return "", fmt.Errorf("no treasury address configured for chain %s", chainID)
	}

	address, exists := chainAddresses[assetID]
	if !exists {
		// Try fallback to default
		if defaultAddress, defaultExists := chainAddresses["default"]; defaultExists {
			return defaultAddress, nil
		}
		return "", fmt.Errorf("no treasury address configured for asset %s on chain %s", assetID, chainID)
	}

	return address, nil
}
