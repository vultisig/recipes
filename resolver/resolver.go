package resolver

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

type MagicResolver struct {
	treasuryConfig map[string]map[string]string
}

func NewMagicResolver() *MagicResolver {
	return &MagicResolver{
		treasuryConfig: map[string]map[string]string{
			"ethereum": {
				"eth":  "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
				"usdc": "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
				"dai":  "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
				"weth": "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			},
			"bitcoin": {
				"btc": "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
			},
			"arbitrum": {
				"eth":  "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
				"usdc": "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			},
		},
	}
}

// Main resolver method with switch case for magic constant types
func (r *MagicResolver) Resolve(magicConstant types.MagicConstant, chainID, assetID string) (string, error) {
	switch magicConstant {
	case types.MagicConstant_MAGIC_CONSTANT_VULTISIG_TREASURY:
		return r.resolveTreasury(chainID, assetID)
	default:
		return "", fmt.Errorf("unsupported magic constant: %v", magicConstant)
	}
}

// Treasury-specific resolver method
func (r *MagicResolver) resolveTreasury(chainID, assetID string) (string, error) {
	chainAddresses, exists := r.treasuryConfig[chainID]
	if !exists {
		return "", fmt.Errorf("no treasury address configured for chain %s", chainID)
	}

	address, exists := chainAddresses[assetID]
	if !exists {
		return "", fmt.Errorf("no treasury address configured for asset %s on chain %s", assetID, chainID)
	}

	return address, nil
}
