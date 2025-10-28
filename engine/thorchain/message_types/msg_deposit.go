package message_types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgDeposit represents THORChain's MsgDeposit message
// Based on THORChain's cosmos SDK implementation for deposits/swaps
type MsgDeposit struct {
	Coins  []sdk.Coin `json:"coins"`
	Memo   string     `json:"memo"`
	Signer string     `json:"signer"`
}

// ValidateBasic implements sdk.Msg interface
func (m *MsgDeposit) ValidateBasic() error {
	if len(m.Coins) == 0 {
		return fmt.Errorf("coins cannot be empty")
	}
	if m.Signer == "" {
		return fmt.Errorf("signer cannot be empty")
	}
	return nil
}

// GetSigners implements sdk.Msg interface
func (m *MsgDeposit) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Signer)
	return []sdk.AccAddress{signer}
}

// Route implements sdk.Msg interface
func (m *MsgDeposit) Route() string {
	return "thorchain"
}

// Type implements sdk.Msg interface  
func (m *MsgDeposit) Type() string {
	return "deposit"
}

// ProtoMessage implements proto.Message interface
func (m *MsgDeposit) ProtoMessage() {}

// Reset implements proto.Message interface
func (m *MsgDeposit) Reset() {
	*m = MsgDeposit{}
}

// String implements proto.Message interface
func (m *MsgDeposit) String() string {
	return fmt.Sprintf("MsgDeposit{Coins: %v, Memo: %s, Signer: %s}", m.Coins, m.Memo, m.Signer)
}