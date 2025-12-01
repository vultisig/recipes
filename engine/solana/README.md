# Solana Engine - IDL Support

This directory contains the Solana transaction validation engine and IDL (Interface Description Language) support. IDL files are located in the `../../idl/` directory and define instruction schemas, account structures, and argument types for Solana programs.

## Directory Structure

```
engine/solana/
├── README.md          # This file - IDL integration guide
├── solana.go          # Main Solana engine implementation
├── loader.go          # IDL loading and type definitions
├── assert.go          # Transaction validation logic
├── compare/           # Solana-specific comparers
└── *_test.go          # Test files

../../idl/
├── embed.go           # Go embed directive for IDL files
├── spl_token.json     # SPL Token program IDL
└── system.json        # Solana System program IDL
```

## Adding New IDL Support

### 1. Create IDL JSON File

Add your new IDL file as `<protocol_name>.json` in the `../../idl/` directory. The IDL must follow this structure:

```json
{
  "version": "1.0.0",
  "name": "<protocol_name>",
  "instructions": [
    {
      "name": "<instruction_name>",
      "accounts": [
        {
          "name": "<account_name>",
          "isMut": true|false,
          "isSigner": true|false
        }
      ],
      "args": [
        {
          "name": "<argument_name>",
          "type": "<argument_type>"
        }
      ],
      "metadata": {
        "discriminator": [<byte_array>]
      }
    }
  ]
}
```

**Key Requirements:**
- Each instruction **must** have a `metadata.discriminator` field with a non-empty byte array
- The discriminator identifies the instruction in the transaction data (function selector, usually 8 bytes for Anchor programs, but it's not a standard, that's why it must be defined)
- Account names and argument names will be prefixed with `account_` and `arg_` respectively for rule constraints

### 2. Supported Argument Types

Currently supported argument types (defined in `loader.go:39-43`):

- `"u8"` - 8-bit unsigned integer
- `"u64"` - 64-bit unsigned integer
- `"publicKey"` - Solana public key (32 bytes)

**Adding New Argument Types:**

If you need a new argument type not listed above:

1. **Define the type constant** in `loader.go:39-43`:
   ```go
   const (
       argU8        argType = "u8"
       argU64       argType = "u64"
       argPublicKey argType = "publicKey"
       argYourNewType argType = "yourNewType"  // Add this
   )
   ```

2. **Add decoder support** in `assert.go:147-165`:
   ```go
   switch arg.Type {
   case argU8:
       err := decodeAndAssert(decoder, constraints, name, compare.NewUint8)
   case argU64:
       err := decodeAndAssert(decoder, constraints, name, compare.NewUint64)
   case argPublicKey:
       err := decodeAndAssert(decoder, constraints, name, solcmp.NewPubKey)
   case argYourNewType:  // Add this case
       err := decodeAndAssert(decoder, constraints, name, yourNewTypeComparer)
   }
   ```

3. **Create the comparer function** if needed (similar to `compare.NewUint8`, `solcmp.NewPubKey`)

### 3. Integration Points

The IDL files are automatically loaded by the Solana engine. Key integration locations:

- **Loading**: `loader.go:57` - `loadIDLDir()` function
- **Engine initialization**: `solana.go:17` - Called during `NewSolana()`
- **Registry**: `../registry.go:56` - Solana engine registered in global registry
- **Validation**: `solana.go:65-91` - IDL used for transaction validation
- **Discriminator validation**: `assert.go:77-94` - Function selector validation
- **Argument parsing**: `assert.go:130-168` - Borsh deserialization of instruction data

### 4. Embed Configuration

The `../../idl/embed.go` file uses Go's embed directive to include all `.json` files:

```go
//go:embed *.json
var Dir embed.FS
```

No changes needed here - new `.json` files are automatically included.

### 5. Testing Your IDL

After adding a new IDL file:

1. **Static testing**: Add test cases in `solana_static_test.go`
2. **Dynamic testing**: You can add new test cases in `solana_dynamic_test.go`, but new IDL would automatically be tested with current testdata
3. **Test discriminator**: Verify discriminator bytes match actual program instruction function selector

### 6. Example: Adding Jupiter Swap IDL (just for reference)

```json
{
  "version": "1.0.0",
  "name": "jupiter_swap",
  "instructions": [
    {
      "name": "swap",
      "accounts": [
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "userSourceTokenAccount",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "userDestinationTokenAccount",
          "isMut": true,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "amountIn",
          "type": "u64"
        },
        {
          "name": "minimumAmountOut",
          "type": "u64"
        }
      ],
      "metadata": {
        "discriminator": [248, 198, 158, 145, 225, 117, 135, 200]
      }
    }
  ]
}
```

### 7. Rules and Constraints

Once your IDL is added, you can create policy/rules that reference:

- **Accounts**: `account_<name>` (e.g., `account_userSourceTokenAccount`)
- **Arguments**: `arg_<name>` (e.g., `arg_amountIn`)
- **Protocol**: The filename without `.json` extension (e.g., `jupiter_swap`)
- **Function**: The instruction name (e.g., `swap`)

The discriminator is automatically validated - the transaction data must start with the exact discriminator bytes defined in the IDL.
