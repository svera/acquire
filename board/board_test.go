package board

import (
	"github.com/svera/acquire/corporation"
	"reflect"
	//"sort"
	"github.com/svera/acquire/tile"
	"testing"
)

func TestPutTile(t *testing.T) {
	board := New()
	tile := tile.New(5, "B", tile.Orphan{})
	board.PutTile(tile)
	if board.grid[5]["B"].Content().Type() != "orphan" {
		t.Errorf("Position %d%s was not put on the board", 5, "B")
	}
}

func TestTileFoundCorporation(t *testing.T) {
	board := New()
	board.grid[5]["D"] = tile.New(5, "D", tile.Orphan{})
	board.grid[6]["C"] = tile.New(6, "C", tile.Orphan{})
	board.grid[6]["E"] = tile.New(6, "E", tile.Orphan{})
	board.grid[7]["D"] = tile.New(7, "D", tile.Orphan{})
	foundingTile := tile.New(6, "D", tile.Orphan{})
	found, corporationTiles := board.TileFoundCorporation(
		foundingTile,
	)

	expectedCorporationTiles := []tile.Interface{
		foundingTile,
		board.grid[5]["D"],
		board.grid[6]["C"],
		board.grid[6]["E"],
		board.grid[7]["D"],
	}

	if !found {
		t.Errorf("TileFoundCorporation() must return true")
	}
	if !slicesSameCells(corporationTiles, expectedCorporationTiles) {
		t.Errorf("Position %d%s must found a corporation with tiles %v, got %v instead", 6, "D", expectedCorporationTiles, corporationTiles)
	}
}

func TestTileNotFoundCorporation(t *testing.T) {
	board := New()
	corp, _ := corporation.New("Test 1", 0)
	found, corporationTiles := board.TileFoundCorporation(tile.New(6, "D", tile.Orphan{}))
	if found {
		t.Errorf("Position %d%s must not found a corporation, got %v instead", 6, "D", corporationTiles)
	}

	board.grid[5]["E"] = tile.New(5, "E", tile.Orphan{})
	board.grid[7]["E"] = tile.New(5, "E", corp)
	board.grid[6]["D"] = tile.New(6, "D", tile.Orphan{})
	board.grid[6]["F"] = tile.New(6, "F", tile.Orphan{})

	found, corporationTiles = board.TileFoundCorporation(tile.New(6, "E", tile.Orphan{}))
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
	corp1, _ := corporation.New("Test 1", 0)
	corp2, _ := corporation.New("Test 2", 1)
	corp3, _ := corporation.New("Test 3", 2)
	corp4, _ := corporation.New("Test 3", 2)
	board.grid[2]["E"] = tile.New(2, "E", corp1)
	board.grid[3]["E"] = tile.New(3, "E", corp1)
	board.grid[4]["E"] = tile.New(4, "E", corp1)
	board.grid[5]["E"] = tile.New(5, "E", corp1)
	board.grid[7]["E"] = tile.New(7, "E", corp2)
	board.grid[8]["E"] = tile.New(8, "E", corp2)
	board.grid[9]["E"] = tile.New(9, "E", corp2)
	board.grid[10]["E"] = tile.New(10, "E", corp2)
	board.grid[11]["E"] = tile.New(11, "E", corp2)
	board.grid[6]["B"] = tile.New(6, "B", corp3)
	board.grid[6]["C"] = tile.New(6, "C", corp3)
	board.grid[6]["D"] = tile.New(6, "D", corp3)
	board.grid[6]["F"] = tile.New(6, "F", corp4)
	board.grid[6]["G"] = tile.New(6, "G", corp4)

	expectedCorporations := []corporation.Interface{corp1, corp2, corp3, corp4}
	merge, corporations := board.TileMergeCorporations(tile.New(6, "E", tile.Orphan{}))
	//sort.Ints(corporationIds)
	if !slicesSameCorporations(corporations, expectedCorporations) {
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
	corp2, _ := corporation.New("Test 2", 1)
	board.grid[3]["E"] = tile.New(3, "E", tile.Orphan{})
	board.grid[5]["E"] = tile.New(5, "E", corp2)
	board.grid[6]["E"] = tile.New(6, "E", corp2)

	expectedCorporationsMerged := []corporation.Interface{}
	merge, corporations := board.TileMergeCorporations(tile.New(4, "E", tile.Orphan{}))
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
	corp2, _ := corporation.New("Test 2", 1)
	board.grid[5]["E"] = tile.New(5, "E", tile.Orphan{})
	board.grid[7]["E"] = tile.New(7, "E", corp2)
	board.grid[8]["E"] = tile.New(8, "E", corp2)
	board.grid[6]["D"] = tile.New(6, "D", tile.Orphan{})
	board.grid[6]["F"] = tile.New(6, "F", tile.Orphan{})
	growerTile := tile.New(6, "E", tile.Orphan{})

	expectedTilesToAppend := []tile.Interface{
		board.grid[5]["E"],
		board.grid[6]["D"],
		growerTile,
		board.grid[7]["E"],
	}
	expectedCorporationToGrow := corp2
	grow, tilesToAppend, corporationToGrow := board.TileGrowCorporation(growerTile)
	if !slicesSameCells(tilesToAppend, expectedTilesToAppend) {
		t.Errorf(
			"Position %d%s must grow corporation %s by %v, got %v in corporation %s instead",
			6,
			"E",
			expectedCorporationToGrow.Name(),
			expectedTilesToAppend,
			tilesToAppend,
			corporationToGrow.Name(),
		)
	}
	if !grow {
		t.Errorf("TileGrowCorporation() must return true")
	}
}

func TestTileDontGrowCorporation(t *testing.T) {
	board := New()
	corp2, _ := corporation.New("Test 2", 1)

	board.grid[7]["E"] = tile.New(7, "E", corp2)
	board.grid[8]["E"] = tile.New(8, "E", corp2)

	grow, _, _ := board.TileGrowCorporation(tile.New(6, "C", tile.Orphan{}))
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
	tl := tile.New(1, "A", tile.Orphan{})
	expectedAdjacentCells := []tile.Interface{
		tile.New(2, "A", tile.Orphan{}),
		tile.New(1, "B", tile.Orphan{}),
	}

	adjacentCells := brd.AdjacentCells(tl)
	if !slicesSameCells(adjacentCells, expectedAdjacentCells) {
		t.Errorf(
			"Position %d%s expected to have adjacent tiles %v, got %v",
			tl.Number, tl.Letter, expectedAdjacentCells, adjacentCells,
		)
	}
}

func TestSetTiles(t *testing.T) {
	brd := New()
	corp, _ := corporation.New("Test", 1)
	tl1 := tile.New(1, "A", corp)
	tl2 := tile.New(1, "B", corp)
	tls := []tile.Interface{tl1, tl2}
	brd.SetTiles(corp, tls)
	if brd.Cell(tl1.Number(), tl1.Letter()).Content() != corp || brd.Cell(tl2.Number(), tl2.Letter()).Content() != corp {
		t.Errorf(
			"Cells %d%s and %d%s expected to belong to corporation",
			tl1.Number(), tl1.Letter(), tl2.Number(), tl2.Letter(),
		)
	}
}

// Compare coordinates of tiles from two slices, order independent
func slicesSameCells(slice1 []tile.Interface, slice2 []tile.Interface) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	var inSlice bool
	for _, val1 := range slice1 {
		inSlice = false
		for _, val2 := range slice2 {
			if val1.Number() == val2.Number() && val1.Letter() == val2.Letter() {
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

// Compare corporations from two slices, order independent
func slicesSameCorporations(slice1 []corporation.Interface, slice2 []corporation.Interface) bool {
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
