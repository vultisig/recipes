// Package zcash provides SDK functionality for Zcash transaction signing and broadcasting.
package zcash

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/vultisig/mobile-tss-lib/tss"
	"golang.org/x/crypto/blake2b"
)

// TxBroadcaster is the interface for broadcasting transactions to the Zcash network.
type TxBroadcaster interface {
	BroadcastTransaction(signedTx []byte) (string, error)
}

// SDK represents the Zcash SDK for transaction signing and broadcasting.
type SDK struct {
	broadcaster TxBroadcaster
}

// NewSDK creates a new Zcash SDK instance.
func NewSDK(broadcaster TxBroadcaster) *SDK {
	return &SDK{
		broadcaster: broadcaster,
	}
}

// TxInput represents a transaction input for signing.
type TxInput struct {
	TxHash   string // Transaction hash (hex string)
	Index    uint32 // Output index in the previous transaction
	Value    uint64 // Value in zatoshis
	Script   []byte // Script code for signing (P2PKH script)
	Sequence uint32 // Sequence number (typically 0xffffffff)
}

// TxOutput represents a transaction output.
type TxOutput struct {
	Value  int64  // Value in zatoshis
	Script []byte // Output script (P2PKH, P2SH, or OP_RETURN)
}

// UnsignedTx represents an unsigned Zcash transaction with all necessary info for signing.
type UnsignedTx struct {
	Inputs   []TxInput
	Outputs  []*TxOutput
	PubKey   []byte   // 33-byte compressed public key
	RawBytes []byte   // Serialized unsigned transaction
	SigHashes [][]byte // Pre-computed signature hashes for each input
}

// ConsensusBranchID is the NU6 consensus branch ID for signature hash personalization.
// Although we use v4 transactions (Sapling format), we must use the
// consensus branch ID of the current epoch for signature hashing.
// NU6 activated on November 23, 2024.
const ConsensusBranchID = 0xC8E71055

// Zcash v4 transaction constants
const (
	zcashV4Version        = uint32(0x80000004) // Version 4 with overwintered flag
	zcashSaplingVersionID = uint32(0x892F2085) // Sapling version group ID
)

// Sign applies TSS signatures to an unsigned transaction and returns the signed transaction bytes.
// signatures: Map where key is derived from message hash (SHA256 + Base64), value is the signature.
func (sdk *SDK) Sign(unsignedTx *UnsignedTx, signatures map[string]tss.KeysignResponse) ([]byte, error) {
	var buf bytes.Buffer

	// Zcash v4 transaction format (Sapling)
	// Version (4 bytes, little-endian) - version 4 with overwintered flag
	if err := binary.Write(&buf, binary.LittleEndian, zcashV4Version); err != nil {
		return nil, fmt.Errorf("zcash: failed to write version: %w", err)
	}

	// Version group ID (4 bytes, little-endian) - Sapling
	if err := binary.Write(&buf, binary.LittleEndian, zcashSaplingVersionID); err != nil {
		return nil, fmt.Errorf("zcash: failed to write version group ID: %w", err)
	}

	// Transparent inputs count
	writeCompactSize(&buf, uint64(len(unsignedTx.Inputs)))

	// Transparent inputs with signatures
	for i, input := range unsignedTx.Inputs {
		// Previous output hash (32 bytes, reversed)
		txHashBytes, err := hex.DecodeString(input.TxHash)
		if err != nil {
			return nil, fmt.Errorf("zcash: invalid tx hash for input %d: %w", i, err)
		}
		// Reverse for little-endian
		for j := len(txHashBytes) - 1; j >= 0; j-- {
			buf.WriteByte(txHashBytes[j])
		}

		// Previous output index (4 bytes, little-endian)
		if err := binary.Write(&buf, binary.LittleEndian, input.Index); err != nil {
			return nil, fmt.Errorf("zcash: failed to write input index: %w", err)
		}

		// Get signature for this input
		if i >= len(unsignedTx.SigHashes) {
			return nil, fmt.Errorf("zcash: missing sig hash for input %d", i)
		}
		sigHash := unsignedTx.SigHashes[i]
		derivedKey := sdk.DeriveKeyFromMessage(sigHash)

		sig, exists := signatures[derivedKey]
		if !exists {
			return nil, fmt.Errorf("zcash: missing signature for input %d (key: %s)", i, derivedKey)
		}

		// Build scriptSig for P2PKH: <sig_length> <sig> <pubkey_length> <pubkey>
		derSig, err := hex.DecodeString(trim0x(sig.DerSignature))
		if err != nil {
			return nil, fmt.Errorf("zcash: failed to decode DER signature for input %d: %w", i, err)
		}

		// Append SIGHASH_ALL
		fullSig := append(derSig, byte(txscript.SigHashAll))

		// scriptSig: <sig_length> <sig> <pubkey_length> <pubkey>
		scriptSig := make([]byte, 0, 2+len(fullSig)+len(unsignedTx.PubKey))
		scriptSig = append(scriptSig, byte(len(fullSig)))
		scriptSig = append(scriptSig, fullSig...)
		scriptSig = append(scriptSig, byte(len(unsignedTx.PubKey)))
		scriptSig = append(scriptSig, unsignedTx.PubKey...)

		// Script length
		writeCompactSize(&buf, uint64(len(scriptSig)))
		buf.Write(scriptSig)

		// Sequence (4 bytes)
		if err := binary.Write(&buf, binary.LittleEndian, uint32(0xffffffff)); err != nil {
			return nil, fmt.Errorf("zcash: failed to write sequence: %w", err)
		}
	}

	// Transparent outputs count
	writeCompactSize(&buf, uint64(len(unsignedTx.Outputs)))

	// Transparent outputs
	for i, output := range unsignedTx.Outputs {
		if err := binary.Write(&buf, binary.LittleEndian, uint64(output.Value)); err != nil {
			return nil, fmt.Errorf("zcash: failed to write output value %d: %w", i, err)
		}
		writeCompactSize(&buf, uint64(len(output.Script)))
		buf.Write(output.Script)
	}

	// Lock time (4 bytes, little-endian)
	if err := binary.Write(&buf, binary.LittleEndian, uint32(0)); err != nil {
		return nil, fmt.Errorf("zcash: failed to write lock time: %w", err)
	}

	// Expiry height (4 bytes, little-endian)
	if err := binary.Write(&buf, binary.LittleEndian, uint32(0)); err != nil {
		return nil, fmt.Errorf("zcash: failed to write expiry height: %w", err)
	}

	// Value balance (8 bytes, little-endian) - 0 for transparent-only
	if err := binary.Write(&buf, binary.LittleEndian, int64(0)); err != nil {
		return nil, fmt.Errorf("zcash: failed to write value balance: %w", err)
	}

	// Shielded spends count - 0
	buf.WriteByte(0x00)

	// Shielded outputs count - 0
	buf.WriteByte(0x00)

	// JoinSplits count - 0 (for Sapling v4)
	buf.WriteByte(0x00)

	return buf.Bytes(), nil
}

// Broadcast submits a signed transaction to the Zcash network.
func (sdk *SDK) Broadcast(signedTx []byte) (string, error) {
	if sdk.broadcaster == nil {
		return "", fmt.Errorf("zcash: broadcaster not configured")
	}

	return sdk.broadcaster.BroadcastTransaction(signedTx)
}

// Send is a convenience method that signs and broadcasts the transaction.
func (sdk *SDK) Send(unsignedTx *UnsignedTx, signatures map[string]tss.KeysignResponse) (string, error) {
	signedTx, err := sdk.Sign(unsignedTx, signatures)
	if err != nil {
		return "", fmt.Errorf("zcash: failed to sign transaction: %w", err)
	}

	txHash, err := sdk.Broadcast(signedTx)
	if err != nil {
		return "", fmt.Errorf("zcash: failed to broadcast transaction: %w", err)
	}

	return txHash, nil
}

// CalculateSigHash computes the signature hash for a Zcash transparent input.
// This uses the ZIP-243 signature hash algorithm for v4 (Sapling) transactions.
func (sdk *SDK) CalculateSigHash(inputs []TxInput, outputs []*TxOutput, inputIndex int) ([]byte, error) {
	if inputIndex < 0 || inputIndex >= len(inputs) {
		return nil, fmt.Errorf("zcash: input index %d out of range (0-%d)", inputIndex, len(inputs)-1)
	}

	var preimage bytes.Buffer

	// 1. nVersion | nVersionGroupId (header)
	if err := binary.Write(&preimage, binary.LittleEndian, zcashV4Version); err != nil {
		return nil, fmt.Errorf("zcash: failed to write version: %w", err)
	}
	if err := binary.Write(&preimage, binary.LittleEndian, zcashSaplingVersionID); err != nil {
		return nil, fmt.Errorf("zcash: failed to write version group ID: %w", err)
	}

	// 2. hashPrevouts - BLAKE2b-256 of all prevouts
	hashPrevouts, err := sdk.calcHashPrevouts(inputs)
	if err != nil {
		return nil, fmt.Errorf("zcash: failed to calculate hashPrevouts: %w", err)
	}
	preimage.Write(hashPrevouts)

	// 3. hashSequence - BLAKE2b-256 of all sequences
	hashSequence := sdk.calcHashSequence(inputs)
	preimage.Write(hashSequence)

	// 4. hashOutputs - BLAKE2b-256 of all outputs
	hashOutputs := sdk.calcHashOutputs(outputs)
	preimage.Write(hashOutputs)

	// 5. hashJoinSplits - 32 zero bytes (no joinsplits)
	preimage.Write(make([]byte, 32))

	// 6. hashShieldedSpends - 32 zero bytes (no shielded spends)
	preimage.Write(make([]byte, 32))

	// 7. hashShieldedOutputs - 32 zero bytes (no shielded outputs)
	preimage.Write(make([]byte, 32))

	// 8. nLockTime
	if err := binary.Write(&preimage, binary.LittleEndian, uint32(0)); err != nil {
		return nil, fmt.Errorf("zcash: failed to write lock time: %w", err)
	}

	// 9. nExpiryHeight
	if err := binary.Write(&preimage, binary.LittleEndian, uint32(0)); err != nil {
		return nil, fmt.Errorf("zcash: failed to write expiry height: %w", err)
	}

	// 10. valueBalance (8 bytes) - 0 for transparent-only
	if err := binary.Write(&preimage, binary.LittleEndian, int64(0)); err != nil {
		return nil, fmt.Errorf("zcash: failed to write value balance: %w", err)
	}

	// 11. nHashType
	if err := binary.Write(&preimage, binary.LittleEndian, uint32(txscript.SigHashAll)); err != nil {
		return nil, fmt.Errorf("zcash: failed to write hash type: %w", err)
	}

	// For SIGHASH_ALL, include the input being signed
	input := inputs[inputIndex]

	// prevout (txid + index)
	txHashBytes, err := hex.DecodeString(input.TxHash)
	if err != nil {
		return nil, fmt.Errorf("zcash: invalid tx hash hex for input %d: %w", inputIndex, err)
	}
	// Reverse for little-endian
	for j := len(txHashBytes) - 1; j >= 0; j-- {
		preimage.WriteByte(txHashBytes[j])
	}
	if err := binary.Write(&preimage, binary.LittleEndian, input.Index); err != nil {
		return nil, fmt.Errorf("zcash: failed to write input index: %w", err)
	}

	// scriptCode (with length prefix)
	writeCompactSize(&preimage, uint64(len(input.Script)))
	preimage.Write(input.Script)

	// amount (value of the input)
	if err := binary.Write(&preimage, binary.LittleEndian, input.Value); err != nil {
		return nil, fmt.Errorf("zcash: failed to write input value: %w", err)
	}

	// nSequence
	if err := binary.Write(&preimage, binary.LittleEndian, uint32(0xffffffff)); err != nil {
		return nil, fmt.Errorf("zcash: failed to write sequence: %w", err)
	}

	// Final hash using BLAKE2b-256 with personalization
	return sdk.blake2bSigHash(preimage.Bytes())
}

// SerializeUnsignedTx creates raw unsigned transaction bytes.
// Uses Zcash v4 (Sapling) format for compatibility with recipes engine.
func (sdk *SDK) SerializeUnsignedTx(inputs []TxInput, outputs []*TxOutput) ([]byte, error) {
	var buf bytes.Buffer

	// Version (4 bytes, little-endian) - version 4 with overwintered flag
	if err := binary.Write(&buf, binary.LittleEndian, zcashV4Version); err != nil {
		return nil, fmt.Errorf("zcash: failed to write version: %w", err)
	}

	// Version group ID (4 bytes, little-endian) - Sapling
	if err := binary.Write(&buf, binary.LittleEndian, zcashSaplingVersionID); err != nil {
		return nil, fmt.Errorf("zcash: failed to write version group ID: %w", err)
	}

	// Transparent inputs count (compactSize)
	writeCompactSize(&buf, uint64(len(inputs)))

	// Transparent inputs
	for i, input := range inputs {
		// Previous output hash (32 bytes)
		txHashBytes, err := hex.DecodeString(input.TxHash)
		if err != nil {
			return nil, fmt.Errorf("zcash: invalid tx hash for input %d: %w", i, err)
		}
		// Reverse for little-endian
		for j := len(txHashBytes) - 1; j >= 0; j-- {
			buf.WriteByte(txHashBytes[j])
		}

		// Previous output index (4 bytes, little-endian)
		if err := binary.Write(&buf, binary.LittleEndian, input.Index); err != nil {
			return nil, fmt.Errorf("zcash: failed to write input index: %w", err)
		}

		// Script length (compactSize) - empty for unsigned
		buf.WriteByte(0x00)

		// Sequence (4 bytes, little-endian) - 0xffffffff
		if err := binary.Write(&buf, binary.LittleEndian, uint32(0xffffffff)); err != nil {
			return nil, fmt.Errorf("zcash: failed to write sequence: %w", err)
		}
	}

	// Transparent outputs count (compactSize)
	writeCompactSize(&buf, uint64(len(outputs)))

	// Transparent outputs
	for i, output := range outputs {
		if output.Value < 0 {
			return nil, fmt.Errorf("zcash: invalid negative output value at index %d: %d", i, output.Value)
		}
		// Value (8 bytes, little-endian)
		if err := binary.Write(&buf, binary.LittleEndian, uint64(output.Value)); err != nil {
			return nil, fmt.Errorf("zcash: failed to write output value: %w", err)
		}

		// Script length (compactSize)
		writeCompactSize(&buf, uint64(len(output.Script)))

		// Script
		buf.Write(output.Script)
	}

	// Lock time (4 bytes, little-endian)
	if err := binary.Write(&buf, binary.LittleEndian, uint32(0)); err != nil {
		return nil, fmt.Errorf("zcash: failed to write lock time: %w", err)
	}

	// Expiry height (4 bytes, little-endian) - 0 for no expiry
	if err := binary.Write(&buf, binary.LittleEndian, uint32(0)); err != nil {
		return nil, fmt.Errorf("zcash: failed to write expiry height: %w", err)
	}

	// Value balance (8 bytes, little-endian) - 0 for transparent-only
	if err := binary.Write(&buf, binary.LittleEndian, int64(0)); err != nil {
		return nil, fmt.Errorf("zcash: failed to write value balance: %w", err)
	}

	// Shielded spends count (compactSize) - 0 for transparent-only
	buf.WriteByte(0x00)

	// Shielded outputs count (compactSize) - 0 for transparent-only
	buf.WriteByte(0x00)

	// JoinSplits count (compactSize) - 0 for transparent-only (Sapling v4)
	buf.WriteByte(0x00)

	return buf.Bytes(), nil
}

// ComputeTxHash computes the transaction hash from signed transaction bytes.
// For Zcash v4 transactions, hash is double SHA256 of the serialized tx.
func (sdk *SDK) ComputeTxHash(signedTx []byte) string {
	hash := chainhash.DoubleHashH(signedTx)
	return hash.String()
}

// DeriveKeyFromMessage derives a map key from a message hash using SHA256 + Base64.
// This is used to look up signatures in the TSS response map.
func (sdk *SDK) DeriveKeyFromMessage(messageHash []byte) string {
	return DeriveKeyFromMessage(messageHash)
}

// DeriveKeyFromMessage derives a map key from a message hash using SHA256 + Base64.
// This standalone function can be used without an SDK instance.
func DeriveKeyFromMessage(messageHash []byte) string {
	hash := sha256.Sum256(messageHash)
	return base64.StdEncoding.EncodeToString(hash[:])
}

// blake2bSigHash computes BLAKE2b-256 with Zcash signature hash personalization.
func (sdk *SDK) blake2bSigHash(data []byte) ([]byte, error) {
	// Personalization: "ZcashSigHash" (12 bytes) + branch ID (4 bytes, little-endian)
	personalization := make([]byte, 16)
	copy(personalization, "ZcashSigHash")
	binary.LittleEndian.PutUint32(personalization[12:], ConsensusBranchID)

	h, err := blake2b.New256(personalization)
	if err != nil {
		return nil, fmt.Errorf("failed to create BLAKE2b hasher: %w", err)
	}
	h.Write(data)
	return h.Sum(nil), nil
}

// calcHashPrevouts computes BLAKE2b-256 of all input prevouts.
func (sdk *SDK) calcHashPrevouts(inputs []TxInput) ([]byte, error) {
	var buf bytes.Buffer
	for i, input := range inputs {
		txHashBytes, err := hex.DecodeString(input.TxHash)
		if err != nil {
			return nil, fmt.Errorf("invalid tx hash hex for input %d: %w", i, err)
		}
		// Reverse for little-endian
		for j := len(txHashBytes) - 1; j >= 0; j-- {
			buf.WriteByte(txHashBytes[j])
		}
		if err := binary.Write(&buf, binary.LittleEndian, input.Index); err != nil {
			return nil, fmt.Errorf("failed to write input index: %w", err)
		}
	}

	personalization := make([]byte, 16)
	copy(personalization, "ZcashPrevoutHash")
	h, _ := blake2b.New256(personalization)
	h.Write(buf.Bytes())
	return h.Sum(nil), nil
}

// calcHashSequence computes BLAKE2b-256 of all input sequences.
func (sdk *SDK) calcHashSequence(inputs []TxInput) []byte {
	var buf bytes.Buffer
	for range inputs {
		_ = binary.Write(&buf, binary.LittleEndian, uint32(0xffffffff))
	}

	personalization := make([]byte, 16)
	copy(personalization, "ZcashSequencHash")
	h, _ := blake2b.New256(personalization)
	h.Write(buf.Bytes())
	return h.Sum(nil)
}

// calcHashOutputs computes BLAKE2b-256 of all outputs.
func (sdk *SDK) calcHashOutputs(outputs []*TxOutput) []byte {
	var buf bytes.Buffer
	for _, output := range outputs {
		_ = binary.Write(&buf, binary.LittleEndian, uint64(output.Value))
		writeCompactSize(&buf, uint64(len(output.Script)))
		buf.Write(output.Script)
	}

	personalization := make([]byte, 16)
	copy(personalization, "ZcashOutputsHash")
	h, _ := blake2b.New256(personalization)
	h.Write(buf.Bytes())
	return h.Sum(nil)
}

// writeCompactSize writes a variable-length integer (Bitcoin/Zcash compact size).
func writeCompactSize(buf *bytes.Buffer, n uint64) {
	switch {
	case n < 0xFD:
		buf.WriteByte(byte(n))
	case n <= 0xFFFF:
		buf.WriteByte(0xFD)
		_ = binary.Write(buf, binary.LittleEndian, uint16(n))
	case n <= 0xFFFFFFFF:
		buf.WriteByte(0xFE)
		_ = binary.Write(buf, binary.LittleEndian, uint32(n))
	default:
		buf.WriteByte(0xFF)
		_ = binary.Write(buf, binary.LittleEndian, n)
	}
}

// trim0x removes the "0x" prefix from a hex string if present.
func trim0x(s string) string {
	if len(s) >= 2 && (s[0:2] == "0x" || s[0:2] == "0X") {
		return s[2:]
	}
	return s
}

