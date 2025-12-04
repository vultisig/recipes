package maya

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vultisig/recipes/chain/cosmos"
	"github.com/vultisig/recipes/types"
)

// NewChain creates a new MAYAChain chain instance.
func NewChain() *cosmos.Chain {
	return cosmos.NewChain(cosmos.ChainConfig{
		ID:           "mayachain",
		Name:         "MAYAChain",
		Description:  "MAYAChain is a decentralized cross-chain liquidity protocol forked from THORChain.",
		Bech32Prefix: "maya",
		Protocols:    []string{"cacao", "send", "mayachain_swap"},
		MessageTypeRegistry: cosmos.NewMessageTypeRegistry(map[string]cosmos.MessageType{
			cosmos.TypeUrlCosmosMsgSend:    cosmos.MessageTypeSend,
			cosmos.TypeUrlCustomMsgSend:    cosmos.MessageTypeSend,
			cosmos.TypeUrlCustomMsgDeposit: cosmos.MessageTypeDeposit,
		}),
		RegisterExtraTypes: func(ir codectypes.InterfaceRegistry) {
			// Register the generated protobuf MsgDeposit for MAYAChain swaps
			ir.RegisterImplementations((*sdk.Msg)(nil), &types.MsgDeposit{})
		},
		GetProtocol: func(id string) (types.Protocol, error) {
			switch id {
			case "cacao", "send":
				return NewCACAO(), nil
			case "mayachain_swap":
				return NewSwap(), nil
			}
			return nil, fmt.Errorf("protocol %q not found on MAYAChain", id)
		},
		CustomFromExtractor: cosmos.DefaultFromExtractorWithDeposit,
	})
}

