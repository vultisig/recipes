// Package cosmos provides a shared engine for Cosmos SDK-based blockchains.
package cosmos

import (
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/vultisig/recipes/chain/cosmos"
	"github.com/vultisig/recipes/engine/compare"
	"github.com/vultisig/recipes/resolver"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
)

// Config holds chain-specific configuration for the Cosmos engine.
type Config struct {
	// ChainID is the identifier used for magic constant resolution (e.g., "cosmos", "mayachain", "thorchain").
	ChainID string

	// SupportedChains is the list of common.Chain values this engine supports.
	SupportedChains []common.Chain

	// MessageTypeRegistry maps TypeUrls to MessageTypes for this chain.
	MessageTypeRegistry *cosmos.MessageTypeRegistry

	// ProtocolMessageTypes maps protocol IDs to allowed message types.
	// e.g., "atom" -> MessageTypeSend, "thorchain_swap" -> MessageTypeDeposit
	ProtocolMessageTypes map[string]cosmos.MessageType

	// RegisterExtraTypes is an optional function to register additional protobuf types
	RegisterExtraTypes func(ir codectypes.InterfaceRegistry)
}

// Engine is a generic Cosmos engine that can be configured for different chains.
type Engine struct {
	config Config
	cdc    codec.Codec
}

// NewEngine creates a new Cosmos engine with the given configuration.
func NewEngine(config Config) *Engine {
	ir := codectypes.NewInterfaceRegistry()

	// Register crypto types (required for PubKey interfaces)
	cryptocodec.RegisterInterfaces(ir)

	// Register bank message types
	banktypes.RegisterInterfaces(ir)

	// Register staking message types (MsgDelegate, MsgUndelegate, MsgBeginRedelegate, ...)
	stakingtypes.RegisterInterfaces(ir)

	// Register distribution message types (MsgWithdrawDelegatorReward, ...)
	distributiontypes.RegisterInterfaces(ir)

	// Register any extra types specific to this chain
	if config.RegisterExtraTypes != nil {
		config.RegisterExtraTypes(ir)
	}

	return &Engine{
		config: config,
		cdc:    codec.NewProtoCodec(ir),
	}
}

// Supports returns true if this engine supports the given chain.
func (e *Engine) Supports(chain common.Chain) bool {
	for _, c := range e.config.SupportedChains {
		if c == chain {
			return true
		}
	}
	return false
}

// Evaluate validates a Cosmos transaction against policy rules.
func (e *Engine) Evaluate(rule *types.Rule, txBytes []byte) error {
	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
	}

	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}

	txData, err := e.parseTransaction(txBytes)
	if err != nil {
		return fmt.Errorf("failed to parse %s transaction: %w", e.config.ChainID, err)
	}

	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return fmt.Errorf("transaction must have at least one message")
	}

	if len(txData.Body.Messages) != 1 {
		return fmt.Errorf("only single-message transactions supported, got %d messages", len(txData.Body.Messages))
	}

	msg := txData.Body.Messages[0]
	mt, err := e.detectMessageType(msg)
	if err != nil {
		return fmt.Errorf("unsupported message type: %w", err)
	}

	if err := e.ensureResourceMessageCompatibility(r, mt); err != nil {
		return err
	}

	if err := e.validateTarget(r, rule.GetTarget(), txData, mt); err != nil {
		return fmt.Errorf("failed to validate target: %w", err)
	}

	if err := e.validateParameterConstraints(r, rule.GetParameterConstraints(), txData, mt); err != nil {
		return fmt.Errorf("failed to validate parameter constraints: %w", err)
	}

	return nil
}

// parseTransaction parses Cosmos transaction bytes into a Cosmos SDK transaction.
func (e *Engine) parseTransaction(txBytes []byte) (*tx.Tx, error) {
	const maxTxBytes = 32 * 1024 // 32 KB
	if len(txBytes) > maxTxBytes {
		return nil, fmt.Errorf("transaction too large: %d bytes (max %d)", len(txBytes), maxTxBytes)
	}

	if len(txBytes) == 0 {
		return nil, fmt.Errorf("empty transaction data")
	}

	var txData tx.Tx
	if err := e.cdc.Unmarshal(txBytes, &txData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal protobuf transaction: %w", err)
	}

	return &txData, nil
}

// detectMessageType determines the message type based on TypeUrl.
func (e *Engine) detectMessageType(msg *codectypes.Any) (cosmos.MessageType, error) {
	if msg == nil {
		return cosmos.MessageTypeUnknown, fmt.Errorf("nil message")
	}

	messageType, err := e.config.MessageTypeRegistry.GetMessageType(msg.TypeUrl)
	if err != nil {
		return cosmos.MessageTypeUnknown, fmt.Errorf("failed to detect message type: %w", err)
	}

	return messageType, nil
}

// ensureResourceMessageCompatibility rejects mismatched resource/message combinations.
func (e *Engine) ensureResourceMessageCompatibility(resource *types.ResourcePath, mt cosmos.MessageType) error {
	expectedMT, exists := e.config.ProtocolMessageTypes[resource.GetProtocolId()]
	if !exists {
		return fmt.Errorf("unsupported protocol: %s", resource.GetProtocolId())
	}

	if mt != expectedMT {
		return fmt.Errorf("resource %s.%s only allows %s, got %s",
			resource.GetProtocolId(), resource.GetFunctionId(), expectedMT, mt)
	}
	return nil
}

// unpackMsgSend unpacks a message to bank MsgSend type.
func (e *Engine) unpackMsgSend(msg *codectypes.Any) (*banktypes.MsgSend, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil message")
	}

	var sdkMsg sdk.Msg
	if err := e.cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return nil, fmt.Errorf("failed to unpack sdk.Msg: %w (typeUrl=%s)", err, msg.TypeUrl)
	}

	msgSend, ok := sdkMsg.(*banktypes.MsgSend)
	if !ok {
		return nil, fmt.Errorf("expected bank MsgSend, got: %T", sdkMsg)
	}

	return msgSend, nil
}

// unpackMsgDeposit unpacks a message to MsgDeposit type.
func (e *Engine) unpackMsgDeposit(msg *codectypes.Any) (*types.MsgDeposit, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil message")
	}

	var sdkMsg sdk.Msg
	if err := e.cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return nil, fmt.Errorf("failed to unpack sdk.Msg: %w (typeUrl=%s)", err, msg.TypeUrl)
	}

	msgDeposit, ok := sdkMsg.(*types.MsgDeposit)
	if !ok {
		return nil, fmt.Errorf("expected MsgDeposit, got: %T", sdkMsg)
	}

	return msgDeposit, nil
}

// unpackMsgBeginRedelegate unpacks a message to staking MsgBeginRedelegate type.
func (e *Engine) unpackMsgBeginRedelegate(msg *codectypes.Any) (*stakingtypes.MsgBeginRedelegate, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil message")
	}

	var sdkMsg sdk.Msg
	if err := e.cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return nil, fmt.Errorf("failed to unpack sdk.Msg: %w (typeUrl=%s)", err, msg.TypeUrl)
	}

	msgRedelegate, ok := sdkMsg.(*stakingtypes.MsgBeginRedelegate)
	if !ok {
		return nil, fmt.Errorf("expected staking MsgBeginRedelegate, got: %T", sdkMsg)
	}

	return msgRedelegate, nil
}

// unpackMsgWithdrawDelegatorReward unpacks a message to distribution MsgWithdrawDelegatorReward type.
func (e *Engine) unpackMsgWithdrawDelegatorReward(msg *codectypes.Any) (*distributiontypes.MsgWithdrawDelegatorReward, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil message")
	}

	var sdkMsg sdk.Msg
	if err := e.cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return nil, fmt.Errorf("failed to unpack sdk.Msg: %w (typeUrl=%s)", err, msg.TypeUrl)
	}

	msgWithdraw, ok := sdkMsg.(*distributiontypes.MsgWithdrawDelegatorReward)
	if !ok {
		return nil, fmt.Errorf("expected distribution MsgWithdrawDelegatorReward, got: %T", sdkMsg)
	}

	return msgWithdraw, nil
}


// validateTarget validates the transaction target against the rule target.
func (e *Engine) validateTarget(resource *types.ResourcePath, target *types.Target, txData *tx.Tx, mt cosmos.MessageType) error {
	if target == nil || target.GetTargetType() == types.TargetType_TARGET_TYPE_UNSPECIFIED {
		return nil
	}

	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return fmt.Errorf("no messages in transaction")
	}

	// Target validation only supported for MsgSend
	if mt != cosmos.MessageTypeSend {
		return fmt.Errorf("target validation only supported for MsgSend transactions, use TARGET_TYPE_UNSPECIFIED for %s", mt)
	}

	msg := txData.Body.Messages[0]
	msgSend, err := e.unpackMsgSend(msg)
	if err != nil {
		return fmt.Errorf("failed to unpack message for target validation: %w", err)
	}

	switch target.GetTargetType() {
	case types.TargetType_TARGET_TYPE_ADDRESS:
		expectedAddress := target.GetAddress()
		if expectedAddress == "" {
			return fmt.Errorf("target address cannot be empty")
		}
		if msgSend.ToAddress != expectedAddress {
			return fmt.Errorf("target address mismatch: expected=%s, actual=%s",
				expectedAddress, msgSend.ToAddress)
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

// validateParameterConstraints validates all parameter constraints.
func (e *Engine) validateParameterConstraints(resource *types.ResourcePath, constraints []*types.ParameterConstraint, txData *tx.Tx, mt cosmos.MessageType) error {
	for _, constraint := range constraints {
		paramName := constraint.GetParameterName()

		value, err := e.extractParameterValue(paramName, txData, mt)
		if err != nil {
			return fmt.Errorf("failed to extract parameter %s: %w", paramName, err)
		}

		if err := e.assertArgsByType(resource.ChainId, paramName, value, constraints); err != nil {
			return fmt.Errorf("constraint validation failed for parameter %s: %w", paramName, err)
		}
	}
	return nil
}

// extractParameterValue extracts the actual value from transaction for the given parameter name.
func (e *Engine) extractParameterValue(paramName string, txData *tx.Tx, mt cosmos.MessageType) (any, error) {
	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return nil, fmt.Errorf("no messages in transaction")
	}

	msg := txData.Body.Messages[0]

	switch mt {
	case cosmos.MessageTypeSend:
		msgSend, err := e.unpackMsgSend(msg)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack message: %w", err)
		}
		return e.extractParameterFromMsgSend(paramName, msgSend, txData)

	case cosmos.MessageTypeDeposit:
		msgDeposit, err := e.unpackMsgDeposit(msg)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack message: %w", err)
		}
		return e.extractParameterFromMsgDeposit(paramName, msgDeposit)

	case cosmos.MessageTypeBeginRedelegate:
		msgRedelegate, err := e.unpackMsgBeginRedelegate(msg)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack message: %w", err)
		}
		return e.extractParameterFromMsgBeginRedelegate(paramName, msgRedelegate)

	case cosmos.MessageTypeWithdrawDelegatorReward:
		msgWithdraw, err := e.unpackMsgWithdrawDelegatorReward(msg)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack message: %w", err)
		}
		return e.extractParameterFromMsgWithdrawDelegatorReward(paramName, msgWithdraw)

	default:
		return nil, fmt.Errorf("unsupported message type: %s", mt)
	}
}

// extractParameterFromMsgSend extracts parameters from MsgSend.
func (e *Engine) extractParameterFromMsgSend(paramName string, msgSend *banktypes.MsgSend, txData *tx.Tx) (any, error) {
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

// extractParameterFromMsgDeposit extracts parameters from MsgDeposit (for swaps).
func (e *Engine) extractParameterFromMsgDeposit(paramName string, msgDeposit *types.MsgDeposit) (any, error) {
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
		return coin.Asset.Symbol, nil
	default:
		return nil, fmt.Errorf("unsupported parameter: %s", paramName)
	}
}

// extractParameterFromMsgBeginRedelegate extracts parameters from staking MsgBeginRedelegate.
//
// MsgBeginRedelegate moves stake from one validator (src) to another (dst) for a given
// delegator. We expose the addresses, the amount and the denom so that policy rules can
// constrain redelegation by validator allowlist or amount caps. We also enforce two
// invariants here that any healthy redelegate tx must satisfy:
//   - amount > 0 (zero-amount redelegations are nonsensical and rejected by the chain)
//   - validator_src_address != validator_dst_address (chain rejects, but we fail fast)
//
// Bech32 prefix validation for the addresses is intentionally not enforced at this layer:
// the Cosmos SDK proto-decode will accept any string, and chain-specific prefix checks
// belong on the per-chain engine wiring (where Bech32Prefix is known) — this generic
// extractor mirrors how MsgSend handles ToAddress without prefix checks.
func (e *Engine) extractParameterFromMsgBeginRedelegate(paramName string, msg *stakingtypes.MsgBeginRedelegate) (any, error) {
	if msg.ValidatorSrcAddress == "" {
		return nil, fmt.Errorf("redelegate validator_src_address required")
	}
	if msg.ValidatorDstAddress == "" {
		return nil, fmt.Errorf("redelegate validator_dst_address required")
	}
	if msg.ValidatorSrcAddress == msg.ValidatorDstAddress {
		return nil, fmt.Errorf("redelegate src and dst validators must differ: %s", msg.ValidatorSrcAddress)
	}
	if msg.Amount.Amount.IsNil() || !msg.Amount.Amount.IsPositive() {
		return nil, fmt.Errorf("redelegate amount must be > 0")
	}

	switch paramName {
	case "delegator_address":
		return msg.DelegatorAddress, nil
	case "validator_src_address":
		return msg.ValidatorSrcAddress, nil
	case "validator_dst_address":
		return msg.ValidatorDstAddress, nil
	case "amount":
		return msg.Amount.Amount.BigInt(), nil
	case "denom":
		return msg.Amount.Denom, nil
	default:
		return nil, fmt.Errorf("unsupported parameter: %s", paramName)
	}
}

// extractParameterFromMsgWithdrawDelegatorReward extracts parameters from
// distribution MsgWithdrawDelegatorReward.
//
// This message has no amount field — the chain pays out whatever rewards have accrued
// from the delegator/validator pair. Policy rules can therefore only constrain by
// delegator and validator address (e.g., validator allowlist).
func (e *Engine) extractParameterFromMsgWithdrawDelegatorReward(paramName string, msg *distributiontypes.MsgWithdrawDelegatorReward) (any, error) {
	switch paramName {
	case "delegator_address":
		return msg.DelegatorAddress, nil
	case "validator_address":
		return msg.ValidatorAddress, nil
	default:
		return nil, fmt.Errorf("unsupported parameter: %s", paramName)
	}
}

// assertArgsByType validates constraints using the appropriate comparator based on Go type.
func (e *Engine) assertArgsByType(chainId, inputName string, arg any, constraints []*types.ParameterConstraint) error {
	switch actual := arg.(type) {
	case string:
		err := compare.AssertArg(
			chainId,
			constraints,
			inputName,
			actual,
			compare.NewString,
		)
		if err != nil {
			return fmt.Errorf("failed to assert string parameter: %w", err)
		}

	case *big.Int:
		err := compare.AssertArg(
			chainId,
			constraints,
			inputName,
			actual,
			compare.NewBigInt,
		)
		if err != nil {
			return fmt.Errorf("failed to assert big.Int parameter: %w", err)
		}

	default:
		return fmt.Errorf("unsupported parameter type: %T", actual)
	}
	return nil
}

// ExtractTxBytes extracts transaction bytes from a base64-encoded Cosmos transaction.
func (e *Engine) ExtractTxBytes(txData string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(txData)
}

