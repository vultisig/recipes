package evm

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	reth "github.com/vultisig/recipes/chain/evm/ethereum"
	sdk "github.com/vultisig/recipes/sdk"
	"github.com/vultisig/recipes/sdk/evm/codegen/erc20"
	"golang.org/x/sync/errgroup"
)

var ZeroAddress common.Address

type UnsignedTx []byte

// go-ethereum *ethclient.Client
type rpcClient interface {
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	FeeHistory(
		ctx context.Context,
		blockCount uint64,
		lastBlock *big.Int,
		rewardPercentiles []float64,
	) (*ethereum.FeeHistory, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
}

// go-ethereum *rpc.Client
type rpcClientRaw interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
}

type SDK struct {
	chainID      *big.Int
	rpcClient    rpcClient
	rpcClientRaw rpcClientRaw
}

func NewSDK(chainID *big.Int, rpcClient rpcClient, rpcClientRaw rpcClientRaw) *SDK {
	return &SDK{
		chainID:      chainID,
		rpcClient:    rpcClient,
		rpcClientRaw: rpcClientRaw,
	}
}

// MakeAnyTransfer if asset is not native, it will use the ERC20 to transfer the asset
func (sdk *SDK) MakeAnyTransfer(
	ctx context.Context,
	from, to, asset common.Address,
	amount *big.Int,
	nonceOffset uint64,
) (UnsignedTx, error) {
	if asset == ZeroAddress {
		tx, err := sdk.MakeTxTransferNative(ctx, from, to, amount, nonceOffset)
		if err != nil {
			return nil, fmt.Errorf("sdk.MakeTxTransferNative: %w", err)
		}
		return tx, nil
	}

	tx, err := sdk.MakeTxTransferERC20(ctx, from, to, asset, amount, nonceOffset)
	if err != nil {
		return nil, fmt.Errorf("sdk.MakeTxTransferERC20: %w", err)
	}
	return tx, nil
}

func (sdk *SDK) MakeTxTransferNative(
	ctx context.Context,
	from, to common.Address,
	value *big.Int,
	nonceOffset uint64,
) (UnsignedTx, error) {
	tx, err := sdk.MakeTx(
		ctx,
		from,
		to,
		value,
		nil,
		nonceOffset,
	)
	if err != nil {
		return nil, fmt.Errorf("sdk.MakeTx: %w", err)
	}
	return tx, nil
}

func (sdk *SDK) MakeTxTransferERC20(
	ctx context.Context,
	from, to, contractAddress common.Address,
	amount *big.Int,
	nonceOffset uint64,
) (UnsignedTx, error) {
	tx, err := sdk.MakeTx(
		ctx,
		from,
		contractAddress,
		big.NewInt(0),
		erc20.NewErc20().PackTransfer(to, amount),
		nonceOffset,
	)
	if err != nil {
		return nil, fmt.Errorf("sdk.MakeTx: %w", err)
	}
	return tx, nil
}

func (sdk *SDK) MakeTx(
	ctx context.Context,
	from, to common.Address,
	value *big.Int,
	data []byte,
	nonceOffset uint64,
) (UnsignedTx, error) {
	nonce, gasLimit, gasTipCap, maxFeePerGas, accessList, err := sdk.estimateTx(ctx, from, to, value, data, nonceOffset, 0)
	if err != nil {
		return nil, fmt.Errorf("sdk.estimateTx: %w", err)
	}

	tx, err := sdk.encodeDynamicFeeTx(
		nonce,
		to,
		gasTipCap,
		maxFeePerGas,
		gasLimit,
		value,
		data,
		accessList,
	)
	if err != nil {
		return nil, fmt.Errorf("sdk.encodeDynamicFeeTx: %w", err)
	}
	return tx, nil
}

func (sdk *SDK) MakeTxWithGasLimit(
	ctx context.Context,
	from, to common.Address,
	value *big.Int,
	data []byte,
	nonceOffset uint64,
	gasLimit uint64,
) (UnsignedTx, error) {
	nonce, _, gasTipCap, maxFeePerGas, accessList, err := sdk.estimateTx(ctx, from, to, value, data, nonceOffset, gasLimit)
	if err != nil {
		return nil, fmt.Errorf("sdk.estimateTx: %w", err)
	}

	tx, err := sdk.encodeDynamicFeeTx(
		nonce,
		to,
		gasTipCap,
		maxFeePerGas,
		gasLimit,
		value,
		data,
		accessList,
	)
	if err != nil {
		return nil, fmt.Errorf("sdk.encodeDynamicFeeTx: %w", err)
	}
	return tx, nil
}

func (sdk *SDK) Send(ctx context.Context, inTx UnsignedTx, r, s, v []byte) (*types.Transaction, error) {
	outTx, err := sdk.appendSignature(inTx, r, s, v)
	if err != nil {
		return nil, fmt.Errorf("sdk.appendSignature: %w", err)
	}

	err = sdk.broadcast(ctx, outTx)
	if err != nil {
		return nil, fmt.Errorf("sdk.broadcast: %w", err)
	}

	return outTx, nil
}

func (sdk *SDK) broadcast(ctx context.Context, tx *types.Transaction) error {
	err := sdk.rpcClient.SendTransaction(ctx, tx)
	if err != nil {
		return fmt.Errorf("sdk.rpcClient.SendTransaction: %w", err)
	}
	return nil
}

// appendSignature : v is the 'RecoveryID'
func (sdk *SDK) appendSignature(inTx UnsignedTx, r, s, v []byte) (*types.Transaction, error) {
	var sig []byte
	sig = append(sig, r...)
	sig = append(sig, s...)
	sig = append(sig, v...)

	inTxDecoded, err := reth.DecodeUnsignedPayload(inTx)
	if err != nil {
		return nil, fmt.Errorf("reth.DecodeUnsignedPayload: %w", err)
	}

	outTx, err := types.NewTx(inTxDecoded).WithSignature(types.LatestSignerForChainID(sdk.chainID), sig)
	if err != nil {
		return nil, fmt.Errorf("types.NewTx.WithSignature: %w", err)
	}

	return outTx, nil
}

type createAccessListArgs struct {
	From                 string `json:"from,omitempty"`
	To                   string `json:"to,omitempty"`
	Gas                  string `json:"gas,omitempty"`
	GasPrice             string `json:"gasPrice,omitempty"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
	Value                string `json:"value,omitempty"`
	Data                 string `json:"data,omitempty"`
}

type createAccessListRes struct {
	AccessList types.AccessList `json:"accessList"`
	GasUsed    string           `json:"gasUsed"`
}

func (sdk *SDK) estimateTx(
	ctx context.Context,
	from, to common.Address,
	value *big.Int,
	data []byte,
	nonceOffset uint64,
	fixedGasLimit uint64,
) (uint64, uint64, *big.Int, *big.Int, types.AccessList, error) {
	var eg errgroup.Group
	var gasLimit uint64
	if fixedGasLimit > 0 {
		gasLimit = fixedGasLimit
	} else {
		eg.Go(func() error {
			r, e := sdk.rpcClient.EstimateGas(ctx, ethereum.CallMsg{
				From:  from,
				To:    &to,
				Data:  data,
				Value: value,
			})
			if e != nil {
				return fmt.Errorf("sdk.rpcClient.EstimateGas: %v", e)
			}
			gasLimit = r + r/2
			return nil
		})
	}

	var gasTipCap *big.Int
	eg.Go(func() error {
		r, e := sdk.rpcClient.SuggestGasTipCap(ctx)
		if e != nil {
			return fmt.Errorf("sdk.rpcClient.SuggestGasTipCap: %v", e)
		}
		gasTipCap = addGas(r, 2)
		return nil
	})

	var baseFee *big.Int
	eg.Go(func() error {
		feeHistory, e := sdk.rpcClient.FeeHistory(ctx, 1, nil, nil)
		if e != nil {
			return fmt.Errorf("sdk.rpcClient.FeeHistory: %v", e)
		}
		if len(feeHistory.BaseFee) == 0 {
			return fmt.Errorf("feeHistory.BaseFee is empty")
		}
		baseFee = addGas(feeHistory.BaseFee[0], 2)
		return nil
	})

	var nonce uint64
	eg.Go(func() error {
		r, e := sdk.rpcClient.PendingNonceAt(ctx, from)
		if e != nil {
			return fmt.Errorf("sdk.rpcClient.PendingNonceAt: %v", e)
		}
		nonce = r
		return nil
	})
	err := eg.Wait()
	if err != nil {
		return 0, 0, nil, nil, nil, fmt.Errorf("eg.Wait: %v", err)
	}

	maxFeePerGas := new(big.Int).Add(gasTipCap, baseFee)

	gasTipCapHex := "0x0"
	if gasTipCap != nil {
		gasTipCapHex = hexutil.EncodeBig(gasTipCap)
	}

	maxFeePerGasHex := "0x0"
	if maxFeePerGas != nil {
		maxFeePerGasHex = hexutil.EncodeBig(maxFeePerGas)
	}

	valueHex := "0x0"
	if value != nil {
		valueHex = hexutil.EncodeBig(value)
	}

	var dataHex string // omitempty in JSON
	if data != nil {
		dataHex = hexutil.Encode(data)
	}

	gas := hexutil.EncodeUint64(gasLimit)

	var callRes createAccessListRes
	_ = sdk.rpcClientRaw.CallContext(
		ctx,
		&callRes,
		"eth_createAccessList",
		createAccessListArgs{
			From:                 from.Hex(),
			To:                   to.Hex(),
			Gas:                  gas,
			MaxPriorityFeePerGas: gasTipCapHex,
			MaxFeePerGas:         maxFeePerGasHex,
			Value:                valueHex,
			Data:                 dataHex,
		},
		"latest",
	)
	return nonce + nonceOffset, gasLimit, gasTipCap, maxFeePerGas, callRes.AccessList, nil
}

func (sdk *SDK) encodeDynamicFeeTx(
	nonce uint64,
	to common.Address,
	maxPriorityFeePerGas, maxFeePerGas *big.Int,
	gas uint64,
	value *big.Int,
	data []byte,
	accessList types.AccessList,
) (UnsignedTx, error) {
	bytes, err := rlp.EncodeToBytes(reth.DynamicFeeTxWithoutSignature{
		ChainID:    sdk.chainID,
		Nonce:      nonce,
		GasTipCap:  maxPriorityFeePerGas,
		GasFeeCap:  maxFeePerGas,
		Gas:        gas,
		To:         &to,
		Value:      value,
		Data:       data,
		AccessList: accessList,
	})
	if err != nil {
		return nil, fmt.Errorf("rlp.EncodeToBytes: %v", err)
	}

	res := append([]byte{types.DynamicFeeTxType}, bytes...)
	return res, nil
}

// addGas : in + in/v
func addGas(in *big.Int, v uint64) *big.Int {
	return new(big.Int).Add(in, new(big.Int).Div(in, new(big.Int).SetUint64(v)))
}

// DeriveSigningHashes derives the signing hash from unsigned EVM transaction bytes.
// For EVM, this computes the EIP-155 signing hash and then SHA256 for the lookup key.
// Returns a single DerivedHash since EVM transactions have one signature.
func (s *SDK) DeriveSigningHashes(txBytes []byte, _ sdk.DeriveOptions) ([]sdk.DerivedHash, error) {
	// Decode the unsigned transaction
	txData, err := reth.DecodeUnsignedPayload(txBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode unsigned payload: %w", err)
	}

	// Compute the signing hash using the chain-specific signer
	tx := types.NewTx(txData)
	signer := types.LatestSignerForChainID(s.chainID)
	txHashToSign := signer.Hash(tx)

	// SHA256 of the signing hash for the lookup key
	msgHash := sha256.Sum256(txHashToSign.Bytes())

	return []sdk.DerivedHash{
		{
			Message: txHashToSign.Bytes(),
			Hash:    msgHash[:],
		},
	}, nil
}
