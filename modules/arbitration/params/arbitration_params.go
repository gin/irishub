package arbitrationparams

import (
	"github.com/irisnet/irishub/modules/params"
	sdk "github.com/irisnet/irishub/types"
	"time"
)

var ComplaintRetrospectParameter ComplaintRetrospectParam

var _ params.SignalParameter = (*ComplaintRetrospectParam)(nil)

type ComplaintRetrospectParam struct {
	Value      time.Duration
	paramSpace params.Subspace
}

func (param *ComplaintRetrospectParam) InitGenesis(genesisState interface{}) {
	param.Value = genesisState.(time.Duration)
}

func (param *ComplaintRetrospectParam) SetReadWriter(paramSpace params.Subspace) {
	param.paramSpace = paramSpace
}

func (param *ComplaintRetrospectParam) GetStoreKey() []byte {
	return []byte("arbitrationComplaintRetrospect")
}

func (param *ComplaintRetrospectParam) SaveValue(ctx sdk.Context) {
	param.paramSpace.Set(ctx, param.GetStoreKey(), param.Value)
}

func (param *ComplaintRetrospectParam) LoadValue(ctx sdk.Context) bool {
	if param.paramSpace.Has(ctx, param.GetStoreKey()) == false {
		return false
	}
	param.paramSpace.Get(ctx, param.GetStoreKey(), &param.Value)
	return true
}

var ArbitrationTimelimitParameter ArbitrationTimelimitParam

var _ params.SignalParameter = (*ArbitrationTimelimitParam)(nil)

type ArbitrationTimelimitParam struct {
	Value      time.Duration
	paramSpace params.Subspace
}

func (param *ArbitrationTimelimitParam) InitGenesis(genesisState interface{}) {
	param.Value = genesisState.(time.Duration)
}

func (param *ArbitrationTimelimitParam) SetReadWriter(paramSpace params.Subspace) {
	param.paramSpace = paramSpace
}

func (param *ArbitrationTimelimitParam) GetStoreKey() []byte {
	return []byte("ArbitrationTimelimit")
}

func (param *ArbitrationTimelimitParam) SaveValue(ctx sdk.Context) {
	param.paramSpace.Set(ctx, param.GetStoreKey(), param.Value)
}

func (param *ArbitrationTimelimitParam) LoadValue(ctx sdk.Context) bool {
	if param.paramSpace.Has(ctx, param.GetStoreKey()) == false {
		return false
	}
	param.paramSpace.Get(ctx, param.GetStoreKey(), &param.Value)
	return true
}
