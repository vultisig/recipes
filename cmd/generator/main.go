package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vultisig/recipes/chain"
	"github.com/vultisig/recipes/ethereum"
	"github.com/vultisig/recipes/internal/generator"
)

func main() {
	// Define command-line flags
	outputPath := flag.String("output", "RESOURCES.md", "Output file path")
	tokenListPath := flag.String("tokenlist", "tokenlist.json", "Path to token list JSON file")
	abiDirPath := flag.String("abi", "abi", "Path to directory containing ABI JSON files")
	flag.Parse()

	// Register chains and protocols
	if err := registerChains(*tokenListPath, *abiDirPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error registering chains: %v\n", err)
		os.Exit(1)
	}

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

func registerChains(tokenListPath, abiDirPath string) error {
	// Register Ethereum chain with token list and ABIs if provided
	ethereumChain, err := ethereum.InitEthereum(tokenListPath, abiDirPath)
	if err != nil {
		return fmt.Errorf("error initializing Ethereum: %w", err)
	}
	chain.RegisterChain(ethereumChain)

	return nil
}
