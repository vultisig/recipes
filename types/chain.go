package types

import "github.com/vultisig/mobile-tss-lib/tss"

// Chain represents a blockchain network
type Chain interface {
	// ID returns the unique identifier for the chain
	ID() string

	// Name returns a human-readable name for the chain
	Name() string

	// Description returns a brief description of the chain
	Description() string

	// SupportedProtocols returns the list of protocol IDs supported by this chain
	SupportedProtocols() []string

	// ParseTransaction decodes a raw transaction hex string into a generic DecodedTransaction object.
	ParseTransaction(txHex string) (DecodedTransaction, error)

	// GetProtocol returns a specific protocol handler supported by this chain.
	GetProtocol(id string) (Protocol, error)

	// ComputeTxHash
	// for plugins we can't use `proposedTxHex` to compute the hash because it doesn't include the signature,
	// we need to properly decode tx bytes, append signature to it, and compute hash using the particular chain library
	// `sigs` is slice, not a map, because we need to preserve the order of signatures
	// for R,S ordered apply for BTC for example
	ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error)
}
