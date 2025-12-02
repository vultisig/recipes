package solana

import (
	"fmt"

	chainsolana "github.com/vultisig/recipes/chain/solana"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
)

type Solana struct {
	chain *chainsolana.Chain
	idl   map[protocolID]idl
}

func NewSolana() (*Solana, error) {
	idls, err := loadIDLDir()
	if err != nil {
		return nil, fmt.Errorf("failed to load idl dir: %w", err)
	}

	return &Solana{
		chain: chainsolana.NewChain(),
		idl:   idls,
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

	// Use chain package to parse transaction (using bytes directly)
	decodedTx, err := s.chain.ParseTransactionBytes(txBytes)
	if err != nil {
		return fmt.Errorf("failed to decode tx payload: %w", err)
	}

	parsedTx, ok := decodedTx.(*chainsolana.ParsedSolanaTransaction)
	if !ok {
		return fmt.Errorf("unexpected transaction type: %T", decodedTx)
	}

	tx := parsedTx.GetTransaction()
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

	idlProtocolSchema, ok := s.idl[r.ProtocolId]
	if !ok {
		return fmt.Errorf("unknown protocol id: %s", r.ProtocolId)
	}
	idlInstSchema, err := findInstruction(idlProtocolSchema.Instructions, r.FunctionId)
	if err != nil {
		return fmt.Errorf("failed to find instruction: %w", err)
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

	err = assertTarget(r, rule.GetTarget(), programID)
	if err != nil {
		return fmt.Errorf("failed to assert target: %w", err)
	}

	err = assertAccounts(rule.GetParameterConstraints(), tx.Message, idlInstSchema.Accounts)
	if err != nil {
		return fmt.Errorf("failed to assert accounts: %w", err)
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
