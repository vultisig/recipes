package metarule

import (
	"fmt"

	"github.com/vultisig/recipes/internal/conv"
	"github.com/vultisig/recipes/solana"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
	"google.golang.org/protobuf/proto"
)

type MetaRule struct{}

func NewMetaRule() *MetaRule {
	return &MetaRule{}
}

// TryFormat meta-rule to exact rule. For example:
// solana.send -> solana.sol.transfer or solana.spl_token.transfer
// solana.sol.transfer -> unmodified solana.sol.transfer
// *.*.* (any 3 fields rule) -> unmodified *.*.*
func (m *MetaRule) TryFormat(in *types.Rule) (*types.Rule, error) {
	r, err := util.ParseResource(in.GetResource())
	if err != nil {
		return nil, fmt.Errorf("failed to parse resource: %w", err)
	}
	if r.GetFunctionId() != "" {
		// it's not a meta-rule
		return in, nil
	}

	chain, err := common.FromString(r.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse chain id: %w", err)
	}
	switch chain {
	case common.Solana:
		out, er := m.handleSolana(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle solana: %w", er)
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

func (m *MetaRule) handleSolana(in *types.Rule, r *types.ResourcePath) (*types.Rule, error) {
	switch r.GetProtocolId() {
	case "send":
		recipient, er := m.getConstraint(in, "recipient")
		if er != nil {
			return nil, fmt.Errorf("failed to parse `recipient`: %w", er)
		}

		amount, er := m.getConstraint(in, "amount")
		if er != nil {
			return nil, fmt.Errorf("failed to parse `amount`: %w", er)
		}

		out := proto.Clone(in)

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

			out.Resource = "solana.sol.transfer"
			out.Target = outTarget
			out.ParameterConstraints = []*types.ParameterConstraint{{
				ParameterName: "amount",
				Constraint:    amount,
			}}
			return out, nil
		}

		// SPL token transfer
		out.Resource = "solana.spl_token.transfer"
		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "destination",
			Constraint:    recipient,
		}, {
			ParameterName: "amount",
			Constraint:    amount,
		}, {
			ParameterName: "source",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "authority",
			Constraint:    anyConstraint(),
		}}
		return out, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id: %s", r.GetProtocolId())
	}
}

func anyConstraint() *types.Constraint {
	return &types.Constraint{
		Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
	}
}
