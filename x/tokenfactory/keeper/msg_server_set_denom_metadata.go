package keeper

import (
	"context"
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/you/nuahchain/x/tokenfactory/types"
)

type denomUnitJSON struct {
	Denom    string   `json:"denom"`
	Exponent uint32   `json:"exponent"`
	Aliases  []string `json:"aliases"`
}

func (k msgServer) SetDenomMetadata(goCtx context.Context, msg *types.MsgSetDenomMetadata) (*types.MsgSetDenomMetadataResponse, error) {
	admin, err := k.addressCodec.StringToBytes(msg.Owner)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.MustBeAdmin(ctx, msg.Base, admin); err != nil {
		return nil, err
	}

	var units []denomUnitJSON
	if err := json.Unmarshal([]byte(msg.DenomUnits), &units); err != nil {
		return nil, errorsmod.Wrap(err, "invalid denomUnits json")
	}

	bu := make([]*banktypes.DenomUnit, 0, len(units))
	for _, u := range units {
		bu = append(bu, &banktypes.DenomUnit{Denom: u.Denom, Exponent: u.Exponent, Aliases: u.Aliases})
	}

	md := banktypes.Metadata{
		Base:        msg.Base,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Display:     msg.Display,
		DenomUnits:  bu,
		Description: msg.Description,
	}

	if err := k.bankKeeper.SetDenomMetaData(ctx, md); err != nil {
		return nil, err
	}

	return &types.MsgSetDenomMetadataResponse{}, nil
}
