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
			"eth":     "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
			"usdc":    "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
			"dai":     "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
			"weth":    "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
			"default": "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
		},
		"bitcoin": {
			"btc":     "bc1qelza2cr7w92dmzgkmhdn0a4vcqpe9rfpfknr6a",
			"default": "bc1qelza2cr7w92dmzgkmhdn0a4vcqpe9rfpfknr6a",
		},
		"arbitrum": {
			"eth":     "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
			"usdc":    "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
			"default": "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
		},
	}
}

func (r *TreasuryResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_VULTISIG_TREASURY
}

func (r *TreasuryResolver) Resolve(constant types.MagicConstant, chainID, assetID string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("TreasuryResolver does not support type: %v", constant)
	}
	chainAddresses, exists := r.treasuryConfig[chainID]
	if !exists {
		return "", "", fmt.Errorf("no treasury address configured for chain %s", chainID)
	}

	address, exists := chainAddresses[assetID]
	if !exists {
		// Try fallback to default
		if defaultAddress, defaultExists := chainAddresses["default"]; defaultExists {
			return defaultAddress, "", nil
		}
		return "", "", fmt.Errorf("no treasury address configured for asset %s on chain %s", assetID, chainID)
	}
	// TODO implement memo when supported chains require it
	return address, "", nil
}
