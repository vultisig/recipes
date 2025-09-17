package solana

import (
	"fmt"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
)

type Solana struct {
	nativeSymbol string
	idl          map[protocolID]idl
}

func NewSolana() (*Solana, error) {
	idls, err := loadIDLDir()
	if err != nil {
		return nil, fmt.Errorf("failed to load idl dir: %w", err)
	}

	return &Solana{
		idl: idls,
	}, nil
}

func (s *Solana) Supports(chain common.Chain) bool {
	return chain == common.Solana
}

func (s *Solana) Evaluate(rule *types.Rule, txBytes []byte) error {
	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
	}

	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}

	tx, err := solana.TransactionFromDecoder(bin.NewBorshDecoder(txBytes))
	if err != nil {
		return fmt.Errorf("failed to decode tx payload: %w", err)
	}

	if len(tx.Message.Instructions) != 1 {
		return fmt.Errorf(
			"only single instruction transactions are allowed, got %d instructions",
			len(tx.Message.Instructions),
		)
	}
	inst := tx.Message.Instructions[0]

	programID, err := tx.ResolveProgramIDIndex(inst.ProgramIDIndex)
	if err != nil {
		return fmt.Errorf("failed to resolve program id: %w", err)
	}

	err = assertTarget(r, rule.GetTarget(), programID)
	if err != nil {
		return fmt.Errorf("failed to assert target: %w", err)
	}

	idlProtocolSchema, ok := s.idl[r.ProtocolId]
	if !ok {
		return fmt.Errorf("unknown protocol id: %s", r.ProtocolId)
	}
	idlInstSchema, err := findInstruction(idlProtocolSchema.Instructions, r.FunctionId)
	if err != nil {
		return fmt.Errorf("failed to find instruction: %w", err)
	}
	if len(idlInstSchema.Metadata.Discriminator) == 0 {
		return fmt.Errorf("instruction %s.%s is missing discriminator in metadata", r.ProtocolId, r.FunctionId)
	}

	err = assertAccounts(rule.GetParameterConstraints(), tx.Message, idlInstSchema.Accounts)
	if err != nil {
		return fmt.Errorf("failed to assert accounts: %w", err)
	}

	err = assertArgs(
		rule.GetParameterConstraints(),
		inst.Data,
		idlInstSchema.Args,
		idlInstSchema.Metadata.Discriminator,
	)
	if err != nil {
		return fmt.Errorf("failed to assert args: %w", err)
	}
	return nil
}

func findInstruction(insts []idlInstruction, name string) (idlInstruction, error) {
	for _, inst := range insts {
		if inst.Name == name {
			return inst, nil
		}
	}
	return idlInstruction{}, fmt.Errorf("instruction not found: %s", name)
}
