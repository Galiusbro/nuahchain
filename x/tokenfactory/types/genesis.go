package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:   DefaultParams(),
		DenomMap: []Denom{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	denomIndexMap := make(map[string]struct{})

	for _, elem := range gs.DenomMap {
		index := fmt.Sprint(elem.Denom)
		if _, ok := denomIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for denom")
		}
		denomIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
