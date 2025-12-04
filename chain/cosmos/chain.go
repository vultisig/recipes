package cosmos

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/types"
)

// ChainConfig holds configuration for a Cosmos-based chain.
type ChainConfig struct {
	// ID is the unique chain identifier (e.g., "cosmos", "mayachain", "thorchain")
	ID string
	// Name is the human-readable chain name
	Name string
	// Description is a detailed description of the chain
	Description string
	// Bech32Prefix is the address prefix (e.g., "cosmos", "maya", "thor")
	Bech32Prefix string
	// Protocols is the list of supported protocol IDs
	Protocols []string
	// MessageTypeRegistry maps TypeUrls to MessageTypes
	MessageTypeRegistry *MessageTypeRegistry
	// RegisterExtraTypes is an optional function to register additional protobuf types
	RegisterExtraTypes func(ir codectypes.InterfaceRegistry)
	// GetProtocol returns a protocol handler for the given ID
	GetProtocol func(id string) (types.Protocol, error)
	// CustomFromExtractor is an optional custom function to extract From address
	CustomFromExtractor func(*tx.Tx, codec.Codec, string) string
	// CustomToExtractor is an optional custom function to extract To address
	CustomToExtractor func(*tx.Tx, codec.Codec) string
}

// Chain implements the types.Chain interface for Cosmos SDK-based blockchains.
type Chain struct {
	config ChainConfig
	cdc    codec.Codec
}

// NewChain creates a new Cosmos chain instance with the given configuration.
func NewChain(config ChainConfig) *Chain {
	ir := codectypes.NewInterfaceRegistry()

	// Register crypto types (required for PubKey interfaces)
	cryptocodec.RegisterInterfaces(ir)

	// Register bank message types
	banktypes.RegisterInterfaces(ir)

	// Register any extra types specific to this chain
	if config.RegisterExtraTypes != nil {
		config.RegisterExtraTypes(ir)
	}

	return &Chain{
		config: config,
		cdc:    codec.NewProtoCodec(ir),
	}
}

// ID returns the unique identifier for the chain.
func (c *Chain) ID() string {
	return c.config.ID
}

// Name returns a human-readable name for the chain.
func (c *Chain) Name() string {
	return c.config.Name
}

// Description returns a detailed description of the chain.
func (c *Chain) Description() string {
	return c.config.Description
}

// SupportedProtocols returns the list of protocol IDs supported by this chain.
func (c *Chain) SupportedProtocols() []string {
	return c.config.Protocols
}

// Codec returns the codec used by this chain.
func (c *Chain) Codec() codec.Codec {
	return c.cdc
}

// Config returns the chain configuration.
func (c *Chain) Config() ChainConfig {
	return c.config
}

// ParseTransaction decodes a raw Cosmos transaction from hex string.
func (c *Chain) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	txBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex: %w", err)
	}

	return c.ParseTransactionBytes(txBytes)
}

// ParseTransactionBytes decodes a raw Cosmos transaction from bytes.
func (c *Chain) ParseTransactionBytes(txBytes []byte) (types.DecodedTransaction, error) {
	const maxTxBytes = 32 * 1024 // 32 KB
	if len(txBytes) > maxTxBytes {
		return nil, fmt.Errorf("transaction too large: %d bytes (max %d)", len(txBytes), maxTxBytes)
	}

	if len(txBytes) == 0 {
		return nil, fmt.Errorf("empty transaction data")
	}

	var txData tx.Tx
	if err := c.cdc.Unmarshal(txBytes, &txData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal protobuf transaction: %w", err)
	}

	return &ParsedCosmosTransaction{
		Tx:            &txData,
		Cdc:           c.cdc,
		ChainID:       c.config.ID,
		Bech32Prefix:  c.config.Bech32Prefix,
		FromExtractor: c.config.CustomFromExtractor,
	}, nil
}

// ComputeTxHash computes the transaction hash from the proposed transaction and signatures.
func (c *Chain) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	if len(sigs) != 1 {
		return "", fmt.Errorf("expected exactly one signature, got %d", len(sigs))
	}

	// Unmarshal the unsigned transaction
	var unsignedTx tx.Tx
	if err := c.cdc.Unmarshal(proposedTx, &unsignedTx); err != nil {
		return "", fmt.Errorf("failed to unmarshal unsigned transaction: %w", err)
	}

	// Verify the transaction has AuthInfo with SignerInfos (containing public key)
	if unsignedTx.AuthInfo == nil || len(unsignedTx.AuthInfo.SignerInfos) == 0 {
		return "", fmt.Errorf("unsigned transaction missing AuthInfo or SignerInfos")
	}

	sig := sigs[0]

	// Decode R and S from hex strings
	rBytes, err := hex.DecodeString(CleanHex(sig.R))
	if err != nil {
		return "", fmt.Errorf("failed to decode R: %w", err)
	}
	sBytes, err := hex.DecodeString(CleanHex(sig.S))
	if err != nil {
		return "", fmt.Errorf("failed to decode S: %w", err)
	}

	if len(rBytes) != 32 {
		return "", fmt.Errorf("r must be 32 bytes, got %d", len(rBytes))
	}
	if len(sBytes) != 32 {
		return "", fmt.Errorf("s must be 32 bytes, got %d", len(sBytes))
	}

	// Normalize S to low-S form
	sLow, err := NormalizeLowS(sBytes)
	if err != nil {
		return "", fmt.Errorf("low-S normalization failed: %w", err)
	}

	// Create signature bytes (r || s format for secp256k1)
	sigBytes := make([]byte, 64)
	copy(sigBytes[:32], rBytes)
	copy(sigBytes[32:], sLow)

	// Set the signature on the transaction
	unsignedTx.Signatures = [][]byte{sigBytes}

	// Marshal the signed transaction
	signedTxBytes, err := c.cdc.Marshal(&unsignedTx)
	if err != nil {
		return "", fmt.Errorf("failed to marshal signed transaction: %w", err)
	}

	// Compute SHA256 hash
	hash := sha256.Sum256(signedTxBytes)
	return strings.ToUpper(hex.EncodeToString(hash[:])), nil
}

// GetProtocol returns a protocol handler for the given ID.
func (c *Chain) GetProtocol(id string) (types.Protocol, error) {
	if c.config.GetProtocol != nil {
		return c.config.GetProtocol(id)
	}
	return nil, fmt.Errorf("protocol %q not found on %s", id, c.config.Name)
}

// From returns the sender address from the first message.
func (p *ParsedCosmosTransaction) From() string {
	// Use custom extractor if provided
	if p.FromExtractor != nil {
		return p.FromExtractor(p.Tx, p.Cdc, p.Bech32Prefix)
	}

	// Default extraction logic for MsgSend
	if p.Tx.Body == nil || len(p.Tx.Body.Messages) == 0 {
		return ""
	}

	msg := p.Tx.Body.Messages[0]
	var sdkMsg sdk.Msg
	if err := p.Cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return ""
	}

	if msgSend, ok := sdkMsg.(*banktypes.MsgSend); ok {
		return msgSend.FromAddress
	}
	return ""
}

// To returns the recipient address from the first message.
func (p *ParsedCosmosTransaction) To() string {
	if p.Tx.Body == nil || len(p.Tx.Body.Messages) == 0 {
		return ""
	}

	msg := p.Tx.Body.Messages[0]
	var sdkMsg sdk.Msg
	if err := p.Cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return ""
	}

	if msgSend, ok := sdkMsg.(*banktypes.MsgSend); ok {
		return msgSend.ToAddress
	}
	return ""
}

// DefaultFromExtractorWithDeposit creates a From extractor that handles both MsgSend and MsgDeposit.
// This is used by chains like THORChain and MAYAChain that support MsgDeposit.
func DefaultFromExtractorWithDeposit(tx *tx.Tx, cdc codec.Codec, bech32Prefix string) string {
	if tx.Body == nil || len(tx.Body.Messages) == 0 {
		return ""
	}

	msg := tx.Body.Messages[0]
	var sdkMsg sdk.Msg
	if err := cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return ""
	}

	switch m := sdkMsg.(type) {
	case *banktypes.MsgSend:
		return m.FromAddress
	case *types.MsgDeposit:
		if len(m.Signer) > 0 {
			addr, err := bech32.ConvertAndEncode(bech32Prefix, m.Signer)
			if err != nil {
				return ""
			}
			return addr
		}
	}
	return ""
}

