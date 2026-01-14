// Package evm provides a generic EVM chain implementation that can be used
// for any Ethereum-compatible blockchain.
package evm

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vultisig/mobile-tss-lib/tss"

	"github.com/vultisig/recipes/chain/evm/ethereum"
	vultisigTypes "github.com/vultisig/recipes/types"
)

// ChainConfig contains the configuration for an EVM chain.
type ChainConfig struct {
	// ID is the unique identifier for the chain (e.g., "ethereum", "bsc")
	ID string
	// Name is the human-readable name (e.g., "Ethereum", "BNB Smart Chain")
	Name string
	// Description is a brief description of the chain
	Description string
	// EVMChainID is the numeric chain ID used in EIP-155
	EVMChainID int64
	// NativeProtocol is the native token protocol ID (e.g., "eth", "bnb", "avax")
	NativeProtocol string
}

// Chain implements the types.Chain interface for any EVM-compatible blockchain.
type Chain struct {
	config          ChainConfig
	genericERC20ABI *ethereum.ABI
}

// NewChain creates a new EVM chain instance with the given configuration.
func NewChain(config ChainConfig) *Chain {
	chain := &Chain{config: config}
	// Pre-load the generic ERC20 ABI
	var err error
	chain.genericERC20ABI, err = ethereum.ParseABI([]byte(ethereum.GenericERC20ABIJson))
	if err != nil {
		panic(fmt.Sprintf("FATAL: Failed to parse internal generic ERC20 ABI: %v", err))
	}
	return chain
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
	return []string{c.config.NativeProtocol}
}

// ParsedEVMTransaction implements the vultisigTypes.DecodedTransaction interface for EVM chains.
type ParsedEVMTransaction struct {
	tx      *types.Transaction
	chainID string
}

// ChainIdentifier returns the chain identifier.
func (p *ParsedEVMTransaction) ChainIdentifier() string { return p.chainID }

// Hash returns the transaction hash.
func (p *ParsedEVMTransaction) Hash() string { return p.tx.Hash().Hex() }

// From returns the sender's address (empty for unsigned transactions).
func (p *ParsedEVMTransaction) From() string { return "" }

// To returns the recipient's address or contract address.
func (p *ParsedEVMTransaction) To() string {
	if p.tx.To() == nil {
		return "" // Contract creation
	}
	return p.tx.To().Hex()
}

// Value returns the amount of native currency transferred.
func (p *ParsedEVMTransaction) Value() *big.Int { return p.tx.Value() }

// Data returns the transaction input data.
func (p *ParsedEVMTransaction) Data() []byte { return p.tx.Data() }

// Nonce returns the transaction nonce.
func (p *ParsedEVMTransaction) Nonce() uint64 { return p.tx.Nonce() }

// GasPrice returns the transaction gas price.
func (p *ParsedEVMTransaction) GasPrice() *big.Int { return p.tx.GasPrice() }

// GasLimit returns the transaction gas limit.
func (p *ParsedEVMTransaction) GasLimit() uint64 { return p.tx.Gas() }

// ParseTransaction decodes a raw EVM transaction from its hex representation.
func (c *Chain) ParseTransaction(txHex string) (vultisigTypes.DecodedTransaction, error) {
	rawTxBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex transaction: %w", err)
	}

	txData, err := ethereum.DecodeUnsignedPayload(rawTxBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode unsigned payload: %w", err)
	}

	tx := types.NewTx(txData)
	return &ParsedEVMTransaction{tx: tx, chainID: c.config.ID}, nil
}

// GetProtocol returns a protocol handler for the given ID.
func (c *Chain) GetProtocol(id string) (vultisigTypes.Protocol, error) {
	lowerID := strings.ToLower(id)
	if lowerID == c.config.NativeProtocol {
		return ethereum.NewNativeEVMProtocol(
			c.config.ID,
			c.config.NativeProtocol,
			fmt.Sprintf("%s Native Token", c.config.Name),
			fmt.Sprintf("Native currency of %s", c.config.Name),
		), nil
	}
	return nil, fmt.Errorf("protocol %q not found on %s", id, c.config.Name)
}

func (c *Chain) ExtractTxBytes(txData string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(txData)
}

// ComputeTxHash computes the transaction hash from the proposed transaction and signatures.
func (c *Chain) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	if len(sigs) != 1 {
		return "", fmt.Errorf("expected exactly one signature, got %d", len(sigs))
	}

	payloadDecoded, err := ethereum.DecodeUnsignedPayload(proposedTx)
	if err != nil {
		return "", fmt.Errorf("%s.DecodeUnsignedPayload: %w", c.config.ID, err)
	}

	var sig []byte
	sig = append(sig, common.FromHex(sigs[0].R)...)
	sig = append(sig, common.FromHex(sigs[0].S)...)
	sig = append(sig, common.FromHex(sigs[0].RecoveryID)...)

	tx, err := types.NewTx(payloadDecoded).WithSignature(
		types.LatestSignerForChainID(big.NewInt(c.config.EVMChainID)),
		sig,
	)
	if err != nil {
		return "", fmt.Errorf("types.NewTx.WithSignature: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// Predefined EVM chain configurations
var (
	EthereumConfig = ChainConfig{
		ID:             "ethereum",
		Name:           "Ethereum",
		Description:    "Ethereum is a decentralized, open-source blockchain with smart contract functionality.",
		EVMChainID:     1,
		NativeProtocol: "eth",
	}

	ArbitrumConfig = ChainConfig{
		ID:             "arbitrum",
		Name:           "Arbitrum",
		Description:    "Arbitrum is a Layer 2 scaling solution for Ethereum using optimistic rollups.",
		EVMChainID:     42161,
		NativeProtocol: "eth",
	}

	AvalancheConfig = ChainConfig{
		ID:             "avalanche",
		Name:           "Avalanche",
		Description:    "Avalanche is a high-performance, scalable blockchain platform for decentralized applications.",
		EVMChainID:     43114,
		NativeProtocol: "avax",
	}

	BaseConfig = ChainConfig{
		ID:             "base",
		Name:           "Base",
		Description:    "Base is a secure, low-cost, developer-friendly Ethereum L2 built on the OP Stack.",
		EVMChainID:     8453,
		NativeProtocol: "eth",
	}

	BSCConfig = ChainConfig{
		ID:             "bsc",
		Name:           "BNB Smart Chain",
		Description:    "BNB Smart Chain (BSC) is an EVM-compatible blockchain designed for fast, low-cost transactions.",
		EVMChainID:     56,
		NativeProtocol: "bnb",
	}

	BlastConfig = ChainConfig{
		ID:             "blast",
		Name:           "Blast",
		Description:    "Blast is an Ethereum L2 with native yield for ETH and stablecoins.",
		EVMChainID:     81457,
		NativeProtocol: "eth",
	}

	CronosConfig = ChainConfig{
		ID:             "cronos",
		Name:           "Cronos",
		Description:    "Cronos is an EVM-compatible blockchain built by Crypto.com.",
		EVMChainID:     25,
		NativeProtocol: "cro",
	}

	OptimismConfig = ChainConfig{
		ID:             "optimism",
		Name:           "Optimism",
		Description:    "Optimism is a Layer 2 scaling solution for Ethereum using optimistic rollups.",
		EVMChainID:     10,
		NativeProtocol: "eth",
	}

	PolygonConfig = ChainConfig{
		ID:             "polygon",
		Name:           "Polygon",
		Description:    "Polygon is a scaling solution for Ethereum that provides faster and cheaper transactions.",
		EVMChainID:     137,
		NativeProtocol: "matic",
	}

	ZksyncConfig = ChainConfig{
		ID:             "zksync",
		Name:           "zkSync Era",
		Description:    "zkSync Era is a Layer 2 scaling solution for Ethereum using zero-knowledge proofs.",
		EVMChainID:     324,
		NativeProtocol: "eth",
	}
)

// AllEVMChainConfigs returns all predefined EVM chain configurations.
func AllEVMChainConfigs() []ChainConfig {
	return []ChainConfig{
		EthereumConfig,
		ArbitrumConfig,
		AvalancheConfig,
		BaseConfig,
		BSCConfig,
		BlastConfig,
		CronosConfig,
		OptimismConfig,
		PolygonConfig,
		ZksyncConfig,
	}
}

// NewChainFromConfig creates a new EVM chain instance that implements types.Chain.
func NewChainFromConfig(config ChainConfig) vultisigTypes.Chain {
	return NewChain(config)
}
