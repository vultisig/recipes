# Rules System: MetaRules vs Direct Rules

## Overview

Vultisig's policy engine supports two rule types that govern how apps interact with blockchain networks:

- **MetaRules**: High-level, chain-agnostic abstractions automatically mapped to specific implementations.
- **Direct Rules**: Low-level, protocol-specific rules that strictly follow blockchain or contract semantics.

All app transactions must comply with these rules and be approved by the Verifier before signing.

***

## MetaRules: Protocol-Agnostic Abstractions

MetaRules summarize cross-chain actions (like send, swap) via simplified resource identifiers such as `{chain}.send`. The system expands these abstractions into concrete, protocol-specific Direct Rules tailored for each network.

### Example MetaRules

| MetaRule | Expands to (Examples)                                               | Chains Supported |
| :-- |:--------------------------------------------------------------------| :-- |
| `ethereum.send` | `ethereum.eth.transfer` (native), `ethereum.erc20.transfer` (ERC20) | Ethereum |
| `solana.send` | `solana.system.transfer`, `solana.spl_token.transfer`               | Solana |
| `bitcoin.send` | `bitcoin.btc.transfer`                                              | Bitcoin |
| `ethereum.swap` | 1inch, Uniswap V2/V3 calls                                          | Ethereum |
| `solana.swap` | Jupiter aggregator calls                                            | Solana |

### MetaRule Expansion (Go Example)

```go
func (m *MetaRule) TryFormat(resource string, constraints map[string]*types.Constraint) ([]string, error) {
	parts := strings.Split(resource, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid meta-rule format: %s", resource)
	}
	chain := parts[0]
	protocol := metaProtocol(parts[1])
	switch protocol {
	case send:
		return m.expandSendRule(chain, constraints)
	case swap:
		return m.expandSwapRule(chain, constraints)
	default:
		return nil, fmt.Errorf("unsupported meta-protocol: %s", protocol)
	}
}
```

***

## Direct Rules: Protocol-Specific Operations

Direct Rules are mapped directly to the target protocol's ABI (EVM) or IDL (Solana), providing fine-grained policy enforcement for app calls.

### Rule Structure (`protobuf`)

```protobuf
message Rule {
  string resource = 1;
  Effect effect = 2;
  string description = 3;
  map<string, Constraint> constraints = 4;
  Authorization authorization = 5;
  string id = 6;
  repeated ParameterConstraint parameter_constraints = 7;
  Target target = 13;
}
```


### Example Direct Rules

#### ERC-20 Transfer (Ethereum)

```json
{
  "resource": "ethereum.erc20.transfer",
  "effect": "EFFECT_ALLOW",
  "target": {
    "target_type": "TARGET_TYPE_ADDRESS",
    "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
  },
  "parameter_constraints": [
    {
      "parameter_name": "to",
      "constraint": {
        "type": "CONSTRAINT_TYPE_WHITELIST",
        "whitelist": ["0x742b15...a92E"],
        "required": true
      }
    }
  ]
}
```


#### Solana SPL Token Transfer

```json
{
  "resource": "solana.spl_token.transfer",
  "effect": "EFFECT_ALLOW",
  "parameter_constraints": [
    {
      "parameter_name": "amount",
      "constraint": {
        "type": "CONSTRAINT_TYPE_RANGE",
        "min_value": "1000000",
        "max_value": "100000000",
        "required": true
      }
    }
  ]
}
```


***

## Supported Blockchains \& Protocols

| Chain | Features Supported |
| :-- | :-- |
| Ethereum | ERC-20, Uniswap, 1inch, custom contracts |
| Bitcoin | Native BTC transfers |
| Solana | SPL tokens, Jupiter aggregator, Metaplex |
| Polygon | Polygon bridges, DEXs |
| ARB, BSC etc | Layer2 \& EVM-specific features |
| THORChain | Cross-chain liquidity/swaps |
| XRPL | XRP Ledger native operations |


***

## Constraint Types

Based on `constraint.proto`:


| Type | Description                                               |
| :-- |:----------------------------------------------------------|
| CONSTRAINT_TYPE_FIXED | Enforce fixed value for parameter                         |
| CONSTRAINT_TYPE_MAX | Maximum allowed value                                     |
| CONSTRAINT_TYPE_MIN | Minimum allowed value                                     |
| CONSTRAINT_TYPE_MAGIC_CONSTANT | Special constant (system address, treasury, router, etc.) |
| CONSTRAINT_TYPE_ANY | Accept any value                                          |
| CONSTRAINT_TYPE_REGEXP | Regular expression match for string params                |
| CONSTRAINT_TYPE_UNSPECIFIED | Not specified - usually treated as deny                   |

### Magic Constants

| Name | Value | Purpose                    |
| :-- | :-- |:---------------------------|
| VULTISIG_TREASURY | 1 | Treasury address           |
| THORCHAIN_VAULT | 2 | Router/vault for THORChain |
| THORCHAIN_ROUTER | 3 | Router for THORChain swaps |

***

## Rule Development Guidelines

- **MetaRules**: Use for simple, cross-chain actions or consistent business logic.
- **Direct Rules**: Apply for advanced protocol features, fine-grained control, or where ABI details matter.
- **Constraints**: Use principle of least privilege; always review all allowed parameters and document constraints.
- **Testing**: Validate with protocol ABIs/IDLs and ensure all value constraints are respected before allow.
- **Magic Constants**: Use system-defined constants for seamless app integrations and parameter substitutions.

This markdown integrates both architectural policy logic and concrete constraint semantics, staying consistent with the system's design and codebase.

