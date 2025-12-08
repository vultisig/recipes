// Package utxo provides shared types and utilities for UTXO-based blockchain SDKs.
// This package contains chain-agnostic functionality that can be used by Bitcoin,
// Litecoin, Dogecoin, Bitcoin Cash, Zcash, and other UTXO-based chains.
package utxo

// UTXO represents an unspent transaction output.
// This is a universal type that works across all UTXO chains.
type UTXO struct {
	TxHash   string // Transaction hash (hex string, no 0x prefix)
	Index    uint32 // Output index in the transaction
	Value    uint64 // Value in the chain's smallest unit (satoshis, litoshis, etc.)
	PkScript []byte // Optional: locking script (for input type detection)
}

// Output represents a transaction output to create.
type Output struct {
	Address string // Destination address (format depends on chain)
	Amount  int64  // Amount in the chain's smallest unit
	Data    []byte // Optional: OP_RETURN data (if Address is empty)
}

// SelectionResult contains the result of UTXO selection.
type SelectionResult struct {
	Selected     []UTXO // Selected UTXOs
	TotalValue   uint64 // Total value of selected UTXOs
	TargetAmount uint64 // Original target amount
	Fee          uint64 // Estimated fee
	Change       int64  // Change amount (can be negative if insufficient funds)
}

// ChainParams contains chain-specific parameters for UTXO operations.
// Each chain implementation should define its own ChainParams.
type ChainParams struct {
	// Name is the human-readable chain name (e.g., "Bitcoin", "Litecoin")
	Name string

	// Ticker is the chain's ticker symbol (e.g., "BTC", "LTC")
	Ticker string

	// DustLimit is the minimum output value in the smallest unit.
	// Outputs below this value are considered "dust" and may not be relayed.
	DustLimit int64

	// SupportsSegWit indicates whether the chain supports SegWit transactions.
	// This affects size estimation (vbytes vs bytes).
	SupportsSegWit bool

	// SigHashType is the default signature hash type for the chain.
	// Most chains use 0x01 (SIGHASH_ALL), but Bitcoin Cash uses 0x41.
	SigHashType byte

	// InputSizes contains the sizes for different input types.
	// Keys: "p2pkh", "p2wpkh", "p2tr", "p2wsh", "p2sh"
	InputSizes map[string]int

	// OutputSizes contains the sizes for different output types.
	// Keys: "p2pkh", "p2wpkh", "p2tr", "p2wsh", "p2sh", "opreturn"
	OutputSizes map[string]int

	// TxOverhead is the base transaction overhead in bytes/vbytes.
	TxOverhead int
}

// Common signature hash types
const (
	SigHashAll         byte = 0x01
	SigHashAllForkID   byte = 0x41 // Bitcoin Cash
	SigHashNone        byte = 0x02
	SigHashSingle      byte = 0x03
	SigHashAnyOneCanPay byte = 0x80
)

// Predefined chain parameters for common UTXO chains.
var (
	// BitcoinParams contains parameters for Bitcoin mainnet.
	BitcoinParams = ChainParams{
		Name:           "Bitcoin",
		Ticker:         "BTC",
		DustLimit:      546,
		SupportsSegWit: true,
		SigHashType:    SigHashAll,
		InputSizes: map[string]int{
			"p2pkh":  148,
			"p2wpkh": 68,
			"p2tr":   58,
			"p2wsh":  91,
			"p2sh":   91,
		},
		OutputSizes: map[string]int{
			"p2pkh":    34,
			"p2wpkh":   31,
			"p2tr":     43,
			"p2wsh":    43,
			"p2sh":     32,
			"opreturn": 11, // Base size, add data length
		},
		TxOverhead: 11,
	}

	// LitecoinParams contains parameters for Litecoin mainnet.
	LitecoinParams = ChainParams{
		Name:           "Litecoin",
		Ticker:         "LTC",
		DustLimit:      100000, // 0.001 LTC
		SupportsSegWit: true,
		SigHashType:    SigHashAll,
		InputSizes: map[string]int{
			"p2pkh":  148,
			"p2wpkh": 68,
			"p2tr":   58,
			"p2wsh":  91,
			"p2sh":   91,
		},
		OutputSizes: map[string]int{
			"p2pkh":    34,
			"p2wpkh":   31,
			"p2tr":     43,
			"p2wsh":    43,
			"p2sh":     32,
			"opreturn": 11,
		},
		TxOverhead: 11,
	}

	// DogecoinParams contains parameters for Dogecoin mainnet.
	DogecoinParams = ChainParams{
		Name:           "Dogecoin",
		Ticker:         "DOGE",
		DustLimit:      100000000, // 1 DOGE
		SupportsSegWit: false,      // Dogecoin doesn't support SegWit
		SigHashType:    SigHashAll,
		InputSizes: map[string]int{
			"p2pkh": 148,
			"p2sh":  91,
		},
		OutputSizes: map[string]int{
			"p2pkh":    34,
			"p2sh":     32,
			"opreturn": 11,
		},
		TxOverhead: 10,
	}

	// BitcoinCashParams contains parameters for Bitcoin Cash mainnet.
	BitcoinCashParams = ChainParams{
		Name:           "Bitcoin Cash",
		Ticker:         "BCH",
		DustLimit:      546,
		SupportsSegWit: false, // BCH doesn't support SegWit
		SigHashType:    SigHashAllForkID,
		InputSizes: map[string]int{
			"p2pkh": 148,
			"p2sh":  91,
		},
		OutputSizes: map[string]int{
			"p2pkh":    34,
			"p2sh":     32,
			"opreturn": 11,
		},
		TxOverhead: 10,
	}

	// ZcashParams contains parameters for Zcash mainnet (transparent only).
	ZcashParams = ChainParams{
		Name:           "Zcash",
		Ticker:         "ZEC",
		DustLimit:      546,
		SupportsSegWit: false, // Zcash doesn't use SegWit
		SigHashType:    SigHashAll,
		InputSizes: map[string]int{
			"p2pkh": 148,
			"p2sh":  91,
		},
		OutputSizes: map[string]int{
			"p2pkh":    34,
			"p2sh":     32,
			"opreturn": 11,
		},
		TxOverhead: 18, // Zcash has additional overhead (version group, expiry, etc.)
	}
)

// GetInputSize returns the size for a given input type.
// Returns the p2pkh size as default if type not found.
func (p ChainParams) GetInputSize(inputType string) int {
	if size, ok := p.InputSizes[inputType]; ok {
		return size
	}
	// Default to P2WPKH for SegWit chains, P2PKH for others
	if p.SupportsSegWit {
		if size, ok := p.InputSizes["p2wpkh"]; ok {
			return size
		}
	}
	return p.InputSizes["p2pkh"]
}

// GetOutputSize returns the size for a given output type.
// Returns the p2pkh size as default if type not found.
func (p ChainParams) GetOutputSize(outputType string) int {
	if size, ok := p.OutputSizes[outputType]; ok {
		return size
	}
	// Default to P2WPKH for SegWit chains, P2PKH for others
	if p.SupportsSegWit {
		if size, ok := p.OutputSizes["p2wpkh"]; ok {
			return size
		}
	}
	return p.OutputSizes["p2pkh"]
}
