package btc

// Size constants for Bitcoin transaction fee estimation (virtual bytes).
const (
	// TxOverheadVBytes is the base transaction overhead.
	TxOverheadVBytes = 11

	// Input sizes (virtual bytes)
	P2PKHInputVBytes  = 148 // Legacy P2PKH input
	P2WPKHInputVBytes = 68  // Native SegWit P2WPKH input
	P2TRInputVBytes   = 58  // Taproot P2TR input
	P2WSHInputVBytes  = 91  // P2WSH input

	// Output sizes (bytes)
	P2PKHOutputVBytes  = 34 // Legacy P2PKH output
	P2WPKHOutputVBytes = 31 // Native SegWit P2WPKH output
	P2TROutputVBytes   = 43 // Taproot P2TR output
	P2WSHOutputVBytes  = 43 // P2WSH output
	OPReturnBaseVBytes = 11 // OP_RETURN output base (without data)
)

// EstimateTxVBytes estimates the virtual size of a transaction.
func EstimateTxVBytes(numInputs, numOutputs int, opReturnDataLen int) int {
	// Assume P2WPKH for modern wallets
	size := TxOverheadVBytes
	size += numInputs * P2WPKHInputVBytes
	size += numOutputs * P2WPKHOutputVBytes
	if opReturnDataLen > 0 {
		size += OPReturnBaseVBytes + opReturnDataLen
	}
	return size
}

// CalculateFee returns the fee in satoshis for a given vbyte size and fee rate.
func CalculateFee(vbytes int, satsPerVByte uint64) uint64 {
	return uint64(vbytes) * satsPerVByte
}
