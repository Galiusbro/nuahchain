package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/you/nuahchain/x/tokenfactory/types"
)

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Owner); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	full := k.FullDenom(msg.Owner, msg.Denom)

	ok, err := k.Denom.Has(ctx, full)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, types.ErrDenomExists
	}

	denom := types.Denom{
		Owner:              msg.Owner,
		Denom:              full,
		Description:        msg.Description,
		Ticker:             msg.Ticker,
		Precision:          msg.Precision,
		Url:                msg.Url,
		MaxSupply:          msg.MaxSupply,
		Supply:             msg.Supply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	if err := k.Denom.Set(ctx, denom.Denom, denom); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Owner); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	full := k.FullDenom(msg.Owner, msg.Denom)

	val, err := k.Denom.Get(ctx, full)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	if msg.Owner != val.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	denom := types.Denom{
		Owner:              msg.Owner,
		Denom:              full,
		Description:        msg.Description,
		Ticker:             msg.Ticker,
		Precision:          msg.Precision,
		Url:                msg.Url,
		MaxSupply:          msg.MaxSupply,
		Supply:             msg.Supply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	if err := k.Denom.Set(ctx, denom.Denom, denom); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update denom")
	}

	return &types.MsgUpdateDenomResponse{}, nil
}

func (k msgServer) DeleteDenom(goCtx context.Context, msg *types.MsgDeleteDenom) (*types.MsgDeleteDenomResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Owner); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	full := k.FullDenom(msg.Owner, msg.Denom)

	val, err := k.Denom.Get(ctx, full)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	if msg.Owner != val.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Denom.Remove(ctx, full); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove denom")
	}

	return &types.MsgDeleteDenomResponse{}, nil
}
