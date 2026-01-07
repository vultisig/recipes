package defi

// DeFi Protocol identifiers
type Protocol string

const (
	// LP Protocols
	ProtocolUniswapV3 Protocol = "uniswap_v3"
	ProtocolRaydium   Protocol = "raydium"
	ProtocolOrca      Protocol = "orca"
	ProtocolMeteora   Protocol = "meteora"

	// Lending Protocols
	ProtocolAave     Protocol = "aave"
	ProtocolCompound Protocol = "compound"
	ProtocolKamino   Protocol = "kamino"
	ProtocolMarginfi Protocol = "marginfi"
	ProtocolSolend   Protocol = "solend"

	// Perps Protocols
	ProtocolGMX          Protocol = "gmx"
	ProtocolLighter      Protocol = "lighter"
	ProtocolHyperliquid  Protocol = "hyperliquid"
	ProtocolJupiterPerps Protocol = "jupiter_perps"
	ProtocolDrift        Protocol = "drift"

	// Prediction Market Protocols
	ProtocolPolymarket Protocol = "polymarket"
)

// LP Action types
type LPAction string

const (
	LPActionAdd         LPAction = "add"
	LPActionRemove      LPAction = "remove"
	LPActionCollectFees LPAction = "collect_fees"
)

// Lend Action types
type LendAction string

const (
	LendActionSupply   LendAction = "supply"
	LendActionWithdraw LendAction = "withdraw"
	LendActionBorrow   LendAction = "borrow"
	LendActionRepay    LendAction = "repay"
)

// Perps Action types
type PerpsAction string

const (
	PerpsActionOpenLong     PerpsAction = "open_long"
	PerpsActionOpenShort    PerpsAction = "open_short"
	PerpsActionClose        PerpsAction = "close"
	PerpsActionAdjustMargin PerpsAction = "adjust_margin"
)

// Bet Action types
type BetAction string

const (
	BetActionBuy    BetAction = "buy"    // Buy outcome tokens (outcome is encoded in tokenId)
	BetActionSell   BetAction = "sell"   // Sell outcome tokens
	BetActionCancel BetAction = "cancel" // Cancel an open order
)

// IsEVMProtocol returns true if the protocol is EVM-based
func IsEVMProtocol(p Protocol) bool {
	switch p {
	case ProtocolUniswapV3, ProtocolAave, ProtocolCompound, ProtocolGMX, ProtocolLighter, ProtocolHyperliquid, ProtocolPolymarket:
		return true
	default:
		return false
	}
}

// IsSolanaProtocol returns true if the protocol is Solana-based
func IsSolanaProtocol(p Protocol) bool {
	switch p {
	case ProtocolRaydium, ProtocolOrca, ProtocolMeteora, ProtocolKamino, ProtocolMarginfi, ProtocolSolend, ProtocolJupiterPerps, ProtocolDrift:
		return true
	default:
		return false
	}
}

