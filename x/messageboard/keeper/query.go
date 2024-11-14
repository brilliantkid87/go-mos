package keeper

import (
	"messageboard/x/messageboard/types"
)

var _ types.QueryServer = Keeper{}
