package zcash

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/vultisig/recipes/types"
)

// ZcashNetworkParams holds Zcash-specific network parameters
type ZcashNetworkParams struct {
	Name            string
	AddressPrefix   []byte // t1 for mainnet transparent
	ScriptPrefix    []byte // t3 for mainnet P2SH
	Bech32HRP       string // not used for transparent, but for future
}

var (
	// ZcashMainNetParams defines the network parameters for Zcash mainnet
	ZcashMainNetParams = &ZcashNetworkParams{
		Name:          "mainnet",
		AddressPrefix: []byte{0x1C, 0xB8}, // t1
		ScriptPrefix:  []byte{0x1C, 0xBD}, // t3
	}
)

// ParsedZcashTransaction implements the types.DecodedTransaction interface for Zcash.
type ParsedZcashTransaction struct {
	tx         *ZcashTransaction
	txHash     string
	rawData    []byte
	network    *ZcashNetworkParams
}

// ZcashTransaction represents a Zcash transparent transaction (v4 or v5)
type ZcashTransaction struct {
	Version         int32
	VersionGroupID  uint32
	Inputs          []*ZcashInput
	Outputs         []*ZcashOutput
	LockTime        uint32
	ExpiryHeight    uint32
	ValueBalance    int64  // for sapling (not used in transparent-only)
	ShieldedSpends  []byte // empty for transparent
	ShieldedOutputs []byte // empty for transparent
	JoinSplits      []byte // empty for transparent
	BindingSig      []byte // empty for transparent
}

// ZcashInput represents a transaction input
type ZcashInput struct {
	PreviousOutPoint wire.OutPoint
	SignatureScript  []byte
	Sequence         uint32
}

// ZcashOutput represents a transaction output
type ZcashOutput struct {
	Value    int64
	PkScript []byte
}

// ChainIdentifier returns "zcash".
func (p *ParsedZcashTransaction) ChainIdentifier() string { return "zcash" }

// Hash returns the transaction hash.
func (p *ParsedZcashTransaction) Hash() string { return p.txHash }

// From returns empty string as Zcash doesn't have a simple from address in unsigned transactions
func (p *ParsedZcashTransaction) From() string { return "" }

// To returns the first output address if it can be decoded, otherwise empty
func (p *ParsedZcashTransaction) To() string {
	if len(p.tx.Outputs) > 0 {
		addr, err := ExtractZcashAddress(p.tx.Outputs[0].PkScript, p.network)
		if err == nil {
			return addr
		}
	}
	return ""
}

// Value returns the value of the first output in zatoshis
func (p *ParsedZcashTransaction) Value() *big.Int {
	if len(p.tx.Outputs) > 0 {
		return new(big.Int).SetInt64(p.tx.Outputs[0].Value)
	}
	return big.NewInt(0)
}

// Data returns the raw transaction data
func (p *ParsedZcashTransaction) Data() []byte { return p.rawData }

// Nonce returns 0 as Zcash doesn't use nonces like Ethereum
func (p *ParsedZcashTransaction) Nonce() uint64 { return 0 }

// GasPrice returns 0 as Zcash doesn't use gas
func (p *ParsedZcashTransaction) GasPrice() *big.Int { return big.NewInt(0) }

// GasLimit returns 0 as Zcash doesn't use gas
func (p *ParsedZcashTransaction) GasLimit() uint64 { return 0 }

// GetTransaction returns the underlying Zcash transaction
func (p *ParsedZcashTransaction) GetTransaction() *ZcashTransaction {
	return p.tx
}

// ParseZcashTransaction decodes a raw Zcash transaction from its hex representation.
func ParseZcashTransaction(txHex string) (types.DecodedTransaction, error) {
	return ParseZcashTransactionWithNetwork(txHex, ZcashMainNetParams)
}

// ParseZcashTransactionWithNetwork decodes a raw Zcash transaction with a specific network.
func ParseZcashTransactionWithNetwork(txHex string, network *ZcashNetworkParams) (types.DecodedTransaction, error) {
	rawTxBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex transaction: %w", err)
	}

	tx, err := deserializeZcashTransaction(rawTxBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize zcash transaction: %w", err)
	}

	// Calculate transaction hash
	txHash, err := calculateZcashTxHash(rawTxBytes, tx.Version)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate transaction hash: %w", err)
	}

	return &ParsedZcashTransaction{
		tx:      tx,
		txHash:  txHash,
		rawData: rawTxBytes,
		network: network,
	}, nil
}

// deserializeZcashTransaction parses raw bytes into a ZcashTransaction
func deserializeZcashTransaction(data []byte) (*ZcashTransaction, error) {
	r := bytes.NewReader(data)
	tx := &ZcashTransaction{}

	// Read header (4 bytes for version, includes overwintered flag)
	var header uint32
	if err := binary.Read(r, binary.LittleEndian, &header); err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	// Check if overwintered (bit 31 set)
	isOverwintered := (header >> 31) == 1
	tx.Version = int32(header & 0x7FFFFFFF)

	if isOverwintered {
		// Read version group ID (4 bytes)
		if err := binary.Read(r, binary.LittleEndian, &tx.VersionGroupID); err != nil {
			return nil, fmt.Errorf("failed to read version group ID: %w", err)
		}
	}

	// Read inputs
	inputCount, err := readVarInt(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read input count: %w", err)
	}

	tx.Inputs = make([]*ZcashInput, inputCount)
	for i := uint64(0); i < inputCount; i++ {
		input, er := readZcashInput(r)
		if er != nil {
			return nil, fmt.Errorf("failed to read input %d: %w", i, er)
		}
		tx.Inputs[i] = input
	}

	// Read outputs
	outputCount, err := readVarInt(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read output count: %w", err)
	}

	tx.Outputs = make([]*ZcashOutput, outputCount)
	for i := uint64(0); i < outputCount; i++ {
		output, er := readZcashOutput(r)
		if er != nil {
			return nil, fmt.Errorf("failed to read output %d: %w", i, er)
		}
		tx.Outputs[i] = output
	}

	// Read lock time
	if err := binary.Read(r, binary.LittleEndian, &tx.LockTime); err != nil {
		return nil, fmt.Errorf("failed to read lock time: %w", err)
	}

	if isOverwintered {
		// Read expiry height
		if err := binary.Read(r, binary.LittleEndian, &tx.ExpiryHeight); err != nil {
			return nil, fmt.Errorf("failed to read expiry height: %w", err)
		}

		// For v4 transactions (Sapling), read additional fields
		if tx.Version >= 4 {
			// Read value balance (8 bytes)
			if err := binary.Read(r, binary.LittleEndian, &tx.ValueBalance); err != nil {
				return nil, fmt.Errorf("failed to read value balance: %w", err)
			}

			// Read shielded spends count (should be 0 for transparent)
			spendCount, er := readVarInt(r)
			if er != nil {
				return nil, fmt.Errorf("failed to read shielded spends count: %w", er)
			}
			if spendCount > 0 {
				return nil, fmt.Errorf("shielded spends not supported, got %d", spendCount)
			}

			// Read shielded outputs count (should be 0 for transparent)
			outputsCount, er := readVarInt(r)
			if er != nil {
				return nil, fmt.Errorf("failed to read shielded outputs count: %w", er)
			}
			if outputsCount > 0 {
				return nil, fmt.Errorf("shielded outputs not supported, got %d", outputsCount)
			}

			// For v4, there may also be joinsplits (should be 0 for transparent)
			if tx.VersionGroupID == 0x892F2085 { // Sapling version group
				joinSplitCount, er := readVarInt(r)
				if er != nil {
					return nil, fmt.Errorf("failed to read joinsplit count: %w", er)
				}
				if joinSplitCount > 0 {
					return nil, fmt.Errorf("joinsplits not supported, got %d", joinSplitCount)
				}
			}
		}

		// For v5 transactions (NU5), the format is slightly different
		if tx.Version == 5 {
			// Read consensus branch ID (4 bytes) - already read as VersionGroupID
			// v5 transactions have a different structure for orchard actions
			// For transparent-only, we skip these
		}
	}

	return tx, nil
}

func readZcashInput(r *bytes.Reader) (*ZcashInput, error) {
	input := &ZcashInput{}

	// Read previous output hash (32 bytes)
	var hash [32]byte
	if _, err := r.Read(hash[:]); err != nil {
		return nil, fmt.Errorf("failed to read prev hash: %w", err)
	}
	// Reverse hash for chainhash
	for i, j := 0, len(hash)-1; i < j; i, j = i+1, j-1 {
		hash[i], hash[j] = hash[j], hash[i]
	}
	prevHash, err := chainhash.NewHash(hash[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create prev hash: %w", err)
	}
	input.PreviousOutPoint.Hash = *prevHash

	// Read previous output index (4 bytes)
	if err := binary.Read(r, binary.LittleEndian, &input.PreviousOutPoint.Index); err != nil {
		return nil, fmt.Errorf("failed to read prev index: %w", err)
	}

	// Read signature script
	scriptLen, err := readVarInt(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read script length: %w", err)
	}
	input.SignatureScript = make([]byte, scriptLen)
	if _, err := r.Read(input.SignatureScript); err != nil {
		return nil, fmt.Errorf("failed to read signature script: %w", err)
	}

	// Read sequence (4 bytes)
	if err := binary.Read(r, binary.LittleEndian, &input.Sequence); err != nil {
		return nil, fmt.Errorf("failed to read sequence: %w", err)
	}

	return input, nil
}

func readZcashOutput(r *bytes.Reader) (*ZcashOutput, error) {
	output := &ZcashOutput{}

	// Read value (8 bytes)
	if err := binary.Read(r, binary.LittleEndian, &output.Value); err != nil {
		return nil, fmt.Errorf("failed to read value: %w", err)
	}

	// Read pk script
	scriptLen, err := readVarInt(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read script length: %w", err)
	}
	output.PkScript = make([]byte, scriptLen)
	if _, err := r.Read(output.PkScript); err != nil {
		return nil, fmt.Errorf("failed to read pk script: %w", err)
	}

	return output, nil
}

func readVarInt(r *bytes.Reader) (uint64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	switch {
	case b < 0xFD:
		return uint64(b), nil
	case b == 0xFD:
		var v uint16
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return 0, err
		}
		return uint64(v), nil
	case b == 0xFE:
		var v uint32
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return 0, err
		}
		return uint64(v), nil
	default:
		var v uint64
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return 0, err
		}
		return v, nil
	}
}

func calculateZcashTxHash(rawTx []byte, version int32) (string, error) {
	// For v4 transactions, hash is double SHA256 of the serialized tx
	// Note: Zcash v5 uses BLAKE2b for txid, but we don't support v5 yet
	if version >= 5 {
		return "", fmt.Errorf("v5 transaction hash calculation not implemented")
	}

	hash := chainhash.DoubleHashH(rawTx)
	return hash.String(), nil
}

// ExtractZcashAddress extracts the address from a Zcash output script
func ExtractZcashAddress(pkScript []byte, network *ZcashNetworkParams) (string, error) {
	// Check for P2PKH (OP_DUP OP_HASH160 <20 bytes> OP_EQUALVERIFY OP_CHECKSIG)
	if len(pkScript) == 25 && pkScript[0] == 0x76 && pkScript[1] == 0xa9 &&
		pkScript[2] == 0x14 && pkScript[23] == 0x88 && pkScript[24] == 0xac {
		// Extract the 20-byte hash
		hash := pkScript[3:23]
		return encodeZcashAddress(network.AddressPrefix, hash), nil
	}

	// Check for P2SH (OP_HASH160 <20 bytes> OP_EQUAL)
	if len(pkScript) == 23 && pkScript[0] == 0xa9 && pkScript[1] == 0x14 && pkScript[22] == 0x87 {
		hash := pkScript[2:22]
		return encodeZcashAddress(network.ScriptPrefix, hash), nil
	}

	// Check for OP_RETURN
	if len(pkScript) > 0 && pkScript[0] == 0x6a {
		return "", fmt.Errorf("OP_RETURN output has no address")
	}

	return "", fmt.Errorf("unsupported script type")
}

// encodeZcashAddress encodes a hash with prefix to a base58check address
func encodeZcashAddress(prefix []byte, hash []byte) string {
	// Zcash uses 2-byte prefix + 20-byte hash
	data := make([]byte, len(prefix)+len(hash))
	copy(data, prefix)
	copy(data[len(prefix):], hash)
	return base58CheckEncode(data)
}

// base58CheckEncode encodes data to base58check format
func base58CheckEncode(data []byte) string {
	// Add 4-byte checksum
	checksum := chainhash.DoubleHashB(data)[:4]
	
	// Create new slice to avoid mutating input
	payload := make([]byte, len(data)+4)
	copy(payload, data)
	copy(payload[len(data):], checksum)

	// Base58 encode
	return base58Encode(payload)
}

// base58Encode encodes bytes to base58
func base58Encode(data []byte) string {
	const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	// Count leading zeros
	var zeros int
	for _, b := range data {
		if b != 0 {
			break
		}
		zeros++
	}

	// Convert to big integer
	num := new(big.Int).SetBytes(data)
	base := big.NewInt(58)
	mod := new(big.Int)

	var result []byte
	for num.Sign() > 0 {
		num.DivMod(num, base, mod)
		result = append([]byte{alphabet[mod.Int64()]}, result...)
	}

	// Add leading 1s for zeros
	for i := 0; i < zeros; i++ {
		result = append([]byte{'1'}, result...)
	}

	return string(result)
}

// GetAllOutputs returns all outputs with their addresses and values
func (p *ParsedZcashTransaction) GetAllOutputs() []struct {
	Address string
	Value   int64
} {
	var outputs []struct {
		Address string
		Value   int64
	}

	for _, txOut := range p.tx.Outputs {
		output := struct {
			Address string
			Value   int64
		}{
			Value: txOut.Value,
		}

		// Try to extract address
		addr, err := ExtractZcashAddress(txOut.PkScript, p.network)
		if err == nil {
			output.Address = addr
		}

		outputs = append(outputs, output)
	}

	return outputs
}

// GetInputCount returns the number of inputs
func (p *ParsedZcashTransaction) GetInputCount() int {
	return len(p.tx.Inputs)
}

// GetOutputCount returns the number of outputs
func (p *ParsedZcashTransaction) GetOutputCount() int {
	return len(p.tx.Outputs)
}

// GetLockTime returns the transaction lock time
func (p *ParsedZcashTransaction) GetLockTime() uint32 {
	return p.tx.LockTime
}

// GetVersion returns the transaction version
func (p *ParsedZcashTransaction) GetVersion() int32 {
	return p.tx.Version
}

// GetExpiryHeight returns the transaction expiry height
func (p *ParsedZcashTransaction) GetExpiryHeight() uint32 {
	return p.tx.ExpiryHeight
}

