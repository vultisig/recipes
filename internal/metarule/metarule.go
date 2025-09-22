package metarule

import (
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/vultisig/recipes/sdk/evm"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
	"google.golang.org/protobuf/proto"
)

type MetaRule struct{}

func NewMetaRule() *MetaRule {
	return &MetaRule{}
}

type metaProtocol string

const (
	send metaProtocol = "send"
	swap metaProtocol = "swap"
)

// TryFormat meta-rule to exact rule(s). For example:
// solana.send -> solana.system.transfer or solana.spl_token.transfer
// solana.system.transfer -> unmodified solana.system.transfer
// *.*.* (any 3 fields rule) -> unmodified *.*.*
func (m *MetaRule) TryFormat(in *types.Rule) ([]*types.Rule, error) {
	r, err := util.ParseResource(in.GetResource())
	if err != nil {
		return nil, fmt.Errorf("failed to parse resource: %w", err)
	}
	if r.GetFunctionId() != "" {
		// it's not a meta-rule
		return []*types.Rule{in}, nil
	}

	chain, err := common.FromString(r.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse chain id: %w", err)
	}
	switch {
	case chain == common.Solana:
		out, er := m.handleSolana(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle solana: %w", er)
		}
		return out, nil
	case chain.IsEvm():
		out, er := m.handleEVM(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle evm: %w", er)
		}
		return out, nil
	default:
		return nil, fmt.Errorf(
			"got meta format (%s) but chain not supported: %s",
			in.GetResource(),
			chain.String(),
		)
	}
}

func (m *MetaRule) getConstraint(rule *types.Rule, name string) (*types.Constraint, error) {
	for _, c := range rule.GetParameterConstraints() {
		if c.GetParameterName() == name {
			return c.GetConstraint(), nil
		}
	}
	return nil, fmt.Errorf("failed to find constraint: %s", name)
}

func (m *MetaRule) handleSolana(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		recipient, er := m.getConstraint(in, "recipient")
		if er != nil {
			return nil, fmt.Errorf("failed to parse `recipient`: %w", er)
		}

		amount, er := m.getConstraint(in, "amount")
		if er != nil {
			return nil, fmt.Errorf("failed to parse `amount`: %w", er)
		}

		out := proto.Clone(in).(*types.Rule)

		if in.GetTarget().GetAddress() == solana.SystemProgramID.String() {
			// native transfer
			var outTarget *types.Target
			switch recipient.GetType() {
			case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: recipient.GetFixedValue(),
					},
				}
			case types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
					Target: &types.Target_MagicConstant{
						MagicConstant: recipient.GetMagicConstantValue(),
					},
				}
			default:
				return nil, fmt.Errorf(
					"invalid constraint type for `recipient`: %s",
					recipient.GetType().String(),
				)
			}

			out.Resource = "solana.system.transfer"
			out.Target = outTarget
			out.ParameterConstraints = []*types.ParameterConstraint{{
				ParameterName: "account_from",
				Constraint:    anyConstraint(),
			}, {
				ParameterName: "account_to",
				Constraint:    recipient,
			}, {
				ParameterName: "arg_lamports",
				Constraint:    amount,
			}}
			return []*types.Rule{out}, nil
		}

		// SPL token transfer
		out.Resource = "solana.spl_token.transfer"
		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "account_source",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "account_destination",
			Constraint:    recipient,
		}, {
			ParameterName: "account_authority",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "arg_amount",
			Constraint:    amount,
		}}
		return []*types.Rule{out}, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id: %s", r.GetProtocolId())
	}
}

func (m *MetaRule) handleEVM(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		recipient, err := m.getConstraint(in, "recipient")
		if err != nil {
			return nil, fmt.Errorf("failed to parse `recipient`: %w", err)
		}
		amount, err := m.getConstraint(in, "amount")
		if err != nil {
			return nil, fmt.Errorf("failed to parse `amount`: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)

		chain, err := common.FromString(r.GetChainId())
		if err != nil {
			return nil, fmt.Errorf("invalid chainID: %w", err)
		}

		nativeSymbol, err := chain.NativeSymbol()
		if err != nil {
			return nil, fmt.Errorf("failed to find native symbol: %w", err)
		}

		if in.GetTarget().GetAddress() == evm.ZeroAddress.String() {
			var outTarget *types.Target
			switch recipient.GetType() {
			case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: recipient.GetFixedValue(),
					},
				}
			case types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
					Target: &types.Target_MagicConstant{
						MagicConstant: recipient.GetMagicConstantValue(),
					},
				}
			default:
				return nil, fmt.Errorf(
					"invalid constraint type for `recipient`: %s",
					recipient.GetType().String(),
				)
			}

			out.Resource = fmt.Sprintf("%s.%s.transfer", strings.ToLower(chain.String()), nativeSymbol)
			out.Target = outTarget
			out.ParameterConstraints = []*types.ParameterConstraint{{
				ParameterName: "amount",
				Constraint:    amount,
			}}

			return []*types.Rule{out}, nil
		}

		// erc20 token transfer
		out.Resource = fmt.Sprintf("%s.erc20.transfer", strings.ToLower(chain.String()))
		out.Target = in.GetTarget()
		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "recipient",
			Constraint:    recipient,
		}, {
			ParameterName: "amount",
			Constraint:    amount,
		}}
		return []*types.Rule{out}, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id: %s", r.GetProtocolId())
	}
}

func anyConstraint() *types.Constraint {
	return &types.Constraint{
		Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
	}
}
