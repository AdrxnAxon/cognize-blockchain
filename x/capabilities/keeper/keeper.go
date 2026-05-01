package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "cosmossdk.io/store/types"

	"github.com/cognize/axon/x/capabilities/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
	bankKeeper    types.BankKeeper
	stakingKeeper types.StakingKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		bankKeeper:    bankKeeper,
		stakingKeeper: stakingKeeper,
	}
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	// TODO: implement
}

func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.DefaultParams()
}
