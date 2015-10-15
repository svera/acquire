package board

import (
	"github.com/svera/acquire/game/tileset"
)

const CellEmpty = -1
const CellUsed = 9

var letters [9]string

type Board struct {
	grid *[13]map[string]int
}

func init() {
	letters = [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
}

func New() *Board {
	board := Board{
		grid: new([13]map[string]int),
	}

	for number := 1; number < 13; number++ {
		board.grid[number] = make(map[string]int)
		for _, letter := range letters {
			board.grid[number][letter] = CellEmpty
		}
	}

	return &board
}

func (b *Board) Cell(t tileset.Tile) int {
	return b.grid[t.Number][t.Letter]
}

// Checks if the passed tile founds a new corporation, returns a slice of tiles
// composing this corporation
func (b *Board) TileFoundCorporation(t tileset.Tile) []tileset.Tile {
	var newCorporationTiles []tileset.Tile
	adjacent := b.AdjacentTiles(t)
	for _, adjacentTile := range adjacent {
		if b.Cell(adjacentTile) == CellUsed {
			newCorporationTiles = append(newCorporationTiles, adjacentTile)
		}
	}
	if len(newCorporationTiles) > 0 {
		newCorporationTiles = append(newCorporationTiles, t)
	}
	return newCorporationTiles
}

// Checks if the passed tile merges two or more corporations, returns a slice of
// corporation ids to be merged
func (b *Board) TileMergeCorporations(t tileset.Tile) []int {
	var mergedCorporations []int
	adjacent := b.AdjacentTiles(t)
	for _, adjacentTile := range adjacent {
		if b.Cell(adjacentTile) != CellEmpty && b.Cell(adjacentTile) != CellUsed {
			mergedCorporations = append(mergedCorporations, b.Cell(adjacentTile))
		}
	}
	return mergedCorporations
}

// Check if the passed tile grows a corporation
func (b *Board) TileGrowCorporation(t tileset.Tile) ([]tileset.Tile, int) {
	var tilesToAppend []tileset.Tile
	var corporationToGrow int = -1
	adjacentCorporations := 0
	adjacent := b.AdjacentTiles(t)
	for _, adjacentTile := range adjacent {
		if b.Cell(adjacentTile) != CellEmpty && b.Cell(adjacentTile) != CellUsed {
			adjacentCorporations++
			if adjacentCorporations == 2 {
				return []tileset.Tile{}, -1
			}
			corporationToGrow = b.Cell(adjacentTile)
		}
		if b.Cell(adjacentTile) == CellUsed {
			tilesToAppend = append(tilesToAppend, adjacentTile)
		}
	}
	if adjacentCorporations == 0 || len(tilesToAppend) == 0 {
		return []tileset.Tile{}, -1
	}
	tilesToAppend = append(tilesToAppend, t)
	return tilesToAppend, corporationToGrow
}

func (b *Board) PutTile(t tileset.Tile) {
	b.grid[t.Number][t.Letter] = CellUsed
}

func (b *Board) AdjacentTiles(t tileset.Tile) []tileset.Tile {
	var adjacent []tileset.Tile

	if t.Letter > "A" {
		adjacent = append(adjacent, tileset.Tile{Number: t.Number, Letter: previousLetter(t.Letter)})
	}
	if t.Letter < "I" {
		adjacent = append(adjacent, tileset.Tile{Number: t.Number, Letter: nextLetter(t.Letter)})
	}
	if t.Number > 1 {
		adjacent = append(adjacent, tileset.Tile{Number: t.Number - 1, Letter: t.Letter})
	}
	if t.Number < 13 {
		adjacent = append(adjacent, tileset.Tile{Number: t.Number + 1, Letter: t.Letter})
	}
	return adjacent
}

func getAdjacentLetter(letter string, delta int) string {
	for i, currentLetter := range letters {
		if currentLetter == letter {
			return letters[i+delta]
		}
	}
	return ""
}

func previousLetter(letter string) string {
	return getAdjacentLetter(letter, -1)
}

func nextLetter(letter string) string {
	return getAdjacentLetter(letter, +1)
}
