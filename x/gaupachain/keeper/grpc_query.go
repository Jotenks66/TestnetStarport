package keeper

import (
	"github.com/GaupaLabs/gaupachain/x/gaupachain/types"
)

var _ types.QueryServer = Keeper{}
