package game

import "testing"

func TestPutTile(t *testing.T) {
	board := NewBoard()
	tile := Tile{
		number: 5,
		letter: "B",
	}
	board.PutTile(tile)
	if board.grid[5]["B"] != boardCellUsed {
		t.Errorf("Tile %d%s was not put on the board", 5, "B")
	}
}
