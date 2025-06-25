package bitcoin

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/vultisig/mobile-tss-lib/tss"

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

func (b *Bitcoin) ComputeTxHash(proposedTxHex string, sigs []tss.KeysignResponse) (string, error) {
	// TODO: @webpiratt: check and add unit tests for witness and non-witness transactions

	txBytes, err := hex.DecodeString(proposedTxHex)
	if err != nil {
		return "", fmt.Errorf("hex.DecodeString: %w", err)
	}

	var tx wire.MsgTx
	err = tx.Deserialize(bytes.NewReader(txBytes))
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

// ValidateInvariants ensures that a btc transfer always has a change output to prevent malicious transactions
func (b *Bitcoin) ValidateInvariants(tx types.DecodedTransaction) error {
	btcTx, ok := tx.(*ParsedBitcoinTransaction)
	if !ok {
		return fmt.Errorf("expected Bitcoin transaction, got %T", tx)
	}
	// Get all outputs
	outputs := btcTx.GetAllOutputs()

	// Validate that there are exactly 2 outputs
	if len(outputs) != 2 {
		return fmt.Errorf("transaction must have exactly 2 outputs, got %d", len(outputs))
	}

	// TODO: Validate that the last output address is the sender as the change output
	// Placeholder - need to implement sender address extraction from inputs
	_ = outputs[1]
	return nil
}

// NewBitcoin creates a new Bitcoin chain instance
func NewChain() types.Chain {
	return &Bitcoin{}
}
