package btc

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/vultisig/mobile-tss-lib/tss"
)

// btcd rpcclient interface - methods we use from the client
type rpcClient interface {
	SendRawTransaction(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error)
	GetBlockCount() (int64, error)
	Ping() error
}

// SDK represents the Bitcoin SDK for transaction signing and broadcasting
type SDK struct {
	chainID   *big.Int
	rpcClient rpcClient
}

// NewSDK creates a new Bitcoin SDK instance
func NewSDK(chainID *big.Int, rpcClient rpcClient) *SDK {
	return &SDK{
		chainID:   chainID,
		rpcClient: rpcClient,
	}
}

// Sign converts TSS signatures into fully signed Bitcoin transaction bytes
// psbtBytes: The unsigned PSBT transaction
// signatures: Map where key is hash of message that was signed, value is the signature
func (sdk *SDK) Sign(psbtBytes []byte, signatures map[string]tss.KeysignResponse) ([]byte, error) {
	// Parse PSBT
	// TODO confirm psbt format
	pkt, err := psbt.NewFromRawBytes(bytes.NewReader(psbtBytes), true)
	if err != nil {
		return nil, fmt.Errorf("parse psbt: %w", err)
	}

	if pkt.UnsignedTx == nil {
		return nil, fmt.Errorf("PSBT missing unsigned transaction")
	}

	for i, input := range pkt.Inputs {
		var sig *tss.KeysignResponse
		var messageHash string

		// TODO: Need to calculate the signature hash for this input to find matching signature
		calculatedHash, err := sdk.calculateInputSignatureHash(pkt, i)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate signature hash for input %d: %w", i, err)
		}

		messageHash = hex.EncodeToString(calculatedHash)

		// Find signature matching this input's message hash
		sigResponse, exists := signatures[messageHash]
		if !exists {
			return nil, fmt.Errorf("missing signature for input %d (message hash: %s)", i, messageHash)
		}
		sig = &sigResponse

		// Get sighash type (default to SIGHASH_ALL)
		sighashType := input.SighashType
		if sighashType == 0 {
			sighashType = txscript.SigHashAll
		}

		// Convert and append DER signature
		der, err := hex.DecodeString(trim0x(sig.DerSignature))
		if err != nil {
			return nil, fmt.Errorf("failed to decode DER signature for input %d: %w", i, err)
		}
		fullSig := append(der, byte(sighashType))

		pubkey, err := sdk.extractPubkeyForInput(&input)
		if err != nil {
			return nil, fmt.Errorf("failed to extract pubkey for input %d: %w", i, err)
		}

		// Add signature to PSBT
		partialSig := &psbt.PartialSig{
			PubKey:    pubkey,
			Signature: fullSig,
		}
		input.PartialSigs = append(input.PartialSigs, partialSig)
	}

	// Finalize all inputs that have signatures
	for i := 0; i < len(pkt.Inputs); i++ {
		err := psbt.Finalize(pkt, i)
		if err != nil {
			return nil, fmt.Errorf("failed to finalize input %d: %w", i, err)
		}
	}

	// Extract final signed transaction
	finalTx, err := psbt.Extract(pkt)
	if err != nil {
		return nil, fmt.Errorf("extract final tx: %w", err)
	}

	// Serialize transaction bytes
	var buf bytes.Buffer
	if err := finalTx.Serialize(&buf); err != nil {
		return nil, fmt.Errorf("serialize final tx: %w", err)
	}

	return buf.Bytes(), nil
}

func trim0x(s string) string {
	if len(s) >= 2 && (s[0:2] == "0x" || s[0:2] == "0X") {
		return s[2:]
	}
	return s
}

// calculateInputSignatureHash calculates the signature hash for a specific input
// This should match the logic in verifier.vault package
func (sdk *SDK) calculateInputSignatureHash(pkt *psbt.Packet, inputIndex int) ([]byte, error) {
	// TODO: Implement signature hash calculation
	// This needs to match exactly what was sent to TSS
	// Should follow the same logic as verifier.vault package for building message hash key
	return nil, fmt.Errorf("signature hash calculation not implemented - need to match verifier.vault logic")
}

// extractPubkeyForInput extracts the public key for a PSBT input
func (sdk *SDK) extractPubkeyForInput(input *psbt.PInput) ([]byte, error) {
	// Look for public key in BIP32 derivation
	if len(input.Bip32Derivation) > 0 {
		for _, derivation := range input.Bip32Derivation {
			if len(derivation.PubKey) == 33 {
				return derivation.PubKey, nil
			}
		}
	}

	return nil, fmt.Errorf("no public key found in PSBT input")
}

// Broadcast submits signed transaction to Bitcoin network
func (sdk *SDK) Broadcast(signedTxBytes []byte) error {
	if sdk.rpcClient == nil {
		return fmt.Errorf("rpc client not configured")
	}

	var tx wire.MsgTx
	if err := tx.Deserialize(bytes.NewReader(signedTxBytes)); err != nil {
		return fmt.Errorf("deserialize signed tx: %w", err)
	}

	_, err := sdk.rpcClient.SendRawTransaction(&tx, false)
	if err != nil {
		return fmt.Errorf("sendrawtransaction failed: %w", err)
	}

	return nil
}

// Send is a convenience method that signs and broadcasts the transaction
func (sdk *SDK) Send(psbt []byte, signatures map[string]tss.KeysignResponse) error {
	// Sign the transaction
	signedTxBytes, err := sdk.Sign(psbt, signatures)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Broadcast the signed transaction
	err = sdk.Broadcast(signedTxBytes)
	if err != nil {
		return fmt.Errorf("failed to broadcast transaction: %w", err)
	}

	return nil
}
