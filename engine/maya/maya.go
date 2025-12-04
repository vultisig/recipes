package maya

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
	chainmaya "github.com/vultisig/recipes/chain/maya"
	stdcompare "github.com/vultisig/recipes/engine/compare"
	"github.com/vultisig/recipes/resolver"
	vtypes "github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
)

// Maya represents the MAYAChain engine implementation
type Maya struct {
	chain *chainmaya.Chain
	cdc   codec.Codec
}

// NewMaya creates a new Maya engine instance
func NewMaya() *Maya {
	ir := codectypes.NewInterfaceRegistry()

	// Register crypto types (required for PubKey interfaces)
	cryptocodec.RegisterInterfaces(ir)

	// Register bank message types
	banktypes.RegisterInterfaces(ir)

	// Register the generated protobuf MsgDeposit for MAYAChain swaps
	ir.RegisterImplementations((*sdk.Msg)(nil), &vtypes.MsgDeposit{})

	return &Maya{
		chain: chainmaya.NewChain(),
		cdc:   codec.NewProtoCodec(ir),
	}
}

// Supports returns true if this engine supports the given chain
func (m *Maya) Supports(chain common.Chain) bool {
	return chain == common.MayaChain
}

// Evaluate validates a MAYAChain transaction against policy rules
func (m *Maya) Evaluate(rule *vtypes.Rule, txBytes []byte) error {
	if rule.GetEffect().String() != vtypes.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
	}

	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}

	txData, err := m.parseTransaction(txBytes)
	if err != nil {
		return fmt.Errorf("failed to parse MAYAChain transaction: %w", err)
	}

	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return fmt.Errorf("transaction must have at least one message")
	}

	if len(txData.Body.Messages) != 1 {
		return fmt.Errorf("only single-message transactions supported, got %d messages", len(txData.Body.Messages))
	}

	msg := txData.Body.Messages[0]
	mt, err := m.detectMessageType(msg)
	if err != nil {
		return fmt.Errorf("unsupported message type: %w", err)
	}

	if err := m.ensureResourceMessageCompatibility(r, mt); err != nil {
		return err
	}

	if err := m.validateTarget(r, rule.GetTarget(), txData); err != nil {
		return fmt.Errorf("failed to validate target: %w", err)
	}

	if err := m.validateParameterConstraints(r, rule.GetParameterConstraints(), txData); err != nil {
		return fmt.Errorf("failed to validate parameter constraints: %w", err)
	}

	return nil
}

// parseTransaction parses MAYAChain transaction bytes into a Cosmos SDK transaction
func (m *Maya) parseTransaction(txBytes []byte) (*tx.Tx, error) {
	const maxTxBytes = 32 * 1024
	if len(txBytes) > maxTxBytes {
		return nil, fmt.Errorf("transaction too large: %d bytes (max %d)", len(txBytes), maxTxBytes)
	}

	if len(txBytes) == 0 {
		return nil, fmt.Errorf("empty transaction data")
	}

	var txData tx.Tx
	if err := m.cdc.Unmarshal(txBytes, &txData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal protobuf transaction: %w", err)
	}

	return &txData, nil
}

// ensureResourceMessageCompatibility rejects mismatched resource/message combinations.
func (m *Maya) ensureResourceMessageCompatibility(resource *vtypes.ResourcePath, mt chainmaya.MessageType) error {
	switch resource.GetProtocolId() {
	case "send", "cacao":
		if mt != chainmaya.MessageTypeSend {
			return fmt.Errorf("resource %s.%s only allows MsgSend, got %s", resource.GetProtocolId(), resource.GetFunctionId(), mt)
		}
	case "mayachain_swap":
		if mt != chainmaya.MessageTypeDeposit {
			return fmt.Errorf("resource %s.%s only allows MsgDeposit, got %s", resource.GetProtocolId(), resource.GetFunctionId(), mt)
		}
	default:
		return fmt.Errorf("unsupported protocol: %s", resource.GetProtocolId())
	}
	return nil
}

// unpackMsgSend unpacks a message to bank MsgSend type
func (m *Maya) unpackMsgSend(msg *codectypes.Any) (*banktypes.MsgSend, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil message")
	}

	var sdkMsg sdk.Msg
	if err := m.cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return nil, fmt.Errorf("failed to unpack sdk.Msg: %w (typeUrl=%s)", err, msg.TypeUrl)
	}

	msgSend, ok := sdkMsg.(*banktypes.MsgSend)
	if !ok {
		return nil, fmt.Errorf("expected bank MsgSend, got: %T", sdkMsg)
	}

	return msgSend, nil
}

// unpackMsgDeposit unpacks a message to MAYAChain MsgDeposit type
func (m *Maya) unpackMsgDeposit(msg *codectypes.Any) (*vtypes.MsgDeposit, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil message")
	}

	var sdkMsg sdk.Msg
	if err := m.cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return nil, fmt.Errorf("failed to unpack sdk.Msg: %w (typeUrl=%s)", err, msg.TypeUrl)
	}

	msgDeposit, ok := sdkMsg.(*vtypes.MsgDeposit)
	if !ok {
		return nil, fmt.Errorf("expected MAYAChain MsgDeposit, got: %T", sdkMsg)
	}

	return msgDeposit, nil
}

// detectMessageType determines the message type based on TypeUrl
func (m *Maya) detectMessageType(msg *codectypes.Any) (chainmaya.MessageType, error) {
	if msg == nil {
		return chainmaya.MessageType(0), fmt.Errorf("nil message")
	}

	messageType, err := chainmaya.GetMessageTypeFromTypeUrl(msg.TypeUrl)
	if err != nil {
		return chainmaya.MessageType(0), fmt.Errorf("failed to detect message type: %w", err)
	}

	return messageType, nil
}

// unpackMessage unpacks a message to the appropriate type based on its TypeUrl
func (m *Maya) unpackMessage(msg *codectypes.Any) (interface{}, chainmaya.MessageType, error) {
	messageType, err := m.detectMessageType(msg)
	if err != nil {
		return nil, chainmaya.MessageType(0), err
	}

	switch messageType {
	case chainmaya.MessageTypeSend:
		msgSend, err := m.unpackMsgSend(msg)
		if err != nil {
			return nil, chainmaya.MessageTypeSend, fmt.Errorf("failed to unpack MsgSend: %w", err)
		}
		return msgSend, chainmaya.MessageTypeSend, nil

	case chainmaya.MessageTypeDeposit:
		msgDeposit, err := m.unpackMsgDeposit(msg)
		if err != nil {
			return nil, chainmaya.MessageTypeDeposit, fmt.Errorf("failed to unpack MsgDeposit: %w", err)
		}
		return msgDeposit, chainmaya.MessageTypeDeposit, nil

	default:
		return nil, chainmaya.MessageType(0), fmt.Errorf("unsupported message type: %s", messageType)
	}
}

// validateTarget validates the transaction target against the rule target
func (m *Maya) validateTarget(resource *vtypes.ResourcePath, target *vtypes.Target, txData *tx.Tx) error {
	if target == nil || target.GetTargetType() == vtypes.TargetType_TARGET_TYPE_UNSPECIFIED {
		return nil
	}

	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return fmt.Errorf("no messages in transaction")
	}

	msg := txData.Body.Messages[0]
	unpackedMsg, messageType, err := m.unpackMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to unpack message for target validation: %w", err)
	}

	if messageType != chainmaya.MessageTypeSend {
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
func (m *Maya) validateParameterConstraints(resource *vtypes.ResourcePath, constraints []*vtypes.ParameterConstraint, txData *tx.Tx) error {
	for _, constraint := range constraints {
		paramName := constraint.GetParameterName()

		value, err := m.extractParameterValue(paramName, txData)
		if err != nil {
			return fmt.Errorf("failed to extract parameter %s: %w", paramName, err)
		}

		if err := m.assertArgsByType(resource.ChainId, paramName, value, constraints); err != nil {
			return fmt.Errorf("constraint validation failed for parameter %s: %w", paramName, err)
		}
	}
	return nil
}

// extractParameterValue extracts the actual value from transaction for the given parameter name
func (m *Maya) extractParameterValue(paramName string, txData *tx.Tx) (any, error) {
	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return nil, fmt.Errorf("no messages in transaction")
	}

	msg := txData.Body.Messages[0]
	unpackedMsg, messageType, err := m.unpackMessage(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack message: %w", err)
	}

	switch messageType {
	case chainmaya.MessageTypeSend:
		return m.extractParameterFromMsgSend(paramName, unpackedMsg.(*banktypes.MsgSend), txData)
	case chainmaya.MessageTypeDeposit:
		return m.extractParameterFromMsgDeposit(paramName, unpackedMsg.(*vtypes.MsgDeposit))
	default:
		return nil, fmt.Errorf("unsupported message type: %s", messageType)
	}
}

// extractParameterFromMsgSend extracts parameters from MsgSend
func (m *Maya) extractParameterFromMsgSend(paramName string, msgSend *banktypes.MsgSend, txData *tx.Tx) (any, error) {
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
func (m *Maya) extractParameterFromMsgDeposit(paramName string, msgDeposit *vtypes.MsgDeposit) (any, error) {
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
		return strings.ToUpper(coin.Asset.Symbol), nil
	default:
		return nil, fmt.Errorf("unsupported parameter: %s", paramName)
	}
}

// assertArgsByType validates constraints using the appropriate comparator based on Go type
func (m *Maya) assertArgsByType(chainId, inputName string, arg any, constraints []*vtypes.ParameterConstraint) error {
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

