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
