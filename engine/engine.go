package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/kaptinlin/jsonschema"
	"github.com/vultisig/recipes/internal/metarule"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
	"google.golang.org/protobuf/types/known/structpb"
)

type Engine struct {
	logger   *log.Logger
	registry *ChainEngineRegistry
}

func NewEngine() (*Engine, error) {
	reg, err := NewChainEngineRegistry()
	if err != nil {
		return nil, fmt.Errorf("failed to create registry: %w", err)
	}

	return &Engine{
		logger:   log.New(io.Discard, "", 0),
		registry: reg,
	}, nil
}

func (e *Engine) SetLogger(log *log.Logger) {
	e.logger = log
}

func (e *Engine) Evaluate(policy *types.Policy, chain common.Chain, txBytes []byte) (*types.Rule, error) {
	var errs []error
	for _, ruleRaw := range policy.GetRules() {
		if ruleRaw == nil {
			continue
		}

		rules, err := metarule.NewMetaRule().TryFormat(ruleRaw)
		if err != nil {
			return nil, fmt.Errorf("failed to format rule: %w", err)
		}

		for _, rule := range rules {
			resourcePathString := rule.GetResource()
			resourcePath, er := util.ParseResource(resourcePathString)
			if er != nil {
				e.logger.Printf(
					"Skipping rule %s: invalid resource path %s: %v",
					rule.GetId(),
					resourcePathString,
					er,
				)
				continue
			}

			if resourcePath.ChainId != strings.ToLower(chain.String()) {
				e.logger.Printf(
					"Skipping rule %s: target chain %s is not '%s'",
					rule.GetId(),
					resourcePath.ChainId,
					chain.String(),
				)
				continue
			}

			e.logger.Printf("Evaluating rule: %s: %s", rule.GetId(), resourcePathString)
			e.logger.Printf("Targeting: Chain='%s', Asset='%s', Function='%s'",
				resourcePath.ChainId, resourcePath.ProtocolId, resourcePath.FunctionId)

			// Get the appropriate engine for this chain
			chainEngine, er := e.registry.GetEngine(chain)
			if er != nil {
				e.logger.Printf("No engine available for chain %s: %v", chain.String(), er)
				continue
			}

			// Evaluate using the chain-specific engine
			er = chainEngine.Evaluate(rule, txBytes)
			if er != nil {
				errs = append(errs, fmt.Errorf("%s(%w)", resourcePathString, er))
				e.logger.Printf("Failed to evaluate tx for %s: %v", chain.String(), er)
				continue
			}

			e.logger.Printf("Tx validated for %s", chain.String())
			return rule, nil
		}
	}
	if len(errs) == 0 {
		return nil, errors.New("no matching rule")
	}

	var errStrs []string
	for _, err := range errs {
		errStrs = append(errStrs, err.Error())
	}
	return nil, fmt.Errorf("failed to evaluate tx: %s", strings.Join(errStrs, " "))
}

func (e *Engine) ValidatePolicyWithSchema(policy *types.Policy, schema *types.RecipeSchema) error {
	// Basic policy validation
	if len(policy.GetRules()) == 0 {
		return fmt.Errorf("policy has no rules")
	}

	if policy.GetId() != schema.GetPluginId() {
		return fmt.Errorf("policy ID %s does not match schema plugin ID %s",
			policy.GetId(), schema.GetPluginId())
	}

	if err := e.validateConfiguration(policy, schema); err != nil {
		return fmt.Errorf("configuration validation failed: %w", err)
	}

	// Build supported resources map from schema
	supportedResources := make(map[string]*types.ResourcePattern)
	for _, resourcePattern := range schema.GetSupportedResources() {
		if resourcePattern.GetResourcePath() != nil {
			resourcePath := resourcePattern.GetResourcePath().GetFull()
			supportedResources[resourcePath] = resourcePattern
		}
	}

	// Validate each rule
	for _, rule := range policy.GetRules() {
		if rule == nil {
			continue
		}

		// Check if resource is supported
		resourcePattern, supported := supportedResources[rule.GetResource()]
		if !supported {
			return fmt.Errorf("rule %s uses unsupported resource: %s",
				rule.GetId(), rule.GetResource())
		}

		// Validate parameter constraints against schema capabilities
		if err := e.validateParameterConstraints(rule, resourcePattern); err != nil {
			return fmt.Errorf("rule %s constraint validation failed: %w",
				rule.GetId(), err)
		}
	}

	return nil
}

func (e *Engine) validateConfiguration(policy *types.Policy, schema *types.RecipeSchema) error {
	emptyJson := []byte("{}")
	configJson, err := json.Marshal(schema.GetConfiguration())
	if err != nil {
		return fmt.Errorf("failed to marshal schema: %w", err)
	}
	if schema.GetConfiguration() == nil {
		configJson = emptyJson
	}

	compiler := jsonschema.NewCompiler()
	jsonSchema, err := compiler.Compile(configJson)
	if err != nil {
		return fmt.Errorf("failed to compile schema: %w", err)
	}

	configurationData := policy.GetConfiguration()
	if configurationData == nil {
		configurationData = &structpb.Struct{}
	}

	policyJson, err := json.Marshal(configurationData)
	if err != nil {
		return fmt.Errorf("failed to marshal policy: %w", err)
	}
	if configurationData == nil {
		policyJson = emptyJson
	}

	result := jsonSchema.Validate(policyJson)
	if !result.IsValid() {
		joinedErrors := ""
		errorMessages := []string{}
		for _, err := range result.Errors {
			errorMessages = append(errorMessages, err.Error())
		}
		joinedErrors = strings.Join(errorMessages, ", ")

		return fmt.Errorf("failed to validate policy configuration data: %s", joinedErrors)
	}

	return nil
}

func (e *Engine) validateParameterConstraints(rule *types.Rule, resourcePattern *types.ResourcePattern) error {
	// Build map of parameter capabilities from schema
	paramCapabilities := make(map[string]*types.ParameterConstraintCapability)
	for _, paramCap := range resourcePattern.GetParameterCapabilities() {
		paramCapabilities[paramCap.GetParameterName()] = paramCap
	}

	// Check each parameter constraint in the rule
	for _, paramConstraint := range rule.GetParameterConstraints() {
		paramName := paramConstraint.GetParameterName()

		// Check if parameter is supported
		paramCap, exists := paramCapabilities[paramName]
		if !exists {
			return fmt.Errorf("parameter %s not supported by schema", paramName)
		}

		// Check if constraint type is supported
		if paramCap.GetSupportedTypes() == paramConstraint.GetConstraint().GetType() {
			continue
		}

		return fmt.Errorf("parameter %s does not support constraint type %s",
			paramName, paramConstraint.GetConstraint().GetType())
	}

	return nil
}
