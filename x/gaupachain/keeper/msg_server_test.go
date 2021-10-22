package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/GaupaLabs/gaupachain/testutil/keeper"
	"github.com/GaupaLabs/gaupachain/x/gaupachain/keeper"
	"github.com/GaupaLabs/gaupachain/x/gaupachain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.GaupachainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
