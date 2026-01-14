package dash

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

// Dash implements the Chain interface for the Dash blockchain.
// Dash does not support SegWit, so all transactions use legacy format.
type Dash struct{}

// ID returns the unique identifier for the Dash chain
func (d *Dash) ID() string {
	return "dash"
}

// Name returns a human-readable name for the Dash chain
func (d *Dash) Name() string {
	return "Dash"
}

// SupportedProtocols returns the list of protocol IDs supported by the Dash chain
func (d *Dash) SupportedProtocols() []string {
	return []string{"dash"}
}

// Description returns a human-readable description for the Dash chain
func (d *Dash) Description() string {
	return "Dash is a cryptocurrency focused on digital cash functionality, featuring InstantSend and PrivateSend capabilities."
}

func (d *Dash) GetProtocol(id string) (types.Protocol, error) {
	if id == "dash" {
		return NewDASH(), nil
	}
	return nil, fmt.Errorf("protocol %q not found or not supported on Dash", id)
}

func (d *Dash) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	return ParseDashTransaction(txHex)
}

func (d *Dash) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	var tx wire.MsgTx
	err := tx.Deserialize(bytes.NewReader(proposedTx))
	if err != nil {
		return "", fmt.Errorf("tx.Deserialize: %w", err)
	}

	if len(tx.TxIn) != len(sigs) {
		return "", fmt.Errorf("input count (%d) does not match sigs count (%d)", len(tx.TxIn), len(sigs))
	}

	// Dash does not support SegWit - always use legacy signature scripts
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

		// For legacy P2PKH, extract the public key from the pre-populated scriptSig
		pubKey, er := utxo.ExtractPubKeyFromScriptSig(in.SignatureScript)
		if er != nil {
			return "", fmt.Errorf("failed to extract pubkey from scriptSig for input %d: %w", i, er)
		}

		// Build scriptSig with DER signature and public key
		scriptSig, er := txscript.NewScriptBuilder().AddData(derSig).AddData(pubKey).Script()
		if er != nil {
			return "", fmt.Errorf("txscript.NewScriptBuilder: %w", er)
		}
		in.SignatureScript = scriptSig
		in.Witness = nil
	}

	return tx.TxHash().String(), nil
}

// NewChain creates a new Dash chain instance
func NewChain() types.Chain {
	return &Dash{}
}

