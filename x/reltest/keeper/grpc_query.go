package keeper

import (
	"github.com/ilgooz/reltest/x/reltest/types"
)

var _ types.QueryServer = Keeper{}
