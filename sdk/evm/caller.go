package evm

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

type callableContract interface {
	Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract
}

// CallReadonly calls a read-only(view) function on a smart contract, and returns exact type.
// Use codegen structs for 'callableContract' arg, and respective func pack/unpack.
//
// Example:
// import "github.com/vultisig/recipes/sdk/evm/codegen/erc20"
// balance, err := evm.CallReadonly(ctx, rpc, erc20.NewErc20(), usdtContractAddress, erc20.PackBalanceOf(owner), erc20.UnpackBalanceOf, nil)
//
// Note: for non-view functions, use SDK.MakeTx method.
func CallReadonly[T any](
	ctx context.Context,
	rpc bind.ContractBackend,
	contract callableContract,
	address common.Address,
	data []byte,
	unpack func([]byte) (T, error),
	inOpts *bind.CallOpts,
) (T, error) {
	var z T

	opts := &bind.CallOpts{Context: ctx}
	if inOpts != nil {
		opts = inOpts
	}

	r, err := bind.Call(
		contract.Instance(rpc, address),
		opts,
		data,
		unpack,
	)
	if err != nil {
		return z, fmt.Errorf("bind.Call: %w", err)
	}
	return r, nil
}
