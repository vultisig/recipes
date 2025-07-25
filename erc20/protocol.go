package erc20

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vultisig/recipes/sdk/evm/codegen/erc20"
	"github.com/vultisig/recipes/types"
)

type Erc20 struct {
	vultisigChainID string
	funcs           []*types.Function
	funcMap         map[string]*types.Function
}

func NewProtocol(vultisigChainID string) types.Protocol {
	fns := []*types.Function{{
		ID: "transfer",
		Parameters: []*types.FunctionParam{{
			Name: "recipient",
			Type: "address",
		}, {
			Name: "amount",
			Type: "decimal",
		}, {
			Name: "token",
			Type: "address",
		}},
	}}
	funcMap := map[string]*types.Function{}
	for _, fn := range fns {
		funcMap[fn.ID] = fn
	}

	return &Erc20{
		vultisigChainID: vultisigChainID,
		funcs:           fns,
		funcMap:         funcMap,
	}
}

func (p *Erc20) ID() string {
	return "erc20"
}

func (p *Erc20) Name() string {
	return "ERC20"
}

func (p *Erc20) ChainID() string {
	return p.vultisigChainID
}

func (p *Erc20) Description() string {
	return "ERC20 token standard"
}

func (p *Erc20) Functions() []*types.Function {
	return p.funcs
}

func (p *Erc20) GetFunction(id string) (*types.Function, error) {
	f, ok := p.funcMap[id]
	if !ok {
		return nil, fmt.Errorf("function not found: id=%s", id)
	}
	return f, nil
}

func (p *Erc20) MatchFunctionCall(
	decodedTx types.DecodedTransaction,
	policyMatcher *types.PolicyFunctionMatcher,
) (bool, map[string]interface{}, error) {
	switch policyMatcher.FunctionID {
	case "transfer":
		recipientRaw, err := findFixedArg(policyMatcher.Constraints, "recipient")
		if err != nil {
			return false, nil, fmt.Errorf("failed to get fixed arg: %w", err)
		}
		amountRaw, err := findFixedArg(policyMatcher.Constraints, "amount")
		if err != nil {
			return false, nil, fmt.Errorf("failed to get fixed arg: %w", err)
		}
		tokenRaw, err := findFixedArg(policyMatcher.Constraints, "token")
		if err != nil {
			return false, nil, fmt.Errorf("failed to get fixed arg: %w", err)
		}

		recipient := common.HexToAddress(recipientRaw)
		token := common.HexToAddress(tokenRaw)
		amount, ok := new(big.Int).SetString(amountRaw, 10)
		if !ok {
			return false, nil, fmt.Errorf("failed to create big int: %s", amountRaw)
		}

		expectedData := erc20.NewErc20().PackTransfer(recipient, amount)

		to := common.HexToAddress(decodedTx.To())
		if to.Cmp(token) != 0 {
			return false, nil, fmt.Errorf("tx 'to' invalid: (%s != %s)", to.Hex(), token.Hex())
		}
		if decodedTx.Value() != nil && decodedTx.Value().Cmp(big.NewInt(0)) != 0 {
			return false, nil, fmt.Errorf("tx 'value' must be zero: (value=%s)", decodedTx.Value().String())
		}
		if decodedTx.Data() == nil || bytes.Compare(decodedTx.Data(), expectedData) != 0 {
			return false, nil, errors.New("wrong calldata, check args")
		}

		extractedParams := map[string]interface{}{
			"recipient": strings.ToLower(recipient.Hex()), // Normalize to lowercase for policy matching
			"amount":    amount.String(),
			"chainId":   policyMatcher.ResourcePath.ChainId,
			"asset":     strings.ToLower(token.Hex()),
		}

		return true, extractedParams, nil
	default:
		return false, nil, nil
	}
}

func findFixedArg(params []*types.ParameterConstraint, name string) (string, error) {
	for _, param := range params {
		if param.ParameterName == name {
			return param.GetConstraint().GetFixedValue(), nil
		}
	}
	return "", fmt.Errorf("arg not found: %s", name)
}
