package gaia

import (
	"context"

	cosmossdk "github.com/vultisig/recipes/sdk/cosmos"
	"github.com/vultisig/mobile-tss-lib/tss"
)

// Cosmos mainnet endpoints (REST API)
var MainnetEndpoints = []string{
	"https://cosmos-rest.publicnode.com",
	"https://rest.cosmos.directory/cosmoshub",
	"https://lcd-cosmoshub.keplr.app",
}

// Cosmos testnet endpoints
var TestnetEndpoints = []string{
	"https://rest.sentry-01.theta-testnet.polypore.xyz",
}

// SDK represents the Cosmos (GAIA) SDK for transaction signing and broadcasting
type SDK struct {
	*cosmossdk.SDK
}

// NewSDK creates a new Cosmos SDK instance with the given RPC client
func NewSDK(rpcClient cosmossdk.RPCClient) *SDK {
	return &SDK{
		SDK: cosmossdk.NewSDK(rpcClient),
	}
}

// NewMainnetSDK creates a new Cosmos SDK instance configured for mainnet
func NewMainnetSDK() *SDK {
	return NewSDK(cosmossdk.NewHTTPRPCClient(MainnetEndpoints))
}

// NewTestnetSDK creates a new Cosmos SDK instance configured for testnet
func NewTestnetSDK() *SDK {
	return NewSDK(cosmossdk.NewHTTPRPCClient(TestnetEndpoints))
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

