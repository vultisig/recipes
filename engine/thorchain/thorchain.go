package thorchain

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	chainthorchain "github.com/vultisig/recipes/chain/thorchain"
	stdcompare "github.com/vultisig/recipes/engine/compare"
	"github.com/vultisig/recipes/resolver"
	vtypes "github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
)

// Thorchain represents the Thorchain engine implementation
type Thorchain struct {
	chain *chainthorchain.Chain
	cdc   codec.Codec
}

// NewThorchain creates a new Thorchain engine instance
func NewThorchain() *Thorchain {
	ir := codectypes.NewInterfaceRegistry()

	// Register crypto types (required for PubKey interfaces)
	cryptocodec.RegisterInterfaces(ir)

	// Register bank message types
	banktypes.RegisterInterfaces(ir)

	// Register the generated protobuf MsgDeposit for THORChain swaps
	ir.RegisterImplementations((*sdk.Msg)(nil), &vtypes.MsgDeposit{})

	return &Thorchain{
		chain: chainthorchain.NewChain(),
		cdc:   codec.NewProtoCodec(ir),
	}
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

	// Validate message type is supported (MsgSend or MsgDeposit)
	msg := txData.Body.Messages[0]
	mt, err := t.detectMessageType(msg)
	if err != nil {
		return fmt.Errorf("unsupported message type: %w", err)
	}

	// Enforce resource ↔ message-type compatibility
	if err := t.ensureResourceMessageCompatibility(r, mt); err != nil {
		return err
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
// Supports base64-encoded protobuf format with security hardening
func (t *Thorchain) parseTransaction(txBytes []byte) (*tx.Tx, error) {
	const maxTxBytes = 32 * 1024 // 32 KB - sufficient for complex transactions while preventing DoS
	if len(txBytes) > maxTxBytes {
		return nil, fmt.Errorf("transaction too large: %d bytes (max %d)", len(txBytes), maxTxBytes)
	}

	if len(txBytes) == 0 {
		return nil, fmt.Errorf("empty transaction data")
	}

	// Unmarshal protobuf bytes to tx.Tx
	var txData tx.Tx
	if err := t.cdc.Unmarshal(txBytes, &txData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal protobuf transaction: %w", err)
	}

	return &txData, nil
}

// ensureResourceMessageCompatibility rejects mismatched resource/message combinations.
func (t *Thorchain) ensureResourceMessageCompatibility(resource *vtypes.ResourcePath, mt chainthorchain.MessageType) error {
	switch resource.GetProtocolId() {
	case "send":
		if mt != chainthorchain.MessageTypeSend {
			return fmt.Errorf("resource %s.%s only allows MsgSend, got %s", resource.GetProtocolId(), resource.GetFunctionId(), mt)
		}
	case "thorchain_swap":
		if mt != chainthorchain.MessageTypeDeposit {
			return fmt.Errorf("resource %s.%s only allows MsgDeposit, got %s", resource.GetProtocolId(), resource.GetFunctionId(), mt)
		}
	default:
		// If new protocols are introduced, fail closed.
		return fmt.Errorf("unsupported protocol: %s", resource.GetProtocolId())
	}
	return nil
}

// unpackMsgSend unpacks a message to bank MsgSend type
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
		return nil, fmt.Errorf("expected bank MsgSend, got: %T", sdkMsg)
	}

	return msgSend, nil
}

// unpackMsgDeposit unpacks a message to THORChain MsgDeposit type
func (t *Thorchain) unpackMsgDeposit(msg *codectypes.Any) (*vtypes.MsgDeposit, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil message")
	}

	var sdkMsg sdk.Msg
	if err := t.cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return nil, fmt.Errorf("failed to unpack sdk.Msg: %w (typeUrl=%s)", err, msg.TypeUrl)
	}

	msgDeposit, ok := sdkMsg.(*vtypes.MsgDeposit)
	if !ok {
		return nil, fmt.Errorf("expected THORChain MsgDeposit, got: %T", sdkMsg)
	}

	return msgDeposit, nil
}

// detectMessageType determines the message type based on TypeUrl
func (t *Thorchain) detectMessageType(msg *codectypes.Any) (chainthorchain.MessageType, error) {
	if msg == nil {
		return chainthorchain.MessageType(0), fmt.Errorf("nil message")
	}

	messageType, err := chainthorchain.GetMessageTypeFromTypeUrl(msg.TypeUrl)
	if err != nil {
		return chainthorchain.MessageType(0), fmt.Errorf("failed to detect message type: %w", err)
	}

	return messageType, nil
}

// unpackMessage unpacks a message to the appropriate type based on its TypeUrl
// Returns the unpacked message and its type
func (t *Thorchain) unpackMessage(msg *codectypes.Any) (interface{}, chainthorchain.MessageType, error) {
	messageType, err := t.detectMessageType(msg)
	if err != nil {
		return nil, chainthorchain.MessageType(0), err
	}

	switch messageType {
	case chainthorchain.MessageTypeSend:
		msgSend, err := t.unpackMsgSend(msg)
		if err != nil {
			return nil, chainthorchain.MessageTypeSend, fmt.Errorf("failed to unpack MsgSend: %w", err)
		}
		return msgSend, chainthorchain.MessageTypeSend, nil

	case chainthorchain.MessageTypeDeposit:
		msgDeposit, err := t.unpackMsgDeposit(msg)
		if err != nil {
			return nil, chainthorchain.MessageTypeDeposit, fmt.Errorf("failed to unpack MsgDeposit: %w", err)
		}
		return msgDeposit, chainthorchain.MessageTypeDeposit, nil

	default:
		return nil, chainthorchain.MessageType(0), fmt.Errorf("unsupported message type: %s", messageType)
	}
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

	// Use unified message unpacking to support both MsgSend and MsgDeposit
	msg := txData.Body.Messages[0]
	unpackedMsg, messageType, err := t.unpackMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to unpack message for target validation: %w", err)
	}

	if messageType != chainthorchain.MessageTypeSend {
		return fmt.Errorf("target validation only supported for MsgSend transactions, use TARGET_TYPE_UNSPECIFIED for %s", messageType)
	}

	msgSend, ok := unpackedMsg.(*banktypes.MsgSend)
	if !ok {
		return fmt.Errorf("expected MsgSend, got: %T", unpackedMsg)
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

	// Use unified message unpacking to support both MsgSend and MsgDeposit
	msg := txData.Body.Messages[0]
	unpackedMsg, messageType, err := t.unpackMessage(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack message: %w", err)
	}

	switch messageType {
	case chainthorchain.MessageTypeSend:
		return t.extractParameterFromMsgSend(paramName, unpackedMsg.(*banktypes.MsgSend), txData)
	case chainthorchain.MessageTypeDeposit:
		return t.extractParameterFromMsgDeposit(paramName, unpackedMsg.(*vtypes.MsgDeposit))
	default:
		return nil, fmt.Errorf("unsupported message type: %s", messageType)
	}
}

// extractParameterFromMsgSend extracts parameters from MsgSend
func (t *Thorchain) extractParameterFromMsgSend(paramName string, msgSend *banktypes.MsgSend, txData *tx.Tx) (any, error) {
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

// extractParameterFromMsgDeposit extracts parameters from MsgDeposit (for swaps)
func (t *Thorchain) extractParameterFromMsgDeposit(paramName string, msgDeposit *vtypes.MsgDeposit) (any, error) {
	switch paramName {
	case "amount":
		if len(msgDeposit.Coins) == 0 {
			return nil, fmt.Errorf("no coins in deposit message")
		}
		if len(msgDeposit.Coins) != 1 {
			return nil, fmt.Errorf("multi-coin deposits not supported, got %d coins", len(msgDeposit.Coins))
		}
		coin := msgDeposit.Coins[0]
		v := new(big.Int)
		if _, ok := v.SetString(coin.Amount, 10); !ok {
			return nil, fmt.Errorf("invalid amount format: %s", coin.Amount)
		}
		return v, nil
	case "memo":
		// For MsgDeposit, memo contains swap instructions including destination address
		return msgDeposit.Memo, nil
	case "from_asset":
		if len(msgDeposit.Coins) == 0 {
			return nil, fmt.Errorf("no coins in deposit message")
		}
		if len(msgDeposit.Coins) != 1 {
			return nil, fmt.Errorf("multi-coin deposits not supported, got %d coins", len(msgDeposit.Coins))
		}
		coin := msgDeposit.Coins[0]
		if coin.Asset == nil {
			return nil, fmt.Errorf("coin missing asset information")
		}
		// Return the asset symbol in uppercase
		return strings.ToUpper(coin.Asset.Symbol), nil
	default:
		// Note: 'recipient' parameter is not supported for THORChain swaps as there's no explicit
		// destination address - the real destination is encoded in the memo field
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
