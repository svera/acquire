package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tileset"
	"reflect"
	//"sort"
	"github.com/svera/acquire/tile"
	"testing"
)

func TestPutTile(t *testing.T) {
	board := New()
	tile := &tileset.Tile{{Number: 5, Letter: "B"}}
	board.PutTile(tile)
	if board.grid[5]["B"].ContentType() != "orphan" {
		t.Errorf("Position %d%s was not put on the board", 5, "B")
	}
}

func TestTileFoundCorporation(t *testing.T) {
	board := New()
	board.grid[5]["D"] = tile.New(5, "D")
	board.grid[6]["C"] = tile.New(6, "C")
	board.grid[6]["E"] = tile.New(6, "E")
	board.grid[7]["D"] = tile.New(7, "D")
	found, corporationTiles := board.TileFoundCorporation(
		board.Coordinates{Number: 6, Letter: "D"},
	)

	expectedCorporationTiles := []board.Coordinates{
		board.Coordinates{Number: 6, Letter: "D"},
		board.Coordinates{Number: 5, Letter: "D"},
		board.Coordinates{Number: 6, Letter: "C"},
		board.Coordinates{Number: 6, Letter: "E"},
		board.Coordinates{Number: 7, Letter: "D"},
	}

	if !found {
		t.Errorf("TileFoundCorporation() must return true")
	}
	if !slicesSameContent(corporationTiles, expectedCorporationTiles) {
		t.Errorf("Position %d%s must found a corporation with tiles %v, got %v instead", 6, "D", expectedCorporationTiles, corporationTiles)
	}
}

func TestTileNotFoundCorporation(t *testing.T) {
	board := New()
	found, corporationTiles := board.TileFoundCorporation(board.Coordinates{Number: 6, Letter: "D"})
	if found {
		t.Errorf("Position %d%s must not found a corporation, got %v instead", 6, "D", corporationTiles)
	}

	board.grid[5]["E"] = tile.New(5, "E")
	board.grid[7]["E"] = 2
	board.grid[6]["D"] = tile.New(6, "D")
	board.grid[6]["F"] = tile.New(6, "F")

	found, corporationTiles = board.TileFoundCorporation(board.Coordinates{Number: 6, Letter: "E"})
	if found {
		t.Errorf("Position %d%s must not found a corporation, got %v instead", 6, "E", corporationTiles)
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
func TestTileQuadrupleMerge(t *testing.T) {
	board := New()
	corp1 := corporation.New("Test 1", 0, 1)
	corp2 := corporation.New("Test 2", 1, 2)
	corp3 := corporation.New("Test 3", 2, 3)
	board.grid[2]["E"] = corp1
	board.grid[3]["E"] = corp1
	board.grid[4]["E"] = corp1
	board.grid[5]["E"] = corp1
	board.grid[7]["E"] = corp2
	board.grid[8]["E"] = corp2
	board.grid[9]["E"] = corp2
	board.grid[10]["E"] = corp2
	board.grid[11]["E"] = corp2
	board.grid[6]["B"] = corp3
	board.grid[6]["C"] = corp3
	board.grid[6]["D"] = corp3
	board.grid[6]["F"] = corp4
	board.grid[6]["G"] = corp4

	expectedCorporations := []corporation.Interface{corp1, corp2, corp3, corp4}
	merge, corporations := board.TileMergeCorporations(board.Coordinates{Number: 6, Letter: "E"})
	//sort.Ints(corporationIds)
	if !reflect.DeepEqual(corporations, expectedCorporations) {
		t.Errorf("Position %d%s must merge corporations %v, got %v instead", 6, "E", expectedCorporations, corporations)
	}
	if !merge {
		t.Errorf("TileMergeCorporations() must return true")
	}
}

// Testing not a merge as this:
//   3 4 5 6
// E []><[][]
func TestTileDontMerge(t *testing.T) {
	board := New()
	corp2 := corporation.New("Test 2", 1, 2)
	board.grid[3]["E"] = tile.New(3, "E")
	board.grid[5]["E"] = corp2
	board.grid[6]["E"] = corp2

	expectedCorporationsMerged := []corporation.Interface{}
	merge, corporations := board.TileMergeCorporations(board.Coordinates{Number: 4, Letter: "E"})
	if !reflect.DeepEqual(corporations, expectedCorporationsMerged) {
		t.Errorf("Position %d%s must not merge corporations, got %v instead", 4, "E", corporations)
	}
	if merge {
		t.Errorf("TileMergeCorporations() must return false")
	}
}

// Testing growing corporation as this:
//   5 6 7 8
// D   []
// E []><[][]
// F   []
func TestTileGrowCorporation(t *testing.T) {
	board := New()
	corp2 := corporation.New("Test 2", 1, 2)
	board.grid[5]["E"] = tile.New(5, "E")
	board.grid[7]["E"] = corp2
	board.grid[8]["E"] = corp2
	board.grid[6]["D"] = tile.New(6, "D")
	board.grid[6]["F"] = tile.New(6, "F")

	expectedTilesToAppend := []board.Coordinates{
		board.Coordinates{Number: 5, Letter: "E"},
		board.Coordinates{Number: 6, Letter: "D"},
		board.Coordinates{Number: 6, Letter: "E"},
		board.Coordinates{Number: 6, Letter: "F"},
	}
	expectedCorporationToGrow := corp2
	grow, tilesToAppend, corporationToGrow := board.TileGrowCorporation(board.Coordinates{Number: 6, Letter: "E"})
	if !slicesSameContent(tilesToAppend, expectedTilesToAppend) {
		t.Errorf(
			"Position %d%s must grow corporation %d by %v, got %v in corporation %d instead",
			6,
			"E",
			expectedCorporationToGrow,
			expectedTilesToAppend,
			tilesToAppend,
			corporationToGrow,
		)
	}
	if !grow {
		t.Errorf("TileGrowCorporation() must return true")
	}
}

func TestTileDontGrowCorporation(t *testing.T) {
	board := New()
	corp2 := corporation.New("Test 2", 1, 2)

	board.grid[7]["E"] = corp2
	board.grid[8]["E"] = corp2

	grow, _, _ := board.TileGrowCorporation(board.Coordinates{Number: 6, Letter: "C"})
	if grow {
		t.Errorf(
			"Position %d%s must not grow any corporation, but got true",
			6,
			"C",
		)
	}
}

func TestAdjacentCells(t *testing.T) {
	brd := New()
	position := board.Coordinates{Number: 1, Letter: "A"}
	expectedAdjacentCells := []board.Coordinates{
		{Number: 2, Letter: "A"},
		{Number: 1, Letter: "B"},
	}

	adjacentCells := brd.AdjacentCells(position)
	if !slicesSameContent(adjacentCells, expectedAdjacentCells) {
		t.Errorf(
			"Position %d%s expected to have adjacent tiles %v, got %v",
			position.Number, position.Letter, expectedAdjacentCells, adjacentCells,
		)
	}
}

func TestSetTiles(t *testing.T) {
	brd := New()
	corp, _ := corporation.New("Test", 1, 5)
	cell1 := board.Coordinates{Number: 1, Letter: "A"}
	cell2 := board.Coordinates{Number: 1, Letter: "B"}
	cells := []board.Coordinates{cell1, cell2}
	brd.SetTiles(corp, cells)
	if brd.Cell(cell1) != 5 || brd.Cell(cell2) != 5 {
		t.Errorf(
			"Cells %d%s and %d%s expected to belong to corporation %d",
			cell1.Number, cell1.Letter, cell2.Number, cell2.Letter, corp.Id(),
		)
	}
}

func slicesSameContent(slice1 []board.Coordinates, slice2 []board.Coordinates) bool {
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
