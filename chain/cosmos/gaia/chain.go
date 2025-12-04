package gaia

import (
	"fmt"

	"github.com/vultisig/recipes/chain/cosmos"
	"github.com/vultisig/recipes/types"
)

// NewChain creates a new Cosmos/GAIA chain instance.
func NewChain() *cosmos.Chain {
	return cosmos.NewChain(cosmos.ChainConfig{
		ID:           "cosmos",
		Name:         "Cosmos",
		Description:  "Cosmos (GAIA) is the hub of the Cosmos network, enabling cross-chain communication via IBC.",
		Bech32Prefix: "cosmos",
		Protocols:    []string{"atom", "send"},
		MessageTypeRegistry: cosmos.NewMessageTypeRegistry(map[string]cosmos.MessageType{
			cosmos.TypeUrlCosmosMsgSend: cosmos.MessageTypeSend,
		}),
		GetProtocol: func(id string) (types.Protocol, error) {
			switch id {
			case "atom", "send":
				return NewATOM(), nil
			}
			return nil, fmt.Errorf("protocol %q not found on Cosmos", id)
		},
	})
}

