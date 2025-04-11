package aizel_test

import (
	"testing"

	keepertest "github.com/AizelNetwork/osevm/testutil/keeper"
	"github.com/AizelNetwork/osevm/testutil/nullify"
	aizel "github.com/AizelNetwork/osevm/x/aizel/module"
	"github.com/AizelNetwork/osevm/x/aizel/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AizelKeeper(t)
	aizel.InitGenesis(ctx, k, genesisState)
	got := aizel.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
