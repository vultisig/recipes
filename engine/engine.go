package engine

import (
	"io"
	"log"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
)

type Engine struct {
	logger *log.Logger
}

func NewEngine() *Engine {
	// Turn off logging by default
	return &Engine{
		logger: log.New(io.Discard, "", 0),
	}
}

func (e *Engine) SetLogger(log *log.Logger) {
	e.logger = log
}

func (e *Engine) Evaluate(policy types.Policy, chain types.Chain, tx types.DecodedTransaction) (bool, *types.Rule, error) {
	for _, rule := range policy.GetRules() {
		if rule == nil {
			continue
		}

		resourcePathString := rule.GetResource()
		resourcePath, err := util.ParseResource(resourcePathString)
		if err != nil {
			e.logger.Printf("Skipping rule %s: invalid resource path %s: %v", rule.GetId(), resourcePathString, err)
			continue
		}

		if resourcePath.ChainId != chain.ID() {
			e.logger.Printf("Skipping rule %s: target chain %s is not '%s'", rule.GetId(), resourcePath.ChainId, chain.ID())
			continue
		}

		e.logger.Printf("Evaluating rule %s: %s", rule.GetId(), resourcePathString)
		e.logger.Printf("Targeting: Chain='%s', Asset='%s', Function='%s'",
			resourcePath.ChainId, resourcePath.ProtocolId, resourcePath.FunctionId)

		protocol, err := chain.GetProtocol(resourcePath.ProtocolId)
		if err != nil {
			e.logger.Printf("Skipping rule %s: Could not get protocol for asset '%s': %v", rule.GetId(), resourcePath.ProtocolId, err)
			continue
		}
		e.logger.Printf("Using protocol: %s (ID: %s)\n", protocol.Name(), protocol.ID())

		policyMatcher := &types.PolicyFunctionMatcher{
			FunctionID:  resourcePath.FunctionId,
			Constraints: rule.GetParameterConstraints(), // Use generated getter
		}

		matches, _, err := protocol.MatchFunctionCall(tx, policyMatcher)
		if err != nil {
			e.logger.Printf("Error during transaction matching for rule %s, function %s: %v", rule.GetId(), resourcePath.FunctionId, err)
			continue
		}

		if matches {
			e.logger.Printf("Transaction matches rule %s for function %s!\n", rule.GetId(), resourcePath.FunctionId)
			return true, rule, nil
		} else {
			e.logger.Printf("Transaction does not match rule %s for function %s!\n", rule.GetId(), resourcePath.FunctionId)
		}
	}

	return false, nil, nil
}