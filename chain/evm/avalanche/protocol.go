package avalanche

import (
	"github.com/vultisig/recipes/chain/evm/ethereum"
	"github.com/vultisig/recipes/types"
)

// NewAVAX creates a new Avalanche AVAX protocol using the shared native EVM implementation
func NewAVAX() types.Protocol {
	return ethereum.NewNativeEVMProtocol(
		"avalanche",
		"avax",
		"Avalanche AVAX",
		"Native AVAX on the Avalanche C-Chain",
	)
}

