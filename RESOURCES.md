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
