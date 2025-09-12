package solana

import (
	"encoding/json"
	"fmt"
	"math/big"
	"path"
	"regexp"
	"strings"

	stdcompare "github.com/vultisig/recipes/engine/compare"
	"github.com/vultisig/recipes/engine/solana/compare"
	idl_embed "github.com/vultisig/recipes/idl"
	"github.com/vultisig/recipes/resolver"
	solanautil "github.com/vultisig/recipes/solana"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
)

type Solana struct {
	nativeSymbol string
	idl          map[protocolID]idl
}

type idl struct {
	Instructions []idlInstruction `json:"instructions"`
	Accounts     []idlAccount     `json:"accounts,omitempty"`
	Types        []idlType        `json:"types,omitempty"`
	Name         string           `json:"name"`
	Version      string           `json:"version"`
}

type idlInstruction struct {
	Name     string        `json:"name"`
	Accounts []idlAccount  `json:"accounts"`
	Args     []idlArgument `json:"args"`
}

type idlAccount struct {
	Name     string `json:"name"`
	IsMut    bool   `json:"isMut"`
	IsSigner bool   `json:"isSigner"`
}

type idlArgument struct {
	Name string      `json:"name"`
	Type interface{} `json:"type"`
}

type idlType struct {
	Name string                 `json:"name"`
	Type map[string]interface{} `json:"type"`
}

func NewSolana(nativeSymbol string) (*Solana, error) {
	idls, err := loadIDLDir()
	if err != nil {
		return nil, fmt.Errorf("failed to load idl dir: %w", err)
	}

	return &Solana{
		nativeSymbol: strings.ToLower(nativeSymbol),
		idl:          idls,
	}, nil
}

func (s *Solana) Evaluate(rule *types.Rule, txBytes []byte) error {
	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
	}

	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}

	txData, err := solanautil.DecodeTransaction(txBytes)
	if err != nil {
		return fmt.Errorf("failed to decode tx payload: %w", err)
	}

	// Only allow single instruction transactions
	if len(txData.Message.Instructions) != 1 {
		return fmt.Errorf(
			"only single instruction transactions are allowed, got %d instructions",
			len(txData.Message.Instructions),
		)
	}

	err = s.assertTarget(r, rule.GetTarget(), txData)
	if err != nil {
		return fmt.Errorf("failed to assert target: %w", err)
	}

	if r.ProtocolId == s.nativeSymbol {
		er := s.assertArgsNative(r, rule, txData)
		if er != nil {
			return fmt.Errorf("failed to evaluate native: symbol=%s, error=%w", s.nativeSymbol, er)
		}
		return nil
	}

	er := s.assertArgsIDL(r, rule, txData)
	if er != nil {
		return fmt.Errorf("failed to evaluate IDL: %w", er)
	}
	return nil
}

func (s *Solana) assertArgsNative(
	resource *types.ResourcePath,
	rule *types.Rule,
	txData *solanautil.TransactionData,
) error {
	if resource.FunctionId != "transfer" {
		return fmt.Errorf(
			"only 'transfer' function supported for native: symbol=%s, function_id=%s",
			resource.ProtocolId,
			resource.FunctionId,
		)
	}

	if len(rule.GetParameterConstraints()) != 1 {
		return fmt.Errorf(
			"expected 1 parameter constraint for native transfer, got: %d",
			len(rule.GetParameterConstraints()),
		)
	}

	// Find the System Program transfer instruction
	transferInstruction, err := solanautil.FindTransferInstruction(txData)
	if err != nil {
		return fmt.Errorf("no system program transfer instruction found: %w", err)
	}

	// Parse the lamports amount from instruction data
	amount, err := solanautil.ParseTransferAmount(*transferInstruction, txData.Message.AccountKeys)
	if err != nil {
		return fmt.Errorf("failed to parse transfer amount: %w", err)
	}

	err = assertArg(
		resource.ChainId,
		rule.GetParameterConstraints(),
		"amount",
		amount,
		compare.NewU64,
	)
	if err != nil {
		return fmt.Errorf("failed to assert amount arg: %w", err)
	}
	return nil
}

func (s *Solana) assertTarget(
	resource *types.ResourcePath,
	target *types.Target,
	txData *solanautil.TransactionData,
) error {
	targetKind := target.GetTargetType()
	switch targetKind {
	case types.TargetType_TARGET_TYPE_ADDRESS:
		// For Solana, we need to check the appropriate account based on the instruction type
		var targetAddr string
		if resource.ProtocolId == s.nativeSymbol {
			// For native transfers, check the destination account
			targetAddr = s.getTransferDestination(txData)
		} else {
			// For program instructions, check based on the specific program logic
			targetAddr = s.getProgramDestination(resource, txData)
		}

		if targetAddr == "" || targetAddr != target.GetAddress() {
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_target_address=%s",
				targetAddr,
				target.GetAddress(),
			)
		}
		return nil

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

		var targetAddr string
		if resource.ProtocolId == s.nativeSymbol {
			targetAddr = s.getTransferDestination(txData)
		} else {
			targetAddr = s.getProgramDestination(resource, txData)
		}

		if targetAddr == "" || targetAddr != resolvedAddr {
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_magic_const_resolved=%s",
				targetAddr,
				resolvedAddr,
			)
		}
		return nil

	default:
		return fmt.Errorf("unknown target type: %s", targetKind.String())
	}
}

func (s *Solana) getTransferDestination(txData *solanautil.TransactionData) string {
	// Find the transfer instruction and get its destination
	transferInst, err := solanautil.FindTransferInstruction(txData)
	if err != nil {
		return ""
	}

	destination, err := solanautil.GetTransferDestination(*transferInst, txData.Message.AccountKeys)
	if err != nil {
		return ""
	}

	return destination.String()
}

func (s *Solana) getProgramDestination(resource *types.ResourcePath, txData *solanautil.TransactionData) string {
	// This would need to be implemented based on specific program logic
	// For now, return empty string as placeholder
	return ""
}

type protocolID = string

func loadIDLDir() (map[protocolID]idl, error) {
	base := "."

	entries, err := idl_embed.Dir.ReadDir(base)
	if err != nil {
		return nil, fmt.Errorf("failed to read idl dir: err=%w", err)
	}

	idls := make(map[string]idl)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		filepath := path.Join(base, entry.Name())
		file, er := idl_embed.Dir.Open(filepath)
		if er != nil {
			return nil, fmt.Errorf("failed to open idl json: path=%s, err=%w", filepath, er)
		}

		var idl idl
		er = json.NewDecoder(file).Decode(&idl)
		_ = file.Close()
		if er != nil {
			return nil, fmt.Errorf("failed to parse idl json: %w", er)
		}

		idls[strings.TrimSuffix(entry.Name(), ".json")] = idl
	}
	return idls, nil
}

func (s *Solana) assertArgsIDL(
	resource *types.ResourcePath,
	rule *types.Rule,
	txData *solanautil.TransactionData,
) error {
	idlItem, ok := s.idl[resource.ProtocolId]
	if !ok {
		return fmt.Errorf("failed to get idl: protocolId=%s", resource.ProtocolId)
	}

	// Find the instruction in the IDL
	var instruction *idlInstruction
	for _, inst := range idlItem.Instructions {
		if inst.Name == resource.FunctionId {
			instruction = &inst
			break
		}
	}

	if instruction == nil {
		return fmt.Errorf("failed to find idl instruction: %s", resource.FunctionId)
	}

	// Find the matching instruction in the transaction
	// For now, we'll use the first instruction that matches the expected program
	if len(txData.Message.Instructions) == 0 {
		return fmt.Errorf("no instructions found in transaction")
	}

	// Use the first instruction
	txInstruction := &txData.Message.Instructions[0]

	// Convert IDL instruction to the format expected by the parser
	solanaInstruction := &solanautil.IDLInstruction{
		Name:     instruction.Name,
		Accounts: make([]solanautil.IDLAccount, len(instruction.Accounts)),
		Args:     make([]solanautil.IDLArgument, len(instruction.Args)),
	}

	for i, acc := range instruction.Accounts {
		solanaInstruction.Accounts[i] = solanautil.IDLAccount{
			Name:     acc.Name,
			IsMut:    acc.IsMut,
			IsSigner: acc.IsSigner,
		}
	}

	for i, arg := range instruction.Args {
		solanaInstruction.Args[i] = solanautil.IDLArgument{
			Name: arg.Name,
			Type: arg.Type,
		}
	}

	// Parse instruction arguments - for proper implementation we'd use the compiled instruction data
	args, err := solanautil.ParseInstructionArgs(txInstruction.Data, solanaInstruction)
	if err != nil {
		return fmt.Errorf("failed to parse instruction args: %w", err)
	}

	if len(rule.GetParameterConstraints()) != len(args) {
		return fmt.Errorf(
			"constraints must be same list as IDL args: constraints_len=%d, idl_args_len=%d",
			len(rule.GetParameterConstraints()),
			len(args),
		)
	}

	for i, arg := range args {
		argDef := instruction.Args[i]
		switch actual := arg.(type) {
		case string:
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				argDef.Name,
				actual,
				stdcompare.NewString,
			)
			if er != nil {
				return fmt.Errorf("failed to assert string arg: %w", er)
			}

		case []byte:
			// Convert to string for comparison (Solana public keys are often represented as base58 strings)
			pubkeyStr := string(actual)
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				argDef.Name,
				pubkeyStr,
				compare.NewPubkey,
			)
			if er != nil {
				return fmt.Errorf("failed to assert pubkey arg: %w", er)
			}

		case *big.Int:
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				argDef.Name,
				actual,
				stdcompare.NewBigInt,
			)
			if er != nil {
				return fmt.Errorf("failed to assert bigint arg: %w", er)
			}

		case uint64:
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				argDef.Name,
				actual,
				compare.NewU64,
			)
			if er != nil {
				return fmt.Errorf("failed to assert u64 arg: %w", er)
			}

		case bool:
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				argDef.Name,
				actual,
				stdcompare.NewBool,
			)
			if er != nil {
				return fmt.Errorf("failed to assert bool arg: %w", er)
			}

		default:
			return fmt.Errorf("unsupported arg type: %T", actual)
		}
	}
	return nil
}

func assertArg[T any](
	chain string,
	expectedList []*types.ParameterConstraint,
	expectedName string,
	actual T,
	makeComparer stdcompare.Constructor[T],
) error {
	const magicAssetIdDefault = "default"

	for _, constraint := range expectedList {
		if constraint.GetParameterName() == expectedName {
			kind := constraint.GetConstraint().GetType()

			switch kind {
			case types.ConstraintType_CONSTRAINT_TYPE_ANY:
				return nil

			case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
				comparer, err := makeComparer(constraint.GetConstraint().GetFixedValue())
				if err != nil {
					return fmt.Errorf(
						"failed to build exact fixed type from constraint: %s",
						constraint.GetConstraint().GetFixedValue(),
					)
				}
				if comparer.Fixed(actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare fixed values: expected=%v, actual=%v",
					constraint.GetConstraint().GetFixedValue(),
					actual,
				)

			case types.ConstraintType_CONSTRAINT_TYPE_MIN:
				comparer, err := makeComparer(constraint.GetConstraint().GetMinValue())
				if err != nil {
					return fmt.Errorf(
						"failed to build exact min type from constraint: %s",
						constraint.GetConstraint().GetMinValue(),
					)
				}
				if comparer.Min(actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare min values: expected=%v, actual=%v",
					constraint.GetConstraint().GetMinValue(),
					actual,
				)

			case types.ConstraintType_CONSTRAINT_TYPE_MAX:
				comparer, err := makeComparer(constraint.GetConstraint().GetMaxValue())
				if err != nil {
					return fmt.Errorf(
						"failed to build exact max type from constraint: %s",
						constraint.GetConstraint().GetMaxValue(),
					)
				}
				if comparer.Max(actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare max values: expected=%v, actual=%v",
					constraint.GetConstraint().GetMaxValue(),
					actual,
				)

			case types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT:
				resolve, err := resolver.NewMagicConstantRegistry().GetResolver(
					constraint.GetConstraint().GetMagicConstantValue(),
				)
				if err != nil {
					return fmt.Errorf(
						"failed to get magic const resolver: magic_const=%s",
						constraint.GetConstraint().GetMagicConstantValue().String(),
					)
				}

				resolvedAddr, _, err := resolve.Resolve(
					constraint.GetConstraint().GetMagicConstantValue(),
					chain,
					magicAssetIdDefault,
				)
				if err != nil {
					return fmt.Errorf(
						"failed to resolve magic const: magic_const=%s",
						constraint.GetConstraint().GetMagicConstantValue().String(),
					)
				}

				comparer, err := makeComparer(resolvedAddr)
				if err != nil {
					return fmt.Errorf(
						"failed to build exact type from magic_const: resolved=%s",
						resolvedAddr,
					)
				}
				if comparer.Magic(actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare magic values: expected(resolved magic addr)=%v, actual(in tx)=%v",
					resolvedAddr,
					actual,
				)

			case types.ConstraintType_CONSTRAINT_TYPE_REGEXP:
				strVal := fmt.Sprintf("%v", actual)
				ok, err := regexp.MatchString(
					constraint.GetConstraint().GetRegexpValue(),
					strVal,
				)
				if err != nil {
					return fmt.Errorf("regexp match failed: expected=%v, actual=%v",
						constraint.GetConstraint().GetRegexpValue(), actual)
				}
				if ok {
					return nil
				}
				return fmt.Errorf("regexp value constraint failed: expected=%v, actual=%v",
					constraint.GetConstraint().GetRegexpValue(), actual)

			default:
				return fmt.Errorf("unknown constraint type: %s", constraint.GetConstraint().GetType())
			}
		}
	}
	return fmt.Errorf("arg not found: %s", expectedName)
}
