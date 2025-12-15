package btc

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/btcsuite/btcd/wire"
)

// CreatePSBT creates a PSBT from a wire.MsgTx with BIP32 derivations for each input.
func CreatePSBT(tx *wire.MsgTx, pubkey []byte) (*psbt.Packet, error) {
	if len(pubkey) != 33 {
		return nil, fmt.Errorf("invalid compressed public key length: expected 33, got %d", len(pubkey))
	}

	packet, err := psbt.NewFromUnsignedTx(tx.Copy())
	if err != nil {
		return nil, fmt.Errorf("failed to create PSBT from transaction: %w", err)
	}

	for i := range packet.Inputs {
		packet.Inputs[i].Bip32Derivation = []*psbt.Bip32Derivation{{
			PubKey:    pubkey,
			Bip32Path: nil,
		}}
	}

	return packet, nil
}
