package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vultisig/recipes/arbitrum"
	"github.com/vultisig/recipes/base"
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
	ethereumChain := ethereum.NewEthereum()
	err := ethereumChain.InitEthereum(tokenListPath, abiDirPath)
	if err != nil {
		return fmt.Errorf("error initializing Ethereum: %w", err)
	}
	chain.RegisterChain(ethereumChain)

	// Register Arbitrum chain with token list and ABIs if provided
	arbitrumChain := arbitrum.NewArbitrum()
	err = arbitrumChain.InitEthereum(tokenListPath, abiDirPath)
	if err != nil {
		return fmt.Errorf("error initializing Arbitrum: %w", err)
	}
	chain.RegisterChain(arbitrumChain)

	// Register Base chain with token list and ABIs if provided
	baseChain := base.NewBase()
	err = baseChain.InitEthereum(tokenListPath, abiDirPath)
	if err != nil {
		return fmt.Errorf("error initializing Base: %w", err)
	}
	chain.RegisterChain(baseChain)

	return nil
}
