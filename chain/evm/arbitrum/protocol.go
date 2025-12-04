package arbitrum

import (
	"github.com/vultisig/recipes/chain/evm/ethereum"
	"github.com/vultisig/recipes/types"
)

// NewArbitrumETH creates a new Arbitrum ETH protocol using the shared native EVM implementation
func NewArbitrumETH() types.Protocol {
	return ethereum.NewNativeEVMProtocol(
		"arbitrum",
		"eth",
		"Arbitrum ETH",
		"Native Ether on the Arbitrum L2 blockchain",
	)
}
