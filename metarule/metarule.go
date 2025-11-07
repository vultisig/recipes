package metarule

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/vultisig/recipes/metarule/internal/thorchain"
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

func getOneInchSpender(chain common.Chain) (string, error) {
	oneInchSpenders := map[common.Chain]string{
		common.Ethereum:  "0x111111125421ca6dc452d289314280a0f8842a65",
		common.Arbitrum:  "0x111111125421ca6dc452d289314280a0f8842a65",
		common.Avalanche: "0x111111125421ca6dc452d289314280a0f8842a65",
		common.BscChain:  "0x111111125421ca6dc452d289314280a0f8842a65",
		common.Base:      "0x111111125421ca6dc452d289314280a0f8842a65",
		common.Optimism:  "0x111111125421ca6dc452d289314280a0f8842a65",
		common.Polygon:   "0x111111125421ca6dc452d289314280a0f8842a65",
		common.Zksync:    "0x6fd4383cb451173d5f9304f041c7bcbf27d561ff",
	}

	spender, ok := oneInchSpenders[chain]
	if !ok {
		return "", fmt.Errorf("no 1inch spender address configured for chain: %s", chain.String())
	}
	return spender, nil
}

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
	case chain == common.Bitcoin:
		out, er := m.handleBitcoin(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle bitcoin: %w", er)
		}
		return out, nil
	case chain == common.XRP:
		out, er := m.handleXRP(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle xrp: %w", er)
		}
		return out, nil
	case chain == common.THORChain:
		out, er := m.handleTHORChain(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle thorchain: %w", er)
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

type swapConstraints struct {
	fromAsset   *types.Constraint
	fromAddress *types.Constraint
	fromAmount  *types.Constraint
	toChain     *types.Constraint
	toAsset     *types.Constraint
	toAddress   *types.Constraint
}

func getSwapConstraints(rule *types.Rule) (swapConstraints, error) {
	res := swapConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "from_asset":
			res.fromAsset = c.GetConstraint()
		case "from_address":
			res.fromAddress = c.GetConstraint()
		case "from_amount":
			res.fromAmount = c.GetConstraint()
		case "to_chain":
			res.toChain = c.GetConstraint()
		case "to_asset":
			res.toAsset = c.GetConstraint()
		case "to_address":
			res.toAddress = c.GetConstraint()
		}
	}

	if res.fromAsset == nil {
		return res, fmt.Errorf("failed to find constraint: from_asset")
	}
	if res.fromAddress == nil {
		return res, fmt.Errorf("failed to find constraint: from_address")
	}
	if res.fromAmount == nil {
		return res, fmt.Errorf("failed to find constraint: from_amount")
	}
	if res.toChain == nil {
		return res, fmt.Errorf("failed to find constraint: to_chain")
	}
	if res.toAsset == nil {
		return res, fmt.Errorf("failed to find constraint: to_asset")
	}
	if res.toAddress == nil {
		return res, fmt.Errorf("failed to find constraint: to_address")
	}

	return res, nil
}

type sendConstraints struct {
	asset       *types.Constraint
	fromAddress *types.Constraint
	amount      *types.Constraint
	toAddress   *types.Constraint
	// memo not supported yet
}

func getSendConstraints(rule *types.Rule) (sendConstraints, error) {
	res := sendConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "asset":
			res.asset = c.GetConstraint()
		case "from_address":
			res.fromAddress = c.GetConstraint()
		case "amount":
			res.amount = c.GetConstraint()
		case "to_address":
			res.toAddress = c.GetConstraint()
			// memo not supported yet
		}
	}

	if res.asset == nil {
		return res, fmt.Errorf("failed to find constraint: asset")
	}
	if res.fromAddress == nil {
		return res, fmt.Errorf("failed to find constraint: from_address")
	}
	if res.amount == nil {
		return res, fmt.Errorf("failed to find constraint: amount")
	}
	if res.toAddress == nil {
		return res, fmt.Errorf("failed to find constraint: to_address")
	}

	return res, nil
}

type sendUtxoConstraints struct {
	changeAddress *types.Constraint
	recipient     *types.Constraint
	amount        *types.Constraint
}

func getSendUtxoConstraints(rule *types.Rule) (sendUtxoConstraints, error) {
	res := sendUtxoConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "change_address":
			res.changeAddress = c.GetConstraint()
		case "recipient":
			res.recipient = c.GetConstraint()
		case "amount":
			res.amount = c.GetConstraint()
		}
	}

	if res.changeAddress == nil {
		return res, fmt.Errorf("failed to find constraint: change_address")
	}
	if res.recipient == nil {
		return res, fmt.Errorf("failed to find constraint: recipient")
	}
	if res.amount == nil {
		return res, fmt.Errorf("failed to find constraint: amount")
	}

	return res, nil
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
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse `send` constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)

		if c.asset.GetFixedValue() == "" {
			// native transfer
			var outTarget *types.Target
			switch c.toAddress.GetType() {
			case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: c.toAddress.GetFixedValue(),
					},
				}
			case types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
					Target: &types.Target_MagicConstant{
						MagicConstant: c.toAddress.GetMagicConstantValue(),
					},
				}
			default:
				return nil, fmt.Errorf(
					"invalid constraint type for `to_address`: %s",
					c.toAddress.GetType().String(),
				)
			}

			out.Resource = "solana.system.transfer"
			out.Target = outTarget
			out.ParameterConstraints = []*types.ParameterConstraint{{
				ParameterName: "account_from",
				Constraint:    c.fromAddress,
			}, {
				ParameterName: "account_to",
				Constraint:    c.toAddress,
			}, {
				ParameterName: "arg_lamports",
				Constraint:    c.amount,
			}}
			return []*types.Rule{out}, nil
		}

		const onlyFixed = "must be fixed constraint for spl token transfer"
		if c.fromAddress.GetFixedValue() == "" {
			return nil, fmt.Errorf("`from_address` " + onlyFixed)
		}
		if c.toAddress.GetFixedValue() == "" {
			return nil, fmt.Errorf("`to_address` " + onlyFixed)
		}

		src, err := DeriveATA(c.fromAddress.GetFixedValue(), c.asset.GetFixedValue())
		if err != nil {
			return nil, fmt.Errorf("failed to derive src ATA: %w", err)
		}
		dst, err := DeriveATA(c.toAddress.GetFixedValue(), c.asset.GetFixedValue())
		if err != nil {
			return nil, fmt.Errorf("failed to derive dst ATA: %w", err)
		}

		// SPL token transfer
		out.Resource = "solana.spl_token.transfer"
		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "account_source",
			Constraint:    fixed(src),
		}, {
			ParameterName: "account_destination",
			Constraint:    fixed(dst),
		}, {
			ParameterName: "account_authority",
			Constraint:    c.fromAddress,
		}, {
			ParameterName: "arg_amount",
			Constraint:    c.amount,
		}}
		return []*types.Rule{out}, nil
	case swap:
		c, err := getSwapConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse `swap` constraints: %w", err)
		}

		rules, err := m.createJupiterRule(in, c)
		if err != nil {
			return nil, fmt.Errorf("failed to create jupiter rules: %w", err)
		}
		return rules, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id: %s", r.GetProtocolId())
	}
}

func (m *MetaRule) handleEVM(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse `send` constraints: %w", err)
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

		if c.asset.GetFixedValue() == "" {
			var outTarget *types.Target
			switch c.toAddress.GetType() {
			case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: c.toAddress.GetFixedValue(),
					},
				}
			case types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
					Target: &types.Target_MagicConstant{
						MagicConstant: c.toAddress.GetMagicConstantValue(),
					},
				}
			default:
				return nil, fmt.Errorf(
					"invalid constraint type for `to_address`: %s",
					c.toAddress.GetType().String(),
				)
			}

			out.Resource = fmt.Sprintf(
				"%s.%s.transfer",
				strings.ToLower(chain.String()),
				strings.ToLower(nativeSymbol),
			)
			out.Target = outTarget
			out.ParameterConstraints = []*types.ParameterConstraint{{
				ParameterName: "amount",
				Constraint:    c.amount,
			}}

			return []*types.Rule{out}, nil
		}

		// erc20 token transfer
		out.Resource = fmt.Sprintf("%s.erc20.transfer", strings.ToLower(chain.String()))
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: c.asset.GetFixedValue(),
			},
		}
		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "recipient",
			Constraint:    c.toAddress,
		}, {
			ParameterName: "amount",
			Constraint:    c.amount,
		}}
		return []*types.Rule{out}, nil
	case swap:
		c, err := getSwapConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse swap constraints: %w", err)
		}

		chain, err := common.FromString(r.GetChainId())
		if err != nil {
			return nil, fmt.Errorf("invalid chainID: %w", err)
		}

		rules := make([]*types.Rule, 0)

		var router *types.Constraint
		if r.GetChainId() == strings.ToLower(c.toChain.GetFixedValue()) {
			// same chain - 1inch
			oneinchRouter, out, er := oneinchSwap(chain, c)
			if er != nil {
				return nil, fmt.Errorf("failed to create oneinch swap rule: %w", er)
			}

			rules = append(rules, out)
			router = fixed(oneinchRouter)
		} else {
			// cross-chain - ThorChain
			// here we don't care is a bridge direction supported — it's a plugin responsibility
			// we build a safe ThorChain rule mapping to the swap request
			out, er := thorchainSwap(chain, c)
			if er != nil {
				return nil, fmt.Errorf("failed to create thorchain swap rule: %w", er)
			}

			rules = append(rules, out)
			router = &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT,
				Value: &types.Constraint_MagicConstantValue{
					MagicConstantValue: types.MagicConstant_THORCHAIN_ROUTER,
				},
			}
		}

		if c.fromAsset.GetFixedValue() != "" {
			approve := proto.Clone(in).(*types.Rule)
			approve.Resource = fmt.Sprintf("%s.erc20.approve", strings.ToLower(chain.String()))
			approve.Target = &types.Target{
				TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
				Target: &types.Target_Address{
					Address: c.fromAsset.GetFixedValue(),
				},
			}
			approve.ParameterConstraints = []*types.ParameterConstraint{
				{
					ParameterName: "amount",
					Constraint: &types.Constraint{
						Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
						Value: &types.Constraint_MinValue{
							MinValue: c.fromAmount.GetFixedValue(),
						},
					},
				},
				{
					ParameterName: "spender",
					Constraint:    router,
				},
			}
			rules = append(rules, approve)
		}

		return rules, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id: %s", r.GetProtocolId())
	}
}

func thorchainSwap(chain common.Chain, c swapConstraints) (*types.Rule, error) {
	asset := c.fromAsset.GetFixedValue()
	amount := c.fromAmount
	if asset == "" {
		asset = evm.ZeroAddress.String()
		amount = fixed("0")
	}

	chainInt, err := common.FromString(c.toChain.GetFixedValue())
	if err != nil {
		return nil, fmt.Errorf("failed to parse chain id: %w", err)
	}

	thorAsset, err := thorchain.MakeAsset(chainInt, c.toAsset.GetFixedValue())
	if err != nil {
		return nil, fmt.Errorf("failed to make thor asset: %w", err)
	}

	shortCode := thorchain.ShortCode(thorAsset)
	var assetPattern string
	if shortCode != "" {
		assetPattern = fmt.Sprintf("(%s|%s)",
			regexp.QuoteMeta(thorAsset),
			regexp.QuoteMeta(shortCode))
	} else {
		assetPattern = regexp.QuoteMeta(thorAsset)
	}

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: fmt.Sprintf("%s.thorchain_router.depositWithExpiry", strings.ToLower(chain.String())),
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
			Target: &types.Target_MagicConstant{
				MagicConstant: types.MagicConstant_THORCHAIN_ROUTER,
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "vault",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT,
					Value: &types.Constraint_MagicConstantValue{
						MagicConstantValue: types.MagicConstant_THORCHAIN_VAULT,
					},
				},
			},
			{
				ParameterName: "asset",
				Constraint:    fixed(asset),
			},
			{
				ParameterName: "amount",
				Constraint:    amount,
			},
			{
				ParameterName: "memo",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_REGEXP,
					Value: &types.Constraint_RegexpValue{
						// =:<asset>:<address>:<optional_params>
						// The = is shorthand for the SWAP command
						// Validates the destination asset in ThorChain notation
						// Validates the destination address
						// Allows optional parameters (streaming options, min amount, affiliate, etc.)
						RegexpValue: fmt.Sprintf(
							"^=:%s:%s:.*",
							assetPattern,
							regexp.QuoteMeta(c.toAddress.GetFixedValue()),
						),
					},
				},
			},
			{
				ParameterName: "expiration",
				Constraint:    anyConstraint(),
			},
		},
	}

	return rule, nil
}

func oneinchSwap(chain common.Chain, c swapConstraints) (string, *types.Rule, error) {
	const oneinchNative = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"

	oneinchRouterAddr, err := getOneInchSpender(chain)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get 1inch spender: %w", err)
	}

	srcToken := c.fromAsset.GetFixedValue()
	if c.fromAsset.GetFixedValue() == "" {
		srcToken = oneinchNative
	}

	dstToken := c.toAsset.GetFixedValue()
	if c.toAsset.GetFixedValue() == "" {
		dstToken = oneinchNative
	}

	out := &types.Rule{
		Effect: types.Effect_EFFECT_ALLOW,
	}
	out.Resource = fmt.Sprintf("%s.routerV6_1inch.swap", strings.ToLower(chain.String()))
	out.Target = &types.Target{
		TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
		Target: &types.Target_Address{
			Address: oneinchRouterAddr,
		},
	}
	out.ParameterConstraints = []*types.ParameterConstraint{
		{
			ParameterName: "executor",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
			},
		},
		{
			ParameterName: "desc.srcToken",
			Constraint:    fixed(srcToken),
		},
		{
			ParameterName: "desc.dstToken",
			Constraint:    fixed(dstToken),
		},
		{
			ParameterName: "desc.srcReceiver",
			Constraint:    anyConstraint(),
		},
		{
			ParameterName: "desc.dstReceiver",
			Constraint:    c.toAddress,
		},
		{
			ParameterName: "desc.amount",
			Constraint:    c.fromAmount,
		},
		{
			ParameterName: "desc.minReturnAmount",
			Constraint:    anyConstraint(),
		},
		{
			ParameterName: "desc.flags",
			Constraint:    anyConstraint(),
		},
		{
			ParameterName: "data",
			Constraint:    anyConstraint(),
		}}

	return oneinchRouterAddr, out, nil
}

func (m *MetaRule) handleBitcoin(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		c, err := getSendUtxoConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse `send utxo` constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "bitcoin.btc.transfer"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_UNSPECIFIED,
		}

		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "output_address_0",
			Constraint:    c.recipient,
		}, {
			ParameterName: "output_value_0",
			Constraint:    c.amount,
		}, {
			ParameterName: "output_address_1",
			Constraint:    c.changeAddress,
		}, {
			ParameterName: "output_value_1",
			Constraint:    anyConstraint(),
		}}
		return []*types.Rule{out}, nil
	case swap:
		c, err := getSwapConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse swap constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "bitcoin.btc.transfer"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_UNSPECIFIED,
		}

		chainInt, err := common.FromString(c.toChain.GetFixedValue())
		if err != nil {
			return nil, fmt.Errorf("failed to parse chain id: %w", err)
		}

		thorAsset, err := thorchain.MakeAsset(chainInt, c.toAsset.GetFixedValue())
		if err != nil {
			return nil, fmt.Errorf("failed to make thor asset: %w", err)
		}

		// Create asset pattern that accepts both full form and shortform
		shortCode := thorchain.ShortCode(thorAsset)
		var assetPattern string
		if shortCode != "" {
			// Accept both full form and shortform: (BTC\.BTC|b)
			assetPattern = fmt.Sprintf("(%s|%s)",
				regexp.QuoteMeta(thorAsset),
				regexp.QuoteMeta(shortCode))
		} else {
			// Fallback to full asset name only
			assetPattern = regexp.QuoteMeta(thorAsset)
		}

		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "output_address_0",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT,
				Value: &types.Constraint_MagicConstantValue{
					MagicConstantValue: types.MagicConstant_THORCHAIN_VAULT,
				},
			},
		}, {
			ParameterName: "output_value_0",
			Constraint:    c.fromAmount,
		}, {
			ParameterName: "output_address_1",
			Constraint:    c.fromAddress, // change
		}, {
			ParameterName: "output_value_1",
			Constraint:    anyConstraint(), // change
		}, {
			ParameterName: "output_data_2",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_REGEXP,
				Value: &types.Constraint_RegexpValue{
					RegexpValue: fmt.Sprintf(
						"^=:%s:%s:.*", // swap_command:asset:address:any(streaming options, min amount out, etc.)
						assetPattern,
						regexp.QuoteMeta(c.toAddress.GetFixedValue()),
					),
				},
			},
		}}
		return []*types.Rule{out}, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id: %s", r.GetProtocolId())
	}
}

func (m *MetaRule) handleXRP(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse send constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "ripple.xrp.transfer"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_UNSPECIFIED,
		}

		out.ParameterConstraints = []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint:    c.toAddress,
			},
			{
				ParameterName: "amount",
				Constraint:    c.amount,
			},
		}

		return []*types.Rule{out}, nil
	case swap:
		c, err := getSwapConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse swap constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "ripple.swap"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
			Target: &types.Target_MagicConstant{
				MagicConstant: types.MagicConstant_THORCHAIN_VAULT,
			},
		}

		chainInt, err := common.FromString(c.toChain.GetFixedValue())
		if err != nil {
			return nil, fmt.Errorf("failed to parse chain id: %w", err)
		}

		thorAsset, err := thorchain.MakeAsset(chainInt, c.toAsset.GetFixedValue())
		if err != nil {
			return nil, fmt.Errorf("failed to make thor asset: %w", err)
		}

		// Create asset pattern that accepts both full form and shortform
		shortCode := thorchain.ShortCode(thorAsset)
		var assetPattern string
		if shortCode != "" {
			// Accept both full form and shortform: (BTC\.BTC|b)
			assetPattern = fmt.Sprintf("(%s|%s)",
				regexp.QuoteMeta(thorAsset),
				regexp.QuoteMeta(shortCode))
		} else {
			// Fallback to full asset name only
			assetPattern = regexp.QuoteMeta(thorAsset)
		}

		out.ParameterConstraints = []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT,
					Value: &types.Constraint_MagicConstantValue{
						MagicConstantValue: types.MagicConstant_THORCHAIN_VAULT,
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint:    c.fromAmount,
			},
			{
				ParameterName: "memo",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_REGEXP,
					Value: &types.Constraint_RegexpValue{
						RegexpValue: fmt.Sprintf(
							"^=:%s:%s:.*", // swap_command:asset:address:any(streaming options, min amount out, etc.)
							assetPattern,
							regexp.QuoteMeta(c.toAddress.GetFixedValue()),
						),
					},
				},
			},
		}

		return []*types.Rule{out}, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id for XRP: %s", r.GetProtocolId())
	}
}

func (m *MetaRule) handleTHORChain(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse send constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "thorchain.rune.transfer"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_UNSPECIFIED,
		}

		out.ParameterConstraints = []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint:    c.toAddress,
			},
			{
				ParameterName: "amount",
				Constraint:    c.amount,
			},
		}

		return []*types.Rule{out}, nil
	case swap:
		c, err := getSwapConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse swap constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "thorchain.thorchain_swap.swap"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_UNSPECIFIED,
		}

		chainInt, err := common.FromString(c.toChain.GetFixedValue())
		if err != nil {
			return nil, fmt.Errorf("failed to parse chain id: %w", err)
		}

		thorAsset, err := thorchain.MakeAsset(chainInt, c.toAsset.GetFixedValue())
		if err != nil {
			return nil, fmt.Errorf("failed to make thor asset: %w", err)
		}

		// Create asset pattern that accepts both full form and shortform
		shortCode := thorchain.ShortCode(thorAsset)
		var assetPattern string
		if shortCode != "" {
			assetPattern = fmt.Sprintf("(%s|%s)",
				regexp.QuoteMeta(thorAsset),
				regexp.QuoteMeta(shortCode))
		} else {
			assetPattern = regexp.QuoteMeta(thorAsset)
		}

		out.ParameterConstraints = []*types.ParameterConstraint{
			{
				ParameterName: "amount",
				Constraint:    c.fromAmount,
			},
			{
				ParameterName: "denom",
				Constraint:    c.fromAsset,
			},
			{
				ParameterName: "memo",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_REGEXP,
					Value: &types.Constraint_RegexpValue{
						RegexpValue: fmt.Sprintf(
							"^=:%s:%s:.*",
							assetPattern,
							regexp.QuoteMeta(c.toAddress.GetFixedValue()),
						),
					},
				},
			},
		}

		return []*types.Rule{out}, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id for THORChain: %s", r.GetProtocolId())
	}
}

func getFixed(c *types.Constraint) (string, error) {
	if c.GetType() != types.ConstraintType_CONSTRAINT_TYPE_FIXED {
		return "", fmt.Errorf("invalid constraint type: %s", c.GetType())
	}

	return c.GetFixedValue(), nil
}

// DeriveATA derives the Associated Token Account address for a given owner and mint
func DeriveATA(ownerStr, mintStr string) (string, error) {
	owner, err := solana.PublicKeyFromBase58(ownerStr)
	if err != nil {
		return "", fmt.Errorf("invalid owner address: %w", err)
	}

	mint, err := solana.PublicKeyFromBase58(mintStr)
	if err != nil {
		return "", fmt.Errorf("invalid mint address: %w", err)
	}

	ataAddr, _, err := solana.FindProgramAddress(
		[][]byte{
			owner.Bytes(),
			solana.TokenProgramID.Bytes(),
			mint.Bytes(),
		},
		solana.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"),
	)
	if err != nil {
		return "", fmt.Errorf("failed to derive ATA: %w", err)
	}

	return ataAddr.String(), nil
}

func (m *MetaRule) createJupiterRule(_ *types.Rule, c swapConstraints) ([]*types.Rule, error) {
	toChainStr, err := getFixed(c.toChain)
	if err != nil {
		return nil, fmt.Errorf("failed to get fixed value for toChain: %w", err)
	}
	if !strings.EqualFold(toChainStr, common.Solana.String()) {
		return nil, fmt.Errorf("only solana->solana allowed for jupiter, got toChain: %q", toChainStr)
	}

	const (
		jupAddr  = "JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4"
		jupEvent = "D8cy77BBepLMngZx6ZukaTff5hCt1HrWyKk3Hnd9oitf"
	)

	fromAssetStr, err := getFixed(c.fromAsset)
	if err != nil {
		return nil, fmt.Errorf("failed to get fixed value: %w", err)
	}

	toAssetStr, err := getFixed(c.toAsset)
	if err != nil {
		return nil, fmt.Errorf("failed to get fixed value for toAsset: %w", err)
	}

	fromAddressStr, err := getFixed(c.fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get fixed value for fromAddress: %w", err)
	}

	toAddressStr, err := getFixed(c.toAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get fixed value for toAddress: %w", err)
	}

	getMint := func(asset string) string {
		if asset == "" {
			return solana.SolMint.String()
		}
		return asset
	}

	sourceMint := getMint(fromAssetStr)
	destMint := getMint(toAssetStr)

	sourceMintConstraint := c.fromAsset
	if fromAssetStr == "" {
		sourceMintConstraint = fixed(sourceMint)
	}

	destinationMintConstraint := c.toAsset
	if toAssetStr == "" {
		destinationMintConstraint = fixed(destMint)
	}

	sourceATA, err := DeriveATA(fromAddressStr, sourceMint)
	if err != nil {
		return nil, fmt.Errorf("failed to derive source ATA for owner %s and mint %s: %w", fromAddressStr, sourceMint, err)
	}

	destATA, err := DeriveATA(toAddressStr, destMint)
	if err != nil {
		return nil, fmt.Errorf("failed to derive destination ATA for owner %s and mint %s: %w", toAddressStr, destMint, err)
	}

	sourceTokenAccountConstraint := fixed(sourceATA)
	destinationTokenAccountConstraint := fixed(destATA)

	createSystemTransferRule := func(to, amount *types.Constraint) *types.Rule {
		return &types.Rule{
			Resource: "solana.system.transfer",
			Effect:   types.Effect_EFFECT_ALLOW,
			ParameterConstraints: []*types.ParameterConstraint{{
				ParameterName: "account_from",
				Constraint:    c.fromAddress,
			}, {
				ParameterName: "account_to",
				Constraint:    to,
			}, {
				ParameterName: "arg_lamports",
				Constraint:    amount,
			}},
			Target: &types.Target{
				TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
				Target: &types.Target_Address{
					Address: solana.SystemProgramID.String(),
				},
			},
		}
	}

	var rules []*types.Rule

	// Allow System Program transfer for funding rent-exempt accounts and wrapping SOL
	// For native SOL swaps, Jupiter wraps SOL into WSOL by transferring to the WSOL ATA
	// Typical amounts: rent ~2,039,280 lamports + wrapped SOL amount for swap
	rules = append(rules,
		createSystemTransferRule(sourceTokenAccountConstraint, c.fromAmount),
		createSystemTransferRule(destinationTokenAccountConstraint, &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
			Value: &types.Constraint_MaxValue{
				MaxValue: "10000000",
			},
		}),
	)

	createATARule := func(payer, ata, owner, mint *types.Constraint) *types.Rule {
		return &types.Rule{
			Resource: "solana.associated_token_account.create",
			Effect:   types.Effect_EFFECT_ALLOW,
			ParameterConstraints: []*types.ParameterConstraint{{
				ParameterName: "account_payer",
				Constraint:    payer,
			}, {
				ParameterName: "account_associated_token_account",
				Constraint:    ata,
			}, {
				ParameterName: "account_owner",
				Constraint:    owner,
			}, {
				ParameterName: "account_mint",
				Constraint:    mint,
			}, {
				ParameterName: "account_system_program",
				Constraint:    fixed(solana.SystemProgramID.String()),
			}, {
				ParameterName: "account_token_program",
				Constraint:    fixed(solana.TokenProgramID.String()),
			}},
			Target: &types.Target{
				TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
				Target: &types.Target_Address{
					Address: "ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL",
				},
			},
		}
	}

	// Add create for source and destination ATAs
	// Jupiter requires wrapped SOL token accounts even for native SOL swaps
	rules = append(rules,
		createATARule(c.fromAddress, sourceTokenAccountConstraint, c.fromAddress, sourceMintConstraint),
		createATARule(c.fromAddress, destinationTokenAccountConstraint, c.toAddress, destinationMintConstraint),
	)

	// SPL Token syncNative - syncs native SOL balance in WSOL account after wrapping
	// This is required after transferring SOL to a WSOL ATA to update the token account balance
	rules = append(rules, &types.Rule{
		Resource: "solana.spl_token.syncNative",
		Effect:   types.Effect_EFFECT_ALLOW,
		ParameterConstraints: []*types.ParameterConstraint{{
			ParameterName: "account_account",
			Constraint:    sourceTokenAccountConstraint,
		}},
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.TokenProgramID.String(),
			},
		},
	})

	// SPL Token Approve - account_source must be an Associated Token Account (ATA)
	// The account_source constraint ensures the approved token account is derived from the ATA program
	rules = append(rules, &types.Rule{
		Resource: "solana.spl_token.approve",
		Effect:   types.Effect_EFFECT_ALLOW,
		ParameterConstraints: []*types.ParameterConstraint{{
			ParameterName: "account_source",
			// sourceTokenAccountConstraint is derived using DeriveATA() which uses
			// the Associated Token Account Program (ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL)
			Constraint: sourceTokenAccountConstraint,
		}, {
			ParameterName: "account_delegate",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				Value: &types.Constraint_FixedValue{
					FixedValue: jupAddr,
				},
			},
		}, {
			ParameterName: "account_owner",
			Constraint:    c.fromAddress,
		}, {
			ParameterName: "arg_amount",
			Constraint:    c.fromAmount,
		}},
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				// Target is the Token Program that executes the approve instruction
				Address: solana.TokenProgramID.String(),
			},
		},
	})

	// SPL Token CloseAccount - closes wSOL account and unwraps back to native SOL
	rules = append(rules, &types.Rule{
		Resource: "solana.spl_token.closeAccount",
		Effect:   types.Effect_EFFECT_ALLOW,
		ParameterConstraints: []*types.ParameterConstraint{{
			ParameterName: "account_account",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "account_destination",
			Constraint:    c.fromAddress,
		}, {
			ParameterName: "account_owner",
			Constraint:    c.fromAddress,
		}},
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.TokenProgramID.String(),
			},
		},
	})

	// Jupiter route instruction (exact input amount)
	jupiterRouteRule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.jupiter_aggregatorv6.route",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: jupAddr,
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{{
			ParameterName: "account_token_program",
			Constraint:    fixed(solana.TokenProgramID.String()),
		}, {
			ParameterName: "account_user_transfer_authority",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "account_user_source_token_account",
			Constraint:    sourceTokenAccountConstraint,
		}, {
			ParameterName: "account_user_destination_token_account",
			Constraint:    destinationTokenAccountConstraint,
		}, {
			ParameterName: "account_destination_token_account",
			Constraint:    anyConstraint(), // Jupiter infrastructure
		}, {
			ParameterName: "account_destination_mint",
			Constraint:    destinationMintConstraint,
		}, {
			ParameterName: "account_platform_fee_account",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "account_event_authority",
			Constraint:    fixed(jupEvent),
		}, {
			ParameterName: "account_program",
			Constraint:    fixed(jupAddr),
		}, {
			ParameterName: "arg_route_plan",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "arg_slippage_bps",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "arg_platform_fee_bps",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "arg_in_amount",
			Constraint:    c.fromAmount,
		}, {
			ParameterName: "arg_quoted_out_amount",
			Constraint:    anyConstraint(),
		}},
	}

	// Jupiter shared_accounts_route instruction (optimized for multi-hop swaps)
	jupiterSharedAccountsRouteRule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.jupiter_aggregatorv6.shared_accounts_route",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: jupAddr,
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{{
			ParameterName: "account_token_program",
			Constraint:    fixed(solana.TokenProgramID.String()),
		}, {
			ParameterName: "account_program_authority",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "account_user_transfer_authority",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "account_source_token_account",
			Constraint:    sourceTokenAccountConstraint,
		}, {
			ParameterName: "account_program_source_token_account",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "account_program_destination_token_account",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "account_destination_token_account",
			Constraint:    destinationTokenAccountConstraint,
		}, {
			ParameterName: "account_source_mint",
			Constraint:    sourceMintConstraint,
		}, {
			ParameterName: "account_destination_mint",
			Constraint:    destinationMintConstraint,
		}, {
			ParameterName: "account_platform_fee_account",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "account_token_2022_program",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "account_event_authority",
			Constraint:    fixed(jupEvent),
		}, {
			ParameterName: "account_program",
			Constraint:    fixed(jupAddr),
		}, {
			ParameterName: "arg_id",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "arg_route_plan",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "arg_in_amount",
			Constraint:    c.fromAmount,
		}, {
			ParameterName: "arg_quoted_out_amount",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "arg_slippage_bps",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "arg_platform_fee_bps",
			Constraint:    anyConstraint(),
		}},
	}

	rules = append(rules, jupiterRouteRule, jupiterSharedAccountsRouteRule)

	return rules, nil
}

func fixed(in string) *types.Constraint {
	return &types.Constraint{
		Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
		Value: &types.Constraint_FixedValue{
			FixedValue: in,
		},
	}
}

func anyConstraint() *types.Constraint {
	return &types.Constraint{
		Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
	}
}
