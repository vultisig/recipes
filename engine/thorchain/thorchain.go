package thorchain

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stdcompare "github.com/vultisig/recipes/engine/compare"
	"github.com/vultisig/recipes/resolver"
	vtypes "github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
	// thortypes "gitlab.com/thorchain/thornode/x/thorchain/types" // when you add MsgDeposit
)

// Thorchain represents the Thorchain engine implementation
type Thorchain struct {
	cdc codec.Codec
}

// NewThorchain creates a new Thorchain engine instance
func NewThorchain() *Thorchain {
	ir := codectypes.NewInterfaceRegistry()

	banktypes.RegisterInterfaces(ir)

	return &Thorchain{cdc: codec.NewProtoCodec(ir)}
}

// Supports returns true if this engine supports the given chain
func (t *Thorchain) Supports(chain common.Chain) bool {
	return chain == common.THORChain
}

// Evaluate validates a Thorchain transaction against policy rules
func (t *Thorchain) Evaluate(rule *vtypes.Rule, txBytes []byte) error {
	// Validate rule effect is ALLOW (following existing pattern from other engines)
	if rule.GetEffect().String() != vtypes.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
	}

	// Parse resource to extract protocol and function information
	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}

	txData, err := t.parseTransaction(txBytes)
	if err != nil {
		return fmt.Errorf("failed to parse Thorchain transaction: %w", err)
	}

	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return fmt.Errorf("transaction must have at least one message")
	}

	if len(txData.Body.Messages) != 1 {
		return fmt.Errorf("only single-message transactions supported, got %d messages", len(txData.Body.Messages))
	}

	// TODO add MsgDeposit support for thorchain swaps
	msg := txData.Body.Messages[0]
	if err := t.validateMessageType(msg); err != nil {
		return fmt.Errorf("unsupported message type: %w", err)
	}

	if err := t.validateTarget(r, rule.GetTarget(), txData); err != nil {
		return fmt.Errorf("failed to validate target: %w", err)
	}

	if err := t.validateParameterConstraints(r, rule.GetParameterConstraints(), txData); err != nil {
		return fmt.Errorf("failed to validate parameter constraints: %w", err)
	}

	return nil
}

// parseTransaction parses Thorchain transaction bytes into a Cosmos SDK transaction
// Supports protobuf, JSON, base64-encoded, and hex-encoded formats with security hardening
func (t *Thorchain) parseTransaction(txBytes []byte) (*tx.Tx, error) {
	const maxTxBytes = 32 * 1024 // 32 KB - sufficient for complex transactions while preventing DoS
	if len(txBytes) > maxTxBytes {
		return nil, fmt.Errorf("transaction too large: %d bytes (max %d)", len(txBytes), maxTxBytes)
	}

	if len(txBytes) == 0 {
		return nil, fmt.Errorf("empty transaction data")
	}

	// Try to parse as protobuf first (standard Cosmos SDK format)
	var txData tx.Tx
	if err := t.cdc.Unmarshal(txBytes, &txData); err == nil {
		return &txData, nil
	}

	// If protobuf parsing fails, try JSON format
	if err := json.Unmarshal(txBytes, &txData); err == nil {
		return &txData, nil
	}

	// If JSON parsing fails, try base64 (common transport encoding for protobuf bytes)
	if b64, err := base64.StdEncoding.DecodeString(string(txBytes)); err == nil {
		// Check decoded size to prevent DoS via large decoded payloads
		if len(b64) > maxTxBytes {
			return nil, fmt.Errorf("decoded base64 transaction too large: %d bytes (max %d)", len(b64), maxTxBytes)
		}
		if err := t.cdc.Unmarshal(b64, &txData); err == nil {
			return &txData, nil
		}
		if err := json.Unmarshal(b64, &txData); err == nil {
			return &txData, nil
		}
	}

	// Finally, try hex (handle 0x/0X prefixes)
	s := strings.TrimPrefix(string(txBytes), "0x")
	s = strings.TrimPrefix(s, "0X")
	decodedBytes, hexErr := hex.DecodeString(s)
	if hexErr != nil {
		return nil, fmt.Errorf("failed to parse transaction as protobuf, JSON, base64, or hex")
	}

	// Check decoded size to prevent DoS via large decoded payloads
	if len(decodedBytes) > maxTxBytes {
		return nil, fmt.Errorf("decoded hex transaction too large: %d bytes (max %d)", len(decodedBytes), maxTxBytes)
	}

	// Try parsing decoded hex bytes as protobuf
	if err := t.cdc.Unmarshal(decodedBytes, &txData); err == nil {
		return &txData, nil
	}

	// Try parsing decoded hex bytes as JSON
	if err := json.Unmarshal(decodedBytes, &txData); err == nil {
		return &txData, nil
	}

	return nil, fmt.Errorf("failed to parse transaction as protobuf, JSON, base64, or hex")
}

// unpackMsgSend unpacks a message to MsgSend type
func (t *Thorchain) unpackMsgSend(msg *codectypes.Any) (*banktypes.MsgSend, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil message")
	}

	var sdkMsg sdk.Msg
	if err := t.cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return nil, fmt.Errorf("failed to unpack sdk.Msg: %w (typeUrl=%s)", err, msg.TypeUrl)
	}

	msgSend, ok := sdkMsg.(*banktypes.MsgSend)
	if !ok {
		return nil, fmt.Errorf("expected MsgSend, got: %T", sdkMsg)
	}

	return msgSend, nil
}

// validateMessageType validates that the message is MsgSend (direct transfers only)
func (t *Thorchain) validateMessageType(msg *codectypes.Any) error {
	_, err := t.unpackMsgSend(msg)
	if err != nil {
		// TODO: Add support for MsgDeposit
		return fmt.Errorf("only MsgSend transactions are supported for now, got: %s", msg.TypeUrl)
	}
	return nil
}

// validateTarget validates the transaction target against the rule target
func (t *Thorchain) validateTarget(resource *vtypes.ResourcePath, target *vtypes.Target, txData *tx.Tx) error {
	if target == nil || target.GetTargetType() == vtypes.TargetType_TARGET_TYPE_UNSPECIFIED {
		return nil // No target validation required
	}

	// Extract the first message (assuming single message transactions for now)
	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return fmt.Errorf("no messages in transaction")
	}

	// Unpack the first message as MsgSend
	// TODO add support for MsgDeposit
	msg := txData.Body.Messages[0]
	msgSend, err := t.unpackMsgSend(msg)
	if err != nil {
		return err
	}

	switch target.GetTargetType() {
	case vtypes.TargetType_TARGET_TYPE_ADDRESS:
		expectedAddress := target.GetAddress()
		if expectedAddress == "" {
			return fmt.Errorf("target address cannot be empty")
		}
		if msgSend.ToAddress != expectedAddress {
			return fmt.Errorf("target address mismatch: expected=%s, actual=%s",
				expectedAddress, msgSend.ToAddress)
		}

	case vtypes.TargetType_TARGET_TYPE_MAGIC_CONSTANT:
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

		if msgSend.ToAddress != resolvedAddr {
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_magic_const_resolved=%s",
				msgSend.ToAddress,
				resolvedAddr,
			)
		}

	default:
		return fmt.Errorf("unsupported target type: %s", target.GetTargetType())
	}

	return nil
}

// validateParameterConstraints validates all parameter constraints
func (t *Thorchain) validateParameterConstraints(resource *vtypes.ResourcePath, constraints []*vtypes.ParameterConstraint, txData *tx.Tx) error {
	for _, constraint := range constraints {
		paramName := constraint.GetParameterName()

		value, err := t.extractParameterValue(paramName, txData)
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
func (t *Thorchain) extractParameterValue(paramName string, txData *tx.Tx) (any, error) {
	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return nil, fmt.Errorf("no messages in transaction")
	}

	// Unpack the first message as MsgSend
	// TODO add support for MsgDeposit
	msg := txData.Body.Messages[0]
	msgSend, err := t.unpackMsgSend(msg)
	if err != nil {
		return nil, err
	}

	switch paramName {
	case "recipient":
		return msgSend.ToAddress, nil
	case "amount":
		if len(msgSend.Amount) == 0 {
			return nil, fmt.Errorf("no amount in message")
		}
		if len(msgSend.Amount) != 1 {
			return nil, fmt.Errorf("multi-coin transfers not supported, got %d coins", len(msgSend.Amount))
		}
		coin := msgSend.Amount[0]
		return coin.Amount.BigInt(), nil
	case "memo":
		return txData.Body.Memo, nil
	case "denom":
		if len(msgSend.Amount) == 0 {
			return nil, fmt.Errorf("no amount in message")
		}
		if len(msgSend.Amount) != 1 {
			return nil, fmt.Errorf("multi-coin transfers not supported, got %d coins", len(msgSend.Amount))
		}
		return msgSend.Amount[0].Denom, nil
	default:
		return nil, fmt.Errorf("unsupported parameter: %s", paramName)
	}
}

// assertArgsByType validates constraints using the appropriate comparator based on Go type
func (t *Thorchain) assertArgsByType(chainId, inputName string, arg any, constraints []*vtypes.ParameterConstraint) error {
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
