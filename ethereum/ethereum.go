package ethereum

import (
	"github.com/vultisig/recipes/types"
)

// Ethereum implements the Chain interface for the Ethereum blockchain
type Ethereum struct {
	// abiRegistry maintains a map of loaded ABIs
	abiRegistry map[string]*ABI
	// tokenList holds the loaded token list
	tokenList *TokenList
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
	// Start with the native ETH protocol
	protocols := []string{"eth"}

	// Add all loaded token protocols
	if e.tokenList != nil {
		for _, token := range e.tokenList.Tokens {
			protocols = append(protocols, token.Symbol)
		}
	}

	// Add all loaded ABI protocols
	for abiName := range e.abiRegistry {
		protocols = append(protocols, abiName)
	}

	return protocols
}

// LoadABI loads an ABI definition and registers it with a given name
func (e *Ethereum) LoadABI(name string, abiJSON []byte) error {
	abi, err := ParseABI(abiJSON)
	if err != nil {
		return err
	}

	// Register the ABI
	if e.abiRegistry == nil {
		e.abiRegistry = make(map[string]*ABI)
	}
	e.abiRegistry[name] = abi

	return nil
}

// LoadTokenList loads a token list from JSON
func (e *Ethereum) LoadTokenList(tokenListJSON []byte) error {
	tokenList, err := ParseTokenList(tokenListJSON)
	if err != nil {
		return err
	}

	e.tokenList = tokenList
	return nil
}

// GetABI returns the ABI for a given name
func (e *Ethereum) GetABI(name string) (*ABI, bool) {
	if e.abiRegistry == nil {
		return nil, false
	}

	abi, ok := e.abiRegistry[name]
	return abi, ok
}

// GetToken returns the token for a given symbol
func (e *Ethereum) GetToken(symbol string) (*Token, bool) {
	if e.tokenList == nil {
		return nil, false
	}

	for _, token := range e.tokenList.Tokens {
		if token.Symbol == symbol {
			return &token, true
		}
	}

	return nil, false
}

// NewEthereum creates a new Ethereum chain instance
func NewEthereum() types.Chain {
	return &Ethereum{
		abiRegistry: make(map[string]*ABI),
	}
}
