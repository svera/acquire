package board

import (
	"github.com/svera/acquire/interfaces"
	"github.com/svera/acquire/mocks"
	"reflect"
	"testing"
)

func TestPutTile(t *testing.T) {
	board := New()

	tile := &mocks.Tile{FakeNumber: 5, FakeLetter: "B"}
	board.PutTile(tile)
	if board.grid[5]["B"].Type() != "unincorporated" {
		t.Errorf("Position %d%s was not put on the board", 5, "B")
	}
}

func TestTileFoundCorporation(t *testing.T) {
	board := New()
	board.PutTile(&mocks.Tile{FakeNumber: 5, FakeLetter: "D"})
	board.PutTile(&mocks.Tile{FakeNumber: 6, FakeLetter: "C"})
	board.PutTile(&mocks.Tile{FakeNumber: 6, FakeLetter: "E"})
	board.PutTile(&mocks.Tile{FakeNumber: 7, FakeLetter: "D"})
	foundingTile := &mocks.Tile{FakeNumber: 6, FakeLetter: "D"}
	found, corporationTiles := board.TileFoundCorporation(
		foundingTile,
	)

	expectedCorporationTiles := []interfaces.Tile{
		foundingTile,
		board.grid[5]["D"].(interfaces.Tile),
		board.grid[6]["C"].(interfaces.Tile),
		board.grid[6]["E"].(interfaces.Tile),
		board.grid[7]["D"].(interfaces.Tile),
	}

	if !found {
		t.Errorf("TileFoundCorporation() must return true")
	}
	if !slicesSameCells(corporationTiles, expectedCorporationTiles) {
		t.Errorf("Position %d%s must found a corporation with tiles %v, got %v instead", 6, "D", expectedCorporationTiles, corporationTiles)
	}
}

func TestTileDoesNotFoundCorporation(t *testing.T) {
	board := New()
	found, corporationTiles := board.TileFoundCorporation(&mocks.Tile{FakeNumber: 6, FakeLetter: "D"})
	if found {
		t.Errorf("Position %d%s must not found a corporation, got %v instead", 6, "D", corporationTiles)
	}

	board.PutTile(&mocks.Tile{FakeNumber: 5, FakeLetter: "E"})
	board.grid[7]["E"] = &mocks.Corporation{}
	board.PutTile(&mocks.Tile{FakeNumber: 6, FakeLetter: "D"})
	board.PutTile(&mocks.Tile{FakeNumber: 6, FakeLetter: "F"})

	found, corporationTiles = board.TileFoundCorporation(&mocks.Tile{FakeNumber: 6, FakeLetter: "E"})
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
	corp1 := &mocks.Corporation{FakeSize: 4}
	corp2 := &mocks.Corporation{FakeSize: 5}
	corp3 := &mocks.Corporation{FakeSize: 3}
	corp4 := &mocks.Corporation{FakeSize: 2}

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

	expectedCorporations := map[string][]interfaces.Corporation{
		"acquirer": []interfaces.Corporation{corp2},
		"defunct":  []interfaces.Corporation{corp1, corp3, corp4},
	}
	merge, corporations := board.TileMergeCorporations(&mocks.Tile{FakeNumber: 6, FakeLetter: "E"})

	if !slicesSameCorporations(corporations["acquirer"], expectedCorporations["acquirer"]) ||
		!slicesSameCorporations(corporations["defunct"], expectedCorporations["defunct"]) {
		t.Errorf("Position %d%s must merge corporations %v, got %v instead", 6, "E", expectedCorporations, corporations)
	}
	if !merge {
		t.Errorf("TileMergeCorporations() must return true")
	}
}

// Testing quadruple merge tie as this:
//   4 5 6 7 8
// C     []
// D     []
// E [][]><[][]
// F     []
// G     []
func TestTileQuadrupleMergeTie(t *testing.T) {
	board := New()
	corp1 := &mocks.Corporation{FakeSize: 2}
	corp2 := &mocks.Corporation{FakeSize: 2}
	corp3 := &mocks.Corporation{FakeSize: 2}
	corp4 := &mocks.Corporation{FakeSize: 2}

	board.grid[4]["E"] = corp1
	board.grid[5]["E"] = corp1
	board.grid[7]["E"] = corp2
	board.grid[8]["E"] = corp2
	board.grid[6]["C"] = corp3
	board.grid[6]["D"] = corp3
	board.grid[6]["F"] = corp4
	board.grid[6]["G"] = corp4

	expectedCorporations := map[string][]interfaces.Corporation{
		"acquirer": []interfaces.Corporation{corp1, corp2, corp3, corp4},
		"defunct":  []interfaces.Corporation{},
	}
	merge, corporations := board.TileMergeCorporations(&mocks.Tile{FakeNumber: 6, FakeLetter: "E"})

	if !slicesSameCorporations(corporations["acquirer"], expectedCorporations["acquirer"]) ||
		!slicesSameCorporations(corporations["defunct"], expectedCorporations["defunct"]) {
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
	corp2 := &mocks.Corporation{}
	board.PutTile(&mocks.Tile{FakeNumber: 3, FakeLetter: "E"})
	board.grid[5]["E"] = corp2
	board.grid[6]["E"] = corp2

	expectedCorporationsMerged := map[string][]interfaces.Corporation{}
	merge, corporations := board.TileMergeCorporations(&mocks.Tile{FakeNumber: 4, FakeLetter: "E"})
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
	corp2 := &mocks.Corporation{}
	board.PutTile(&mocks.Tile{FakeNumber: 5, FakeLetter: "E"})
	board.grid[7]["E"] = corp2
	board.grid[8]["E"] = corp2
	board.PutTile(&mocks.Tile{FakeNumber: 6, FakeLetter: "D"})
	board.PutTile(&mocks.Tile{FakeNumber: 6, FakeLetter: "F"})
	growerTile := &mocks.Tile{FakeNumber: 6, FakeLetter: "E"}

	expectedTilesToAppend := []interfaces.Tile{
		board.grid[5]["E"].(interfaces.Tile),
		board.grid[6]["D"].(interfaces.Tile),
		growerTile,
		board.grid[6]["F"].(interfaces.Tile),
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
	corp2 := &mocks.Corporation{}

	board.grid[7]["E"] = corp2
	board.grid[8]["E"] = corp2

	grow, _, _ := board.TileGrowCorporation(&mocks.Tile{FakeNumber: 6, FakeLetter: "C"})
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
	tl := &mocks.Tile{FakeNumber: 1, FakeLetter: "A"}

	adjacentCells := brd.AdjacentCells(tl.Number(), tl.Letter())
	if len(adjacentCells) != 2 {
		t.Errorf(
			"Position %d%s expected to have adjacent 2 adjacent tiles, got %d",
			tl.Number(), tl.Letter(), len(adjacentCells),
		)
	}
}

func TestSetOwner(t *testing.T) {
	brd := New()
	corp := &mocks.Corporation{}
	tl1 := &mocks.Tile{FakeNumber: 1, FakeLetter: "A"}
	tl2 := &mocks.Tile{FakeNumber: 1, FakeLetter: "B"}
	tls := []interfaces.Tile{tl1, tl2}
	brd.SetOwner(corp, tls)
	if brd.Cell(tl1.Number(), tl1.Letter()) != corp || brd.Cell(tl2.Number(), tl2.Letter()) != corp {
		t.Errorf(
			"Cells %d%s and %d%s expected to belong to corporation",
			tl1.Number(), tl1.Letter(), tl2.Number(), tl2.Letter(),
		)
	}
}

// Compare coordinates of tiles from two slices, order independent
func slicesSameCells(slice1 []interfaces.Tile, slice2 []interfaces.Tile) bool {
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
func slicesSameCorporations(slice1 []interfaces.Corporation, slice2 []interfaces.Corporation) bool {
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
