package keeper_test

import (
	"context"
	"testing"

	"github.com/alice/checkers/x/checkers"
	"github.com/alice/checkers/x/checkers/types"
    "github.com/alice/checkers/x/checkers/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
    alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
    bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
    carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func TestCreateGame(t *testing.T) {
    msgServer, context := setupMsgServerCreateGame(t)
    createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
        Creator: alice,
        Red:     bob,
        Black:   carol,
    })
    require.Nil(t, err)
    require.EqualValues(t, types.MsgCreateGameResponse{
        IdValue: "1",
    }, *createResponse)
}

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, context.Context) {
    k, ctx := setupKeeper(t)
    checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
    return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}