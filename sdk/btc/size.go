package btc

// Size constants for Bitcoin transaction fee estimation.
// These are virtual bytes (vbytes) for SegWit transactions.
const (
	// TxOverheadVBytes is the base transaction overhead:
	// version (4) + locktime (4) + segwit marker/flag (0.5) + input count (1) + output count (1)
	TxOverheadVBytes = 11

	// Input sizes (virtual bytes)
	P2PKHInputVBytes  = 148 // Legacy P2PKH input
	P2WPKHInputVBytes = 68  // Native SegWit P2WPKH input
	P2TRInputVBytes   = 58  // Taproot P2TR input (Schnorr signature)
	P2WSHInputVBytes  = 91  // P2WSH input (conservative estimate for single-sig)

	// Output sizes (bytes, same as vbytes for outputs)
	P2PKHOutputVBytes    = 34 // Legacy P2PKH output
	P2WPKHOutputVBytes   = 31 // Native SegWit P2WPKH output
	P2TROutputVBytes     = 43 // Taproot P2TR output
	P2WSHOutputVBytes    = 43 // P2WSH output
	OPReturnBaseVBytes   = 11 // OP_RETURN output base (without data)
	OPReturnPerByteVSize = 1  // Additional vbytes per byte of OP_RETURN data
)

// InputType represents the type of transaction input for size estimation.
type InputType int

const (
	InputP2PKH InputType = iota
	InputP2WPKH
	InputP2TR
	InputP2WSH
)

// OutputType represents the type of transaction output for size estimation.
type OutputType int

const (
	OutputP2PKH OutputType = iota
	OutputP2WPKH
	OutputP2TR
	OutputP2WSH
	OutputOPReturn
)

// InputSize returns the virtual size in bytes for a given input type.
func InputSize(inputType InputType) int {
	switch inputType {
	case InputP2PKH:
		return P2PKHInputVBytes
	case InputP2WPKH:
		return P2WPKHInputVBytes
	case InputP2TR:
		return P2TRInputVBytes
	case InputP2WSH:
		return P2WSHInputVBytes
	default:
		return P2WPKHInputVBytes // Default to SegWit
	}
}

// OutputSize returns the size in bytes for a given output type.
// For OP_RETURN, pass the data length as the second parameter.
func OutputSize(outputType OutputType, dataLen ...int) int {
	switch outputType {
	case OutputP2PKH:
		return P2PKHOutputVBytes
	case OutputP2WPKH:
		return P2WPKHOutputVBytes
	case OutputP2TR:
		return P2TROutputVBytes
	case OutputP2WSH:
		return P2WSHOutputVBytes
	case OutputOPReturn:
		if len(dataLen) > 0 {
			return OPReturnBaseVBytes + dataLen[0]
		}
		return OPReturnBaseVBytes
	default:
		return P2WPKHOutputVBytes // Default to SegWit
	}
}

// EstimateTxSize estimates the virtual size of a transaction.
// Assumes all inputs are the same type (SegWit by default).
func EstimateTxSize(numInputs, numOutputs int, isSegwit bool) int {
	inputSize := P2WPKHInputVBytes
	outputSize := P2WPKHOutputVBytes
	if !isSegwit {
		inputSize = P2PKHInputVBytes
		outputSize = P2PKHOutputVBytes
	}

	return TxOverheadVBytes + (numInputs * inputSize) + (numOutputs * outputSize)
}

// EstimateTxSizeWithTypes provides more accurate size estimation
// when input and output types are known.
func EstimateTxSizeWithTypes(inputs []InputType, outputs []OutputType, opReturnDataLen int) int {
	size := TxOverheadVBytes

	for _, input := range inputs {
		size += InputSize(input)
	}

	for _, output := range outputs {
		if output == OutputOPReturn {
			size += OutputSize(output, opReturnDataLen)
		} else {
			size += OutputSize(output)
		}
	}

	return size
}

// CalculateFee returns the fee in satoshis for a given virtual size and fee rate.
func CalculateFee(vbytes int, satsPerVByte uint64) uint64 {
	return uint64(vbytes) * satsPerVByte
}

// EstimateFee estimates the fee for a transaction with the given parameters.
// This is a convenience function that combines size estimation and fee calculation.
func EstimateFee(numInputs, numOutputs int, isSegwit bool, satsPerVByte uint64) uint64 {
	size := EstimateTxSize(numInputs, numOutputs, isSegwit)
	return CalculateFee(size, satsPerVByte)
}
