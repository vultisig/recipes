package ethereum

import (
	"encoding/json"
	"fmt"
)

// Token represents a token in a token list
type Token struct {
	ChainId  int      `json:"chainId"`
	Address  string   `json:"address"`
	Name     string   `json:"name"`
	Symbol   string   `json:"symbol"`
	Decimals int      `json:"decimals"`
	LogoURI  string   `json:"logoURI,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

// TokenList represents a token list according to the Uniswap TokenList standard
type TokenList struct {
	Name      string   `json:"name"`
	LogoURI   string   `json:"logoURI,omitempty"`
	Keywords  []string `json:"keywords,omitempty"`
	Timestamp string   `json:"timestamp"`
	Version   struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
	Tokens []Token `json:"tokens"`
}

// ParseTokenList parses a token list JSON into a TokenList struct
func ParseTokenList(tokenListJSON []byte) (*TokenList, error) {
	var tokenList TokenList
	if err := json.Unmarshal(tokenListJSON, &tokenList); err != nil {
		return nil, fmt.Errorf("error unmarshaling token list: %w", err)
	}

	return &tokenList, nil
}

// GetTokenBySymbol returns a token by its symbol
func (t *TokenList) GetTokenBySymbol(symbol string) (*Token, bool) {
	for i, token := range t.Tokens {
		if token.Symbol == symbol {
			return &t.Tokens[i], true
		}
	}
	return nil, false
}

// GetTokenByAddress returns a token by its address
func (t *TokenList) GetTokenByAddress(address string) (*Token, bool) {
	for i, token := range t.Tokens {
		if token.Address == address {
			return &t.Tokens[i], true
		}
	}
	return nil, false
}
