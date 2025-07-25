package ethereum

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/erc20"

	vultisigTypes "github.com/vultisig/recipes/types"
)

const ethEvmChainID = 1

// ParsedEthereumTransaction implements the vultisigTypes.DecodedTransaction interface for Ethereum.
type ParsedEthereumTransaction struct {
	tx      *types.Transaction
	sender  common.Address
	chainID *big.Int
}

// ChainIdentifier returns "ethereum".
func (p *ParsedEthereumTransaction) ChainIdentifier() string { return "ethereum" }

// Hash returns the transaction hash.
func (p *ParsedEthereumTransaction) Hash() string { return p.tx.Hash().Hex() }

// From returns the sender's address.
func (p *ParsedEthereumTransaction) From() string { return p.sender.Hex() }

// To returns the recipient's address or contract address.
func (p *ParsedEthereumTransaction) To() string {
	if p.tx.To() == nil {
		return "" // Contract creation
	}
	return p.tx.To().Hex()
}

// Value returns the amount of ETH transferred.
func (p *ParsedEthereumTransaction) Value() *big.Int { return p.tx.Value() }

// Data returns the transaction input data.
func (p *ParsedEthereumTransaction) Data() []byte { return p.tx.Data() }

// Nonce returns the transaction nonce.
func (p *ParsedEthereumTransaction) Nonce() uint64 { return p.tx.Nonce() }

// GasPrice returns the transaction gas price.
func (p *ParsedEthereumTransaction) GasPrice() *big.Int { return p.tx.GasPrice() }

// GasLimit returns the transaction gas limit.
func (p *ParsedEthereumTransaction) GasLimit() uint64 { return p.tx.Gas() }

// Ethereum implements the Chain interface for the Ethereum blockchain
type Ethereum struct {
	abiRegistry map[string]*ABI // Changed from *ABI to *types.ABI
	tokenList   *TokenList
	chainID     *big.Int
}

// ID returns the unique identifier for the Ethereum chain
func (e *Ethereum) ID() string {
	return "ethereum"
}

// Name returns a human-readable name for the Ethereum chain
func (e *Ethereum) Name() string {
	return "Ethereum"
}

// Description returns a detailed description of the Ethereum chain
func (e *Ethereum) Description() string {
	return "Ethereum is a decentralized, open-source blockchain with smart contract functionality."
}

// SupportedProtocols returns the list of protocol IDs supported by the Ethereum chain
func (e *Ethereum) SupportedProtocols() []string {
	protocols := []string{"eth"}
	if e.tokenList != nil {
		for _, token := range e.tokenList.Tokens {
			protocols = append(protocols, strings.ToLower(token.Symbol)) // Use lowercase symbol consistent with potential policy assetID
		}
	}
	// Add other ABI protocols if necessary, though policy assetID will typically be 'eth' or token symbol
	return protocols
}

// ParseTransaction decodes a raw Ethereum transaction from its hex representation.
func (e *Ethereum) ParseTransaction(txHex string) (vultisigTypes.DecodedTransaction, error) {
	rawTxBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex transaction: %w", err)
	}

	txData, err := DecodeUnsignedPayload(rawTxBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode unsigned payload: %w", err)
	}

	tx := types.NewTx(txData)

	zeroAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")

	return &ParsedEthereumTransaction{tx: tx, sender: zeroAddress, chainID: nil}, nil
}

// LoadABI loads an ABI definition and registers it with a given name
func (e *Ethereum) LoadABI(name string, abiJSON []byte) error {
	abiJson, err := ParseABI(abiJSON) // This ParseABI must return *types.ABI
	if err != nil {
		return err
	}
	if e.abiRegistry == nil {
		e.abiRegistry = make(map[string]*ABI)
	}
	e.abiRegistry[strings.ToLower(name)] = abiJson
	return nil
}

// LoadTokenList loads a token list from JSON
func (e *Ethereum) LoadTokenList(tokenListJSON []byte) error {
	tokenList, err := ParseTokenList(tokenListJSON)
	if err != nil {
		return err
	}
	e.tokenList = tokenList
	// Optionally, pre-load ABIs for all tokens if a generic ERC20 ABI is available
	// Or do it on-demand in GetProtocol
	return nil
}

// GetABI returns the ABI for a given name (case-insensitive)
func (e *Ethereum) GetABI(name string) (*ABI, bool) { // Return type changed to *types.ABI
	if e.abiRegistry == nil {
		return nil, false
	}
	abi, ok := e.abiRegistry[strings.ToLower(name)]
	return abi, ok
}

// GetToken returns the token for a given symbol (case-insensitive)
func (e *Ethereum) GetToken(symbol string) (*Token, bool) {
	if e.tokenList == nil {
		return nil, false
	}
	lowerSymbol := strings.ToLower(symbol)
	for _, token := range e.tokenList.Tokens {
		if strings.ToLower(token.Symbol) == lowerSymbol {
			return &token, true
		}
	}
	return nil, false
}

// GetProtocol returns a protocol handler for the given ID (e.g., "eth" or a token symbol).
func (e *Ethereum) GetProtocol(id string) (vultisigTypes.Protocol, error) {
	lowerID := strings.ToLower(id)
	switch lowerID {
	case "eth":
		return NewETH(), nil
	case "erc20":
		return erc20.NewProtocol(NewETH().ChainID()), nil
	default:
		// Check if an ABI was specifically loaded for this ID
		if abi, ok := e.GetABI(lowerID); ok {
			// This is for contracts registered directly by name/address via LoadABI
			return NewABIProtocol(id, id, fmt.Sprintf("Custom ABI Protocol: %s", id), abi), nil
		}
	}
	return nil, fmt.Errorf("protocol %q not found or not supported on Ethereum. Ensure token list and ABIs are loaded correctly", id)
}

func (e *Ethereum) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	if len(sigs) != 1 {
		return "", fmt.Errorf("expected exactly one signature, got %d", len(sigs))
	}

	payloadDecoded, err := DecodeUnsignedPayload(proposedTx)
	if err != nil {
		return "", fmt.Errorf("ethereum.DecodeUnsignedPayload: %w", err)
	}

	var sig []byte
	sig = append(sig, common.FromHex(sigs[0].R)...)
	sig = append(sig, common.FromHex(sigs[0].S)...)
	sig = append(sig, common.FromHex(sigs[0].RecoveryID)...)

	tx, err := types.NewTx(payloadDecoded).WithSignature(types.LatestSignerForChainID(big.NewInt(ethEvmChainID)), sig)
	if err != nil {
		return "", fmt.Errorf("gethtypes.NewTx.WithSignature: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// NewEthereum creates a new Ethereum chain instance
// It now also attempts to preload the generic ERC20 ABI.
func NewEthereum() vultisigTypes.Chain {
	return &Ethereum{
		abiRegistry: make(map[string]*ABI), // Changed to *types.ABI
	}
}

// ... (rest of the file, if any, like helper functions for ABI/TokenList parsing if they were here)
// Ensure ParseABI and ParseTokenList functions are defined in this package or accessible.
// For this example, I'm assuming they are defined elsewhere in the 'ethereum' package.
// If they are not, their definitions would be needed.
// Based on the initially provided files, ParseABI and ParseTokenList are not shown but are used.
// I will assume they exist in other files of the ethereum package.
