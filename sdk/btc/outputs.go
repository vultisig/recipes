package btc

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// CreateOutput creates a wire.TxOut from an address and amount.
func CreateOutput(address string, amount int64, network *chaincfg.Params) (*wire.TxOut, error) {
	addr, err := btcutil.DecodeAddress(address, network)
	if err != nil {
		return nil, fmt.Errorf("failed to decode address: %w", err)
	}

	pkScript, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create pkScript: %w", err)
	}

	return wire.NewTxOut(amount, pkScript), nil
}

// CreateOPReturnOutput creates an OP_RETURN output for data embedding.
// The data must be 80 bytes or less.
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

	return wire.NewTxOut(0, pkScript), nil
}

// PubKeyToP2WPKHAddress derives a P2WPKH address from a compressed public key.
func PubKeyToP2WPKHAddress(pubkey []byte, network *chaincfg.Params) (string, error) {
	if len(pubkey) != 33 {
		return "", fmt.Errorf("invalid compressed public key length: expected 33, got %d", len(pubkey))
	}

	witnessProg := btcutil.Hash160(pubkey)
	addr, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, network)
	if err != nil {
		return "", fmt.Errorf("failed to create P2WPKH address: %w", err)
	}

	return addr.EncodeAddress(), nil
}

// IsWitnessOutput returns true if the pkScript is a witness program.
func IsWitnessOutput(pkScript []byte) bool {
	return txscript.IsPayToWitnessPubKeyHash(pkScript) ||
		txscript.IsPayToWitnessScriptHash(pkScript) ||
		txscript.IsPayToTaproot(pkScript)
}
