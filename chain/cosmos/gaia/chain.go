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
		Protocols:    []string{"atom", "send", "staking_redelegate", "staking_withdraw_rewards"},
		MessageTypeRegistry: cosmos.NewMessageTypeRegistry(map[string]cosmos.MessageType{
			cosmos.TypeUrlCosmosMsgSend:                    cosmos.MessageTypeSend,
			cosmos.TypeUrlCosmosMsgBeginRedelegate:         cosmos.MessageTypeBeginRedelegate,
			cosmos.TypeUrlCosmosMsgWithdrawDelegatorReward: cosmos.MessageTypeWithdrawDelegatorReward,
		}),
		GetProtocol: func(id string) (types.Protocol, error) {
			switch id {
			case "atom", "send":
				return NewATOM(), nil
			case "staking_redelegate":
				return NewStakingRedelegate(), nil
			case "staking_withdraw_rewards":
				return NewStakingWithdrawRewards(), nil
			}
			return nil, fmt.Errorf("protocol %q not found on Cosmos", id)
		},
	})
}
