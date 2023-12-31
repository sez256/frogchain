package investibc

import (
	"math/rand"

	"frogchain/testutil/sample"
	investibcsimulation "frogchain/x/investibc/simulation"
	"frogchain/x/investibc/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = investibcsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgSetAdminAccount = "op_weight_msg_set_admin_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetAdminAccount int = 100

	opWeightMsgDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeposit int = 100

	opWeightMsgInitIcaAccount = "op_weight_msg_register_ica_account"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitIcaAccount int = 100

	opWeightMsgSetDepositDenom = "op_weight_msg_set_deposit_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetDepositDenom int = 100

	opWeightMsgWithdraw = "op_weight_msg_withdraw"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdraw int = 100

	opWeightMsgSetLiquidityDenom = "op_weight_msg_set_liquidity_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetLiquidityDenom int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	investibcGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&investibcGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSetAdminAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetAdminAccount, &weightMsgSetAdminAccount, nil,
		func(_ *rand.Rand) {
			weightMsgSetAdminAccount = defaultWeightMsgSetAdminAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetAdminAccount,
		investibcsimulation.SimulateMsgSetAdminAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeposit, &weightMsgDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgDeposit = defaultWeightMsgDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeposit,
		investibcsimulation.SimulateMsgDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInitIcaAccount int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInitIcaAccount, &weightMsgInitIcaAccount, nil,
		func(_ *rand.Rand) {
			weightMsgInitIcaAccount = defaultWeightMsgInitIcaAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitIcaAccount,
		investibcsimulation.SimulateMsgInitIcaAccount(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetDepositDenom int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetDepositDenom, &weightMsgSetDepositDenom, nil,
		func(_ *rand.Rand) {
			weightMsgSetDepositDenom = defaultWeightMsgSetDepositDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetDepositDenom,
		investibcsimulation.SimulateMsgSetDepositDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdraw int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdraw, &weightMsgWithdraw, nil,
		func(_ *rand.Rand) {
			weightMsgWithdraw = defaultWeightMsgWithdraw
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdraw,
		investibcsimulation.SimulateMsgWithdraw(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetLiquidityDenom int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetLiquidityDenom, &weightMsgSetLiquidityDenom, nil,
		func(_ *rand.Rand) {
			weightMsgSetLiquidityDenom = defaultWeightMsgSetLiquidityDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetLiquidityDenom,
		investibcsimulation.SimulateMsgSetLiquidityDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetAdminAccount,
			defaultWeightMsgSetAdminAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				investibcsimulation.SimulateMsgSetAdminAccount(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeposit,
			defaultWeightMsgDeposit,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				investibcsimulation.SimulateMsgDeposit(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitIcaAccount,
			defaultWeightMsgInitIcaAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				investibcsimulation.SimulateMsgInitIcaAccount(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetDepositDenom,
			defaultWeightMsgSetDepositDenom,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				investibcsimulation.SimulateMsgSetDepositDenom(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWithdraw,
			defaultWeightMsgWithdraw,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				investibcsimulation.SimulateMsgWithdraw(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),

		simulation.NewWeightedProposalMsg(
			opWeightMsgSetLiquidityDenom,
			defaultWeightMsgSetLiquidityDenom,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				investibcsimulation.SimulateMsgSetLiquidityDenom(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
