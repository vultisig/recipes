package types

// PolicyFunctionMatcher provides criteria for matching a function call against a policy.
// It encapsulates the function to be matched and the constraints on its parameters.
type PolicyFunctionMatcher struct {
	FunctionID string // The identifier of the function to match (e.g., "transfer", "approve").
	// Constraints should hold the specific parameter constraints relevant to this function call.
	Constraints []*ParameterConstraint // Changed from []*Constraint to []*ParameterConstraint
}
