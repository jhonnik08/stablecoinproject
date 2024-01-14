package stablecoinproject_test

import (
	"testing"

	keepertest "github.com/jhonnik08/stablecoinproject/testutil/keeper"
	"github.com/jhonnik08/stablecoinproject/testutil/nullify"
	"github.com/jhonnik08/stablecoinproject/x/stablecoinproject/module"
	"github.com/jhonnik08/stablecoinproject/x/stablecoinproject/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StablecoinprojectKeeper(t)
	stablecoinproject.InitGenesis(ctx, k, genesisState)
	got := stablecoinproject.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
