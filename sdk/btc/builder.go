package btc

import (
	"bytes"
	"cmp"
	"errors"
	"fmt"
	"slices"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// Build creates an unsigned transaction from the given parameters.
// Returns a BuildResult with PSBT ready for signing.
//
// Parameters:
//   - utxos: Available UTXOs to spend from
//   - outputs: Pre-built transaction outputs (change output should have Value=0)
//   - changeOutputIndex: Index of the change output in outputs slice
//   - feeRate: Fee rate in sats/vbyte
//   - pubKey: Compressed public key (33 bytes) for PSBT derivation
func (b *Builder) Build(
	utxos []UTXO,
	outputs []*wire.TxOut,
	changeOutputIndex int,
	feeRate uint64,
	pubKey []byte,
) (*BuildResult, error) {
	if err := b.validate(utxos, outputs, changeOutputIndex, feeRate, pubKey); err != nil {
		return nil, err
	}

	// Calculate total output amount (excluding change) and OP_RETURN data length
	var totalOutputAmount int64
	var opReturnDataLen int
	for i, out := range outputs {
		if i != changeOutputIndex {
			totalOutputAmount += out.Value
		}
		// Detect OP_RETURN for size estimation (OP_RETURN = 0x6a)
		if len(out.PkScript) > 0 && out.PkScript[0] == 0x6a {
			opReturnDataLen += len(out.PkScript) - 2 // Subtract OP_RETURN + push opcode
		}
	}

	// Select UTXOs (largest-first)
	selectedUTXOs, totalInputValue, err := b.selectUTXOs(utxos, uint64(totalOutputAmount), feeRate, len(outputs), opReturnDataLen)
	if err != nil {
		return nil, err
	}

	// Create the transaction
	tx := wire.NewMsgTx(wire.TxVersion)

	// Add inputs
	for _, utxo := range selectedUTXOs {
		txHash, err := chainhash.NewHashFromStr(utxo.TxHash)
		if err != nil {
			return nil, fmt.Errorf("invalid UTXO tx hash %s: %w", utxo.TxHash, err)
		}
		tx.AddTxIn(&wire.TxIn{
			PreviousOutPoint: wire.OutPoint{Hash: *txHash, Index: utxo.Index},
			Sequence:         wire.MaxTxInSequenceNum,
		})
	}

	// Add outputs
	for _, out := range outputs {
		tx.AddTxOut(out)
	}

	// Calculate fee
	estimatedSize := EstimateTxVBytes(len(selectedUTXOs), len(outputs), opReturnDataLen)
	fee := CalculateFee(estimatedSize, feeRate)

	// Calculate and set change
	changeAmount := int64(totalInputValue) - totalOutputAmount - int64(fee)

	if changeAmount > b.DustLimit {
		tx.TxOut[changeOutputIndex].Value = changeAmount
	} else if changeAmount > 0 {
		// Change is dust, add to fee
		fee += uint64(changeAmount)
		changeAmount = 0
		tx.TxOut[changeOutputIndex].Value = 0
	} else if changeAmount < 0 {
		return nil, fmt.Errorf("insufficient funds: need %d more satoshis", -changeAmount)
	}

	// Create PSBT
	pkt, err := CreatePSBT(tx, pubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create PSBT: %w", err)
	}

	return &BuildResult{
		Packet:        pkt,
		SelectedUTXOs: selectedUTXOs,
		Fee:           fee,
		ChangeAmount:  changeAmount,
		ChangeIndex:   changeOutputIndex,
	}, nil
}

// PopulatePSBTMetadata populates WitnessUtxo/NonWitnessUtxo for PSBT inputs.
// Call this after Build() if you need to sign the transaction.
//
// For witness inputs (P2WPKH, P2WSH, P2TR), only WitnessUtxo is needed.
// For legacy inputs (P2PKH), NonWitnessUtxo (full previous tx) is needed.
func PopulatePSBTMetadata(result *BuildResult, fetcher PrevTxFetcher) error {
	pkt := result.Packet

	for i := range pkt.Inputs {
		prevOutPoint := pkt.UnsignedTx.TxIn[i].PreviousOutPoint
		txHash := prevOutPoint.Hash.String()

		// Check if we have PkScript from UTXO
		var matchedUTXO *UTXO
		for j := range result.SelectedUTXOs {
			if result.SelectedUTXOs[j].TxHash == txHash && result.SelectedUTXOs[j].Index == prevOutPoint.Index {
				matchedUTXO = &result.SelectedUTXOs[j]
				break
			}
		}

		// If we have PkScript and it's witness, we can populate directly
		if matchedUTXO != nil && len(matchedUTXO.PkScript) > 0 && IsWitnessOutput(matchedUTXO.PkScript) {
			pkt.Inputs[i].WitnessUtxo = &wire.TxOut{
				Value:    int64(matchedUTXO.Value),
				PkScript: matchedUTXO.PkScript,
			}
			continue
		}

		// Need to fetch the previous transaction
		rawTx, err := fetcher.GetRawTransaction(txHash)
		if err != nil {
			return fmt.Errorf("failed to get previous tx %s: %w", txHash, err)
		}

		var prevTx wire.MsgTx
		if err := prevTx.Deserialize(bytes.NewReader(rawTx)); err != nil {
			return fmt.Errorf("failed to deserialize previous tx %s: %w", txHash, err)
		}

		if int(prevOutPoint.Index) >= len(prevTx.TxOut) {
			return fmt.Errorf("invalid output index %d for tx %s", prevOutPoint.Index, txHash)
		}

		prevOutput := prevTx.TxOut[prevOutPoint.Index]
		if IsWitnessOutput(prevOutput.PkScript) {
			pkt.Inputs[i].WitnessUtxo = prevOutput
		} else {
			pkt.Inputs[i].NonWitnessUtxo = &prevTx
		}
	}

	return nil
}

// selectUTXOs selects UTXOs using largest-first strategy.
func (b *Builder) selectUTXOs(utxos []UTXO, targetAmount uint64, feeRate uint64, numOutputs int, opReturnDataLen int) ([]UTXO, uint64, error) {
	// Sort by value descending
	sorted := make([]UTXO, len(utxos))
	copy(sorted, utxos)
	slices.SortFunc(sorted, func(a, b UTXO) int {
		return cmp.Compare(b.Value, a.Value)
	})

	var selected []UTXO
	var totalValue uint64

	for _, utxo := range sorted {
		selected = append(selected, utxo)
		totalValue += utxo.Value

		fee := CalculateFee(EstimateTxVBytes(len(selected), numOutputs, opReturnDataLen), feeRate)
		if totalValue >= targetAmount+fee {
			return selected, totalValue, nil
		}
	}

	fee := CalculateFee(EstimateTxVBytes(len(selected), numOutputs, opReturnDataLen), feeRate)
	return nil, 0, fmt.Errorf("insufficient funds: have %d, need %d (amount: %d, fee: %d)",
		totalValue, targetAmount+fee, targetAmount, fee)
}

func (b *Builder) validate(utxos []UTXO, outputs []*wire.TxOut, changeOutputIndex int, feeRate uint64, pubKey []byte) error {
	if len(utxos) == 0 {
		return errors.New("no UTXOs provided")
	}
	if feeRate == 0 {
		return errors.New("fee rate must be greater than 0")
	}
	if len(pubKey) != 33 {
		return fmt.Errorf("invalid public key length: expected 33, got %d", len(pubKey))
	}
	if len(outputs) == 0 {
		return errors.New("no outputs provided")
	}
	if changeOutputIndex < 0 || changeOutputIndex >= len(outputs) {
		return fmt.Errorf("invalid change output index: %d", changeOutputIndex)
	}
	return nil
}

// IsWitnessOutput returns true if the pkScript is a witness program.
func IsWitnessOutput(pkScript []byte) bool {
	return txscript.IsPayToWitnessPubKeyHash(pkScript) ||
		txscript.IsPayToWitnessScriptHash(pkScript) ||
		txscript.IsPayToTaproot(pkScript)
}
