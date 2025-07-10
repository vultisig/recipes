# Cryptocurrency Wallet Policy Resources

This document lists all available resources that can be used in your wallet plugin policies.
Each resource represents an action that can be performed by a plugin, subject to policy constraints.

## Available Resources


### arbitrum.eth.transfer

**Chain:** Arbitrum One  
**Protocol:** Arbitrum One  
**Function:** Transfer Arbitrum One ETH  

Transfer Ether to another address on Arbitrum One

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| recipient | address | The Arbitrum One address of the recipient |
| amount | decimal | The amount of Ether to transfer |


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.eth.transfer",
  "effect": "ALLOW",
  "constraints": {
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### arbitrum.usdc.allowance

**Chain:** Arbitrum One  
**Protocol:** USD Coin  
**Function:** USD Coin allowance  

Get the amount of USD Coin tokens allowed to be spent by an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| owner | address | owner parameter of type address |
| spender | address | spender parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdc.allowance",
  "effect": "ALLOW",
  "constraints": {
    "owner": {
      "type": "fixed",
      "value": "example_value"
    },
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### arbitrum.usdc.approve

**Chain:** Arbitrum One  
**Protocol:** USD Coin  
**Function:** USD Coin approve  

Approve an address to spend USD Coin tokens

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| spender | address | spender parameter of type address |
| amount | decimal | The amount of USD Coin tokens |


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdc.approve",
  "effect": "ALLOW",
  "constraints": {
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### arbitrum.usdc.balanceOf

**Chain:** Arbitrum One  
**Protocol:** USD Coin  
**Function:** USD Coin balanceOf  

Get the USD Coin token balance of an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| account | address | account parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdc.balanceOf",
  "effect": "ALLOW",
  "constraints": {
    "account": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### arbitrum.usdc.decimals

**Chain:** Arbitrum One  
**Protocol:** USD Coin  
**Function:** USD Coin decimals  

Call the decimals function on USD Coin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdc.decimals",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### arbitrum.usdc.name

**Chain:** Arbitrum One  
**Protocol:** USD Coin  
**Function:** USD Coin name  

Call the name function on USD Coin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdc.name",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### arbitrum.usdc.symbol

**Chain:** Arbitrum One  
**Protocol:** USD Coin  
**Function:** USD Coin symbol  

Call the symbol function on USD Coin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdc.symbol",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### arbitrum.usdc.totalSupply

**Chain:** Arbitrum One  
**Protocol:** USD Coin  
**Function:** USD Coin totalSupply  

Call the totalSupply function on USD Coin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdc.totalSupply",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### arbitrum.usdc.transfer

**Chain:** Arbitrum One  
**Protocol:** USD Coin  
**Function:** USD Coin transfer  

Transfer USD Coin tokens to another address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of USD Coin tokens |


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdc.transfer",
  "effect": "ALLOW",
  "constraints": {
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### arbitrum.usdc.transferFrom

**Chain:** Arbitrum One  
**Protocol:** USD Coin  
**Function:** USD Coin transferFrom  

Transfer USD Coin tokens from one address to another

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| sender | address | sender parameter of type address |
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of USD Coin tokens |


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdc.transferFrom",
  "effect": "ALLOW",
  "constraints": {
    "sender": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### arbitrum.usdt.allowance

**Chain:** Arbitrum One  
**Protocol:** Tether USD  
**Function:** Tether USD allowance  

Get the amount of Tether USD tokens allowed to be spent by an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| owner | address | owner parameter of type address |
| spender | address | spender parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdt.allowance",
  "effect": "ALLOW",
  "constraints": {
    "owner": {
      "type": "fixed",
      "value": "example_value"
    },
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### arbitrum.usdt.approve

**Chain:** Arbitrum One  
**Protocol:** Tether USD  
**Function:** Tether USD approve  

Approve an address to spend Tether USD tokens

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| spender | address | spender parameter of type address |
| amount | decimal | The amount of Tether USD tokens |


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdt.approve",
  "effect": "ALLOW",
  "constraints": {
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### arbitrum.usdt.balanceOf

**Chain:** Arbitrum One  
**Protocol:** Tether USD  
**Function:** Tether USD balanceOf  

Get the Tether USD token balance of an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| account | address | account parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdt.balanceOf",
  "effect": "ALLOW",
  "constraints": {
    "account": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### arbitrum.usdt.decimals

**Chain:** Arbitrum One  
**Protocol:** Tether USD  
**Function:** Tether USD decimals  

Call the decimals function on Tether USD token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdt.decimals",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### arbitrum.usdt.name

**Chain:** Arbitrum One  
**Protocol:** Tether USD  
**Function:** Tether USD name  

Call the name function on Tether USD token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdt.name",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### arbitrum.usdt.symbol

**Chain:** Arbitrum One  
**Protocol:** Tether USD  
**Function:** Tether USD symbol  

Call the symbol function on Tether USD token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdt.symbol",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### arbitrum.usdt.totalSupply

**Chain:** Arbitrum One  
**Protocol:** Tether USD  
**Function:** Tether USD totalSupply  

Call the totalSupply function on Tether USD token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdt.totalSupply",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### arbitrum.usdt.transfer

**Chain:** Arbitrum One  
**Protocol:** Tether USD  
**Function:** Tether USD transfer  

Transfer Tether USD tokens to another address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of Tether USD tokens |


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdt.transfer",
  "effect": "ALLOW",
  "constraints": {
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### arbitrum.usdt.transferFrom

**Chain:** Arbitrum One  
**Protocol:** Tether USD  
**Function:** Tether USD transferFrom  

Transfer Tether USD tokens from one address to another

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| sender | address | sender parameter of type address |
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of Tether USD tokens |


**Example Policy Rule:**

```json
{
  "resource": "arbitrum.usdt.transferFrom",
  "effect": "ALLOW",
  "constraints": {
    "sender": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### bitcoin.btc.transfer

**Chain:** Bitcoin  
**Protocol:** Bitcoin  
**Function:** Transfer BTC  

Transfer Bitcoin to another address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| recipient | address | The Bitcoin address of the recipient |
| amount | decimal | The amount of Bitcoin to transfer |


**Example Policy Rule:**

```json
{
  "resource": "bitcoin.btc.transfer",
  "effect": "ALLOW",
  "constraints": {
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.dai.allowance

**Chain:** Ethereum  
**Protocol:** Dai Stablecoin  
**Function:** Dai Stablecoin allowance  

Get the amount of Dai Stablecoin tokens allowed to be spent by an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| owner | address | owner parameter of type address |
| spender | address | spender parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.dai.allowance",
  "effect": "ALLOW",
  "constraints": {
    "owner": {
      "type": "fixed",
      "value": "example_value"
    },
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.dai.approve

**Chain:** Ethereum  
**Protocol:** Dai Stablecoin  
**Function:** Dai Stablecoin approve  

Approve an address to spend Dai Stablecoin tokens

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| spender | address | spender parameter of type address |
| amount | decimal | The amount of Dai Stablecoin tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.dai.approve",
  "effect": "ALLOW",
  "constraints": {
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.dai.balanceOf

**Chain:** Ethereum  
**Protocol:** Dai Stablecoin  
**Function:** Dai Stablecoin balanceOf  

Get the Dai Stablecoin token balance of an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| account | address | account parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.dai.balanceOf",
  "effect": "ALLOW",
  "constraints": {
    "account": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.dai.decimals

**Chain:** Ethereum  
**Protocol:** Dai Stablecoin  
**Function:** Dai Stablecoin decimals  

Call the decimals function on Dai Stablecoin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.dai.decimals",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.dai.name

**Chain:** Ethereum  
**Protocol:** Dai Stablecoin  
**Function:** Dai Stablecoin name  

Call the name function on Dai Stablecoin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.dai.name",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.dai.symbol

**Chain:** Ethereum  
**Protocol:** Dai Stablecoin  
**Function:** Dai Stablecoin symbol  

Call the symbol function on Dai Stablecoin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.dai.symbol",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.dai.totalSupply

**Chain:** Ethereum  
**Protocol:** Dai Stablecoin  
**Function:** Dai Stablecoin totalSupply  

Call the totalSupply function on Dai Stablecoin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.dai.totalSupply",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.dai.transfer

**Chain:** Ethereum  
**Protocol:** Dai Stablecoin  
**Function:** Dai Stablecoin transfer  

Transfer Dai Stablecoin tokens to another address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of Dai Stablecoin tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.dai.transfer",
  "effect": "ALLOW",
  "constraints": {
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.dai.transferFrom

**Chain:** Ethereum  
**Protocol:** Dai Stablecoin  
**Function:** Dai Stablecoin transferFrom  

Transfer Dai Stablecoin tokens from one address to another

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| sender | address | sender parameter of type address |
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of Dai Stablecoin tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.dai.transferFrom",
  "effect": "ALLOW",
  "constraints": {
    "sender": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.eth.transfer

**Chain:** Ethereum  
**Protocol:** Ethereum  
**Function:** Transfer Ethereum ETH  

Transfer Ether to another address on Ethereum

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| recipient | address | The Ethereum address of the recipient |
| amount | decimal | The amount of Ether to transfer |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.eth.transfer",
  "effect": "ALLOW",
  "constraints": {
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.WETH

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.WETH  

Call the WETH function on uniswapv2_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.WETH",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.uniswapv2_router.addLiquidity

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.addLiquidity  

Add liquidity to a token pair with slippage protection

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA - Token contract address for liquidity pair (Must be a valid ERC20 token contract address) |
| tokenB | address | tokenB - Token contract address for liquidity pair (Must be a valid ERC20 token contract address) |
| amountADesired | decimal | amountADesired - Desired amount (maximum you're willing to provide) (Must be positive and within reasonable bounds) |
| amountBDesired | decimal | amountBDesired - Desired amount (maximum you're willing to provide) (Must be positive and within reasonable bounds) |
| amountAMin | decimal | amountAMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| amountBMin | decimal | amountBMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| to | address | to parameter of type address (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.addLiquidity",
  "effect": "ALLOW",
  "constraints": {
    "tokenA": {
      "type": "fixed",
      "value": "example_value"
    },
    "tokenB": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountADesired": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountBDesired": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountAMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountBMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.addLiquidityETH

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.addLiquidityETH  

Add liquidity to an ETH/token pair with slippage protection

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token - Token contract address for liquidity pair (Must be a valid ERC20 token contract address) |
| amountTokenDesired | decimal | amountTokenDesired - Desired amount (maximum you're willing to provide) (Must be positive and within reasonable bounds) |
| amountTokenMin | decimal | amountTokenMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| amountETHMin | decimal | amountETHMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| to | address | to parameter of type address (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |
| value | decimal | The amount of ETH to send with the transaction (Must be positive and within reasonable bounds) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.addLiquidityETH",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountTokenDesired": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountTokenMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountETHMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.factory

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.factory  

Call the factory function on uniswapv2_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.factory",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.uniswapv2_router.getAmountIn

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.getAmountIn  

Call the getAmountIn function on uniswapv2_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 (Must be positive and within reasonable bounds) |
| reserveIn | decimal | reserveIn parameter of type uint256 |
| reserveOut | decimal | reserveOut parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.getAmountIn",
  "effect": "ALLOW",
  "constraints": {
    "amountOut": {
      "type": "fixed",
      "value": "example_value"
    },
    "reserveIn": {
      "type": "fixed",
      "value": "example_value"
    },
    "reserveOut": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.getAmountOut

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.getAmountOut  

Call the getAmountOut function on uniswapv2_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 (Must be positive and within reasonable bounds) |
| reserveIn | decimal | reserveIn parameter of type uint256 |
| reserveOut | decimal | reserveOut parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.getAmountOut",
  "effect": "ALLOW",
  "constraints": {
    "amountIn": {
      "type": "fixed",
      "value": "example_value"
    },
    "reserveIn": {
      "type": "fixed",
      "value": "example_value"
    },
    "reserveOut": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.getAmountsIn

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.getAmountsIn  

Call the getAmountsIn function on uniswapv2_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 (Must be positive and within reasonable bounds) |
| path | array | path parameter of type address[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.getAmountsIn",
  "effect": "ALLOW",
  "constraints": {
    "amountOut": {
      "type": "fixed",
      "value": "example_value"
    },
    "path": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.getAmountsOut

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.getAmountsOut  

Call the getAmountsOut function on uniswapv2_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 (Must be positive and within reasonable bounds) |
| path | array | path parameter of type address[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.getAmountsOut",
  "effect": "ALLOW",
  "constraints": {
    "amountIn": {
      "type": "fixed",
      "value": "example_value"
    },
    "path": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.quote

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.quote  

Call the quote function on uniswapv2_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountA | decimal | amountA parameter of type uint256 (Must be positive and within reasonable bounds) |
| reserveA | decimal | reserveA parameter of type uint256 |
| reserveB | decimal | reserveB parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.quote",
  "effect": "ALLOW",
  "constraints": {
    "amountA": {
      "type": "fixed",
      "value": "example_value"
    },
    "reserveA": {
      "type": "fixed",
      "value": "example_value"
    },
    "reserveB": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.removeLiquidity

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.removeLiquidity  

Remove liquidity from a token pair with minimum output protection

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA - Token contract address for liquidity pair (Must be a valid ERC20 token contract address) |
| tokenB | address | tokenB - Token contract address for liquidity pair (Must be a valid ERC20 token contract address) |
| liquidity | decimal | Amount of LP tokens to remove |
| amountAMin | decimal | amountAMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| amountBMin | decimal | amountBMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| to | address | to parameter of type address (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.removeLiquidity",
  "effect": "ALLOW",
  "constraints": {
    "tokenA": {
      "type": "fixed",
      "value": "example_value"
    },
    "tokenB": {
      "type": "fixed",
      "value": "example_value"
    },
    "liquidity": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountAMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountBMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.removeLiquidityETH

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.removeLiquidityETH  

Remove liquidity from an ETH/token pair with minimum output protection

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token - Token contract address for liquidity pair (Must be a valid ERC20 token contract address) |
| liquidity | decimal | Amount of LP tokens to remove |
| amountTokenMin | decimal | amountTokenMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| amountETHMin | decimal | amountETHMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| to | address | to parameter of type address (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.removeLiquidityETH",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "liquidity": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountTokenMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountETHMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.removeLiquidityETHSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.removeLiquidityETHSupportingFeeOnTransferTokens  

Remove liquidity from ETH/token pair supporting fee-on-transfer tokens

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token - Token contract address for liquidity pair (Must be a valid ERC20 token contract address) |
| liquidity | decimal | Amount of LP tokens to remove |
| amountTokenMin | decimal | amountTokenMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| amountETHMin | decimal | amountETHMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| to | address | to parameter of type address (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.removeLiquidityETHSupportingFeeOnTransferTokens",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "liquidity": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountTokenMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountETHMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.removeLiquidityETHWithPermit

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.removeLiquidityETHWithPermit  

Remove ETH/token liquidity using permit signature

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token - Token contract address for liquidity pair (Must be a valid ERC20 token contract address) |
| liquidity | decimal | Amount of LP tokens to remove |
| amountTokenMin | decimal | amountTokenMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| amountETHMin | decimal | amountETHMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| to | address | to parameter of type address (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |
| approveMax | boolean | Whether to approve maximum uint256 amount for permit |
| v | decimal | v component of permit signature for gas-less approval |
| r | string | r component of permit signature for gas-less approval |
| s | string | s component of permit signature for gas-less approval |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.removeLiquidityETHWithPermit",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "liquidity": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountTokenMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountETHMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },
    "approveMax": {
      "type": "fixed",
      "value": "example_value"
    },
    "v": {
      "type": "fixed",
      "value": "example_value"
    },
    "r": {
      "type": "fixed",
      "value": "example_value"
    },
    "s": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens  

Call the removeLiquidityETHWithPermitSupportingFeeOnTransferTokens function on uniswapv2_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address (Must be a valid ERC20 token contract address) |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 (Must be positive and within reasonable bounds) |
| amountETHMin | decimal | amountETHMin parameter of type uint256 (Must be positive and within reasonable bounds) |
| to | address | to parameter of type address (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |
| approveMax | boolean | approveMax parameter of type bool |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "liquidity": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountTokenMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountETHMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },
    "approveMax": {
      "type": "fixed",
      "value": "example_value"
    },
    "v": {
      "type": "fixed",
      "value": "example_value"
    },
    "r": {
      "type": "fixed",
      "value": "example_value"
    },
    "s": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.removeLiquidityWithPermit

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.removeLiquidityWithPermit  

Remove liquidity using permit signature for gas-less approval

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA - Token contract address for liquidity pair (Must be a valid ERC20 token contract address) |
| tokenB | address | tokenB - Token contract address for liquidity pair (Must be a valid ERC20 token contract address) |
| liquidity | decimal | Amount of LP tokens to remove |
| amountAMin | decimal | amountAMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| amountBMin | decimal | amountBMin - Minimum amount for slippage protection (typically 95-99% of desired) (Must be positive and within reasonable bounds) |
| to | address | to parameter of type address (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |
| approveMax | boolean | Whether to approve maximum uint256 amount for permit |
| v | decimal | v component of permit signature for gas-less approval |
| r | string | r component of permit signature for gas-less approval |
| s | string | s component of permit signature for gas-less approval |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.removeLiquidityWithPermit",
  "effect": "ALLOW",
  "constraints": {
    "tokenA": {
      "type": "fixed",
      "value": "example_value"
    },
    "tokenB": {
      "type": "fixed",
      "value": "example_value"
    },
    "liquidity": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountAMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountBMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },
    "approveMax": {
      "type": "fixed",
      "value": "example_value"
    },
    "v": {
      "type": "fixed",
      "value": "example_value"
    },
    "r": {
      "type": "fixed",
      "value": "example_value"
    },
    "s": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.swapETHForExactTokens

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.swapETHForExactTokens  

Swap ETH for exact amount of tokens with maximum input protection

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | Exact amount of output tokens desired (Must be positive and within reasonable bounds) |
| path | array | Array of token addresses representing the swap path. First address is input token, last is output token |
| to | address | Address that will receive the output tokens. Should be wallet address or approved contract (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |
| value | decimal | The amount of ETH to send with the transaction (Must be positive and within reasonable bounds) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.swapETHForExactTokens",
  "effect": "ALLOW",
  "constraints": {
    "amountOut": {
      "type": "fixed",
      "value": "example_value"
    },
    "path": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.swapExactETHForTokens

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.swapExactETHForTokens  

Swap exact ETH for tokens with minimum output protection

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOutMin | decimal | Minimum amount of output tokens (slippage protection). Should be calculated as: expectedOutput * (1 - slippageTolerance) (Must be positive and within reasonable bounds) |
| path | array | Array of token addresses representing the swap path. First address is input token, last is output token |
| to | address | Address that will receive the output tokens. Should be wallet address or approved contract (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |
| value | decimal | The amount of ETH to send with the transaction (Must be positive and within reasonable bounds) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.swapExactETHForTokens",
  "effect": "ALLOW",
  "constraints": {
    "amountOutMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "path": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.swapExactETHForTokensSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.swapExactETHForTokensSupportingFeeOnTransferTokens  

Call the swapExactETHForTokensSupportingFeeOnTransferTokens function on uniswapv2_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOutMin | decimal | amountOutMin parameter of type uint256 (Must be positive and within reasonable bounds) |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |
| value | decimal | The amount of ETH to send with the transaction (Must be positive and within reasonable bounds) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.swapExactETHForTokensSupportingFeeOnTransferTokens",
  "effect": "ALLOW",
  "constraints": {
    "amountOutMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "path": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.swapExactTokensForETH

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.swapExactTokensForETH  

Swap exact tokens for ETH with minimum output protection

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | Exact amount of input tokens to swap (Must be positive and within reasonable bounds) |
| amountOutMin | decimal | Minimum amount of output tokens (slippage protection). Should be calculated as: expectedOutput * (1 - slippageTolerance) (Must be positive and within reasonable bounds) |
| path | array | Array of token addresses representing the swap path. First address is input token, last is output token |
| to | address | Address that will receive the output tokens. Should be wallet address or approved contract (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.swapExactTokensForETH",
  "effect": "ALLOW",
  "constraints": {
    "amountIn": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountOutMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "path": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.swapExactTokensForETHSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.swapExactTokensForETHSupportingFeeOnTransferTokens  

Call the swapExactTokensForETHSupportingFeeOnTransferTokens function on uniswapv2_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 (Must be positive and within reasonable bounds) |
| amountOutMin | decimal | amountOutMin parameter of type uint256 (Must be positive and within reasonable bounds) |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.swapExactTokensForETHSupportingFeeOnTransferTokens",
  "effect": "ALLOW",
  "constraints": {
    "amountIn": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountOutMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "path": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.swapExactTokensForTokens

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.swapExactTokensForTokens  

Swap exact tokens for other tokens with minimum output protection

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | Exact amount of input tokens to swap (Must be positive and within reasonable bounds) |
| amountOutMin | decimal | Minimum amount of output tokens (slippage protection). Should be calculated as: expectedOutput * (1 - slippageTolerance) (Must be positive and within reasonable bounds) |
| path | array | Array of token addresses representing the swap path. First address is input token, last is output token |
| to | address | Address that will receive the output tokens. Should be wallet address or approved contract (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.swapExactTokensForTokens",
  "effect": "ALLOW",
  "constraints": {
    "amountIn": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountOutMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "path": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.swapExactTokensForTokensSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.swapExactTokensForTokensSupportingFeeOnTransferTokens  

Call the swapExactTokensForTokensSupportingFeeOnTransferTokens function on uniswapv2_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 (Must be positive and within reasonable bounds) |
| amountOutMin | decimal | amountOutMin parameter of type uint256 (Must be positive and within reasonable bounds) |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.swapExactTokensForTokensSupportingFeeOnTransferTokens",
  "effect": "ALLOW",
  "constraints": {
    "amountIn": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountOutMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "path": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.swapTokensForExactETH

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.swapTokensForExactETH  

Swap tokens for exact ETH with maximum input protection

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | Exact amount of output tokens desired (Must be positive and within reasonable bounds) |
| amountInMax | decimal | Maximum amount of input tokens willing to spend (slippage protection for exact output swaps) (Must be positive and within reasonable bounds) |
| path | array | Array of token addresses representing the swap path. First address is input token, last is output token |
| to | address | Address that will receive the output tokens. Should be wallet address or approved contract (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.swapTokensForExactETH",
  "effect": "ALLOW",
  "constraints": {
    "amountOut": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountInMax": {
      "type": "fixed",
      "value": "example_value"
    },
    "path": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv2_router.swapTokensForExactTokens

**Chain:** Ethereum  
**Protocol:** uniswapv2_router  
**Function:** uniswapv2_router.swapTokensForExactTokens  

Swap tokens for exact amount of other tokens with maximum input protection

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | Exact amount of output tokens desired (Must be positive and within reasonable bounds) |
| amountInMax | decimal | Maximum amount of input tokens willing to spend (slippage protection for exact output swaps) (Must be positive and within reasonable bounds) |
| path | array | Array of token addresses representing the swap path. First address is input token, last is output token |
| to | address | Address that will receive the output tokens. Should be wallet address or approved contract (Must be a valid Ethereum address, avoid zero address) |
| deadline | decimal | Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer) |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv2_router.swapTokensForExactTokens",
  "effect": "ALLOW",
  "constraints": {
    "amountOut": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountInMax": {
      "type": "fixed",
      "value": "example_value"
    },
    "path": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.usdc.allowance

**Chain:** Ethereum  
**Protocol:** USD Coin  
**Function:** USD Coin allowance  

Get the amount of USD Coin tokens allowed to be spent by an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| owner | address | owner parameter of type address |
| spender | address | spender parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdc.allowance",
  "effect": "ALLOW",
  "constraints": {
    "owner": {
      "type": "fixed",
      "value": "example_value"
    },
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.usdc.approve

**Chain:** Ethereum  
**Protocol:** USD Coin  
**Function:** USD Coin approve  

Approve an address to spend USD Coin tokens

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| spender | address | spender parameter of type address |
| amount | decimal | The amount of USD Coin tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdc.approve",
  "effect": "ALLOW",
  "constraints": {
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.usdc.balanceOf

**Chain:** Ethereum  
**Protocol:** USD Coin  
**Function:** USD Coin balanceOf  

Get the USD Coin token balance of an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| account | address | account parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdc.balanceOf",
  "effect": "ALLOW",
  "constraints": {
    "account": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.usdc.decimals

**Chain:** Ethereum  
**Protocol:** USD Coin  
**Function:** USD Coin decimals  

Call the decimals function on USD Coin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdc.decimals",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.usdc.name

**Chain:** Ethereum  
**Protocol:** USD Coin  
**Function:** USD Coin name  

Call the name function on USD Coin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdc.name",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.usdc.symbol

**Chain:** Ethereum  
**Protocol:** USD Coin  
**Function:** USD Coin symbol  

Call the symbol function on USD Coin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdc.symbol",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.usdc.totalSupply

**Chain:** Ethereum  
**Protocol:** USD Coin  
**Function:** USD Coin totalSupply  

Call the totalSupply function on USD Coin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdc.totalSupply",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.usdc.transfer

**Chain:** Ethereum  
**Protocol:** USD Coin  
**Function:** USD Coin transfer  

Transfer USD Coin tokens to another address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of USD Coin tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdc.transfer",
  "effect": "ALLOW",
  "constraints": {
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.usdc.transferFrom

**Chain:** Ethereum  
**Protocol:** USD Coin  
**Function:** USD Coin transferFrom  

Transfer USD Coin tokens from one address to another

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| sender | address | sender parameter of type address |
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of USD Coin tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdc.transferFrom",
  "effect": "ALLOW",
  "constraints": {
    "sender": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.usdt.allowance

**Chain:** Ethereum  
**Protocol:** Tether USD  
**Function:** Tether USD allowance  

Get the amount of Tether USD tokens allowed to be spent by an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| owner | address | owner parameter of type address |
| spender | address | spender parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdt.allowance",
  "effect": "ALLOW",
  "constraints": {
    "owner": {
      "type": "fixed",
      "value": "example_value"
    },
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.usdt.approve

**Chain:** Ethereum  
**Protocol:** Tether USD  
**Function:** Tether USD approve  

Approve an address to spend Tether USD tokens

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| spender | address | spender parameter of type address |
| amount | decimal | The amount of Tether USD tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdt.approve",
  "effect": "ALLOW",
  "constraints": {
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.usdt.balanceOf

**Chain:** Ethereum  
**Protocol:** Tether USD  
**Function:** Tether USD balanceOf  

Get the Tether USD token balance of an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| account | address | account parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdt.balanceOf",
  "effect": "ALLOW",
  "constraints": {
    "account": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.usdt.decimals

**Chain:** Ethereum  
**Protocol:** Tether USD  
**Function:** Tether USD decimals  

Call the decimals function on Tether USD token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdt.decimals",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.usdt.name

**Chain:** Ethereum  
**Protocol:** Tether USD  
**Function:** Tether USD name  

Call the name function on Tether USD token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdt.name",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.usdt.symbol

**Chain:** Ethereum  
**Protocol:** Tether USD  
**Function:** Tether USD symbol  

Call the symbol function on Tether USD token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdt.symbol",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.usdt.totalSupply

**Chain:** Ethereum  
**Protocol:** Tether USD  
**Function:** Tether USD totalSupply  

Call the totalSupply function on Tether USD token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdt.totalSupply",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.usdt.transfer

**Chain:** Ethereum  
**Protocol:** Tether USD  
**Function:** Tether USD transfer  

Transfer Tether USD tokens to another address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of Tether USD tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdt.transfer",
  "effect": "ALLOW",
  "constraints": {
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.usdt.transferFrom

**Chain:** Ethereum  
**Protocol:** Tether USD  
**Function:** Tether USD transferFrom  

Transfer Tether USD tokens from one address to another

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| sender | address | sender parameter of type address |
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of Tether USD tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.usdt.transferFrom",
  "effect": "ALLOW",
  "constraints": {
    "sender": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.wbtc.allowance

**Chain:** Ethereum  
**Protocol:** Wrapped Bitcoin  
**Function:** Wrapped Bitcoin allowance  

Get the amount of Wrapped Bitcoin tokens allowed to be spent by an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| owner | address | owner parameter of type address |
| spender | address | spender parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.wbtc.allowance",
  "effect": "ALLOW",
  "constraints": {
    "owner": {
      "type": "fixed",
      "value": "example_value"
    },
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.wbtc.approve

**Chain:** Ethereum  
**Protocol:** Wrapped Bitcoin  
**Function:** Wrapped Bitcoin approve  

Approve an address to spend Wrapped Bitcoin tokens

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| spender | address | spender parameter of type address |
| amount | decimal | The amount of Wrapped Bitcoin tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.wbtc.approve",
  "effect": "ALLOW",
  "constraints": {
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.wbtc.balanceOf

**Chain:** Ethereum  
**Protocol:** Wrapped Bitcoin  
**Function:** Wrapped Bitcoin balanceOf  

Get the Wrapped Bitcoin token balance of an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| account | address | account parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.wbtc.balanceOf",
  "effect": "ALLOW",
  "constraints": {
    "account": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.wbtc.decimals

**Chain:** Ethereum  
**Protocol:** Wrapped Bitcoin  
**Function:** Wrapped Bitcoin decimals  

Call the decimals function on Wrapped Bitcoin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.wbtc.decimals",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.wbtc.name

**Chain:** Ethereum  
**Protocol:** Wrapped Bitcoin  
**Function:** Wrapped Bitcoin name  

Call the name function on Wrapped Bitcoin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.wbtc.name",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.wbtc.symbol

**Chain:** Ethereum  
**Protocol:** Wrapped Bitcoin  
**Function:** Wrapped Bitcoin symbol  

Call the symbol function on Wrapped Bitcoin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.wbtc.symbol",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.wbtc.totalSupply

**Chain:** Ethereum  
**Protocol:** Wrapped Bitcoin  
**Function:** Wrapped Bitcoin totalSupply  

Call the totalSupply function on Wrapped Bitcoin token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.wbtc.totalSupply",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.wbtc.transfer

**Chain:** Ethereum  
**Protocol:** Wrapped Bitcoin  
**Function:** Wrapped Bitcoin transfer  

Transfer Wrapped Bitcoin tokens to another address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of Wrapped Bitcoin tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.wbtc.transfer",
  "effect": "ALLOW",
  "constraints": {
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.wbtc.transferFrom

**Chain:** Ethereum  
**Protocol:** Wrapped Bitcoin  
**Function:** Wrapped Bitcoin transferFrom  

Transfer Wrapped Bitcoin tokens from one address to another

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| sender | address | sender parameter of type address |
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of Wrapped Bitcoin tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.wbtc.transferFrom",
  "effect": "ALLOW",
  "constraints": {
    "sender": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.weth.allowance

**Chain:** Ethereum  
**Protocol:** Wrapped Ether  
**Function:** Wrapped Ether allowance  

Get the amount of Wrapped Ether tokens allowed to be spent by an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| owner | address | owner parameter of type address |
| spender | address | spender parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.weth.allowance",
  "effect": "ALLOW",
  "constraints": {
    "owner": {
      "type": "fixed",
      "value": "example_value"
    },
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.weth.approve

**Chain:** Ethereum  
**Protocol:** Wrapped Ether  
**Function:** Wrapped Ether approve  

Approve an address to spend Wrapped Ether tokens

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| spender | address | spender parameter of type address |
| amount | decimal | The amount of Wrapped Ether tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.weth.approve",
  "effect": "ALLOW",
  "constraints": {
    "spender": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.weth.balanceOf

**Chain:** Ethereum  
**Protocol:** Wrapped Ether  
**Function:** Wrapped Ether balanceOf  

Get the Wrapped Ether token balance of an address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| account | address | account parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.weth.balanceOf",
  "effect": "ALLOW",
  "constraints": {
    "account": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.weth.decimals

**Chain:** Ethereum  
**Protocol:** Wrapped Ether  
**Function:** Wrapped Ether decimals  

Call the decimals function on Wrapped Ether token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.weth.decimals",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.weth.name

**Chain:** Ethereum  
**Protocol:** Wrapped Ether  
**Function:** Wrapped Ether name  

Call the name function on Wrapped Ether token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.weth.name",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.weth.symbol

**Chain:** Ethereum  
**Protocol:** Wrapped Ether  
**Function:** Wrapped Ether symbol  

Call the symbol function on Wrapped Ether token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.weth.symbol",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.weth.totalSupply

**Chain:** Ethereum  
**Protocol:** Wrapped Ether  
**Function:** Wrapped Ether totalSupply  

Call the totalSupply function on Wrapped Ether token

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.weth.totalSupply",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.weth.transfer

**Chain:** Ethereum  
**Protocol:** Wrapped Ether  
**Function:** Wrapped Ether transfer  

Transfer Wrapped Ether tokens to another address

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of Wrapped Ether tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.weth.transfer",
  "effect": "ALLOW",
  "constraints": {
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.weth.transferFrom

**Chain:** Ethereum  
**Protocol:** Wrapped Ether  
**Function:** Wrapped Ether transferFrom  

Transfer Wrapped Ether tokens from one address to another

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| sender | address | sender parameter of type address |
| recipient | address | recipient parameter of type address |
| amount | decimal | The amount of Wrapped Ether tokens |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.weth.transferFrom",
  "effect": "ALLOW",
  "constraints": {
    "sender": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```



## Using Wildcards

You can use wildcards in resource paths to match multiple resources:

* `chain.*.*` - Match all functions in all protocols on a chain
* `chain.protocol.*` - Match all functions in a specific protocol on a chain

## Available Constraint Types

| Type | Description | Example |
|------|-------------|---------|
| fixed | Exact match required | `{"type": "fixed", "value": "0.1"}` |
| max | Maximum value | `{"type": "max", "value": "1.0"}` |
| min | Minimum value | `{"type": "min", "value": "0.01"}` |
| range | Value within range | `{"type": "range", "value": {"min": "0.1", "max": "1.0"}}` |
| whitelist | Value from allowed list | `{"type": "whitelist", "value": ["address1", "address2"]}` |
| max_per_period | Limit actions per time period | `{"type": "max_per_period", "value": 3, "period": "day"}` |
