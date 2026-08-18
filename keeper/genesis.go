package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/nft-transfer/types"
)

// InitGenesis initializes the ibc nft-transfer state.
// Port binding is done by the host chain's IBC router (ibc-go v10).
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	k.SetPort(ctx, state.PortId)

	for _, trace := range state.Traces {
		k.SetClassTrace(ctx, trace)
	}

	if err := k.SetParams(ctx, state.Params); err != nil {
		panic(fmt.Sprintf("SetParams failed: %v", err))
	}
}

// ExportGenesis exports ibc nft-transfer module's portID and class trace info into its genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		PortId: k.GetPort(ctx),
		Traces: k.GetAllClassTraces(ctx),
		Params: k.GetParams(ctx),
	}
}

// SetPort sets the port ID for the nft-transfer module
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.PortKey), []byte(portID))
}

// GetPort gets the port ID for the nft-transfer module
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get([]byte(types.PortKey)))
}
