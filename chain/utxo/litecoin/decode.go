package litecoin

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/vultisig/recipes/types"
)

// LitecoinMainNetParams holds Litecoin-specific network parameters.
var LitecoinMainNetParams = &chaincfg.Params{
	Name:             "mainnet",
	Net:              0xdbb6c0fb, // LTC mainnet magic
	PubKeyHashAddrID: 0x30,       // L prefix for P2PKH
	ScriptHashAddrID: 0x32,       // M prefix for P2SH
	Bech32HRPSegwit:  "ltc",      // ltc1 prefix for SegWit
}

// ParsedLitecoinTransaction implements the types.DecodedTransaction interface for Litecoin.
type ParsedLitecoinTransaction struct {
	tx      *wire.MsgTx
	txHash  string
	rawData []byte
	network *chaincfg.Params
}

// ChainIdentifier returns "litecoin".
func (p *ParsedLitecoinTransaction) ChainIdentifier() string { return "litecoin" }

// Hash returns the transaction hash.
func (p *ParsedLitecoinTransaction) Hash() string { return p.txHash }

// From returns empty string as Litecoin doesn't have a simple from address in unsigned transactions
func (p *ParsedLitecoinTransaction) From() string { return "" }

// To returns the first output address if it can be decoded, otherwise empty
func (p *ParsedLitecoinTransaction) To() string {
	if len(p.tx.TxOut) > 0 {
		// Extract address from the first output
		_, addrs, _, err := txscript.ExtractPkScriptAddrs(p.tx.TxOut[0].PkScript, p.network)
		if err == nil && len(addrs) > 0 {
			return addrs[0].EncodeAddress()
		}
	}
	return ""
}

// Value returns the value of the first output in litoshi (1 LTC = 100,000,000 litoshi)
func (p *ParsedLitecoinTransaction) Value() *big.Int {
	if len(p.tx.TxOut) > 0 {
		return new(big.Int).SetInt64(p.tx.TxOut[0].Value)
	}
	return big.NewInt(0)
}

// Data returns the raw transaction data
func (p *ParsedLitecoinTransaction) Data() []byte { return p.rawData }

// Nonce returns 0 as Litecoin doesn't use nonces like Ethereum
func (p *ParsedLitecoinTransaction) Nonce() uint64 { return 0 }

// GasPrice returns 0 as Litecoin doesn't use gas
func (p *ParsedLitecoinTransaction) GasPrice() *big.Int { return big.NewInt(0) }

// GasLimit returns 0 as Litecoin doesn't use gas
func (p *ParsedLitecoinTransaction) GasLimit() uint64 { return 0 }

// ParseLitecoinTransaction decodes a raw Litecoin transaction from its hex representation.
func ParseLitecoinTransaction(txHex string) (types.DecodedTransaction, error) {
	return ParseLitecoinTransactionWithNetwork(txHex, LitecoinMainNetParams)
}

// ParseLitecoinTransactionWithNetwork decodes a raw Litecoin transaction with a specific network.
func ParseLitecoinTransactionWithNetwork(txHex string, network *chaincfg.Params) (types.DecodedTransaction, error) {
	rawTxBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex transaction: %w", err)
	}

	// Deserialize the transaction
	tx := &wire.MsgTx{}
	err = tx.Deserialize(bytes.NewReader(rawTxBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize transaction: %w", err)
	}

	// Calculate transaction hash
	txHash := tx.TxHash().String()

	return &ParsedLitecoinTransaction{
		tx:      tx,
		txHash:  txHash,
		rawData: rawTxBytes,
		network: network,
	}, nil
}

// GetTransaction returns the underlying wire.MsgTx
func (p *ParsedLitecoinTransaction) GetTransaction() *wire.MsgTx {
	return p.tx
}

// GetNetwork returns the network parameters
func (p *ParsedLitecoinTransaction) GetNetwork() *chaincfg.Params {
	return p.network
}

// GetOutputCount returns the number of outputs
func (p *ParsedLitecoinTransaction) GetOutputCount() int {
	return len(p.tx.TxOut)
}

// GetInputCount returns the number of inputs
func (p *ParsedLitecoinTransaction) GetInputCount() int {
	return len(p.tx.TxIn)
}

// IsSegWit returns true if the transaction uses SegWit
func (p *ParsedLitecoinTransaction) IsSegWit() bool {
	return p.tx.HasWitness()
}

