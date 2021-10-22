package gaupachain_test

import (
	"testing"

	keepertest "github.com/GaupaLabs/gaupachain/testutil/keeper"
	"github.com/GaupaLabs/gaupachain/x/gaupachain"
	"github.com/GaupaLabs/gaupachain/x/gaupachain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GaupachainKeeper(t)
	gaupachain.InitGenesis(ctx, *k, genesisState)
	got := gaupachain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
