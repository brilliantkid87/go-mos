package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "messageboard/testutil/keeper"
	"messageboard/x/messageboard/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MessageboardKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
