package bsc

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

const bscEvmChainID = 56

// ParsedBscTransaction implements the vultisigTypes.DecodedTransaction interface for BSC.
type ParsedBscTransaction struct {
	tx      *types.Transaction
	sender  common.Address
	chainID *big.Int
}

// ChainIdentifier returns "bsc".
func (p *ParsedBscTransaction) ChainIdentifier() string { return "bsc" }

// Hash returns the transaction hash.
func (p *ParsedBscTransaction) Hash() string { return p.tx.Hash().Hex() }

// From returns the sender's address.
func (p *ParsedBscTransaction) From() string { return p.sender.Hex() }

// To returns the recipient's address or contract address.
func (p *ParsedBscTransaction) To() string {
	if p.tx.To() == nil {
		return "" // Contract creation
	}
	return p.tx.To().Hex()
}

// Value returns the amount of BNB transferred.
func (p *ParsedBscTransaction) Value() *big.Int { return p.tx.Value() }

// Data returns the transaction input data.
func (p *ParsedBscTransaction) Data() []byte { return p.tx.Data() }

// Nonce returns the transaction nonce.
func (p *ParsedBscTransaction) Nonce() uint64 { return p.tx.Nonce() }

// GasPrice returns the transaction gas price.
func (p *ParsedBscTransaction) GasPrice() *big.Int { return p.tx.GasPrice() }

// GasLimit returns the transaction gas limit.
func (p *ParsedBscTransaction) GasLimit() uint64 { return p.tx.Gas() }

// Bsc implements the Chain interface for the BNB Smart Chain blockchain
type Bsc struct {
	abiRegistry     map[string]*ethereum.ABI
	tokenList       *ethereum.TokenList
	genericERC20ABI *ethereum.ABI
}

// ID returns the unique identifier for the BSC chain
func (b *Bsc) ID() string {
	return "bsc"
}

// Name returns a human-readable name for the BSC chain
func (b *Bsc) Name() string {
	return "BNB Smart Chain"
}

// Description returns a detailed description of the BSC chain
func (b *Bsc) Description() string {
	return "BNB Smart Chain (BSC) is an EVM-compatible blockchain designed for fast, low-cost transactions."
}

// SupportedProtocols returns the list of protocol IDs supported by the BSC chain
func (b *Bsc) SupportedProtocols() []string {
	protocols := []string{"bnb"}
	if b.tokenList != nil {
		for _, token := range b.tokenList.Tokens {
			protocols = append(protocols, strings.ToLower(token.Symbol))
		}
	}
	return protocols
}

// ParseTransaction decodes a raw BSC transaction from its hex representation.
func (b *Bsc) ParseTransaction(txHex string) (vultisigTypes.DecodedTransaction, error) {
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

	return &ParsedBscTransaction{tx: tx, sender: zeroAddress, chainID: nil}, nil
}

// LoadABI loads an ABI definition and registers it with a given name
func (b *Bsc) LoadABI(name string, abiJSON []byte) error {
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
func (b *Bsc) LoadTokenList(tokenListJSON []byte) error {
	tokenList, err := ethereum.ParseTokenList(tokenListJSON)
	if err != nil {
		return err
	}
	b.tokenList = tokenList
	return nil
}

// GetABI returns the ABI for a given name (case-insensitive)
func (b *Bsc) GetABI(name string) (*ethereum.ABI, bool) {
	if b.abiRegistry == nil {
		return nil, false
	}
	abi, ok := b.abiRegistry[strings.ToLower(name)]
	return abi, ok
}

// GetToken returns the token for a given symbol (case-insensitive)
func (b *Bsc) GetToken(symbol string) (*ethereum.Token, bool) {
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

// GetProtocol returns a protocol handler for the given ID (e.g., "bnb" or a token symbol).
func (b *Bsc) GetProtocol(id string) (vultisigTypes.Protocol, error) {
	lowerID := strings.ToLower(id)
	if lowerID == "bnb" {
		return NewBNB(), nil
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
		return ethereum.NewABIProtocol(token.Symbol, token.Name, fmt.Sprintf("BEP20 Token: %s on %s", token.Name, token.Address), b.genericERC20ABI), nil
	}

	// Check if an ABI was specifically loaded for this ID
	if abi, ok := b.GetABI(lowerID); ok {
		return ethereum.NewABIProtocol(id, id, fmt.Sprintf("Custom ABI Protocol: %s", id), abi), nil
	}

	return nil, fmt.Errorf("protocol %q not found or not supported on BSC. Ensure token list and ABIs are loaded correctly", id)
}

func (b *Bsc) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	if len(sigs) != 1 {
		return "", fmt.Errorf("expected exactly one signature, got %d", len(sigs))
	}

	payloadDecoded, err := ethereum.DecodeUnsignedPayload(proposedTx)
	if err != nil {
		return "", fmt.Errorf("bsc.DecodeUnsignedPayload: %w", err)
	}

	var sig []byte
	sig = append(sig, common.FromHex(sigs[0].R)...)
	sig = append(sig, common.FromHex(sigs[0].S)...)
	sig = append(sig, common.FromHex(sigs[0].RecoveryID)...)

	tx, err := types.NewTx(payloadDecoded).WithSignature(types.LatestSignerForChainID(big.NewInt(bscEvmChainID)), sig)
	if err != nil {
		return "", fmt.Errorf("gethtypes.NewTx.WithSignature: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// NewBsc creates a new BSC chain instance
func NewBsc() vultisigTypes.Chain {
	bscChain := &Bsc{
		abiRegistry: make(map[string]*ethereum.ABI),
	}
	var err error
	bscChain.genericERC20ABI, err = ethereum.ParseABI([]byte(ethereum.GenericERC20ABIJson))
	if err != nil {
		panic(fmt.Sprintf("FATAL: Failed to parse internal generic ERC20 ABI: %v", err))
	}
	return bscChain
}

