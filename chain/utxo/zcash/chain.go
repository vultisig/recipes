package zcash

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/vultisig/mobile-tss-lib/tss"

	"github.com/vultisig/recipes/chain/utxo"
	"github.com/vultisig/recipes/sdk/zcash"
	"github.com/vultisig/recipes/types"
)

// Zcash implements the Chain interface for the Zcash blockchain
type Zcash struct{}

// ID returns the unique identifier for the Zcash chain
func (z *Zcash) ID() string {
	return "zcash"
}

// Name returns a human-readable name for the Zcash chain
func (z *Zcash) Name() string {
	return "Zcash"
}

// SupportedProtocols returns the list of protocol IDs supported by the Zcash chain
func (z *Zcash) SupportedProtocols() []string {
	return []string{"zec"}
}

// Description returns a human-readable description for the Zcash chain
func (z *Zcash) Description() string {
	return "Zcash is a privacy-focused cryptocurrency with transparent and shielded transaction options."
}

func (z *Zcash) GetProtocol(id string) (types.Protocol, error) {
	if id == "zec" {
		return NewZEC(), nil
	}
	return nil, fmt.Errorf("protocol %q not found or not supported on Zcash", id)
}

func (z *Zcash) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	return ParseZcashTransaction(txHex)
}

func (z *Zcash) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	// Parse the transaction
	tx, err := deserializeZcashTransaction(proposedTx)
	if err != nil {
		return "", fmt.Errorf("deserializeZcashTransaction: %w", err)
	}

	if len(tx.Inputs) != len(sigs) {
		return "", fmt.Errorf("input count (%d) does not match sigs count (%d)", len(tx.Inputs), len(sigs))
	}

	// For each input, apply the signature
	for i, input := range tx.Inputs {
		selectedSig := sigs[i]
		r, er := hex.DecodeString(selectedSig.R)
		if er != nil {
			return "", fmt.Errorf("hex.DecodeString(selectedSig.R): %w", er)
		}
		s, er := hex.DecodeString(selectedSig.S)
		if er != nil {
			return "", fmt.Errorf("hex.DecodeString(selectedSig.S): %w", er)
		}

		// Encode signature in DER format as required by BIP66
		derSig := utxo.EncodeDERSignature(r, s)
		derSig = append(derSig, byte(txscript.SigHashAll))

		// For transparent Zcash transactions, extract the public key from the pre-populated scriptSig
		pubKey, er := utxo.ExtractPubKeyFromScriptSig(input.SignatureScript)
		if er != nil {
			return "", fmt.Errorf("failed to extract pubkey from scriptSig for input %d: %w", i, er)
		}

		// Build scriptSig with DER signature and public key
		scriptSig, er := txscript.NewScriptBuilder().AddData(derSig).AddData(pubKey).Script()
		if er != nil {
			return "", fmt.Errorf("txscript.NewScriptBuilder: %w", er)
		}
		input.SignatureScript = scriptSig
	}

	// Re-serialize the transaction with signatures
	signedTx, err := serializeZcashTransaction(tx)
	if err != nil {
		return "", fmt.Errorf("serializeZcashTransaction: %w", err)
	}

	// Calculate hash
	hash := chainhash.DoubleHashH(signedTx)
	return hash.String(), nil
}

// serializeZcashTransaction serializes a ZcashTransaction back to bytes
func serializeZcashTransaction(tx *ZcashTransaction) ([]byte, error) {
	var buf bytes.Buffer

	// Write header with overwintered flag
	header := uint32(tx.Version) | 0x80000000 // Set overwintered flag
	if err := writeUint32(&buf, header); err != nil {
		return nil, err
	}

	// Write version group ID
	if err := writeUint32(&buf, tx.VersionGroupID); err != nil {
		return nil, err
	}

	// Write inputs
	if err := writeVarInt(&buf, uint64(len(tx.Inputs))); err != nil {
		return nil, err
	}
	for _, input := range tx.Inputs {
		if err := writeZcashInput(&buf, input); err != nil {
			return nil, err
		}
	}

	// Write outputs
	if err := writeVarInt(&buf, uint64(len(tx.Outputs))); err != nil {
		return nil, err
	}
	for _, output := range tx.Outputs {
		if err := writeZcashOutput(&buf, output); err != nil {
			return nil, err
		}
	}

	// Write lock time
	if err := writeUint32(&buf, tx.LockTime); err != nil {
		return nil, err
	}

	// Write expiry height
	if err := writeUint32(&buf, tx.ExpiryHeight); err != nil {
		return nil, err
	}

	// For v4+, write value balance and empty shielded sections
	if tx.Version >= 4 {
		if err := writeInt64(&buf, tx.ValueBalance); err != nil {
			return nil, err
		}
		// Empty shielded spends
		if err := writeVarInt(&buf, 0); err != nil {
			return nil, err
		}
		// Empty shielded outputs
		if err := writeVarInt(&buf, 0); err != nil {
			return nil, err
		}
		// Empty joinsplits for Sapling
		if tx.VersionGroupID == 0x892F2085 {
			if err := writeVarInt(&buf, 0); err != nil {
				return nil, err
			}
		}
	}

	return buf.Bytes(), nil
}

func writeUint32(buf *bytes.Buffer, v uint32) error {
	b := make([]byte, 4)
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	_, err := buf.Write(b)
	return err
}

func writeInt64(buf *bytes.Buffer, v int64) error {
	b := make([]byte, 8)
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
	_, err := buf.Write(b)
	return err
}

func writeVarInt(buf *bytes.Buffer, v uint64) error {
	switch {
	case v < 0xFD:
		return buf.WriteByte(byte(v))
	case v <= 0xFFFF:
		if err := buf.WriteByte(0xFD); err != nil {
			return err
		}
		b := make([]byte, 2)
		b[0] = byte(v)
		b[1] = byte(v >> 8)
		_, err := buf.Write(b)
		return err
	case v <= 0xFFFFFFFF:
		if err := buf.WriteByte(0xFE); err != nil {
			return err
		}
		return writeUint32(buf, uint32(v))
	default:
		if err := buf.WriteByte(0xFF); err != nil {
			return err
		}
		return writeInt64(buf, int64(v))
	}
}

func writeZcashInput(buf *bytes.Buffer, input *ZcashInput) error {
	// Write previous outpoint hash (reversed)
	hash := input.PreviousOutPoint.Hash
	reversed := make([]byte, 32)
	for i := 0; i < 32; i++ {
		reversed[i] = hash[31-i]
	}
	if _, err := buf.Write(reversed); err != nil {
		return err
	}

	// Write previous outpoint index
	if err := writeUint32(buf, input.PreviousOutPoint.Index); err != nil {
		return err
	}

	// Write signature script
	if err := writeVarInt(buf, uint64(len(input.SignatureScript))); err != nil {
		return err
	}
	if _, err := buf.Write(input.SignatureScript); err != nil {
		return err
	}

	// Write sequence
	return writeUint32(buf, input.Sequence)
}

func writeZcashOutput(buf *bytes.Buffer, output *ZcashOutput) error {
	// Write value
	if err := writeInt64(buf, output.Value); err != nil {
		return err
	}

	// Write pk script
	if err := writeVarInt(buf, uint64(len(output.PkScript))); err != nil {
		return err
	}
	_, err := buf.Write(output.PkScript)
	return err
}

func (z *Zcash) ExtractTxBytes(txData string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(txData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %w", err)
	}
	// Strip embedded metadata (sighashes, pubkey)
	txBytes, _, _, err := zcash.ParseWithMetadata(b)
	if err != nil {
		return nil, fmt.Errorf("failed to parse metadata: %w", err)
	}
	return txBytes, nil
}

// NewChain creates a new Zcash chain instance
func NewChain() types.Chain {
	return &Zcash{}
}

