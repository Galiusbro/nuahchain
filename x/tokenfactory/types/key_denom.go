package types

import "cosmossdk.io/collections"

// DenomKey is the prefix to retrieve all Denom
var DenomKey = collections.NewPrefix("denom/value/")
