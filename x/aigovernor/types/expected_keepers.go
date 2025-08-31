package types

import (
	"context"

	"cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AuthKeeper defines the expected interface for the Auth module.
type AuthKeeper interface {
	AddressCodec() address.Codec
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

type TokenFactoryKeeper interface {
	CreateDenom(ctx context.Context, creatorAddr string, subdenom string) (newTokenDenom string, err error)
	MintTokens(ctx context.Context, tokens sdk.Coin, mintToAddress string) error
	BurnTokens(ctx context.Context, tokens sdk.Coin, burnFromAddress string) error
	ChangeAdmin(ctx context.Context, denom, newAdmin, creator string) error
}
