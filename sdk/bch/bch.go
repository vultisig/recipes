// Package bch provides Bitcoin Cash specific transaction signing functionality.
// BCH requires SIGHASH_FORKID and BIP143-style signature hashing even for legacy P2PKH.
package bch

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	bchtxscript "github.com/gcash/bchd/txscript"
	bchwire "github.com/gcash/bchd/wire"
	"github.com/vultisig/mobile-tss-lib/tss"
)

const (
	// SigHashForkID is the BCH-specific sighash flag that must be OR'd with sighash type.
	// This provides replay protection between BCH and BTC.
	SigHashForkID = 0x40
)

// rpcClient interface - methods we use from the client
type rpcClient interface {
	SendRawTransaction(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error)
}

// SDK represents the Bitcoin Cash SDK for transaction signing and broadcasting.
// BCH uses BIP143-style signature hashing with SIGHASH_FORKID for all transactions.
type SDK struct {
	rpcClient rpcClient
}

// NewSDK creates a new Bitcoin Cash SDK instance.
func NewSDK(rpcClient rpcClient) *SDK {
	return &SDK{
		rpcClient: rpcClient,
	}
}

// Sign converts TSS signatures into fully signed Bitcoin Cash transaction bytes.
// psbtBytes: The unsigned PSBT transaction
// signatures: Map where key is hash of message that was signed, value is the signature
//
// Note: BCH cannot use btcd's psbt.Finalize() because it validates sighash types
// and doesn't understand BCH's SIGHASH_FORKID (0x41). We manually build the scriptSig
// using bchd's ScriptBuilder for proper canonical encoding.
func (sdk *SDK) Sign(psbtBytes []byte, signatures map[string]tss.KeysignResponse) ([]byte, error) {
	pkt, err := psbt.NewFromRawBytes(bytes.NewReader(psbtBytes), false)
	if err != nil {
		return nil, fmt.Errorf("parse psbt: %w", err)
	}

	if pkt.UnsignedTx == nil {
		return nil, fmt.Errorf("PSBT missing unsigned transaction")
	}

	// Create a copy of the unsigned tx to add signatures
	finalTx := pkt.UnsignedTx.Copy()

	for i := range pkt.Inputs {
		input := &pkt.Inputs[i]

		// Calculate BCH-specific signature hash for this input
		sigHash, err := sdk.CalculateInputSignatureHash(pkt, i)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate signature hash for input %d: %w", i, err)
		}

		// Derive key in Signature map
		derivedKey := sdk.deriveKeyFromMessage(sigHash)

		// Find signature matching this input's derived key
		sigResponse, exists := signatures[derivedKey]
		if !exists {
			return nil, fmt.Errorf("missing signature for input %d (derived key: %s)", i, derivedKey)
		}

		// BCH requires SIGHASH_ALL | SIGHASH_FORKID (0x41)
		sighashType := byte(bchtxscript.SigHashAll | SigHashForkID)

		// Convert and append DER signature with BCH sighash type
		der, err := hex.DecodeString(trim0x(sigResponse.DerSignature))
		if err != nil {
			return nil, fmt.Errorf("failed to decode DER signature for input %d: %w", i, err)
		}
		fullSig := make([]byte, len(der)+1)
		copy(fullSig, der)
		fullSig[len(der)] = sighashType

		pubkey, err := sdk.extractPubkeyForInput(input)
		if err != nil {
			return nil, fmt.Errorf("failed to extract pubkey for input %d: %w", i, err)
		}

		// Build P2PKH scriptSig using bchd's ScriptBuilder for proper canonical encoding
		// Format: <sig> <pubkey>
		builder := bchtxscript.NewScriptBuilder()
		builder.AddData(fullSig)
		builder.AddData(pubkey)
		scriptSig, err := builder.Script()
		if err != nil {
			return nil, fmt.Errorf("failed to build scriptSig for input %d: %w", i, err)
		}

		finalTx.TxIn[i].SignatureScript = scriptSig
	}

	// Serialize transaction bytes
	var buf bytes.Buffer
	if err := finalTx.Serialize(&buf); err != nil {
		return nil, fmt.Errorf("serialize final tx: %w", err)
	}

	return buf.Bytes(), nil
}

// CalculateInputSignatureHash calculates the BCH-specific signature hash for a PSBT input.
// BCH uses BIP143-style sighash with SIGHASH_FORKID for ALL transactions (including legacy P2PKH).
func (sdk *SDK) CalculateInputSignatureHash(pkt *psbt.Packet, inputIndex int) ([]byte, error) {
	if inputIndex >= len(pkt.Inputs) {
		return nil, fmt.Errorf("input index %d out of range", inputIndex)
	}

	input := &pkt.Inputs[inputIndex]

	// Get the previous output value and script
	var prevValue int64
	var prevScript []byte

	if input.WitnessUtxo != nil {
		prevValue = input.WitnessUtxo.Value
		prevScript = input.WitnessUtxo.PkScript
	} else if input.NonWitnessUtxo != nil {
		outIndex := pkt.UnsignedTx.TxIn[inputIndex].PreviousOutPoint.Index
		if int(outIndex) >= len(input.NonWitnessUtxo.TxOut) {
			return nil, fmt.Errorf("invalid previous output index %d", outIndex)
		}
		prevOutput := input.NonWitnessUtxo.TxOut[outIndex]
		prevValue = prevOutput.Value
		prevScript = prevOutput.PkScript
	} else {
		return nil, fmt.Errorf("input %d missing both witness and non-witness UTXO", inputIndex)
	}

	// Convert btcd wire.MsgTx to bchd wire.MsgTx
	bchTx, err := convertToBCHTx(pkt.UnsignedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to BCH tx: %w", err)
	}

	// BCH uses SIGHASH_ALL | SIGHASH_FORKID
	sighashType := bchtxscript.SigHashAll | SigHashForkID

	// Calculate signature hashes (needed for BIP143)
	sigHashes := bchtxscript.NewTxSigHashes(bchTx)

	// Use BIP143-style signature hash algorithm (required for BCH post-fork)
	sigHash, _, err := bchtxscript.CalcSignatureHash(
		prevScript,
		sigHashes,
		bchtxscript.SigHashType(sighashType),
		bchTx,
		inputIndex,
		prevValue,
		true, // useBip143SigHashAlgo - always true for BCH post-fork
	)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate BCH signature hash: %w", err)
	}

	return sigHash, nil
}

// Broadcast submits signed transaction to Bitcoin Cash network.
func (sdk *SDK) Broadcast(signedTxBytes []byte) error {
	if sdk.rpcClient == nil {
		return fmt.Errorf("rpc client not configured")
	}

	var tx wire.MsgTx
	if err := tx.Deserialize(bytes.NewReader(signedTxBytes)); err != nil {
		return fmt.Errorf("deserialize signed tx: %w", err)
	}

	_, err := sdk.rpcClient.SendRawTransaction(&tx, false)
	if err != nil {
		return fmt.Errorf("sendrawtransaction failed: %w", err)
	}

	return nil
}

// convertToBCHTx converts a btcd wire.MsgTx to bchd wire.MsgTx.
// Both use the same wire format, so we serialize/deserialize.
func convertToBCHTx(btcTx *wire.MsgTx) (*bchwire.MsgTx, error) {
	var buf bytes.Buffer
	if err := btcTx.Serialize(&buf); err != nil {
		return nil, fmt.Errorf("failed to serialize btcd tx: %w", err)
	}

	bchTx := &bchwire.MsgTx{}
	if err := bchTx.Deserialize(bytes.NewReader(buf.Bytes())); err != nil {
		return nil, fmt.Errorf("failed to deserialize as bchd tx: %w", err)
	}

	return bchTx, nil
}

// deriveKeyFromMessage derives a key from a message hash.
func (sdk *SDK) deriveKeyFromMessage(messageHash []byte) string {
	hash := sha256.Sum256(messageHash)
	return base64.StdEncoding.EncodeToString(hash[:])
}

// extractPubkeyForInput extracts the public key for a PSBT input.
func (sdk *SDK) extractPubkeyForInput(input *psbt.PInput) ([]byte, error) {
	if len(input.Bip32Derivation) > 0 {
		for _, derivation := range input.Bip32Derivation {
			if len(derivation.PubKey) == 33 {
				return derivation.PubKey, nil
			}
		}
	}
	return nil, fmt.Errorf("no public key found in PSBT input")
}

func trim0x(s string) string {
	if len(s) >= 2 && (s[0:2] == "0x" || s[0:2] == "0X") {
		return s[2:]
	}
	return s
}
