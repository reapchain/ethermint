package v2

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	paramtypes "github.com/reapchain/cosmos-sdk/x/params/types"
	"github.com/reapchain/ethermint/x/evm/types"
)

// MigrateStore sets the default AllowUnprotectedTxs parameter.
func MigrateStore(ctx sdk.Context, paramstore *paramtypes.Subspace) error {
	if !paramstore.HasKeyTable() {
		ps := paramstore.WithKeyTable(types.ParamKeyTable())
		paramstore = &ps
	}

	// add RejectUnprotected
	paramstore.Set(ctx, types.ParamStoreKeyAllowUnprotectedTxs, types.DefaultAllowUnprotectedTxs)
	return nil
}
