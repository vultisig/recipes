package types

import "math/big"

// DecodedTransaction represents a transaction after basic chain-level parsing.
// It provides a common interface for accessing transaction details across different chains.
type DecodedTransaction interface {
	// ChainIdentifier returns the identifier of the chain this transaction belongs to (e.g., "ethereum").
	ChainIdentifier() string
	// Hash returns the transaction hash as a string.
	Hash() string
	// From returns the sender's address as a string. May be empty if sender cannot be determined or is not applicable.
	From() string
	// To returns the recipient's address (for native transfers or contract address) as a string.
	To() string
	// Value returns the amount of native currency transferred in the transaction.
	Value() *big.Int
	// Data returns the input data for contract calls or additional data for the transaction.
	Data() []byte
	// Nonce returns the transaction nonce.
	Nonce() uint64
	// GasPrice returns the price per unit of gas for the transaction.
	GasPrice() *big.Int
	// GasLimit returns the maximum amount of gas the transaction is allowed to consume.
	GasLimit() uint64
}
