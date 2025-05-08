package policy

import (
	"encoding/json"
	"fmt"
	"time"
)

// Policy represents a set of rules that determine what a plugin is allowed to do
type Policy struct {
	// ID is a unique identifier for the policy
	ID string `json:"id"`

	// Name is a human-readable name for the policy
	Name string `json:"name"`

	// Description provides details about what the policy allows
	Description string `json:"description"`

	// Version is the policy version
	Version string `json:"version"`

	// Author is the identifier of the plugin developer
	Author string `json:"author"`

	// Rules is an ordered list of permission rules
	Rules []*Rule `json:"rules"`

	// CreatedAt is when the policy was created
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt is when the policy was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// NewPolicy creates a new policy with the given attributes
func NewPolicy(id, name, description, version, author string, rules []*Rule) *Policy {
	now := time.Now()
	return &Policy{
		ID:          id,
		Name:        name,
		Description: description,
		Version:     version,
		Author:      author,
		Rules:       rules,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// AddRule adds a new rule to the policy
func (p *Policy) AddRule(rule *Rule) {
	p.Rules = append(p.Rules, rule)
	p.UpdatedAt = time.Now()
}

// RemoveRule removes a rule at the specified index
func (p *Policy) RemoveRule(index int) error {
	if index < 0 || index >= len(p.Rules) {
		return fmt.Errorf("index out of range")
	}

	p.Rules = append(p.Rules[:index], p.Rules[index+1:]...)
	p.UpdatedAt = time.Now()
	return nil
}

// Evaluate evaluates a transaction against the policy
func (p *Policy) Evaluate(tx interface{}) (bool, string, error) {
	// Start with default deny
	allowed := false
	reason := "No matching rule found"

	// Check each rule in order
	for i, rule := range p.Rules {
		match, err := rule.Matches(tx)
		if err != nil {
			return false, "", fmt.Errorf("error evaluating rule %d: %w", i, err)
		}

		if match {
			allowed = rule.Effect == EffectAllow
			if allowed {
				reason = fmt.Sprintf("Allowed by rule %d: %s", i, rule.Resource)
			} else {
				reason = fmt.Sprintf("Denied by rule %d: %s", i, rule.Resource)
			}
			return allowed, reason, nil
		}
	}

	// No matching rule found, default to deny
	return allowed, reason, nil
}

// String returns a string representation of the policy
func (p *Policy) String() string {
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling policy: %w", err).Error()
	}
	return string(b)
}

// Marshal serializes the policy to JSON
func (p *Policy) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

// Unmarshal deserializes the policy from JSON
func Unmarshal(data []byte) (*Policy, error) {
	var p Policy
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}
	return &p, nil
}
