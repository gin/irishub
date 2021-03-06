package simulation

import (
	"encoding/json"
	"math/rand"
	"testing"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/irisnet/irishub/types"
	"github.com/irisnet/irishub/modules/bank"
	"github.com/irisnet/irishub/modules/stake"
	"github.com/irisnet/irishub/modules/gov"
	"github.com/irisnet/irishub/modules/mock"
	"github.com/irisnet/irishub/modules/mock/simulation"
	distr "github.com/irisnet/irishub/modules/distribution"
	"github.com/irisnet/irishub/modules/guardian"
	protocolKeeper "github.com/irisnet/irishub/app/protocol/keeper"
	govtypes "github.com/irisnet/irishub/types/gov"
)

// TestGovWithRandomMessages
func TestGovWithRandomMessages(t *testing.T) {
	mapp := mock.NewApp()

	bank.RegisterCodec(mapp.Cdc)
	gov.RegisterCodec(mapp.Cdc)

	bankKeeper := mapp.BankKeeper
	stakeKey := mapp.KeyStake
	stakeTKey := mapp.TkeyStake
	paramKey := mapp.KeyParams
	govKey := sdk.NewKVStoreKey("gov")
	distrKey := sdk.NewKVStoreKey("distr")
	guardianKey := sdk.NewKVStoreKey("guardian")
    protocolKey := sdk.NewKVStoreKey("protocol")
	protocolKeeper := protocolKeeper.NewKeeper(mapp.Cdc,protocolKey)

	paramKeeper := mapp.ParamsKeeper
	stakeKeeper := stake.NewKeeper(
		mapp.Cdc, stakeKey,
		stakeTKey, bankKeeper,
		paramKeeper.Subspace(stake.DefaultParamspace),
		stake.DefaultCodespace,
	)
	distrKeeper := distr.NewKeeper(
		mapp.Cdc,
		distrKey,
		mapp.ParamsKeeper.Subspace(distr.DefaultParamspace),
		mapp.BankKeeper, &stakeKeeper, mapp.FeeCollectionKeeper,
		distr.DefaultCodespace,
	)
	guardianKeeper := guardian.NewKeeper(
		mapp.Cdc,
		guardianKey,
		guardian.DefaultCodespace,
	)
	govKeeper := gov.NewKeeper(
		mapp.Cdc,
		govKey,
		distrKeeper,
		bankKeeper,
		guardianKeeper,
		stakeKeeper,
		protocolKeeper,
		govtypes.DefaultCodespace,
	)

	mapp.Router().AddRoute("gov", []*sdk.KVStoreKey{govKey, mapp.KeyAccount, stakeKey, paramKey}, gov.NewHandler(govKeeper))
	mapp.SetEndBlocker(func(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
		gov.EndBlocker(ctx, govKeeper)
		return abci.ResponseEndBlock{}
	})

	err := mapp.CompleteSetup(govKey,protocolKey)
	if err != nil {
		panic(err)
	}

	appStateFn := func(r *rand.Rand, accs []simulation.Account) json.RawMessage {
		simulation.RandomSetGenesis(r, mapp, accs, []string{"stake"})
		return json.RawMessage("{}")
	}

	setup := func(r *rand.Rand, accs []simulation.Account) {
		ctx := mapp.NewContext(false, abci.Header{})
		stake.InitGenesis(ctx, stakeKeeper, stake.DefaultGenesisState())

		gov.InitGenesis(ctx, govKeeper, gov.DefaultGenesisState())
	}

	// Test with unscheduled votes
	simulation.Simulate(
		t, mapp.BaseApp, appStateFn,
		[]simulation.WeightedOperation{
			{2, SimulateMsgSubmitProposal(govKeeper, stakeKeeper)},
			{3, SimulateMsgDeposit(govKeeper, stakeKeeper)},
			{20, SimulateMsgVote(govKeeper, stakeKeeper)},
		}, []simulation.RandSetup{
			setup,
		}, []simulation.Invariant{
			//AllInvariants(),
		}, 10, 100,
		false,
	)

	// Test with scheduled votes
	simulation.Simulate(
		t, mapp.BaseApp, appStateFn,
		[]simulation.WeightedOperation{
			{10, SimulateSubmittingVotingAndSlashingForProposal(govKeeper, stakeKeeper)},
			{5, SimulateMsgDeposit(govKeeper, stakeKeeper)},
		}, []simulation.RandSetup{
			setup,
		}, []simulation.Invariant{
			AllInvariants(),
		}, 10, 100,
		false,
	)
}
