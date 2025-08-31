package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/you/nuahchain/x/tokenfactory/types"
)

func (k msgServer) Mint(ctx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Owner); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgMintResponse{}, nil
}
