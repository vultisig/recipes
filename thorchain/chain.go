package thorchain

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/vultisig/mobile-tss-lib/tss"
	vultisigtypes "github.com/vultisig/recipes/types"
)

// Thorchain implements the Chain interface for the Thorchain blockchain
type Thorchain struct{}

// ID returns the unique identifier for the Thorchain
func (t *Thorchain) ID() string {
	return "thorchain"
}

// Name returns a human-readable name for the Thorchain
func (t *Thorchain) Name() string {
	return "Thorchain"
}

// Description returns a detailed description of the Thorchain
func (t *Thorchain) Description() string {
	return "Thorchain is a decentralized cross-chain liquidity protocol that enables users to swap assets across different blockchains. It uses proof-of-stake consensus with TSS for cross-chain operations."
}

// SupportedProtocols returns the list of protocol IDs supported by the Thorchain
func (t *Thorchain) SupportedProtocols() []string {
	return []string{"rune", "tcy"}
}

// ParseTransaction decodes a raw Thorchain transaction from its hex representation
func (t *Thorchain) ParseTransaction(txHex string) (vultisigtypes.DecodedTransaction, error) {
	return ParseThorchainTransaction(txHex)
}

// GetProtocol returns a protocol handler for the given ID
func (t *Thorchain) GetProtocol(id string) (vultisigtypes.Protocol, error) {
	lowerID := strings.ToLower(id)
	switch lowerID {
	case "rune":
		return NewRUNE(), nil
	case "tcy":
		return NewTCY(), nil
	default:
		return nil, fmt.Errorf("protocol %q not found or not supported on Thorchain", id)
	}
}

// ComputeTxHash computes the transaction hash with signatures applied
// For Thorchain transactions, this follows Cosmos SDK transaction signing and hashing
func (t *Thorchain) ComputeTxHash(proposedTxHex []byte, sigs []tss.KeysignResponse) (string, error) {
	if len(sigs) == 0 {
		return "", fmt.Errorf("no signatures provided")
	}

	// Create codec for transaction handling
	cdc := createCosmosCodec()

	// Unmarshal the transaction into proper Cosmos SDK Tx structure
	var cosmosTx tx.Tx
	err := cdc.Unmarshal(proposedTxHex, &cosmosTx)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal transaction: %w", err)
	}

	// Prepare signatures for injection
	signatures := make([][]byte, len(sigs))
	for i, sig := range sigs {
		// Convert signature components to bytes
		rBytes, err := hex.DecodeString(sig.R)
		if err != nil {
			return "", fmt.Errorf("invalid signature R component: %w", err)
		}
		sBytes, err := hex.DecodeString(sig.S)
		if err != nil {
			return "", fmt.Errorf("invalid signature S component: %w", err)
		}
		recoveryBytes, err := hex.DecodeString(sig.RecoveryID)
		if err != nil {
			return "", fmt.Errorf("invalid recovery ID: %w", err)
		}

		// Combine signature components in standard format
		sigBytes := make([]byte, 0, len(rBytes)+len(sBytes)+len(recoveryBytes))
		sigBytes = append(sigBytes, rBytes...)
		sigBytes = append(sigBytes, sBytes...)
		sigBytes = append(sigBytes, recoveryBytes...)
		signatures[i] = sigBytes
	}

	// Inject signatures into the transaction's AuthInfo
	if cosmosTx.AuthInfo == nil {
		return "", fmt.Errorf("transaction AuthInfo is nil")
	}

	// Update signatures in the transaction (not AuthInfo)
	cosmosTx.Signatures = signatures

	// Marshal the signed transaction back to bytes
	signedTxBytes, err := cdc.Marshal(&cosmosTx)
	if err != nil {
		return "", fmt.Errorf("failed to marshal signed transaction: %w", err)
	}

	// Compute the final transaction hash using SHA256 (standard for Cosmos)
	hash := computeThorchainTxHash(signedTxBytes)

	return hash, nil
}

// NewThorchain creates a new Thorchain instance
func NewThorchain() vultisigtypes.Chain {
	return &Thorchain{}
}
