// Package cosmos provides common types and utilities for Cosmos SDK-based blockchains.
package cosmos

import (
	"fmt"
	"math/big"
	"strings"
)

// Secp256k1 curve order for signature normalization
var SecpN, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
var SecpHalfN = new(big.Int).Rsh(new(big.Int).Set(SecpN), 1)

// NormalizeLowS normalizes S to low-S form as required by Cosmos SDK.
// This ensures signature malleability protection per BIP-62.
func NormalizeLowS(s []byte) ([]byte, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("empty s")
	}
	var sb big.Int
	sb.SetBytes(s)
	if sb.Sign() <= 0 || sb.Cmp(SecpN) >= 0 {
		return nil, fmt.Errorf("s not in [1, N-1]")
	}
	if sb.Cmp(SecpHalfN) > 0 {
		sb.Sub(SecpN, &sb)
	}
	out := sb.Bytes()
	// left-pad to 32 bytes
	if len(out) < 32 {
		pad := make([]byte, 32-len(out))
		out = append(pad, out...)
	}
	return out, nil
}

// CleanHex removes 0x prefix from hex strings.
func CleanHex(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		return s[2:]
	}
	return s
}

