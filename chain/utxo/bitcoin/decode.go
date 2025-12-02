package bitcoin

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

// ParsedBitcoinTransaction implements the types.DecodedTransaction interface for Bitcoin.
type ParsedBitcoinTransaction struct {
	tx      *wire.MsgTx
	txHash  string
	rawData []byte
	network *chaincfg.Params
}

// ChainIdentifier returns "bitcoin".
func (p *ParsedBitcoinTransaction) ChainIdentifier() string { return "bitcoin" }

// Hash returns the transaction hash.
func (p *ParsedBitcoinTransaction) Hash() string { return p.txHash }

// From returns empty string as Bitcoin doesn't have a simple from address in unsigned transactions
func (p *ParsedBitcoinTransaction) From() string { return "" }

// To returns the first output address if it can be decoded, otherwise empty
func (p *ParsedBitcoinTransaction) To() string {
	if len(p.tx.TxOut) > 0 {
		// Extract address from the first output
		_, addrs, _, err := txscript.ExtractPkScriptAddrs(p.tx.TxOut[0].PkScript, p.network)
		if err == nil && len(addrs) > 0 {
			return addrs[0].EncodeAddress()
		}
	}
	return ""
}

// Value returns the value of the first output in satoshis
func (p *ParsedBitcoinTransaction) Value() *big.Int {
	if len(p.tx.TxOut) > 0 {
		return new(big.Int).SetInt64(p.tx.TxOut[0].Value)
	}
	return big.NewInt(0)
}

// Data returns the raw transaction data
func (p *ParsedBitcoinTransaction) Data() []byte { return p.rawData }

// Nonce returns 0 as Bitcoin doesn't use nonces like Ethereum
func (p *ParsedBitcoinTransaction) Nonce() uint64 { return 0 }

// GasPrice returns 0 as Bitcoin doesn't use gas
func (p *ParsedBitcoinTransaction) GasPrice() *big.Int { return big.NewInt(0) }

// GasLimit returns 0 as Bitcoin doesn't use gas
func (p *ParsedBitcoinTransaction) GasLimit() uint64 { return 0 }

// ParseBitcoinTransaction decodes a raw Bitcoin transaction from its hex representation.
func ParseBitcoinTransaction(txHex string) (types.DecodedTransaction, error) {
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

	// Use mainnet params by default
	network := &chaincfg.MainNetParams

	// Calculate transaction hash
	txHash := tx.TxHash().String()

	return &ParsedBitcoinTransaction{
		tx:      tx,
		txHash:  txHash,
		rawData: rawTxBytes,
		network: network,
	}, nil
}

// ParseBitcoinTransactionWithNetwork decodes a raw Bitcoin transaction with a specific network.
func ParseBitcoinTransactionWithNetwork(txHex string, network *chaincfg.Params) (types.DecodedTransaction, error) {
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

	return &ParsedBitcoinTransaction{
		tx:      tx,
		txHash:  txHash,
		rawData: rawTxBytes,
		network: network,
	}, nil
}

// GetAllOutputs returns all outputs with their addresses and values
func (p *ParsedBitcoinTransaction) GetAllOutputs() []struct {
	Address string
	Value   int64
} {
	var outputs []struct {
		Address string
		Value   int64
	}

	for _, txOut := range p.tx.TxOut {
		output := struct {
			Address string
			Value   int64
		}{
			Value: txOut.Value,
		}

		// Try to extract address
		_, addrs, _, err := txscript.ExtractPkScriptAddrs(txOut.PkScript, p.network)
		if err == nil && len(addrs) > 0 {
			output.Address = addrs[0].EncodeAddress()
		}

		outputs = append(outputs, output)
	}

	return outputs
}

// GetInputCount returns the number of inputs
func (p *ParsedBitcoinTransaction) GetInputCount() int {
	return len(p.tx.TxIn)
}

// GetOutputCount returns the number of outputs
func (p *ParsedBitcoinTransaction) GetOutputCount() int {
	return len(p.tx.TxOut)
}

// GetLockTime returns the transaction lock time
func (p *ParsedBitcoinTransaction) GetLockTime() uint32 {
	return p.tx.LockTime
}

// GetVersion returns the transaction version
func (p *ParsedBitcoinTransaction) GetVersion() int32 {
	return p.tx.Version
}

// IsSegWit returns true if the transaction uses SegWit
func (p *ParsedBitcoinTransaction) IsSegWit() bool {
	return p.tx.HasWitness()
}

// GetTransaction returns the underlying wire.MsgTx
func (p *ParsedBitcoinTransaction) GetTransaction() *wire.MsgTx {
	return p.tx
}

// GetNetwork returns the network parameters
func (p *ParsedBitcoinTransaction) GetNetwork() *chaincfg.Params {
	return p.network
}

