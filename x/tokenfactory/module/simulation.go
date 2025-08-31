package tokenfactory

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/you/nuahchain/testutil/sample"
	tokenfactorysimulation "github.com/you/nuahchain/x/tokenfactory/simulation"
	"github.com/you/nuahchain/x/tokenfactory/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tokenfactoryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DenomMap: []types.Denom{{Owner: sample.AccAddress(),
			Denom: "0",
		}, {Owner: sample.AccAddress(),
			Denom: "1",
		}}}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tokenfactoryGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateDenom          = "op_weight_msg_tokenfactory"
		defaultWeightMsgCreateDenom int = 100
	)

	var weightMsgCreateDenom int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateDenom, &weightMsgCreateDenom, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDenom = defaultWeightMsgCreateDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDenom,
		tokenfactorysimulation.SimulateMsgCreateDenom(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateDenom          = "op_weight_msg_tokenfactory"
		defaultWeightMsgUpdateDenom int = 100
	)

	var weightMsgUpdateDenom int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateDenom, &weightMsgUpdateDenom, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDenom = defaultWeightMsgUpdateDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDenom,
		tokenfactorysimulation.SimulateMsgUpdateDenom(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteDenom          = "op_weight_msg_tokenfactory"
		defaultWeightMsgDeleteDenom int = 100
	)

	var weightMsgDeleteDenom int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteDenom, &weightMsgDeleteDenom, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDenom = defaultWeightMsgDeleteDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDenom,
		tokenfactorysimulation.SimulateMsgDeleteDenom(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgMint          = "op_weight_msg_tokenfactory"
		defaultWeightMsgMint int = 100
	)

	var weightMsgMint int
	simState.AppParams.GetOrGenerate(opWeightMsgMint, &weightMsgMint, nil,
		func(_ *rand.Rand) {
			weightMsgMint = defaultWeightMsgMint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMint,
		tokenfactorysimulation.SimulateMsgMint(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgBurn          = "op_weight_msg_tokenfactory"
		defaultWeightMsgBurn int = 100
	)

	var weightMsgBurn int
	simState.AppParams.GetOrGenerate(opWeightMsgBurn, &weightMsgBurn, nil,
		func(_ *rand.Rand) {
			weightMsgBurn = defaultWeightMsgBurn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurn,
		tokenfactorysimulation.SimulateMsgBurn(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgChangeAdmin          = "op_weight_msg_tokenfactory"
		defaultWeightMsgChangeAdmin int = 100
	)

	var weightMsgChangeAdmin int
	simState.AppParams.GetOrGenerate(opWeightMsgChangeAdmin, &weightMsgChangeAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgChangeAdmin = defaultWeightMsgChangeAdmin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgChangeAdmin,
		tokenfactorysimulation.SimulateMsgChangeAdmin(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgSetDenomMetadata          = "op_weight_msg_tokenfactory"
		defaultWeightMsgSetDenomMetadata int = 100
	)

	var weightMsgSetDenomMetadata int
	simState.AppParams.GetOrGenerate(opWeightMsgSetDenomMetadata, &weightMsgSetDenomMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgSetDenomMetadata = defaultWeightMsgSetDenomMetadata
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetDenomMetadata,
		tokenfactorysimulation.SimulateMsgSetDenomMetadata(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
