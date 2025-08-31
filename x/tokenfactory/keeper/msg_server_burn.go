package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/you/nuahchain/x/tokenfactory/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	admin, err := k.addressCodec.StringToBytes(msg.Owner)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.MustBeAdmin(ctx, msg.Denom, admin); err != nil {
		return nil, err
	}

	if msg.Amount <= 0 {
		return nil, types.ErrInvalidAmount
	}

	coins := sdk.NewCoins(sdk.NewCoin(msg.Denom, sdkmath.NewInt(msg.Amount)))
	adminAddr := sdk.AccAddress(admin)

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, adminAddr, types.ModuleAccountName, coins); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleAccountName, coins); err != nil {
		return nil, err
	}

	return &types.MsgBurnResponse{}, nil
}
