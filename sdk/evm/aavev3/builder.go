package aavev3

import (
	"context"
	"fmt"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type TxData struct {
	To    ethcommon.Address
	Data  []byte
	Value *big.Int
}

var interestRateVariable = big.NewInt(2)

func BuildDepositTx(ctx context.Context, client *Client, asset ethcommon.Address, amount string, user ethcommon.Address) ([]TxData, error) {
	decimals, err := client.GetTokenDecimals(ctx, asset)
	if err != nil {
		return nil, fmt.Errorf("get token decimals: %w", err)
	}

	parsed, err := ParseAmount(amount, int(decimals))
	if err != nil {
		return nil, fmt.Errorf("parse amount: %w", err)
	}

	pool := client.PoolAddress()

	approveData := erc20Codec.PackApprove(pool, parsed)
	supplyData := poolCodec.PackSupply(asset, parsed, user, 0)

	return []TxData{
		{To: asset, Data: approveData, Value: big.NewInt(0)},
		{To: pool, Data: supplyData, Value: big.NewInt(0)},
	}, nil
}

func BuildWithdrawTx(ctx context.Context, client *Client, asset ethcommon.Address, amount string, user ethcommon.Address) ([]TxData, error) {
	decimals, err := client.GetTokenDecimals(ctx, asset)
	if err != nil {
		return nil, fmt.Errorf("get token decimals: %w", err)
	}

	parsed, err := ParseAmount(amount, int(decimals))
	if err != nil {
		return nil, fmt.Errorf("parse amount: %w", err)
	}

	pool := client.PoolAddress()
	withdrawData := poolCodec.PackWithdraw(asset, parsed, user)

	return []TxData{
		{To: pool, Data: withdrawData, Value: big.NewInt(0)},
	}, nil
}

func BuildBorrowTx(ctx context.Context, client *Client, asset ethcommon.Address, amount string, user ethcommon.Address) ([]TxData, error) {
	decimals, err := client.GetTokenDecimals(ctx, asset)
	if err != nil {
		return nil, fmt.Errorf("get token decimals: %w", err)
	}

	parsed, err := ParseAmount(amount, int(decimals))
	if err != nil {
		return nil, fmt.Errorf("parse amount: %w", err)
	}

	pool := client.PoolAddress()
	borrowData := poolCodec.PackBorrow(asset, parsed, interestRateVariable, 0, user)

	return []TxData{
		{To: pool, Data: borrowData, Value: big.NewInt(0)},
	}, nil
}

func BuildRepayTx(ctx context.Context, client *Client, asset ethcommon.Address, amount string, user ethcommon.Address) ([]TxData, error) {
	decimals, err := client.GetTokenDecimals(ctx, asset)
	if err != nil {
		return nil, fmt.Errorf("get token decimals: %w", err)
	}

	parsed, err := ParseAmount(amount, int(decimals))
	if err != nil {
		return nil, fmt.Errorf("parse amount: %w", err)
	}

	pool := client.PoolAddress()

	approveData := erc20Codec.PackApprove(pool, parsed)
	repayData := poolCodec.PackRepay(asset, parsed, interestRateVariable, user)

	return []TxData{
		{To: asset, Data: approveData, Value: big.NewInt(0)},
		{To: pool, Data: repayData, Value: big.NewInt(0)},
	}, nil
}
