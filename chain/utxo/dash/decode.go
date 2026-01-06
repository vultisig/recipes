package dash

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

// DashMainNetParams holds Dash-specific network parameters.
var DashMainNetParams = &chaincfg.Params{
	Name:             "mainnet",
	Net:              0xbf0c6bbd, // Dash mainnet magic
	PubKeyHashAddrID: 0x4c,       // X prefix for P2PKH
	ScriptHashAddrID: 0x10,       // 7 prefix for P2SH
	Bech32HRPSegwit:  "",         // Dash doesn't use bech32
}

// ParsedDashTransaction implements the types.DecodedTransaction interface for Dash.
type ParsedDashTransaction struct {
	tx      *wire.MsgTx
	txHash  string
	rawData []byte
	network *chaincfg.Params
}

// ChainIdentifier returns "dash".
func (p *ParsedDashTransaction) ChainIdentifier() string { return "dash" }

// Hash returns the transaction hash.
func (p *ParsedDashTransaction) Hash() string { return p.txHash }

// From returns empty string as Dash doesn't have a simple from address in unsigned transactions
func (p *ParsedDashTransaction) From() string { return "" }

// To returns the first output address if it can be decoded, otherwise empty
func (p *ParsedDashTransaction) To() string {
	if len(p.tx.TxOut) > 0 {
		// Extract address from the first output
		_, addrs, _, err := txscript.ExtractPkScriptAddrs(p.tx.TxOut[0].PkScript, p.network)
		if err == nil && len(addrs) > 0 {
			return addrs[0].EncodeAddress()
		}
	}
	return ""
}

// Value returns the value of the first output in duffs (1 DASH = 100,000,000 duffs)
func (p *ParsedDashTransaction) Value() *big.Int {
	if len(p.tx.TxOut) > 0 {
		return new(big.Int).SetInt64(p.tx.TxOut[0].Value)
	}
	return big.NewInt(0)
}

// Data returns the raw transaction data
func (p *ParsedDashTransaction) Data() []byte { return p.rawData }

// Nonce returns 0 as Dash doesn't use nonces like Ethereum
func (p *ParsedDashTransaction) Nonce() uint64 { return 0 }

// GasPrice returns 0 as Dash doesn't use gas
func (p *ParsedDashTransaction) GasPrice() *big.Int { return big.NewInt(0) }

// GasLimit returns 0 as Dash doesn't use gas
func (p *ParsedDashTransaction) GasLimit() uint64 { return 0 }

// ParseDashTransaction decodes a raw Dash transaction from its hex representation.
func ParseDashTransaction(txHex string) (types.DecodedTransaction, error) {
	return ParseDashTransactionWithNetwork(txHex, DashMainNetParams)
}

// ParseDashTransactionWithNetwork decodes a raw Dash transaction with a specific network.
func ParseDashTransactionWithNetwork(txHex string, network *chaincfg.Params) (types.DecodedTransaction, error) {
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

	return &ParsedDashTransaction{
		tx:      tx,
		txHash:  txHash,
		rawData: rawTxBytes,
		network: network,
	}, nil
}

// GetTransaction returns the underlying wire.MsgTx
func (p *ParsedDashTransaction) GetTransaction() *wire.MsgTx {
	return p.tx
}

// GetNetwork returns the network parameters
func (p *ParsedDashTransaction) GetNetwork() *chaincfg.Params {
	return p.network
}

// GetOutputCount returns the number of outputs
func (p *ParsedDashTransaction) GetOutputCount() int {
	return len(p.tx.TxOut)
}

// GetInputCount returns the number of inputs
func (p *ParsedDashTransaction) GetInputCount() int {
	return len(p.tx.TxIn)
}

