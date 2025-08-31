package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/you/nuahchain/x/tokenfactory/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
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

	recipientAddr, err := k.addressCodec.StringToBytes(msg.Recipient)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid recipient address")
	}

	coins := sdk.NewCoins(sdk.NewCoin(msg.Denom, sdk.NewInt(msg.Amount)))

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleAccountName, coins); err != nil {
		return nil, err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccountName, sdk.AccAddress(recipientAddr), coins); err != nil {
		return nil, err
	}

	return &types.MsgMintResponse{}, nil
}
