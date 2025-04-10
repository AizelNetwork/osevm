package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "github.com/evmos/os/testutil/keeper"
    "github.com/evmos/os/x/aizel/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AizelKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
