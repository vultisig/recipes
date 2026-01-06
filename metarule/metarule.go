package metarule

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gagliardetto/solana-go"
	chainEVM "github.com/vultisig/recipes/chain/evm"
	"github.com/vultisig/recipes/metarule/internal/mayachain"
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
	send   metaProtocol = "send"
	swap   metaProtocol = "swap"
	bridge metaProtocol = "bridge"
	lp     metaProtocol = "lp"
	lend   metaProtocol = "lend"
	perps  metaProtocol = "perps"
	bet    metaProtocol = "bet"
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
	case chain == common.Tron:
		out, er := m.handleTron(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle tron: %w", er)
		}
		return out, nil
	case chain == common.Litecoin:
		out, er := m.handleLitecoin(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle litecoin: %w", er)
		}
		return out, nil
	case chain == common.Dogecoin:
		out, er := m.handleDogecoin(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle dogecoin: %w", er)
		}
		return out, nil
	case chain == common.BitcoinCash:
		out, er := m.handleBitcoinCash(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle bitcoin cash: %w", er)
		}
		return out, nil
	case chain == common.Zcash:
		out, er := m.handleZcash(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle zcash: %w", er)
		}
		return out, nil
	case chain == common.GaiaChain:
		out, er := m.handleGaia(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle gaia: %w", er)
		}
		return out, nil
	case chain == common.MayaChain:
		out, er := m.handleMaya(in, r)
		if er != nil {
			return nil, fmt.Errorf("failed to handle maya: %w", er)
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

// bridgeConstraints holds parameters for bridge operations (same asset across chains)
type bridgeConstraints struct {
	fromAsset   *types.Constraint // Token address on source chain (empty for native)
	fromAddress *types.Constraint // Sender address
	fromAmount  *types.Constraint // Amount to bridge
	toChain     *types.Constraint // Destination chain
	toAsset     *types.Constraint // Token address on destination (empty for native)
	toAddress   *types.Constraint // Recipient address on destination
}

func getBridgeConstraints(rule *types.Rule) (bridgeConstraints, error) {
	res := bridgeConstraints{}

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

	// from_asset can be empty (native token), so we initialize it if nil
	if res.fromAsset == nil {
		res.fromAsset = fixed("")
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
	// to_asset can be empty (native token), so we initialize it if nil
	if res.toAsset == nil {
		res.toAsset = fixed("")
	}
	if res.toAddress == nil {
		return res, fmt.Errorf("failed to find constraint: to_address")
	}

	return res, nil
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
			out.Resource = "solana.system.transfer"
			out.Target = &types.Target{
				TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
				Target: &types.Target_Address{
					Address: solana.SystemProgramID.String(),
				},
			}
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
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.TokenProgramID.String(),
			},
		}
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
		return []*types.Rule{out, {
			Resource: "solana.associated_token_account.create",
			Effect:   types.Effect_EFFECT_ALLOW,
			ParameterConstraints: []*types.ParameterConstraint{{
				ParameterName: "account_payer",
				Constraint:    c.fromAddress,
			}, {
				ParameterName: "account_associated_token_account",
				Constraint:    fixed(dst),
			}, {
				ParameterName: "account_owner",
				Constraint:    c.toAddress,
			}, {
				ParameterName: "account_mint",
				Constraint:    c.asset,
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
					Address: solana.SPLAssociatedTokenAccountProgramID.String(),
				},
			},
		}}, nil
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
	case lp:
		return m.handleSolanaLP(in, r)
	case lend:
		return m.handleSolanaLend(in, r)
	case perps:
		return m.handleSolanaPerps(in, r)
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
					ParameterName: "spender",
					Constraint:    router,
				},
				{
					ParameterName: "amount",
					Constraint:    c.fromAmount,
				},
			}
			// Bundle approve before swap so clients can sign 2 sequential messages:
			// approve tx, then swap tx (see iOS KeysignMessageFactory).
			rules = append([]*types.Rule{approve}, rules...)
		}

		return rules, nil
	case bridge:
		c, err := getBridgeConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse bridge constraints: %w", err)
		}

		chain, err := common.FromString(r.GetChainId())
		if err != nil {
			return nil, fmt.Errorf("invalid chainID: %w", err)
		}

		rules := make([]*types.Rule, 0)

		// Determine bridge provider based on route
		// LiFi is the default for cross-chain EVM bridges
		router := &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT,
			Value: &types.Constraint_MagicConstantValue{
				MagicConstantValue: types.MagicConstant_LIFI_ROUTER,
			},
		}

		// Create bridge rule for LiFi Diamond contract call
		// LiFi uses a generic call pattern through its Diamond proxy
		bridgeRule := proto.Clone(in).(*types.Rule)
		bridgeRule.Resource = fmt.Sprintf("%s.lifi_bridge.bridge", strings.ToLower(chain.String()))
		bridgeRule.Effect = types.Effect_EFFECT_ALLOW
		bridgeRule.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
			Target: &types.Target_MagicConstant{
				MagicConstant: types.MagicConstant_LIFI_ROUTER,
			},
		}
		bridgeRule.ParameterConstraints = []*types.ParameterConstraint{
			{
				ParameterName: "from_asset",
				Constraint:    c.fromAsset,
			},
			{
				ParameterName: "from_amount",
				Constraint:    c.fromAmount,
			},
			{
				ParameterName: "to_chain",
				Constraint:    c.toChain,
			},
			{
				ParameterName: "to_asset",
				Constraint:    c.toAsset,
			},
			{
				ParameterName: "to_address",
				Constraint:    c.toAddress,
			},
		}
		rules = append(rules, bridgeRule)

		// If bridging ERC20 token, add approval rule
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
	case lp:
		return m.handleEVMLP(in, r)
	case lend:
		return m.handleEVMLend(in, r)
	case perps:
		return m.handleEVMPerps(in, r)
	case bet:
		return m.handleEVMBet(in, r)
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
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse `send` constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "bitcoin.btc.transfer"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_UNSPECIFIED,
		}

		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "output_address_0",
			Constraint:    c.toAddress,
		}, {
			ParameterName: "output_value_0",
			Constraint:    c.amount,
		}, {
			ParameterName: "output_address_1",
			Constraint:    c.fromAddress,
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

		// Extract from_asset value, default to RUNE if empty (native token)
		fromAssetValue := c.fromAsset.GetFixedValue()
		if fromAssetValue == "" {
			fromAssetValue = "RUNE"
		}

		out.ParameterConstraints = []*types.ParameterConstraint{
			{
				ParameterName: "from_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAssetValue,
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

func (m *MetaRule) handleZcash(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse `send` constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "zcash.zec.transfer"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_UNSPECIFIED,
		}

		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "output_address_0",
			Constraint:    c.toAddress,
		}, {
			ParameterName: "output_value_0",
			Constraint:    c.amount,
		}, {
			ParameterName: "output_address_1",
			Constraint:    c.fromAddress,
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
		out.Resource = "zcash.zec.transfer"
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
			// Accept both full form and shortform: (ZEC\.ZEC|z)
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
		return nil, fmt.Errorf("unsupported protocol id for Zcash: %s", r.GetProtocolId())
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
		solana.SPLAssociatedTokenAccountProgramID,
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
					Address: solana.SPLAssociatedTokenAccountProgramID.String(),
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

func (m *MetaRule) handleTron(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse send constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "tron.trx.transfer"
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
		out.Resource = "tron.swap"
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
			assetPattern = fmt.Sprintf("(%s|%s)",
				regexp.QuoteMeta(thorAsset),
				regexp.QuoteMeta(shortCode))
		} else {
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
		return nil, fmt.Errorf("unsupported protocol id for Tron: %s", r.GetProtocolId())
	}
}

func (m *MetaRule) handleLitecoin(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse `send` constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "litecoin.ltc.transfer"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_UNSPECIFIED,
		}

		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "output_address_0",
			Constraint:    c.toAddress,
		}, {
			ParameterName: "output_value_0",
			Constraint:    c.amount,
		}, {
			ParameterName: "output_address_1",
			Constraint:    c.fromAddress,
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
		out.Resource = "litecoin.ltc.transfer"
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

		shortCode := thorchain.ShortCode(thorAsset)
		var assetPattern string
		if shortCode != "" {
			assetPattern = fmt.Sprintf("(%s|%s)",
				regexp.QuoteMeta(thorAsset),
				regexp.QuoteMeta(shortCode))
		} else {
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
			Constraint:    c.fromAddress,
		}, {
			ParameterName: "output_value_1",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "output_data_2",
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
		}}
		return []*types.Rule{out}, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id: %s", r.GetProtocolId())
	}
}

func (m *MetaRule) handleDogecoin(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse `send` constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "dogecoin.doge.transfer"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_UNSPECIFIED,
		}

		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "output_address_0",
			Constraint:    c.toAddress,
		}, {
			ParameterName: "output_value_0",
			Constraint:    c.amount,
		}, {
			ParameterName: "output_address_1",
			Constraint:    c.fromAddress,
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
		out.Resource = "dogecoin.doge.transfer"
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

		shortCode := thorchain.ShortCode(thorAsset)
		var assetPattern string
		if shortCode != "" {
			assetPattern = fmt.Sprintf("(%s|%s)",
				regexp.QuoteMeta(thorAsset),
				regexp.QuoteMeta(shortCode))
		} else {
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
			Constraint:    c.fromAddress,
		}, {
			ParameterName: "output_value_1",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "output_data_2",
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
		}}
		return []*types.Rule{out}, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id: %s", r.GetProtocolId())
	}
}

func (m *MetaRule) handleBitcoinCash(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse `send` constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "bitcoincash.bch.transfer"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_UNSPECIFIED,
		}

		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "output_address_0",
			Constraint:    c.toAddress,
		}, {
			ParameterName: "output_value_0",
			Constraint:    c.amount,
		}, {
			ParameterName: "output_address_1",
			Constraint:    c.fromAddress,
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
		out.Resource = "bitcoincash.bch.transfer"
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

		shortCode := thorchain.ShortCode(thorAsset)
		var assetPattern string
		if shortCode != "" {
			assetPattern = fmt.Sprintf("(%s|%s)",
				regexp.QuoteMeta(thorAsset),
				regexp.QuoteMeta(shortCode))
		} else {
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
			Constraint:    c.fromAddress,
		}, {
			ParameterName: "output_value_1",
			Constraint:    anyConstraint(),
		}, {
			ParameterName: "output_data_2",
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
		}}
		return []*types.Rule{out}, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id: %s", r.GetProtocolId())
	}
}

func (m *MetaRule) handleGaia(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse send constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "cosmos.atom.transfer"
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
		out.Resource = "cosmos.atom.transfer"
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
		return nil, fmt.Errorf("unsupported protocol id for Gaia: %s", r.GetProtocolId())
	}
}

func (m *MetaRule) handleMaya(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	switch metaProtocol(r.GetProtocolId()) {
	case send:
		c, err := getSendConstraints(in)
		if err != nil {
			return nil, fmt.Errorf("failed to parse send constraints: %w", err)
		}

		out := proto.Clone(in).(*types.Rule)
		out.Resource = "mayachain.cacao.transfer"
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
		out.Resource = "mayachain.mayachain_swap.swap"
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_UNSPECIFIED,
		}

		chainInt, err := common.FromString(c.toChain.GetFixedValue())
		if err != nil {
			return nil, fmt.Errorf("failed to parse chain id: %w", err)
		}

		// For Maya, we need to handle the asset differently
		// Maya uses its own asset notation similar to THORChain
		mayaAsset, err := mayachain.MakeAsset(chainInt, c.toAsset.GetFixedValue())
		if err != nil {
			return nil, fmt.Errorf("failed to make asset: %w", err)
		}

		shortCode := mayachain.ShortCode(mayaAsset)
		var assetPattern string
		if shortCode != "" {
			assetPattern = fmt.Sprintf("(%s|%s)",
				regexp.QuoteMeta(mayaAsset),
				regexp.QuoteMeta(shortCode))
		} else {
			assetPattern = regexp.QuoteMeta(mayaAsset)
		}

		// Extract from_asset value, default to CACAO if empty (native token)
		fromAssetValue := c.fromAsset.GetFixedValue()
		if fromAssetValue == "" {
			fromAssetValue = "CACAO"
		}

		out.ParameterConstraints = []*types.ParameterConstraint{
			{
				ParameterName: "from_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAssetValue,
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
		return nil, fmt.Errorf("unsupported protocol id for Maya: %s", r.GetProtocolId())
	}
}

// ============================================================================
// DeFi Meta-Protocol Handlers - EVM
// ============================================================================

func (m *MetaRule) handleEVMLP(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	c, err := getLPConstraints(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LP constraints: %w", err)
	}

	chain, err := common.FromString(r.GetChainId())
	if err != nil {
		return nil, fmt.Errorf("invalid chainID: %w", err)
	}

	protocol := c.protocol.GetFixedValue()
	action := c.action.GetFixedValue()

	// Get UniswapV3 Position Manager address for the chain
	positionManager, err := getUniswapV3PositionManager(chain)
	if err != nil {
		return nil, fmt.Errorf("failed to get position manager: %w", err)
	}

	rules := make([]*types.Rule, 0)

	switch protocol {
	case "uniswap_v3":
		switch action {
		case "add":
			// For ERC20 tokens, we must allow the exact approval required *before* mint.
			// This ensures clients can sign 2 sequential messages: approve tx, then mint tx.
			if c.token0.GetFixedValue() == "" {
				return nil, fmt.Errorf("token0 must be fixed for uniswap_v3 LP add")
			}
			if c.token1.GetFixedValue() == "" {
				return nil, fmt.Errorf("token1 must be fixed for uniswap_v3 LP add")
			}

			rules = append(rules,
				createERC20ApproveRule(chain, c.token0.GetFixedValue(), positionManager, c.amount0),
				createERC20ApproveRule(chain, c.token1.GetFixedValue(), positionManager, c.amount1),
			)

			// Create mint or increaseLiquidity rule
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.uniswapV3_nonfungible_position_manager.mint", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: positionManager},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "params.token0", Constraint: c.token0},
					{ParameterName: "params.token1", Constraint: c.token1},
					{ParameterName: "params.amount0Desired", Constraint: c.amount0},
					{ParameterName: "params.amount1Desired", Constraint: c.amount1},
					{ParameterName: "params.amount0Min", Constraint: c.minAmount0},
					{ParameterName: "params.amount1Min", Constraint: c.minAmount1},
					{ParameterName: "params.recipient", Constraint: c.recipient},
					{ParameterName: "params.fee", Constraint: anyConstraint()},
					{ParameterName: "params.tickLower", Constraint: anyConstraint()},
					{ParameterName: "params.tickUpper", Constraint: anyConstraint()},
					{ParameterName: "params.deadline", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, rule)

		case "remove":
			// decreaseLiquidity + collect
			decreaseRule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.uniswapV3_nonfungible_position_manager.decreaseLiquidity", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: positionManager},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "params.tokenId", Constraint: c.pool},
					{ParameterName: "params.liquidity", Constraint: anyConstraint()},
					{ParameterName: "params.amount0Min", Constraint: c.minAmount0},
					{ParameterName: "params.amount1Min", Constraint: c.minAmount1},
					{ParameterName: "params.deadline", Constraint: anyConstraint()},
				},
			}
			collectRule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.uniswapV3_nonfungible_position_manager.collect", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: positionManager},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "params.tokenId", Constraint: c.pool},
					{ParameterName: "params.recipient", Constraint: c.recipient},
					{ParameterName: "params.amount0Max", Constraint: anyConstraint()},
					{ParameterName: "params.amount1Max", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, decreaseRule, collectRule)

		case "collect_fees":
			// collect only
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.uniswapV3_nonfungible_position_manager.collect", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: positionManager},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "params.tokenId", Constraint: c.pool},
					{ParameterName: "params.recipient", Constraint: c.recipient},
					{ParameterName: "params.amount0Max", Constraint: c.amount0},
					{ParameterName: "params.amount1Max", Constraint: c.amount1},
				},
			}
			rules = append(rules, rule)

		default:
			return nil, fmt.Errorf("unsupported LP action: %s", action)
		}
	default:
		return nil, fmt.Errorf("unsupported LP protocol for EVM: %s", protocol)
	}

	return rules, nil
}

func (m *MetaRule) handleEVMLend(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	c, err := getLendConstraints(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse lend constraints: %w", err)
	}

	chain, err := common.FromString(r.GetChainId())
	if err != nil {
		return nil, fmt.Errorf("invalid chainID: %w", err)
	}

	protocol := c.protocol.GetFixedValue()
	action := c.action.GetFixedValue()

	rules := make([]*types.Rule, 0)

	switch protocol {
	case "aave":
		poolAddr, err := getAaveV3Pool(chain)
		if err != nil {
			return nil, fmt.Errorf("failed to get AAVE pool: %w", err)
		}

		switch action {
		case "supply":
			if c.asset.GetFixedValue() == "" {
				return nil, fmt.Errorf("asset must be fixed for aave supply (required for ERC20 approve)")
			}
			// Approve first so clients can sign approve tx then supply tx.
			rules = append(rules, createERC20ApproveRule(chain, c.asset.GetFixedValue(), poolAddr, c.amount))

			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.aaveV3_pool.supply", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: poolAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "asset", Constraint: c.asset},
					{ParameterName: "amount", Constraint: c.amount},
					{ParameterName: "onBehalfOf", Constraint: c.onBehalfOf},
					{ParameterName: "referralCode", Constraint: fixed("0")},
				},
			}
			rules = append(rules, rule)

		case "withdraw":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.aaveV3_pool.withdraw", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: poolAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "asset", Constraint: c.asset},
					{ParameterName: "amount", Constraint: c.amount},
					{ParameterName: "to", Constraint: c.onBehalfOf},
				},
			}
			rules = append(rules, rule)

		case "borrow":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.aaveV3_pool.borrow", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: poolAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "asset", Constraint: c.asset},
					{ParameterName: "amount", Constraint: c.amount},
					{ParameterName: "interestRateMode", Constraint: anyConstraint()},
					{ParameterName: "referralCode", Constraint: fixed("0")},
					{ParameterName: "onBehalfOf", Constraint: c.onBehalfOf},
				},
			}
			rules = append(rules, rule)

		case "repay":
			if c.asset.GetFixedValue() == "" {
				return nil, fmt.Errorf("asset must be fixed for aave repay (required for ERC20 approve)")
			}
			// Approve first so clients can sign approve tx then repay tx.
			rules = append(rules, createERC20ApproveRule(chain, c.asset.GetFixedValue(), poolAddr, c.amount))

			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.aaveV3_pool.repay", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: poolAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "asset", Constraint: c.asset},
					{ParameterName: "amount", Constraint: c.amount},
					{ParameterName: "interestRateMode", Constraint: anyConstraint()},
					{ParameterName: "onBehalfOf", Constraint: c.onBehalfOf},
				},
			}
			rules = append(rules, rule)

		default:
			return nil, fmt.Errorf("unsupported lend action: %s", action)
		}

	case "compound":
		cometAddr, err := getCompoundV3Comet(chain)
		if err != nil {
			return nil, fmt.Errorf("failed to get Compound comet: %w", err)
		}

		switch action {
		case "supply":
			if c.asset.GetFixedValue() == "" {
				return nil, fmt.Errorf("asset must be fixed for compound supply (required for ERC20 approve)")
			}
			// Approve first so clients can sign approve tx then supply tx.
			rules = append(rules, createERC20ApproveRule(chain, c.asset.GetFixedValue(), cometAddr, c.amount))

			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.compoundV3_comet.supply", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: cometAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "asset", Constraint: c.asset},
					{ParameterName: "amount", Constraint: c.amount},
				},
			}
			rules = append(rules, rule)

		case "withdraw":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.compoundV3_comet.withdraw", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: cometAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "asset", Constraint: c.asset},
					{ParameterName: "amount", Constraint: c.amount},
				},
			}
			rules = append(rules, rule)

		default:
			return nil, fmt.Errorf("unsupported lend action for Compound: %s", action)
		}

	default:
		return nil, fmt.Errorf("unsupported lend protocol for EVM: %s", protocol)
	}

	return rules, nil
}

func (m *MetaRule) handleEVMPerps(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	c, err := getPerpsConstraints(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse perps constraints: %w", err)
	}

	chain, err := common.FromString(r.GetChainId())
	if err != nil {
		return nil, fmt.Errorf("invalid chainID: %w", err)
	}

	protocol := c.protocol.GetFixedValue()
	action := c.action.GetFixedValue()

	rules := make([]*types.Rule, 0)

	switch protocol {
	case "gmx":
		routerAddr, err := getGMXV2ExchangeRouter(chain)
		if err != nil {
			return nil, fmt.Errorf("failed to get GMX router: %w", err)
		}

		// Determine isLong based on action
		isLong := action == "open_long"

		// createOrder for open/close
		switch action {
		case "open_long", "open_short":
			// Approve collateral first (exactly bounded by collateral_delta constraint).
			if c.collateralToken.GetFixedValue() != "" {
				rules = append(rules, createERC20ApproveRule(chain, c.collateralToken.GetFixedValue(), routerAddr, c.collateralDelta))
			}

			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.gmxV2_exchange_router.createOrder", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: routerAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "params.market", Constraint: c.market},
					{ParameterName: "params.initialCollateralToken", Constraint: c.collateralToken},
					{ParameterName: "params.sizeDeltaUsd", Constraint: c.sizeDelta},
					{ParameterName: "params.initialCollateralDeltaAmount", Constraint: c.collateralDelta},
					{ParameterName: "params.acceptablePrice", Constraint: c.acceptablePrice},
					{ParameterName: "params.executionFee", Constraint: c.executionFee},
					{ParameterName: "params.isLong", Constraint: fixed(fmt.Sprintf("%t", isLong))},
					{ParameterName: "params.orderType", Constraint: anyConstraint()},
					{ParameterName: "params.receiver", Constraint: anyConstraint()},
					{ParameterName: "params.callbackContract", Constraint: anyConstraint()},
					{ParameterName: "params.uiFeeReceiver", Constraint: anyConstraint()},
					{ParameterName: "params.swapPath", Constraint: anyConstraint()},
					{ParameterName: "params.triggerPrice", Constraint: anyConstraint()},
					{ParameterName: "params.callbackGasLimit", Constraint: anyConstraint()},
					{ParameterName: "params.minOutputAmount", Constraint: anyConstraint()},
					{ParameterName: "params.decreasePositionSwapType", Constraint: anyConstraint()},
					{ParameterName: "params.shouldUnwrapNativeToken", Constraint: anyConstraint()},
					{ParameterName: "params.referralCode", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, rule)

		case "close":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.gmxV2_exchange_router.createOrder", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: routerAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "params.market", Constraint: c.market},
					{ParameterName: "params.sizeDeltaUsd", Constraint: c.sizeDelta},
					{ParameterName: "params.acceptablePrice", Constraint: c.acceptablePrice},
					{ParameterName: "params.executionFee", Constraint: c.executionFee},
					{ParameterName: "params.orderType", Constraint: anyConstraint()},
					{ParameterName: "params.receiver", Constraint: anyConstraint()},
					{ParameterName: "params.callbackContract", Constraint: anyConstraint()},
					{ParameterName: "params.uiFeeReceiver", Constraint: anyConstraint()},
					{ParameterName: "params.initialCollateralToken", Constraint: anyConstraint()},
					{ParameterName: "params.swapPath", Constraint: anyConstraint()},
					{ParameterName: "params.initialCollateralDeltaAmount", Constraint: anyConstraint()},
					{ParameterName: "params.triggerPrice", Constraint: anyConstraint()},
					{ParameterName: "params.callbackGasLimit", Constraint: anyConstraint()},
					{ParameterName: "params.minOutputAmount", Constraint: anyConstraint()},
					{ParameterName: "params.decreasePositionSwapType", Constraint: anyConstraint()},
					{ParameterName: "params.isLong", Constraint: anyConstraint()},
					{ParameterName: "params.shouldUnwrapNativeToken", Constraint: anyConstraint()},
					{ParameterName: "params.referralCode", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, rule)

		default:
			return nil, fmt.Errorf("unsupported perps action: %s", action)
		}

	case "hyperliquid":
		// Hyperliquid is a custom L1 but has an Arbitrum bridge for deposits/withdrawals
		// For trading, it uses EIP-712 style signing on the HL L1
		bridgeAddr, err := getHyperliquidBridge(chain)
		if err != nil {
			return nil, fmt.Errorf("failed to get Hyperliquid bridge: %w", err)
		}

		switch action {
		case "open_long", "open_short":
			// isBid = true for long, false for short
			isBid := action == "open_long"
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.hyperliquid_bridge.order", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: bridgeAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "user", Constraint: anyConstraint()},
					{ParameterName: "isBid", Constraint: fixed(fmt.Sprintf("%t", isBid))},
					{ParameterName: "asset", Constraint: c.market},
					{ParameterName: "limitPx", Constraint: c.acceptablePrice},
					{ParameterName: "sz", Constraint: c.sizeDelta},
					{ParameterName: "reduceOnly", Constraint: fixed("false")},
					{ParameterName: "orderType", Constraint: anyConstraint()},
					{ParameterName: "cloid", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, rule)

		case "close":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.hyperliquid_bridge.order", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: bridgeAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "user", Constraint: anyConstraint()},
					{ParameterName: "isBid", Constraint: anyConstraint()},
					{ParameterName: "asset", Constraint: c.market},
					{ParameterName: "limitPx", Constraint: c.acceptablePrice},
					{ParameterName: "sz", Constraint: c.sizeDelta},
					{ParameterName: "reduceOnly", Constraint: fixed("true")},
					{ParameterName: "orderType", Constraint: anyConstraint()},
					{ParameterName: "cloid", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, rule)

		case "adjust_margin":
			// Hyperliquid margin adjustment via deposit/withdraw to bridge
			depositRule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.hyperliquid_bridge.batchedDepositWithPermit", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: bridgeAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "destination", Constraint: anyConstraint()},
					{ParameterName: "usd", Constraint: c.collateralDelta},
				},
			}
			withdrawRule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.hyperliquid_bridge.initiateWithdrawal", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: bridgeAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "usd", Constraint: c.collateralDelta},
				},
			}
			rules = append(rules, depositRule, withdrawRule)

		default:
			return nil, fmt.Errorf("unsupported perps action: %s", action)
		}

	default:
		return nil, fmt.Errorf("unsupported perps protocol for EVM: %s", protocol)
	}

	return rules, nil
}

// ============================================================================
// DeFi Meta-Protocol Handlers - Solana
// ============================================================================

func (m *MetaRule) handleSolanaLP(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	c, err := getLPConstraints(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LP constraints: %w", err)
	}

	protocol := c.protocol.GetFixedValue()
	action := c.action.GetFixedValue()

	rules := make([]*types.Rule, 0)

	switch protocol {
	case "raydium":
		programID := "CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK"

		switch action {
		case "add":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.raydium_clmm.openPosition",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_amount0Max", Constraint: c.amount0},
					{ParameterName: "arg_amount1Max", Constraint: c.amount1},
					{ParameterName: "arg_liquidity", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, rule)

		case "remove":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.raydium_clmm.decreaseLiquidity",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_amount0Min", Constraint: c.minAmount0},
					{ParameterName: "arg_amount1Min", Constraint: c.minAmount1},
					{ParameterName: "arg_liquidity", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, rule)

		case "collect_fees":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.raydium_clmm.collectFee",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_amount0Max", Constraint: c.amount0},
					{ParameterName: "arg_amount1Max", Constraint: c.amount1},
				},
			}
			rules = append(rules, rule)

		default:
			return nil, fmt.Errorf("unsupported LP action: %s", action)
		}

	case "orca":
		programID := "whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc"

		switch action {
		case "add":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.orca_whirlpool.increaseLiquidity",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_liquidityAmount", Constraint: anyConstraint()},
					{ParameterName: "arg_tokenMaxA", Constraint: c.amount0},
					{ParameterName: "arg_tokenMaxB", Constraint: c.amount1},
				},
			}
			rules = append(rules, rule)

		case "remove":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.orca_whirlpool.decreaseLiquidity",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_liquidityAmount", Constraint: anyConstraint()},
					{ParameterName: "arg_tokenMinA", Constraint: c.minAmount0},
					{ParameterName: "arg_tokenMinB", Constraint: c.minAmount1},
				},
			}
			rules = append(rules, rule)

		case "collect_fees":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.orca_whirlpool.collectFees",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{},
			}
			rules = append(rules, rule)

		default:
			return nil, fmt.Errorf("unsupported LP action: %s", action)
		}

	default:
		return nil, fmt.Errorf("unsupported LP protocol for Solana: %s", protocol)
	}

	return rules, nil
}

func (m *MetaRule) handleSolanaLend(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	c, err := getLendConstraints(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse lend constraints: %w", err)
	}

	protocol := c.protocol.GetFixedValue()
	action := c.action.GetFixedValue()

	rules := make([]*types.Rule, 0)

	switch protocol {
	case "kamino":
		programID := "KLend2g3cP87ber41GYr72yfE9j6eBJYwRqVNMi6mHL"

		switch action {
		case "supply":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.kamino_lending.deposit",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_liquidityAmount", Constraint: c.amount},
				},
			}
			rules = append(rules, rule)

		case "withdraw":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.kamino_lending.withdraw",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_collateralAmount", Constraint: c.amount},
				},
			}
			rules = append(rules, rule)

		case "borrow":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.kamino_lending.borrow",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_liquidityAmount", Constraint: c.amount},
				},
			}
			rules = append(rules, rule)

		case "repay":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.kamino_lending.repay",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_liquidityAmount", Constraint: c.amount},
				},
			}
			rules = append(rules, rule)

		default:
			return nil, fmt.Errorf("unsupported lend action: %s", action)
		}

	default:
		return nil, fmt.Errorf("unsupported lend protocol for Solana: %s", protocol)
	}

	return rules, nil
}

func (m *MetaRule) handleSolanaPerps(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	c, err := getPerpsConstraints(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse perps constraints: %w", err)
	}

	protocol := c.protocol.GetFixedValue()
	action := c.action.GetFixedValue()

	rules := make([]*types.Rule, 0)

	switch protocol {
	case "jupiter_perps":
		programID := "PERPHjGBqRHArX4DySjwM6UJHiR3sWAatqfdBS2qQJu"

		switch action {
		case "open_long", "open_short":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.jupiter_perpetuals.openPosition",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_size", Constraint: c.sizeDelta},
					{ParameterName: "arg_collateral", Constraint: c.collateralDelta},
					{ParameterName: "arg_price", Constraint: c.acceptablePrice},
					{ParameterName: "arg_side", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, rule)

		case "close":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.jupiter_perpetuals.closePosition",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_price", Constraint: c.acceptablePrice},
				},
			}
			rules = append(rules, rule)

		case "adjust_margin":
			// Add or remove collateral
			addRule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.jupiter_perpetuals.addCollateral",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_collateral", Constraint: c.collateralDelta},
				},
			}
			removeRule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.jupiter_perpetuals.removeCollateral",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_collateralUsd", Constraint: c.collateralDelta},
				},
			}
			rules = append(rules, addRule, removeRule)

		default:
			return nil, fmt.Errorf("unsupported perps action: %s", action)
		}

	case "drift":
		programID := "dRiftyHA39MWEi3m9aunc5MzRF1JYuBsbn6VPcn33UH"

		switch action {
		case "open_long", "open_short", "close":
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.drift_protocol.placePerpOrder",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: programID},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "arg_baseAssetAmount", Constraint: c.sizeDelta},
					{ParameterName: "arg_price", Constraint: c.acceptablePrice},
					{ParameterName: "arg_marketIndex", Constraint: anyConstraint()},
					{ParameterName: "arg_orderType", Constraint: anyConstraint()},
					{ParameterName: "arg_marketType", Constraint: anyConstraint()},
					{ParameterName: "arg_direction", Constraint: anyConstraint()},
					{ParameterName: "arg_userOrderId", Constraint: anyConstraint()},
					{ParameterName: "arg_reduceOnly", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, rule)

		default:
			return nil, fmt.Errorf("unsupported perps action: %s", action)
		}

	default:
		return nil, fmt.Errorf("unsupported perps protocol for Solana: %s", protocol)
	}

	return rules, nil
}

// ============================================================================
// Bet Meta-Protocol Handler - EVM
// ============================================================================

func (m *MetaRule) handleEVMBet(in *types.Rule, r *types.ResourcePath) ([]*types.Rule, error) {
	c, err := getBetConstraints(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse bet constraints: %w", err)
	}

	chain, err := common.FromString(r.GetChainId())
	if err != nil {
		return nil, fmt.Errorf("invalid chainID: %w", err)
	}

	protocol := c.protocol.GetFixedValue()
	action := c.action.GetFixedValue()

	rules := make([]*types.Rule, 0)

	switch protocol {
	case "polymarket":
		exchangeAddr, err := chainEVM.GetPolymarketCTFExchange(chain)
		if err != nil {
			return nil, fmt.Errorf("failed to get Polymarket exchange: %w", err)
		}

		usdcAddr, err := chainEVM.GetPolymarketUSDC(chain)
		if err != nil {
			return nil, fmt.Errorf("failed to get Polymarket USDC: %w", err)
		}

		ctfAddr, err := chainEVM.GetPolymarketCTF(chain)
		if err != nil {
			return nil, fmt.Errorf("failed to get Polymarket CTF: %w", err)
		}

		switch action {
		case "buy", "sell":
			// Verified PolygonScan ABI: `fillOrder((..., bytes signature) order, uint256 fillAmount)`
			// `order.side` is BUY/SELL (not YES/NO). YES/NO is encoded in `order.tokenId`.
			side := "0"
			if action == "sell" {
				side = "1"
			}

			// Bundle prerequisites first so clients can sign sequential txs:
			// 1) CTF setApprovalForAll, 2) ERC20 approve (buy only), 3) fillOrder/fillOrders.
			rules = append(rules, &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.polymarket_ctf.setApprovalForAll", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: ctfAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "operator", Constraint: fixed(exchangeAddr)},
					{ParameterName: "approved", Constraint: fixed("true")},
				},
			})

			if action == "buy" {
				rules = append(rules, createERC20ApproveRule(chain, usdcAddr, exchangeAddr, c.amount))
			}

			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.polymarket_ctf_exchange.fillOrder", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: exchangeAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "order.tokenId", Constraint: c.market},
					{ParameterName: "order.maker", Constraint: c.maker},
					{ParameterName: "order.side", Constraint: fixed(side)},
					{ParameterName: "fillAmount", Constraint: c.amount},

					{ParameterName: "order.salt", Constraint: anyConstraint()},
					{ParameterName: "order.signer", Constraint: anyConstraint()},
					{ParameterName: "order.taker", Constraint: anyConstraint()},
					{ParameterName: "order.makerAmount", Constraint: anyConstraint()},
					{ParameterName: "order.takerAmount", Constraint: anyConstraint()},
					{ParameterName: "order.expiration", Constraint: anyConstraint()},
					{ParameterName: "order.nonce", Constraint: anyConstraint()},
					{ParameterName: "order.feeRateBps", Constraint: anyConstraint()},
					{ParameterName: "order.signatureType", Constraint: anyConstraint()},
					{ParameterName: "order.signature", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, rule)

			// Also allow fillOrders for batch orders
			batchRule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.polymarket_ctf_exchange.fillOrders", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: exchangeAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "orders", Constraint: anyConstraint()},
					{ParameterName: "fillAmounts", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, batchRule)

		case "cancel":
			// cancelOrder(Order order) - order includes signature bytes in the ABI
			rule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.polymarket_ctf_exchange.cancelOrder", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: exchangeAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "order.tokenId", Constraint: c.market},
					{ParameterName: "order.maker", Constraint: c.maker},
					{ParameterName: "order.salt", Constraint: anyConstraint()},
					{ParameterName: "order.signer", Constraint: anyConstraint()},
					{ParameterName: "order.taker", Constraint: anyConstraint()},
					{ParameterName: "order.makerAmount", Constraint: anyConstraint()},
					{ParameterName: "order.takerAmount", Constraint: anyConstraint()},
					{ParameterName: "order.expiration", Constraint: anyConstraint()},
					{ParameterName: "order.nonce", Constraint: anyConstraint()},
					{ParameterName: "order.feeRateBps", Constraint: anyConstraint()},
					{ParameterName: "order.side", Constraint: anyConstraint()},
					{ParameterName: "order.signatureType", Constraint: anyConstraint()},
					{ParameterName: "order.signature", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, rule)

			// Also allow batch cancel
			batchCancelRule := &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("%s.polymarket_ctf_exchange.cancelOrders", strings.ToLower(chain.String())),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: exchangeAddr},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{ParameterName: "orders", Constraint: anyConstraint()},
				},
			}
			rules = append(rules, batchCancelRule)

		default:
			return nil, fmt.Errorf("unsupported bet action: %s", action)
		}

	default:
		return nil, fmt.Errorf("unsupported bet protocol for EVM: %s", protocol)
	}

	return rules, nil
}

// ============================================================================
// DeFi Constraint Parsers
// ============================================================================

type lpConstraints struct {
	action     *types.Constraint
	protocol   *types.Constraint
	pool       *types.Constraint
	token0     *types.Constraint
	token1     *types.Constraint
	amount0    *types.Constraint
	amount1    *types.Constraint
	minAmount0 *types.Constraint
	minAmount1 *types.Constraint
	recipient  *types.Constraint
}

func getLPConstraints(rule *types.Rule) (lpConstraints, error) {
	res := lpConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "action":
			res.action = c.GetConstraint()
		case "protocol":
			res.protocol = c.GetConstraint()
		case "pool":
			res.pool = c.GetConstraint()
		case "token0":
			res.token0 = c.GetConstraint()
		case "token1":
			res.token1 = c.GetConstraint()
		case "amount0":
			res.amount0 = c.GetConstraint()
		case "amount1":
			res.amount1 = c.GetConstraint()
		case "min_amount0":
			res.minAmount0 = c.GetConstraint()
		case "min_amount1":
			res.minAmount1 = c.GetConstraint()
		case "recipient":
			res.recipient = c.GetConstraint()
		}
	}

	if res.action == nil {
		return res, fmt.Errorf("missing required constraint: action")
	}
	if res.protocol == nil {
		return res, fmt.Errorf("missing required constraint: protocol")
	}

	// Set defaults for optional fields
	if res.minAmount0 == nil {
		res.minAmount0 = anyConstraint()
	}
	if res.minAmount1 == nil {
		res.minAmount1 = anyConstraint()
	}
	if res.amount0 == nil {
		res.amount0 = anyConstraint()
	}
	if res.amount1 == nil {
		res.amount1 = anyConstraint()
	}
	if res.recipient == nil {
		res.recipient = anyConstraint()
	}
	if res.pool == nil {
		res.pool = anyConstraint()
	}
	if res.token0 == nil {
		res.token0 = anyConstraint()
	}
	if res.token1 == nil {
		res.token1 = anyConstraint()
	}

	return res, nil
}

type lendConstraints struct {
	action     *types.Constraint
	protocol   *types.Constraint
	asset      *types.Constraint
	amount     *types.Constraint
	onBehalfOf *types.Constraint
}

func getLendConstraints(rule *types.Rule) (lendConstraints, error) {
	res := lendConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "action":
			res.action = c.GetConstraint()
		case "protocol":
			res.protocol = c.GetConstraint()
		case "asset":
			res.asset = c.GetConstraint()
		case "amount":
			res.amount = c.GetConstraint()
		case "on_behalf_of":
			res.onBehalfOf = c.GetConstraint()
		}
	}

	if res.action == nil {
		return res, fmt.Errorf("missing required constraint: action")
	}
	if res.protocol == nil {
		return res, fmt.Errorf("missing required constraint: protocol")
	}
	if res.asset == nil {
		return res, fmt.Errorf("missing required constraint: asset")
	}
	if res.amount == nil {
		return res, fmt.Errorf("missing required constraint: amount")
	}
	if res.onBehalfOf == nil {
		res.onBehalfOf = anyConstraint()
	}

	return res, nil
}

type perpsConstraints struct {
	action          *types.Constraint
	protocol        *types.Constraint
	market          *types.Constraint
	sizeDelta       *types.Constraint
	collateralDelta *types.Constraint
	collateralToken *types.Constraint
	acceptablePrice *types.Constraint
	executionFee    *types.Constraint
}

func getPerpsConstraints(rule *types.Rule) (perpsConstraints, error) {
	res := perpsConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "action":
			res.action = c.GetConstraint()
		case "protocol":
			res.protocol = c.GetConstraint()
		case "market":
			res.market = c.GetConstraint()
		case "size_delta":
			res.sizeDelta = c.GetConstraint()
		case "collateral_delta":
			res.collateralDelta = c.GetConstraint()
		case "collateral_token":
			res.collateralToken = c.GetConstraint()
		case "acceptable_price":
			res.acceptablePrice = c.GetConstraint()
		case "execution_fee":
			res.executionFee = c.GetConstraint()
		}
	}

	if res.action == nil {
		return res, fmt.Errorf("missing required constraint: action")
	}
	if res.protocol == nil {
		return res, fmt.Errorf("missing required constraint: protocol")
	}

	// Set defaults for optional fields
	if res.market == nil {
		res.market = anyConstraint()
	}
	if res.sizeDelta == nil {
		res.sizeDelta = anyConstraint()
	}
	if res.collateralDelta == nil {
		res.collateralDelta = anyConstraint()
	}
	if res.collateralToken == nil {
		res.collateralToken = anyConstraint()
	}
	if res.acceptablePrice == nil {
		res.acceptablePrice = anyConstraint()
	}
	if res.executionFee == nil {
		res.executionFee = anyConstraint()
	}

	return res, nil
}

type betConstraints struct {
	action   *types.Constraint
	protocol *types.Constraint
	market   *types.Constraint
	amount   *types.Constraint
	price    *types.Constraint
	maker    *types.Constraint
}

func getBetConstraints(rule *types.Rule) (betConstraints, error) {
	res := betConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "action":
			res.action = c.GetConstraint()
		case "protocol":
			res.protocol = c.GetConstraint()
		case "market":
			res.market = c.GetConstraint()
		case "amount":
			res.amount = c.GetConstraint()
		case "price":
			res.price = c.GetConstraint()
		case "maker":
			res.maker = c.GetConstraint()
		}
	}

	if res.action == nil {
		return res, fmt.Errorf("missing required constraint: action")
	}
	if res.protocol == nil {
		return res, fmt.Errorf("missing required constraint: protocol")
	}

	// Set defaults for optional fields
	if res.market == nil {
		res.market = anyConstraint()
	}
	if res.amount == nil {
		res.amount = anyConstraint()
	}
	if res.price == nil {
		res.price = anyConstraint()
	}
	if res.maker == nil {
		res.maker = anyConstraint()
	}

	return res, nil
}

// ============================================================================
// DeFi Helper Functions
// ============================================================================

func getUniswapV3PositionManager(chain common.Chain) (string, error) {
	addrs := map[common.Chain]string{
		common.Ethereum: "0xC36442b4a4522E871399CD717aBDD847Ab11FE88",
		common.Arbitrum: "0xC36442b4a4522E871399CD717aBDD847Ab11FE88",
		common.Optimism: "0xC36442b4a4522E871399CD717aBDD847Ab11FE88",
		common.Base:     "0x03a520b32C04BF3bEEf7BEb72E919cf822Ed34f1",
		common.Polygon:  "0xC36442b4a4522E871399CD717aBDD847Ab11FE88",
	}
	addr, ok := addrs[chain]
	if !ok {
		return "", fmt.Errorf("UniswapV3 not available on chain: %s", chain.String())
	}
	return addr, nil
}

func getAaveV3Pool(chain common.Chain) (string, error) {
	addrs := map[common.Chain]string{
		common.Ethereum:  "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2",
		common.Arbitrum:  "0x794a61358D6845594F94dc1DB02A252b5b4814aD",
		common.Optimism:  "0x794a61358D6845594F94dc1DB02A252b5b4814aD",
		common.Base:      "0xA238Dd80C259a72e81d7e4664a9801593F98d1c5",
		common.Polygon:   "0x794a61358D6845594F94dc1DB02A252b5b4814aD",
		common.Avalanche: "0x794a61358D6845594F94dc1DB02A252b5b4814aD",
	}
	addr, ok := addrs[chain]
	if !ok {
		return "", fmt.Errorf("AAVE V3 not available on chain: %s", chain.String())
	}
	return addr, nil
}

func getCompoundV3Comet(chain common.Chain) (string, error) {
	// Default to USDC market
	addrs := map[common.Chain]string{
		common.Ethereum: "0xc3d688B66703497DAA19211EEdff47f25384cdc3",
		common.Arbitrum: "0xA5EDBDD9646f8dFF606d7448e414884C7d905dCA",
		common.Base:     "0xb125E6687d4313864e53df431d5425969c15Eb2F",
		common.Polygon:  "0xF25212E676D1F7F89Cd72fFEe66158f541246445",
	}
	addr, ok := addrs[chain]
	if !ok {
		return "", fmt.Errorf("compound V3 not available on chain: %s", chain.String())
	}
	return addr, nil
}

func getGMXV2ExchangeRouter(chain common.Chain) (string, error) {
	addrs := map[common.Chain]string{
		common.Arbitrum:  "0x7C68C7866A64FA2160F78EEaE12217FFbf871fa8",
		common.Avalanche: "0x79be2F4eC8A4143BaF963206cF133f3710856D0a",
	}
	addr, ok := addrs[chain]
	if !ok {
		return "", fmt.Errorf("GMX V2 not available on chain: %s", chain.String())
	}
	return addr, nil
}

func getHyperliquidBridge(chain common.Chain) (string, error) {
	// Hyperliquid is a custom L1, but has a bridge on Arbitrum for deposits/withdrawals
	addrs := map[common.Chain]string{
		common.Arbitrum: "0x2Df1c51E09aECF9cacB7bc98cB1742757f163dF7",
	}
	addr, ok := addrs[chain]
	if !ok {
		return "", fmt.Errorf("hyperliquid bridge not available on chain: %s", chain.String())
	}
	return addr, nil
}

func createERC20ApproveRule(chain common.Chain, tokenAddr, spender string, amount *types.Constraint) *types.Rule {
	return &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: fmt.Sprintf("%s.erc20.approve", strings.ToLower(chain.String())),
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target:     &types.Target_Address{Address: tokenAddr},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{ParameterName: "spender", Constraint: fixed(spender)},
			{ParameterName: "amount", Constraint: amount},
		},
	}
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
