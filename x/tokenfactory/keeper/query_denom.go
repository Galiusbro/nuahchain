package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/you/nuahchain/x/tokenfactory/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListDenom(ctx context.Context, req *types.QueryAllDenomRequest) (*types.QueryAllDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	denoms, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Denom,
		req.Pagination,
		func(_ string, value types.Denom) (types.Denom, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDenomResponse{Denom: denoms, Pagination: pageRes}, nil
}

func (q queryServer) GetDenom(ctx context.Context, req *types.QueryGetDenomRequest) (*types.QueryGetDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Denom.Get(ctx, req.Denom)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetDenomResponse{Denom: val}, nil
}
