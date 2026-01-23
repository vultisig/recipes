package btc

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/vultisig/mobile-tss-lib/tss"
	sdk "github.com/vultisig/recipes/sdk"
)

// btcd rpcclient interface - methods we use from the client
type rpcClient interface {
	SendRawTransaction(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error)
}

// SDK represents the Bitcoin SDK for transaction signing and broadcasting
type SDK struct {
	rpcClient rpcClient
}

// NewSDK creates a new Bitcoin SDK instance
func NewSDK(rpcClient rpcClient) *SDK {
	return &SDK{
		rpcClient: rpcClient,
	}
}

// Sign converts TSS signatures into fully signed Bitcoin transaction bytes
// psbtBytes: The unsigned PSBT transaction
// signatures: Map where key is hash of message that was signed, value is the signature
func (sdk *SDK) Sign(psbtBytes []byte, signatures map[string]tss.KeysignResponse) ([]byte, error) {
	// Parse PSBT
	pkt, err := psbt.NewFromRawBytes(bytes.NewReader(psbtBytes), false)
	if err != nil {
		return nil, fmt.Errorf("parse psbt: %w", err)
	}

	if pkt.UnsignedTx == nil {
		return nil, fmt.Errorf("PSBT missing unsigned transaction")
	}

	for i := range pkt.Inputs {
		input := &pkt.Inputs[i]
		var sig *tss.KeysignResponse

		// Calculate signature hash for this input
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
		sig = &sigResponse

		// Get sighash type (default to SIGHASH_ALL)
		sighashType := input.SighashType
		if sighashType == 0 {
			sighashType = txscript.SigHashAll
		}

		// Convert and append DER signature
		der, err := hex.DecodeString(trim0x(sig.DerSignature))
		if err != nil {
			return nil, fmt.Errorf("failed to decode DER signature for input %d: %w", i, err)
		}
		fullSig := append(der, byte(sighashType))

		pubkey, err := sdk.extractPubkeyForInput(input)
		if err != nil {
			return nil, fmt.Errorf("failed to extract pubkey for input %d: %w", i, err)
		}

		// Add signature to PSBT
		partialSig := &psbt.PartialSig{
			PubKey:    pubkey,
			Signature: fullSig,
		}
		input.PartialSigs = append(input.PartialSigs, partialSig)
	}

	// Finalize all inputs that have signatures
	for i := 0; i < len(pkt.Inputs); i++ {
		err := psbt.Finalize(pkt, i)
		if err != nil {
			return nil, fmt.Errorf("failed to finalize input %d: %w", i, err)
		}
	}

	// Extract final signed transaction
	finalTx, err := psbt.Extract(pkt)
	if err != nil {
		return nil, fmt.Errorf("extract final tx: %w", err)
	}

	// Serialize transaction bytes
	var buf bytes.Buffer
	if err := finalTx.Serialize(&buf); err != nil {
		return nil, fmt.Errorf("serialize final tx: %w", err)
	}

	return buf.Bytes(), nil
}

// Broadcast submits signed transaction to Bitcoin network
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

// Send is a convenience method that signs and broadcasts the transaction
func (sdk *SDK) Send(psbt []byte, signatures map[string]tss.KeysignResponse) error {
	// Sign the transaction
	signedTxBytes, err := sdk.Sign(psbt, signatures)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Broadcast the signed transaction
	err = sdk.Broadcast(signedTxBytes)
	if err != nil {
		return fmt.Errorf("failed to broadcast transaction: %w", err)
	}

	return nil
}

// deriveKeyFromMessage derives a key from a message hash
func (sdk *SDK) deriveKeyFromMessage(messageHash []byte) string {
	hash := sha256.Sum256(messageHash)

	return base64.StdEncoding.EncodeToString(hash[:])
}

func trim0x(s string) string {
	if len(s) >= 2 && (s[0:2] == "0x" || s[0:2] == "0X") {
		return s[2:]
	}
	return s
}

// calculateInputSignatureHash calculates the signature hash for a specific PSBT input
func (sdk *SDK) CalculateInputSignatureHash(pkt *psbt.Packet, inputIndex int) ([]byte, error) {
	if inputIndex >= len(pkt.Inputs) {
		return nil, fmt.Errorf("input index %d out of range", inputIndex)
	}

	input := &pkt.Inputs[inputIndex]
	tx := pkt.UnsignedTx

	// Get sighash type (default to SIGHASH_ALL)
	sighashType := input.SighashType
	if sighashType == 0 {
		sighashType = txscript.SigHashAll
	}

	// Check if this is a witness transaction
	if input.WitnessUtxo != nil {
		prevOutput := input.WitnessUtxo

		prevFetcher := txscript.NewMultiPrevOutFetcher(nil)
		for i, psbtInput := range pkt.Inputs {
			if psbtInput.WitnessUtxo != nil {
				prevFetcher.AddPrevOut(tx.TxIn[i].PreviousOutPoint, psbtInput.WitnessUtxo)
			} else if psbtInput.NonWitnessUtxo != nil {
				outIndex := tx.TxIn[i].PreviousOutPoint.Index
				if int(outIndex) < len(psbtInput.NonWitnessUtxo.TxOut) {
					prevFetcher.AddPrevOut(tx.TxIn[i].PreviousOutPoint, psbtInput.NonWitnessUtxo.TxOut[outIndex])
				}
			}
		}

		sigHashes := txscript.NewTxSigHashes(tx, prevFetcher)

		// Determine the scriptCode for sighash calculation.
		// P2WSH: use WitnessScript;
		// P2SH-P2WPKH: derive P2PKH script from redeem witness program;
		// native P2WPKH: derive P2PKH script from the witness program.
		// P2TR: Calculate Taproot signature hash
		var sigScript []byte
		switch {
		case txscript.IsPayToWitnessScriptHash(prevOutput.PkScript):
			// P2WSH or P2SH-P2WSH
			sigScript = input.WitnessScript
		case txscript.IsPayToWitnessPubKeyHash(prevOutput.PkScript):
			// Native P2WPKH or P2SH-P2WPKH - create canonical P2PKH script
			var err error
			if input.RedeemScript != nil {
				// P2SH-P2WPKH
				sigScript, err = createP2PKHScriptFromWitnessProgram(input.RedeemScript)
			} else {
				// Native P2WPKH
				sigScript, err = createP2PKHScriptFromWitnessProgram(prevOutput.PkScript)
			}
			if err != nil {
				return nil, fmt.Errorf("native P2WPKH: %w", err)
			}
		case txscript.IsPayToTaproot(prevOutput.PkScript):
			// P2TR immediate return Taproot SigHash
			return txscript.CalcTaprootSignatureHash(
				sigHashes,
				txscript.SigHashType(sighashType),
				tx,
				inputIndex,
				prevFetcher,
			)
		default:
			// Shouldn't land here for witness inputs, safe fallback.
			sigScript = prevOutput.PkScript
		}

		return txscript.CalcWitnessSigHash(
			sigScript,
			sigHashes,
			txscript.SigHashType(sighashType),
			tx,
			inputIndex,
			prevOutput.Value,
		)
	} else if input.NonWitnessUtxo != nil {
		// Non-witness input (legacy P2PKH, P2SH)
		prevTx := input.NonWitnessUtxo
		outIndex := tx.TxIn[inputIndex].PreviousOutPoint.Index
		if int(outIndex) >= len(prevTx.TxOut) {
			return nil, fmt.Errorf("invalid previous output index %d", outIndex)
		}

		prevOutput := prevTx.TxOut[outIndex]

		// Determine the signing script
		var sigScript []byte
		if input.RedeemScript != nil {
			sigScript = input.RedeemScript // P2SH
		} else {
			sigScript = prevOutput.PkScript // P2PKH
		}

		return txscript.CalcSignatureHash(
			sigScript,
			txscript.SigHashType(sighashType),
			tx,
			inputIndex,
		)
	}

	return nil, fmt.Errorf("input %d missing both witness and non-witness UTXO", inputIndex)
}

// createP2PKHScriptFromWitnessProgram extracts pubkey hash from P2WPKH witness program
// and creates the canonical P2PKH script for signature hash calculation
func createP2PKHScriptFromWitnessProgram(witnessProgram []byte) ([]byte, error) {
	if len(witnessProgram) != 22 || witnessProgram[0] != 0x00 || witnessProgram[1] != 0x14 {
		return nil, fmt.Errorf("invalid P2WPKH witness program format")
	}

	// Extract the 20-byte pubkey hash
	pkh := witnessProgram[2:22]

	// Create P2PKH address and script (use mainnet params for script generation)
	addr, err := btcutil.NewAddressPubKeyHash(pkh, &chaincfg.MainNetParams)
	if err != nil {
		return nil, fmt.Errorf("create P2PKH address: %w", err)
	}

	return txscript.PayToAddrScript(addr)
}

// extractPubkeyForInput extracts the public key for a PSBT input
func (s *SDK) extractPubkeyForInput(input *psbt.PInput) ([]byte, error) {
	if len(input.Bip32Derivation) > 0 {
		for _, derivation := range input.Bip32Derivation {
			if len(derivation.PubKey) == 33 {
				return derivation.PubKey, nil
			}
		}
	}

	return nil, fmt.Errorf("no public key found in PSBT input")
}

// DeriveSigningHashes derives the signing hashes from PSBT transaction bytes.
// For Bitcoin, this returns one hash per input since each input requires a separate signature.
func (s *SDK) DeriveSigningHashes(psbtBytes []byte, _ sdk.DeriveOptions) ([]sdk.DerivedHash, error) {
	pkt, err := psbt.NewFromRawBytes(bytes.NewReader(psbtBytes), false)
	if err != nil {
		return nil, fmt.Errorf("parse psbt: %w", err)
	}

	if pkt.UnsignedTx == nil {
		return nil, fmt.Errorf("PSBT missing unsigned transaction")
	}

	hashes := make([]sdk.DerivedHash, 0, len(pkt.Inputs))
	for i := range pkt.Inputs {
		sigHash, err := s.CalculateInputSignatureHash(pkt, i)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate signature hash for input %d: %w", i, err)
		}

		// The Hash field is sha256 of the sighash for lookup purposes
		hash := sha256.Sum256(sigHash)

		hashes = append(hashes, sdk.DerivedHash{
			Message: sigHash,
			Hash:    hash[:],
		})
	}

	return hashes, nil
}
