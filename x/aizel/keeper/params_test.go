package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "github.com/AizelNetwork/osevm/testutil/keeper"
    "github.com/AizelNetwork/osevm/x/aizel/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AizelKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
