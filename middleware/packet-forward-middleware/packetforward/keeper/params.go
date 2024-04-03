package keeper

import (
	"github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetFeePercentage retrieves the fee percentage for forwarded packets from the store.
func (k Keeper) GetFeePercentage(ctx sdk.Context) sdkmath.LegacyDec {
	return k.GetParams(ctx).FeePercentage

}

// GetParams returns the total set of pfm parameters.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(k.GetFeePercentage(ctx))
}

// SetParams sets the total set of pfm parameters.
func (k Keeper) SetParams(ctx sdk.Context, p types.Params) error {
	if err := p.Validate(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&p)
	store.Set(types.ParamsKey, bz)
	return nil
}
