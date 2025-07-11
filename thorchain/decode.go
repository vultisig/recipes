package thorchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	vultisigTypes "github.com/vultisig/recipes/types"
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
	parsedMemo *ParsedMemo
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

func (p *ParsedThorchainTransaction) GetParsedMemo() *ParsedMemo {
	return p.parsedMemo
}

func (p *ParsedThorchainTransaction) GetMemoType() MemoType {
	if p.parsedMemo != nil {
		return p.parsedMemo.Type
	}
	return MemoTypeTransfer
}

// createCosmosCodec creates a codec for decoding Cosmos SDK transactions
func createCosmosCodec() codec.Codec {
	interfaceRegistry := types.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)

	return codec.NewProtoCodec(interfaceRegistry)
}

// ParseThorchainTransaction decodes a raw Thorchain transaction using Cosmos SDK protobuf parsing
func ParseThorchainTransaction(txHex string) (vultisigTypes.DecodedTransaction, error) {
	// Remove 0x prefix if present
	txHex = strings.TrimPrefix(txHex, "0x")

	// Validate hex format and decode to bytes
	rawTxBytes, err := hex.DecodeString(txHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex transaction: %w", err)
	}

	// Validate minimum transaction size
	if len(rawTxBytes) < 10 {
		return nil, fmt.Errorf("transaction too short: %d bytes (minimum ~10 bytes expected)", len(rawTxBytes))
	}

	// Generate transaction hash using Cosmos SDK standard
	txHash := computeThorchainTxHash(rawTxBytes)

	// Create codec for transaction decoding
	cdc := createCosmosCodec()

	// Decode the transaction using Cosmos SDK protobuf
	var cosmosTxRaw tx.TxRaw
	err = cdc.Unmarshal(rawTxBytes, &cosmosTxRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to decode protobuf transaction: %w", err)
	}

	// Decode the TxRaw into a proper Tx structure
	var cosmosSDKTx tx.Tx

	// Initialize and unmarshal TxBody
	cosmosSDKTx.Body = &tx.TxBody{}
	err = cdc.Unmarshal(cosmosTxRaw.BodyBytes, cosmosSDKTx.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode transaction body: %w", err)
	}

	// Initialize and unmarshal AuthInfo
	cosmosSDKTx.AuthInfo = &tx.AuthInfo{}
	err = cdc.Unmarshal(cosmosTxRaw.AuthInfoBytes, cosmosSDKTx.AuthInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to decode auth info: %w", err)
	}

	cosmosSDKTx.Signatures = cosmosTxRaw.Signatures

	// Extract transaction information from decoded Cosmos SDK transaction
	parsed := &ParsedThorchainTransaction{
		txHash:  txHash,
		rawData: rawTxBytes,
	}

	// Extract auth info (gas, fee, sequence, account number)
	if cosmosSDKTx.AuthInfo != nil {
		if cosmosSDKTx.AuthInfo.Fee != nil {
			parsed.gasLimit = cosmosSDKTx.AuthInfo.Fee.GasLimit

			// Calculate gas price with proper validation
			if len(cosmosSDKTx.AuthInfo.Fee.Amount) > 0 && parsed.gasLimit > 0 {
				feeCoin := cosmosSDKTx.AuthInfo.Fee.Amount[0]

				// Verify fee denomination matches native token (rune) or known tokens
				if feeCoin.Denom == "rune" || feeCoin.Denom == "tcy" {
					feeAmount := feeCoin.Amount.BigInt()
					gasLimit := big.NewInt(int64(parsed.gasLimit))

					// Calculate gas price (fee amount / gas limit)
					parsed.gasPrice = new(big.Int).Div(feeAmount, gasLimit)
				}
			}
		}

		if len(cosmosSDKTx.AuthInfo.SignerInfos) > 0 {
			parsed.sequence = cosmosSDKTx.AuthInfo.SignerInfos[0].Sequence
		}
	}

	// Extract message information
	if cosmosSDKTx.Body != nil {
		parsed.memo = cosmosSDKTx.Body.Memo
		parsed.parsedMemo = ParseThorchainMemo(cosmosSDKTx.Body.Memo)

		if len(cosmosSDKTx.Body.Messages) > 0 {
			// Handle the first message (most common case for transfers)
			msg := cosmosSDKTx.Body.Messages[0]

			// Decode the message based on type URL
			switch msg.TypeUrl {
			case "/cosmos.bank.v1beta1.MsgSend":
				parsed.msgType = "MsgSend"

				var bankMsg banktypes.MsgSend
				err = cdc.Unmarshal(msg.Value, &bankMsg)
				if err == nil {
					parsed.fromAddr = bankMsg.FromAddress
					parsed.toAddr = bankMsg.ToAddress

					if len(bankMsg.Amount) > 0 {
						coin := bankMsg.Amount[0]
						parsed.amount = coin.Amount.BigInt()
						parsed.denom = coin.Denom
					}
				}

			default:
				// Handle other message types or set as unknown
				parsed.msgType = extractMessageType(msg.TypeUrl)
			}
		}
	}

	return parsed, nil
}

// extractMessageType extracts a readable message type from the type URL
func extractMessageType(typeURL string) string {
	// Extract the message type from the full type URL
	// e.g., "/cosmos.bank.v1beta1.MsgSend" -> "MsgSend"
	parts := strings.Split(typeURL, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return "Unknown"
}

// ValidateThorchainTransactionHex is an alias for backward compatibility
func ValidateThorchainTransactionHex(txHex string) (vultisigTypes.DecodedTransaction, error) {
	return ParseThorchainTransaction(txHex)
}

// computeThorchainTxHash generates a transaction hash using SHA256 (like Cosmos)
func computeThorchainTxHash(txBytes []byte) string {
	hash := sha256.Sum256(txBytes)
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

// ValidateThorchainAddress validates a Thorchain address format and checksum
func ValidateThorchainAddress(address string) error {
	if address == "" {
		return fmt.Errorf("address cannot be empty")
	}

	// Check prefix
	if !strings.HasPrefix(address, ThorchainBech32Prefix) {
		return fmt.Errorf("address must start with '%s'", ThorchainBech32Prefix)
	}

	// Check length
	if len(address) < ThorchainMinAddrLength || len(address) > ThorchainMaxAddrLength {
		return fmt.Errorf("address length must be between %d and %d characters", ThorchainMinAddrLength, ThorchainMaxAddrLength)
	}

	// Use bech32 library to decode and validate checksum
	hrp, data, err := bech32.Decode(address)
	if err != nil {
		return fmt.Errorf("invalid bech32 address: %w", err)
	}

	// Verify the human-readable part matches expected prefix
	if hrp != ThorchainBech32Prefix {
		return fmt.Errorf("expected address prefix %s, got %s", ThorchainBech32Prefix, hrp)
	}

	// Verify decoded data is not empty
	if len(data) == 0 {
		return fmt.Errorf("address contains no data")
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
