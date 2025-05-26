package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vultisig/recipes/chain"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	policyPath := flag.String("policy", "", "Path to the policy.json file")
	txHex := flag.String("tx", "", "Hex-encoded transaction string")
	chainID := flag.String("chain", "ethereum", "Chain ID (e.g., 'ethereum', 'bitcoin')")
	flag.Parse()

	if *policyPath == "" {
		log.Fatal("Error: -policy flag is required")
	}
	if *txHex == "" {
		log.Fatal("Error: -tx flag is required")
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
	fmt.Printf("Successfully loaded policy: %s (Name: %s)\n", policy.GetId(), policy.GetName())

	// 2. Initialize Chain based on chainID flag
	selectedChain, err := chain.GetChain(*chainID)
	if err != nil {
		log.Fatalf("Failed to get chain %s: %v", *chainID, err)
	}
	fmt.Printf("Using chain: %s (%s)\n", selectedChain.ID(), selectedChain.Name())

	// If it's Ethereum and tokens or specific ABIs are involved, they might need to be loaded.
	if *chainID == "ethereum" {
		// For example, if a policy refers to "USDC" and the chain needs its token info/ABI:
		// tokenListBytes, _ := os.ReadFile("path/to/your/ethereum_tokenlist.json")
		// if tokenListBytes != nil { selectedChain.(*ethereum.Ethereum).LoadTokenList(tokenListBytes) }
	}

	// 3. Iterate through Policy Rules and Verify Transaction
	transactionAllowedByPolicy := false
	var matchingRuleID string
	var matchingFunctionID string
	var extractedTxParams map[string]interface{}

	// Attempt to parse the transaction once, as it's the same for all rules on this chain.
	decodedTx, err := selectedChain.ParseTransaction(*txHex)
	if err != nil {
		log.Fatalf("Failed to parse %s transaction '%s': %v. Cannot proceed.", *chainID, *txHex, err)
	}
	fmt.Printf("Successfully parsed transaction: Hash=%s, From=%s, To=%s, Value=%s\n",
		decodedTx.Hash(), decodedTx.From(), decodedTx.To(), decodedTx.Value().String())

	for _, rule := range policy.GetRules() { // Use generated GetRules() getter
		if rule == nil { // Defensive check
			continue
		}
		resourcePathString := rule.GetResource() // Use generated getter
		resourcePath, err := util.ParseResource(resourcePathString)
		if err != nil {
			log.Printf("Skipping rule %s: invalid resource path %s: %v", rule.GetId(), resourcePathString, err)
			continue
		}

		if resourcePath.ChainId != *chainID {
			log.Printf("Skipping rule %s: target chain %s is not '%s'", rule.GetId(), resourcePath.ChainId, *chainID)
			continue
		}

		fmt.Printf("\nEvaluating rule: %s (Description: %s)\n", rule.GetId(), rule.GetDescription())
		fmt.Printf("  Targeting: Chain='%s', Asset='%s', Function='%s'\n",
			resourcePath.ChainId, resourcePath.ProtocolId, resourcePath.FunctionId)

		protocol, err := selectedChain.GetProtocol(resourcePath.ProtocolId)
		if err != nil {
			log.Printf("  Skipping rule %s: Could not get protocol for asset '%s': %v", rule.GetId(), resourcePath.ProtocolId, err)
			continue
		}
		fmt.Printf("  Using protocol: %s (ID: %s)\n", protocol.Name(), protocol.ID())

		policyMatcher := &types.PolicyFunctionMatcher{
			FunctionID:  resourcePath.FunctionId,
			Constraints: rule.GetParameterConstraints(), // Use generated getter
		}

		matches, params, err := protocol.MatchFunctionCall(decodedTx, policyMatcher)
		if err != nil {
			log.Printf("  Error during transaction matching for rule %s, function %s: %v", rule.GetId(), resourcePath.FunctionId, err)
			continue
		}

		if matches {
			fmt.Printf("  SUCCESS: Transaction matches rule %s for function %s!\n", rule.GetId(), resourcePath.FunctionId)
			transactionAllowedByPolicy = true
			matchingRuleID = rule.GetId()
			matchingFunctionID = resourcePath.FunctionId
			extractedTxParams = params
			break
		} else {
			fmt.Printf("  INFO: Transaction does not match rule %s for function %s.\n", rule.GetId(), resourcePath.FunctionId)
		}
	}

	// 4. Output Result
	fmt.Println("\n--- Policy Check Result ---")
	if transactionAllowedByPolicy {
		fmt.Printf("Transaction IS ALLOWED by the policy.\n")
		fmt.Printf("Matched Rule ID: %s\n", matchingRuleID)
		fmt.Printf("Matched Function ID: %s\n", matchingFunctionID)
		fmt.Printf("Extracted Transaction Parameters:\n")
		for key, val := range extractedTxParams {
			fmt.Printf("  %s: %v\n", key, val)
		}
	} else {
		fmt.Printf("Transaction IS NOT ALLOWED by the policy.\n")
	}

	// Example of using the chain registry if it were populated
	_ = chain.NewRegistry() // Placeholder for registry usage if needed in future

}
