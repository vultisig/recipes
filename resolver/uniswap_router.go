package resolver

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// Uniswap Universal Router addresses per chain
// Source: https://docs.uniswap.org/contracts/v3/reference/deployments
// Note: Each chain has its own Universal Router address
var uniswapUniversalRouters = map[string]string{
	"Ethereum": "0x66a9893cc07d91d95644aedd05d03f95e1dba8af",
	"Polygon":  "0x1095692a6237d83c6a72f3f5efedb9a670c49223",
	"Arbitrum": "0x5E325eDA8064b456f4781070C0738d849c824258", // Universal Router on Arbitrum
	"Optimism": "0x851116d9223fabed8e56c0e6b8ad0c31d98b3507",
	"Base":     "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD", // Same as old Ethereum address
	"BSC":      "0x1906c1d672b88cd1b9ac7593301ca990f94eae07",
}

// Uniswap SwapRouter02 addresses (fallback for chains without Universal Router)
var uniswapSwapRouter02 = map[string]string{
	"Ethereum": "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45",
	"Polygon":  "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45",
	"Arbitrum": "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45",
	"Optimism": "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45",
	"Base":     "0x2626664c2603336E57B271c5C0b26F421741e481",
	"BSC":      "0xB971eF87ede563556b2ED4b1C0b0019111Dd85d2",
}

type UniswapRouterResolver struct{}

func NewUniswapRouterResolver() Resolver {
	return &UniswapRouterResolver{}
}

func (r *UniswapRouterResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_UNISWAP_UNIVERSAL_ROUTER
}

func (r *UniswapRouterResolver) Resolve(constant types.MagicConstant, chainID, _ string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("UniswapRouterResolver does not support type: %v", constant)
	}

	// Try Universal Router first
	if router, ok := uniswapUniversalRouters[chainID]; ok {
		return router, "", nil
	}

	// Fallback to SwapRouter02
	if router, ok := uniswapSwapRouter02[chainID]; ok {
		return router, "", nil
	}

	return "", "", fmt.Errorf("Uniswap not supported on chain %s", chainID)
}

