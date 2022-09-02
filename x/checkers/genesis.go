package checkers

import (
	"github.com/alice/checkers/x/checkers/keeper"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the storedGame
	for _, elem := range genState.StoredGameList {
		k.SetStoredGame(ctx, *elem)
	}

	// Set if defined
	if genState.NextGame != nil {
		k.SetNextGame(ctx, *genState.NextGame)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all storedGame
	storedGameList := k.GetAllStoredGame(ctx)
	for _, elem := range storedGameList {
		elem := elem
		genesis.StoredGameList = append(genesis.StoredGameList, &elem)
	}

	// Get all nextGame
	nextGame, found := k.GetNextGame(ctx)
	if found {
		genesis.NextGame = &nextGame
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
