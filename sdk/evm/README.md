# EVM Package

The EVM package provides a Go SDK for interacting with EVM compatible blockchains. It simplifies common operations like creating and sending transactions, transferring tokens, and interacting with smart contracts.

## Components

### SDK

The main `SDK` struct provides methods for creating and sending transactions:

```go
import (
    "github.com/vultisig/recipes/sdk/evm"
)

sdk := evm.NewSDK(chainID, rpcClient, rpcClientRaw)

// Transfer native tokens
txNative, err := sdk.MakeTxTransferNative(ctx, fromAddress, toAddress, amount)
resultTxNative, err := sdk.Send(ctx, tx, r, s, v)

// Transfer ERC20 tokens
txERC20, err := sdk.MakeTxTransferERC20(ctx, fromAddress, toAddress, tokenContractAddress, amount)
resultTxERC20, err := sdk.Send(ctx, tx, r, s, v)

// Any transaction with calldata (e.g. ERC20 approve)
tx, err := sdk.MakeTx(ctx, fromAddress, contractAddress, big.NewInt(0), erc20.NewErc20().PackApprove(spender, amount))
tx, err := sdk.Send(ctx, tx, r, s, v)
```

### Caller

The `CallReadonly` function allows you to make read-only(view) calls to smart contracts, without worrying about correct mapping from Go types to Solidity types back and forth:

```go
import (
    "github.com/vultisig/recipes/sdk/evm"
    "github.com/vultisig/recipes/sdk/evm/codegen/erc20"
)

// Get ERC20 token balance
balance, err := evm.CallReadonly(
    ctx,
    rpcClient,
    erc20.NewErc20(),
    contractAddress,
    erc20.NewErc20().PackBalanceOf(ownerAddress),
    erc20.NewErc20().UnpackBalanceOf,
    nil,
)
```

### Codegen Packages

The `codegen` directory contains auto-generated Go bindings for common smart contracts:

- `erc20`: Bindings for the ERC20 token standard;
- `uniswapv2_router`: Bindings for Uniswap V2 Router contract;

These bindings provide type-safe methods for interacting with smart contracts.

#### How to add new bindings

Make sure `abigen` installed locally:

```
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

Make sure your `go/bin` directory is in your `PATH` environment variable, so you can run `abigen` command from anywhere.

Put ABI JSON (you can get it in desired protocol docs or directly in etherscan) in `abi` directory at the project root, then run the code generation tool:

```bash
go run cmd/gen_abi_bind/main.go
```

It will read all ABI files from the `abi` directory and generate Go bindings in the `sdk/evm/codegen/` directory.

For each ABI JSON file it would generate own Go package. Why not group uniswap contracts for example in one package called `uniswapv2` with filenames convention like `router.go`, `factory.go`? It is much safer from human mistake to import Go package with the same name as the contract name, to avoid packing calldata to one contract with args from another.
