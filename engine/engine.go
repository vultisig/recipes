package engine

import (
	"fmt"
	"encoding/json"
	"io"
	"log"
	"strings"

	"github.com/kaptinlin/jsonschema"
	"google.golang.org/protobuf/types/known/structpb"
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

func (e *Engine) Evaluate(policy *types.Policy, chain types.Chain, tx types.DecodedTransaction) (bool, *types.Rule, error) {
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
			FunctionID:   resourcePath.FunctionId,
			Constraints:  rule.GetParameterConstraints(), // Use generated getter
			ResourcePath: resourcePath,
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
	configJson, err := json.Marshal(schema.GetConfiguration())
	if err != nil {
		return fmt.Errorf("failed to marshal schema: %w", err)
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
		constraintType := paramConstraint.GetConstraint().GetType()
		supported := false
		for _, supportedType := range paramCap.GetSupportedTypes() {
			if supportedType == constraintType {
				supported = true
				break
			}
		}

		if !supported {
			return fmt.Errorf("parameter %s does not support constraint type %s",
				paramName, constraintType.String())
		}
	}

	return nil
}
