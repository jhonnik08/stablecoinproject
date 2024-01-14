package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "github.com/jhonnik08/stablecoinproject/testutil/keeper"
    "github.com/jhonnik08/stablecoinproject/x/stablecoinproject/types"
    "github.com/jhonnik08/stablecoinproject/x/stablecoinproject/keeper"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.StablecoinprojectKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}