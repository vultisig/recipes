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

#### Adding a New Chain

1. Implement the `Chain` interface:
```go
type MyChain struct{}

func (c *MyChain) ID() string {
    return "my-chain"
}

func (c *MyChain) Name() string {
    return "My Chain"
}

func (c *MyChain) SupportedProtocols() []string {
    return []string{"my-protocol"}
}
```

2. Register your chain:
```go
chain.RegisterChain(&MyChain{})
```

#### Adding a New Protocol

1. Create a protocol validator:
```go
type MyValidator struct{}

func (v *MyValidator) ValidateTransaction(tx *Transaction) error {
    return nil
}
```

2. Register the validator:
```go
validator.RegisterValidator("my-protocol", &MyValidator{})
```

## Development

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