package btc

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// CreateOutput creates a wire.TxOut from an address and amount.
// Automatically detects address type and creates the appropriate pkScript.
func CreateOutput(address string, amount int64, network *chaincfg.Params) (*wire.TxOut, error) {
	pkScript, err := AddressToPkScript(address, network)
	if err != nil {
		return nil, fmt.Errorf("failed to create pkScript for address %s: %w", address, err)
	}

	return wire.NewTxOut(amount, pkScript), nil
}

// CreateOPReturnOutput creates an OP_RETURN output for data embedding.
// The data must be 80 bytes or less (Bitcoin standard relay policy).
func CreateOPReturnOutput(data []byte) (*wire.TxOut, error) {
	if len(data) > 80 {
		return nil, fmt.Errorf("OP_RETURN data exceeds 80 bytes: %d", len(data))
	}

	builder := txscript.NewScriptBuilder()
	builder.AddOp(txscript.OP_RETURN)
	builder.AddData(data)

	pkScript, err := builder.Script()
	if err != nil {
		return nil, fmt.Errorf("failed to build OP_RETURN script: %w", err)
	}

	// OP_RETURN outputs have 0 value
	return wire.NewTxOut(0, pkScript), nil
}

// AddressToPkScript converts an address string to its corresponding pkScript.
// Supports P2PKH, P2SH, P2WPKH, P2WSH, and P2TR address formats.
func AddressToPkScript(address string, network *chaincfg.Params) ([]byte, error) {
	addr, err := btcutil.DecodeAddress(address, network)
	if err != nil {
		return nil, fmt.Errorf("failed to decode address: %w", err)
	}

	pkScript, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create pkScript: %w", err)
	}

	return pkScript, nil
}

// PubKeyToP2WPKHAddress derives a P2WPKH (native SegWit) address from a compressed public key.
func PubKeyToP2WPKHAddress(pubkey []byte, network *chaincfg.Params) (string, error) {
	if len(pubkey) != 33 {
		return "", fmt.Errorf("invalid compressed public key length: expected 33, got %d", len(pubkey))
	}

	// Create witness pubkey hash address
	witnessProg := btcutil.Hash160(pubkey)
	addr, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, network)
	if err != nil {
		return "", fmt.Errorf("failed to create P2WPKH address: %w", err)
	}

	return addr.EncodeAddress(), nil
}

// PubKeyToP2PKHAddress derives a P2PKH (legacy) address from a compressed public key.
func PubKeyToP2PKHAddress(pubkey []byte, network *chaincfg.Params) (string, error) {
	if len(pubkey) != 33 {
		return "", fmt.Errorf("invalid compressed public key length: expected 33, got %d", len(pubkey))
	}

	pubkeyHash := btcutil.Hash160(pubkey)
	addr, err := btcutil.NewAddressPubKeyHash(pubkeyHash, network)
	if err != nil {
		return "", fmt.Errorf("failed to create P2PKH address: %w", err)
	}

	return addr.EncodeAddress(), nil
}

// DetectOutputType determines the output type from a pkScript.
func DetectOutputType(pkScript []byte) OutputType {
	switch {
	case txscript.IsPayToWitnessPubKeyHash(pkScript):
		return OutputP2WPKH
	case txscript.IsPayToPubKeyHash(pkScript):
		return OutputP2PKH
	case txscript.IsPayToTaproot(pkScript):
		return OutputP2TR
	case txscript.IsPayToWitnessScriptHash(pkScript):
		return OutputP2WSH
	case len(pkScript) > 0 && pkScript[0] == txscript.OP_RETURN:
		return OutputOPReturn
	default:
		return OutputP2WPKH // Default assumption
	}
}

// DetectInputType determines the input type from a pkScript of the UTXO being spent.
func DetectInputType(pkScript []byte) InputType {
	switch {
	case txscript.IsPayToWitnessPubKeyHash(pkScript):
		return InputP2WPKH
	case txscript.IsPayToPubKeyHash(pkScript):
		return InputP2PKH
	case txscript.IsPayToTaproot(pkScript):
		return InputP2TR
	case txscript.IsPayToWitnessScriptHash(pkScript):
		return InputP2WSH
	default:
		return InputP2WPKH // Default assumption for modern wallets
	}
}

// IsWitnessOutput returns true if the pkScript is a witness program (SegWit or Taproot).
func IsWitnessOutput(pkScript []byte) bool {
	return txscript.IsPayToWitnessPubKeyHash(pkScript) ||
		txscript.IsPayToWitnessScriptHash(pkScript) ||
		txscript.IsPayToTaproot(pkScript)
}
