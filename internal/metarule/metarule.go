package metarule

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/vultisig/recipes/internal/metarule/thorchain"
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
	recipient *types.Constraint
	amount    *types.Constraint
}

func getSendConstraints(rule *types.Rule) (sendConstraints, error) {
	res := sendConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "recipient":
			res.recipient = c.GetConstraint()
		case "amount":
			res.amount = c.GetConstraint()
		}
	}

	if res.recipient == nil {
		return res, fmt.Errorf("failed to find constraint: recipient")
	}
	if res.amount == nil {
		return res, fmt.Errorf("failed to find constraint: amount")
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

		if in.GetTarget().GetAddress() == solana.SystemProgramID.String() {
			// native transfer
			var outTarget *types.Target
			switch c.recipient.GetType() {
			case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: c.recipient.GetFixedValue(),
					},
				}
			case types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
					Target: &types.Target_MagicConstant{
						MagicConstant: c.recipient.GetMagicConstantValue(),
					},
				}
			default:
				return nil, fmt.Errorf(
					"invalid constraint type for `recipient`: %s",
					c.recipient.GetType().String(),
				)
			}

			out.Resource = "solana.system.transfer"
			out.Target = outTarget
			out.ParameterConstraints = []*types.ParameterConstraint{{
				ParameterName: "account_from",
				Constraint:    anyConstraint(),
			}, {
				ParameterName: "account_to",
				Constraint:    c.recipient,
			}, {
				ParameterName: "arg_lamports",
				Constraint:    c.amount,
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
			Constraint:    c.recipient,
		}, {
			ParameterName: "account_authority",
			Constraint:    anyConstraint(),
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

		if in.GetTarget().GetAddress() == evm.ZeroAddress.String() {
			var outTarget *types.Target
			switch c.recipient.GetType() {
			case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: c.recipient.GetFixedValue(),
					},
				}
			case types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT:
				outTarget = &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
					Target: &types.Target_MagicConstant{
						MagicConstant: c.recipient.GetMagicConstantValue(),
					},
				}
			default:
				return nil, fmt.Errorf(
					"invalid constraint type for `recipient`: %s",
					c.recipient.GetType().String(),
				)
			}

			out.Resource = fmt.Sprintf("%s.%s.transfer", strings.ToLower(chain.String()), nativeSymbol)
			out.Target = outTarget
			out.ParameterConstraints = []*types.ParameterConstraint{{
				ParameterName: "amount",
				Constraint:    c.amount,
			}}

			return []*types.Rule{out}, nil
		}

		// erc20 token transfer
		out.Resource = fmt.Sprintf("%s.erc20.transfer", strings.ToLower(chain.String()))
		out.Target = in.GetTarget()
		out.ParameterConstraints = []*types.ParameterConstraint{{
			ParameterName: "recipient",
			Constraint:    c.recipient,
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
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
			{
				ParameterName: "spender",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: in.GetTarget().GetAddress(),
					},
					Required: true,
				},
			},
		}

		rules = append(rules, approve)

		out := proto.Clone(in).(*types.Rule)
		out.Resource = fmt.Sprintf("%s.routerV6_1inch.swap", strings.ToLower(chain.String()))
		out.Target = in.GetTarget()
		out.ParameterConstraints = []*types.ParameterConstraint{
			{
				ParameterName: "executor",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
			{
				ParameterName: "desc.srcToken",
				Constraint:    c.fromAsset,
			},
			{
				ParameterName: "desc.dstToken",
				Constraint:    c.toAsset,
			},
			{
				ParameterName: "desc.srcReceiver",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
			{
				ParameterName: "desc.dstReceiver",
				Constraint:    c.fromAddress,
			},
			{
				ParameterName: "desc.amount",
				Constraint:    c.fromAmount,
			},
			{
				ParameterName: "desc.minReturnAmount",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
			{
				ParameterName: "desc.flags",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			},
			{
				ParameterName: "data",
				Constraint: &types.Constraint{
					Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: true,
				},
			}}
		rules = append(rules, out)
		return rules, nil
	default:
		return nil, fmt.Errorf("unsupported protocol id: %s", r.GetProtocolId())
	}
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
				Constraint:    c.recipient,
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

func (m *MetaRule) createJupiterRule(in *types.Rule, c swapConstraints) ([]*types.Rule, error) {
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

	var rules []*types.Rule

	fromAssetStr, err := getFixed(c.fromAsset)
	if err != nil {
		return nil, fmt.Errorf("failed to get fixed value: %w", err)
	}

	toAssetStr, err := getFixed(c.toAsset)
	if err != nil {
		return nil, fmt.Errorf("failed to get fixed value for toAsset: %w", err)
	}

	sourceMintConstraint := c.fromAsset
	if fromAssetStr == "" {
		sourceMintConstraint = fixed(solana.SolMint.String())
	}

	destinationMintConstraint := c.toAsset
	if toAssetStr == "" {
		destinationMintConstraint = fixed(solana.SolMint.String())
	}

	baseConstraints := []*types.ParameterConstraint{{
		ParameterName: "arg_routePlan",
		Constraint:    anyConstraint(),
	}, {
		ParameterName: "arg_slippageBps",
		Constraint:    anyConstraint(),
	}, {
		ParameterName: "arg_platformFeeBps",
		Constraint:    anyConstraint(),
	}}

	jupiterInstructions := []string{
		"route",
		"routeWithTokenLedger",
		"sharedAccountsRoute",
		"sharedAccountsRouteWithTokenLedger",
		"exactOutRoute",
	}

	fromAddressStr, err := getFixed(c.fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get fixed value for fromAddress: %w", err)
	}

	toAddressStr, err := getFixed(c.toAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get fixed value for toAddress: %w", err)
	}

	var sourceTokenAccountConstraint *types.Constraint
	sourceMintForATA := fromAssetStr
	if sourceMintForATA == "" {
		sourceMintForATA = solana.SolMint.String()
	}
	sourceATA, err := DeriveATA(fromAddressStr, sourceMintForATA)
	if err != nil {
		return nil, fmt.Errorf("failed to derive source ATA for owner %s and mint %s: %w", fromAddressStr, sourceMintForATA, err)
	}
	sourceTokenAccountConstraint = fixed(sourceATA)

	var destinationTokenAccountConstraint *types.Constraint
	destMintForATA := toAssetStr
	if destMintForATA == "" {
		destMintForATA = solana.SolMint.String()
	}
	destATA, err := DeriveATA(toAddressStr, destMintForATA)
	if err != nil {
		return nil, fmt.Errorf("failed to derive destination ATA for owner %s and mint %s: %w", toAddressStr, destMintForATA, err)
	}
	destinationTokenAccountConstraint = fixed(destATA)

	// Add createIdempotent for source ATA if source asset is not native SOL
	if fromAssetStr != "" {
		rules = append(rules, &types.Rule{
			Resource: "solana.associated_token_account.createIdempotent",
			Effect:   types.Effect_EFFECT_ALLOW,
			ParameterConstraints: []*types.ParameterConstraint{{
				ParameterName: "account_payer",
				Constraint:    c.fromAddress,
			}, {
				ParameterName: "account_associatedTokenAccount",
				Constraint:    sourceTokenAccountConstraint,
			}, {
				ParameterName: "account_owner",
				Constraint:    c.fromAddress,
			}, {
				ParameterName: "account_mint",
				Constraint:    c.fromAsset,
			}, {
				ParameterName: "account_systemProgram",
				Constraint:    fixed(solana.SystemProgramID.String()),
			}, {
				ParameterName: "account_tokenProgram",
				Constraint:    fixed(solana.TokenProgramID.String()),
			}},
			Target: &types.Target{
				TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
				Target: &types.Target_Address{
					Address: "ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL",
				},
			},
		})
	}

	// Add createIdempotent for destination ATA if destination asset is not native SOL
	if toAssetStr != "" {
		rules = append(rules, &types.Rule{
			Resource: "solana.associated_token_account.createIdempotent",
			Effect:   types.Effect_EFFECT_ALLOW,
			ParameterConstraints: []*types.ParameterConstraint{{
				ParameterName: "account_payer",
				Constraint:    c.fromAddress,
			}, {
				ParameterName: "account_associatedTokenAccount",
				Constraint:    destinationTokenAccountConstraint,
			}, {
				ParameterName: "account_owner",
				Constraint:    c.toAddress,
			}, {
				ParameterName: "account_mint",
				Constraint:    c.toAsset,
			}, {
				ParameterName: "account_systemProgram",
				Constraint:    fixed(solana.SystemProgramID.String()),
			}, {
				ParameterName: "account_tokenProgram",
				Constraint:    fixed(solana.TokenProgramID.String()),
			}},
			Target: &types.Target{
				TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
				Target: &types.Target_Address{
					Address: "ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL",
				},
			},
		})
	}

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

	for _, instruction := range jupiterInstructions {
		out := proto.Clone(in).(*types.Rule)
		out.Resource = "solana.jupiter_aggregatorv6." + instruction
		out.Target = &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: jupAddr,
			},
		}

		constraints := []*types.ParameterConstraint{{
			ParameterName: "account_tokenProgram",
			Constraint:    fixed(solana.TokenProgramID.String()),
		}, {
			ParameterName: "account_userTransferAuthority",
			Constraint:    anyConstraint(),
		}}

		if instruction == "sharedAccountsRoute" || instruction == "sharedAccountsRouteWithTokenLedger" {
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_programAuthority",
				Constraint:    anyConstraint(),
			})
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_sourceTokenAccount",
				Constraint:    sourceTokenAccountConstraint,
			})
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_programSourceTokenAccount",
				Constraint:    anyConstraint(),
			})
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_programDestinationTokenAccount",
				Constraint:    anyConstraint(),
			})
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_destinationTokenAccount",
				Constraint:    destinationTokenAccountConstraint,
			})
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_sourceMint",
				Constraint:    sourceMintConstraint,
			})
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_destinationMint",
				Constraint:    destinationMintConstraint,
			})
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_platformFeeAccount",
				Constraint:    anyConstraint(),
			})
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_token2022Program",
				Constraint:    anyConstraint(),
			})
		} else {
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_userSourceTokenAccount",
				Constraint:    sourceTokenAccountConstraint,
			})
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_userDestinationTokenAccount",
				Constraint:    destinationTokenAccountConstraint,
			})
			if instruction == "exactOutRoute" {
				constraints = append(constraints, &types.ParameterConstraint{
					ParameterName: "account_sourceMint",
					Constraint:    sourceMintConstraint,
				})
			}
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_destinationMint",
				Constraint:    destinationMintConstraint,
			})
		}

		if instruction == "routeWithTokenLedger" || instruction == "sharedAccountsRouteWithTokenLedger" {
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "account_tokenLedger",
				Constraint:    anyConstraint(),
			})
		}

		constraints = append(constraints, &types.ParameterConstraint{
			ParameterName: "account_eventAuthority",
			Constraint:    fixed(jupEvent),
		})
		constraints = append(constraints, &types.ParameterConstraint{
			ParameterName: "account_program",
			Constraint:    fixed(jupAddr),
		})

		if instruction == "sharedAccountsRoute" || instruction == "sharedAccountsRouteWithTokenLedger" {
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "arg_id",
				Constraint:    anyConstraint(),
			})
		}

		constraints = append(constraints, baseConstraints...)

		if instruction == "route" || instruction == "sharedAccountsRoute" {
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "arg_inAmount",
				Constraint:    c.fromAmount,
			})
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "arg_quotedOutAmount",
				Constraint:    anyConstraint(),
			})
		} else if instruction == "exactOutRoute" {
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "arg_outAmount",
				Constraint:    anyConstraint(),
			})
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "arg_quotedInAmount",
				Constraint:    anyConstraint(),
			})
		} else {
			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: "arg_quotedOutAmount",
				Constraint:    anyConstraint(),
			})
		}

		out.ParameterConstraints = constraints
		rules = append(rules, out)
	}

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
