package maya

import (
	"context"

	cosmossdk "github.com/vultisig/recipes/sdk/cosmos"
	"github.com/vultisig/mobile-tss-lib/tss"
)

// MAYAChain mainnet endpoints (REST API)
var MainnetEndpoints = []string{
	"https://mayanode.mayachain.info",
	"https://maya-api.polkachu.com",
}

// MAYAChain stagenet endpoints
var StagenetEndpoints = []string{
	"https://stagenet.mayanode.mayachain.info",
}

// SDK represents the MAYAChain SDK for transaction signing and broadcasting
type SDK struct {
	*cosmossdk.SDK
}

// NewSDK creates a new MAYAChain SDK instance with the given RPC client
func NewSDK(rpcClient cosmossdk.RPCClient) *SDK {
	return &SDK{
		SDK: cosmossdk.NewSDK(rpcClient),
	}
}

// NewMainnetSDK creates a new MAYAChain SDK instance configured for mainnet
func NewMainnetSDK() *SDK {
	return NewSDK(cosmossdk.NewHTTPRPCClient(MainnetEndpoints))
}

// NewStagenetSDK creates a new MAYAChain SDK instance configured for stagenet
func NewStagenetSDK() *SDK {
	return NewSDK(cosmossdk.NewHTTPRPCClient(StagenetEndpoints))
}

// Re-export types for convenience
type BroadcastTxResponse = cosmossdk.BroadcastTxResponse
type TxResponse = cosmossdk.TxResponse
type RPCClient = cosmossdk.RPCClient
type HTTPRPCClient = cosmossdk.HTTPRPCClient

// NewHTTPRPCClient creates a new HTTP RPC client with the given endpoints
func NewHTTPRPCClient(endpoints []string) *HTTPRPCClient {
	return cosmossdk.NewHTTPRPCClient(endpoints)
}

// GetPubKeyFromBytes creates a Cosmos public key from compressed bytes
var GetPubKeyFromBytes = cosmossdk.GetPubKeyFromBytes

// Send is a convenience method that signs and broadcasts the transaction
func (s *SDK) Send(ctx context.Context, unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse, pubKey []byte) (*BroadcastTxResponse, error) {
	return s.SDK.Send(ctx, unsignedTxBytes, signatures, pubKey)
}

