package bitcoincash

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/wire"
	"github.com/gcash/bchd/chaincfg"
	"github.com/gcash/bchutil"
	"github.com/vultisig/recipes/types"
)

// ParsedBitcoinCashTransaction implements the types.DecodedTransaction interface for Bitcoin Cash.
type ParsedBitcoinCashTransaction struct {
	tx      *wire.MsgTx
	txHash  string
	rawData []byte
}

// ChainIdentifier returns "bitcoincash".
func (p *ParsedBitcoinCashTransaction) ChainIdentifier() string { return "bitcoincash" }

// Hash returns the transaction hash.
func (p *ParsedBitcoinCashTransaction) Hash() string { return p.txHash }

// From returns empty string as Bitcoin Cash doesn't have a simple from address in unsigned transactions
func (p *ParsedBitcoinCashTransaction) From() string { return "" }

// To returns the first output address in CashAddr format if it can be decoded, otherwise empty
func (p *ParsedBitcoinCashTransaction) To() string {
	if len(p.tx.TxOut) > 0 {
		addr, err := ExtractCashAddress(p.tx.TxOut[0].PkScript)
		if err == nil {
			return addr
		}
	}
	return ""
}

// Value returns the value of the first output in satoshis
func (p *ParsedBitcoinCashTransaction) Value() *big.Int {
	if len(p.tx.TxOut) > 0 {
		return new(big.Int).SetInt64(p.tx.TxOut[0].Value)
	}
	return big.NewInt(0)
}

// Data returns the raw transaction data
func (p *ParsedBitcoinCashTransaction) Data() []byte { return p.rawData }

// Nonce returns 0 as Bitcoin Cash doesn't use nonces like Ethereum
func (p *ParsedBitcoinCashTransaction) Nonce() uint64 { return 0 }

// GasPrice returns 0 as Bitcoin Cash doesn't use gas
func (p *ParsedBitcoinCashTransaction) GasPrice() *big.Int { return big.NewInt(0) }

// GasLimit returns 0 as Bitcoin Cash doesn't use gas
func (p *ParsedBitcoinCashTransaction) GasLimit() uint64 { return 0 }

// ParseBitcoinCashTransaction decodes a raw Bitcoin Cash transaction from its hex representation.
func ParseBitcoinCashTransaction(txHex string) (types.DecodedTransaction, error) {
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

	return &ParsedBitcoinCashTransaction{
		tx:      tx,
		txHash:  txHash,
		rawData: rawTxBytes,
	}, nil
}

// ExtractCashAddress extracts a CashAddr format address from a BCH output script.
// Returns addresses in format like "qz839f9pg5af9x9dk9ewwe97kqvtlh7un550ua8lgs" (without prefix).
func ExtractCashAddress(pkScript []byte) (string, error) {
	// Check for P2PKH (OP_DUP OP_HASH160 <20 bytes> OP_EQUALVERIFY OP_CHECKSIG)
	if len(pkScript) == 25 && pkScript[0] == 0x76 && pkScript[1] == 0xa9 &&
		pkScript[2] == 0x14 && pkScript[23] == 0x88 && pkScript[24] == 0xac {
		hash := pkScript[3:23]
		addr, err := bchutil.NewAddressPubKeyHash(hash, &chaincfg.MainNetParams)
		if err != nil {
			return "", fmt.Errorf("failed to create CashAddr P2PKH: %w", err)
		}
		// Return without "bitcoincash:" prefix to match THORChain format
		return addr.EncodeAddress(), nil
	}

	// Check for P2SH (OP_HASH160 <20 bytes> OP_EQUAL)
	if len(pkScript) == 23 && pkScript[0] == 0xa9 && pkScript[1] == 0x14 && pkScript[22] == 0x87 {
		hash := pkScript[2:22]
		addr, err := bchutil.NewAddressScriptHashFromHash(hash, &chaincfg.MainNetParams)
		if err != nil {
			return "", fmt.Errorf("failed to create CashAddr P2SH: %w", err)
		}
		return addr.EncodeAddress(), nil
	}

	// Check for OP_RETURN
	if len(pkScript) > 0 && pkScript[0] == 0x6a {
		return "", fmt.Errorf("OP_RETURN output has no address")
	}

	return "", fmt.Errorf("unsupported script type")
}

// GetTransaction returns the underlying wire.MsgTx
func (p *ParsedBitcoinCashTransaction) GetTransaction() *wire.MsgTx {
	return p.tx
}

// GetOutputCount returns the number of outputs
func (p *ParsedBitcoinCashTransaction) GetOutputCount() int {
	return len(p.tx.TxOut)
}

// GetInputCount returns the number of inputs
func (p *ParsedBitcoinCashTransaction) GetInputCount() int {
	return len(p.tx.TxIn)
}

