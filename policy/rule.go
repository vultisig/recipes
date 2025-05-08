package policy

import (
	"fmt"
	"strings"

	"github.com/vultisig/recipes/transaction"
)

// Effect represents the result of a rule evaluation
type Effect string

const (
	// EffectAllow indicates that the rule allows the action
	EffectAllow Effect = "ALLOW"

	// EffectDeny indicates that the rule denies the action
	EffectDeny Effect = "DENY"
)

// Authorization represents how a transaction should be authorized
type Authorization struct {
	// Type of authorization (prompt, allow, deny)
	Type string `json:"type"`

	// Message to display when prompting the user
	Message string `json:"message,omitempty"`
}

// Rule represents a permission rule in a policy
type Rule struct {
	// Resource is the resource pattern this rule applies to
	Resource string `json:"resource"`

	// Effect is the outcome if this rule matches (ALLOW or DENY)
	Effect Effect `json:"effect"`

	// Description provides human-readable details about this rule
	Description string `json:"description,omitempty"`

	// Constraints defines restrictions on the resource
	Constraints map[string]*Constraint `json:"constraints,omitempty"`

	// Authorization defines how this should be authorized (if needed)
	Authorization *Authorization `json:"authorization,omitempty"`
}

// NewRule creates a new rule with the given attributes
func NewRule(resource string, effect Effect, description string) *Rule {
	return &Rule{
		Resource:    resource,
		Effect:      effect,
		Description: description,
		Constraints: make(map[string]*Constraint),
	}
}

// AddConstraint adds a constraint to the rule
func (r *Rule) AddConstraint(name string, constraint *Constraint) {
	r.Constraints[name] = constraint
}

// SetAuthorization sets the authorization for this rule
func (r *Rule) SetAuthorization(authType string, message string) {
	r.Authorization = &Authorization{
		Type:    authType,
		Message: message,
	}
}

// Matches checks if a transaction matches this rule
func (r *Rule) Matches(tx interface{}) (bool, error) {
	// Convert to our transaction type
	transaction, ok := tx.(*transaction.Transaction)
	if !ok {
		return false, fmt.Errorf("expected transaction.Transaction, got %T", tx)
	}

	// Check if resource matches
	if !r.resourceMatches(transaction.ResourcePath) {
		return false, nil
	}

	// Check all constraints
	if len(r.Constraints) > 0 {
		for name, constraint := range r.Constraints {
			paramValue, ok := transaction.GetParam(name)
			if !ok && constraint.Required {
				return false, nil
			}

			if ok {
				matches, err := constraint.Validate(paramValue)
				if err != nil {
					return false, fmt.Errorf("error validating constraint %s: %w", name, err)
				}

				if !matches {
					return false, nil
				}
			}
		}
	}

	// All constraints satisfied
	return true, nil
}

// resourceMatches checks if a transaction resource matches this rule's resource pattern
func (r *Rule) resourceMatches(transactionResource string) bool {
	// Exact match
	if r.Resource == transactionResource {
		return true
	}

	// Wildcard match
	if strings.HasSuffix(r.Resource, "*") {
		prefix := r.Resource[:len(r.Resource)-1]
		return strings.HasPrefix(transactionResource, prefix)
	}

	return false
}
