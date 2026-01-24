package tron

import (
	cryptoSha256 "crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	chaintron "github.com/vultisig/recipes/chain/tron"
	stdcompare "github.com/vultisig/recipes/engine/compare"
	"github.com/vultisig/recipes/resolver"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
)

// Tron represents the TRON engine implementation
type Tron struct {
	chain *chaintron.Chain
}

// NewTron creates a new Tron engine instance
func NewTron() *Tron {
	return &Tron{
		chain: chaintron.NewChain(),
	}
}

// Supports returns true if this engine supports the given chain
func (t *Tron) Supports(chain common.Chain) bool {
	return chain == common.Tron
}

// Evaluate validates a TRON transaction against policy rules
func (t *Tron) Evaluate(rule *types.Rule, txBytes []byte) error {
	fmt.Printf("[TRON DEBUG] Evaluate called with %d bytes: %s\n", len(txBytes), hex.EncodeToString(txBytes[:min(len(txBytes), 100)]))

	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
	}

	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}

	decodedTx, err := t.chain.ParseTransactionBytes(txBytes)
	if err != nil {
		fmt.Printf("[TRON DEBUG] ParseTransactionBytes failed: %v\n", err)
		return fmt.Errorf("failed to parse TRON transaction: %w", err)
	}

	parsedTx, ok := decodedTx.(*chaintron.ParsedTronTransaction)
	if !ok {
		return fmt.Errorf("unexpected transaction type: %T", decodedTx)
	}

	rawData := parsedTx.GetRawData()
	if rawData == nil {
		return fmt.Errorf("transaction raw data is nil")
	}

	if len(rawData.Contract) == 0 {
		return fmt.Errorf("transaction has no contracts")
	}

	if len(rawData.Contract) != 1 {
		return fmt.Errorf("only single-contract transactions supported, got %d contracts", len(rawData.Contract))
	}

	contract := rawData.Contract[0]
	switch contract.Type {
	case "TransferContract":
		if r.GetProtocolId() != "trx" {
			return fmt.Errorf("unexpected protocol for TransferContract: %s", r.GetProtocolId())
		}
		if err := t.validateTarget(r, rule.GetTarget(), parsedTx); err != nil {
			return fmt.Errorf("failed to validate target: %w", err)
		}
		if err := t.validateParameterConstraints(r, rule.GetParameterConstraints(), parsedTx); err != nil {
			return fmt.Errorf("failed to validate parameter constraints: %w", err)
		}
	case "TriggerSmartContract":
		if r.GetProtocolId() != "trc20" {
			return fmt.Errorf("unexpected protocol for TriggerSmartContract: %s", r.GetProtocolId())
		}
		if err := t.validateTRC20Transfer(r, rule, parsedTx); err != nil {
			return fmt.Errorf("failed to validate TRC-20 transfer: %w", err)
		}
	default:
		return fmt.Errorf("unsupported contract type: %s", contract.Type)
	}

	return nil
}

// validateTarget validates the transaction target against the rule target
func (t *Tron) validateTarget(resource *types.ResourcePath, target *types.Target, tx *chaintron.ParsedTronTransaction) error {
	if target == nil || target.GetTargetType() == types.TargetType_TARGET_TYPE_UNSPECIFIED {
		return nil
	}

	actualDestination := tx.To()

	switch target.GetTargetType() {
	case types.TargetType_TARGET_TYPE_ADDRESS:
		expectedAddress := target.GetAddress()
		if expectedAddress == "" {
			return fmt.Errorf("target address cannot be empty")
		}
		if actualDestination != expectedAddress {
			return fmt.Errorf("target address mismatch: expected=%s, actual=%s",
				expectedAddress, actualDestination)
		}

	case types.TargetType_TARGET_TYPE_MAGIC_CONSTANT:
		resolve, err := resolver.NewMagicConstantRegistry().GetResolver(target.GetMagicConstant())
		if err != nil {
			return fmt.Errorf(
				"failed to get resolver: magic_const=%s",
				target.GetMagicConstant().String(),
			)
		}

		resolvedAddr, _, err := resolve.Resolve(
			target.GetMagicConstant(),
			resource.ChainId,
			"default",
		)
		if err != nil {
			return fmt.Errorf(
				"failed to resolve magic const: value=%s, error=%w",
				target.GetMagicConstant().String(),
				err,
			)
		}

		if actualDestination != resolvedAddr {
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_magic_const_resolved=%s",
				actualDestination,
				resolvedAddr,
			)
		}

	default:
		return fmt.Errorf("unsupported target type: %s", target.GetTargetType())
	}

	return nil
}

// validateParameterConstraints validates all parameter constraints
func (t *Tron) validateParameterConstraints(resource *types.ResourcePath, constraints []*types.ParameterConstraint, tx *chaintron.ParsedTronTransaction) error {
	for _, constraint := range constraints {
		paramName := constraint.GetParameterName()

		value, err := t.extractParameterValue(paramName, tx)
		if err != nil {
			return fmt.Errorf("failed to extract parameter %s: %w", paramName, err)
		}

		if err := t.assertArgsByType(resource.ChainId, paramName, value, constraints); err != nil {
			return fmt.Errorf("constraint validation failed for parameter %s: %w", paramName, err)
		}
	}
	return nil
}

// extractParameterValue extracts the actual value from transaction for the given parameter name
func (t *Tron) extractParameterValue(paramName string, tx *chaintron.ParsedTronTransaction) (interface{}, error) {
	switch paramName {
	case "recipient":
		return tx.To(), nil
	case "amount":
		return tx.GetAmount(), nil
	case "memo":
		return tx.GetMemo(), nil
	default:
		return nil, fmt.Errorf("unsupported parameter: %s", paramName)
	}
}

// assertArgsByType validates constraints using the appropriate comparator based on Go type
func (t *Tron) assertArgsByType(chainId, inputName string, arg interface{}, constraints []*types.ParameterConstraint) error {
	switch actual := arg.(type) {
	case string:
		err := stdcompare.AssertArg(
			chainId,
			constraints,
			inputName,
			actual,
			stdcompare.NewString,
		)
		if err != nil {
			return fmt.Errorf("failed to assert string parameter: %w", err)
		}

	case *big.Int:
		err := stdcompare.AssertArg(
			chainId,
			constraints,
			inputName,
			actual,
			stdcompare.NewBigInt,
		)
		if err != nil {
			return fmt.Errorf("failed to assert big.Int parameter: %w", err)
		}

	default:
		return fmt.Errorf("unsupported parameter type: %T", actual)
	}
	return nil
}

// validateTRC20Transfer validates a TRC-20 token transfer (TriggerSmartContract)
func (t *Tron) validateTRC20Transfer(resource *types.ResourcePath, rule *types.Rule, tx *chaintron.ParsedTronTransaction) error {
	callData := tx.GetCallData()
	if callData == "" {
		return fmt.Errorf("TRC-20 transfer missing call data")
	}

	if len(callData) < 8 {
		return fmt.Errorf("call data too short: %s", callData)
	}

	funcSelector := callData[:8]
	if funcSelector != "a9059cbb" {
		return fmt.Errorf("invalid function selector, expected transfer(address,uint256), got: %s", funcSelector)
	}

	if len(callData) < 136 {
		return fmt.Errorf("call data incomplete for transfer: need 136 chars (4+32+32 bytes as hex), got %d", len(callData))
	}

	recipientHex := callData[8:72]
	recipientAddr, err := t.hexToTronAddress(recipientHex)
	if err != nil {
		return fmt.Errorf("failed to decode recipient address: %w", err)
	}

	amountHex := callData[72:136]
	amount := new(big.Int)
	amount.SetString(amountHex, 16)

	contractAddr := tx.GetContractAddress()

	for _, constraint := range rule.GetParameterConstraints() {
		switch constraint.GetParameterName() {
		case "recipient":
			expectedRecipient := constraint.GetConstraint().GetFixedValue()
			if expectedRecipient != "" && !strings.EqualFold(recipientAddr, expectedRecipient) {
				return fmt.Errorf("recipient mismatch: expected %s, got %s", expectedRecipient, recipientAddr)
			}
			magicConst := constraint.GetConstraint().GetMagicConstantValue()
			if magicConst != types.MagicConstant_UNSPECIFIED {
				resolve, err := resolver.NewMagicConstantRegistry().GetResolver(magicConst)
				if err != nil {
					return fmt.Errorf("failed to get resolver: %w", err)
				}
				resolvedAddr, _, err := resolve.Resolve(
					magicConst,
					resource.ChainId,
					"default",
				)
				if err != nil {
					return fmt.Errorf("failed to resolve magic constant: %w", err)
				}
				if !strings.EqualFold(recipientAddr, resolvedAddr) {
					return fmt.Errorf("recipient mismatch: expected %s (resolved), got %s", resolvedAddr, recipientAddr)
				}
			}

		case "amount":
			err := stdcompare.AssertArg(
				resource.ChainId,
				rule.GetParameterConstraints(),
				"amount",
				amount,
				stdcompare.NewBigInt,
			)
			if err != nil {
				return fmt.Errorf("amount constraint failed: %w", err)
			}

		case "from_asset":
			expectedContract := constraint.GetConstraint().GetFixedValue()
			if expectedContract != "" && !strings.EqualFold(contractAddr, expectedContract) {
				return fmt.Errorf("contract address mismatch: expected %s, got %s", expectedContract, contractAddr)
			}

		case "memo":
			memo := tx.GetMemo()
			err := stdcompare.AssertArg(
				resource.ChainId,
				rule.GetParameterConstraints(),
				"memo",
				memo,
				stdcompare.NewString,
			)
			if err != nil {
				return fmt.Errorf("memo constraint failed: %w", err)
			}
		}
	}

	return nil
}

// hexToTronAddress converts a hex-encoded address (with leading zeros) to TRON base58 address
func (t *Tron) hexToTronAddress(hexAddr string) (string, error) {
	hexAddr = strings.TrimPrefix(hexAddr, "0x")
	hexAddr = strings.TrimLeft(hexAddr, "0")
	if len(hexAddr) < 40 {
		hexAddr = strings.Repeat("0", 40-len(hexAddr)) + hexAddr
	}
	hexAddr = "41" + hexAddr

	addrBytes, err := hex.DecodeString(hexAddr)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex address: %w", err)
	}

	return tronBase58CheckEncode(addrBytes), nil
}

// tronBase58CheckEncode encodes bytes to TRON base58check format
func tronBase58CheckEncode(input []byte) string {
	firstHash := sha256(input)
	secondHash := sha256(firstHash)
	checksum := secondHash[:4]

	result := make([]byte, len(input)+4)
	copy(result, input)
	copy(result[len(input):], checksum)

	return base58Encode(result)
}

// sha256 computes SHA256 hash
func sha256(data []byte) []byte {
	h := cryptoSha256.Sum256(data)
	return h[:]
}

// base58Encode encodes bytes to base58
func base58Encode(input []byte) string {
	const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	x := new(big.Int).SetBytes(input)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := new(big.Int)

	var result []byte
	for x.Cmp(zero) > 0 {
		x.DivMod(x, base, mod)
		result = append([]byte{alphabet[mod.Int64()]}, result...)
	}

	for _, b := range input {
		if b != 0 {
			break
		}
		result = append([]byte{alphabet[0]}, result...)
	}

	return string(result)
}

// ExtractTxBytes extracts transaction bytes from a base64-encoded Tron transaction.
func (t *Tron) ExtractTxBytes(txData string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(txData)
}
