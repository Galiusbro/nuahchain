package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/you/nuahchain/x/tokenfactory/keeper"
	"github.com/you/nuahchain/x/tokenfactory/types"
)

func TestDenomMsgServerCreate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)
	owner, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateDenom{Owner: owner,
			Denom: strconv.Itoa(i),
		}
		_, err := srv.CreateDenom(f.ctx, expected)
		require.NoError(t, err)
		rst, err := f.keeper.Denom.Get(f.ctx, f.keeper.FullDenom(expected.Owner, expected.Denom))
		require.NoError(t, err)
		require.Equal(t, expected.Owner, rst.Owner)
	}
}

func TestDenomMsgServerUpdate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	owner, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	expected := &types.MsgCreateDenom{Owner: owner,
		Denom: strconv.Itoa(0),
	}
	_, err = srv.CreateDenom(f.ctx, expected)
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgUpdateDenom
		err     error
	}{
		{
			desc: "invalid address",
			request: &types.MsgUpdateDenom{Owner: "invalid",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			desc: "unauthorized",
			request: &types.MsgUpdateDenom{Owner: unauthorizedAddr,
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "key not found",
			request: &types.MsgUpdateDenom{Owner: owner,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "completed",
			request: &types.MsgUpdateDenom{Owner: owner,
				Denom: strconv.Itoa(0),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.UpdateDenom(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, err := f.keeper.Denom.Get(f.ctx, f.keeper.FullDenom(expected.Owner, expected.Denom))
				require.NoError(t, err)
				require.Equal(t, expected.Owner, rst.Owner)
			}
		})
	}
}

func TestDenomMsgServerDelete(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	owner, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	_, err = srv.CreateDenom(f.ctx, &types.MsgCreateDenom{Owner: owner,
		Denom: strconv.Itoa(0),
	})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgDeleteDenom
		err     error
	}{
		{
			desc: "invalid address",
			request: &types.MsgDeleteDenom{Owner: "invalid",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			desc: "unauthorized",
			request: &types.MsgDeleteDenom{Owner: unauthorizedAddr,
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "key not found",
			request: &types.MsgDeleteDenom{Owner: owner,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "completed",
			request: &types.MsgDeleteDenom{Owner: owner,
				Denom: strconv.Itoa(0),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.DeleteDenom(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				full := f.keeper.FullDenom(tc.request.Owner, tc.request.Denom)
				found, err := f.keeper.Denom.Has(f.ctx, full)
				require.NoError(t, err)
				require.False(t, found)
			}
		})
	}
}
