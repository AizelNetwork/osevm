package keeper

import (
	"github.com/AizelNetwork/osevm/x/aizel/types"
)

var _ types.QueryServer = Keeper{}
