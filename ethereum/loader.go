package ethereum

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vultisig/recipes/protocol"
	"github.com/vultisig/recipes/types"
)

// DefaultERC20ABIPath is the path to the default ERC20 ABI definition
const DefaultERC20ABIPath = "abis/erc20.json"

// RegisterEthereumProtocols registers all Ethereum protocols
func RegisterEthereumProtocols(ethereumChain *Ethereum, erc20ABI *ABI) error {
	// Register native ETH protocol
	protocol.RegisterProtocol(NewETH())

	// Register protocol validators
	registerProtocolValidators()

	// Register token protocols from token list with dynamically loaded ERC20 functions
	if ethereumChain.tokenList != nil && erc20ABI != nil {
		for _, token := range ethereumChain.tokenList.Tokens {
			// Create token-specific ERC20 protocol using ERC20 ABI
			tokenProtocolID := token.Symbol
			tokenName := token.Name
			tokenDescription := fmt.Sprintf("%s token on Ethereum (ERC20)", token.Name)

			// Generate token-specific functions from ERC20 ABI
			tokenProtocol := NewABIProtocolWithCustomization(
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
		}
	}

	// Register ABI protocols (non-token contracts)
	for name, abi := range ethereumChain.abiRegistry {
		// Skip registering the ERC20 ABI directly, as it's used for tokens
		if name != "erc20" {
			description := fmt.Sprintf("Protocol generated from %s ABI", name)

			// Use the generic ABI protocol creation - validators are automatically applied
			abiProtocol := NewABIProtocol(name, name, description, abi)
			protocol.RegisterProtocol(abiProtocol)
		}
	}

	return nil
}

// registerProtocolValidators registers all available protocol validators
func registerProtocolValidators() {

	// Register Uniswap v2 validator
	uniswapValidator := NewUniswapV2Validator()
	GlobalValidatorRegistry.RegisterValidator(uniswapValidator.GetProtocolID(), uniswapValidator)

}

// LoadTokenListFromFile loads a token list from a file
func LoadTokenListFromFile(path string, ethereumChain *Ethereum) error {
	// Read the token list file
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading token list file: %w", err)
	}

	// Parse the token list
	return ethereumChain.LoadTokenList(data)
}

// LoadABIFromFile loads an ABI from a file
func LoadABIFromFile(path string, name string, ethereumChain *Ethereum) error {
	// Read the ABI file
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading ABI file: %w", err)
	}

	// Parse the ABI
	return ethereumChain.LoadABI(name, data)
}

// LoadABIsFromDirectory loads all ABIs from a directory
func LoadABIsFromDirectory(dirPath string, ethereumChain *Ethereum) error {
	// Get all .json files in the directory
	files, err := filepath.Glob(filepath.Join(dirPath, "*.json"))
	if err != nil {
		return fmt.Errorf("error finding ABI files: %w", err)
	}

	for _, filePath := range files {
		// Use the filename (without extension) as the ABI name
		name := filepath.Base(filePath)
		name = name[:len(name)-len(filepath.Ext(name))]

		// Load the ABI
		if err := LoadABIFromFile(filePath, name, ethereumChain); err != nil {
			return fmt.Errorf("error loading ABI from %s: %w", filePath, err)
		}
	}

	return nil
}

// InitEthereum initializes Ethereum with token list and ABIs
func InitEthereum(tokenListPath string, abiDirPath string) (types.Chain, error) {
	// Create Ethereum chain
	ethereumChain := NewEthereum().(*Ethereum)

	// Load token list if provided
	if tokenListPath != "" {
		if err := LoadTokenListFromFile(tokenListPath, ethereumChain); err != nil {
			return nil, fmt.Errorf("error loading token list: %w", err)
		}
	}

	// Load ABIs if directory provided
	if abiDirPath != "" {
		if err := LoadABIsFromDirectory(abiDirPath, ethereumChain); err != nil {
			return nil, fmt.Errorf("error loading ABIs: %w", err)
		}
	}

	// Get the ERC20 ABI for token protocols
	var erc20ABI *ABI
	if abi, ok := ethereumChain.GetABI("erc20"); ok {
		erc20ABI = abi
	} else {
		// Try to load the default ERC20 ABI
		fmt.Println("Searching for ERC20 ABI in", abiDirPath)
		erc20Path := filepath.Join(abiDirPath, "erc20.json")
		if _, err := os.Stat(erc20Path); err == nil {
			if err := LoadABIFromFile(erc20Path, "erc20", ethereumChain); err != nil {
				return nil, fmt.Errorf("error loading ERC20 ABI: %w", err)
			}
			erc20ABI, _ = ethereumChain.GetABI("erc20")
		} else {
			return nil, fmt.Errorf("ERC20 ABI not found, required for token protocols")
		}
	}

	// Register protocols
	if err := RegisterEthereumProtocols(ethereumChain, erc20ABI); err != nil {
		return nil, fmt.Errorf("error registering protocols: %w", err)
	}

	return ethereumChain, nil
}
