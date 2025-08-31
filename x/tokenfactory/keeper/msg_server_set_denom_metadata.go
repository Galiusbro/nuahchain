package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/you/nuahchain/x/tokenfactory/types"
)

func (k msgServer) SetDenomMetadata(ctx context.Context, msg *types.MsgSetDenomMetadata) (*types.MsgSetDenomMetadataResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Owner); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgSetDenomMetadataResponse{}, nil
}
