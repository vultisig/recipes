package bitcoin

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	v1 "github.com/vultisig/commondata/go/vultisig/vault/v1"
	"github.com/vultisig/mobile-tss-lib/tss"

	"github.com/vultisig/recipes/address"
	"github.com/vultisig/recipes/common"

	"github.com/vultisig/recipes/types"
)

// Bitcoin implements the Chain interface for the Bitcoin blockchain
type Bitcoin struct{}

// ID returns the unique identifier for the Bitcoin chain
func (b *Bitcoin) ID() string {
	return "bitcoin"
}

// Name returns a human-readable name for the Bitcoin chain
func (b *Bitcoin) Name() string {
	return "Bitcoin"
}

// SupportedProtocols returns the list of protocol IDs supported by the Bitcoin chain
func (b *Bitcoin) SupportedProtocols() []string {
	return []string{"btc"}
}

// Description returns a human-readable description for the Bitcoin chain
func (b *Bitcoin) Description() string {
	return "Bitcoin is a digital currency that is not controlled by any government or financial institution."
}

func (b *Bitcoin) GetProtocol(id string) (types.Protocol, error) {
	if id == "btc" {
		return NewBTC(), nil
	}
	return nil, fmt.Errorf("protocol %q not found or not supported on Bitcoin", id)
}

func (b *Bitcoin) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	return ParseBitcoinTransaction(txHex)
}

func (b *Bitcoin) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	// TODO: @webpiratt: check and add unit tests for witness and non-witness transactions
	var tx wire.MsgTx
	err := tx.Deserialize(bytes.NewReader(proposedTx))
	if err != nil {
		return "", fmt.Errorf("tx.Deserialize: %w", err)
	}

	if len(tx.TxIn) != len(sigs) {
		return "", fmt.Errorf("input count (%d) does not match sigs count (%d)", len(tx.TxIn), len(sigs))
	}

	witness := tx.HasWitness()

	for i, in := range tx.TxIn {
		selectedSig := sigs[i]
		r, er := hex.DecodeString(selectedSig.R)
		if er != nil {
			return "", fmt.Errorf("hex.DecodeString(selectedSig.R): %w", er)
		}
		s, er := hex.DecodeString(selectedSig.S)
		if er != nil {
			return "", fmt.Errorf("hex.DecodeString(selectedSig.S): %w", er)
		}

		var sig []byte
		sig = append(sig, r...)
		sig = append(sig, s...)
		sig = append(sig, byte(txscript.SigHashAll))

		if witness {
			witnessPubKey := in.Witness[1] // must be set by tx proposer
			in.Witness = wire.TxWitness{sig, witnessPubKey}
			in.SignatureScript = nil
		} else {
			scriptSig, er2 := txscript.NewScriptBuilder().AddData(sig).Script()
			if er2 != nil {
				return "", fmt.Errorf("txscript.NewScriptBuilder: %w", er2)
			}
			in.SignatureScript = scriptSig
			in.Witness = nil
		}
	}

	return tx.TxHash().String(), nil
}

// ValidateInvariants ensures that a btc transfer follows required invariants
func (b *Bitcoin) ValidateInvariants(context map[string]interface{}, tx types.DecodedTransaction) error {
	btcTx, ok := tx.(*ParsedBitcoinTransaction)
	if !ok {
		return fmt.Errorf("expected Bitcoin transaction, got %T", tx)
	}
	vaultInterface, exists := context["vault"]
	if !exists {
		return fmt.Errorf("vault is required for Bitcoin invariant validation")
	}

	vault, ok := vaultInterface.(*v1.Vault)
	if !ok {
		return fmt.Errorf("expected *v1.Vault, got %T", vaultInterface)
	}

	// 1. Validate transaction structure (extensible for future tx types)
	if err := b.checkTransactionStructure(btcTx); err != nil {
		return fmt.Errorf("transaction structure validation failed: %w", err)
	}

	// 2. Validate change output back to sender (derived from vault)
	if err := b.validateChangeOutput(vault, btcTx); err != nil {
		return fmt.Errorf("change output validation failed: %w", err)
	}

	return nil
}

// checkTransactionStructure validates basic structure requirements
func (b *Bitcoin) checkTransactionStructure(btcTx *ParsedBitcoinTransaction) error {
	outputs := btcTx.GetAllOutputs()

	// Support standard patterns:
	// 1 output: entire UTXO spent (no change)
	// 2 outputs: recipient + change
	// Can expand to other transaction types here
	if len(outputs) > 2 {
		return fmt.Errorf("transaction must have exactly 2 outputs for standard transfer, got %d", len(outputs))
	}

	return nil
}

// validateChangeOutput validates that the last output goes back to vault-derived address
func (b *Bitcoin) validateChangeOutput(vault *v1.Vault, btcTx *ParsedBitcoinTransaction) error {
	outputs := btcTx.GetAllOutputs()
	if len(outputs) == 0 {
		return fmt.Errorf("transaction has no outputs")
	}

	// For 1-output transactions, no change validation needed (entire UTXO spent)
	if len(outputs) == 1 {
		return nil
	}

	// Last output should always be change back to sender (vault)
	lastOutput := outputs[len(outputs)-1]

	// Derive expected vault address for this chain
	expectedChangeAddr, err := b.deriveVaultAddress(vault)
	if err != nil {
		return fmt.Errorf("failed to derive vault address: %w", err)
	}

	// Check if last output matches vault address
	if lastOutput.Address != expectedChangeAddr {
		return fmt.Errorf("last output address %s does not match vault address %s",
			lastOutput.Address, expectedChangeAddr)
	}

	return nil
}

// deriveVaultAddress derives the Bitcoin address from vault
func (b *Bitcoin) deriveVaultAddress(vault *v1.Vault) (string, error) {
	// Use your address derivation utilities
	vaultAddr, _, _, err := address.GetAddress(
		vault.PublicKeyEcdsa,
		vault.HexChainCode,
		common.Bitcoin,
	)
	if err != nil {
		return "", fmt.Errorf("failed to derive Bitcoin address: %w", err)
	}

	return vaultAddr, nil
}

// NewBitcoin creates a new Bitcoin chain instance
func NewChain() types.Chain {
	return &Bitcoin{}
}
