package dogecoin

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

// DogecoinMainNetParams holds Dogecoin-specific network parameters.
var DogecoinMainNetParams = &chaincfg.Params{
	Name:             "mainnet",
	Net:              0xc0c0c0c0, // DOGE mainnet magic
	PubKeyHashAddrID: 0x1e,       // D prefix for P2PKH
	ScriptHashAddrID: 0x16,       // 9 or A prefix for P2SH
	Bech32HRPSegwit:  "",         // DOGE doesn't use bech32
}

// ParsedDogecoinTransaction implements the types.DecodedTransaction interface for Dogecoin.
type ParsedDogecoinTransaction struct {
	tx      *wire.MsgTx
	txHash  string
	rawData []byte
	network *chaincfg.Params
}

// ChainIdentifier returns "dogecoin".
func (p *ParsedDogecoinTransaction) ChainIdentifier() string { return "dogecoin" }

// Hash returns the transaction hash.
func (p *ParsedDogecoinTransaction) Hash() string { return p.txHash }

// From returns empty string as Dogecoin doesn't have a simple from address in unsigned transactions
func (p *ParsedDogecoinTransaction) From() string { return "" }

// To returns the first output address if it can be decoded, otherwise empty
func (p *ParsedDogecoinTransaction) To() string {
	if len(p.tx.TxOut) > 0 {
		// Extract address from the first output
		_, addrs, _, err := txscript.ExtractPkScriptAddrs(p.tx.TxOut[0].PkScript, p.network)
		if err == nil && len(addrs) > 0 {
			return addrs[0].EncodeAddress()
		}
	}
	return ""
}

// Value returns the value of the first output in koinu (1 DOGE = 100,000,000 koinu)
func (p *ParsedDogecoinTransaction) Value() *big.Int {
	if len(p.tx.TxOut) > 0 {
		return new(big.Int).SetInt64(p.tx.TxOut[0].Value)
	}
	return big.NewInt(0)
}

// Data returns the raw transaction data
func (p *ParsedDogecoinTransaction) Data() []byte { return p.rawData }

// Nonce returns 0 as Dogecoin doesn't use nonces like Ethereum
func (p *ParsedDogecoinTransaction) Nonce() uint64 { return 0 }

// GasPrice returns 0 as Dogecoin doesn't use gas
func (p *ParsedDogecoinTransaction) GasPrice() *big.Int { return big.NewInt(0) }

// GasLimit returns 0 as Dogecoin doesn't use gas
func (p *ParsedDogecoinTransaction) GasLimit() uint64 { return 0 }

// ParseDogecoinTransaction decodes a raw Dogecoin transaction from its hex representation.
func ParseDogecoinTransaction(txHex string) (types.DecodedTransaction, error) {
	return ParseDogecoinTransactionWithNetwork(txHex, DogecoinMainNetParams)
}

// ParseDogecoinTransactionWithNetwork decodes a raw Dogecoin transaction with a specific network.
func ParseDogecoinTransactionWithNetwork(txHex string, network *chaincfg.Params) (types.DecodedTransaction, error) {
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

	return &ParsedDogecoinTransaction{
		tx:      tx,
		txHash:  txHash,
		rawData: rawTxBytes,
		network: network,
	}, nil
}

// GetTransaction returns the underlying wire.MsgTx
func (p *ParsedDogecoinTransaction) GetTransaction() *wire.MsgTx {
	return p.tx
}

// GetNetwork returns the network parameters
func (p *ParsedDogecoinTransaction) GetNetwork() *chaincfg.Params {
	return p.network
}

// GetOutputCount returns the number of outputs
func (p *ParsedDogecoinTransaction) GetOutputCount() int {
	return len(p.tx.TxOut)
}

// GetInputCount returns the number of inputs
func (p *ParsedDogecoinTransaction) GetInputCount() int {
	return len(p.tx.TxIn)
}

