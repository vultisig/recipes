# Vultisig Recipes

A collection of blockchain integration recipes and utilities for the Vultisig platform.

## Overview

This repository contains a set of tools and utilities for integrating various blockchain networks with Vultisig. It provides a standardized way to handle different blockchain protocols and their specific requirements.

## Features

- Chain Registry System
  - Centralized registry for managing blockchain integrations
  - Support for multiple blockchain networks
  - Easy addition of new blockchain implementations
  - Thread-safe operations

- Bitcoin Integration
  - Standardized Bitcoin chain implementation
  - Protocol support for BTC
  - Extensible design for future protocol additions

## Getting Started

### Prerequisites

- Go 1.24 or later
- Git
- For protobuf generation (only needed if modifying `.proto` files):
  - protoc v30.2+
  - protoc-gen-go: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - protoc-gen-go-grpc: `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
  - Or use buf (see `buf.gen.yaml`)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/vultisig/recipes.git
cd recipes
```

2. Install dependencies:
```bash
go mod download
```

### Usage

## Adding a New Chain

Adding a new blockchain requires implementing several layers. Each layer has a specific responsibility.

### Layer Overview

| Layer | Location | Purpose |
|-------|----------|---------|
| **Chain** | `chain/` | Parse transactions, define protocols |
| **Engine** | `engine/` | Validate transactions against policy rules |
| **SDK** | `sdk/` | Sign and broadcast transactions (optional) |
| **MetaRule** | `metarule/` | Transform `send`/`swap` intents into chain-specific rules |
| **Resolver** | `resolver/` | Resolve magic constants (vault addresses, etc.) |

---

### 1. Chain Layer (`chain/`)

Defines how to parse and identify transactions.

- **Location**: Nest under chain family if applicable
  - UTXO: `chain/utxo/<name>/`
  - EVM: `chain/evm/<name>/`
  - Standalone: `chain/<name>/`

- **Files needed**:
  - `chain.go` — Implements `types.Chain` interface (ID, Name, ParseTransaction)
  - `protocol.go` — Defines native token and supported functions
  - `decode.go` — Transaction deserialization logic (if complex)

- **Register** in `chain/registry.go`

---

### 2. Engine Layer (`engine/`)

Validates transactions against policy rules.

- **Location**: Mirror the chain layer structure (`engine/utxo/<name>/`)

- **Implements**: `Supports(chain)` and `Evaluate(rule, txBytes)`
  - UTXO chains can wrap the generic `engine/utxo.Engine`
  - EVM chains can use the shared EVM engine

- **Register** in `engine/registry.go`

---

### 3. SDK Layer (`sdk/`) — Optional

Handles transaction signing and broadcasting. Only needed if your chain has unique signing requirements not covered by existing implementations.

- **Location**: `sdk/<name>/`
- **Provides**: `Sign()`, `Broadcast()`, `Send()` methods

---

### 4. MetaRule Layer (`metarule/`)

Enables **Recurring Send** and **Recurring Swap** by transforming high-level intents into chain-specific rules.

- **`metarule/metarule.go`** — Add a `handle<Chain>()` function for `send` and `swap` protocols
- **`metarule/internal/thorchain/thorchain.go`** — Add ThorChain asset mapping:
  - Network constant (e.g., `zec network = "ZEC"`)
  - Add to `parseNetwork()` switch
  - Add asset to pools list (e.g., `{asset: "ZEC.ZEC"}`)
  - Add shortcode to `ShortCode()` if applicable

---

### 5. Resolver Layer (`resolver/`)

Resolves magic constants like `THORCHAIN_VAULT` to actual addresses.

- **`resolver/thorchain_common.go`** — Add chain to `getThorChainSymbol()` switch

---

### 6. vultisig-go Integration

Ensure the chain is defined in `vultisig-go/common/chain.go`:
- Add chain constant
- Add to `chainToString` map
- Implement `NativeSymbol()` for the chain

---

### Checklist

- [ ] `chain/` — Chain implementation + register
- [ ] `engine/` — Engine implementation + register  
- [ ] `sdk/` — SDK if needed
- [ ] `metarule/metarule.go` — Handler for send/swap
- [ ] `metarule/internal/thorchain/` — ThorChain asset mapping
- [ ] `resolver/thorchain_common.go` — Vault resolution
- [ ] `vultisig-go/common/chain.go` — Chain definition
- [ ] Tests for each layer
- [ ] Run `go run cmd/generator/main.go --output RESOURCES.md`

## Development

### Pre-Commit Checklist

Before committing changes, run the following commands to ensure CI will pass:

```bash
# 1. Build
go build ./...

# 2. Run tests
go test ./...

# 3. Run linter (optional, but recommended)
golangci-lint run ./...

# 4. Generate documentation (required if chain/protocol changes)
go run cmd/generator/main.go --output RESOURCES.md

# 5. Generate protobuf (only if .proto files changed, requires protoc)
go generate ./...
```

### Running Tests

Run all tests:
```bash
go test ./...
```

Run specific test suites:
```bash
go test ./validator/... # Run validator tests
go test ./chain/...    # Run chain tests
```

Run test engine with coverage:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Test Guidelines

- Write unit tests for all new validators
- Include both positive and negative test cases
- Mock external dependencies when testing
- Use test helpers from `testdata` package
- Follow table-driven test patterns

### Documentation

The project includes automatically generated documentation. To update the documentation:

```bash
go run cmd/generator/main.go --output RESOURCES.md
```

Note: The CI pipeline will check if the documentation is up to date. Make sure to run this command and commit any changes when modifying the codebase.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request


For support, please open an issue in the GitHub repository or contact the Vultisig team. 