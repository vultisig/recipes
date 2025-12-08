package utxo

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
)

// SelectionStrategy defines how UTXOs should be selected.
type SelectionStrategy int

const (
	// LargestFirst selects UTXOs starting with the largest values.
	// This minimizes the number of inputs and thus transaction fees.
	LargestFirst SelectionStrategy = iota

	// SmallestFirst selects UTXOs starting with the smallest values.
	// This helps consolidate small UTXOs but may increase fees.
	SmallestFirst

	// SelectAll uses all available UTXOs (for consolidation transactions).
	SelectAll
)

// SelectionParams contains parameters for UTXO selection.
type SelectionParams struct {
	// UTXOs is the list of available UTXOs to select from.
	UTXOs []UTXO

	// TargetAmount is the amount needed for outputs (excluding fees).
	TargetAmount uint64

	// FeeRate is the fee rate in smallest units per byte/vbyte.
	FeeRate uint64

	// NumOutputs is the expected number of outputs (including change).
	NumOutputs int

	// Strategy is the UTXO selection strategy to use.
	Strategy SelectionStrategy

	// ChainParams contains chain-specific parameters for size estimation.
	ChainParams ChainParams

	// InputType is the expected input type (e.g., "p2wpkh", "p2pkh").
	// If empty, defaults based on chain's SegWit support.
	InputType string
}

// Select selects UTXOs to cover the target amount plus estimated fees.
// Returns the selection result including selected UTXOs, total value, and fee.
func Select(params SelectionParams) (*SelectionResult, error) {
	if len(params.UTXOs) == 0 {
		return nil, errors.New("no UTXOs provided")
	}

	if params.FeeRate == 0 {
		return nil, errors.New("fee rate must be greater than 0")
	}

	// Handle SelectAll strategy
	if params.Strategy == SelectAll {
		return selectAll(params)
	}

	// Sort UTXOs based on strategy
	sorted := make([]UTXO, len(params.UTXOs))
	copy(sorted, params.UTXOs)

	switch params.Strategy {
	case LargestFirst:
		slices.SortFunc(sorted, func(a, b UTXO) int {
			return cmp.Compare(b.Value, a.Value) // Descending
		})
	case SmallestFirst:
		slices.SortFunc(sorted, func(a, b UTXO) int {
			return cmp.Compare(a.Value, b.Value) // Ascending
		})
	}

	// Determine input type for size estimation
	inputType := params.InputType
	if inputType == "" {
		if params.ChainParams.SupportsSegWit {
			inputType = "p2wpkh"
		} else {
			inputType = "p2pkh"
		}
	}

	// Select UTXOs until we have enough
	var selected []UTXO
	var totalValue uint64

	for _, utxo := range sorted {
		selected = append(selected, utxo)
		totalValue += utxo.Value

		// Estimate fee with current selection
		fee := EstimateFee(EstimateFeeParams{
			NumInputs:   len(selected),
			NumOutputs:  params.NumOutputs,
			FeeRate:     params.FeeRate,
			ChainParams: params.ChainParams,
			InputType:   inputType,
		})

		// Check if we have enough
		needed := params.TargetAmount + fee
		if totalValue >= needed {
			change := int64(totalValue) - int64(params.TargetAmount) - int64(fee)
			return &SelectionResult{
				Selected:     selected,
				TotalValue:   totalValue,
				TargetAmount: params.TargetAmount,
				Fee:          fee,
				Change:       change,
			}, nil
		}
	}

	// All UTXOs exhausted, calculate shortfall
	fee := EstimateFee(EstimateFeeParams{
		NumInputs:   len(selected),
		NumOutputs:  params.NumOutputs,
		FeeRate:     params.FeeRate,
		ChainParams: params.ChainParams,
		InputType:   inputType,
	})
	needed := params.TargetAmount + fee
	shortfall := needed - totalValue

	return nil, fmt.Errorf("insufficient funds: have %d, need %d (amount: %d, fee: %d, shortfall: %d)",
		totalValue, needed, params.TargetAmount, fee, shortfall)
}

// selectAll returns all UTXOs without filtering.
func selectAll(params SelectionParams) (*SelectionResult, error) {
	var totalValue uint64
	for _, utxo := range params.UTXOs {
		totalValue += utxo.Value
	}

	inputType := params.InputType
	if inputType == "" {
		if params.ChainParams.SupportsSegWit {
			inputType = "p2wpkh"
		} else {
			inputType = "p2pkh"
		}
	}

	fee := EstimateFee(EstimateFeeParams{
		NumInputs:   len(params.UTXOs),
		NumOutputs:  params.NumOutputs,
		FeeRate:     params.FeeRate,
		ChainParams: params.ChainParams,
		InputType:   inputType,
	})

	change := int64(totalValue) - int64(params.TargetAmount) - int64(fee)

	return &SelectionResult{
		Selected:     params.UTXOs,
		TotalValue:   totalValue,
		TargetAmount: params.TargetAmount,
		Fee:          fee,
		Change:       change,
	}, nil
}

// SelectWithDustHandling selects UTXOs and handles dust change appropriately.
// If the change would be below the dust limit, it's added to the fee instead.
func SelectWithDustHandling(params SelectionParams) (*SelectionResult, error) {
	result, err := Select(params)
	if err != nil {
		return nil, err
	}

	// If change is dust, add it to fee
	if result.Change > 0 && result.Change < params.ChainParams.DustLimit {
		result.Fee += uint64(result.Change)
		result.Change = 0
	}

	return result, nil
}
