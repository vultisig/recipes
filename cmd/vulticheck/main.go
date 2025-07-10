package main

import (
	"flag"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/vultisig/recipes/chain"
	"github.com/vultisig/recipes/engine"
	"github.com/vultisig/recipes/types"
)

func main() {
	policyPath := flag.String("policy", "", "Path to the policy.json file")
	schemaPath := flag.String("schema", "", "Path to the schema.json file")
	txHex := flag.String("tx", "", "Hex-encoded transaction string")
	chainID := flag.String("chain", "ethereum", "Chain ID (e.g., 'ethereum', 'bitcoin')")
	flag.Parse()

	if *policyPath == "" {
		log.Fatal("Error: -policy flag is required")
	}

	// 1. Load and Parse Policy
	policyFileBytes, err := os.ReadFile(*policyPath)
	if err != nil {
		log.Fatalf("Failed to read policy file %s: %v", *policyPath, err)
	}

	var policy types.Policy
	// protojson.UnmarshalOptions{DiscardUnknown: true} can be useful if policy.json has extra fields
	if err := protojson.Unmarshal(policyFileBytes, &policy); err != nil {
		log.Fatalf("Failed to unmarshal policy JSON: %v", err)
	}
	log.Printf("Successfully loaded policy: %s (Name: %s)\n", policy.GetId(), policy.GetName())

	// 2. Initialize Engine
	eng := engine.NewEngine()
	eng.SetLogger(log.Default())

	// 3. Load Schema and ValidatePolicy (if path given)
	var schema *types.RecipeSchema
	if *schemaPath != "" {
		schemaFileBytes, err := os.ReadFile(*schemaPath)
		if err != nil {
			log.Fatalf("Failed to read schema file %s: %v", *schemaPath, err)
		}

		schema = &types.RecipeSchema{}
		if err := protojson.Unmarshal(schemaFileBytes, schema); err != nil {
			log.Fatalf("Failed to unmarshal schema JSON: %v", err)
		}
		log.Printf("Successfully loaded schema for plugin: %s (Version: %d)",
			schema.GetPluginName(), schema.GetPluginVersion())

		err = eng.ValidatePolicyWithSchema(&policy, schema)
		if err != nil {
			log.Fatalf("Failed to validate policy: %v", err)
		}
		log.Printf("Successfully validated policy: %s", policy.GetName())
	}

	// 4. Decode and Evaluate transaction (if tx flag given)
	if *txHex != "" {
		// Initialise chain
		selectedChain, err := chain.GetChain(*chainID)
		if err != nil {
			log.Fatalf("Failed to get chain %s: %v", *chainID, err)
		}
		log.Printf("Using chain: %s (%s)\n", selectedChain.ID(), selectedChain.Name())

		// Attempt to parse the transaction once, as it's the same for all rules on this chain.
		decodedTx, err := selectedChain.ParseTransaction(*txHex)
		if err != nil {
			log.Fatalf("Failed to parse %s transaction '%s': %v. Cannot proceed.", *chainID, *txHex, err)
		}
		log.Printf("Successfully parsed transaction: Hash=%s, From=%s, To=%s, Value=%s\n",
			decodedTx.Hash(), decodedTx.From(), decodedTx.To(), decodedTx.Value().String())

		transactionAllowedByPolicy, matchingRule, err := eng.Evaluate(&policy, selectedChain, decodedTx)
		if err != nil {
			log.Printf("Failed to evaluate transaction: %v", err)
		}

		log.Printf("Transaction allowed by policy: %t", transactionAllowedByPolicy)
		log.Printf("Matching rule: %s", matchingRule.GetId())
	}

	// 5. Final check, either txHex or Schema should be present
	if *schemaPath == "" && *txHex == "" {
		log.Fatal("Must provide -schema or -tx")
	}
}
