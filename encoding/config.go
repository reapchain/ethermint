package encoding

import (
	amino "github.com/reapchain/cosmos-sdk/codec"
	"github.com/reapchain/cosmos-sdk/codec/types"
	"github.com/reapchain/cosmos-sdk/simapp/params"
	"github.com/reapchain/cosmos-sdk/types/module"
	"github.com/reapchain/cosmos-sdk/x/auth/tx"

	enccodec "github.com/reapchain/ethermint/encoding/codec"
)

// MakeConfig creates an EncodingConfig for testing
func MakeConfig(mb module.BasicManager) params.EncodingConfig {
	cdc := amino.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := amino.NewProtoCodec(interfaceRegistry)

	encodingConfig := params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             cdc,
	}

	enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	mb.RegisterLegacyAminoCodec(encodingConfig.Amino)
	enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	mb.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
