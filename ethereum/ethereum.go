package ethereum

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	vultisigTypes "github.com/vultisig/recipes/types"
)

// GenericERC20ABI is a standard ABI for ERC20 transfer function.
// [{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]
const GenericERC20ABIJson = `[{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"}]`

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
	abiRegistry     map[string]*ABI // Changed from *ABI to *types.ABI
	tokenList       *TokenList
	genericERC20ABI *ABI // Changed from *ABI to *types.ABI
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
	if lowerID == "eth" {
		return NewETH(), nil
	}

	// Check if it's a token symbol from the loaded token list
	if token, ok := e.GetToken(lowerID); ok {
		// Use the generic ERC20 ABI for this token
		// The NewABIProtocol will use this ABI to understand 'transfer', 'decimals', etc.
		if e.genericERC20ABI == nil {
			// Attempt to load generic ERC20 ABI if not already loaded (e.g. by NewEthereum)
			var err error
			e.genericERC20ABI, err = ParseABI([]byte(GenericERC20ABIJson))
			if err != nil {
				return nil, fmt.Errorf("failed to parse internal generic ERC20 ABI: %w", err)
			}
		}
		// The protocol ID for ABIProtocol should be the token symbol for clarity.
		// The name and description can also come from the token.
		return NewABIProtocol(token.Symbol, token.Name, fmt.Sprintf("ERC20 Token: %s on %s", token.Name, token.Address), e.genericERC20ABI), nil
	}

	// Check if an ABI was specifically loaded for this ID
	if abi, ok := e.GetABI(lowerID); ok {
		// This is for contracts registered directly by name/address via LoadABI
		return NewABIProtocol(id, id, fmt.Sprintf("Custom ABI Protocol: %s", id), abi), nil
	}

	return nil, fmt.Errorf("protocol %q not found or not supported on Ethereum. Ensure token list and ABIs are loaded correctly", id)
}

// NewEthereum creates a new Ethereum chain instance
// It now also attempts to preload the generic ERC20 ABI.
func NewEthereum() vultisigTypes.Chain {
	ethChain := &Ethereum{
		abiRegistry: make(map[string]*ABI), // Changed to *types.ABI
	}
	// Pre-load the generic ERC20 ABI
	var err error
	// This ParseABI must return *types.ABI
	ethChain.genericERC20ABI, err = ParseABI([]byte(GenericERC20ABIJson))
	if err != nil {
		// This is a critical internal ABI; panic or log an error.
		// For a CLI tool, panicking might be okay if it's considered essential.
		panic(fmt.Sprintf("FATAL: Failed to parse internal generic ERC20 ABI: %v", err))
	}
	return ethChain
}

// ... (rest of the file, if any, like helper functions for ABI/TokenList parsing if they were here)
// Ensure ParseABI and ParseTokenList functions are defined in this package or accessible.
// For this example, I'm assuming they are defined elsewhere in the 'ethereum' package.
// If they are not, their definitions would be needed.
// Based on the initially provided files, ParseABI and ParseTokenList are not shown but are used.
// I will assume they exist in other files of the ethereum package.
