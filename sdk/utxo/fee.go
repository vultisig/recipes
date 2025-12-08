package utxo

// EstimateFeeParams contains parameters for fee estimation.
type EstimateFeeParams struct {
	// NumInputs is the number of transaction inputs.
	NumInputs int

	// NumOutputs is the number of transaction outputs.
	NumOutputs int

	// FeeRate is the fee rate in smallest units per byte/vbyte.
	FeeRate uint64

	// ChainParams contains chain-specific size parameters.
	ChainParams ChainParams

	// InputType is the type of inputs (e.g., "p2wpkh", "p2pkh").
	// If empty, uses chain's default.
	InputType string

	// OutputType is the type of outputs (e.g., "p2wpkh", "p2pkh").
	// If empty, uses chain's default.
	OutputType string

	// OPReturnSize is the size of OP_RETURN data (0 if none).
	OPReturnSize int
}

// EstimateTxSize estimates the transaction size in bytes/vbytes.
func EstimateTxSize(params EstimateFeeParams) int {
	inputType := params.InputType
	if inputType == "" {
		if params.ChainParams.SupportsSegWit {
			inputType = "p2wpkh"
		} else {
			inputType = "p2pkh"
		}
	}

	outputType := params.OutputType
	if outputType == "" {
		if params.ChainParams.SupportsSegWit {
			outputType = "p2wpkh"
		} else {
			outputType = "p2pkh"
		}
	}

	size := params.ChainParams.TxOverhead
	size += params.NumInputs * params.ChainParams.GetInputSize(inputType)
	size += params.NumOutputs * params.ChainParams.GetOutputSize(outputType)

	// Add OP_RETURN size if present
	if params.OPReturnSize > 0 {
		size += params.ChainParams.GetOutputSize("opreturn") + params.OPReturnSize
	}

	return size
}

// EstimateFee estimates the transaction fee.
func EstimateFee(params EstimateFeeParams) uint64 {
	size := EstimateTxSize(params)
	return CalculateFee(size, params.FeeRate)
}

// CalculateFee calculates fee from size and rate.
func CalculateFee(sizeBytes int, feeRate uint64) uint64 {
	return uint64(sizeBytes) * feeRate
}

// EstimateTxSizeDetailed provides detailed size estimation with mixed input/output types.
type DetailedSizeParams struct {
	// InputTypes is a list of input types for each input.
	InputTypes []string

	// OutputTypes is a list of output types for each output.
	OutputTypes []string

	// OPReturnSize is the size of OP_RETURN data (0 if none).
	OPReturnSize int

	// ChainParams contains chain-specific size parameters.
	ChainParams ChainParams
}

// EstimateTxSizeDetailed estimates size with mixed input/output types.
func EstimateTxSizeDetailed(params DetailedSizeParams) int {
	size := params.ChainParams.TxOverhead

	for _, inputType := range params.InputTypes {
		size += params.ChainParams.GetInputSize(inputType)
	}

	for _, outputType := range params.OutputTypes {
		if outputType == "opreturn" {
			size += params.ChainParams.GetOutputSize("opreturn") + params.OPReturnSize
		} else {
			size += params.ChainParams.GetOutputSize(outputType)
		}
	}

	return size
}

// MinimumFee returns the minimum fee for a transaction to be relayed.
// This is typically 1 sat/vbyte for most networks.
func MinimumFee(txSize int) uint64 {
	return uint64(txSize) // 1 sat/byte minimum
}

// RecommendedFee returns a recommended fee based on priority.
type FeePriority int

const (
	FeeLow    FeePriority = iota // Economy, may take longer to confirm
	FeeMedium                     // Standard, typically next few blocks
	FeeHigh                       // Priority, likely next block
)

// SuggestedFeeRates contains suggested fee rates for different priorities.
// These are placeholders; actual rates should come from a fee estimation service.
type SuggestedFeeRates struct {
	Low    uint64 // sats/vbyte for low priority
	Medium uint64 // sats/vbyte for medium priority
	High   uint64 // sats/vbyte for high priority
}

// DefaultFeeRates returns conservative default fee rates.
// In production, these should come from a fee estimation API.
func DefaultFeeRates() SuggestedFeeRates {
	return SuggestedFeeRates{
		Low:    1,
		Medium: 5,
		High:   20,
	}
}
