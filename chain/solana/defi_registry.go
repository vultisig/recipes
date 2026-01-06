package solana

import "github.com/gagliardetto/solana-go"

// DeFi protocol program IDs on Solana

// LP Protocol Program IDs
var (
	// Raydium CLMM (Concentrated Liquidity Market Maker)
	RaydiumCLMMProgramID = solana.MustPublicKeyFromBase58("CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK")

	// Orca Whirlpools
	OrcaWhirlpoolProgramID = solana.MustPublicKeyFromBase58("whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc")

	// Meteora Pools
	MeteoraProgramID = solana.MustPublicKeyFromBase58("LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo")
)

// Lending Protocol Program IDs
var (
	// Kamino Finance Lending
	KaminoLendingProgramID = solana.MustPublicKeyFromBase58("KLend2g3cP87ber41GYr72yfE9j6eBJYwRqVNMi6mHL")

	// Marginfi
	MarginfiProgramID = solana.MustPublicKeyFromBase58("MFv2hWf31Z9kbCa1snEPYctwafyhdvnV7FZnsebVacA")

	// Solend
	SolendProgramID = solana.MustPublicKeyFromBase58("So1endDq2YkqhipRh3WViPa8hdiSpxWy6z3Z6tMCpAo")
)

// Perps Protocol Program IDs
var (
	// Jupiter Perpetuals
	JupiterPerpsProgramID = solana.MustPublicKeyFromBase58("PERPHjGBqRHArX4DySjwM6UJHiR3sWAatqfdBS2qQJu")

	// Drift Protocol
	DriftProgramID = solana.MustPublicKeyFromBase58("dRiftyHA39MWEi3m9aunc5MzRF1JYuBsbn6VPcn33UH")
)

// Protocol enum for routing
type DeFiProtocol string

const (
	// LP Protocols
	ProtocolRaydium  DeFiProtocol = "raydium"
	ProtocolOrca     DeFiProtocol = "orca"
	ProtocolMeteora  DeFiProtocol = "meteora"

	// Lending Protocols
	ProtocolKamino   DeFiProtocol = "kamino"
	ProtocolMarginfi DeFiProtocol = "marginfi"
	ProtocolSolend   DeFiProtocol = "solend"

	// Perps Protocols
	ProtocolJupiterPerps DeFiProtocol = "jupiter_perps"
	ProtocolDrift        DeFiProtocol = "drift"
)

// GetLPProgramID returns the program ID for an LP protocol
func GetLPProgramID(protocol DeFiProtocol) (solana.PublicKey, error) {
	switch protocol {
	case ProtocolRaydium:
		return RaydiumCLMMProgramID, nil
	case ProtocolOrca:
		return OrcaWhirlpoolProgramID, nil
	case ProtocolMeteora:
		return MeteoraProgramID, nil
	default:
		return solana.PublicKey{}, nil
	}
}

// GetLendingProgramID returns the program ID for a lending protocol
func GetLendingProgramID(protocol DeFiProtocol) (solana.PublicKey, error) {
	switch protocol {
	case ProtocolKamino:
		return KaminoLendingProgramID, nil
	case ProtocolMarginfi:
		return MarginfiProgramID, nil
	case ProtocolSolend:
		return SolendProgramID, nil
	default:
		return solana.PublicKey{}, nil
	}
}

// GetPerpsProgramID returns the program ID for a perps protocol
func GetPerpsProgramID(protocol DeFiProtocol) (solana.PublicKey, error) {
	switch protocol {
	case ProtocolJupiterPerps:
		return JupiterPerpsProgramID, nil
	case ProtocolDrift:
		return DriftProgramID, nil
	default:
		return solana.PublicKey{}, nil
	}
}


