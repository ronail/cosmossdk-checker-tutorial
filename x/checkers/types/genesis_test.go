package types

import (
	"testing"

	"github.com/alice/checkers/x/checkers/rules"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func GetStoredGame1() *StoredGame {
	return &StoredGame{
		Black: bob,
		Red:   carol,
		Index: "1",
		Game:  rules.New().String(),
		Turn:  "b",
	}
}

func TestDefaultGenesisIsCorrect(t *testing.T) {
	require.EqualValues(t,
		&GenesisState{
			SystemInfo: &SystemInfo{
				NextId: uint64(DefaultIndex),
			},
			StoredGameList: []*StoredGame{},
			NextGame:       &NextGame{"", uint64(1)},
		},
		DefaultGenesis())
}
