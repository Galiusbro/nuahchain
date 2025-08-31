package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/you/nuahchain/x/tokenfactory/types"
)

func (k msgServer) ChangeAdmin(goCtx context.Context, msg *types.MsgChangeAdmin) (*types.MsgChangeAdminResponse, error) {
	admin, err := k.addressCodec.StringToBytes(msg.Owner)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.MustBeAdmin(ctx, msg.Denom, admin); err != nil {
		return nil, err
	}

	newAdmin, err := k.addressCodec.StringToBytes(msg.NewAdmin)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid new admin address")
	}

	d, _ := k.GetDenom(ctx, msg.Denom)
	d.Owner = sdk.AccAddress(newAdmin).String()
	k.SetDenom(ctx, d)

	return &types.MsgChangeAdminResponse{}, nil
}
