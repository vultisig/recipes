package solana

import (
	"fmt"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
)

// TransactionData represents a decoded Solana transaction using proper types
type TransactionData struct {
	Transaction *solana.Transaction `json:"transaction"`
	Message     *solana.Message     `json:"message"`
}

// IDL types for instruction parsing (kept for compatibility)
type IDLInstruction struct {
	Name     string        `json:"name"`
	Accounts []IDLAccount  `json:"accounts"`
	Args     []IDLArgument `json:"args"`
}

type IDLAccount struct {
	Name     string `json:"name"`
	IsMut    bool   `json:"isMut"`
	IsSigner bool   `json:"isSigner"`
}

type IDLArgument struct {
	Name string      `json:"name"`
	Type interface{} `json:"type"`
}

// Well-known program IDs using solana-go constants
var (
	SystemProgramID   = system.ProgramID
	SPLTokenProgramID = token.ProgramID
	// Token 2022 Program ID - using actual public key
	SPLToken2022ID = solana.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")
)

// DecodeTransaction decodes a base64-encoded Solana transaction using solana-go
func DecodeTransaction(txBytes []byte) (*TransactionData, error) {
	// Parse the transaction using solana-go
	tx, err := solana.TransactionFromDecoder(bin.NewBorshDecoder(txBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to decode Solana transaction: %w", err)
	}

	return &TransactionData{
		Transaction: tx,
		Message:     &tx.Message,
	}, nil
}

// ParseTransferAmount extracts the lamports amount from a System Program transfer instruction
func ParseTransferAmount(instruction solana.CompiledInstruction, accounts []solana.PublicKey) (uint64, error) {
	// Check if this is a system program instruction
	programID := accounts[instruction.ProgramIDIndex]
	if !programID.Equals(SystemProgramID) {
		return 0, fmt.Errorf("not a system program instruction")
	}

	// For system transfers, we need to manually decode the instruction data
	// System Transfer instruction format: [instruction_type:4bytes][lamports:8bytes]
	if len(instruction.Data) < 12 {
		return 0, fmt.Errorf("insufficient data for system transfer instruction")
	}

	// Decode instruction type (should be 2 for Transfer)
	instrType := bin.NewBorshDecoder(instruction.Data[:4])
	var instructionIndex uint32
	err := instrType.Decode(&instructionIndex)
	if err != nil || instructionIndex != 2 {
		return 0, fmt.Errorf("not a transfer instruction")
	}

	// Decode lamports amount
	amountDecoder := bin.NewBorshDecoder(instruction.Data[4:12])
	var lamports uint64
	err = amountDecoder.Decode(&lamports)
	if err != nil {
		return 0, fmt.Errorf("failed to decode transfer amount: %w", err)
	}

	return lamports, nil
}

// ParseSPLTransferAmount extracts the amount from an SPL token transfer instruction
func ParseSPLTransferAmount(instruction solana.CompiledInstruction, accounts []solana.PublicKey) (uint64, error) {
	// Check if this is an SPL token program instruction
	programID := accounts[instruction.ProgramIDIndex]
	if !IsSPLTokenProgram(programID) {
		return 0, fmt.Errorf("not an SPL token program instruction")
	}

	// For SPL token transfers, we need to manually decode based on instruction type
	if len(instruction.Data) < 9 {
		return 0, fmt.Errorf("insufficient data for token transfer instruction")
	}

	// First byte is the instruction type
	instructionType := instruction.Data[0]

	var amount uint64
	switch instructionType {
	case 3: // Transfer instruction
		if len(instruction.Data) < 9 {
			return 0, fmt.Errorf("insufficient data for transfer instruction")
		}
		// Amount is bytes 1-8
		amountDecoder := bin.NewBorshDecoder(instruction.Data[1:9])
		err := amountDecoder.Decode(&amount)
		if err != nil {
			return 0, fmt.Errorf("failed to decode transfer amount: %w", err)
		}
	case 12: // TransferChecked instruction
		if len(instruction.Data) < 10 {
			return 0, fmt.Errorf("insufficient data for transfer checked instruction")
		}
		// Amount is bytes 1-8, decimals is byte 9
		amountDecoder := bin.NewBorshDecoder(instruction.Data[1:9])
		err := amountDecoder.Decode(&amount)
		if err != nil {
			return 0, fmt.Errorf("failed to decode transfer checked amount: %w", err)
		}
	default:
		return 0, fmt.Errorf("not a transfer instruction, got instruction type: %d", instructionType)
	}

	return amount, nil
}

// GetTransferDestination extracts the destination account from a transfer instruction
func GetTransferDestination(instruction solana.CompiledInstruction, accounts []solana.PublicKey) (solana.PublicKey, error) {
	programID := accounts[instruction.ProgramIDIndex]

	if programID.Equals(SystemProgramID) {
		// For system transfers, destination is the second account
		if len(instruction.Accounts) < 2 {
			return solana.PublicKey{}, fmt.Errorf("insufficient accounts for system transfer")
		}
		return accounts[instruction.Accounts[1]], nil
	}

	if IsSPLTokenProgram(programID) {
		// For SPL token transfers, we need to check the instruction type
		if len(instruction.Data) < 1 {
			return solana.PublicKey{}, fmt.Errorf("insufficient instruction data")
		}

		instructionType := instruction.Data[0]
		switch instructionType {
		case 3: // Transfer instruction
			// Destination is the second account
			if len(instruction.Accounts) < 2 {
				return solana.PublicKey{}, fmt.Errorf("insufficient accounts for token transfer")
			}
			return accounts[instruction.Accounts[1]], nil
		case 12: // TransferChecked instruction
			// For TransferChecked, destination is the third account
			if len(instruction.Accounts) < 3 {
				return solana.PublicKey{}, fmt.Errorf("insufficient accounts for token transfer checked")
			}
			return accounts[instruction.Accounts[2]], nil
		default:
			return solana.PublicKey{}, fmt.Errorf("not a token transfer instruction")
		}
	}

	return solana.PublicKey{}, fmt.Errorf("unsupported program for transfer: %s", programID.String())
}

// ParseInstructionArgs parses instruction arguments based on IDL (simplified version for compatibility)
func ParseInstructionArgs(data []byte, instruction *IDLInstruction) ([]interface{}, error) {
	// This is a simplified parser that maintains compatibility with the existing interface
	// In practice, you'd use the proper instruction decoders from solana-go

	args := make([]interface{}, len(instruction.Args))

	// For basic instruction types, we can still do simple parsing
	if len(data) < 4 {
		return nil, fmt.Errorf("insufficient data for instruction parsing")
	}

	// Skip instruction discriminator (first 4 or 8 bytes depending on instruction)
	offset := 4
	if len(data) >= 8 {
		offset = 8
	}

	for i, arg := range instruction.Args {
		switch argType := arg.Type.(string); argType {
		case "u64":
			if offset+8 > len(data) {
				return nil, fmt.Errorf("insufficient data for u64 argument")
			}
			// Use solana-go's binary decoder for proper endianness
			decoder := bin.NewBorshDecoder(data[offset : offset+8])
			var value uint64
			err := decoder.Decode(&value)
			if err != nil {
				return nil, fmt.Errorf("failed to decode u64: %w", err)
			}
			args[i] = value
			offset += 8

		case "u8":
			if offset+1 > len(data) {
				return nil, fmt.Errorf("insufficient data for u8 argument")
			}
			args[i] = data[offset]
			offset += 1

		case "publicKey":
			if offset+32 > len(data) {
				return nil, fmt.Errorf("insufficient data for publicKey argument")
			}
			pubkey := solana.PublicKeyFromBytes(data[offset : offset+32])
			args[i] = pubkey.String() // Convert to string for compatibility
			offset += 32

		default:
			return nil, fmt.Errorf("unsupported argument type: %s", argType)
		}
	}

	return args, nil
}

// GetProgramNameByID returns the human-readable name for a program ID
func GetProgramNameByID(programID solana.PublicKey) string {
	switch {
	case programID.Equals(SystemProgramID):
		return "system"
	case IsSPLTokenProgram(programID):
		return "spl_token"
	default:
		return "unknown"
	}
}

// IsSPLTokenProgram checks if the given program ID is an SPL Token program
func IsSPLTokenProgram(programID solana.PublicKey) bool {
	return programID.Equals(SPLTokenProgramID) || programID.Equals(SPLToken2022ID)
}

// IsSystemProgram checks if the given program ID is the System Program
func IsSystemProgram(programID solana.PublicKey) bool {
	return programID.Equals(SystemProgramID)
}

// FindTransferInstruction finds the first transfer instruction in a transaction
func FindTransferInstruction(txData *TransactionData) (*solana.CompiledInstruction, error) {
	if txData.Message == nil || len(txData.Message.Instructions) == 0 {
		return nil, fmt.Errorf("no instructions in transaction")
	}

	accounts := txData.Message.AccountKeys

	for _, instruction := range txData.Message.Instructions {
		programID := accounts[instruction.ProgramIDIndex]

		// Check if it's a system program transfer
		if programID.Equals(SystemProgramID) {
			// Check if it's a transfer instruction (instruction type 2)
			if len(instruction.Data) >= 4 {
				instrDecoder := bin.NewBorshDecoder(instruction.Data[:4])
				var instructionIndex uint32
				err := instrDecoder.Decode(&instructionIndex)
				if err == nil && instructionIndex == 2 {
					return &instruction, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("no transfer instruction found")
}
