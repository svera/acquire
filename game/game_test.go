package game

import (
	"testing"
	"github.com/svera/acquire/game/board"
	"github.com/svera/acquire/game/corporation"
	"github.com/svera/acquire/game/player"
	"github.com/svera/acquire/game/tileset"
)

func TestNewGame(t *testing.T) {
	var players []*player.Player
	players = append(players, player.New("Test"))

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
	_, err := New(board, players, corporations, tileset)

	if err == nil {
		t.Errorf("Game must not be created with less than 3 players, got %d", len(players))
	}
}

func TestNewGameInitsPlayersTilesets(t *testing.T) {
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
	New(board, players, corporations, tileset)

	for i, player := range players {
		if len(player.Tiles()) != 6 {
			t.Errorf("Players must have 6 tiles at the beginning, player %d got %d", i, len(player.Tiles()))
		}
	}
}
