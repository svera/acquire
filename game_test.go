package game

import (
	"github.com/svera/acquire/game/board"
	"github.com/svera/acquire/game/corporation"
	"github.com/svera/acquire/game/player"
	"github.com/svera/acquire/game/tileset"
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	players, corporations, board, tileset := setup()
	players = players[:1]

	_, err := New(board, players, corporations, tileset)

	if err == nil {
		t.Errorf("Game must not be created with less than 3 players, got %d", len(players))
	}
}

func TestNewGameInitsPlayersTilesets(t *testing.T) {
	players, corporations, board, tileset := setup()
	New(board, players, corporations, tileset)

	for i, player := range players {
		if len(player.Tiles()) != 6 {
			t.Errorf("Players must have 6 tiles at the beginning, player %d got %d", i, len(player.Tiles()))
		}
	}
}

func TestAreEndConditionsReached(t *testing.T) {
	players, corporations, board, tileset := setup()
	game, _ := New(board, players, corporations, tileset)

	if game.AreEndConditionsReached() {
		t.Errorf("End game conditions not reached (no active corporations) but detected as it were")
	}

	corporations[0].Size = func() int {
		return 41
	}

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (a corporation bigger than 40 tiles) but not detected")
	}

	corporations[0].Size = func() int {
		return 11
	}

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (all active corporations safe) but not detected")
	}

	corporations[0].Size = func() int {
		return 11
	}

	corporations[1].Size = func() int {
		return 2
	}

	if game.AreEndConditionsReached() {
		t.Errorf("End game conditions not reached but detected as it were")
	}

}

func TestGetMainStockHolders(t *testing.T) {
	players, corporations, board, tileset := setup()
	game, _ := New(board, players, corporations, tileset)
	players[0].Shares = func(c *corporation.Corporation) int {
		return 8
	}
	stockHolders := game.GetMainStockHolders(corporations[0])
	expectedStockHolders := map[string][]*player.Player{
		"majority": {players[0]},
		"minority": {players[0]},
	}
	if !reflect.DeepEqual(stockHolders, expectedStockHolders) {
		t.Errorf(
			"If there's just one player with stock in a defunct corporation, " +
				"he/she must get both majority and minority bonuses",
		)
	}

	players[1].Shares = func(c *corporation.Corporation) int {
		return 5
	}
	stockHolders = game.GetMainStockHolders(corporations[0])
	expectedStockHolders = map[string][]*player.Player{
		"majority": {players[0]},
		"minority": {players[1]},
	}
	if !reflect.DeepEqual(stockHolders, expectedStockHolders) {
		t.Errorf(
			"Wrong main stock holders",
		)
	}

	players[1].Shares = func(c *corporation.Corporation) int {
		return 8
	}
	players[2].Shares = func(c *corporation.Corporation) int {
		return 5
	}
	stockHolders = game.GetMainStockHolders(corporations[0])
	expectedStockHolders = map[string][]*player.Player{
		"majority": {players[0], players[1]},
		"minority": {},
	}
	if !reflect.DeepEqual(stockHolders, expectedStockHolders) {
		t.Errorf(
			"If there are two or more majority stock holders in a defunct corporation, " +
				"the majority bonus must be splitted between them (no minority bonus given)",
		)
	}

	players[1].Shares = func(c *corporation.Corporation) int {
		return 5
	}
	players[2].Shares = func(c *corporation.Corporation) int {
		return 5
	}
	stockHolders = game.GetMainStockHolders(corporations[0])
	expectedStockHolders = map[string][]*player.Player{
		"majority": {players[0]},
		"minority": {players[1], players[2]},
	}
	if !reflect.DeepEqual(stockHolders, expectedStockHolders) {
		t.Errorf(
			"If there are two or more minority stock holders in a defunct corporation, " +
				"the minority bonus must be splitted between them",
		)
	}
}

func setup() ([]*player.Player, [7]*corporation.Corporation, *board.Board, *tileset.Tileset) {
	var players []*player.Player
	players = append(players, player.New("Test1"))
	players = append(players, player.New("Test2"))
	players = append(players, player.New("Test3"))

	var corporations [7]*corporation.Corporation
	corporations[0], _ = corporation.New("A", 0)
	corporations[1], _ = corporation.New("B", 0)
	corporations[2], _ = corporation.New("C", 1)
	corporations[3], _ = corporation.New("D", 1)
	corporations[4], _ = corporation.New("E", 1)
	corporations[5], _ = corporation.New("F", 2)
	corporations[6], _ = corporation.New("G", 2)

	board := board.New()
	tileset := tileset.New()
	return players, corporations, board, tileset
}
