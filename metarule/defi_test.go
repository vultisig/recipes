package metarule

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/recipes/types"
)

func TestTryFormat_EVM_LP_Add(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ethereum.lp",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "add",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "uniswap_v3",
					},
				},
			},
			{
				ParameterName: "token0",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // USDC
					},
				},
			},
			{
				ParameterName: "token1",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // WETH
					},
				},
			},
			{
				ParameterName: "amount0",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "1000000000", // 1000 USDC
					},
				},
			},
			{
				ParameterName: "amount1",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "500000000000000000", // 0.5 ETH
					},
				},
			},
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0x1234567890123456789012345678901234567890",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.GreaterOrEqual(t, len(result), 3) // Approve token0 + approve token1 + mint

	// Approvals must come before the mint
	assert.Equal(t, "ethereum.erc20.approve", result[0].Resource)
	assert.Equal(t, "ethereum.erc20.approve", result[1].Resource)
	assert.Equal(t, "ethereum.uniswapV3_nonfungible_position_manager.mint", result[2].Resource)
	assert.Equal(t, "0xC36442b4a4522E871399CD717aBDD847Ab11FE88", result[2].Target.GetAddress())
}

func TestTryFormat_EVM_LP_CollectFees(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ethereum.lp",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "collect_fees",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "uniswap_v3",
					},
				},
			},
			{
				ParameterName: "pool",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "12345", // Token ID
					},
				},
			},
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0x1234567890123456789012345678901234567890",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	// Should be the collect function
	assert.Equal(t, "ethereum.uniswapV3_nonfungible_position_manager.collect", result[0].Resource)
}

func TestTryFormat_EVM_Lend_Supply(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ethereum.lend",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "supply",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "aave",
					},
				},
			},
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // USDC
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "1000000000",
					},
				},
			},
			{
				ParameterName: "on_behalf_of",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0x1234567890123456789012345678901234567890",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.GreaterOrEqual(t, len(result), 2) // Approve + supply

	// Approve must come before supply
	assert.Equal(t, "ethereum.erc20.approve", result[0].Resource)
	assert.Equal(t, "ethereum.aaveV3_pool.supply", result[1].Resource)
	assert.Equal(t, "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2", result[1].Target.GetAddress())
}

func TestTryFormat_EVM_Lend_Borrow(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "arbitrum.lend",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "borrow",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "aave",
					},
				},
			},
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0xaf88d065e77c8cC2239327C5EDb3A432268e5831", // USDC on Arbitrum
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "500000000", // 500 USDC
					},
				},
			},
			{
				ParameterName: "on_behalf_of",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0x1234567890123456789012345678901234567890",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1) // Borrow only (no approve needed)

	assert.Equal(t, "arbitrum.aaveV3_pool.borrow", result[0].Resource)
	assert.Equal(t, "0x794a61358D6845594F94dc1DB02A252b5b4814aD", result[0].Target.GetAddress())
}

func TestTryFormat_EVM_Perps_OpenLong(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "arbitrum.perps",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "open_long",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "gmx",
					},
				},
			},
			{
				ParameterName: "market",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0x70d95587d40A2caf56bd97485aB3Eec10Bee6336", // ETH/USD
					},
				},
			},
			{
				ParameterName: "collateral_token",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0xaf88d065e77c8cC2239327C5EDb3A432268e5831", // USDC
					},
				},
			},
			{
				ParameterName: "size_delta",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "10000000000000000000000000000000", // $10,000 size
					},
				},
			},
			{
				ParameterName: "collateral_delta",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "1000000000", // 1000 USDC collateral
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.GreaterOrEqual(t, len(result), 1)

	// Approve must come before createOrder (USDC collateral)
	require.GreaterOrEqual(t, len(result), 2)
	assert.Equal(t, "arbitrum.erc20.approve", result[0].Resource)
	assert.Equal(t, "arbitrum.gmxV2_exchange_router.createOrder", result[1].Resource)
	assert.Equal(t, "0x7C68C7866A64FA2160F78EEaE12217FFbf871fa8", result[1].Target.GetAddress())
}

func TestTryFormat_EVM_Perps_Hyperliquid_OpenLong(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "arbitrum.perps",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "open_long",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "hyperliquid",
					},
				},
			},
			{
				ParameterName: "market",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1", // BTC asset index
					},
				},
			},
			{
				ParameterName: "size_delta",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "1000000000", // Size
					},
				},
			},
			{
				ParameterName: "acceptable_price",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "50000000000", // $50,000 limit price
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.GreaterOrEqual(t, len(result), 1)

	// First rule should be order
	assert.Equal(t, "arbitrum.hyperliquid_bridge.order", result[0].Resource)
	assert.Equal(t, "0x2Df1c51E09aECF9cacB7bc98cB1742757f163dF7", result[0].Target.GetAddress())
}

func TestTryFormat_Solana_LP_Add(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.lp",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "add",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "raydium",
					},
				},
			},
			{
				ParameterName: "amount0",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "1000000000",
					},
				},
			},
			{
				ParameterName: "amount1",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "500000000",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, "solana.raydium_clmm.openPosition", result[0].Resource)
	assert.Equal(t, "CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK", result[0].Target.GetAddress())
}

func TestTryFormat_Solana_LP_CollectFees(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.lp",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "collect_fees",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "orca",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, "solana.orca_whirlpool.collectFees", result[0].Resource)
	assert.Equal(t, "whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc", result[0].Target.GetAddress())
}

func TestTryFormat_Solana_Lend_Supply(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.lend",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "supply",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "kamino",
					},
				},
			},
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v", // USDC
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "1000000000",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, "solana.kamino_lending.deposit", result[0].Resource)
	assert.Equal(t, "KLend2g3cP87ber41GYr72yfE9j6eBJYwRqVNMi6mHL", result[0].Target.GetAddress())
}

func TestTryFormat_Solana_Perps_OpenLong(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.perps",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "open_long",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "jupiter_perps",
					},
				},
			},
			{
				ParameterName: "size_delta",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "1000000000",
					},
				},
			},
			{
				ParameterName: "collateral_delta",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "100000000",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, "solana.jupiter_perpetuals.openPosition", result[0].Resource)
	assert.Equal(t, "PERPHjGBqRHArX4DySjwM6UJHiR3sWAatqfdBS2qQJu", result[0].Target.GetAddress())
}

func TestTryFormat_Solana_Perps_Drift(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.perps",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "open_short",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "drift",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, "solana.drift_protocol.placePerpOrder", result[0].Resource)
	assert.Equal(t, "dRiftyHA39MWEi3m9aunc5MzRF1JYuBsbn6VPcn33UH", result[0].Target.GetAddress())
}

func TestTryFormat_LP_MissingAction(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ethereum.lp",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "uniswap_v3",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing required constraint: action")
}

func TestTryFormat_Lend_MissingProtocol(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ethereum.lend",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "supply",
					},
				},
			},
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing required constraint: protocol")
}

func TestTryFormat_LP_UnsupportedProtocol(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ethereum.lp",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "add",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "unsupported_protocol",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported LP protocol")
}

func TestTryFormat_Perps_UnsupportedAction(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "arbitrum.perps",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "unsupported_action",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "gmx",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported perps action")
}

// ============================================================================
// Bet Meta-Protocol Tests
// ============================================================================

func TestTryFormat_EVM_Bet_Buy(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "polygon.bet",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "buy",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "polymarket",
					},
				},
			},
			{
				ParameterName: "market",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "21742633143463906290569050155826241533067272736897614950488156847949938836455", // Example token ID
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: "100000000", // 100 USDC (6 decimals)
					},
				},
			},
			{
				ParameterName: "maker",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0x1234567890123456789012345678901234567890",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.GreaterOrEqual(t, len(result), 3)

	// 1) CTF approval, 2) USDC approve, 3) fillOrder
	assert.Equal(t, "polygon.polymarket_ctf.setApprovalForAll", result[0].Resource)
	assert.Equal(t, "0x4D97DCd97eC945f40cF65F87097ACe5EA0476045", result[0].Target.GetAddress())
	assert.Equal(t, "polygon.erc20.approve", result[1].Resource)
	assert.Equal(t, "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174", result[1].Target.GetAddress())
	assert.Equal(t, "polygon.polymarket_ctf_exchange.fillOrder", result[2].Resource)
	assert.Equal(t, "0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E", result[2].Target.GetAddress())
}

func TestTryFormat_EVM_Bet_Sell(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "polygon.bet",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "sell",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "polymarket",
					},
				},
			},
			{
				ParameterName: "market",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "21742633143463906290569050155826241533067272736897614950488156847949938836455",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.GreaterOrEqual(t, len(result), 2)

	// 1) CTF approval, 2) fillOrder
	assert.Equal(t, "polygon.polymarket_ctf.setApprovalForAll", result[0].Resource)
	assert.Equal(t, "0x4D97DCd97eC945f40cF65F87097ACe5EA0476045", result[0].Target.GetAddress())
	assert.Equal(t, "polygon.polymarket_ctf_exchange.fillOrder", result[1].Resource)
}

func TestTryFormat_EVM_Bet_Cancel(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "polygon.bet",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "cancel",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "polymarket",
					},
				},
			},
			{
				ParameterName: "market",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "21742633143463906290569050155826241533067272736897614950488156847949938836455",
					},
				},
			},
			{
				ParameterName: "maker",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "0x1234567890123456789012345678901234567890",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.GreaterOrEqual(t, len(result), 1)

	// First rule should be cancelOrder
	assert.Equal(t, "polygon.polymarket_ctf_exchange.cancelOrder", result[0].Resource)
	assert.Equal(t, "0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E", result[0].Target.GetAddress())
}

func TestTryFormat_Bet_MissingAction(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "polygon.bet",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "polymarket",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing required constraint: action")
}

func TestTryFormat_Bet_MissingProtocol(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "polygon.bet",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "buy",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing required constraint: protocol")
}

func TestTryFormat_Bet_UnsupportedProtocol(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "polygon.bet",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "buy",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "unsupported_protocol",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported bet protocol")
}

func TestTryFormat_Bet_UnsupportedAction(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "polygon.bet",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "unsupported_action",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "polymarket",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported bet action")
}

func TestTryFormat_Bet_UnsupportedChain(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ethereum.bet", // Polymarket is only on Polygon
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "action",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "buy",
					},
				},
			},
			{
				ParameterName: "protocol",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "polymarket",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "polymarket CTF Exchange not available on chain")
}

