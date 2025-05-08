package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vultisig/recipes/internal/generator"
)

func main() {
	// Define command-line flags
	outputPath := flag.String("output", "RESOURCES.md", "Output file path")
	flag.Parse()

	// Create generator
	gen, err := generator.NewGenerator(*outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating generator: %v\n", err)
		os.Exit(1)
	}

	// Generate documentation
	if err := gen.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating documentation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Documentation generated successfully at %s\n", *outputPath)
}
