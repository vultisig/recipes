package base

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

const baseEvmChainID = 8453

// ParsedBaseTransaction implements the vultisigTypes.DecodedTransaction interface for Base.
type ParsedBaseTransaction struct {
	tx      *types.Transaction
	sender  common.Address
	chainID *big.Int
}

// ChainIdentifier returns "base".
func (p *ParsedBaseTransaction) ChainIdentifier() string { return "base" }

// Hash returns the transaction hash.
func (p *ParsedBaseTransaction) Hash() string { return p.tx.Hash().Hex() }

// From returns the sender's address.
func (p *ParsedBaseTransaction) From() string { return p.sender.Hex() }

// To returns the recipient's address or contract address.
func (p *ParsedBaseTransaction) To() string {
	if p.tx.To() == nil {
		return "" // Contract creation
	}
	return p.tx.To().Hex()
}

// Value returns the amount of ETH transferred.
func (p *ParsedBaseTransaction) Value() *big.Int { return p.tx.Value() }

// Data returns the transaction input data.
func (p *ParsedBaseTransaction) Data() []byte { return p.tx.Data() }

// Nonce returns the transaction nonce.
func (p *ParsedBaseTransaction) Nonce() uint64 { return p.tx.Nonce() }

// GasPrice returns the transaction gas price.
func (p *ParsedBaseTransaction) GasPrice() *big.Int { return p.tx.GasPrice() }

// GasLimit returns the transaction gas limit.
func (p *ParsedBaseTransaction) GasLimit() uint64 { return p.tx.Gas() }

// Base implements the Chain interface for the Base blockchain
type Base struct {
	abiRegistry     map[string]*ethereum.ABI
	tokenList       *ethereum.TokenList
	genericERC20ABI *ethereum.ABI
}

// ID returns the unique identifier for the Base chain
func (b *Base) ID() string {
	return "base"
}

// Name returns a human-readable name for the Base chain
func (b *Base) Name() string {
	return "Base"
}

// Description returns a detailed description of the Base chain
func (b *Base) Description() string {
	return "Base is a secure, low-cost, developer-friendly Ethereum L2 built on the OP Stack."
}

// SupportedProtocols returns the list of protocol IDs supported by the Base chain
func (b *Base) SupportedProtocols() []string {
	protocols := []string{"eth"}
	if b.tokenList != nil {
		for _, token := range b.tokenList.Tokens {
			protocols = append(protocols, strings.ToLower(token.Symbol))
		}
	}
	return protocols
}

// ParseTransaction decodes a raw Base transaction from its hex representation.
func (b *Base) ParseTransaction(txHex string) (vultisigTypes.DecodedTransaction, error) {
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

	return &ParsedBaseTransaction{tx: tx, sender: zeroAddress, chainID: nil}, nil
}

// LoadABI loads an ABI definition and registers it with a given name
func (b *Base) LoadABI(name string, abiJSON []byte) error {
	abiJson, err := ethereum.ParseABI(abiJSON)
	if err != nil {
		return err
	}
	if b.abiRegistry == nil {
		b.abiRegistry = make(map[string]*ethereum.ABI)
	}
	b.abiRegistry[strings.ToLower(name)] = abiJson
	return nil
}

// LoadTokenList loads a token list from JSON
func (b *Base) LoadTokenList(tokenListJSON []byte) error {
	tokenList, err := ethereum.ParseTokenList(tokenListJSON)
	if err != nil {
		return err
	}
	b.tokenList = tokenList
	return nil
}

// GetABI returns the ABI for a given name (case-insensitive)
func (b *Base) GetABI(name string) (*ethereum.ABI, bool) {
	if b.abiRegistry == nil {
		return nil, false
	}
	abi, ok := b.abiRegistry[strings.ToLower(name)]
	return abi, ok
}

// GetToken returns the token for a given symbol (case-insensitive)
func (b *Base) GetToken(symbol string) (*ethereum.Token, bool) {
	if b.tokenList == nil {
		return nil, false
	}
	lowerSymbol := strings.ToLower(symbol)
	for _, token := range b.tokenList.Tokens {
		if strings.ToLower(token.Symbol) == lowerSymbol {
			return &token, true
		}
	}
	return nil, false
}

// GetProtocol returns a protocol handler for the given ID (e.g., "eth" or a token symbol).
func (b *Base) GetProtocol(id string) (vultisigTypes.Protocol, error) {
	lowerID := strings.ToLower(id)
	if lowerID == "eth" {
		return NewBaseETH(), nil
	}

	// Check if it's a token symbol from the loaded token list
	if token, ok := b.GetToken(lowerID); ok {
		if b.genericERC20ABI == nil {
			var err error
			b.genericERC20ABI, err = ethereum.ParseABI([]byte(ethereum.GenericERC20ABIJson))
			if err != nil {
				return nil, fmt.Errorf("failed to parse internal generic ERC20 ABI: %w", err)
			}
		}
		return ethereum.NewABIProtocol(token.Symbol, token.Name, fmt.Sprintf("ERC20 Token: %s on %s", token.Name, token.Address), b.genericERC20ABI), nil
	}

	// Check if an ABI was specifically loaded for this ID
	if abi, ok := b.GetABI(lowerID); ok {
		return ethereum.NewABIProtocol(id, id, fmt.Sprintf("Custom ABI Protocol: %s", id), abi), nil
	}

	return nil, fmt.Errorf("protocol %q not found or not supported on Base. Ensure token list and ABIs are loaded correctly", id)
}

func (b *Base) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	if len(sigs) != 1 {
		return "", fmt.Errorf("expected exactly one signature, got %d", len(sigs))
	}

	payloadDecoded, err := ethereum.DecodeUnsignedPayload(proposedTx)
	if err != nil {
		return "", fmt.Errorf("base.DecodeUnsignedPayload: %w", err)
	}

	var sig []byte
	sig = append(sig, common.FromHex(sigs[0].R)...)
	sig = append(sig, common.FromHex(sigs[0].S)...)
	sig = append(sig, common.FromHex(sigs[0].RecoveryID)...)

	tx, err := types.NewTx(payloadDecoded).WithSignature(types.LatestSignerForChainID(big.NewInt(baseEvmChainID)), sig)
	if err != nil {
		return "", fmt.Errorf("gethtypes.NewTx.WithSignature: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// NewBase creates a new Base chain instance
func NewBase() vultisigTypes.Chain {
	baseChain := &Base{
		abiRegistry: make(map[string]*ethereum.ABI),
	}
	var err error
	baseChain.genericERC20ABI, err = ethereum.ParseABI([]byte(ethereum.GenericERC20ABIJson))
	if err != nil {
		panic(fmt.Sprintf("FATAL: Failed to parse internal generic ERC20 ABI: %v", err))
	}
	return baseChain
}

