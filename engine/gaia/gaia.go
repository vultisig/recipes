package gaia

import (
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	chaingaia "github.com/vultisig/recipes/chain/gaia"
	stdcompare "github.com/vultisig/recipes/engine/compare"
	"github.com/vultisig/recipes/resolver"
	vtypes "github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
)

// Gaia represents the Cosmos (GAIA) engine implementation
type Gaia struct {
	chain *chaingaia.Chain
	cdc   codec.Codec
}

// NewGaia creates a new Gaia engine instance
func NewGaia() *Gaia {
	ir := codectypes.NewInterfaceRegistry()

	// Register crypto types (required for PubKey interfaces)
	cryptocodec.RegisterInterfaces(ir)

	// Register bank message types
	banktypes.RegisterInterfaces(ir)

	return &Gaia{
		chain: chaingaia.NewChain(),
		cdc:   codec.NewProtoCodec(ir),
	}
}

// Supports returns true if this engine supports the given chain
func (g *Gaia) Supports(chain common.Chain) bool {
	return chain == common.GaiaChain
}

// Evaluate validates a Cosmos transaction against policy rules
func (g *Gaia) Evaluate(rule *vtypes.Rule, txBytes []byte) error {
	// Validate rule effect is ALLOW (following existing pattern from other engines)
	if rule.GetEffect().String() != vtypes.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
	}

	// Parse resource to extract protocol and function information
	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}

	txData, err := g.parseTransaction(txBytes)
	if err != nil {
		return fmt.Errorf("failed to parse Cosmos transaction: %w", err)
	}

	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return fmt.Errorf("transaction must have at least one message")
	}

	if len(txData.Body.Messages) != 1 {
		return fmt.Errorf("only single-message transactions supported, got %d messages", len(txData.Body.Messages))
	}

	// Validate message type is supported (only MsgSend for GAIA)
	msg := txData.Body.Messages[0]
	mt, err := g.detectMessageType(msg)
	if err != nil {
		return fmt.Errorf("unsupported message type: %w", err)
	}

	// Enforce resource ↔ message-type compatibility
	if err := g.ensureResourceMessageCompatibility(r, mt); err != nil {
		return err
	}

	if err := g.validateTarget(r, rule.GetTarget(), txData); err != nil {
		return fmt.Errorf("failed to validate target: %w", err)
	}

	if err := g.validateParameterConstraints(r, rule.GetParameterConstraints(), txData); err != nil {
		return fmt.Errorf("failed to validate parameter constraints: %w", err)
	}

	return nil
}

// parseTransaction parses Cosmos transaction bytes into a Cosmos SDK transaction
func (g *Gaia) parseTransaction(txBytes []byte) (*tx.Tx, error) {
	const maxTxBytes = 32 * 1024 // 32 KB
	if len(txBytes) > maxTxBytes {
		return nil, fmt.Errorf("transaction too large: %d bytes (max %d)", len(txBytes), maxTxBytes)
	}

	if len(txBytes) == 0 {
		return nil, fmt.Errorf("empty transaction data")
	}

	var txData tx.Tx
	if err := g.cdc.Unmarshal(txBytes, &txData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal protobuf transaction: %w", err)
	}

	return &txData, nil
}

// ensureResourceMessageCompatibility rejects mismatched resource/message combinations.
func (g *Gaia) ensureResourceMessageCompatibility(resource *vtypes.ResourcePath, mt chaingaia.MessageType) error {
	switch resource.GetProtocolId() {
	case "atom", "send":
		if mt != chaingaia.MessageTypeSend {
			return fmt.Errorf("resource %s.%s only allows MsgSend, got %s", resource.GetProtocolId(), resource.GetFunctionId(), mt)
		}
	default:
		return fmt.Errorf("unsupported protocol: %s", resource.GetProtocolId())
	}
	return nil
}

// unpackMsgSend unpacks a message to bank MsgSend type
func (g *Gaia) unpackMsgSend(msg *codectypes.Any) (*banktypes.MsgSend, error) {
	if msg == nil {
		return nil, fmt.Errorf("nil message")
	}

	var sdkMsg sdk.Msg
	if err := g.cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return nil, fmt.Errorf("failed to unpack sdk.Msg: %w (typeUrl=%s)", err, msg.TypeUrl)
	}

	msgSend, ok := sdkMsg.(*banktypes.MsgSend)
	if !ok {
		return nil, fmt.Errorf("expected bank MsgSend, got: %T", sdkMsg)
	}

	return msgSend, nil
}

// detectMessageType determines the message type based on TypeUrl
func (g *Gaia) detectMessageType(msg *codectypes.Any) (chaingaia.MessageType, error) {
	if msg == nil {
		return chaingaia.MessageType(0), fmt.Errorf("nil message")
	}

	messageType, err := chaingaia.GetMessageTypeFromTypeUrl(msg.TypeUrl)
	if err != nil {
		return chaingaia.MessageType(0), fmt.Errorf("failed to detect message type: %w", err)
	}

	return messageType, nil
}

// validateTarget validates the transaction target against the rule target
func (g *Gaia) validateTarget(resource *vtypes.ResourcePath, target *vtypes.Target, txData *tx.Tx) error {
	if target == nil || target.GetTargetType() == vtypes.TargetType_TARGET_TYPE_UNSPECIFIED {
		return nil // No target validation required
	}

	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return fmt.Errorf("no messages in transaction")
	}

	msg := txData.Body.Messages[0]
	msgSend, err := g.unpackMsgSend(msg)
	if err != nil {
		return fmt.Errorf("failed to unpack message for target validation: %w", err)
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
func (g *Gaia) validateParameterConstraints(resource *vtypes.ResourcePath, constraints []*vtypes.ParameterConstraint, txData *tx.Tx) error {
	for _, constraint := range constraints {
		paramName := constraint.GetParameterName()

		value, err := g.extractParameterValue(paramName, txData)
		if err != nil {
			return fmt.Errorf("failed to extract parameter %s: %w", paramName, err)
		}

		if err := g.assertArgsByType(resource.ChainId, paramName, value, constraints); err != nil {
			return fmt.Errorf("constraint validation failed for parameter %s: %w", paramName, err)
		}
	}
	return nil
}

// extractParameterValue extracts the actual value from transaction for the given parameter name
func (g *Gaia) extractParameterValue(paramName string, txData *tx.Tx) (any, error) {
	if txData.Body == nil || len(txData.Body.Messages) == 0 {
		return nil, fmt.Errorf("no messages in transaction")
	}

	msg := txData.Body.Messages[0]
	msgSend, err := g.unpackMsgSend(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack message: %w", err)
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
func (g *Gaia) assertArgsByType(chainId, inputName string, arg any, constraints []*vtypes.ParameterConstraint) error {
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

