package btc

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/vultisig/recipes/sdk/utxo"
)

// Re-export common UTXO types for convenience.
// Users can import either sdk/btc or sdk/utxo.
type (
	// UTXO represents an unspent transaction output.
	UTXO = utxo.UTXO

	// Output represents a transaction output to create.
	Output = utxo.Output
)

// PreviousTx contains raw transaction data for PSBT metadata population.
// Only needed if UTXOs don't include PkScript.
type PreviousTx struct {
	TxHash string // Transaction hash
	RawTx  []byte // Raw transaction bytes
}

// ChainConfig contains Bitcoin-specific parameters for transaction building.
type ChainConfig struct {
	Network   *chaincfg.Params // btcd network params
	DustLimit int64            // Minimum output value (546 for Bitcoin mainnet)
}

// Predefined configurations for Bitcoin networks.
var (
	MainnetConfig = ChainConfig{
		Network:   &chaincfg.MainNetParams,
		DustLimit: 546,
	}
	TestnetConfig = ChainConfig{
		Network:   &chaincfg.TestNet3Params,
		DustLimit: 546,
	}
	RegtestConfig = ChainConfig{
		Network:   &chaincfg.RegressionNetParams,
		DustLimit: 546,
	}
)

// BuildParams contains all inputs needed to build a Bitcoin transaction.
type BuildParams struct {
	// Required
	UTXOs      []UTXO   // Available UTXOs to spend from
	FeeRate    uint64   // Fee rate in sats/vbyte
	FromPubKey []byte   // Compressed public key (33 bytes) for PSBT derivation
	Outputs    []Output // Desired outputs (recipient, memo, etc.)

	// Optional
	ChangeAddress string       // If empty, derived from FromPubKey as P2WPKH
	PreviousTxs   []PreviousTx // Raw txs for PSBT metadata (if UTXOs lack PkScript)
	SelectAll     bool         // If true, use all UTXOs (for consolidation)
}

// SendParams is a convenience struct for simple send transactions.
type SendParams struct {
	UTXOs         []UTXO       // Available UTXOs to spend from
	FeeRate       uint64       // Fee rate in sats/vbyte
	FromPubKey    []byte       // Compressed public key (33 bytes)
	ToAddress     string       // Recipient address
	Amount        uint64       // Amount to send in satoshis
	ChangeAddress string       // Optional: explicit change address
	PreviousTxs   []PreviousTx // Optional: raw txs for PSBT metadata
}

// SwapParams is a convenience struct for THORChain-style swap transactions.
type SwapParams struct {
	UTXOs         []UTXO       // Available UTXOs to spend from
	FeeRate       uint64       // Fee rate in sats/vbyte
	FromPubKey    []byte       // Compressed public key (33 bytes)
	VaultAddress  string       // THORChain/Maya inbound vault address
	Amount        uint64       // Amount to swap in satoshis
	Memo          string       // Swap memo (will be OP_RETURN output)
	ChangeAddress string       // Optional: explicit change address
	PreviousTxs   []PreviousTx // Optional: raw txs for PSBT metadata
}

// BuildResult contains the built transaction and metadata.
type BuildResult struct {
	PSBT            []byte // Serialized PSBT ready for signing
	UnsignedTxBytes []byte // Raw unsigned tx bytes (for policy evaluation)
	SelectedUTXOs   []UTXO // Which UTXOs were selected
	Fee             uint64 // Calculated fee in satoshis
	ChangeAmount    int64  // Change output amount (0 if no change)
	ChangeIndex     int    // Index of change output (-1 if no change)
	EstimatedVBytes int    // Estimated transaction virtual size
}
