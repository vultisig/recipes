package thorchain

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vultisig/recipes/chain/cosmos"
	"github.com/vultisig/recipes/types"
)

// NewChain creates a new THORChain chain instance.
func NewChain() *cosmos.Chain {
	return cosmos.NewChain(cosmos.ChainConfig{
		ID:           "thorchain",
		Name:         "THORChain",
		Description:  "THORChain is a decentralized liquidity protocol for cross-chain swaps.",
		Bech32Prefix: "thor",
		Protocols:    []string{"rune", "send", "thorchain_swap"},
		MessageTypeRegistry: cosmos.NewMessageTypeRegistry(map[string]cosmos.MessageType{
			cosmos.TypeUrlCosmosMsgSend:    cosmos.MessageTypeSend,
			cosmos.TypeUrlCustomMsgSend:    cosmos.MessageTypeSend,
			cosmos.TypeUrlCustomMsgDeposit: cosmos.MessageTypeDeposit,
		}),
		RegisterExtraTypes: func(ir codectypes.InterfaceRegistry) {
			// Register the generated protobuf MsgDeposit for THORChain swaps
			ir.RegisterImplementations((*sdk.Msg)(nil), &types.MsgDeposit{})
		},
		GetProtocol: func(id string) (types.Protocol, error) {
			switch id {
			case "rune", "send":
				return NewRUNE(), nil
			case "thorchain_swap":
				return NewSwap(), nil
			}
			return nil, fmt.Errorf("protocol %q not found on THORChain", id)
		},
		CustomFromExtractor: cosmos.DefaultFromExtractorWithDeposit,
	})
}

