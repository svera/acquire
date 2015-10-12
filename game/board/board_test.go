package board

import (
	"github.com/svera/acquire/game/tileset"
	"testing"
)

func TestPutTile(t *testing.T) {
	board := New()
	tile := tileset.Tile{Number: 5, Letter: "B"}
	board.PutTile(tile)
	if board.grid[5]["B"] != boardCellUsed {
		t.Errorf("Tile %d%s was not put on the board", 5, "B")
	}
}
