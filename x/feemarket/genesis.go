package feemarket

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	abci "github.com/reapchain/reapchain-core/abci/types"

	"github.com/reapchain/ethermint/x/feemarket/keeper"
	"github.com/reapchain/ethermint/x/feemarket/types"
)

// InitGenesis initializes genesis state based on exported genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	data types.GenesisState,
) []abci.ValidatorUpdate {
	k.SetParams(ctx, data.Params)
	k.SetBlockGasUsed(ctx, data.BlockGas)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis exports genesis state of the fee market module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:   k.GetParams(ctx),
		BlockGas: k.GetBlockGasUsed(ctx),
	}
}
