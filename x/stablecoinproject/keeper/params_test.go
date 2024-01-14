package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "github.com/jhonnik08/stablecoinproject/testutil/keeper"
    "github.com/jhonnik08/stablecoinproject/x/stablecoinproject/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.StablecoinprojectKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
