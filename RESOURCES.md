# Cryptocurrency Wallet Policy Resources

This document lists all available resources that can be used in your wallet plugin policies.
Each resource represents an action that can be performed by a plugin, subject to policy constraints.

## Available Resources


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


### ethereum.aerodrome_factory.MAX_FEE

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.MAX_FEE  

Call the MAX_FEE function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.MAX_FEE",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_factory.ZERO_FEE_INDICATOR

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.ZERO_FEE_INDICATOR  

Call the ZERO_FEE_INDICATOR function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.ZERO_FEE_INDICATOR",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_factory.allPools

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.allPools  

Call the allPools function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| param0 | decimal |  parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.allPools",
  "effect": "ALLOW",
  "constraints": {
    "param0": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.allPoolsLength

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.allPoolsLength  

Call the allPoolsLength function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.allPoolsLength",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_factory.createPool

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.createPool  

Call the createPool function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| fee | decimal | fee parameter of type uint24 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.createPool",
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
    "fee": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.createPool

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.createPool  

Call the createPool function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| stable | boolean | stable parameter of type bool |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.createPool",
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
    "stable": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.customFee

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.customFee  

Call the customFee function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| param0 | address |  parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.customFee",
  "effect": "ALLOW",
  "constraints": {
    "param0": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.feeManager

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.feeManager  

Call the feeManager function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.feeManager",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_factory.getFee

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.getFee  

Call the getFee function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| pool | address | pool parameter of type address |
| _stable | boolean | _stable parameter of type bool |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.getFee",
  "effect": "ALLOW",
  "constraints": {
    "pool": {
      "type": "fixed",
      "value": "example_value"
    },
    "_stable": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.getPool

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.getPool  

Call the getPool function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| fee | decimal | fee parameter of type uint24 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.getPool",
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
    "fee": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.getPool

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.getPool  

Call the getPool function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| stable | boolean | stable parameter of type bool |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.getPool",
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
    "stable": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.implementation

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.implementation  

Call the implementation function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.implementation",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_factory.isPaused

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.isPaused  

Call the isPaused function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.isPaused",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_factory.isPool

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.isPool  

Call the isPool function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| pool | address | pool parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.isPool",
  "effect": "ALLOW",
  "constraints": {
    "pool": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.pauser

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.pauser  

Call the pauser function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.pauser",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_factory.setCustomFee

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.setCustomFee  

Call the setCustomFee function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| pool | address | pool parameter of type address |
| fee | decimal | fee parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.setCustomFee",
  "effect": "ALLOW",
  "constraints": {
    "pool": {
      "type": "fixed",
      "value": "example_value"
    },
    "fee": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.setFee

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.setFee  

Call the setFee function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _stable | boolean | _stable parameter of type bool |
| _fee | decimal | _fee parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.setFee",
  "effect": "ALLOW",
  "constraints": {
    "_stable": {
      "type": "fixed",
      "value": "example_value"
    },
    "_fee": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.setFeeManager

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.setFeeManager  

Call the setFeeManager function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _feeManager | address | _feeManager parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.setFeeManager",
  "effect": "ALLOW",
  "constraints": {
    "_feeManager": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.setPauseState

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.setPauseState  

Call the setPauseState function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _state | boolean | _state parameter of type bool |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.setPauseState",
  "effect": "ALLOW",
  "constraints": {
    "_state": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.setPauser

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.setPauser  

Call the setPauser function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _pauser | address | _pauser parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.setPauser",
  "effect": "ALLOW",
  "constraints": {
    "_pauser": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.setVoter

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.setVoter  

Call the setVoter function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _voter | address | _voter parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.setVoter",
  "effect": "ALLOW",
  "constraints": {
    "_voter": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_factory.stableFee

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.stableFee  

Call the stableFee function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.stableFee",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_factory.volatileFee

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.volatileFee  

Call the volatileFee function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.volatileFee",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_factory.voter

**Chain:** Ethereum  
**Protocol:** aerodrome_factory  
**Function:** aerodrome_factory.voter  

Call the voter function on aerodrome_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_factory.voter",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_router.ETHER

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.ETHER  

Call the ETHER function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.ETHER",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_router.UNSAFE_swapExactTokensForTokens

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.UNSAFE_swapExactTokensForTokens  

Call the UNSAFE_swapExactTokensForTokens function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amounts | decimal | amounts parameter of type uint256[] |
| routes | array | routes parameter of type struct IRouter.Route[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.UNSAFE_swapExactTokensForTokens",
  "effect": "ALLOW",
  "constraints": {
    "amounts": {
      "type": "fixed",
      "value": "example_value"
    },
    "routes": {
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


### ethereum.aerodrome_router.addLiquidity

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.addLiquidity  

Call the addLiquidity function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| stable | boolean | stable parameter of type bool |
| amountADesired | decimal | amountADesired parameter of type uint256 |
| amountBDesired | decimal | amountBDesired parameter of type uint256 |
| amountAMin | decimal | amountAMin parameter of type uint256 |
| amountBMin | decimal | amountBMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.addLiquidity",
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
    "stable": {
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


### ethereum.aerodrome_router.addLiquidityETH

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.addLiquidityETH  

Call the addLiquidityETH function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| stable | boolean | stable parameter of type bool |
| amountTokenDesired | decimal | amountTokenDesired parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.addLiquidityETH",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "stable": {
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


### ethereum.aerodrome_router.defaultFactory

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.defaultFactory  

Call the defaultFactory function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.defaultFactory",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_router.factoryRegistry

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.factoryRegistry  

Call the factoryRegistry function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.factoryRegistry",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_router.generateZapInParams

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.generateZapInParams  

Call the generateZapInParams function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| stable | boolean | stable parameter of type bool |
| _factory | address | _factory parameter of type address |
| amountInA | decimal | amountInA parameter of type uint256 |
| amountInB | decimal | amountInB parameter of type uint256 |
| routesA | array | routesA parameter of type struct IRouter.Route[] |
| routesB | array | routesB parameter of type struct IRouter.Route[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.generateZapInParams",
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
    "stable": {
      "type": "fixed",
      "value": "example_value"
    },
    "_factory": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountInA": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountInB": {
      "type": "fixed",
      "value": "example_value"
    },
    "routesA": {
      "type": "fixed",
      "value": "example_value"
    },
    "routesB": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_router.generateZapOutParams

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.generateZapOutParams  

Call the generateZapOutParams function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| stable | boolean | stable parameter of type bool |
| _factory | address | _factory parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| routesA | array | routesA parameter of type struct IRouter.Route[] |
| routesB | array | routesB parameter of type struct IRouter.Route[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.generateZapOutParams",
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
    "stable": {
      "type": "fixed",
      "value": "example_value"
    },
    "_factory": {
      "type": "fixed",
      "value": "example_value"
    },
    "liquidity": {
      "type": "fixed",
      "value": "example_value"
    },
    "routesA": {
      "type": "fixed",
      "value": "example_value"
    },
    "routesB": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_router.getAmountsOut

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.getAmountsOut  

Call the getAmountsOut function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| routes | array | routes parameter of type struct IRouter.Route[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.getAmountsOut",
  "effect": "ALLOW",
  "constraints": {
    "amountIn": {
      "type": "fixed",
      "value": "example_value"
    },
    "routes": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_router.getReserves

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.getReserves  

Call the getReserves function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| stable | boolean | stable parameter of type bool |
| _factory | address | _factory parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.getReserves",
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
    "stable": {
      "type": "fixed",
      "value": "example_value"
    },
    "_factory": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_router.isTrustedForwarder

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.isTrustedForwarder  

Call the isTrustedForwarder function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| forwarder | address | forwarder parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.isTrustedForwarder",
  "effect": "ALLOW",
  "constraints": {
    "forwarder": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_router.poolFor

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.poolFor  

Call the poolFor function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| stable | boolean | stable parameter of type bool |
| _factory | address | _factory parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.poolFor",
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
    "stable": {
      "type": "fixed",
      "value": "example_value"
    },
    "_factory": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_router.quoteAddLiquidity

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.quoteAddLiquidity  

Call the quoteAddLiquidity function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| stable | boolean | stable parameter of type bool |
| _factory | address | _factory parameter of type address |
| amountADesired | decimal | amountADesired parameter of type uint256 |
| amountBDesired | decimal | amountBDesired parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.quoteAddLiquidity",
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
    "stable": {
      "type": "fixed",
      "value": "example_value"
    },
    "_factory": {
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

  }
}
```


### ethereum.aerodrome_router.quoteRemoveLiquidity

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.quoteRemoveLiquidity  

Call the quoteRemoveLiquidity function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| stable | boolean | stable parameter of type bool |
| _factory | address | _factory parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.quoteRemoveLiquidity",
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
    "stable": {
      "type": "fixed",
      "value": "example_value"
    },
    "_factory": {
      "type": "fixed",
      "value": "example_value"
    },
    "liquidity": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_router.quoteStableLiquidityRatio

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.quoteStableLiquidityRatio  

Call the quoteStableLiquidityRatio function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| _factory | address | _factory parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.quoteStableLiquidityRatio",
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
    "_factory": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.aerodrome_router.removeLiquidity

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.removeLiquidity  

Call the removeLiquidity function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| stable | boolean | stable parameter of type bool |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountAMin | decimal | amountAMin parameter of type uint256 |
| amountBMin | decimal | amountBMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.removeLiquidity",
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
    "stable": {
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


### ethereum.aerodrome_router.removeLiquidityETH

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.removeLiquidityETH  

Call the removeLiquidityETH function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| stable | boolean | stable parameter of type bool |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.removeLiquidityETH",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "stable": {
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


### ethereum.aerodrome_router.removeLiquidityETHSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.removeLiquidityETHSupportingFeeOnTransferTokens  

Call the removeLiquidityETHSupportingFeeOnTransferTokens function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| stable | boolean | stable parameter of type bool |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.removeLiquidityETHSupportingFeeOnTransferTokens",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "stable": {
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


### ethereum.aerodrome_router.sortTokens

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.sortTokens  

Call the sortTokens function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.sortTokens",
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

  }
}
```


### ethereum.aerodrome_router.swapExactETHForTokens

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.swapExactETHForTokens  

Call the swapExactETHForTokens function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| routes | array | routes parameter of type struct IRouter.Route[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.swapExactETHForTokens",
  "effect": "ALLOW",
  "constraints": {
    "amountOutMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "routes": {
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


### ethereum.aerodrome_router.swapExactETHForTokensSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.swapExactETHForTokensSupportingFeeOnTransferTokens  

Call the swapExactETHForTokensSupportingFeeOnTransferTokens function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| routes | array | routes parameter of type struct IRouter.Route[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.swapExactETHForTokensSupportingFeeOnTransferTokens",
  "effect": "ALLOW",
  "constraints": {
    "amountOutMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "routes": {
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


### ethereum.aerodrome_router.swapExactTokensForETH

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.swapExactTokensForETH  

Call the swapExactTokensForETH function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| routes | array | routes parameter of type struct IRouter.Route[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.swapExactTokensForETH",
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
    "routes": {
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


### ethereum.aerodrome_router.swapExactTokensForETHSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.swapExactTokensForETHSupportingFeeOnTransferTokens  

Call the swapExactTokensForETHSupportingFeeOnTransferTokens function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| routes | array | routes parameter of type struct IRouter.Route[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.swapExactTokensForETHSupportingFeeOnTransferTokens",
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
    "routes": {
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


### ethereum.aerodrome_router.swapExactTokensForTokens

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.swapExactTokensForTokens  

Call the swapExactTokensForTokens function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| routes | array | routes parameter of type struct IRouter.Route[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.swapExactTokensForTokens",
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
    "routes": {
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


### ethereum.aerodrome_router.swapExactTokensForTokensSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.swapExactTokensForTokensSupportingFeeOnTransferTokens  

Call the swapExactTokensForTokensSupportingFeeOnTransferTokens function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| routes | array | routes parameter of type struct IRouter.Route[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.swapExactTokensForTokensSupportingFeeOnTransferTokens",
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
    "routes": {
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


### ethereum.aerodrome_router.voter

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.voter  

Call the voter function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.voter",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_router.weth

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.weth  

Call the weth function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.weth",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.aerodrome_router.zapIn

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.zapIn  

Call the zapIn function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenIn | address | tokenIn parameter of type address |
| amountInA | decimal | amountInA parameter of type uint256 |
| amountInB | decimal | amountInB parameter of type uint256 |
| zapInPool | string | zapInPool parameter of type struct IRouter.Zap |
| routesA | array | routesA parameter of type struct IRouter.Route[] |
| routesB | array | routesB parameter of type struct IRouter.Route[] |
| to | address | to parameter of type address |
| stake | boolean | stake parameter of type bool |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.zapIn",
  "effect": "ALLOW",
  "constraints": {
    "tokenIn": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountInA": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountInB": {
      "type": "fixed",
      "value": "example_value"
    },
    "zapInPool": {
      "type": "fixed",
      "value": "example_value"
    },
    "routesA": {
      "type": "fixed",
      "value": "example_value"
    },
    "routesB": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "stake": {
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


### ethereum.aerodrome_router.zapOut

**Chain:** Ethereum  
**Protocol:** aerodrome_router  
**Function:** aerodrome_router.zapOut  

Call the zapOut function on aerodrome_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenOut | address | tokenOut parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| zapOutPool | string | zapOutPool parameter of type struct IRouter.Zap |
| routesA | array | routesA parameter of type struct IRouter.Route[] |
| routesB | array | routesB parameter of type struct IRouter.Route[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.aerodrome_router.zapOut",
  "effect": "ALLOW",
  "constraints": {
    "tokenOut": {
      "type": "fixed",
      "value": "example_value"
    },
    "liquidity": {
      "type": "fixed",
      "value": "example_value"
    },
    "zapOutPool": {
      "type": "fixed",
      "value": "example_value"
    },
    "routesA": {
      "type": "fixed",
      "value": "example_value"
    },
    "routesB": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.camelot_factory.baseFeeConfiguration

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.baseFeeConfiguration  

Call the baseFeeConfiguration function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.baseFeeConfiguration",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.camelot_factory.createPool

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.createPool  

Call the createPool function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.createPool",
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

  }
}
```


### ethereum.camelot_factory.defaultCommunityFee

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.defaultCommunityFee  

Call the defaultCommunityFee function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.defaultCommunityFee",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.camelot_factory.farmingAddress

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.farmingAddress  

Call the farmingAddress function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.farmingAddress",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.camelot_factory.owner

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.owner  

Call the owner function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.owner",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.camelot_factory.poolByPair

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.poolByPair  

Call the poolByPair function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| param0 | address |  parameter of type address |
| param1 | address |  parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.poolByPair",
  "effect": "ALLOW",
  "constraints": {
    "param0": {
      "type": "fixed",
      "value": "example_value"
    },
    "param1": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.camelot_factory.poolDeployer

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.poolDeployer  

Call the poolDeployer function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.poolDeployer",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.camelot_factory.setBaseFeeConfiguration

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.setBaseFeeConfiguration  

Call the setBaseFeeConfiguration function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| alpha1 | decimal | alpha1 parameter of type uint16 |
| alpha2 | decimal | alpha2 parameter of type uint16 |
| beta1 | decimal | beta1 parameter of type uint32 |
| beta2 | decimal | beta2 parameter of type uint32 |
| gamma1 | decimal | gamma1 parameter of type uint16 |
| gamma2 | decimal | gamma2 parameter of type uint16 |
| volumeBeta | decimal | volumeBeta parameter of type uint32 |
| volumeGamma | decimal | volumeGamma parameter of type uint16 |
| baseFee | decimal | baseFee parameter of type uint16 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.setBaseFeeConfiguration",
  "effect": "ALLOW",
  "constraints": {
    "alpha1": {
      "type": "fixed",
      "value": "example_value"
    },
    "alpha2": {
      "type": "fixed",
      "value": "example_value"
    },
    "beta1": {
      "type": "fixed",
      "value": "example_value"
    },
    "beta2": {
      "type": "fixed",
      "value": "example_value"
    },
    "gamma1": {
      "type": "fixed",
      "value": "example_value"
    },
    "gamma2": {
      "type": "fixed",
      "value": "example_value"
    },
    "volumeBeta": {
      "type": "fixed",
      "value": "example_value"
    },
    "volumeGamma": {
      "type": "fixed",
      "value": "example_value"
    },
    "baseFee": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.camelot_factory.setDefaultCommunityFee

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.setDefaultCommunityFee  

Call the setDefaultCommunityFee function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| newDefaultCommunityFee | decimal | newDefaultCommunityFee parameter of type uint8 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.setDefaultCommunityFee",
  "effect": "ALLOW",
  "constraints": {
    "newDefaultCommunityFee": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.camelot_factory.setFarmingAddress

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.setFarmingAddress  

Call the setFarmingAddress function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _farmingAddress | address | _farmingAddress parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.setFarmingAddress",
  "effect": "ALLOW",
  "constraints": {
    "_farmingAddress": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.camelot_factory.setOwner

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.setOwner  

Call the setOwner function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _owner | address | _owner parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.setOwner",
  "effect": "ALLOW",
  "constraints": {
    "_owner": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.camelot_factory.setVaultAddress

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.setVaultAddress  

Call the setVaultAddress function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _vaultAddress | address | _vaultAddress parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.setVaultAddress",
  "effect": "ALLOW",
  "constraints": {
    "_vaultAddress": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.camelot_factory.vaultAddress

**Chain:** Ethereum  
**Protocol:** camelot_factory  
**Function:** camelot_factory.vaultAddress  

Call the vaultAddress function on camelot_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_factory.vaultAddress",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.camelot_router.WETH

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.WETH  

Call the WETH function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.WETH",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.camelot_router.addLiquidity

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.addLiquidity  

Call the addLiquidity function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| amountADesired | decimal | amountADesired parameter of type uint256 |
| amountBDesired | decimal | amountBDesired parameter of type uint256 |
| amountAMin | decimal | amountAMin parameter of type uint256 |
| amountBMin | decimal | amountBMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.addLiquidity",
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


### ethereum.camelot_router.addLiquidityETH

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.addLiquidityETH  

Call the addLiquidityETH function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amountTokenDesired | decimal | amountTokenDesired parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.addLiquidityETH",
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


### ethereum.camelot_router.factory

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.factory  

Call the factory function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.factory",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.camelot_router.getAmountsOut

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.getAmountsOut  

Call the getAmountsOut function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| path | array | path parameter of type address[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.getAmountsOut",
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


### ethereum.camelot_router.getPair

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.getPair  

Call the getPair function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token1 | address | token1 parameter of type address |
| token2 | address | token2 parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.getPair",
  "effect": "ALLOW",
  "constraints": {
    "token1": {
      "type": "fixed",
      "value": "example_value"
    },
    "token2": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.camelot_router.quote

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.quote  

Call the quote function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountA | decimal | amountA parameter of type uint256 |
| reserveA | decimal | reserveA parameter of type uint256 |
| reserveB | decimal | reserveB parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.quote",
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


### ethereum.camelot_router.removeLiquidity

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.removeLiquidity  

Call the removeLiquidity function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountAMin | decimal | amountAMin parameter of type uint256 |
| amountBMin | decimal | amountBMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.removeLiquidity",
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


### ethereum.camelot_router.removeLiquidityETH

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.removeLiquidityETH  

Call the removeLiquidityETH function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.removeLiquidityETH",
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


### ethereum.camelot_router.removeLiquidityETHSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.removeLiquidityETHSupportingFeeOnTransferTokens  

Call the removeLiquidityETHSupportingFeeOnTransferTokens function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.removeLiquidityETHSupportingFeeOnTransferTokens",
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


### ethereum.camelot_router.removeLiquidityETHWithPermit

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.removeLiquidityETHWithPermit  

Call the removeLiquidityETHWithPermit function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| approveMax | boolean | approveMax parameter of type bool |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.removeLiquidityETHWithPermit",
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


### ethereum.camelot_router.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens  

Call the removeLiquidityETHWithPermitSupportingFeeOnTransferTokens function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| approveMax | boolean | approveMax parameter of type bool |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens",
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


### ethereum.camelot_router.removeLiquidityWithPermit

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.removeLiquidityWithPermit  

Call the removeLiquidityWithPermit function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountAMin | decimal | amountAMin parameter of type uint256 |
| amountBMin | decimal | amountBMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| approveMax | boolean | approveMax parameter of type bool |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.removeLiquidityWithPermit",
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


### ethereum.camelot_router.swapExactETHForTokensSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.swapExactETHForTokensSupportingFeeOnTransferTokens  

Call the swapExactETHForTokensSupportingFeeOnTransferTokens function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| referrer | address | referrer parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.swapExactETHForTokensSupportingFeeOnTransferTokens",
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
    "referrer": {
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


### ethereum.camelot_router.swapExactTokensForETHSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.swapExactTokensForETHSupportingFeeOnTransferTokens  

Call the swapExactTokensForETHSupportingFeeOnTransferTokens function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| referrer | address | referrer parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.swapExactTokensForETHSupportingFeeOnTransferTokens",
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
    "referrer": {
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


### ethereum.camelot_router.swapExactTokensForTokensSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** camelot_router  
**Function:** camelot_router.swapExactTokensForTokensSupportingFeeOnTransferTokens  

Call the swapExactTokensForTokensSupportingFeeOnTransferTokens function on camelot_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| referrer | address | referrer parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.camelot_router.swapExactTokensForTokensSupportingFeeOnTransferTokens",
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
    "referrer": {
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
**Function:** Transfer ETH  

Transfer Ether to another address

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


### ethereum.joe_factory.addQuoteAsset

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.addQuoteAsset  

Call the addQuoteAsset function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| quoteAsset | address | quoteAsset parameter of type contract IERC20 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.addQuoteAsset",
  "effect": "ALLOW",
  "constraints": {
    "quoteAsset": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.becomeOwner

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.becomeOwner  

Call the becomeOwner function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.becomeOwner",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.createLBPair

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.createLBPair  

Call the createLBPair function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenX | address | tokenX parameter of type contract IERC20 |
| tokenY | address | tokenY parameter of type contract IERC20 |
| activeId | decimal | activeId parameter of type uint24 |
| binStep | decimal | binStep parameter of type uint16 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.createLBPair",
  "effect": "ALLOW",
  "constraints": {
    "tokenX": {
      "type": "fixed",
      "value": "example_value"
    },
    "tokenY": {
      "type": "fixed",
      "value": "example_value"
    },
    "activeId": {
      "type": "fixed",
      "value": "example_value"
    },
    "binStep": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.forceDecay

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.forceDecay  

Call the forceDecay function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| pair | address | pair parameter of type contract ILBPair |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.forceDecay",
  "effect": "ALLOW",
  "constraints": {
    "pair": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.getAllBinSteps

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getAllBinSteps  

Call the getAllBinSteps function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getAllBinSteps",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.getAllLBPairs

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getAllLBPairs  

Call the getAllLBPairs function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenX | address | tokenX parameter of type contract IERC20 |
| tokenY | address | tokenY parameter of type contract IERC20 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getAllLBPairs",
  "effect": "ALLOW",
  "constraints": {
    "tokenX": {
      "type": "fixed",
      "value": "example_value"
    },
    "tokenY": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.getFeeRecipient

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getFeeRecipient  

Call the getFeeRecipient function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getFeeRecipient",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.getFlashLoanFee

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getFlashLoanFee  

Call the getFlashLoanFee function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getFlashLoanFee",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.getLBPairAtIndex

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getLBPairAtIndex  

Call the getLBPairAtIndex function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| index | decimal | index parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getLBPairAtIndex",
  "effect": "ALLOW",
  "constraints": {
    "index": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.getLBPairImplementation

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getLBPairImplementation  

Call the getLBPairImplementation function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getLBPairImplementation",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.getLBPairInformation

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getLBPairInformation  

Call the getLBPairInformation function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type contract IERC20 |
| tokenB | address | tokenB parameter of type contract IERC20 |
| binStep | decimal | binStep parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getLBPairInformation",
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
    "binStep": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.getMaxFlashLoanFee

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getMaxFlashLoanFee  

Call the getMaxFlashLoanFee function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getMaxFlashLoanFee",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.getMinBinStep

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getMinBinStep  

Call the getMinBinStep function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getMinBinStep",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.getNumberOfLBPairs

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getNumberOfLBPairs  

Call the getNumberOfLBPairs function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getNumberOfLBPairs",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.getNumberOfQuoteAssets

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getNumberOfQuoteAssets  

Call the getNumberOfQuoteAssets function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getNumberOfQuoteAssets",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.getOpenBinSteps

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getOpenBinSteps  

Call the getOpenBinSteps function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getOpenBinSteps",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.getPreset

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getPreset  

Call the getPreset function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| binStep | decimal | binStep parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getPreset",
  "effect": "ALLOW",
  "constraints": {
    "binStep": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.getQuoteAssetAtIndex

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.getQuoteAssetAtIndex  

Call the getQuoteAssetAtIndex function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| index | decimal | index parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.getQuoteAssetAtIndex",
  "effect": "ALLOW",
  "constraints": {
    "index": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.isQuoteAsset

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.isQuoteAsset  

Call the isQuoteAsset function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type contract IERC20 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.isQuoteAsset",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.owner

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.owner  

Call the owner function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.owner",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.pendingOwner

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.pendingOwner  

Call the pendingOwner function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.pendingOwner",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.removePreset

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.removePreset  

Call the removePreset function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| binStep | decimal | binStep parameter of type uint16 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.removePreset",
  "effect": "ALLOW",
  "constraints": {
    "binStep": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.removeQuoteAsset

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.removeQuoteAsset  

Call the removeQuoteAsset function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| quoteAsset | address | quoteAsset parameter of type contract IERC20 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.removeQuoteAsset",
  "effect": "ALLOW",
  "constraints": {
    "quoteAsset": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.renounceOwnership

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.renounceOwnership  

Call the renounceOwnership function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.renounceOwnership",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.revokePendingOwner

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.revokePendingOwner  

Call the revokePendingOwner function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.revokePendingOwner",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_factory.setFeeRecipient

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.setFeeRecipient  

Call the setFeeRecipient function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| feeRecipient | address | feeRecipient parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.setFeeRecipient",
  "effect": "ALLOW",
  "constraints": {
    "feeRecipient": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.setFeesParametersOnPair

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.setFeesParametersOnPair  

Call the setFeesParametersOnPair function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenX | address | tokenX parameter of type contract IERC20 |
| tokenY | address | tokenY parameter of type contract IERC20 |
| binStep | decimal | binStep parameter of type uint16 |
| baseFactor | decimal | baseFactor parameter of type uint16 |
| filterPeriod | decimal | filterPeriod parameter of type uint16 |
| decayPeriod | decimal | decayPeriod parameter of type uint16 |
| reductionFactor | decimal | reductionFactor parameter of type uint16 |
| variableFeeControl | decimal | variableFeeControl parameter of type uint24 |
| protocolShare | decimal | protocolShare parameter of type uint16 |
| maxVolatilityAccumulator | decimal | maxVolatilityAccumulator parameter of type uint24 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.setFeesParametersOnPair",
  "effect": "ALLOW",
  "constraints": {
    "tokenX": {
      "type": "fixed",
      "value": "example_value"
    },
    "tokenY": {
      "type": "fixed",
      "value": "example_value"
    },
    "binStep": {
      "type": "fixed",
      "value": "example_value"
    },
    "baseFactor": {
      "type": "fixed",
      "value": "example_value"
    },
    "filterPeriod": {
      "type": "fixed",
      "value": "example_value"
    },
    "decayPeriod": {
      "type": "fixed",
      "value": "example_value"
    },
    "reductionFactor": {
      "type": "fixed",
      "value": "example_value"
    },
    "variableFeeControl": {
      "type": "fixed",
      "value": "example_value"
    },
    "protocolShare": {
      "type": "fixed",
      "value": "example_value"
    },
    "maxVolatilityAccumulator": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.setFlashLoanFee

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.setFlashLoanFee  

Call the setFlashLoanFee function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| flashLoanFee | decimal | flashLoanFee parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.setFlashLoanFee",
  "effect": "ALLOW",
  "constraints": {
    "flashLoanFee": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.setLBPairIgnored

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.setLBPairIgnored  

Call the setLBPairIgnored function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenX | address | tokenX parameter of type contract IERC20 |
| tokenY | address | tokenY parameter of type contract IERC20 |
| binStep | decimal | binStep parameter of type uint16 |
| ignored | boolean | ignored parameter of type bool |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.setLBPairIgnored",
  "effect": "ALLOW",
  "constraints": {
    "tokenX": {
      "type": "fixed",
      "value": "example_value"
    },
    "tokenY": {
      "type": "fixed",
      "value": "example_value"
    },
    "binStep": {
      "type": "fixed",
      "value": "example_value"
    },
    "ignored": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.setLBPairImplementation

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.setLBPairImplementation  

Call the setLBPairImplementation function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| newLBPairImplementation | address | newLBPairImplementation parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.setLBPairImplementation",
  "effect": "ALLOW",
  "constraints": {
    "newLBPairImplementation": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.setPendingOwner

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.setPendingOwner  

Call the setPendingOwner function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| pendingOwner_ | address | pendingOwner_ parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.setPendingOwner",
  "effect": "ALLOW",
  "constraints": {
    "pendingOwner_": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.setPreset

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.setPreset  

Call the setPreset function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| binStep | decimal | binStep parameter of type uint16 |
| baseFactor | decimal | baseFactor parameter of type uint16 |
| filterPeriod | decimal | filterPeriod parameter of type uint16 |
| decayPeriod | decimal | decayPeriod parameter of type uint16 |
| reductionFactor | decimal | reductionFactor parameter of type uint16 |
| variableFeeControl | decimal | variableFeeControl parameter of type uint24 |
| protocolShare | decimal | protocolShare parameter of type uint16 |
| maxVolatilityAccumulator | decimal | maxVolatilityAccumulator parameter of type uint24 |
| isOpen | boolean | isOpen parameter of type bool |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.setPreset",
  "effect": "ALLOW",
  "constraints": {
    "binStep": {
      "type": "fixed",
      "value": "example_value"
    },
    "baseFactor": {
      "type": "fixed",
      "value": "example_value"
    },
    "filterPeriod": {
      "type": "fixed",
      "value": "example_value"
    },
    "decayPeriod": {
      "type": "fixed",
      "value": "example_value"
    },
    "reductionFactor": {
      "type": "fixed",
      "value": "example_value"
    },
    "variableFeeControl": {
      "type": "fixed",
      "value": "example_value"
    },
    "protocolShare": {
      "type": "fixed",
      "value": "example_value"
    },
    "maxVolatilityAccumulator": {
      "type": "fixed",
      "value": "example_value"
    },
    "isOpen": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_factory.setPresetOpenState

**Chain:** Ethereum  
**Protocol:** joe_factory  
**Function:** joe_factory.setPresetOpenState  

Call the setPresetOpenState function on joe_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| binStep | decimal | binStep parameter of type uint16 |
| isOpen | boolean | isOpen parameter of type bool |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_factory.setPresetOpenState",
  "effect": "ALLOW",
  "constraints": {
    "binStep": {
      "type": "fixed",
      "value": "example_value"
    },
    "isOpen": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.joe_router.WAVAX

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.WAVAX  

Call the WAVAX function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.WAVAX",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_router.addLiquidity

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.addLiquidity  

Call the addLiquidity function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| amountADesired | decimal | amountADesired parameter of type uint256 |
| amountBDesired | decimal | amountBDesired parameter of type uint256 |
| amountAMin | decimal | amountAMin parameter of type uint256 |
| amountBMin | decimal | amountBMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.addLiquidity",
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


### ethereum.joe_router.addLiquidityAVAX

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.addLiquidityAVAX  

Call the addLiquidityAVAX function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amountTokenDesired | decimal | amountTokenDesired parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountAVAXMin | decimal | amountAVAXMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.addLiquidityAVAX",
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
    "amountAVAXMin": {
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


### ethereum.joe_router.factory

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.factory  

Call the factory function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.factory",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.joe_router.getAmountIn

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.getAmountIn  

Call the getAmountIn function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 |
| reserveIn | decimal | reserveIn parameter of type uint256 |
| reserveOut | decimal | reserveOut parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.getAmountIn",
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


### ethereum.joe_router.getAmountOut

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.getAmountOut  

Call the getAmountOut function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| reserveIn | decimal | reserveIn parameter of type uint256 |
| reserveOut | decimal | reserveOut parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.getAmountOut",
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


### ethereum.joe_router.getAmountsIn

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.getAmountsIn  

Call the getAmountsIn function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 |
| path | array | path parameter of type address[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.getAmountsIn",
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


### ethereum.joe_router.getAmountsOut

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.getAmountsOut  

Call the getAmountsOut function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| path | array | path parameter of type address[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.getAmountsOut",
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


### ethereum.joe_router.quote

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.quote  

Call the quote function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountA | decimal | amountA parameter of type uint256 |
| reserveA | decimal | reserveA parameter of type uint256 |
| reserveB | decimal | reserveB parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.quote",
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


### ethereum.joe_router.removeLiquidity

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.removeLiquidity  

Call the removeLiquidity function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountAMin | decimal | amountAMin parameter of type uint256 |
| amountBMin | decimal | amountBMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.removeLiquidity",
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


### ethereum.joe_router.removeLiquidityAVAX

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.removeLiquidityAVAX  

Call the removeLiquidityAVAX function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountAVAXMin | decimal | amountAVAXMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.removeLiquidityAVAX",
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
    "amountAVAXMin": {
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


### ethereum.joe_router.removeLiquidityAVAXSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.removeLiquidityAVAXSupportingFeeOnTransferTokens  

Call the removeLiquidityAVAXSupportingFeeOnTransferTokens function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountAVAXMin | decimal | amountAVAXMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.removeLiquidityAVAXSupportingFeeOnTransferTokens",
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
    "amountAVAXMin": {
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


### ethereum.joe_router.removeLiquidityAVAXWithPermit

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.removeLiquidityAVAXWithPermit  

Call the removeLiquidityAVAXWithPermit function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountAVAXMin | decimal | amountAVAXMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| approveMax | boolean | approveMax parameter of type bool |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.removeLiquidityAVAXWithPermit",
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
    "amountAVAXMin": {
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


### ethereum.joe_router.removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens  

Call the removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountAVAXMin | decimal | amountAVAXMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| approveMax | boolean | approveMax parameter of type bool |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens",
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
    "amountAVAXMin": {
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


### ethereum.joe_router.removeLiquidityWithPermit

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.removeLiquidityWithPermit  

Call the removeLiquidityWithPermit function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountAMin | decimal | amountAMin parameter of type uint256 |
| amountBMin | decimal | amountBMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| approveMax | boolean | approveMax parameter of type bool |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.removeLiquidityWithPermit",
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


### ethereum.joe_router.swapAVAXForExactTokens

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.swapAVAXForExactTokens  

Call the swapAVAXForExactTokens function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.swapAVAXForExactTokens",
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


### ethereum.joe_router.swapExactAVAXForTokens

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.swapExactAVAXForTokens  

Call the swapExactAVAXForTokens function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.swapExactAVAXForTokens",
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


### ethereum.joe_router.swapExactAVAXForTokensSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.swapExactAVAXForTokensSupportingFeeOnTransferTokens  

Call the swapExactAVAXForTokensSupportingFeeOnTransferTokens function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.swapExactAVAXForTokensSupportingFeeOnTransferTokens",
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


### ethereum.joe_router.swapExactTokensForAVAX

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.swapExactTokensForAVAX  

Call the swapExactTokensForAVAX function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.swapExactTokensForAVAX",
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


### ethereum.joe_router.swapExactTokensForAVAXSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.swapExactTokensForAVAXSupportingFeeOnTransferTokens  

Call the swapExactTokensForAVAXSupportingFeeOnTransferTokens function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.swapExactTokensForAVAXSupportingFeeOnTransferTokens",
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


### ethereum.joe_router.swapExactTokensForTokens

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.swapExactTokensForTokens  

Call the swapExactTokensForTokens function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.swapExactTokensForTokens",
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


### ethereum.joe_router.swapExactTokensForTokensSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.swapExactTokensForTokensSupportingFeeOnTransferTokens  

Call the swapExactTokensForTokensSupportingFeeOnTransferTokens function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.swapExactTokensForTokensSupportingFeeOnTransferTokens",
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


### ethereum.joe_router.swapTokensForExactAVAX

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.swapTokensForExactAVAX  

Call the swapTokensForExactAVAX function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 |
| amountInMax | decimal | amountInMax parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.swapTokensForExactAVAX",
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


### ethereum.joe_router.swapTokensForExactTokens

**Chain:** Ethereum  
**Protocol:** joe_router  
**Function:** joe_router.swapTokensForExactTokens  

Call the swapTokensForExactTokens function on joe_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 |
| amountInMax | decimal | amountInMax parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.joe_router.swapTokensForExactTokens",
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


### ethereum.pancakev3_factory.collectProtocol

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.collectProtocol  

Call the collectProtocol function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| pool | address | pool parameter of type address |
| recipient | address | recipient parameter of type address |
| amount0Requested | decimal | amount0Requested parameter of type uint128 |
| amount1Requested | decimal | amount1Requested parameter of type uint128 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.collectProtocol",
  "effect": "ALLOW",
  "constraints": {
    "pool": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount0Requested": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount1Requested": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_factory.createPool

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.createPool  

Call the createPool function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| fee | decimal | fee parameter of type uint24 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.createPool",
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
    "fee": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_factory.enableFeeAmount

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.enableFeeAmount  

Call the enableFeeAmount function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| fee | decimal | fee parameter of type uint24 |
| tickSpacing | decimal | tickSpacing parameter of type int24 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.enableFeeAmount",
  "effect": "ALLOW",
  "constraints": {
    "fee": {
      "type": "fixed",
      "value": "example_value"
    },
    "tickSpacing": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_factory.feeAmountTickSpacing

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.feeAmountTickSpacing  

Call the feeAmountTickSpacing function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| param0 | decimal |  parameter of type uint24 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.feeAmountTickSpacing",
  "effect": "ALLOW",
  "constraints": {
    "param0": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_factory.feeAmountTickSpacingExtraInfo

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.feeAmountTickSpacingExtraInfo  

Call the feeAmountTickSpacingExtraInfo function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| param0 | decimal |  parameter of type uint24 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.feeAmountTickSpacingExtraInfo",
  "effect": "ALLOW",
  "constraints": {
    "param0": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_factory.getPool

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.getPool  

Call the getPool function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| param0 | address |  parameter of type address |
| param1 | address |  parameter of type address |
| param2 | decimal |  parameter of type uint24 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.getPool",
  "effect": "ALLOW",
  "constraints": {
    "param0": {
      "type": "fixed",
      "value": "example_value"
    },
    "param1": {
      "type": "fixed",
      "value": "example_value"
    },
    "param2": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_factory.lmPoolDeployer

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.lmPoolDeployer  

Call the lmPoolDeployer function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.lmPoolDeployer",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.pancakev3_factory.owner

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.owner  

Call the owner function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.owner",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.pancakev3_factory.poolDeployer

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.poolDeployer  

Call the poolDeployer function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.poolDeployer",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.pancakev3_factory.setFeeAmountExtraInfo

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.setFeeAmountExtraInfo  

Call the setFeeAmountExtraInfo function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| fee | decimal | fee parameter of type uint24 |
| whitelistRequested | boolean | whitelistRequested parameter of type bool |
| enabled | boolean | enabled parameter of type bool |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.setFeeAmountExtraInfo",
  "effect": "ALLOW",
  "constraints": {
    "fee": {
      "type": "fixed",
      "value": "example_value"
    },
    "whitelistRequested": {
      "type": "fixed",
      "value": "example_value"
    },
    "enabled": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_factory.setFeeProtocol

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.setFeeProtocol  

Call the setFeeProtocol function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| pool | address | pool parameter of type address |
| feeProtocol0 | decimal | feeProtocol0 parameter of type uint32 |
| feeProtocol1 | decimal | feeProtocol1 parameter of type uint32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.setFeeProtocol",
  "effect": "ALLOW",
  "constraints": {
    "pool": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeProtocol0": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeProtocol1": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_factory.setLmPool

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.setLmPool  

Call the setLmPool function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| pool | address | pool parameter of type address |
| lmPool | address | lmPool parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.setLmPool",
  "effect": "ALLOW",
  "constraints": {
    "pool": {
      "type": "fixed",
      "value": "example_value"
    },
    "lmPool": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_factory.setLmPoolDeployer

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.setLmPoolDeployer  

Call the setLmPoolDeployer function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _lmPoolDeployer | address | _lmPoolDeployer parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.setLmPoolDeployer",
  "effect": "ALLOW",
  "constraints": {
    "_lmPoolDeployer": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_factory.setOwner

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.setOwner  

Call the setOwner function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _owner | address | _owner parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.setOwner",
  "effect": "ALLOW",
  "constraints": {
    "_owner": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_factory.setWhiteListAddress

**Chain:** Ethereum  
**Protocol:** pancakev3_factory  
**Function:** pancakev3_factory.setWhiteListAddress  

Call the setWhiteListAddress function on pancakev3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| user | address | user parameter of type address |
| verified | boolean | verified parameter of type bool |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_factory.setWhiteListAddress",
  "effect": "ALLOW",
  "constraints": {
    "user": {
      "type": "fixed",
      "value": "example_value"
    },
    "verified": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_router.WETH9

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.WETH9  

Call the WETH9 function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.WETH9",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.pancakev3_router.deployer

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.deployer  

Call the deployer function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.deployer",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.pancakev3_router.exactInput

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.exactInput  

Call the exactInput function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct ISwapRouter.ExactInputParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.exactInput",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.pancakev3_router.exactInputSingle

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.exactInputSingle  

Call the exactInputSingle function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct ISwapRouter.ExactInputSingleParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.exactInputSingle",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.pancakev3_router.exactOutput

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.exactOutput  

Call the exactOutput function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct ISwapRouter.ExactOutputParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.exactOutput",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.pancakev3_router.exactOutputSingle

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.exactOutputSingle  

Call the exactOutputSingle function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct ISwapRouter.ExactOutputSingleParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.exactOutputSingle",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.pancakev3_router.factory

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.factory  

Call the factory function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.factory",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.pancakev3_router.multicall

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.multicall  

Call the multicall function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| data | array | data parameter of type bytes[] |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.multicall",
  "effect": "ALLOW",
  "constraints": {
    "data": {
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


### ethereum.pancakev3_router.pancakeV3SwapCallback

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.pancakeV3SwapCallback  

Call the pancakeV3SwapCallback function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amount0Delta | decimal | amount0Delta parameter of type int256 |
| amount1Delta | decimal | amount1Delta parameter of type int256 |
| _data | bytes | _data parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.pancakeV3SwapCallback",
  "effect": "ALLOW",
  "constraints": {
    "amount0Delta": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount1Delta": {
      "type": "fixed",
      "value": "example_value"
    },
    "_data": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_router.refundETH

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.refundETH  

Call the refundETH function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.refundETH",
  "effect": "ALLOW",
  "constraints": {
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_router.selfPermit

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.selfPermit  

Call the selfPermit function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| value | decimal | value parameter of type uint256 |
| deadline | decimal | deadline parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.selfPermit",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "value": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_router.selfPermitAllowed

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.selfPermitAllowed  

Call the selfPermitAllowed function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| nonce | decimal | nonce parameter of type uint256 |
| expiry | decimal | expiry parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.selfPermitAllowed",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "nonce": {
      "type": "fixed",
      "value": "example_value"
    },
    "expiry": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_router.selfPermitAllowedIfNecessary

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.selfPermitAllowedIfNecessary  

Call the selfPermitAllowedIfNecessary function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| nonce | decimal | nonce parameter of type uint256 |
| expiry | decimal | expiry parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.selfPermitAllowedIfNecessary",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "nonce": {
      "type": "fixed",
      "value": "example_value"
    },
    "expiry": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_router.selfPermitIfNecessary

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.selfPermitIfNecessary  

Call the selfPermitIfNecessary function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| value | decimal | value parameter of type uint256 |
| deadline | decimal | deadline parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.selfPermitIfNecessary",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "value": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.pancakev3_router.sweepToken

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.sweepToken  

Call the sweepToken function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.sweepToken",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
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


### ethereum.pancakev3_router.sweepTokenWithFee

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.sweepTokenWithFee  

Call the sweepTokenWithFee function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| feeBips | decimal | feeBips parameter of type uint256 |
| feeRecipient | address | feeRecipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.sweepTokenWithFee",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeBips": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeRecipient": {
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


### ethereum.pancakev3_router.unwrapWETH9

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.unwrapWETH9  

Call the unwrapWETH9 function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.unwrapWETH9",
  "effect": "ALLOW",
  "constraints": {
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
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


### ethereum.pancakev3_router.unwrapWETH9WithFee

**Chain:** Ethereum  
**Protocol:** pancakev3_router  
**Function:** pancakev3_router.unwrapWETH9WithFee  

Call the unwrapWETH9WithFee function on pancakev3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| feeBips | decimal | feeBips parameter of type uint256 |
| feeRecipient | address | feeRecipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.pancakev3_router.unwrapWETH9WithFee",
  "effect": "ALLOW",
  "constraints": {
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeBips": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeRecipient": {
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


### ethereum.quickswapv3_factory.baseFeeConfiguration

**Chain:** Ethereum  
**Protocol:** quickswapv3_factory  
**Function:** quickswapv3_factory.baseFeeConfiguration  

Call the baseFeeConfiguration function on quickswapv3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_factory.baseFeeConfiguration",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.quickswapv3_factory.createPool

**Chain:** Ethereum  
**Protocol:** quickswapv3_factory  
**Function:** quickswapv3_factory.createPool  

Call the createPool function on quickswapv3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_factory.createPool",
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

  }
}
```


### ethereum.quickswapv3_factory.farmingAddress

**Chain:** Ethereum  
**Protocol:** quickswapv3_factory  
**Function:** quickswapv3_factory.farmingAddress  

Call the farmingAddress function on quickswapv3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_factory.farmingAddress",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.quickswapv3_factory.owner

**Chain:** Ethereum  
**Protocol:** quickswapv3_factory  
**Function:** quickswapv3_factory.owner  

Call the owner function on quickswapv3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_factory.owner",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.quickswapv3_factory.poolByPair

**Chain:** Ethereum  
**Protocol:** quickswapv3_factory  
**Function:** quickswapv3_factory.poolByPair  

Call the poolByPair function on quickswapv3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| param0 | address |  parameter of type address |
| param1 | address |  parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_factory.poolByPair",
  "effect": "ALLOW",
  "constraints": {
    "param0": {
      "type": "fixed",
      "value": "example_value"
    },
    "param1": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_factory.poolDeployer

**Chain:** Ethereum  
**Protocol:** quickswapv3_factory  
**Function:** quickswapv3_factory.poolDeployer  

Call the poolDeployer function on quickswapv3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_factory.poolDeployer",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.quickswapv3_factory.setBaseFeeConfiguration

**Chain:** Ethereum  
**Protocol:** quickswapv3_factory  
**Function:** quickswapv3_factory.setBaseFeeConfiguration  

Call the setBaseFeeConfiguration function on quickswapv3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| alpha1 | decimal | alpha1 parameter of type uint16 |
| alpha2 | decimal | alpha2 parameter of type uint16 |
| beta1 | decimal | beta1 parameter of type uint32 |
| beta2 | decimal | beta2 parameter of type uint32 |
| gamma1 | decimal | gamma1 parameter of type uint16 |
| gamma2 | decimal | gamma2 parameter of type uint16 |
| volumeBeta | decimal | volumeBeta parameter of type uint32 |
| volumeGamma | decimal | volumeGamma parameter of type uint16 |
| baseFee | decimal | baseFee parameter of type uint16 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_factory.setBaseFeeConfiguration",
  "effect": "ALLOW",
  "constraints": {
    "alpha1": {
      "type": "fixed",
      "value": "example_value"
    },
    "alpha2": {
      "type": "fixed",
      "value": "example_value"
    },
    "beta1": {
      "type": "fixed",
      "value": "example_value"
    },
    "beta2": {
      "type": "fixed",
      "value": "example_value"
    },
    "gamma1": {
      "type": "fixed",
      "value": "example_value"
    },
    "gamma2": {
      "type": "fixed",
      "value": "example_value"
    },
    "volumeBeta": {
      "type": "fixed",
      "value": "example_value"
    },
    "volumeGamma": {
      "type": "fixed",
      "value": "example_value"
    },
    "baseFee": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_factory.setFarmingAddress

**Chain:** Ethereum  
**Protocol:** quickswapv3_factory  
**Function:** quickswapv3_factory.setFarmingAddress  

Call the setFarmingAddress function on quickswapv3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _farmingAddress | address | _farmingAddress parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_factory.setFarmingAddress",
  "effect": "ALLOW",
  "constraints": {
    "_farmingAddress": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_factory.setOwner

**Chain:** Ethereum  
**Protocol:** quickswapv3_factory  
**Function:** quickswapv3_factory.setOwner  

Call the setOwner function on quickswapv3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _owner | address | _owner parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_factory.setOwner",
  "effect": "ALLOW",
  "constraints": {
    "_owner": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_factory.setVaultAddress

**Chain:** Ethereum  
**Protocol:** quickswapv3_factory  
**Function:** quickswapv3_factory.setVaultAddress  

Call the setVaultAddress function on quickswapv3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _vaultAddress | address | _vaultAddress parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_factory.setVaultAddress",
  "effect": "ALLOW",
  "constraints": {
    "_vaultAddress": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_factory.vaultAddress

**Chain:** Ethereum  
**Protocol:** quickswapv3_factory  
**Function:** quickswapv3_factory.vaultAddress  

Call the vaultAddress function on quickswapv3_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_factory.vaultAddress",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.quickswapv3_router.WNativeToken

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.WNativeToken  

Call the WNativeToken function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.WNativeToken",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.quickswapv3_router.algebraSwapCallback

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.algebraSwapCallback  

Call the algebraSwapCallback function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amount0Delta | decimal | amount0Delta parameter of type int256 |
| amount1Delta | decimal | amount1Delta parameter of type int256 |
| _data | bytes | _data parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.algebraSwapCallback",
  "effect": "ALLOW",
  "constraints": {
    "amount0Delta": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount1Delta": {
      "type": "fixed",
      "value": "example_value"
    },
    "_data": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_router.exactInput

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.exactInput  

Call the exactInput function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct ISwapRouter.ExactInputParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.exactInput",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.quickswapv3_router.exactInputSingle

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.exactInputSingle  

Call the exactInputSingle function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct ISwapRouter.ExactInputSingleParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.exactInputSingle",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.quickswapv3_router.exactInputSingleSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.exactInputSingleSupportingFeeOnTransferTokens  

Call the exactInputSingleSupportingFeeOnTransferTokens function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct ISwapRouter.ExactInputSingleParams |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.exactInputSingleSupportingFeeOnTransferTokens",
  "effect": "ALLOW",
  "constraints": {
    "params": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_router.exactOutput

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.exactOutput  

Call the exactOutput function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct ISwapRouter.ExactOutputParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.exactOutput",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.quickswapv3_router.exactOutputSingle

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.exactOutputSingle  

Call the exactOutputSingle function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct ISwapRouter.ExactOutputSingleParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.exactOutputSingle",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.quickswapv3_router.factory

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.factory  

Call the factory function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.factory",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.quickswapv3_router.multicall

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.multicall  

Call the multicall function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| data | array | data parameter of type bytes[] |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.multicall",
  "effect": "ALLOW",
  "constraints": {
    "data": {
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


### ethereum.quickswapv3_router.poolDeployer

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.poolDeployer  

Call the poolDeployer function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.poolDeployer",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.quickswapv3_router.refundNativeToken

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.refundNativeToken  

Call the refundNativeToken function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.refundNativeToken",
  "effect": "ALLOW",
  "constraints": {
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_router.selfPermit

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.selfPermit  

Call the selfPermit function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| value | decimal | value parameter of type uint256 |
| deadline | decimal | deadline parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.selfPermit",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "value": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_router.selfPermitAllowed

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.selfPermitAllowed  

Call the selfPermitAllowed function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| nonce | decimal | nonce parameter of type uint256 |
| expiry | decimal | expiry parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.selfPermitAllowed",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "nonce": {
      "type": "fixed",
      "value": "example_value"
    },
    "expiry": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_router.selfPermitAllowedIfNecessary

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.selfPermitAllowedIfNecessary  

Call the selfPermitAllowedIfNecessary function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| nonce | decimal | nonce parameter of type uint256 |
| expiry | decimal | expiry parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.selfPermitAllowedIfNecessary",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "nonce": {
      "type": "fixed",
      "value": "example_value"
    },
    "expiry": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_router.selfPermitIfNecessary

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.selfPermitIfNecessary  

Call the selfPermitIfNecessary function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| value | decimal | value parameter of type uint256 |
| deadline | decimal | deadline parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.selfPermitIfNecessary",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "value": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.quickswapv3_router.sweepToken

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.sweepToken  

Call the sweepToken function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.sweepToken",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
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


### ethereum.quickswapv3_router.sweepTokenWithFee

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.sweepTokenWithFee  

Call the sweepTokenWithFee function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| feeBips | decimal | feeBips parameter of type uint256 |
| feeRecipient | address | feeRecipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.sweepTokenWithFee",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeBips": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeRecipient": {
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


### ethereum.quickswapv3_router.unwrapWNativeToken

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.unwrapWNativeToken  

Call the unwrapWNativeToken function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.unwrapWNativeToken",
  "effect": "ALLOW",
  "constraints": {
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
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


### ethereum.quickswapv3_router.unwrapWNativeTokenWithFee

**Chain:** Ethereum  
**Protocol:** quickswapv3_router  
**Function:** quickswapv3_router.unwrapWNativeTokenWithFee  

Call the unwrapWNativeTokenWithFee function on quickswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| feeBips | decimal | feeBips parameter of type uint256 |
| feeRecipient | address | feeRecipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.quickswapv3_router.unwrapWNativeTokenWithFee",
  "effect": "ALLOW",
  "constraints": {
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeBips": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeRecipient": {
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


### ethereum.thorchain_router.RUNE

**Chain:** Ethereum  
**Protocol:** thorchain_router  
**Function:** thorchain_router.RUNE  

Call the RUNE function on thorchain_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.thorchain_router.RUNE",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.thorchain_router.deposit

**Chain:** Ethereum  
**Protocol:** thorchain_router  
**Function:** thorchain_router.deposit  

Call the deposit function on thorchain_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| vault | address | vault parameter of type address payable |
| asset | address | asset parameter of type address |
| amount | decimal | amount parameter of type uint256 |
| memo | string | memo parameter of type string |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.thorchain_router.deposit",
  "effect": "ALLOW",
  "constraints": {
    "vault": {
      "type": "fixed",
      "value": "example_value"
    },
    "asset": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },
    "memo": {
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


### ethereum.thorchain_router.depositWithExpiry

**Chain:** Ethereum  
**Protocol:** thorchain_router  
**Function:** thorchain_router.depositWithExpiry  

Call the depositWithExpiry function on thorchain_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| vault | address | vault parameter of type address payable |
| asset | address | asset parameter of type address |
| amount | decimal | amount parameter of type uint256 |
| memo | string | memo parameter of type string |
| expiration | decimal | expiration parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.thorchain_router.depositWithExpiry",
  "effect": "ALLOW",
  "constraints": {
    "vault": {
      "type": "fixed",
      "value": "example_value"
    },
    "asset": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },
    "memo": {
      "type": "fixed",
      "value": "example_value"
    },
    "expiration": {
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


### ethereum.thorchain_router.returnVaultAssets

**Chain:** Ethereum  
**Protocol:** thorchain_router  
**Function:** thorchain_router.returnVaultAssets  

Call the returnVaultAssets function on thorchain_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| router | address | router parameter of type address |
| asgard | address | asgard parameter of type address payable |
| coins | array | coins parameter of type struct THORChain_Router.Coin[] |
| memo | string | memo parameter of type string |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.thorchain_router.returnVaultAssets",
  "effect": "ALLOW",
  "constraints": {
    "router": {
      "type": "fixed",
      "value": "example_value"
    },
    "asgard": {
      "type": "fixed",
      "value": "example_value"
    },
    "coins": {
      "type": "fixed",
      "value": "example_value"
    },
    "memo": {
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


### ethereum.thorchain_router.transferAllowance

**Chain:** Ethereum  
**Protocol:** thorchain_router  
**Function:** thorchain_router.transferAllowance  

Call the transferAllowance function on thorchain_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| router | address | router parameter of type address |
| newVault | address | newVault parameter of type address |
| asset | address | asset parameter of type address |
| amount | decimal | amount parameter of type uint256 |
| memo | string | memo parameter of type string |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.thorchain_router.transferAllowance",
  "effect": "ALLOW",
  "constraints": {
    "router": {
      "type": "fixed",
      "value": "example_value"
    },
    "newVault": {
      "type": "fixed",
      "value": "example_value"
    },
    "asset": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },
    "memo": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.thorchain_router.transferOut

**Chain:** Ethereum  
**Protocol:** thorchain_router  
**Function:** thorchain_router.transferOut  

Call the transferOut function on thorchain_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| to | address | to parameter of type address payable |
| asset | address | asset parameter of type address |
| amount | decimal | amount parameter of type uint256 |
| memo | string | memo parameter of type string |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.thorchain_router.transferOut",
  "effect": "ALLOW",
  "constraints": {
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "asset": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },
    "memo": {
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


### ethereum.thorchain_router.transferOutAndCall

**Chain:** Ethereum  
**Protocol:** thorchain_router  
**Function:** thorchain_router.transferOutAndCall  

Call the transferOutAndCall function on thorchain_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| aggregator | address | aggregator parameter of type address payable |
| finalToken | address | finalToken parameter of type address |
| to | address | to parameter of type address |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| memo | string | memo parameter of type string |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.thorchain_router.transferOutAndCall",
  "effect": "ALLOW",
  "constraints": {
    "aggregator": {
      "type": "fixed",
      "value": "example_value"
    },
    "finalToken": {
      "type": "fixed",
      "value": "example_value"
    },
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountOutMin": {
      "type": "fixed",
      "value": "example_value"
    },
    "memo": {
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


### ethereum.thorchain_router.vaultAllowance

**Chain:** Ethereum  
**Protocol:** thorchain_router  
**Function:** thorchain_router.vaultAllowance  

Call the vaultAllowance function on thorchain_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| vault | address | vault parameter of type address |
| token | address | token parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.thorchain_router.vaultAllowance",
  "effect": "ALLOW",
  "constraints": {
    "vault": {
      "type": "fixed",
      "value": "example_value"
    },
    "token": {
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


### ethereum.uniswapv3_router.WETH9

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.WETH9  

Call the WETH9 function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.WETH9",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.uniswapv3_router.approveMax

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.approveMax  

Call the approveMax function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.approveMax",
  "effect": "ALLOW",
  "constraints": {
    "token": {
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


### ethereum.uniswapv3_router.approveMaxMinusOne

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.approveMaxMinusOne  

Call the approveMaxMinusOne function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.approveMaxMinusOne",
  "effect": "ALLOW",
  "constraints": {
    "token": {
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


### ethereum.uniswapv3_router.approveZeroThenMax

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.approveZeroThenMax  

Call the approveZeroThenMax function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.approveZeroThenMax",
  "effect": "ALLOW",
  "constraints": {
    "token": {
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


### ethereum.uniswapv3_router.approveZeroThenMaxMinusOne

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.approveZeroThenMaxMinusOne  

Call the approveZeroThenMaxMinusOne function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.approveZeroThenMaxMinusOne",
  "effect": "ALLOW",
  "constraints": {
    "token": {
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


### ethereum.uniswapv3_router.callPositionManager

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.callPositionManager  

Call the callPositionManager function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| data | bytes | data parameter of type bytes |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.callPositionManager",
  "effect": "ALLOW",
  "constraints": {
    "data": {
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


### ethereum.uniswapv3_router.checkOracleSlippage

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.checkOracleSlippage  

Call the checkOracleSlippage function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| path | bytes | path parameter of type bytes |
| maximumTickDivergence | decimal | maximumTickDivergence parameter of type uint24 |
| secondsAgo | decimal | secondsAgo parameter of type uint32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.checkOracleSlippage",
  "effect": "ALLOW",
  "constraints": {
    "path": {
      "type": "fixed",
      "value": "example_value"
    },
    "maximumTickDivergence": {
      "type": "fixed",
      "value": "example_value"
    },
    "secondsAgo": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv3_router.checkOracleSlippage

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.checkOracleSlippage  

Call the checkOracleSlippage function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| paths | array | paths parameter of type bytes[] |
| amounts | decimal | amounts parameter of type uint128[] |
| maximumTickDivergence | decimal | maximumTickDivergence parameter of type uint24 |
| secondsAgo | decimal | secondsAgo parameter of type uint32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.checkOracleSlippage",
  "effect": "ALLOW",
  "constraints": {
    "paths": {
      "type": "fixed",
      "value": "example_value"
    },
    "amounts": {
      "type": "fixed",
      "value": "example_value"
    },
    "maximumTickDivergence": {
      "type": "fixed",
      "value": "example_value"
    },
    "secondsAgo": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv3_router.exactInput

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.exactInput  

Call the exactInput function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct IV3SwapRouter.ExactInputParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.exactInput",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.uniswapv3_router.exactInputSingle

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.exactInputSingle  

Call the exactInputSingle function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct IV3SwapRouter.ExactInputSingleParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.exactInputSingle",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.uniswapv3_router.exactOutput

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.exactOutput  

Call the exactOutput function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct IV3SwapRouter.ExactOutputParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.exactOutput",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.uniswapv3_router.exactOutputSingle

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.exactOutputSingle  

Call the exactOutputSingle function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct IV3SwapRouter.ExactOutputSingleParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.exactOutputSingle",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.uniswapv3_router.factory

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.factory  

Call the factory function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.factory",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.uniswapv3_router.factoryV2

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.factoryV2  

Call the factoryV2 function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.factoryV2",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.uniswapv3_router.getApprovalType

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.getApprovalType  

Call the getApprovalType function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amount | decimal | amount parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.getApprovalType",
  "effect": "ALLOW",
  "constraints": {
    "token": {
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


### ethereum.uniswapv3_router.increaseLiquidity

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.increaseLiquidity  

Call the increaseLiquidity function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct IApproveAndCall.IncreaseLiquidityParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.increaseLiquidity",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.uniswapv3_router.mint

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.mint  

Call the mint function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| params | string | params parameter of type struct IApproveAndCall.MintParams |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.mint",
  "effect": "ALLOW",
  "constraints": {
    "params": {
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


### ethereum.uniswapv3_router.multicall

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.multicall  

Call the multicall function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| previousBlockhash | string | previousBlockhash parameter of type bytes32 |
| data | array | data parameter of type bytes[] |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.multicall",
  "effect": "ALLOW",
  "constraints": {
    "previousBlockhash": {
      "type": "fixed",
      "value": "example_value"
    },
    "data": {
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


### ethereum.uniswapv3_router.multicall

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.multicall  

Call the multicall function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| deadline | decimal | deadline parameter of type uint256 |
| data | array | data parameter of type bytes[] |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.multicall",
  "effect": "ALLOW",
  "constraints": {
    "deadline": {
      "type": "fixed",
      "value": "example_value"
    },
    "data": {
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


### ethereum.uniswapv3_router.multicall

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.multicall  

Call the multicall function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| data | array | data parameter of type bytes[] |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.multicall",
  "effect": "ALLOW",
  "constraints": {
    "data": {
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


### ethereum.uniswapv3_router.positionManager

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.positionManager  

Call the positionManager function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.positionManager",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.uniswapv3_router.pull

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.pull  

Call the pull function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| value | decimal | value parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.pull",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "value": {
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


### ethereum.uniswapv3_router.refundETH

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.refundETH  

Call the refundETH function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.refundETH",
  "effect": "ALLOW",
  "constraints": {
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv3_router.selfPermit

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.selfPermit  

Call the selfPermit function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| value | decimal | value parameter of type uint256 |
| deadline | decimal | deadline parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.selfPermit",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "value": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv3_router.selfPermitAllowed

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.selfPermitAllowed  

Call the selfPermitAllowed function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| nonce | decimal | nonce parameter of type uint256 |
| expiry | decimal | expiry parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.selfPermitAllowed",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "nonce": {
      "type": "fixed",
      "value": "example_value"
    },
    "expiry": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv3_router.selfPermitAllowedIfNecessary

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.selfPermitAllowedIfNecessary  

Call the selfPermitAllowedIfNecessary function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| nonce | decimal | nonce parameter of type uint256 |
| expiry | decimal | expiry parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.selfPermitAllowedIfNecessary",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "nonce": {
      "type": "fixed",
      "value": "example_value"
    },
    "expiry": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv3_router.selfPermitIfNecessary

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.selfPermitIfNecessary  

Call the selfPermitIfNecessary function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| value | decimal | value parameter of type uint256 |
| deadline | decimal | deadline parameter of type uint256 |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.selfPermitIfNecessary",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "value": {
      "type": "fixed",
      "value": "example_value"
    },
    "deadline": {
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv3_router.swapExactTokensForTokens

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.swapExactTokensForTokens  

Call the swapExactTokensForTokens function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.swapExactTokensForTokens",
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv3_router.swapTokensForExactTokens

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.swapTokensForExactTokens  

Call the swapTokensForExactTokens function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 |
| amountInMax | decimal | amountInMax parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.swapTokensForExactTokens",
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
    "value": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv3_router.sweepToken

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.sweepToken  

Call the sweepToken function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.sweepToken",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
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


### ethereum.uniswapv3_router.sweepToken

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.sweepToken  

Call the sweepToken function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.sweepToken",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountMinimum": {
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


### ethereum.uniswapv3_router.sweepTokenWithFee

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.sweepTokenWithFee  

Call the sweepTokenWithFee function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| feeBips | decimal | feeBips parameter of type uint256 |
| feeRecipient | address | feeRecipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.sweepTokenWithFee",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeBips": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeRecipient": {
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


### ethereum.uniswapv3_router.sweepTokenWithFee

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.sweepTokenWithFee  

Call the sweepTokenWithFee function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| feeBips | decimal | feeBips parameter of type uint256 |
| feeRecipient | address | feeRecipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.sweepTokenWithFee",
  "effect": "ALLOW",
  "constraints": {
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeBips": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeRecipient": {
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


### ethereum.uniswapv3_router.uniswapV3SwapCallback

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.uniswapV3SwapCallback  

Call the uniswapV3SwapCallback function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amount0Delta | decimal | amount0Delta parameter of type int256 |
| amount1Delta | decimal | amount1Delta parameter of type int256 |
| _data | bytes | _data parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.uniswapV3SwapCallback",
  "effect": "ALLOW",
  "constraints": {
    "amount0Delta": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount1Delta": {
      "type": "fixed",
      "value": "example_value"
    },
    "_data": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.uniswapv3_router.unwrapWETH9

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.unwrapWETH9  

Call the unwrapWETH9 function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.unwrapWETH9",
  "effect": "ALLOW",
  "constraints": {
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
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


### ethereum.uniswapv3_router.unwrapWETH9

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.unwrapWETH9  

Call the unwrapWETH9 function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.unwrapWETH9",
  "effect": "ALLOW",
  "constraints": {
    "amountMinimum": {
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


### ethereum.uniswapv3_router.unwrapWETH9WithFee

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.unwrapWETH9WithFee  

Call the unwrapWETH9WithFee function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| recipient | address | recipient parameter of type address |
| feeBips | decimal | feeBips parameter of type uint256 |
| feeRecipient | address | feeRecipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.unwrapWETH9WithFee",
  "effect": "ALLOW",
  "constraints": {
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeBips": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeRecipient": {
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


### ethereum.uniswapv3_router.unwrapWETH9WithFee

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.unwrapWETH9WithFee  

Call the unwrapWETH9WithFee function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountMinimum | decimal | amountMinimum parameter of type uint256 |
| feeBips | decimal | feeBips parameter of type uint256 |
| feeRecipient | address | feeRecipient parameter of type address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.unwrapWETH9WithFee",
  "effect": "ALLOW",
  "constraints": {
    "amountMinimum": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeBips": {
      "type": "fixed",
      "value": "example_value"
    },
    "feeRecipient": {
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


### ethereum.uniswapv3_router.wrapETH

**Chain:** Ethereum  
**Protocol:** uniswapv3_router  
**Function:** uniswapv3_router.wrapETH  

Call the wrapETH function on uniswapv3_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| value | decimal | value parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.uniswapv3_router.wrapETH",
  "effect": "ALLOW",
  "constraints": {
    "value": {
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


### ethereum.vvs_factory.INIT_CODE_PAIR_HASH

**Chain:** Ethereum  
**Protocol:** vvs_factory  
**Function:** vvs_factory.INIT_CODE_PAIR_HASH  

Call the INIT_CODE_PAIR_HASH function on vvs_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_factory.INIT_CODE_PAIR_HASH",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.vvs_factory.allPairs

**Chain:** Ethereum  
**Protocol:** vvs_factory  
**Function:** vvs_factory.allPairs  

Call the allPairs function on vvs_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| param0 | decimal |  parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_factory.allPairs",
  "effect": "ALLOW",
  "constraints": {
    "param0": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.vvs_factory.allPairsLength

**Chain:** Ethereum  
**Protocol:** vvs_factory  
**Function:** vvs_factory.allPairsLength  

Call the allPairsLength function on vvs_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_factory.allPairsLength",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.vvs_factory.createPair

**Chain:** Ethereum  
**Protocol:** vvs_factory  
**Function:** vvs_factory.createPair  

Call the createPair function on vvs_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_factory.createPair",
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

  }
}
```


### ethereum.vvs_factory.feeTo

**Chain:** Ethereum  
**Protocol:** vvs_factory  
**Function:** vvs_factory.feeTo  

Call the feeTo function on vvs_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_factory.feeTo",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.vvs_factory.feeToSetter

**Chain:** Ethereum  
**Protocol:** vvs_factory  
**Function:** vvs_factory.feeToSetter  

Call the feeToSetter function on vvs_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_factory.feeToSetter",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.vvs_factory.getPair

**Chain:** Ethereum  
**Protocol:** vvs_factory  
**Function:** vvs_factory.getPair  

Call the getPair function on vvs_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| param0 | address |  parameter of type address |
| param1 | address |  parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_factory.getPair",
  "effect": "ALLOW",
  "constraints": {
    "param0": {
      "type": "fixed",
      "value": "example_value"
    },
    "param1": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.vvs_factory.setFeeTo

**Chain:** Ethereum  
**Protocol:** vvs_factory  
**Function:** vvs_factory.setFeeTo  

Call the setFeeTo function on vvs_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _feeTo | address | _feeTo parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_factory.setFeeTo",
  "effect": "ALLOW",
  "constraints": {
    "_feeTo": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.vvs_factory.setFeeToSetter

**Chain:** Ethereum  
**Protocol:** vvs_factory  
**Function:** vvs_factory.setFeeToSetter  

Call the setFeeToSetter function on vvs_factory

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| _feeToSetter | address | _feeToSetter parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_factory.setFeeToSetter",
  "effect": "ALLOW",
  "constraints": {
    "_feeToSetter": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.vvs_router.WETH

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.WETH  

Call the WETH function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.WETH",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.vvs_router.addLiquidity

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.addLiquidity  

Call the addLiquidity function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| amountADesired | decimal | amountADesired parameter of type uint256 |
| amountBDesired | decimal | amountBDesired parameter of type uint256 |
| amountAMin | decimal | amountAMin parameter of type uint256 |
| amountBMin | decimal | amountBMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.addLiquidity",
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


### ethereum.vvs_router.addLiquidityETH

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.addLiquidityETH  

Call the addLiquidityETH function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| amountTokenDesired | decimal | amountTokenDesired parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.addLiquidityETH",
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


### ethereum.vvs_router.factory

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.factory  

Call the factory function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.factory",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.vvs_router.getAmountIn

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.getAmountIn  

Call the getAmountIn function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 |
| reserveIn | decimal | reserveIn parameter of type uint256 |
| reserveOut | decimal | reserveOut parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.getAmountIn",
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


### ethereum.vvs_router.getAmountOut

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.getAmountOut  

Call the getAmountOut function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| reserveIn | decimal | reserveIn parameter of type uint256 |
| reserveOut | decimal | reserveOut parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.getAmountOut",
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


### ethereum.vvs_router.getAmountsIn

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.getAmountsIn  

Call the getAmountsIn function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 |
| path | array | path parameter of type address[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.getAmountsIn",
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


### ethereum.vvs_router.getAmountsOut

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.getAmountsOut  

Call the getAmountsOut function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| path | array | path parameter of type address[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.getAmountsOut",
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


### ethereum.vvs_router.quote

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.quote  

Call the quote function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountA | decimal | amountA parameter of type uint256 |
| reserveA | decimal | reserveA parameter of type uint256 |
| reserveB | decimal | reserveB parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.quote",
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


### ethereum.vvs_router.removeLiquidity

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.removeLiquidity  

Call the removeLiquidity function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountAMin | decimal | amountAMin parameter of type uint256 |
| amountBMin | decimal | amountBMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.removeLiquidity",
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


### ethereum.vvs_router.removeLiquidityETH

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.removeLiquidityETH  

Call the removeLiquidityETH function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.removeLiquidityETH",
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


### ethereum.vvs_router.removeLiquidityETHSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.removeLiquidityETHSupportingFeeOnTransferTokens  

Call the removeLiquidityETHSupportingFeeOnTransferTokens function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.removeLiquidityETHSupportingFeeOnTransferTokens",
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


### ethereum.vvs_router.removeLiquidityETHWithPermit

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.removeLiquidityETHWithPermit  

Call the removeLiquidityETHWithPermit function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| approveMax | boolean | approveMax parameter of type bool |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.removeLiquidityETHWithPermit",
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


### ethereum.vvs_router.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens  

Call the removeLiquidityETHWithPermitSupportingFeeOnTransferTokens function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountTokenMin | decimal | amountTokenMin parameter of type uint256 |
| amountETHMin | decimal | amountETHMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| approveMax | boolean | approveMax parameter of type bool |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens",
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


### ethereum.vvs_router.removeLiquidityWithPermit

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.removeLiquidityWithPermit  

Call the removeLiquidityWithPermit function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| tokenA | address | tokenA parameter of type address |
| tokenB | address | tokenB parameter of type address |
| liquidity | decimal | liquidity parameter of type uint256 |
| amountAMin | decimal | amountAMin parameter of type uint256 |
| amountBMin | decimal | amountBMin parameter of type uint256 |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| approveMax | boolean | approveMax parameter of type bool |
| v | decimal | v parameter of type uint8 |
| r | string | r parameter of type bytes32 |
| s | string | s parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.removeLiquidityWithPermit",
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


### ethereum.vvs_router.swapETHForExactTokens

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.swapETHForExactTokens  

Call the swapETHForExactTokens function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.swapETHForExactTokens",
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


### ethereum.vvs_router.swapExactETHForTokens

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.swapExactETHForTokens  

Call the swapExactETHForTokens function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.swapExactETHForTokens",
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


### ethereum.vvs_router.swapExactETHForTokensSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.swapExactETHForTokensSupportingFeeOnTransferTokens  

Call the swapExactETHForTokensSupportingFeeOnTransferTokens function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.swapExactETHForTokensSupportingFeeOnTransferTokens",
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


### ethereum.vvs_router.swapExactTokensForETH

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.swapExactTokensForETH  

Call the swapExactTokensForETH function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.swapExactTokensForETH",
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


### ethereum.vvs_router.swapExactTokensForETHSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.swapExactTokensForETHSupportingFeeOnTransferTokens  

Call the swapExactTokensForETHSupportingFeeOnTransferTokens function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.swapExactTokensForETHSupportingFeeOnTransferTokens",
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


### ethereum.vvs_router.swapExactTokensForTokens

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.swapExactTokensForTokens  

Call the swapExactTokensForTokens function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.swapExactTokensForTokens",
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


### ethereum.vvs_router.swapExactTokensForTokensSupportingFeeOnTransferTokens

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.swapExactTokensForTokensSupportingFeeOnTransferTokens  

Call the swapExactTokensForTokensSupportingFeeOnTransferTokens function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountIn | decimal | amountIn parameter of type uint256 |
| amountOutMin | decimal | amountOutMin parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.swapExactTokensForTokensSupportingFeeOnTransferTokens",
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


### ethereum.vvs_router.swapTokensForExactETH

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.swapTokensForExactETH  

Call the swapTokensForExactETH function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 |
| amountInMax | decimal | amountInMax parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.swapTokensForExactETH",
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


### ethereum.vvs_router.swapTokensForExactTokens

**Chain:** Ethereum  
**Protocol:** vvs_router  
**Function:** vvs_router.swapTokensForExactTokens  

Call the swapTokensForExactTokens function on vvs_router

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amountOut | decimal | amountOut parameter of type uint256 |
| amountInMax | decimal | amountInMax parameter of type uint256 |
| path | array | path parameter of type address[] |
| to | address | to parameter of type address |
| deadline | decimal | deadline parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.vvs_router.swapTokensForExactTokens",
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
