package swap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// EVMChainConfig contains configuration for an EVM chain
type EVMChainConfig struct {
	Name    string   // Human-readable name
	ChainID *big.Int // EIP-155 chain ID
	RPCURLs []string // List of RPC URLs (first one is primary)
}

// Default EVM chain configurations with public RPC endpoints
var evmChainConfigs = map[string]EVMChainConfig{
	"Ethereum": {
		Name:    "Ethereum",
		ChainID: big.NewInt(1),
		RPCURLs: []string{
			"https://eth.llamarpc.com",
			"https://ethereum.publicnode.com",
			"https://rpc.ankr.com/eth",
		},
	},
	"BSC": {
		Name:    "BSC",
		ChainID: big.NewInt(56),
		RPCURLs: []string{
			"https://bsc-dataseed.binance.org",
			"https://bsc.publicnode.com",
			"https://rpc.ankr.com/bsc",
		},
	},
	"Polygon": {
		Name:    "Polygon",
		ChainID: big.NewInt(137),
		RPCURLs: []string{
			"https://polygon.llamarpc.com",
			"https://polygon-bor.publicnode.com",
			"https://rpc.ankr.com/polygon",
		},
	},
	"Arbitrum": {
		Name:    "Arbitrum",
		ChainID: big.NewInt(42161),
		RPCURLs: []string{
			"https://arbitrum.llamarpc.com",
			"https://arbitrum-one.publicnode.com",
			"https://rpc.ankr.com/arbitrum",
		},
	},
	"Avalanche": {
		Name:    "Avalanche",
		ChainID: big.NewInt(43114),
		RPCURLs: []string{
			"https://api.avax.network/ext/bc/C/rpc",
			"https://avalanche-c-chain.publicnode.com",
			"https://rpc.ankr.com/avalanche",
		},
	},
	"Base": {
		Name:    "Base",
		ChainID: big.NewInt(8453),
		RPCURLs: []string{
			"https://base.llamarpc.com",
			"https://base.publicnode.com",
			"https://mainnet.base.org",
		},
	},
	"Optimism": {
		Name:    "Optimism",
		ChainID: big.NewInt(10),
		RPCURLs: []string{
			"https://optimism.llamarpc.com",
			"https://optimism.publicnode.com",
			"https://mainnet.optimism.io",
		},
	},
	"Blast": {
		Name:    "Blast",
		ChainID: big.NewInt(81457),
		RPCURLs: []string{
			"https://rpc.blast.io",
			"https://blast.din.dev/rpc",
		},
	},
	"CronosChain": {
		Name:    "CronosChain",
		ChainID: big.NewInt(25),
		RPCURLs: []string{
			"https://evm.cronos.org",
			"https://cronos.publicnode.com",
		},
	},
	"ZkSync": {
		Name:    "ZkSync",
		ChainID: big.NewInt(324),
		RPCURLs: []string{
			"https://mainnet.era.zksync.io",
			"https://zksync.drpc.org",
		},
	},
}

// GetEVMChainConfig returns the configuration for an EVM chain
func GetEVMChainConfig(chain string) (*EVMChainConfig, error) {
	config, ok := evmChainConfigs[chain]
	if !ok {
		return nil, fmt.Errorf("unknown EVM chain: %s", chain)
	}
	return &config, nil
}

// GetEVMChainID returns the EIP-155 chain ID for an EVM chain
func GetEVMChainID(chain string) (*big.Int, error) {
	config, err := GetEVMChainConfig(chain)
	if err != nil {
		return nil, err
	}
	return config.ChainID, nil
}

// NonceClient handles EVM nonce fetching via JSON-RPC
type NonceClient struct {
	httpClient *http.Client
	rpcURLs    map[string][]string // chain name -> RPC URLs
}

// NewNonceClient creates a new nonce client with default RPC URLs
func NewNonceClient() *NonceClient {
	rpcURLs := make(map[string][]string)
	for name, config := range evmChainConfigs {
		rpcURLs[name] = config.RPCURLs
	}

	return &NonceClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		rpcURLs: rpcURLs,
	}
}

// SetRPCURL sets custom RPC URLs for a chain
func (c *NonceClient) SetRPCURL(chain string, urls []string) {
	c.rpcURLs[chain] = urls
}

// jsonRPCRequest represents a JSON-RPC request
type jsonRPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

// jsonRPCResponse represents a JSON-RPC response
type jsonRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *jsonRPCError   `json:"error"`
}

// jsonRPCError represents a JSON-RPC error
type jsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetNonce fetches the pending nonce for an address on an EVM chain
func (c *NonceClient) GetNonce(ctx context.Context, chain, address string) (uint64, error) {
	urls, ok := c.rpcURLs[chain]
	if !ok || len(urls) == 0 {
		return 0, fmt.Errorf("no RPC URL configured for chain: %s", chain)
	}

	// Normalize address
	if !strings.HasPrefix(address, "0x") {
		address = "0x" + address
	}

	var lastErr error
	for _, url := range urls {
		nonce, err := c.getNonceFromRPC(ctx, url, address)
		if err == nil {
			return nonce, nil
		}
		lastErr = err
	}

	return 0, fmt.Errorf("failed to get nonce from all RPCs for %s: %w", chain, lastErr)
}

// getNonceFromRPC fetches the nonce from a specific RPC endpoint
func (c *NonceClient) getNonceFromRPC(ctx context.Context, rpcURL, address string) (uint64, error) {
	req := jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getTransactionCount",
		Params:  []interface{}{address, "pending"},
		ID:      1,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", rpcURL, bytes.NewReader(reqBody))
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response: %w", err)
	}

	var rpcResp jsonRPCResponse
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if rpcResp.Error != nil {
		return 0, fmt.Errorf("RPC error: %s (code: %d)", rpcResp.Error.Message, rpcResp.Error.Code)
	}

	// Parse hex nonce
	var nonceHex string
	if err := json.Unmarshal(rpcResp.Result, &nonceHex); err != nil {
		return 0, fmt.Errorf("failed to unmarshal nonce: %w", err)
	}

	nonce, err := parseHexUint64(nonceHex)
	if err != nil {
		return 0, fmt.Errorf("failed to parse nonce: %w", err)
	}

	return nonce, nil
}

// GetGasPrice fetches the current gas price for an EVM chain
func (c *NonceClient) GetGasPrice(ctx context.Context, chain string) (*big.Int, error) {
	urls, ok := c.rpcURLs[chain]
	if !ok || len(urls) == 0 {
		return nil, fmt.Errorf("no RPC URL configured for chain: %s", chain)
	}

	var lastErr error
	for _, url := range urls {
		gasPrice, err := c.getGasPriceFromRPC(ctx, url)
		if err == nil {
			return gasPrice, nil
		}
		lastErr = err
	}

	return nil, fmt.Errorf("failed to get gas price from all RPCs for %s: %w", chain, lastErr)
}

// getGasPriceFromRPC fetches the gas price from a specific RPC endpoint
func (c *NonceClient) getGasPriceFromRPC(ctx context.Context, rpcURL string) (*big.Int, error) {
	req := jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_gasPrice",
		Params:  []interface{}{},
		ID:      1,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", rpcURL, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var rpcResp jsonRPCResponse
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if rpcResp.Error != nil {
		return nil, fmt.Errorf("RPC error: %s (code: %d)", rpcResp.Error.Message, rpcResp.Error.Code)
	}

	var gasPriceHex string
	if err := json.Unmarshal(rpcResp.Result, &gasPriceHex); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gas price: %w", err)
	}

	gasPrice, ok := parseHexBigInt(gasPriceHex)
	if !ok {
		return nil, fmt.Errorf("failed to parse gas price: %s", gasPriceHex)
	}

	return gasPrice, nil
}

// parseHexUint64 parses a hex string (with 0x prefix) to uint64
func parseHexUint64(hex string) (uint64, error) {
	hex = strings.TrimPrefix(hex, "0x")
	return strconv.ParseUint(hex, 16, 64)
}

// parseHexBigInt parses a hex string (with 0x prefix) to *big.Int
func parseHexBigInt(hex string) (*big.Int, bool) {
	hex = strings.TrimPrefix(hex, "0x")
	n := new(big.Int)
	_, ok := n.SetString(hex, 16)
	return n, ok
}

// DefaultNonceClient is the default nonce client instance
var DefaultNonceClient = NewNonceClient()

// GetNonce fetches the pending nonce using the default client
func GetNonce(ctx context.Context, chain, address string) (uint64, error) {
	return DefaultNonceClient.GetNonce(ctx, chain, address)
}

// GetGasPrice fetches the current gas price using the default client
func GetGasPrice(ctx context.Context, chain string) (*big.Int, error) {
	return DefaultNonceClient.GetGasPrice(ctx, chain)
}

