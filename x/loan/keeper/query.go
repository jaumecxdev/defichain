package keeper

import (
	"defichain/x/loan/types"
)

var _ types.QueryServer = Keeper{}
