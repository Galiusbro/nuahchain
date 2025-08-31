package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/you/nuahchain/x/tokenfactory/keeper"
	"github.com/you/nuahchain/x/tokenfactory/types"
)

func createNDenom(keeper keeper.Keeper, ctx context.Context, n int) []types.Denom {
	items := make([]types.Denom, n)
	for i := range items {
		items[i].Denom = strconv.Itoa(i)
		items[i].Description = strconv.Itoa(i)
		items[i].Ticker = strconv.Itoa(i)
		items[i].Precision = int64(i)
		items[i].Url = strconv.Itoa(i)
		items[i].MaxSupply = int64(i)
		items[i].Supply = int64(i)
		items[i].CanChangeMaxSupply = true
		_ = keeper.Denom.Set(ctx, items[i].Denom, items[i])
	}
	return items
}

func TestDenomQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNDenom(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetDenomRequest
		response *types.QueryGetDenomResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetDenomRequest{
				Denom: msgs[0].Denom,
			},
			response: &types.QueryGetDenomResponse{Denom: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetDenomRequest{
				Denom: msgs[1].Denom,
			},
			response: &types.QueryGetDenomResponse{Denom: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetDenomRequest{
				Denom: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetDenom(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestDenomQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNDenom(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDenomRequest {
		return &types.QueryAllDenomRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListDenom(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Denom), step)
			require.Subset(t, msgs, resp.Denom)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListDenom(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Denom), step)
			require.Subset(t, msgs, resp.Denom)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListDenom(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Denom)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListDenom(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
