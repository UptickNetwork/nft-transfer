package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/nft-transfer/types"
)

// InitGenesis initializes the ibc nft-transfer state and binds to PortID.
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	k.SetPort(ctx, state.PortId)

	for _, trace := range state.Traces {
		k.SetClassTrace(ctx, trace)
	}

	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, state.PortId) {
		// nft-transfer module binds to the nft-transfer port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, state.PortId)
		if err != nil {
			panic(fmt.Sprintf("could not claim port capability: %v", err))
		}
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

// BindPort binds to the nft-transfer port
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	return k.portKeeper.BindPort(ctx, portID)
}

// IsBound checks if the nft-transfer module is already bound to the port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	return k.GetPort(ctx) == portID
}
