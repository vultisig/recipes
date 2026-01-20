package tron

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/types"
)

// Chain implements the types.Chain interface for TRON.
type Chain struct{}

// NewChain creates a new TRON chain instance.
func NewChain() *Chain {
	return &Chain{}
}

// ID returns the unique identifier for the TRON chain.
func (c *Chain) ID() string {
	return "tron"
}

// Name returns a human-readable name for the TRON chain.
func (c *Chain) Name() string {
	return "TRON"
}

// Description returns a detailed description of the TRON chain.
func (c *Chain) Description() string {
	return "TRON is a decentralized blockchain platform for content sharing and dApps."
}

// SupportedProtocols returns the list of protocol IDs supported by TRON.
func (c *Chain) SupportedProtocols() []string {
	return []string{"trx"}
}

// ParsedTronTransaction implements the types.DecodedTransaction interface for TRON.
type ParsedTronTransaction struct {
	rawData     *TronRawData
	rawDataHex  string
	txID        string
}

// TronRawData represents the raw data portion of a TRON transaction
type TronRawData struct {
	Contract      []TronContract `json:"contract"`
	RefBlockBytes string         `json:"ref_block_bytes"`
	RefBlockHash  string         `json:"ref_block_hash"`
	Expiration    int64          `json:"expiration"`
	Timestamp     int64          `json:"timestamp"`
	FeeLimit      int64          `json:"fee_limit,omitempty"`
	Data          string         `json:"data,omitempty"` // Memo field
}

// TronContract represents a contract in a TRON transaction
type TronContract struct {
	Parameter TronParameter `json:"parameter"`
	Type      string        `json:"type"`
}

// TronParameter represents the parameter of a contract
type TronParameter struct {
	Value   TronValue `json:"value"`
	TypeUrl string    `json:"type_url"`
}

// TronValue represents the value of a contract parameter
type TronValue struct {
	Amount       int64  `json:"amount,omitempty"`
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address,omitempty"`
	Data         string `json:"data,omitempty"` // For smart contract calls
}

// ChainIdentifier returns "tron".
func (p *ParsedTronTransaction) ChainIdentifier() string {
	return "tron"
}

// Hash returns the transaction hash.
func (p *ParsedTronTransaction) Hash() string {
	return p.txID
}

// From returns the sender address.
func (p *ParsedTronTransaction) From() string {
	if p.rawData == nil || len(p.rawData.Contract) == 0 {
		return ""
	}
	return p.rawData.Contract[0].Parameter.Value.OwnerAddress
}

// To returns the recipient address.
func (p *ParsedTronTransaction) To() string {
	if p.rawData == nil || len(p.rawData.Contract) == 0 {
		return ""
	}
	return p.rawData.Contract[0].Parameter.Value.ToAddress
}

// Value returns the amount being transferred.
func (p *ParsedTronTransaction) Value() *big.Int {
	if p.rawData == nil || len(p.rawData.Contract) == 0 {
		return nil
	}
	return big.NewInt(p.rawData.Contract[0].Parameter.Value.Amount)
}

// Data returns the transaction data/memo.
func (p *ParsedTronTransaction) Data() []byte {
	if p.rawData == nil {
		return nil
	}
	if p.rawData.Data != "" {
		data, _ := hex.DecodeString(p.rawData.Data)
		return data
	}
	return nil
}

// Nonce returns 0 as TRON doesn't use nonces in the same way.
func (p *ParsedTronTransaction) Nonce() uint64 {
	return 0
}

// GasPrice returns nil as TRON uses a different fee model.
func (p *ParsedTronTransaction) GasPrice() *big.Int {
	return nil
}

// GasLimit returns the fee limit.
func (p *ParsedTronTransaction) GasLimit() uint64 {
	if p.rawData == nil {
		return 0
	}
	return uint64(p.rawData.FeeLimit)
}

// GetRawData returns the underlying TRON raw data.
func (p *ParsedTronTransaction) GetRawData() *TronRawData {
	return p.rawData
}

// GetMemo returns the transaction memo/data as string.
func (p *ParsedTronTransaction) GetMemo() string {
	if p.rawData == nil || p.rawData.Data == "" {
		return ""
	}
	data, err := hex.DecodeString(p.rawData.Data)
	if err != nil {
		return p.rawData.Data
	}
	return string(data)
}

// GetAmount returns the transaction amount in SUN (1 TRX = 1,000,000 SUN).
func (p *ParsedTronTransaction) GetAmount() *big.Int {
	if p.rawData == nil || len(p.rawData.Contract) == 0 {
		return big.NewInt(0)
	}
	return big.NewInt(p.rawData.Contract[0].Parameter.Value.Amount)
}

// ParseTransaction decodes a TRON transaction's raw_data from hex string.
// The input must be the protobuf-encoded raw_data structure (containing ref_block_bytes,
// ref_block_hash, expiration, contract, etc.), NOT the full Transaction message with signatures.
// The transaction hash is computed as SHA256 of this raw_data.
func (c *Chain) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	txBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex: %w", err)
	}

	return c.ParseTransactionBytes(txBytes)
}

// ParseTransactionBytes decodes a TRON transaction's raw_data from bytes.
// The input must be the protobuf-encoded raw_data structure (containing ref_block_bytes,
// ref_block_hash, expiration, contract, etc.), NOT the full Transaction message with signatures.
// This is a simplified implementation that handles basic contract types including TransferContract
// and TriggerSmartContract.
func (c *Chain) ParseTransactionBytes(txBytes []byte) (types.DecodedTransaction, error) {
	const maxTxBytes = 32 * 1024 // 32 KB
	if len(txBytes) > maxTxBytes {
		return nil, fmt.Errorf("transaction too large: %d bytes (max %d)", len(txBytes), maxTxBytes)
	}

	if len(txBytes) == 0 {
		return nil, fmt.Errorf("empty transaction data")
	}

	// Compute transaction ID (SHA256 of raw_data)
	txID := sha256.Sum256(txBytes)
	txIDHex := hex.EncodeToString(txID[:])

	// For now, we'll parse the protobuf manually for basic transfer contracts
	// A full implementation would use the TRON protobuf definitions
	rawData, err := parseProtobufRawData(txBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TRON transaction: %w", err)
	}

	return &ParsedTronTransaction{
		rawData:    rawData,
		rawDataHex: hex.EncodeToString(txBytes),
		txID:       txIDHex,
	}, nil
}

// parseProtobufRawData parses the protobuf-encoded raw_data of a TRON transaction.
// This is a simplified parser that handles basic TransferContract transactions.
func parseProtobufRawData(data []byte) (*TronRawData, error) {
	if len(data) < 10 {
		return nil, fmt.Errorf("data too short for TRON transaction")
	}

	rawData := &TronRawData{}

	// Parse protobuf fields
	pos := 0
	for pos < len(data) {
		// Read field tag as varint (supports field numbers > 15)
		tagVal, n := readVarint(data[pos:])
		if n == 0 {
			return nil, fmt.Errorf("failed to read tag at position %d", pos)
		}
		pos += n
		fieldNum := tagVal >> 3
		wireType := tagVal & 0x7

		var err error
		switch fieldNum {
		case 1: // ref_block_bytes (bytes)
			if wireType != 2 {
				return nil, fmt.Errorf("unexpected wire type for ref_block_bytes")
			}
			rawData.RefBlockBytes, pos, err = readBytesField(data, pos)
			if err != nil {
				return nil, fmt.Errorf("ref_block_bytes: %w", err)
			}

		case 4: // ref_block_hash (bytes)
			if wireType != 2 {
				return nil, fmt.Errorf("unexpected wire type for ref_block_hash")
			}
			rawData.RefBlockHash, pos, err = readBytesField(data, pos)
			if err != nil {
				return nil, fmt.Errorf("ref_block_hash: %w", err)
			}

		case 8: // expiration (int64)
			if wireType != 0 {
				return nil, fmt.Errorf("unexpected wire type for expiration")
			}
			val, n := readVarint(data[pos:])
			if n == 0 {
				return nil, fmt.Errorf("failed to read expiration")
			}
			pos += n
			rawData.Expiration = int64(val)

		case 10: // data (bytes) - memo
			if wireType != 2 {
				return nil, fmt.Errorf("unexpected wire type for data")
			}
			rawData.Data, pos, err = readBytesField(data, pos)
			if err != nil {
				return nil, fmt.Errorf("data: %w", err)
			}

		case 11: // contract (repeated message)
			if wireType != 2 {
				return nil, fmt.Errorf("unexpected wire type for contract")
			}
			length, n := readVarint(data[pos:])
			if n == 0 {
				return nil, fmt.Errorf("failed to read contract length")
			}
			pos += n
			if pos+int(length) > len(data) {
				return nil, fmt.Errorf("contract length exceeds data")
			}
			contract, err := parseContract(data[pos : pos+int(length)])
			if err != nil {
				return nil, fmt.Errorf("failed to parse contract: %w", err)
			}
			rawData.Contract = append(rawData.Contract, contract)
			pos += int(length)

		case 14: // timestamp (int64)
			if wireType != 0 {
				return nil, fmt.Errorf("unexpected wire type for timestamp")
			}
			val, n := readVarint(data[pos:])
			if n == 0 {
				return nil, fmt.Errorf("failed to read timestamp")
			}
			pos += n
			rawData.Timestamp = int64(val)

		case 18: // fee_limit (int64)
			if wireType != 0 {
				return nil, fmt.Errorf("unexpected wire type for fee_limit")
			}
			val, n := readVarint(data[pos:])
			if n == 0 {
				return nil, fmt.Errorf("failed to read fee_limit")
			}
			pos += n
			rawData.FeeLimit = int64(val)

		default:
			// Skip unknown fields based on wire type
			pos, err = skipField(data, pos, wireType)
			if err != nil {
				return nil, fmt.Errorf("failed to skip field %d: %w", fieldNum, err)
			}
		}
	}

	return rawData, nil
}

// readBytesField reads a length-delimited bytes field and returns hex-encoded string
func readBytesField(data []byte, pos int) (string, int, error) {
	length, n := readVarint(data[pos:])
	if n == 0 {
		return "", pos, fmt.Errorf("failed to read length")
	}
	pos += n
	if pos+int(length) > len(data) {
		return "", pos, fmt.Errorf("length %d exceeds data bounds", length)
	}
	result := hex.EncodeToString(data[pos : pos+int(length)])
	return result, pos + int(length), nil
}

// skipField skips over a protobuf field based on its wire type
func skipField(data []byte, pos int, wireType uint64) (int, error) {
	switch wireType {
	case 0: // Varint
		_, n := readVarint(data[pos:])
		if n == 0 {
			return pos, fmt.Errorf("failed to read varint")
		}
		return pos + n, nil
	case 1: // 64-bit (fixed64, sfixed64, double)
		if pos+8 > len(data) {
			return pos, fmt.Errorf("not enough data for 64-bit field")
		}
		return pos + 8, nil
	case 2: // Length-delimited (string, bytes, embedded messages)
		length, n := readVarint(data[pos:])
		if n == 0 {
			return pos, fmt.Errorf("failed to read length")
		}
		pos += n
		if pos+int(length) > len(data) {
			return pos, fmt.Errorf("length exceeds data")
		}
		return pos + int(length), nil
	case 5: // 32-bit (fixed32, sfixed32, float)
		if pos+4 > len(data) {
			return pos, fmt.Errorf("not enough data for 32-bit field")
		}
		return pos + 4, nil
	default:
		return pos, fmt.Errorf("unsupported wire type %d", wireType)
	}
}

// parseContract parses a single contract from protobuf
func parseContract(data []byte) (TronContract, error) {
	contract := TronContract{}
	pos := 0

	for pos < len(data) {
		tagVal, n := readVarint(data[pos:])
		if n == 0 {
			return contract, fmt.Errorf("failed to read tag at position %d", pos)
		}
		pos += n
		fieldNum := tagVal >> 3
		wireType := tagVal & 0x7

		var err error
		switch fieldNum {
		case 1: // type (enum)
			if wireType != 0 {
				return contract, fmt.Errorf("unexpected wire type for type")
			}
			val, n := readVarint(data[pos:])
			if n == 0 {
				return contract, fmt.Errorf("failed to read contract type")
			}
			pos += n
			contract.Type = contractTypeToString(int(val))

		case 2: // parameter (Any)
			if wireType != 2 {
				return contract, fmt.Errorf("unexpected wire type for parameter")
			}
			length, n := readVarint(data[pos:])
			if n == 0 {
				return contract, fmt.Errorf("failed to read parameter length")
			}
			pos += n
			if pos+int(length) > len(data) {
				return contract, fmt.Errorf("parameter length exceeds data")
			}
			param, err := parseParameter(data[pos : pos+int(length)])
			if err != nil {
				return contract, fmt.Errorf("failed to parse parameter: %w", err)
			}
			contract.Parameter = param
			pos += int(length)

		default:
			pos, err = skipField(data, pos, wireType)
			if err != nil {
				return contract, fmt.Errorf("failed to skip field %d: %w", fieldNum, err)
			}
		}
	}

	return contract, nil
}

// parseParameter parses the parameter Any type
func parseParameter(data []byte) (TronParameter, error) {
	param := TronParameter{}
	pos := 0
	var valueBytes []byte

	// First pass: extract type_url and raw value bytes
	for pos < len(data) {
		tagVal, n := readVarint(data[pos:])
		if n == 0 {
			return param, fmt.Errorf("failed to read tag at position %d", pos)
		}
		pos += n
		fieldNum := tagVal >> 3
		wireType := tagVal & 0x7

		var err error
		switch fieldNum {
		case 1: // type_url (string)
			if wireType != 2 {
				return param, fmt.Errorf("unexpected wire type for type_url")
			}
			length, n := readVarint(data[pos:])
			if n == 0 {
				return param, fmt.Errorf("failed to read type_url length")
			}
			pos += n
			if pos+int(length) > len(data) {
				return param, fmt.Errorf("type_url length exceeds data")
			}
			param.TypeUrl = string(data[pos : pos+int(length)])
			pos += int(length)

		case 2: // value (bytes) - store raw bytes for later parsing
			if wireType != 2 {
				return param, fmt.Errorf("unexpected wire type for value")
			}
			length, n := readVarint(data[pos:])
			if n == 0 {
				return param, fmt.Errorf("failed to read value length")
			}
			pos += n
			if pos+int(length) > len(data) {
				return param, fmt.Errorf("value length exceeds data")
			}
			valueBytes = data[pos : pos+int(length)]
			pos += int(length)

		default:
			pos, err = skipField(data, pos, wireType)
			if err != nil {
				return param, fmt.Errorf("failed to skip field %d: %w", fieldNum, err)
			}
		}
	}

	// Second pass: parse value based on type_url
	if valueBytes != nil {
		// Currently only TransferContract is supported (native TRX transfers)
		// Both send and swap operations use TransferContract
		if strings.HasSuffix(param.TypeUrl, "TransferContract") {
			value, err := parseTransferContractValue(valueBytes)
			if err != nil {
				return param, fmt.Errorf("failed to parse TransferContract value: %w", err)
			}
			param.Value = value
		} else {
			// For unsupported contract types, try to extract basic fields
			value, err := parseGenericContractValue(valueBytes)
			if err != nil {
				return param, fmt.Errorf("failed to parse value for %s: %w", param.TypeUrl, err)
			}
			param.Value = value
		}
	}

	return param, nil
}

// parseTransferContractValue parses the value of a TransferContract
func parseTransferContractValue(data []byte) (TronValue, error) {
	value := TronValue{}
	pos := 0

	for pos < len(data) {
		tagVal, n := readVarint(data[pos:])
		if n == 0 {
			return value, fmt.Errorf("failed to read tag at position %d", pos)
		}
		pos += n
		fieldNum := tagVal >> 3
		wireType := tagVal & 0x7

		var err error
		switch fieldNum {
		case 1: // owner_address (bytes)
			if wireType != 2 {
				return value, fmt.Errorf("unexpected wire type for owner_address")
			}
			length, n := readVarint(data[pos:])
			if n == 0 {
				return value, fmt.Errorf("failed to read owner_address length")
			}
			pos += n
			if pos+int(length) > len(data) {
				return value, fmt.Errorf("owner_address length exceeds data")
			}
			value.OwnerAddress = encodeAddress(data[pos : pos+int(length)])
			pos += int(length)

		case 2: // to_address (bytes)
			if wireType != 2 {
				return value, fmt.Errorf("unexpected wire type for to_address")
			}
			length, n := readVarint(data[pos:])
			if n == 0 {
				return value, fmt.Errorf("failed to read to_address length")
			}
			pos += n
			if pos+int(length) > len(data) {
				return value, fmt.Errorf("to_address length exceeds data")
			}
			value.ToAddress = encodeAddress(data[pos : pos+int(length)])
			pos += int(length)

		case 3: // amount (int64)
			if wireType != 0 {
				return value, fmt.Errorf("unexpected wire type for amount")
			}
			val, n := readVarint(data[pos:])
			if n == 0 {
				return value, fmt.Errorf("failed to read amount")
			}
			pos += n
			value.Amount = int64(val)

		default:
			pos, err = skipField(data, pos, wireType)
			if err != nil {
				return value, fmt.Errorf("failed to skip field %d: %w", fieldNum, err)
			}
		}
	}

	return value, nil
}

// parseGenericContractValue attempts to parse a generic contract value
// extracting common fields like owner_address
func parseGenericContractValue(data []byte) (TronValue, error) {
	value := TronValue{}
	pos := 0

	for pos < len(data) {
		tagVal, n := readVarint(data[pos:])
		if n == 0 {
			return value, fmt.Errorf("failed to read tag at position %d", pos)
		}
		pos += n
		fieldNum := tagVal >> 3
		wireType := tagVal & 0x7

		var err error
		switch fieldNum {
		case 1: // owner_address is typically field 1 in most contracts
			if wireType == 2 {
				length, n := readVarint(data[pos:])
				if n == 0 {
					return value, fmt.Errorf("failed to read owner_address length")
				}
				pos += n
				if pos+int(length) > len(data) {
					return value, fmt.Errorf("owner_address length exceeds data")
				}
				value.OwnerAddress = encodeAddress(data[pos : pos+int(length)])
				pos += int(length)
			} else {
				pos, err = skipField(data, pos, wireType)
				if err != nil {
					return value, fmt.Errorf("failed to skip field %d: %w", fieldNum, err)
				}
			}

		default:
			pos, err = skipField(data, pos, wireType)
			if err != nil {
				return value, fmt.Errorf("failed to skip field %d: %w", fieldNum, err)
			}
		}
	}

	return value, nil
}

// readVarint reads a varint from the byte slice and returns the value and bytes consumed.
// Returns (0, 0) if the data is empty or the varint is incomplete/invalid.
func readVarint(data []byte) (uint64, int) {
	if len(data) == 0 {
		return 0, 0
	}
	var result uint64
	var shift uint
	for i, b := range data {
		result |= uint64(b&0x7f) << shift
		if b&0x80 == 0 {
			return result, i + 1
		}
		shift += 7
		if shift >= 64 {
			// Overflow - varint is too large
			return 0, 0
		}
	}
	// Incomplete varint - didn't find terminating byte
	return 0, 0
}

// encodeAddress converts raw address bytes to TRON base58check address.
// TRON addresses are 21 bytes: 1 byte version (0x41 for mainnet) + 20 bytes address.
// The output is base58check encoded (input + 4-byte checksum from double SHA256).
func encodeAddress(data []byte) string {
	if len(data) != 21 {
		// Invalid address length, return hex as fallback
		return hex.EncodeToString(data)
	}
	// TRON uses the same base58check format as Bitcoin, but with version byte 0x41
	// The input already includes the version byte, so we use CheckEncode without adding another
	// Actually, btcsuite's CheckEncode adds version byte, so we need to use base58.Encode with checksum
	return tronBase58CheckEncode(data)
}

// tronBase58CheckEncode encodes bytes to TRON base58check format.
// Unlike Bitcoin's base58check, TRON expects the version byte to already be in the input.
func tronBase58CheckEncode(input []byte) string {
	// Double SHA256 for checksum (same as Bitcoin)
	firstHash := sha256.Sum256(input)
	secondHash := sha256.Sum256(firstHash[:])
	checksum := secondHash[:4]

	// Make a copy to avoid modifying the original slice's underlying array
	// when input is a sub-slice of a larger buffer
	result := make([]byte, len(input)+4)
	copy(result, input)
	copy(result[len(input):], checksum)
	return base58.Encode(result)
}

// contractTypeToString converts contract type enum to string
func contractTypeToString(t int) string {
	switch t {
	case 1:
		return "TransferContract"
	case 2:
		return "TransferAssetContract"
	case 31:
		return "TriggerSmartContract"
	default:
		return fmt.Sprintf("Unknown(%d)", t)
	}
}

// ComputeTxHash computes the transaction hash from the raw_data bytes.
// The proposedTx must be the protobuf-encoded raw_data structure, NOT the full Transaction message.
// TRON transaction ID is SHA256 of the raw_data.
func (c *Chain) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	hash := sha256.Sum256(proposedTx)
	return hex.EncodeToString(hash[:]), nil
}

// GetProtocol returns a protocol handler for the given ID.
func (c *Chain) GetProtocol(id string) (types.Protocol, error) {
	if id == "trx" {
		return NewTRX(), nil
	}
	return nil, fmt.Errorf("protocol %q not found on TRON", id)
}

