package aavev3

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"

	goeth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/vultisig/recipes/sdk/evm/codegen/aavev3_dataprovider"
	"github.com/vultisig/recipes/sdk/evm/codegen/aavev3_pool"
	"github.com/vultisig/recipes/sdk/evm/codegen/erc20"
)

var (
	poolCodec = aavev3_pool.NewAavev3Pool()
	dpCodec   = aavev3_dataprovider.NewAavev3Dataprovider()
	erc20Codec = erc20.NewErc20()
)

type RPCCaller interface {
	CallContract(ctx context.Context, call goeth.CallMsg, blockNumber *big.Int) ([]byte, error)
}

type Client struct {
	rpc    RPCCaller
	deploy Deployment
}

func NewClient(rpc RPCCaller, deploy Deployment) *Client {
	return &Client{rpc: rpc, deploy: deploy}
}

func (c *Client) GetUserAccountData(ctx context.Context, user ethcommon.Address) (*UserAccountData, error) {
	pool := c.deploy.Pool
	data, err := c.rpc.CallContract(ctx, goeth.CallMsg{
		To:   &pool,
		Data: poolCodec.PackGetUserAccountData(user),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("call getUserAccountData: %w", err)
	}
	out, err := poolCodec.UnpackGetUserAccountData(data)
	if err != nil {
		return nil, fmt.Errorf("unpack getUserAccountData: %w", err)
	}
	return userAccountDataFromCodegen(out), nil
}

func (c *Client) GetReserveData(ctx context.Context, asset ethcommon.Address) (*ReserveData, error) {
	pool := c.deploy.Pool
	data, err := c.rpc.CallContract(ctx, goeth.CallMsg{
		To:   &pool,
		Data: poolCodec.PackGetReserveData(asset),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("call getReserveData: %w", err)
	}
	out, err := poolCodec.UnpackGetReserveData(data)
	if err != nil {
		return nil, fmt.Errorf("unpack getReserveData: %w", err)
	}
	return reserveDataFromCodegen(out), nil
}

func (c *Client) GetReserveConfigData(ctx context.Context, asset ethcommon.Address) (*ReserveConfigData, error) {
	dp := c.deploy.DataProvider
	data, err := c.rpc.CallContract(ctx, goeth.CallMsg{
		To:   &dp,
		Data: dpCodec.PackGetReserveConfigurationData(asset),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("call getReserveConfigurationData: %w", err)
	}
	out, err := dpCodec.UnpackGetReserveConfigurationData(data)
	if err != nil {
		return nil, fmt.Errorf("unpack getReserveConfigurationData: %w", err)
	}
	return reserveConfigDataFromCodegen(out), nil
}

func (c *Client) GetTokenDecimals(ctx context.Context, token ethcommon.Address) (uint8, error) {
	data, err := c.rpc.CallContract(ctx, goeth.CallMsg{
		To:   &token,
		Data: erc20Codec.PackDecimals(),
	}, nil)
	if err != nil {
		return 0, fmt.Errorf("call decimals(): %w", err)
	}
	dec, err := erc20Codec.UnpackDecimals(data)
	if err != nil {
		return 0, fmt.Errorf("unpack decimals: %w", err)
	}
	return dec, nil
}

func (c *Client) GetTokenSymbol(ctx context.Context, token ethcommon.Address) (string, error) {
	data, err := c.rpc.CallContract(ctx, goeth.CallMsg{
		To:   &token,
		Data: erc20Codec.PackSymbol(),
	}, nil)
	if err != nil {
		return "", fmt.Errorf("call symbol(): %w", err)
	}
	sym, unpackErr := erc20Codec.UnpackSymbol(data)
	if unpackErr == nil {
		return sym, nil
	}
	return decodeABIString(data)
}

func (c *Client) PoolAddress() ethcommon.Address {
	return c.deploy.Pool
}

func decodeABIString(data []byte) (string, error) {
	if len(data) < 32 {
		return "", fmt.Errorf("data too short: %d bytes", len(data))
	}

	offset := new(big.Int).SetBytes(data[:32])
	if offset.Cmp(big.NewInt(int64(len(data)))) < 0 && offset.Int64() >= 32 {
		off := int(offset.Int64())
		if off+32 <= len(data) {
			length := binary.BigEndian.Uint64(data[off+24 : off+32])
			if off+32+int(length) <= len(data) {
				return string(data[off+32 : off+32+int(length)]), nil
			}
		}
	}

	s := strings.TrimRight(string(data[:32]), "\x00")
	if isPrintable(s) && len(s) > 0 {
		return s, nil
	}

	return "", fmt.Errorf("unable to decode string from ABI data")
}

func isPrintable(s string) bool {
	for _, r := range s {
		if r < 0x20 || r > 0x7e {
			return false
		}
	}
	return true
}
