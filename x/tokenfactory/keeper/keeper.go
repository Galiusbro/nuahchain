package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/you/nuahchain/x/tokenfactory/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte

	Schema collections.Schema
	Params collections.Item[types.Params]

	authKeeper types.AuthKeeper
	bankKeeper types.BankKeeper
	Denom      collections.Map[string, types.Denom]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,

	authKeeper types.AuthKeeper,
	bankKeeper types.BankKeeper,
) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,

		authKeeper: authKeeper,
		bankKeeper: bankKeeper,
		Params:     collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		Denom:      collections.NewMap(sb, types.DenomKey, "denom", collections.StringKey, codec.CollValue[types.Denom](cdc))}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() []byte {
	return k.authority
}

// FullDenom constructs the full denom path for a given owner and subdenom
func (k Keeper) FullDenom(owner, subdenom string) string {
	return fmt.Sprintf("factory/%s/%s", owner, subdenom)
}

// GetDenom returns denom data from store
func (k Keeper) GetDenom(ctx context.Context, denom string) (types.Denom, bool) {
	val, err := k.Denom.Get(ctx, denom)
	if err != nil {
		return types.Denom{}, false
	}
	return val, true
}

// SetDenom stores denom data
func (k Keeper) SetDenom(ctx context.Context, d types.Denom) {
	if err := k.Denom.Set(ctx, d.Denom, d); err != nil {
		panic(err)
	}
}

// MustBeAdmin checks that caller is admin of denom
func (k Keeper) MustBeAdmin(ctx context.Context, denom string, caller sdk.AccAddress) error {
	d, ok := k.GetDenom(ctx, denom)
	if !ok {
		return sdkerrors.ErrInvalidRequest.Wrap("denom does not exist")
	}
	if d.Owner != caller.String() {
		return types.ErrNotAdmin
	}
	return nil
}
