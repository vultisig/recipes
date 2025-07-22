package thorchain

import (
	"fmt"
	"strings"
)

type MemoType string

const (
	MemoTypeSwap MemoType = "SWAP"
	// TODO: Implement additional memo types in future versions
	// MemoTypeAdd      MemoType = "ADD"
	// MemoTypeWithdraw MemoType = "WITHDRAW"
	// MemoTypeDonate   MemoType = "DONATE"
	MemoTypeTransfer MemoType = "TRANSFER"
	MemoTypeUnknown  MemoType = "UNKNOWN"
)

// ParsedMemo represents a parsed Thorchain memo
type ParsedMemo struct {
	Type        MemoType
	Asset       string
	DestAddr    string
	Limit       string
	Affiliate   string
	Fee         string
	Pool        string
	BasisPoints string
	AssetAddr   string
}

// ParseThorchainMemo parses a Thorchain memo string into structured data
func ParseThorchainMemo(memo string) *ParsedMemo {
	parsed := &ParsedMemo{
		Type: MemoTypeTransfer,
	}

	// Handle empty or simple memos
	if memo == "" || !strings.Contains(memo, ":") {
		return parsed
	}

	// Split memo by colons
	parts := strings.Split(strings.ToUpper(memo), ":")
	if len(parts) < 2 {
		parsed.Type = MemoTypeUnknown
		return parsed
	}

	// Parse based on operation type
	operation := parts[0]
	switch operation {
	case "SWAP":
		parsed.Type = MemoTypeSwap
		parsed.parseSwapMemo(parts)
	default:
		parsed.Type = MemoTypeUnknown
	}

	return parsed
}

// parseSwapMemo parses SWAP:ASSET:DEST_ADDR:LIMIT:AFFILIATE:FEE
func (p *ParsedMemo) parseSwapMemo(parts []string) {
	if len(parts) >= 2 {
		p.Asset = parts[1]
	}
	if len(parts) >= 3 {
		p.DestAddr = parts[2]
	}
	if len(parts) >= 4 {
		p.Limit = parts[3]
	}
	if len(parts) >= 5 {
		p.Affiliate = parts[4]
	}
	if len(parts) >= 6 {
		p.Fee = parts[5]
	}
}

// IsSwap returns true if this is a swap operation
func (p *ParsedMemo) IsSwap() bool {
	return p.Type == MemoTypeSwap
}

// IsTransfer returns true if this is a simple transfer (no structured memo)
func (p *ParsedMemo) IsTransfer() bool {
	return p.Type == MemoTypeTransfer
}

// Validate performs basic validation on the parsed memo
func (p *ParsedMemo) Validate() error {
	if p.Type == MemoTypeSwap && p.Asset == "" {
		return fmt.Errorf("swap memo missing target asset")
	}

	return nil
}
