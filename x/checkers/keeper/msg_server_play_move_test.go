package keeper_test

import (
	"context"
	"testing"

	"github.com/alice/checkers/x/checkers"
	"github.com/alice/checkers/x/checkers/keeper"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)


func setupMsgServerWithOneGameForPlayMove(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
    k, ctx := setupKeeper(t)
    checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
    server := keeper.NewMsgServerImpl(*k)
    context := sdk.WrapSDKContext(ctx)
    server.CreateGame(context, &types.MsgCreateGame{
        Creator: alice,
        Red:     bob,
        Black:   carol,
    })
    return server, *k, context
}

func TestPlayMove(t *testing.T) {
    msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t)
    playMoveResponse, err := msgServer.PlayMove(context, &types.MsgPlayMove{
        Creator: carol,
        IdValue: "1",
        FromX:   1,
        FromY:   2,
        ToX:     2,
        ToY:     3,
    })
    require.Nil(t, err)
    require.EqualValues(t, types.MsgPlayMoveResponse{
        IdValue:   "1",
        CapturedX: -1,
        CapturedY: -1,
        Winner:    "*",
    }, *playMoveResponse)
}