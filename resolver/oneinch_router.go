package resolver

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// 1inch Aggregation Router V6 addresses per chain
// Source: https://docs.1inch.io/docs/aggregation-protocol/smart-contract-addresses/
// Note: Each chain has its own router address - not all are the same
var oneInchRouters = map[string]string{
	"Ethereum":  "0x111111125421cA6dc452d289314280a0f8842A65",
	"BSC":       "0x111111125421cA6dc452d289314280a0f8842A65",
	"Polygon":   "0x111111125421cA6dc452d289314280a0f8842A65",
	"Avalanche": "0x652747cb44D5fC52799c3DaEa613c52625588AB5", // Different address on Avalanche
	"Arbitrum":  "0x6b0CE50D408d27ABA09F7e96Ac437011D8CDFbB8", // Different address on Arbitrum
	"Optimism":  "0x111111125421cA6dc452d289314280a0f8842A65",
	"Base":      "0x111111125421cA6dc452d289314280a0f8842A65",
	"Gnosis":    "0xed6c1002450cbf418e96d16361cbed3a84366c43", // Different address on Gnosis
	// Note: Fantom removed - 1inch uses Router V4 on Fantom, not V6, and is being deprecated
}

type OneInchRouterResolver struct{}

func NewOneInchRouterResolver() Resolver {
	return &OneInchRouterResolver{}
}

func (r *OneInchRouterResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_ONEINCH_ROUTER
}

func (r *OneInchRouterResolver) Resolve(constant types.MagicConstant, chainID, _ string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("OneInchRouterResolver does not support type: %v", constant)
	}

	router, ok := oneInchRouters[chainID]
	if !ok {
		return "", "", fmt.Errorf("1inch not supported on chain %s", chainID)
	}

	return router, "", nil
}

