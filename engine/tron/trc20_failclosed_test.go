package tron

// Regression test for the TRON TRC-20 fail-open fix.
//
// Before: validateTRC20Transfer iterated only the constraints present in the
// rule and never called validateTarget, so a rule with zero/incomplete
// constraints ALLOWED a TRC-20 transfer to any recipient for any amount.
// After: every security-critical field (recipient, amount, from_asset) must be
// constrained or the engine refuses (fails closed) — matching the EVM engine.

import (
	"encoding/hex"
	"strings"
	"testing"

	chaintron "github.com/vultisig/recipes/chain/tron"
	"github.com/vultisig/recipes/types"
)

// buildTronTRC20Tx hand-rolls a TriggerSmartContract (TRC-20) tx in the same
// manual-protobuf style as buildTronTransferTx in tron_test.go.
func buildTronTRC20Tx(ownerAddr, contractAddr, callData []byte) []byte {
	v := []byte{}
	v = append(v, 0x0a, byte(len(ownerAddr)))
	v = append(v, ownerAddr...)
	v = append(v, 0x12, byte(len(contractAddr)))
	v = append(v, contractAddr...)
	v = append(v, 0x22) // f4 data, wire type 2
	v = append(v, encodeVarint(uint64(len(callData)))...)
	v = append(v, callData...)

	typeURL := "type.googleapis.com/protocol.TriggerSmartContract"
	p := []byte{}
	p = append(p, 0x0a)
	p = append(p, encodeVarint(uint64(len(typeURL)))...)
	p = append(p, []byte(typeURL)...)
	p = append(p, 0x12)
	p = append(p, encodeVarint(uint64(len(v)))...)
	p = append(p, v...)

	c := []byte{}
	c = append(c, 0x08, 0x1f) // f1 type = 31 (TriggerSmartContract)
	c = append(c, 0x12)
	c = append(c, encodeVarint(uint64(len(p)))...)
	c = append(c, p...)

	rd := []byte{}
	refBlockBytes := []byte{0x12, 0x34}
	rd = append(rd, 0x0a, byte(len(refBlockBytes)))
	rd = append(rd, refBlockBytes...)
	refBlockHash := []byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11}
	rd = append(rd, 0x22, byte(len(refBlockHash)))
	rd = append(rd, refBlockHash...)
	rd = append(rd, 0x40)
	rd = append(rd, encodeVarint(1700000000000)...)
	rd = append(rd, 0x5a)
	rd = append(rd, encodeVarint(uint64(len(c)))...)
	rd = append(rd, c...)
	rd = append(rd, 0x70)
	rd = append(rd, encodeVarint(1699999990000)...)
	return rd
}

func trc20TransferCalldata(recipient20Hex, amountHex64 string) []byte {
	recipient32 := strings.Repeat("0", 64-len(recipient20Hex)) + recipient20Hex
	cd, _ := hex.DecodeString("a9059cbb" + recipient32 + amountHex64)
	return cd
}

func fixed(name, value string) *types.ParameterConstraint {
	return &types.ParameterConstraint{
		ParameterName: name,
		Constraint: &types.Constraint{
			Type:  types.ConstraintType_CONSTRAINT_TYPE_FIXED,
			Value: &types.Constraint_FixedValue{FixedValue: value},
		},
	}
}

func TestTRC20_FailsClosed_OnMissingConstraints(t *testing.T) {
	tr := NewTron()

	owner, _ := hex.DecodeString("41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	contract, _ := hex.DecodeString("41a614f803b6fd780986a42c78ec9c7f77e6ded13d")
	recipient20 := "b614f803b6fd780986a42c78ec9c7f77e6ded13d"
	amount := "0000000000000000000000000000000000000000000000000000000005f5e100" // 100,000,000
	txBytes := buildTronTRC20Tx(owner, contract, trc20TransferCalldata(recipient20, amount))

	// Resolve the exact addresses the engine will compare against.
	recipientB58, err := tr.hexToTronAddress(strings.Repeat("0", 24) + recipient20)
	if err != nil {
		t.Fatalf("recipient encode: %v", err)
	}
	parsed, err := tr.chain.ParseTransactionBytes(txBytes)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	contractB58 := parsed.(*chaintron.ParsedTronTransaction).GetContractAddress()

	cases := []struct {
		name        string
		constraints []*types.ParameterConstraint
		wantErr     string // "" = must be allowed
	}{
		{
			name:        "zero constraints — was FAIL-OPEN, now rejected",
			constraints: nil,
			wantErr:     "missing required constraint",
		},
		{
			name:        "missing recipient (amount+from_asset only) — rejected",
			constraints: []*types.ParameterConstraint{fixed("amount", "100000000"), fixed("from_asset", contractB58)},
			wantErr:     "missing required constraint",
		},
		{
			name: "fully constrained + matching — ALLOWED",
			constraints: []*types.ParameterConstraint{
				fixed("recipient", recipientB58),
				fixed("amount", "100000000"),
				fixed("from_asset", contractB58),
			},
			wantErr: "",
		},
		{
			name: "fully constrained but wrong recipient — rejected by value check (not the gate)",
			constraints: []*types.ParameterConstraint{
				fixed("recipient", "TWrongRecipientAddressXXXXXXXXXXXXXXXXX"),
				fixed("amount", "100000000"),
				fixed("from_asset", contractB58),
			},
			wantErr: "recipient mismatch",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rule := &types.Rule{
				Effect:               types.Effect_EFFECT_ALLOW,
				Resource:             "tron.trc20.transfer",
				ParameterConstraints: tc.constraints,
			}
			err := tr.Evaluate(rule, txBytes)
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("expected ALLOWED (nil), got: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("expected error containing %q, got: %v", tc.wantErr, err)
			}
		})
	}
}
