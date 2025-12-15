package btc

import (
	"github.com/btcsuite/btcd/btcutil/psbt"
)

// UTXO represents an unspent transaction output.
type UTXO struct {
	TxHash   string // Transaction hash (hex string)
	Index    uint32 // Output index in the transaction
	Value    uint64 // Value in satoshis
	PkScript []byte // Optional: locking script (for witness detection)
}

// BuildResult contains the built transaction and metadata.
type BuildResult struct {
	Packet        *psbt.Packet // PSBT packet ready for signing
	SelectedUTXOs []UTXO       // Which UTXOs were selected
	Fee           uint64       // Calculated fee in satoshis
	ChangeAmount  int64        // Change output amount (0 if no change)
	ChangeIndex   int          // Index of change output (-1 if no change)
}

// PrevTxFetcher fetches previous transactions for PSBT metadata population.
// Implement this interface to provide raw transactions on-demand.
type PrevTxFetcher interface {
	GetRawTransaction(txHash string) ([]byte, error)
}

// Builder creates unsigned Bitcoin transactions.
type Builder struct {
	DustLimit int64
}

// NewBuilder creates a new Builder with the given dust limit.
func NewBuilder(dustLimit int64) *Builder {
	return &Builder{
		DustLimit: dustLimit,
	}
}

// Mainnet returns a Builder configured for Bitcoin mainnet.
func Mainnet() *Builder {
	return NewBuilder(546)
}

// Testnet returns a Builder configured for Bitcoin testnet.
func Testnet() *Builder {
	return NewBuilder(546)
}
