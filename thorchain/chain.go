package thorchain

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/types"
)

// Thorchain implements the Chain interface for the Thorchain blockchain
type Thorchain struct {
	network string // mainnet, testnet, etc.
}

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
func (t *Thorchain) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	return ParseThorchainTransaction(txHex)
}

// GetProtocol returns a protocol handler for the given ID
func (t *Thorchain) GetProtocol(id string) (types.Protocol, error) {
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
func (t *Thorchain) ComputeTxHash(proposedTxHex string, sigs []tss.KeysignResponse) (string, error) {
	if len(sigs) == 0 {
		return "", fmt.Errorf("no signatures provided")
	}

	// Remove 0x prefix if present
	proposedTxHex = strings.TrimPrefix(proposedTxHex, "0x")

	// Decode the proposed transaction
	txBytes, err := hex.DecodeString(proposedTxHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode transaction hex: %w", err)
	}

	// Create signed transaction bytes by combining original transaction with signatures
	// This follows the Cosmos SDK pattern where signatures are appended to the transaction
	signedTxData := make([]byte, 0, len(txBytes)+len(sigs)*96) // Estimate signature size
	signedTxData = append(signedTxData, txBytes...)

	// Apply signatures to the transaction data
	for _, sig := range sigs {
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

		// Append signature data in standard format
		signedTxData = append(signedTxData, rBytes...)
		signedTxData = append(signedTxData, sBytes...)
		signedTxData = append(signedTxData, recoveryBytes...)
	}

	// Compute the final transaction hash using SHA256 (standard for Cosmos)
	hash := computeThorchainTxHash(signedTxData)

	return hash, nil
}

// GetNetwork returns the network configuration (mainnet, testnet, etc.)
func (t *Thorchain) GetNetwork() string {
	return t.network
}

// NewThorchain creates a new Thorchain instance
func NewThorchain() types.Chain {
	return &Thorchain{
		network: "mainnet", // Default to mainnet
	}
}

// NewThorchainWithNetwork creates a new Thorchain instance with specified network
func NewThorchainWithNetwork(network string) types.Chain {
	return &Thorchain{
		network: network,
	}
}

// ValidateNetwork validates the network parameter
func ValidateNetwork(network string) error {
	validNetworks := map[string]bool{
		"mainnet":  true,
		"testnet":  true,
		"stagenet": true,
		"chaosnet": true,
	}

	if !validNetworks[network] {
		return fmt.Errorf("invalid network: %s, must be one of: mainnet, testnet, stagenet, chaosnet", network)
	}

	return nil
}
