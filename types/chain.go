package types

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
}
