package keeper

import (
	"github.com/jhonnik08/stablecoinproject/x/stablecoinproject/types"
)

var _ types.QueryServer = Keeper{}
