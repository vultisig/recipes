// Package utxo provides common types and interfaces for UTXO-based blockchains.
package utxo

import (
	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/types"
)

// Chain represents a UTXO-based blockchain.
type Chain interface {
	types.Chain

	// ComputeTxHash computes the transaction hash after applying signatures.
	ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error)
}

// Transaction represents a parsed UTXO transaction with output access.
type Transaction interface {
	// OutputCount returns the number of outputs in the transaction.
	OutputCount() int

	// Output returns the output at the given index.
	Output(index int) (Output, error)
}

// Output represents a single transaction output.
type Output interface {
	// Value returns the output value in the smallest unit (satoshis, zatoshis, etc.).
	Value() int64

	// PkScript returns the output's public key script.
	PkScript() []byte
}

// AddressExtractor extracts an address from a PkScript for a specific chain.
type AddressExtractor interface {
	// ExtractAddress extracts the address from the given PkScript.
	ExtractAddress(pkScript []byte) (string, error)
}

