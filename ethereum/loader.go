package ethereum

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/vultisig/recipes/protocol"
	"github.com/vultisig/recipes/types"
)

// DefaultERC20ABIPath is the path to the default ERC20 ABI definition
const DefaultERC20ABIPath = "abis/erc20.json"

// sync.Once to ensure validators are registered only once
var registerValidatorsOnce sync.Once
var registerValidatorsError error

// RegisterEthereumProtocols registers all Ethereum protocols
func (e *Ethereum) RegisterEthereumProtocols(erc20ABI *ABI) error {
	// Register native ETH protocol
	protocol.RegisterProtocol(NewETH(e.ID(), e.Name()))

	fmt.Printf("Created ETH protocol with chainID: %s\n", e.ID())
	// Register protocol validators only once
	registerValidatorsOnce.Do(func() {
		registerValidatorsError = registerProtocolValidators()
	})

	// Check if validator registration failed
	if registerValidatorsError != nil {
		return fmt.Errorf("failed to register protocol validators: %w", registerValidatorsError)
	}

	// Register token protocols from token list with dynamically loaded ERC20 functions
	if e.tokenList != nil && erc20ABI != nil {
		for _, token := range e.tokenList.Tokens {
			// Create token-specific ERC20 protocol using ERC20 ABI
			tokenProtocolID := token.Symbol
			tokenName := token.Name
			tokenDescription := fmt.Sprintf("%s token on Ethereum (ERC20)", token.Name)

			// Generate token-specific functions from ERC20 ABI
			tokenProtocol := NewABIProtocolWithCustomization(
				e.ID(),
				tokenProtocolID,
				tokenName,
				tokenDescription,
				erc20ABI,
				// Function customizer to make ERC20 functions token-specific
				func(f *types.Function, abiFunc *ABIFunction) {
					// Customize function name and description for the specific token
					f.Name = fmt.Sprintf("%s %s", tokenName, abiFunc.Name)

					// Customize description
					switch abiFunc.Name {
					case "transfer":
						f.Description = fmt.Sprintf("Transfer %s tokens to another address", tokenName)
					case "transferFrom":
						f.Description = fmt.Sprintf("Transfer %s tokens from one address to another", tokenName)
					case "approve":
						f.Description = fmt.Sprintf("Approve an address to spend %s tokens", tokenName)
					case "balanceOf":
						f.Description = fmt.Sprintf("Get the %s token balance of an address", tokenName)
					case "allowance":
						f.Description = fmt.Sprintf("Get the amount of %s tokens allowed to be spent by an address", tokenName)
					default:
						f.Description = fmt.Sprintf("Call the %s function on %s token", abiFunc.Name, tokenName)
					}

					// Customize parameter descriptions for specific token
					for _, param := range f.Parameters {
						if param.Name == "amount" || param.Name == "value" {
							param.Description = fmt.Sprintf("The amount of %s tokens", tokenName)
						}
					}
				},
			)

			protocol.RegisterProtocol(tokenProtocol)
			fmt.Printf("Created Token protocol with chainID: %s, name: %s\n", tokenProtocol.ChainID(), tokenProtocol.Name())
		}
	}

	// Register ABI protocols (non-token contracts)
	for name, abi := range e.abiRegistry {
		// Skip registering the ERC20 ABI directly, as it's used for tokens
		if name != "erc20" {
			description := fmt.Sprintf("Protocol generated from %s ABI", name)

			// Use the generic ABI protocol creation - validators are automatically applied
			abiProtocol := NewABIProtocol(e.ID(), name, name, description, abi)
			protocol.RegisterProtocol(abiProtocol)
		}
	}

	return nil
}

// registerProtocolValidators registers all available protocol validators
func registerProtocolValidators() error {
	// Register Uniswap v2 validator
	uniswapValidator := NewUniswapV2Validator()
	if err := GlobalValidatorRegistry.RegisterValidator(uniswapValidator.GetProtocolID(), uniswapValidator); err != nil {
		return fmt.Errorf("failed to register Uniswap v2 validator: %w", err)
	}

	return nil
}

// LoadTokenListFromFile loads a token list from a file
func (e *Ethereum) LoadTokenListFromFile(path string) error {
	// Read the token list file
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading token list file: %w", err)
	}

	// Parse the token list
	return e.LoadTokenList(data)
}

// LoadABIFromFile loads an ABI from a file
func (e *Ethereum) LoadABIFromFile(path string, name string) error {
	// Read the ABI file
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading ABI file: %w", err)
	}

	// Parse the ABI
	return e.LoadABI(name, data)
}

// LoadABIsFromDirectory loads all ABIs from a directory
func (e *Ethereum) LoadABIsFromDirectory(dirPath string) error {
	// Get all .json files in the directory
	files, err := filepath.Glob(filepath.Join(dirPath, "*.json"))
	if err != nil {
		return fmt.Errorf("error finding ABI files: %w", err)
	}

	for _, filePath := range files {
		// Use the filename (without extension) as the ABI name
		name := filepath.Base(filePath)
		name = name[:len(name)-len(filepath.Ext(name))]

		if !e.supportedABIs[name] {
			fmt.Printf("Skipping ABI %s for chain %s (not supported)\n", name, e.ID())
			continue
		}
		// Load the ABI
		if err := e.LoadABIFromFile(filePath, name); err != nil {
			return fmt.Errorf("error loading ABI from %s: %w", filePath, err)
		}
	}

	return nil
}

// InitEthereum initializes Ethereum with token list and ABIs
func (e *Ethereum) InitEthereum(tokenListPath string, abiDirPath string) error {

	// Load token list if provided
	if tokenListPath != "" {
		if err := e.LoadTokenListFromFile(tokenListPath); err != nil {
			return fmt.Errorf("error loading token list: %w", err)
		}
	}

	// Load ABIs if directory provided
	if abiDirPath != "" {
		if err := e.LoadABIsFromDirectory(abiDirPath); err != nil {
			return fmt.Errorf("error loading ABIs: %w", err)
		}
	}

	// Get the ERC20 ABI for token protocols
	var erc20ABI *ABI
	if abi, ok := e.GetABI("erc20"); ok {
		erc20ABI = abi
	} else {
		// Try to load the default ERC20 ABI
		fmt.Println("Searching for ERC20 ABI in", abiDirPath)
		erc20Path := filepath.Join(abiDirPath, "erc20.json")
		if _, err := os.Stat(erc20Path); err == nil {
			if err := e.LoadABIFromFile(erc20Path, "erc20"); err != nil {
				return fmt.Errorf("error loading ERC20 ABI: %w", err)
			}
			erc20ABI, _ = e.GetABI("erc20")
		} else {
			return fmt.Errorf("ERC20 ABI not found, required for token protocols")
		}
	}

	// Register protocols
	if err := e.RegisterEthereumProtocols(erc20ABI); err != nil {
		return fmt.Errorf("error registering protocols: %w", err)
	}

	return nil
}
