package bitcoincash

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/vultisig/mobile-tss-lib/tss"

	"github.com/vultisig/recipes/types"
)

// BitcoinCash implements the Chain interface for the Bitcoin Cash blockchain.
// Bitcoin Cash does not support SegWit, so all transactions use legacy format.
type BitcoinCash struct{}

// ID returns the unique identifier for the Bitcoin Cash chain
func (b *BitcoinCash) ID() string {
	return "bitcoincash"
}

// Name returns a human-readable name for the Bitcoin Cash chain
func (b *BitcoinCash) Name() string {
	return "Bitcoin Cash"
}

// SupportedProtocols returns the list of protocol IDs supported by the Bitcoin Cash chain
func (b *BitcoinCash) SupportedProtocols() []string {
	return []string{"bch"}
}

// Description returns a human-readable description for the Bitcoin Cash chain
func (b *BitcoinCash) Description() string {
	return "Bitcoin Cash is a peer-to-peer electronic cash system that forked from Bitcoin with larger block sizes."
}

func (b *BitcoinCash) GetProtocol(id string) (types.Protocol, error) {
	if id == "bch" {
		return NewBCH(), nil
	}
	return nil, fmt.Errorf("protocol %q not found or not supported on Bitcoin Cash", id)
}

func (b *BitcoinCash) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	return ParseBitcoinCashTransaction(txHex)
}

func (b *BitcoinCash) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	var tx wire.MsgTx
	err := tx.Deserialize(bytes.NewReader(proposedTx))
	if err != nil {
		return "", fmt.Errorf("tx.Deserialize: %w", err)
	}

	if len(tx.TxIn) != len(sigs) {
		return "", fmt.Errorf("input count (%d) does not match sigs count (%d)", len(tx.TxIn), len(sigs))
	}

	// Bitcoin Cash does not support SegWit - always use legacy signature scripts
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
		// Bitcoin Cash uses SIGHASH_ALL | SIGHASH_FORKID (0x41)
		sig = append(sig, byte(txscript.SigHashAll|0x40))

		scriptSig, er := txscript.NewScriptBuilder().AddData(sig).Script()
		if er != nil {
			return "", fmt.Errorf("txscript.NewScriptBuilder: %w", er)
		}
		in.SignatureScript = scriptSig
		in.Witness = nil
	}

	return tx.TxHash().String(), nil
}

// NewChain creates a new Bitcoin Cash chain instance
func NewChain() types.Chain {
	return &BitcoinCash{}
}

