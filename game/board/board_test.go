package board

import (
	"github.com/svera/acquire/game/tileset"
	"reflect"
	"sort"
	"testing"
)

func TestPutTile(t *testing.T) {
	board := New()
	tile := tileset.Position{Number: 5, Letter: "B"}
	board.PutTile(tile)
	if board.grid[5]["B"] != OrphanTile {
		t.Errorf("Position %d%s was not put on the board", 5, "B")
	}
}

func TestTileFoundCorporation(t *testing.T) {
	board := New()
	board.grid[5]["D"] = OrphanTile
	board.grid[6]["C"] = OrphanTile
	board.grid[6]["E"] = OrphanTile
	board.grid[7]["D"] = OrphanTile
	found, corporationTiles := board.TileFoundCorporation(
		tileset.Position{Number: 6, Letter: "D"},
	)

	expectedCorporationTiles := []tileset.Position{
		tileset.Position{Number: 6, Letter: "D"},
		tileset.Position{Number: 5, Letter: "D"},
		tileset.Position{Number: 6, Letter: "C"},
		tileset.Position{Number: 6, Letter: "E"},
		tileset.Position{Number: 7, Letter: "D"},
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
	found, corporationTiles := board.TileFoundCorporation(tileset.Position{Number: 6, Letter: "D"})
	if found {
		t.Errorf("Position %d%s must not found a corporation, got %v instead", 6, "D", corporationTiles)
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

	expectedCorporationsIds := []int{1, 2, 3, 4}
	merge, corporationIds := board.TileMergeCorporations(tileset.Position{Number: 6, Letter: "E"})
	sort.Ints(corporationIds)
	if !reflect.DeepEqual(corporationIds, expectedCorporationsIds) {
		t.Errorf("Position %d%s must merge corporations %v, got %v instead", 6, "E", expectedCorporationsIds, corporationIds)
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
	board.grid[3]["E"] = OrphanTile
	board.grid[5]["E"] = 2
	board.grid[6]["E"] = 2

	expectedCorporationsMerged := []int{}
	merge, corporationIds := board.TileMergeCorporations(tileset.Position{Number: 4, Letter: "E"})
	sort.Ints(corporationIds)
	if !reflect.DeepEqual(corporationIds, expectedCorporationsMerged) {
		t.Errorf("Position %d%s must not merge corporations, got %v instead", 4, "E", corporationIds)
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
	board.grid[5]["E"] = OrphanTile
	board.grid[7]["E"] = 2
	board.grid[8]["E"] = 2
	board.grid[6]["D"] = OrphanTile
	board.grid[6]["F"] = OrphanTile

	expectedTilesToAppend := []tileset.Position{
		tileset.Position{Number: 5, Letter: "E"},
		tileset.Position{Number: 6, Letter: "D"},
		tileset.Position{Number: 6, Letter: "E"},
		tileset.Position{Number: 6, Letter: "F"},
	}
	expectedCorporationToGrow := 2
	grow, tilesToAppend, corporationToGrow := board.TileGrowCorporation(tileset.Position{Number: 6, Letter: "E"})
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

func slicesSameContent(slice1 []tileset.Position, slice2 []tileset.Position) bool {
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
