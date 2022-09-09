package keeper

import (
	"context"
	"strings"

	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RejectGame(goCtx context.Context, msg *types.MsgRejectGame) (*types.MsgRejectGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.GameIndex)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "game not found %s", msg.GameIndex)
	}

	if strings.Compare(storedGame.Red, msg.Creator) == 0 {
		if 1 < storedGame.MoveCount { // Notice the use of the new field
			return nil, types.ErrRedAlreadyPlayed
		}
	} else if strings.Compare(storedGame.Black, msg.Creator) == 0 {
		if 0 < storedGame.MoveCount { // Notice the use of the new field
			return nil, types.ErrBlackAlreadyPlayed
		}
	} else {
		return nil, types.ErrCreatorNotPlayer
	}

	k.Keeper.RemoveStoredGame(ctx, msg.GameIndex)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.GameRejectedEventType,
			sdk.NewAttribute(types.GameRejectedEventCreator, msg.Creator),
			sdk.NewAttribute(types.GameRejectedEventGameIndex, msg.GameIndex),
		),
	)

	return &types.MsgRejectGameResponse{}, nil
}
