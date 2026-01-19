// Package utxo provides common types and interfaces for UTXO-based blockchains.
package utxo

import (
	"fmt"
	"math/big"

	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/types"
)

// Chain represents a UTXO-based blockchain.
type Chain interface {
	types.Chain

	// ComputeTxHash computes the transaction hash after applying signatures.
	ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error)
}

// Transaction represents a parsed UTXO transaction with output access.
type Transaction interface {
	// OutputCount returns the number of outputs in the transaction.
	OutputCount() int

	// Output returns the output at the given index.
	Output(index int) (Output, error)
}

// Output represents a single transaction output.
type Output interface {
	// Value returns the output value in the smallest unit (satoshis, zatoshis, etc.).
	Value() int64

	// PkScript returns the output's public key script.
	PkScript() []byte
}

// AddressExtractor extracts an address from a PkScript for a specific chain.
type AddressExtractor interface {
	// ExtractAddress extracts the address from the given PkScript.
	ExtractAddress(pkScript []byte) (string, error)
}

// EncodeDERSignature encodes r, s values into DER format as required by BIP66.
// The format is: 0x30 <total_length> 0x02 <r_length> <R> 0x02 <s_length> <S>
func EncodeDERSignature(r, s []byte) []byte {
	// Normalize r and s (remove leading zeros and ensure positive)
	rNorm := normalizeForDER(r)
	sNorm := normalizeForDER(s)

	// Calculate lengths
	rLen := len(rNorm)
	sLen := len(sNorm)
	totalLen := 2 + rLen + 2 + sLen // 0x02 + rLen byte + r + 0x02 + sLen byte + s

	// Build DER signature
	der := make([]byte, 0, 2+totalLen)
	der = append(der, 0x30)           // SEQUENCE tag
	der = append(der, byte(totalLen)) // Total length
	der = append(der, 0x02)           // INTEGER tag for r
	der = append(der, byte(rLen))     // r length
	der = append(der, rNorm...)       // r value
	der = append(der, 0x02)           // INTEGER tag for s
	der = append(der, byte(sLen))     // s length
	der = append(der, sNorm...)       // s value

	return der
}

// normalizeForDER prepares a big-endian integer for DER encoding.
// DER integers must not have leading zeros (except when needed for sign bit)
// and must be positive (prepend 0x00 if high bit is set).
func normalizeForDER(b []byte) []byte {
	// Use big.Int to normalize (removes leading zeros)
	i := new(big.Int).SetBytes(b)
	normalized := i.Bytes()

	// If the integer is zero, return a single zero byte
	if len(normalized) == 0 {
		return []byte{0x00}
	}

	// If high bit is set, prepend 0x00 to indicate positive number
	if normalized[0]&0x80 != 0 {
		return append([]byte{0x00}, normalized...)
	}

	return normalized
}

// ExtractPubKeyFromScriptSig extracts the public key from a P2PKH scriptSig.
// The scriptSig format is: <sig_length> <signature> <pubkey_length> <pubkey>
// This is used to extract a pre-populated public key set by the tx proposer.
func ExtractPubKeyFromScriptSig(scriptSig []byte) ([]byte, error) {
	if len(scriptSig) == 0 {
		return nil, fmt.Errorf("empty scriptSig")
	}

	// First byte is the signature push length
	if len(scriptSig) < 1 {
		return nil, fmt.Errorf("scriptSig too short")
	}

	sigPushLen := int(scriptSig[0])
	if sigPushLen < 1 || sigPushLen > 73 {
		return nil, fmt.Errorf("invalid signature push length: %d", sigPushLen)
	}

	// Skip past the signature push opcode and signature data
	offset := 1 + sigPushLen

	if len(scriptSig) < offset+1 {
		return nil, fmt.Errorf("scriptSig too short for pubkey")
	}

	// Next byte is the pubkey push length
	pubKeyPushLen := int(scriptSig[offset])

	// Compressed pubkey is 33 bytes, uncompressed is 65 bytes
	if pubKeyPushLen != 33 && pubKeyPushLen != 65 {
		return nil, fmt.Errorf("invalid pubkey length: %d", pubKeyPushLen)
	}

	offset++
	if len(scriptSig) < offset+pubKeyPushLen {
		return nil, fmt.Errorf("scriptSig too short for pubkey data")
	}

	pubKey := make([]byte, pubKeyPushLen)
	copy(pubKey, scriptSig[offset:offset+pubKeyPushLen])

	return pubKey, nil
}
