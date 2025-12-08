package btc

import (
	"cmp"
	"errors"
	"fmt"
	"slices"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// TransactionBuilder creates unsigned Bitcoin transactions.
type TransactionBuilder struct {
	config ChainConfig
}

// NewTransactionBuilder creates a new TransactionBuilder with the given chain configuration.
func NewTransactionBuilder(config ChainConfig) *TransactionBuilder {
	return &TransactionBuilder{config: config}
}

// BuildTransaction creates an unsigned transaction from the given parameters.
// Returns a PSBT ready for signing with sdk/btc.Sign().
//
// The builder:
//  1. Selects UTXOs using largest-first strategy (unless SelectAll is true)
//  2. Creates outputs from the provided Output list
//  3. Adds a change output if there's remaining value above dust limit
//  4. Creates a PSBT with BIP32 derivations for signing
//  5. Populates PSBT input metadata (WitnessUtxo/NonWitnessUtxo)
func (b *TransactionBuilder) BuildTransaction(params BuildParams) (*BuildResult, error) {
	if err := b.validateBuildParams(params); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}

	// Calculate total output amount
	var totalOutputAmount int64
	for _, out := range params.Outputs {
		totalOutputAmount += out.Amount
	}

	// Determine change address
	changeAddress := params.ChangeAddress
	if changeAddress == "" {
		var err error
		changeAddress, err = PubKeyToP2WPKHAddress(params.FromPubKey, b.config.Network)
		if err != nil {
			return nil, fmt.Errorf("failed to derive change address: %w", err)
		}
	}

	// Select UTXOs
	selectedUTXOs, totalInputValue, err := b.selectUTXOs(params.UTXOs, uint64(totalOutputAmount), params.FeeRate, len(params.Outputs)+1, params.SelectAll)
	if err != nil {
		return nil, fmt.Errorf("UTXO selection failed: %w", err)
	}

	// Create the transaction
	tx := wire.NewMsgTx(wire.TxVersion)

	// Add inputs from selected UTXOs
	for _, utxo := range selectedUTXOs {
		txHash, err := chainhash.NewHashFromStr(utxo.TxHash)
		if err != nil {
			return nil, fmt.Errorf("invalid UTXO tx hash %s: %w", utxo.TxHash, err)
		}

		txIn := &wire.TxIn{
			PreviousOutPoint: wire.OutPoint{
				Hash:  *txHash,
				Index: utxo.Index,
			},
			Sequence: wire.MaxTxInSequenceNum,
		}
		tx.AddTxIn(txIn)
	}

	// Add outputs
	for _, out := range params.Outputs {
		var txOut *wire.TxOut
		var err error

		if out.Address != "" {
			txOut, err = CreateOutput(out.Address, out.Amount, b.config.Network)
		} else if len(out.Data) > 0 {
			txOut, err = CreateOPReturnOutput(out.Data)
		} else {
			return nil, fmt.Errorf("output must have either Address or Data")
		}

		if err != nil {
			return nil, fmt.Errorf("failed to create output: %w", err)
		}
		tx.AddTxOut(txOut)
	}

	// Calculate fee based on actual transaction size
	estimatedSize := b.estimateTxSize(len(selectedUTXOs), len(params.Outputs)+1, params.Outputs) // +1 for change
	fee := CalculateFee(estimatedSize, params.FeeRate)

	// Calculate change
	changeAmount := int64(totalInputValue) - totalOutputAmount - int64(fee)
	changeIndex := -1

	if changeAmount > b.config.DustLimit {
		// Add change output
		changeOut, err := CreateOutput(changeAddress, changeAmount, b.config.Network)
		if err != nil {
			return nil, fmt.Errorf("failed to create change output: %w", err)
		}
		tx.AddTxOut(changeOut)
		changeIndex = len(tx.TxOut) - 1
	} else if changeAmount > 0 {
		// Change is dust, add it to fee
		fee += uint64(changeAmount)
		changeAmount = 0
	} else if changeAmount < 0 {
		return nil, fmt.Errorf("insufficient funds: need %d more satoshis", -changeAmount)
	}

	// Create PSBT
	pkt, err := CreatePSBT(tx, params.FromPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create PSBT: %w", err)
	}

	// Populate PSBT input metadata
	if err := PopulatePSBTInputs(pkt, selectedUTXOs, params.PreviousTxs); err != nil {
		return nil, fmt.Errorf("failed to populate PSBT inputs: %w", err)
	}

	// Serialize PSBT
	psbtBytes, err := SerializePSBT(pkt)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize PSBT: %w", err)
	}

	// Serialize unsigned tx for policy evaluation
	unsignedTxBytes, err := SerializeUnsignedTx(pkt)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize unsigned tx: %w", err)
	}

	return &BuildResult{
		PSBT:            psbtBytes,
		UnsignedTxBytes: unsignedTxBytes,
		SelectedUTXOs:   selectedUTXOs,
		Fee:             fee,
		ChangeAmount:    changeAmount,
		ChangeIndex:     changeIndex,
		EstimatedVBytes: estimatedSize,
	}, nil
}

// BuildSendTransaction is a convenience method for simple send transactions.
// Creates a 2-output transaction: recipient + change.
func (b *TransactionBuilder) BuildSendTransaction(params SendParams) (*BuildResult, error) {
	return b.BuildTransaction(BuildParams{
		UTXOs:         params.UTXOs,
		FeeRate:       params.FeeRate,
		FromPubKey:    params.FromPubKey,
		ChangeAddress: params.ChangeAddress,
		PreviousTxs:   params.PreviousTxs,
		Outputs: []Output{
			{Address: params.ToAddress, Amount: int64(params.Amount)},
		},
	})
}

// BuildSwapTransaction is a convenience method for THORChain-style swap transactions.
// Creates a 3-output transaction: vault + change + OP_RETURN memo.
// Output order: [vault, change, memo] - this matches THORChain expectations.
func (b *TransactionBuilder) BuildSwapTransaction(params SwapParams) (*BuildResult, error) {
	if params.Memo == "" {
		return nil, errors.New("swap transaction requires a memo")
	}

	// For swap transactions, we build outputs manually to control order
	// THORChain expects: vault output, change output, memo output
	outputs := []Output{
		{Address: params.VaultAddress, Amount: int64(params.Amount)},
		// Note: Change will be added by BuildTransaction at the end
		{Data: []byte(params.Memo)}, // OP_RETURN memo
	}

	return b.BuildTransaction(BuildParams{
		UTXOs:         params.UTXOs,
		FeeRate:       params.FeeRate,
		FromPubKey:    params.FromPubKey,
		ChangeAddress: params.ChangeAddress,
		PreviousTxs:   params.PreviousTxs,
		Outputs:       outputs,
	})
}

// selectUTXOs selects UTXOs to cover the target amount plus fees.
// Uses largest-first selection strategy to minimize number of inputs.
func (b *TransactionBuilder) selectUTXOs(utxos []UTXO, targetAmount uint64, feeRate uint64, numOutputs int, selectAll bool) ([]UTXO, uint64, error) {
	if len(utxos) == 0 {
		return nil, 0, errors.New("no UTXOs provided")
	}

	// If selectAll is true, use all UTXOs
	if selectAll {
		var total uint64
		for _, utxo := range utxos {
			total += utxo.Value
		}
		return utxos, total, nil
	}

	// Sort UTXOs by value (descending - largest first)
	sorted := make([]UTXO, len(utxos))
	copy(sorted, utxos)
	slices.SortFunc(sorted, func(a, b UTXO) int {
		return cmp.Compare(b.Value, a.Value) // Descending order
	})

	var selected []UTXO
	var totalValue uint64

	for _, utxo := range sorted {
		selected = append(selected, utxo)
		totalValue += utxo.Value

		// Estimate fee with current number of inputs
		estimatedSize := EstimateTxSize(len(selected), numOutputs, true)
		fee := CalculateFee(estimatedSize, feeRate)

		// Check if we have enough
		if totalValue >= targetAmount+fee {
			return selected, totalValue, nil
		}
	}

	// Not enough funds
	estimatedSize := EstimateTxSize(len(selected), numOutputs, true)
	fee := CalculateFee(estimatedSize, feeRate)
	needed := targetAmount + fee
	return nil, 0, fmt.Errorf("insufficient funds: have %d satoshis, need %d satoshis (amount: %d, fee: %d)",
		totalValue, needed, targetAmount, fee)
}

// estimateTxSize estimates transaction size accounting for different output types.
func (b *TransactionBuilder) estimateTxSize(numInputs, numOutputs int, outputs []Output) int {
	// Count OP_RETURN data length
	var opReturnDataLen int
	for _, out := range outputs {
		if len(out.Data) > 0 {
			opReturnDataLen += len(out.Data)
		}
	}

	// Build input types (assume P2WPKH for modern wallets)
	inputs := make([]InputType, numInputs)
	for i := range inputs {
		inputs[i] = InputP2WPKH
	}

	// Build output types
	outputTypes := make([]OutputType, 0, numOutputs)
	for _, out := range outputs {
		if len(out.Data) > 0 {
			outputTypes = append(outputTypes, OutputOPReturn)
		} else {
			outputTypes = append(outputTypes, OutputP2WPKH) // Assume P2WPKH outputs
		}
	}
	// Add change output
	outputTypes = append(outputTypes, OutputP2WPKH)

	return EstimateTxSizeWithTypes(inputs, outputTypes, opReturnDataLen)
}

// validateBuildParams validates the build parameters.
func (b *TransactionBuilder) validateBuildParams(params BuildParams) error {
	if len(params.UTXOs) == 0 {
		return errors.New("no UTXOs provided")
	}

	if params.FeeRate == 0 {
		return errors.New("fee rate must be greater than 0")
	}

	if len(params.FromPubKey) != 33 {
		return fmt.Errorf("invalid public key length: expected 33 bytes, got %d", len(params.FromPubKey))
	}

	if len(params.Outputs) == 0 {
		return errors.New("no outputs provided")
	}

	for i, out := range params.Outputs {
		if out.Address == "" && len(out.Data) == 0 {
			return fmt.Errorf("output %d must have either Address or Data", i)
		}
		if out.Address != "" && len(out.Data) > 0 {
			return fmt.Errorf("output %d cannot have both Address and Data", i)
		}
		if out.Address != "" && out.Amount < 0 {
			return fmt.Errorf("output %d has negative amount", i)
		}
		if len(out.Data) > 80 {
			return fmt.Errorf("output %d OP_RETURN data exceeds 80 bytes", i)
		}
	}

	return nil
}
