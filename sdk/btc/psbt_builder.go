package btc

import (
	"bytes"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/btcsuite/btcd/wire"
)

// CreatePSBT creates a PSBT from a wire.MsgTx with BIP32 derivations for each input.
// The pubkey is added to each input's Bip32Derivation field for signing.
func CreatePSBT(tx *wire.MsgTx, pubkey []byte) (*psbt.Packet, error) {
	if len(pubkey) != 33 {
		return nil, fmt.Errorf("invalid compressed public key length: expected 33, got %d", len(pubkey))
	}

	packet, err := psbt.NewFromUnsignedTx(tx.Copy())
	if err != nil {
		return nil, fmt.Errorf("failed to create PSBT from transaction: %w", err)
	}

	// Add BIP32 derivation to each input
	for i := range packet.Inputs {
		derivation := &psbt.Bip32Derivation{
			PubKey:    pubkey,
			Bip32Path: nil, // Path can be empty; signer will use the pubkey directly
		}
		packet.Inputs[i].Bip32Derivation = []*psbt.Bip32Derivation{derivation}
	}

	return packet, nil
}

// PopulatePSBTInputs populates WitnessUtxo or NonWitnessUtxo for each PSBT input.
// This metadata is required for proper signing.
//
// It uses two sources of information:
//   - UTXOs with PkScript: Creates WitnessUtxo directly from the UTXO data
//   - PreviousTxs: For inputs that need full previous transaction (legacy inputs)
//
// For witness inputs (P2WPKH, P2WSH, P2TR), only WitnessUtxo is needed.
// For legacy inputs (P2PKH, P2SH), NonWitnessUtxo (full previous tx) is needed.
func PopulatePSBTInputs(pkt *psbt.Packet, utxos []UTXO, prevTxs []PreviousTx) error {
	// Build a map of previous transactions by hash
	prevTxMap := make(map[string]*wire.MsgTx)
	for _, pt := range prevTxs {
		var msgTx wire.MsgTx
		if err := msgTx.Deserialize(bytes.NewReader(pt.RawTx)); err != nil {
			return fmt.Errorf("failed to deserialize previous tx %s: %w", pt.TxHash, err)
		}
		prevTxMap[pt.TxHash] = &msgTx
	}

	// Build a map of UTXOs by outpoint for quick lookup
	utxoMap := make(map[string]UTXO)
	for _, utxo := range utxos {
		key := fmt.Sprintf("%s:%d", utxo.TxHash, utxo.Index)
		utxoMap[key] = utxo
	}

	// Populate each input
	for i := range pkt.Inputs {
		prevOutPoint := pkt.UnsignedTx.TxIn[i].PreviousOutPoint
		txHash := prevOutPoint.Hash.String()
		outIndex := prevOutPoint.Index

		utxoKey := fmt.Sprintf("%s:%d", txHash, outIndex)
		utxo, hasUTXO := utxoMap[utxoKey]

		// If UTXO has PkScript, we can determine if it's witness or not
		if hasUTXO && len(utxo.PkScript) > 0 {
			if IsWitnessOutput(utxo.PkScript) {
				// Witness input: only need WitnessUtxo
				pkt.Inputs[i].WitnessUtxo = &wire.TxOut{
					Value:    int64(utxo.Value),
					PkScript: utxo.PkScript,
				}
			} else {
				// Legacy input: need full previous transaction
				prevTx, ok := prevTxMap[txHash]
				if !ok {
					return fmt.Errorf("missing previous transaction for legacy input %d (tx: %s)", i, txHash)
				}
				pkt.Inputs[i].NonWitnessUtxo = prevTx
			}
			continue
		}

		// If no PkScript in UTXO, try to get from previous transaction
		prevTx, hasPrevTx := prevTxMap[txHash]
		if !hasPrevTx {
			return fmt.Errorf("input %d: no PkScript in UTXO and no previous transaction provided (tx: %s)", i, txHash)
		}

		if int(outIndex) >= len(prevTx.TxOut) {
			return fmt.Errorf("input %d: invalid output index %d for transaction %s", i, outIndex, txHash)
		}

		prevOutput := prevTx.TxOut[outIndex]

		if IsWitnessOutput(prevOutput.PkScript) {
			pkt.Inputs[i].WitnessUtxo = prevOutput
		} else {
			pkt.Inputs[i].NonWitnessUtxo = prevTx
		}
	}

	return nil
}

// SerializePSBT serializes a PSBT to bytes.
func SerializePSBT(pkt *psbt.Packet) ([]byte, error) {
	var buf bytes.Buffer
	if err := pkt.Serialize(&buf); err != nil {
		return nil, fmt.Errorf("failed to serialize PSBT: %w", err)
	}
	return buf.Bytes(), nil
}

// SerializeUnsignedTx serializes the unsigned transaction from a PSBT.
// This is useful for policy evaluation before signing.
func SerializeUnsignedTx(pkt *psbt.Packet) ([]byte, error) {
	var buf bytes.Buffer
	if err := pkt.UnsignedTx.Serialize(&buf); err != nil {
		return nil, fmt.Errorf("failed to serialize unsigned tx: %w", err)
	}
	return buf.Bytes(), nil
}
