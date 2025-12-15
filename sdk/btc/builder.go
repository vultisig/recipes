package btc

import (
	"bytes"
	"cmp"
	"errors"
	"fmt"
	"slices"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// Build creates an unsigned transaction from the given parameters.
// Returns a BuildResult with PSBT ready for signing.
//
// Parameters:
//   - utxos: Available UTXOs to spend from
//   - outputs: Desired outputs (recipient addresses, OP_RETURN data)
//   - feeRate: Fee rate in sats/vbyte
//   - pubKey: Compressed public key (33 bytes) for PSBT derivation
//   - changeAddr: Change address (if empty, derived from pubKey as P2WPKH)
func (b *Builder) Build(
	utxos []UTXO,
	outputs []Output,
	feeRate uint64,
	pubKey []byte,
	changeAddr string,
) (*BuildResult, error) {
	if err := b.validate(utxos, outputs, feeRate, pubKey); err != nil {
		return nil, err
	}

	// Calculate total output amount and OP_RETURN data length
	var totalOutputAmount int64
	var opReturnDataLen int
	for _, out := range outputs {
		totalOutputAmount += out.Amount
		if len(out.Data) > 0 {
			opReturnDataLen += len(out.Data)
		}
	}

	// Determine change address
	if changeAddr == "" {
		var err error
		changeAddr, err = PubKeyToP2WPKHAddress(pubKey, b.Network)
		if err != nil {
			return nil, fmt.Errorf("failed to derive change address: %w", err)
		}
	}

	// Select UTXOs (largest-first)
	selectedUTXOs, totalInputValue, err := b.selectUTXOs(utxos, uint64(totalOutputAmount), feeRate, len(outputs)+1, opReturnDataLen)
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
		var txOut *wire.TxOut
		var err error
		if out.Address != "" {
			txOut, err = CreateOutput(out.Address, out.Amount, b.Network)
		} else {
			txOut, err = CreateOPReturnOutput(out.Data)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to create output: %w", err)
		}
		tx.AddTxOut(txOut)
	}

	// Calculate fee
	estimatedSize := EstimateTxVBytes(len(selectedUTXOs), len(outputs)+1, opReturnDataLen)
	fee := CalculateFee(estimatedSize, feeRate)

	// Calculate and add change
	changeAmount := int64(totalInputValue) - totalOutputAmount - int64(fee)
	changeIndex := -1

	if changeAmount > b.DustLimit {
		changeOut, err := CreateOutput(changeAddr, changeAmount, b.Network)
		if err != nil {
			return nil, fmt.Errorf("failed to create change output: %w", err)
		}
		tx.AddTxOut(changeOut)
		changeIndex = len(tx.TxOut) - 1
	} else if changeAmount > 0 {
		// Change is dust, add to fee
		fee += uint64(changeAmount)
		changeAmount = 0
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
		ChangeIndex:   changeIndex,
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
		var pkScript []byte
		for _, utxo := range result.SelectedUTXOs {
			if utxo.TxHash == txHash && utxo.Index == prevOutPoint.Index {
				pkScript = utxo.PkScript
				break
			}
		}

		// If we have PkScript and it's witness, we can populate directly
		if len(pkScript) > 0 && IsWitnessOutput(pkScript) {
			for _, utxo := range result.SelectedUTXOs {
				if utxo.TxHash == txHash && utxo.Index == prevOutPoint.Index {
					pkt.Inputs[i].WitnessUtxo = &wire.TxOut{
						Value:    int64(utxo.Value),
						PkScript: pkScript,
					}
					break
				}
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
	if len(utxos) == 0 {
		return nil, 0, errors.New("no UTXOs provided")
	}

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

func (b *Builder) validate(utxos []UTXO, outputs []Output, feeRate uint64, pubKey []byte) error {
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
	for i, out := range outputs {
		if out.Address == "" && len(out.Data) == 0 {
			return fmt.Errorf("output %d must have either Address or Data", i)
		}
		if out.Address != "" && len(out.Data) > 0 {
			return fmt.Errorf("output %d cannot have both Address and Data", i)
		}
		if len(out.Data) > 80 {
			return fmt.Errorf("output %d OP_RETURN data exceeds 80 bytes", i)
		}
	}
	return nil
}
