package arbitrum

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/vultisig/mobile-tss-lib/tss"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/vultisig/recipes/chain/evm/ethereum"
	vultisigTypes "github.com/vultisig/recipes/types"
)

const arbitrumEvmChainID = 42161

// ParsedArbitrumTransaction implements the vultisigTypes.DecodedTransaction interface for Arbitrum.
type ParsedArbitrumTransaction struct {
	tx      *types.Transaction
	sender  common.Address
	chainID *big.Int
}

// ChainIdentifier returns "arbitrum".
func (p *ParsedArbitrumTransaction) ChainIdentifier() string { return "arbitrum" }

// Hash returns the transaction hash.
func (p *ParsedArbitrumTransaction) Hash() string { return p.tx.Hash().Hex() }

// From returns the sender's address.
func (p *ParsedArbitrumTransaction) From() string { return p.sender.Hex() }

// To returns the recipient's address or contract address.
func (p *ParsedArbitrumTransaction) To() string {
	if p.tx.To() == nil {
		return "" // Contract creation
	}
	return p.tx.To().Hex()
}

// Value returns the amount of ETH transferred.
func (p *ParsedArbitrumTransaction) Value() *big.Int { return p.tx.Value() }

// Data returns the transaction input data.
func (p *ParsedArbitrumTransaction) Data() []byte { return p.tx.Data() }

// Nonce returns the transaction nonce.
func (p *ParsedArbitrumTransaction) Nonce() uint64 { return p.tx.Nonce() }

// GasPrice returns the transaction gas price.
func (p *ParsedArbitrumTransaction) GasPrice() *big.Int { return p.tx.GasPrice() }

// GasLimit returns the transaction gas limit.
func (p *ParsedArbitrumTransaction) GasLimit() uint64 { return p.tx.Gas() }

// Arbitrum implements the Chain interface for the Arbitrum blockchain
type Arbitrum struct {
	abiRegistry     map[string]*ethereum.ABI
	tokenList       *ethereum.TokenList
	genericERC20ABI *ethereum.ABI
}

// ID returns the unique identifier for the Arbitrum chain
func (a *Arbitrum) ID() string {
	return "arbitrum"
}

// Name returns a human-readable name for the Arbitrum chain
func (a *Arbitrum) Name() string {
	return "Arbitrum"
}

// Description returns a detailed description of the Arbitrum chain
func (a *Arbitrum) Description() string {
	return "Arbitrum is a Layer 2 scaling solution for Ethereum using optimistic rollups."
}

// SupportedProtocols returns the list of protocol IDs supported by the Arbitrum chain
func (a *Arbitrum) SupportedProtocols() []string {
	protocols := []string{"eth"}
	if a.tokenList != nil {
		for _, token := range a.tokenList.Tokens {
			protocols = append(protocols, strings.ToLower(token.Symbol))
		}
	}
	return protocols
}

// ParseTransaction decodes a raw Arbitrum transaction from its hex representation.
func (a *Arbitrum) ParseTransaction(txHex string) (vultisigTypes.DecodedTransaction, error) {
	rawTxBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex transaction: %w", err)
	}

	txData, err := ethereum.DecodeUnsignedPayload(rawTxBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode unsigned payload: %w", err)
	}

	tx := types.NewTx(txData)

	zeroAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")

	return &ParsedArbitrumTransaction{tx: tx, sender: zeroAddress, chainID: nil}, nil
}

// LoadABI loads an ABI definition and registers it with a given name
func (a *Arbitrum) LoadABI(name string, abiJSON []byte) error {
	abiJson, err := ethereum.ParseABI(abiJSON)
	if err != nil {
		return err
	}
	if a.abiRegistry == nil {
		a.abiRegistry = make(map[string]*ethereum.ABI)
	}
	a.abiRegistry[strings.ToLower(name)] = abiJson
	return nil
}

// LoadTokenList loads a token list from JSON
func (a *Arbitrum) LoadTokenList(tokenListJSON []byte) error {
	tokenList, err := ethereum.ParseTokenList(tokenListJSON)
	if err != nil {
		return err
	}
	a.tokenList = tokenList
	return nil
}

// GetABI returns the ABI for a given name (case-insensitive)
func (a *Arbitrum) GetABI(name string) (*ethereum.ABI, bool) {
	if a.abiRegistry == nil {
		return nil, false
	}
	abi, ok := a.abiRegistry[strings.ToLower(name)]
	return abi, ok
}

// GetToken returns the token for a given symbol (case-insensitive)
func (a *Arbitrum) GetToken(symbol string) (*ethereum.Token, bool) {
	if a.tokenList == nil {
		return nil, false
	}
	lowerSymbol := strings.ToLower(symbol)
	for _, token := range a.tokenList.Tokens {
		if strings.ToLower(token.Symbol) == lowerSymbol {
			return &token, true
		}
	}
	return nil, false
}

// GetProtocol returns a protocol handler for the given ID (e.g., "eth" or a token symbol).
func (a *Arbitrum) GetProtocol(id string) (vultisigTypes.Protocol, error) {
	lowerID := strings.ToLower(id)
	if lowerID == "eth" {
		return NewArbitrumETH(), nil
	}

	// Check if it's a token symbol from the loaded token list
	if token, ok := a.GetToken(lowerID); ok {
		if a.genericERC20ABI == nil {
			var err error
			a.genericERC20ABI, err = ethereum.ParseABI([]byte(ethereum.GenericERC20ABIJson))
			if err != nil {
				return nil, fmt.Errorf("failed to parse internal generic ERC20 ABI: %w", err)
			}
		}
		return ethereum.NewABIProtocol(token.Symbol, token.Name, fmt.Sprintf("ERC20 Token: %s on %s", token.Name, token.Address), a.genericERC20ABI), nil
	}

	// Check if an ABI was specifically loaded for this ID
	if abi, ok := a.GetABI(lowerID); ok {
		return ethereum.NewABIProtocol(id, id, fmt.Sprintf("Custom ABI Protocol: %s", id), abi), nil
	}

	return nil, fmt.Errorf("protocol %q not found or not supported on Arbitrum. Ensure token list and ABIs are loaded correctly", id)
}

func (a *Arbitrum) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	if len(sigs) != 1 {
		return "", fmt.Errorf("expected exactly one signature, got %d", len(sigs))
	}

	payloadDecoded, err := ethereum.DecodeUnsignedPayload(proposedTx)
	if err != nil {
		return "", fmt.Errorf("arbitrum.DecodeUnsignedPayload: %w", err)
	}

	var sig []byte
	sig = append(sig, common.FromHex(sigs[0].R)...)
	sig = append(sig, common.FromHex(sigs[0].S)...)
	sig = append(sig, common.FromHex(sigs[0].RecoveryID)...)

	tx, err := types.NewTx(payloadDecoded).WithSignature(types.LatestSignerForChainID(big.NewInt(arbitrumEvmChainID)), sig)
	if err != nil {
		return "", fmt.Errorf("gethtypes.NewTx.WithSignature: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// NewArbitrum creates a new Arbitrum chain instance
func NewArbitrum() vultisigTypes.Chain {
	arbitrumChain := &Arbitrum{
		abiRegistry: make(map[string]*ethereum.ABI),
	}
	var err error
	arbitrumChain.genericERC20ABI, err = ethereum.ParseABI([]byte(ethereum.GenericERC20ABIJson))
	if err != nil {
		panic(fmt.Sprintf("FATAL: Failed to parse internal generic ERC20 ABI: %v", err))
	}
	return arbitrumChain
}

