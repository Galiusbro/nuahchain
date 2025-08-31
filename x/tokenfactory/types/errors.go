package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/tokenfactory module sentinel errors
var (
	ErrInvalidSigner   = errorsmod.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrDenomExists     = errorsmod.Register(ModuleName, 2, "denom already exists")
	ErrInvalidSubdenom = errorsmod.Register(ModuleName, 3, "invalid subdenom")
	ErrNotAdmin        = errorsmod.Register(ModuleName, 4, "caller is not denom admin")
	ErrInvalidAmount   = errorsmod.Register(ModuleName, 5, "invalid amount")
)
