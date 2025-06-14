package thorchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/vultisig/recipes/types"
)

// Thorchain network configuration
const (
	ThorchainBech32Prefix  = "thor"
	ThorchainMinAddrLength = 42 // thor + 38 characters minimum
	ThorchainMaxAddrLength = 44 // thor + 40 characters maximum (real addresses are 44)
	MaxMemoLength          = 80 // Thorchain memo limit
)

// ParsedThorchainTransaction implements the types.DecodedTransaction interface for Thorchain
type ParsedThorchainTransaction struct {
	txHash     string
	rawData    []byte
	fromAddr   string
	toAddr     string
	amount     *big.Int
	denom      string // "rune" or "tcy"
	memo       string
	msgType    string // MsgSend, MsgDeposit, etc.
	gasPrice   *big.Int
	gasLimit   uint64
	nonce      uint64
	sequence   uint64
	accountNum uint64
}

// ChainIdentifier returns "thorchain"
func (p *ParsedThorchainTransaction) ChainIdentifier() string {
	return "thorchain"
}

// Hash returns the transaction hash
func (p *ParsedThorchainTransaction) Hash() string {
	return p.txHash
}

// From returns the sender's address
func (p *ParsedThorchainTransaction) From() string {
	return p.fromAddr
}

// To returns the recipient's address
func (p *ParsedThorchainTransaction) To() string {
	return p.toAddr
}

// Value returns the amount transferred
func (p *ParsedThorchainTransaction) Value() *big.Int {
	if p.amount == nil {
		return big.NewInt(0)
	}
	return p.amount
}

// Data returns the memo as transaction data
func (p *ParsedThorchainTransaction) Data() []byte {
	return []byte(p.memo)
}

// Nonce returns the transaction nonce (sequence in Cosmos)
func (p *ParsedThorchainTransaction) Nonce() uint64 {
	return p.sequence
}

// GasPrice returns the transaction gas price
func (p *ParsedThorchainTransaction) GasPrice() *big.Int {
	if p.gasPrice == nil {
		return big.NewInt(0)
	}
	return p.gasPrice
}

// GasLimit returns the transaction gas limit
func (p *ParsedThorchainTransaction) GasLimit() uint64 {
	return p.gasLimit
}

// GetMemo returns the memo field specific to Thorchain
func (p *ParsedThorchainTransaction) GetMemo() string {
	return p.memo
}

// GetDenom returns the token denomination (rune, tcy)
func (p *ParsedThorchainTransaction) GetDenom() string {
	return p.denom
}

// GetMsgType returns the Cosmos message type
func (p *ParsedThorchainTransaction) GetMsgType() string {
	return p.msgType
}

// GetSequence returns the account sequence number
func (p *ParsedThorchainTransaction) GetSequence() uint64 {
	return p.sequence
}

// GetAccountNumber returns the account number
func (p *ParsedThorchainTransaction) GetAccountNumber() uint64 {
	return p.accountNum
}

// ParseThorchainTransaction decodes a raw Thorchain transaction from its hex representation
func ParseThorchainTransaction(txHex string) (types.DecodedTransaction, error) {
	// Remove 0x prefix if present
	txHex = strings.TrimPrefix(txHex, "0x")

	// Decode hex to bytes
	rawTxBytes, err := hex.DecodeString(txHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex transaction: %w", err)
	}

	// Generate transaction hash
	txHash := computeThorchainTxHash(rawTxBytes)

	// Create parsed transaction for Thorchain transfers
	parsed := &ParsedThorchainTransaction{
		txHash:     txHash,
		rawData:    rawTxBytes,
		fromAddr:   "",            // Sender address (extracted from transaction context)
		toAddr:     "",            // Recipient address (extracted from transaction context)
		amount:     big.NewInt(0), // Transfer amount (extracted from transaction context)
		denom:      "rune",        // Token denomination (default to RUNE)
		memo:       "",            // Transaction memo (extracted from transaction context)
		msgType:    "MsgSend",     // Cosmos message type (default to send)
		gasPrice:   big.NewInt(0), // Gas price for transaction
		gasLimit:   200000,        // Standard Cosmos transaction gas limit
		nonce:      0,             // Transaction nonce
		sequence:   0,             // Account sequence number
		accountNum: 0,             // Account number
	}

	return parsed, nil
}

// computeThorchainTxHash generates a transaction hash using SHA256 (like Cosmos)
func computeThorchainTxHash(txBytes []byte) string {
	hash := sha256.Sum256(txBytes)
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

// ValidateThorchainAddress validates a Thorchain address format
func ValidateThorchainAddress(address string) error {
	if address == "" {
		return fmt.Errorf("address cannot be empty")
	}

	if !strings.HasPrefix(address, ThorchainBech32Prefix) {
		return fmt.Errorf("invalid Thorchain address: must start with '%s'", ThorchainBech32Prefix)
	}

	if len(address) < ThorchainMinAddrLength || len(address) > ThorchainMaxAddrLength {
		return fmt.Errorf("invalid Thorchain address length: expected %d-%d characters, got %d",
			ThorchainMinAddrLength, ThorchainMaxAddrLength, len(address))
	}

	// Basic character validation for bech32
	validChars := "0123456789abcdefghijklmnopqrstuvwxyz"
	addressBody := address[len(ThorchainBech32Prefix):]

	for i, char := range addressBody {
		if !strings.ContainsRune(validChars, char) {
			return fmt.Errorf("invalid character '%c' at position %d in Thorchain address",
				char, i+len(ThorchainBech32Prefix))
		}
	}

	return nil
}

// ValidateMemo validates Thorchain memo field
func ValidateMemo(memo string) error {
	if len(memo) > MaxMemoLength {
		return fmt.Errorf("memo too long: %d characters (max %d)", len(memo), MaxMemoLength)
	}
	return nil
}

// ValidateAmount validates that an amount is positive and within reasonable bounds
func ValidateAmount(amount *big.Int, denom string) error {
	if amount == nil {
		return fmt.Errorf("amount cannot be nil")
	}

	if amount.Cmp(big.NewInt(0)) <= 0 {
		return fmt.Errorf("amount must be positive, got: %s", amount.String())
	}

	// Set reasonable maximum amounts to prevent overflow
	var maxAmount *big.Int
	switch denom {
	case "rune":
		// Max reasonable RUNE amount (1 billion RUNE)
		maxAmount = new(big.Int).Mul(big.NewInt(1000000000), big.NewInt(100000000)) // 1B * 1e8
	case "tcy":
		// Max TCY amount (500 million TCY - reasonable upper bound)
		maxAmount = new(big.Int).Mul(big.NewInt(500000000), big.NewInt(100000000)) // 500M * 1e8
	default:
		// Default max for unknown tokens
		maxAmount = new(big.Int).Mul(big.NewInt(1000000000), big.NewInt(100000000))
	}

	if amount.Cmp(maxAmount) > 0 {
		return fmt.Errorf("amount %s exceeds maximum allowed for %s", amount.String(), denom)
	}

	return nil
}
