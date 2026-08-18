package keeper

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	porttypes "github.com/cosmos/ibc-go/v10/modules/core/05-port/types"

	"github.com/bianjieai/nft-transfer/types"
)

// Keeper defines the IBC non fungible transfer keeper
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.Codec
	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string

	ics4Wrapper   porttypes.ICS4Wrapper
	channelKeeper types.ChannelKeeper
	portKeeper    types.PortKeeper
	nftKeeper     types.NFTKeeper
	authKeeper    types.AccountKeeper
}

// NewKeeper creates a new IBC nft-transfer Keeper instance
func NewKeeper(
	cdc codec.Codec,
	key storetypes.StoreKey,
	authority string,
	ics4Wrapper porttypes.ICS4Wrapper,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	authKeeper types.AccountKeeper,
	nftKeeper types.NFTKeeper,
) Keeper {
	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		authority:     authority,
		ics4Wrapper:   ics4Wrapper,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		nftKeeper:     nftKeeper,
		authKeeper:    authKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// SetEscrowAddress is a no-op in ibc-go v10 (escrow addresses are computed on-the-fly).
// Kept for API compatibility with ibc_module.go callbacks.
func (k Keeper) SetEscrowAddress(ctx sdk.Context, portID, channelID string) {
	// no-op: escrow address is deterministic via types.GetEscrowAddress(portID, channelID)
}
