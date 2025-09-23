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


### ethereum.routerv6_1inch.advanceEpoch

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.advanceEpoch  

Call the advanceEpoch function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| series | decimal | series parameter of type uint96 |
| amount | decimal | amount parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.advanceEpoch",
  "effect": "ALLOW",
  "constraints": {
    "series": {
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


### ethereum.routerv6_1inch.and

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.and  

Call the and function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| offsets | decimal | offsets parameter of type uint256 |
| data | bytes | data parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.and",
  "effect": "ALLOW",
  "constraints": {
    "offsets": {
      "type": "fixed",
      "value": "example_value"
    },
    "data": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.arbitraryStaticCall

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.arbitraryStaticCall  

Call the arbitraryStaticCall function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| target | address | target parameter of type address |
| data | bytes | data parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.arbitraryStaticCall",
  "effect": "ALLOW",
  "constraints": {
    "target": {
      "type": "fixed",
      "value": "example_value"
    },
    "data": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.bitInvalidatorForOrder

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.bitInvalidatorForOrder  

Call the bitInvalidatorForOrder function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| maker | address | maker parameter of type address |
| slot | decimal | slot parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.bitInvalidatorForOrder",
  "effect": "ALLOW",
  "constraints": {
    "maker": {
      "type": "fixed",
      "value": "example_value"
    },
    "slot": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.bitsInvalidateForOrder

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.bitsInvalidateForOrder  

Call the bitsInvalidateForOrder function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| makerTraits | decimal | makerTraits parameter of type MakerTraits |
| additionalMask | decimal | additionalMask parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.bitsInvalidateForOrder",
  "effect": "ALLOW",
  "constraints": {
    "makerTraits": {
      "type": "fixed",
      "value": "example_value"
    },
    "additionalMask": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.cancelOrder

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.cancelOrder  

Call the cancelOrder function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| makerTraits | decimal | makerTraits parameter of type MakerTraits |
| orderHash | string | orderHash parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.cancelOrder",
  "effect": "ALLOW",
  "constraints": {
    "makerTraits": {
      "type": "fixed",
      "value": "example_value"
    },
    "orderHash": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.cancelOrders

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.cancelOrders  

Call the cancelOrders function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| makerTraits | decimal | makerTraits parameter of type MakerTraits[] |
| orderHashes | array | orderHashes parameter of type bytes32[] |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.cancelOrders",
  "effect": "ALLOW",
  "constraints": {
    "makerTraits": {
      "type": "fixed",
      "value": "example_value"
    },
    "orderHashes": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.checkPredicate

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.checkPredicate  

Call the checkPredicate function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| predicate | bytes | predicate parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.checkPredicate",
  "effect": "ALLOW",
  "constraints": {
    "predicate": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.clipperSwap

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.clipperSwap  

Call the clipperSwap function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| clipperExchange | address | clipperExchange parameter of type contract IClipperExchange |
| srcToken | decimal | srcToken parameter of type Address |
| dstToken | address | dstToken parameter of type contract IERC20 |
| inputAmount | decimal | inputAmount parameter of type uint256 |
| outputAmount | decimal | outputAmount parameter of type uint256 |
| goodUntil | decimal | goodUntil parameter of type uint256 |
| r | string | r parameter of type bytes32 |
| vs | string | vs parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.clipperSwap",
  "effect": "ALLOW",
  "constraints": {
    "clipperExchange": {
      "type": "fixed",
      "value": "example_value"
    },
    "srcToken": {
      "type": "fixed",
      "value": "example_value"
    },
    "dstToken": {
      "type": "fixed",
      "value": "example_value"
    },
    "inputAmount": {
      "type": "fixed",
      "value": "example_value"
    },
    "outputAmount": {
      "type": "fixed",
      "value": "example_value"
    },
    "goodUntil": {
      "type": "fixed",
      "value": "example_value"
    },
    "r": {
      "type": "fixed",
      "value": "example_value"
    },
    "vs": {
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


### ethereum.routerv6_1inch.clipperSwapTo

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.clipperSwapTo  

Call the clipperSwapTo function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| clipperExchange | address | clipperExchange parameter of type contract IClipperExchange |
| recipient | address | recipient parameter of type address payable |
| srcToken | decimal | srcToken parameter of type Address |
| dstToken | address | dstToken parameter of type contract IERC20 |
| inputAmount | decimal | inputAmount parameter of type uint256 |
| outputAmount | decimal | outputAmount parameter of type uint256 |
| goodUntil | decimal | goodUntil parameter of type uint256 |
| r | string | r parameter of type bytes32 |
| vs | string | vs parameter of type bytes32 |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.clipperSwapTo",
  "effect": "ALLOW",
  "constraints": {
    "clipperExchange": {
      "type": "fixed",
      "value": "example_value"
    },
    "recipient": {
      "type": "fixed",
      "value": "example_value"
    },
    "srcToken": {
      "type": "fixed",
      "value": "example_value"
    },
    "dstToken": {
      "type": "fixed",
      "value": "example_value"
    },
    "inputAmount": {
      "type": "fixed",
      "value": "example_value"
    },
    "outputAmount": {
      "type": "fixed",
      "value": "example_value"
    },
    "goodUntil": {
      "type": "fixed",
      "value": "example_value"
    },
    "r": {
      "type": "fixed",
      "value": "example_value"
    },
    "vs": {
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


### ethereum.routerv6_1inch.curveSwapCallback

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.curveSwapCallback  

Call the curveSwapCallback function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| param0 | address |  parameter of type address |
| param1 | address |  parameter of type address |
| inCoin | address | inCoin parameter of type address |
| dx | decimal | dx parameter of type uint256 |
| param4 | decimal |  parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.curveSwapCallback",
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
    "inCoin": {
      "type": "fixed",
      "value": "example_value"
    },
    "dx": {
      "type": "fixed",
      "value": "example_value"
    },
    "param4": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.eip712Domain

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.eip712Domain  

Call the eip712Domain function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.eip712Domain",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.routerv6_1inch.epoch

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.epoch  

Call the epoch function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| maker | address | maker parameter of type address |
| series | decimal | series parameter of type uint96 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.epoch",
  "effect": "ALLOW",
  "constraints": {
    "maker": {
      "type": "fixed",
      "value": "example_value"
    },
    "series": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.epochEquals

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.epochEquals  

Call the epochEquals function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| maker | address | maker parameter of type address |
| series | decimal | series parameter of type uint256 |
| makerEpoch | decimal | makerEpoch parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.epochEquals",
  "effect": "ALLOW",
  "constraints": {
    "maker": {
      "type": "fixed",
      "value": "example_value"
    },
    "series": {
      "type": "fixed",
      "value": "example_value"
    },
    "makerEpoch": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.eq

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.eq  

Call the eq function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| value | decimal | value parameter of type uint256 |
| data | bytes | data parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.eq",
  "effect": "ALLOW",
  "constraints": {
    "value": {
      "type": "fixed",
      "value": "example_value"
    },
    "data": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.ethUnoswap

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.ethUnoswap  

Call the ethUnoswap function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.ethUnoswap",
  "effect": "ALLOW",
  "constraints": {
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
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


### ethereum.routerv6_1inch.ethUnoswap2

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.ethUnoswap2  

Call the ethUnoswap2 function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |
| dex2 | decimal | dex2 parameter of type Address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.ethUnoswap2",
  "effect": "ALLOW",
  "constraints": {
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex2": {
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


### ethereum.routerv6_1inch.ethUnoswap3

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.ethUnoswap3  

Call the ethUnoswap3 function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |
| dex2 | decimal | dex2 parameter of type Address |
| dex3 | decimal | dex3 parameter of type Address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.ethUnoswap3",
  "effect": "ALLOW",
  "constraints": {
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex2": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex3": {
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


### ethereum.routerv6_1inch.ethUnoswapTo

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.ethUnoswapTo  

Call the ethUnoswapTo function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| to | decimal | to parameter of type Address |
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.ethUnoswapTo",
  "effect": "ALLOW",
  "constraints": {
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
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


### ethereum.routerv6_1inch.ethUnoswapTo2

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.ethUnoswapTo2  

Call the ethUnoswapTo2 function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| to | decimal | to parameter of type Address |
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |
| dex2 | decimal | dex2 parameter of type Address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.ethUnoswapTo2",
  "effect": "ALLOW",
  "constraints": {
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex2": {
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


### ethereum.routerv6_1inch.ethUnoswapTo3

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.ethUnoswapTo3  

Call the ethUnoswapTo3 function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| to | decimal | to parameter of type Address |
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |
| dex2 | decimal | dex2 parameter of type Address |
| dex3 | decimal | dex3 parameter of type Address |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.ethUnoswapTo3",
  "effect": "ALLOW",
  "constraints": {
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex2": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex3": {
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


### ethereum.routerv6_1inch.fillContractOrder

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.fillContractOrder  

Call the fillContractOrder function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| order | string | order parameter of type struct IOrderMixin.Order |
| signature | bytes | signature parameter of type bytes |
| amount | decimal | amount parameter of type uint256 |
| takerTraits | decimal | takerTraits parameter of type TakerTraits |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.fillContractOrder",
  "effect": "ALLOW",
  "constraints": {
    "order": {
      "type": "fixed",
      "value": "example_value"
    },
    "signature": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },
    "takerTraits": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.fillContractOrderArgs

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.fillContractOrderArgs  

Call the fillContractOrderArgs function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| order | string | order parameter of type struct IOrderMixin.Order |
| signature | bytes | signature parameter of type bytes |
| amount | decimal | amount parameter of type uint256 |
| takerTraits | decimal | takerTraits parameter of type TakerTraits |
| args | bytes | args parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.fillContractOrderArgs",
  "effect": "ALLOW",
  "constraints": {
    "order": {
      "type": "fixed",
      "value": "example_value"
    },
    "signature": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },
    "takerTraits": {
      "type": "fixed",
      "value": "example_value"
    },
    "args": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.fillOrder

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.fillOrder  

Call the fillOrder function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| order | string | order parameter of type struct IOrderMixin.Order |
| r | string | r parameter of type bytes32 |
| vs | string | vs parameter of type bytes32 |
| amount | decimal | amount parameter of type uint256 |
| takerTraits | decimal | takerTraits parameter of type TakerTraits |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.fillOrder",
  "effect": "ALLOW",
  "constraints": {
    "order": {
      "type": "fixed",
      "value": "example_value"
    },
    "r": {
      "type": "fixed",
      "value": "example_value"
    },
    "vs": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },
    "takerTraits": {
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


### ethereum.routerv6_1inch.fillOrderArgs

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.fillOrderArgs  

Call the fillOrderArgs function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| order | string | order parameter of type struct IOrderMixin.Order |
| r | string | r parameter of type bytes32 |
| vs | string | vs parameter of type bytes32 |
| amount | decimal | amount parameter of type uint256 |
| takerTraits | decimal | takerTraits parameter of type TakerTraits |
| args | bytes | args parameter of type bytes |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.fillOrderArgs",
  "effect": "ALLOW",
  "constraints": {
    "order": {
      "type": "fixed",
      "value": "example_value"
    },
    "r": {
      "type": "fixed",
      "value": "example_value"
    },
    "vs": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },
    "takerTraits": {
      "type": "fixed",
      "value": "example_value"
    },
    "args": {
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


### ethereum.routerv6_1inch.gt

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.gt  

Call the gt function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| value | decimal | value parameter of type uint256 |
| data | bytes | data parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.gt",
  "effect": "ALLOW",
  "constraints": {
    "value": {
      "type": "fixed",
      "value": "example_value"
    },
    "data": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.hashOrder

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.hashOrder  

Call the hashOrder function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| order | string | order parameter of type struct IOrderMixin.Order |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.hashOrder",
  "effect": "ALLOW",
  "constraints": {
    "order": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.increaseEpoch

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.increaseEpoch  

Call the increaseEpoch function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| series | decimal | series parameter of type uint96 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.increaseEpoch",
  "effect": "ALLOW",
  "constraints": {
    "series": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.lt

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.lt  

Call the lt function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| value | decimal | value parameter of type uint256 |
| data | bytes | data parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.lt",
  "effect": "ALLOW",
  "constraints": {
    "value": {
      "type": "fixed",
      "value": "example_value"
    },
    "data": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.not

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.not  

Call the not function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| data | bytes | data parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.not",
  "effect": "ALLOW",
  "constraints": {
    "data": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.or

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.or  

Call the or function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| offsets | decimal | offsets parameter of type uint256 |
| data | bytes | data parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.or",
  "effect": "ALLOW",
  "constraints": {
    "offsets": {
      "type": "fixed",
      "value": "example_value"
    },
    "data": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.owner

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.owner  

Call the owner function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.owner",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.routerv6_1inch.pause

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.pause  

Call the pause function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.pause",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.routerv6_1inch.paused

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.paused  

Call the paused function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.paused",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.routerv6_1inch.permitAndCall

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.permitAndCall  

Call the permitAndCall function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| permit | bytes | permit parameter of type bytes |
| action | bytes | action parameter of type bytes |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.permitAndCall",
  "effect": "ALLOW",
  "constraints": {
    "permit": {
      "type": "fixed",
      "value": "example_value"
    },
    "action": {
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


### ethereum.routerv6_1inch.rawRemainingInvalidatorForOrder

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.rawRemainingInvalidatorForOrder  

Call the rawRemainingInvalidatorForOrder function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| maker | address | maker parameter of type address |
| orderHash | string | orderHash parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.rawRemainingInvalidatorForOrder",
  "effect": "ALLOW",
  "constraints": {
    "maker": {
      "type": "fixed",
      "value": "example_value"
    },
    "orderHash": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.remainingInvalidatorForOrder

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.remainingInvalidatorForOrder  

Call the remainingInvalidatorForOrder function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| maker | address | maker parameter of type address |
| orderHash | string | orderHash parameter of type bytes32 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.remainingInvalidatorForOrder",
  "effect": "ALLOW",
  "constraints": {
    "maker": {
      "type": "fixed",
      "value": "example_value"
    },
    "orderHash": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.renounceOwnership

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.renounceOwnership  

Call the renounceOwnership function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.renounceOwnership",
  "effect": "ALLOW",
  "constraints": {

  }
}
```


### ethereum.routerv6_1inch.rescueFunds

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.rescueFunds  

Call the rescueFunds function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | address | token parameter of type contract IERC20 |
| amount | decimal | amount parameter of type uint256 |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.rescueFunds",
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


### ethereum.routerv6_1inch.simulate

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.simulate  

Call the simulate function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| target | address | target parameter of type address |
| data | bytes | data parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.simulate",
  "effect": "ALLOW",
  "constraints": {
    "target": {
      "type": "fixed",
      "value": "example_value"
    },
    "data": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.swap

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.swap  

Call the swap function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| executor | address | executor parameter of type contract IAggregationExecutor |
| desc | string | desc parameter of type struct GenericRouter.SwapDescription |
| data | bytes | data parameter of type bytes |
| value | decimal | The amount of ETH to send with the transaction |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.swap",
  "effect": "ALLOW",
  "constraints": {
    "executor": {
      "type": "fixed",
      "value": "example_value"
    },
    "desc": {
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


### ethereum.routerv6_1inch.transferOwnership

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.transferOwnership  

Call the transferOwnership function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| newOwner | address | newOwner parameter of type address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.transferOwnership",
  "effect": "ALLOW",
  "constraints": {
    "newOwner": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.uniswapV3SwapCallback

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.uniswapV3SwapCallback  

Call the uniswapV3SwapCallback function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| amount0Delta | decimal | amount0Delta parameter of type int256 |
| amount1Delta | decimal | amount1Delta parameter of type int256 |
| param2 | bytes |  parameter of type bytes |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.uniswapV3SwapCallback",
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
    "param2": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.unoswap

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.unoswap  

Call the unoswap function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | decimal | token parameter of type Address |
| amount | decimal | amount parameter of type uint256 |
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.unoswap",
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
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.unoswap2

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.unoswap2  

Call the unoswap2 function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | decimal | token parameter of type Address |
| amount | decimal | amount parameter of type uint256 |
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |
| dex2 | decimal | dex2 parameter of type Address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.unoswap2",
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
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex2": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.unoswap3

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.unoswap3  

Call the unoswap3 function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| token | decimal | token parameter of type Address |
| amount | decimal | amount parameter of type uint256 |
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |
| dex2 | decimal | dex2 parameter of type Address |
| dex3 | decimal | dex3 parameter of type Address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.unoswap3",
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
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex2": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex3": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.unoswapTo

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.unoswapTo  

Call the unoswapTo function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| to | decimal | to parameter of type Address |
| token | decimal | token parameter of type Address |
| amount | decimal | amount parameter of type uint256 |
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.unoswapTo",
  "effect": "ALLOW",
  "constraints": {
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.unoswapTo2

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.unoswapTo2  

Call the unoswapTo2 function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| to | decimal | to parameter of type Address |
| token | decimal | token parameter of type Address |
| amount | decimal | amount parameter of type uint256 |
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |
| dex2 | decimal | dex2 parameter of type Address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.unoswapTo2",
  "effect": "ALLOW",
  "constraints": {
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex2": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.unoswapTo3

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.unoswapTo3  

Call the unoswapTo3 function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|
| to | decimal | to parameter of type Address |
| token | decimal | token parameter of type Address |
| amount | decimal | amount parameter of type uint256 |
| minReturn | decimal | minReturn parameter of type uint256 |
| dex | decimal | dex parameter of type Address |
| dex2 | decimal | dex2 parameter of type Address |
| dex3 | decimal | dex3 parameter of type Address |


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.unoswapTo3",
  "effect": "ALLOW",
  "constraints": {
    "to": {
      "type": "fixed",
      "value": "example_value"
    },
    "token": {
      "type": "fixed",
      "value": "example_value"
    },
    "amount": {
      "type": "fixed",
      "value": "example_value"
    },
    "minReturn": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex2": {
      "type": "fixed",
      "value": "example_value"
    },
    "dex3": {
      "type": "fixed",
      "value": "example_value"
    },

  }
}
```


### ethereum.routerv6_1inch.unpause

**Chain:** Ethereum  
**Protocol:** routerv6_1inch  
**Function:** routerv6_1inch.unpause  

Call the unpause function on routerv6_1inch

**Parameters:**

| Name | Type | Description |
|------|------|-------------|


**Example Policy Rule:**

```json
{
  "resource": "ethereum.routerv6_1inch.unpause",
  "effect": "ALLOW",
  "constraints": {

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
