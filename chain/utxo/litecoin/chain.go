package litecoin

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/vultisig/mobile-tss-lib/tss"

	"github.com/vultisig/recipes/types"
)

// Litecoin implements the Chain interface for the Litecoin blockchain.
// Litecoin supports SegWit like Bitcoin.
type Litecoin struct{}

// ID returns the unique identifier for the Litecoin chain
func (l *Litecoin) ID() string {
	return "litecoin"
}

// Name returns a human-readable name for the Litecoin chain
func (l *Litecoin) Name() string {
	return "Litecoin"
}

// SupportedProtocols returns the list of protocol IDs supported by the Litecoin chain
func (l *Litecoin) SupportedProtocols() []string {
	return []string{"ltc"}
}

// Description returns a human-readable description for the Litecoin chain
func (l *Litecoin) Description() string {
	return "Litecoin is a peer-to-peer cryptocurrency created as a lighter, faster alternative to Bitcoin."
}

func (l *Litecoin) GetProtocol(id string) (types.Protocol, error) {
	if id == "ltc" {
		return NewLTC(), nil
	}
	return nil, fmt.Errorf("protocol %q not found or not supported on Litecoin", id)
}

func (l *Litecoin) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	return ParseLitecoinTransaction(txHex)
}

func (l *Litecoin) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
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

// NewChain creates a new Litecoin chain instance
func NewChain() types.Chain {
	return &Litecoin{}
}

