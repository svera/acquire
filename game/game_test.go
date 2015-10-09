package game

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	var players []*Player
	players = append(players, NewPlayer("Test"))

	var corporations [7]*Corporation
	corporations[0], _ = NewCorporation("A", 0)
	corporations[1], _ = NewCorporation("B", 0)
	corporations[2], _ = NewCorporation("C", 1)
	corporations[3], _ = NewCorporation("D", 1)
	corporations[4], _ = NewCorporation("E", 1)
	corporations[5], _ = NewCorporation("F", 2)
	corporations[6], _ = NewCorporation("G", 2)

	board := NewBoard()
	tileset := NewTileset()
	_, err := NewGame(board, players, corporations, tileset)

	if err == nil {
		t.Errorf("Game must not be created with less than 3 players, got %d", len(players))
	}
}

func TestNewGameInitsPlayersTilesets(t *testing.T) {
	var players []*Player
	players = append(players, NewPlayer("Test1"))
	players = append(players, NewPlayer("Test2"))
	players = append(players, NewPlayer("Test3"))

	var corporations [7]*Corporation
	corporations[0], _ = NewCorporation("A", 0)
	corporations[1], _ = NewCorporation("B", 0)
	corporations[2], _ = NewCorporation("C", 1)
	corporations[3], _ = NewCorporation("D", 1)
	corporations[4], _ = NewCorporation("E", 1)
	corporations[5], _ = NewCorporation("F", 2)
	corporations[6], _ = NewCorporation("G", 2)

	board := NewBoard()
	tileset := NewTileset()
	NewGame(board, players, corporations, tileset)

	for i, player := range players {
		if len(player.tiles) != 6 {
			t.Errorf("Players must have 6 tiles at the beginning, player %d got %d", i, len(player.tiles))
		}
	}
}
