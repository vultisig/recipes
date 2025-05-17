package generator

import (
	"fmt"
	"os"
	"sort"
	"text/template"

	"github.com/vultisig/recipes/chain"
	"github.com/vultisig/recipes/protocol"
	"github.com/vultisig/recipes/types"
)

// ResourceDoc represents a resource for documentation
type ResourceDoc struct {
	ResourcePath string
	ChainName    string
	ProtocolName string
	FunctionName string
	Description  string
	Parameters   []*types.FunctionParam
}

// Generator generates documentation for available resources
type Generator struct {
	outputPath string
	tmpl       *template.Template
}

// NewGenerator creates a new documentation generator
func NewGenerator(outputPath string) (*Generator, error) {
	// Define template functions
	funcMap := template.FuncMap{
		"last": func(x int, a []*types.FunctionParam) bool {
			return x == len(a)-1
		},
	}

	// Parse the template with functions
	tmpl, err := template.New("doc").Funcs(funcMap).Parse(docTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %w", err)
	}

	return &Generator{
		outputPath: outputPath,
		tmpl:       tmpl,
	}, nil
}

// Generate creates the documentation
func (g *Generator) Generate() error {
	// Create output file
	file, err := os.Create(g.outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer file.Close()

	// Collect resource documentation
	resources := make([]*ResourceDoc, 0)

	// Iterate through chains
	for _, c := range chain.DefaultRegistry.List() {
		// Get protocols for this chain
		protocols := protocol.ListProtocolsByChain(c.ID())

		for _, p := range protocols {
			// Get functions for this protocol
			functions := p.Functions()

			for _, f := range functions {
				// Create resource path
				resourcePath := fmt.Sprintf("%s.%s.%s", c.ID(), p.ID(), f.ID)

				// Create resource doc
				resourceDoc := &ResourceDoc{
					ResourcePath: resourcePath,
					ChainName:    c.Name(),
					ProtocolName: p.Name(),
					FunctionName: f.Name,
					Description:  f.Description,
					Parameters:   f.Parameters,
				}

				resources = append(resources, resourceDoc)
			}
		}
	}

	// sort resources by ResourcePath
	sort.Slice(resources, func(i, j int) bool {
		return resources[i].ResourcePath < resources[j].ResourcePath
	})

	// Execute template
	data := map[string]interface{}{
		"Resources": resources,
	}

	if err := g.tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}

// Documentation template
const docTemplate = `# Cryptocurrency Wallet Policy Resources

This document lists all available resources that can be used in your wallet plugin policies.
Each resource represents an action that can be performed by a plugin, subject to policy constraints.

## Available Resources

{{ range .Resources }}
### {{ .ResourcePath }}

**Chain:** {{ .ChainName }}  
**Protocol:** {{ .ProtocolName }}  
**Function:** {{ .FunctionName }}  

{{ .Description }}

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
{{ range .Parameters }}| {{ .Name }} | {{ .Type }} | {{ .Description }} |
{{ end }}

**Example Policy Rule:**

` + "```json" + `
{
  "resource": "{{ .ResourcePath }}",
  "effect": "ALLOW",
  "constraints": {
{{ range $index, $param := .Parameters }}    "{{ $param.Name }}": {
      "type": "fixed",
      "value": "example_value"
    }{{ if not (last $index $.Parameters) }},{{ end }}
{{ end }}
  }
}
` + "```" + `

{{ end }}

## Using Wildcards

You can use wildcards in resource paths to match multiple resources:

* ` + "`chain.*.*`" + ` - Match all functions in all protocols on a chain
* ` + "`chain.protocol.*`" + ` - Match all functions in a specific protocol on a chain

## Available Constraint Types

| Type | Description | Example |
|------|-------------|---------|
| fixed | Exact match required | ` + "`" + `{"type": "fixed", "value": "0.1"}` + "`" + ` |
| max | Maximum value | ` + "`" + `{"type": "max", "value": "1.0"}` + "`" + ` |
| min | Minimum value | ` + "`" + `{"type": "min", "value": "0.01"}` + "`" + ` |
| range | Value within range | ` + "`" + `{"type": "range", "value": {"min": "0.1", "max": "1.0"}}` + "`" + ` |
| whitelist | Value from allowed list | ` + "`" + `{"type": "whitelist", "value": ["address1", "address2"]}` + "`" + ` |
| max_per_period | Limit actions per time period | ` + "`" + `{"type": "max_per_period", "value": 3, "period": "day"}` + "`" + ` |
`
