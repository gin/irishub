package app

import (
	"fmt"

	"github.com/irisnet/irishub/store"
	sdk "github.com/irisnet/irishub/types"
	dbm "github.com/tendermint/tendermint/libs/db"

	"github.com/irisnet/irishub/app/protocol"
)

// File for storing in-package BaseApp optional functions,
// for options that need access to non-exported fields of the BaseApp

// SetPruning sets a pruning option on the multistore associated with the app
func SetPruning(pruning string) func(*BaseApp) {
	var pruningEnum sdk.PruningStrategy
	switch pruning {
	case "nothing":
		pruningEnum = sdk.PruneNothing
	case "everything":
		pruningEnum = sdk.PruneEverything
	case "syncable":
		pruningEnum = sdk.PruneSyncable
	default:
		panic(fmt.Sprintf("Invalid pruning strategy: %s", pruning))
	}
	return func(bap *BaseApp) {
		bap.cms.SetPruning(pruningEnum)
	}
}

// SetMinimumFees returns an option that sets the minimum fees on the app.
func SetMinimumFees(minFees string) func(*BaseApp) {
	fees, err := sdk.ParseCoins(minFees)
	if err != nil {
		panic(fmt.Sprintf("invalid minimum fees: %v", err))
	}
	return func(bap *BaseApp) { bap.SetMinimumFees(fees) }
}

// nolint - Setter functions
func (app *BaseApp) SetName(name string) {
	if app.sealed {
		panic("SetName() on sealed BaseApp")
	}
	app.name = name
}
func (app *BaseApp) SetDB(db dbm.DB) {
	if app.sealed {
		panic("SetDB() on sealed BaseApp")
	}
	app.db = db
}
func (app *BaseApp) SetCMS(cms store.CommitMultiStore) {
	if app.sealed {
		panic("SetEndBlocker() on sealed BaseApp")
	}
	app.cms = cms
}
func (app *BaseApp) SetTxDecoder(txDecoder sdk.TxDecoder) {
	if app.sealed {
		panic("SetTxDecoder() on sealed BaseApp")
	}
	app.txDecoder = txDecoder
}
func (app *BaseApp) SetAddrPeerFilter(pf sdk.PeerFilter) {
	if app.sealed {
		panic("SetAddrPeerFilter() on sealed BaseApp")
	}
	app.addrPeerFilter = pf
}
func (app *BaseApp) SetPubKeyPeerFilter(pf sdk.PeerFilter) {
	if app.sealed {
		panic("SetPubKeyPeerFilter() on sealed BaseApp")
	}
	app.pubkeyPeerFilter = pf
}

func (app *BaseApp) SetProtocolEngine(pe *protocol.ProtocolEngine) {
	if app.sealed {
		panic("SetPubKeyPeerFilter() on sealed BaseApp")
	}
	app.Engine = pe
}

func (app *BaseApp) Seal()          { app.sealed = true }
func (app *BaseApp) IsSealed() bool { return app.sealed }
func (app *BaseApp) enforceSeal() {
	if !app.sealed {
		panic("enforceSeal() on BaseApp but not sealed")
	}
}
