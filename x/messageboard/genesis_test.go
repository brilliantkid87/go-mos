package messageboard_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "messageboard/testutil/keeper"
	"messageboard/testutil/nullify"
	"messageboard/x/messageboard"
	"messageboard/x/messageboard/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MessageboardKeeper(t)
	messageboard.InitGenesis(ctx, *k, genesisState)
	got := messageboard.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
