package board

import (
	"github.com/svera/acquire/game/tileset"
	"reflect"
	"sort"
	"testing"
)

func TestPutTile(t *testing.T) {
	board := New()
	tile := tileset.Tile{Number: 5, Letter: "B"}
	board.PutTile(tile)
	if board.grid[5]["B"] != CellUsed {
		t.Errorf("Tile %d%s was not put on the board", 5, "B")
	}
}

func TestTileFoundCorporation(t *testing.T) {
	board := New()
	board.grid[5]["D"] = CellUsed
	board.grid[6]["C"] = CellUsed
	board.grid[6]["E"] = CellUsed
	board.grid[7]["D"] = CellUsed
	corporationTiles := board.TileFoundCorporation(tileset.Tile{Number: 6, Letter: "D"})
	expectedCorporationTiles := []tileset.Tile{
		tileset.Tile{Number: 6, Letter: "D"},
		tileset.Tile{Number: 5, Letter: "D"},
		tileset.Tile{Number: 6, Letter: "C"},
		tileset.Tile{Number: 6, Letter: "E"},
		tileset.Tile{Number: 7, Letter: "D"},
	}
	if !slicesContentEquals(corporationTiles, expectedCorporationTiles) {
		t.Errorf("Tile %d%s must found a corporation with tiles %v, got %v instead", 6, "D", expectedCorporationTiles, corporationTiles)
	}
}

func TestTileNotFoundCorporation(t *testing.T) {
	board := New()
	corporationTiles := board.TileFoundCorporation(tileset.Tile{Number: 6, Letter: "D"})
	if len(corporationTiles) != 0 {
		t.Errorf("Tile %d%s must not found a corporation, got %v instead", 6, "D", corporationTiles)
	}
}

// Testing quadruple merge as this:
//   2 3 4 5 6 7 8 9 1011
// B         []
// C         []
// D         []
// E [][][][]><[][][][][]
// F         []
// G         []
func TestTileMergeCorporations(t *testing.T) {
	board := New()
	board.grid[2]["E"] = 1
	board.grid[3]["E"] = 1
	board.grid[4]["E"] = 1
	board.grid[5]["E"] = 1
	board.grid[7]["E"] = 2
	board.grid[8]["E"] = 2
	board.grid[9]["E"] = 2
	board.grid[10]["E"] = 2
	board.grid[11]["E"] = 2
	board.grid[6]["B"] = 3
	board.grid[6]["C"] = 3
	board.grid[6]["D"] = 3
	board.grid[6]["F"] = 4
	board.grid[6]["G"] = 4

	expectedCorporationsMerged := []int{1, 2, 3, 4}
	corporationsMerged := board.TileMergeCorporations(tileset.Tile{Number: 6, Letter: "E"})
	sort.Ints(corporationsMerged)
	if !reflect.DeepEqual(corporationsMerged, expectedCorporationsMerged) {
		t.Errorf("Tile %d%s must merge corporations %v, got %v instead", 6, "E", expectedCorporationsMerged, corporationsMerged)
	}
}

func slicesContentEquals(slice1 []tileset.Tile, slice2 []tileset.Tile) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	var inSlice bool
	for _, val1 := range slice1 {
		inSlice = false
		for _, val2 := range slice2 {
			if val1 == val2 {
				inSlice = true
				break
			}
		}
		if !inSlice {
			return false
		}
	}
	return true
}
