package board

import (
	"github.com/svera/acquire/game/tileset"
	"reflect"
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

func TestCorporationFounded(t *testing.T) {
	board := New()
	board.grid[5]["D"] = boardCellUsed
	board.grid[6]["C"] = boardCellUsed
	board.grid[6]["E"] = boardCellUsed
	board.grid[7]["D"] = boardCellUsed
	corporationTiles := board.CorporationFounded(tileset.Tile{Number: 6, Letter: "D"})
	expectedCorporationTiles := []tileset.Tile{
		tileset.Tile{Number: 6, Letter: "D"},
		tileset.Tile{Number: 5, Letter: "D"},
		tileset.Tile{Number: 6, Letter: "C"},
		tileset.Tile{Number: 6, Letter: "E"},
		tileset.Tile{Number: 7, Letter: "D"},
	}
	if reflect.DeepEqual(corporationTiles, expectedCorporationTiles) {
		t.Errorf("Tile %d%s must found a corporation with tiles %v, got %v instead", 6, "D", expectedCorporationTiles, corporationTiles)
	}
}

func TestCorporationNotFounded(t *testing.T) {
	board := New()
	corporationTiles := board.CorporationFounded(tileset.Tile{Number: 6, Letter: "D"})
	expectedCorporationTiles := []tileset.Tile{}
	if reflect.DeepEqual(corporationTiles, expectedCorporationTiles) {
		t.Errorf("Tile %d%s must not found a corporation, got %v instead", 6, "D", expectedCorporationTiles, corporationTiles)
	}
}
