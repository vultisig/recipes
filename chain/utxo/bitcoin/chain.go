package bitcoin

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/vultisig/mobile-tss-lib/tss"

	"github.com/vultisig/recipes/chain/utxo"
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

		// Encode signature in DER format as required by BIP66
		derSig := utxo.EncodeDERSignature(r, s)
		derSig = append(derSig, byte(txscript.SigHashAll))

		if witness {
			if len(in.Witness) < 2 {
				return "", fmt.Errorf("invalid witness structure for input %d: expected at least 2 elements, got %d", i, len(in.Witness))
			}
			witnessPubKey := in.Witness[1] // must be set by tx proposer
			in.Witness = wire.TxWitness{derSig, witnessPubKey}
			in.SignatureScript = nil
		} else {
			// For legacy P2PKH, extract the public key from the pre-populated scriptSig
			pubKey, er2 := utxo.ExtractPubKeyFromScriptSig(in.SignatureScript)
			if er2 != nil {
				return "", fmt.Errorf("failed to extract pubkey from scriptSig for input %d: %w", i, er2)
			}

			// Build scriptSig with DER signature and public key
			scriptSig, er2 := txscript.NewScriptBuilder().AddData(derSig).AddData(pubKey).Script()
			if er2 != nil {
				return "", fmt.Errorf("txscript.NewScriptBuilder: %w", er2)
			}
			in.SignatureScript = scriptSig
			in.Witness = nil
		}
	}

	return tx.TxHash().String(), nil
}

// NewBitcoin creates a new Bitcoin chain instance
func NewChain() types.Chain {
	return &Bitcoin{}
}

