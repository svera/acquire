package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tileset"
)

const Empty = -1
const OrphanTile = 9

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
			board.grid[number][letter] = Empty
		}
	}

	return &board
}

// Returns a board cell value, which can be -1 if that cell doesn't have
// any tile, 9 if it has an orphan tile or the ID of a corporation otherwise
func (b *Board) Cell(t tileset.Position) int {
	return b.grid[t.Number][t.Letter]
}

// Checks if the passed tile founds a new corporation, returns a slice of tiles
// composing this corporation
func (b *Board) TileFoundCorporation(t tileset.Position) (bool, []tileset.Position) {
	var newCorporationTiles []tileset.Position
	adjacent := b.adjacentOrphanTiles(t)
	for _, adjacentCell := range adjacent {
		newCorporationTiles = append(newCorporationTiles, adjacentCell)
	}
	if len(newCorporationTiles) > 0 {
		newCorporationTiles = append(newCorporationTiles, t)
		return true, newCorporationTiles
	}
	return false, newCorporationTiles
}

// Checks if the passed tile merges two or more corporations, returns a slice of
// corporation IDs to be merged
func (b *Board) TileMergeCorporations(t tileset.Position) (bool, []int) {
	var corporations []int
	adjacent := b.adjacentCorporationTiles(t)
	for _, adjacentCell := range adjacent {
		corporations = append(corporations, b.Cell(adjacentCell))
	}
	if len(corporations) > 1 {
		return true, corporations
	}
	return false, []int{}
}

// Check if the passed tile grows a corporation
func (b *Board) TileGrowCorporation(t tileset.Position) (bool, []tileset.Position, int) {
	tilesToAppend := []tileset.Position{{Number: t.Number, Letter: t.Letter}}
	corporationToGrow := -1
	adjacent := b.adjacentTiles(t)
	for _, adjacentCell := range adjacent {
		if b.Cell(adjacentCell) != OrphanTile {
			if corporationToGrow != -1 {
				return false, []tileset.Position{}, -1
			}
			corporationToGrow = b.Cell(adjacentCell)
		} else {
			tilesToAppend = append(tilesToAppend, adjacentCell)
		}
	}
	if corporationToGrow == -1 {
		return false, []tileset.Position{}, -1
	}
	return true, tilesToAppend, corporationToGrow
}

// Puts the passed tile on the board
func (b *Board) PutTile(t tileset.Position) {
	b.grid[t.Number][t.Letter] = OrphanTile
}

// Returns all cells adjacent to the passed one
func (b *Board) AdjacentCells(t tileset.Position) []tileset.Position {
	var adjacent []tileset.Position

	if t.Letter > "A" {
		adjacent = append(adjacent, tileset.Position{Number: t.Number, Letter: previousLetter(t.Letter)})
	}
	if t.Letter < "I" {
		adjacent = append(adjacent, tileset.Position{Number: t.Number, Letter: nextLetter(t.Letter)})
	}
	if t.Number > 1 {
		adjacent = append(adjacent, tileset.Position{Number: t.Number - 1, Letter: t.Letter})
	}
	if t.Number < 13 {
		adjacent = append(adjacent, tileset.Position{Number: t.Number + 1, Letter: t.Letter})
	}
	return adjacent
}

func (b *Board) adjacentCellsWithFilter(t tileset.Position, filter func(tileset.Position) bool) []tileset.Position {
	var adjacentFilteredCells []tileset.Position
	adjacent := b.AdjacentCells(t)

	for _, adjacentCell := range adjacent {
		if filter(adjacentCell) {
			adjacentFilteredCells = append(adjacentFilteredCells, adjacentCell)
		}
	}
	return adjacentFilteredCells
}

func (b *Board) adjacentTiles(t tileset.Position) []tileset.Position {
	return b.adjacentCellsWithFilter(
		t,
		func(t tileset.Position) bool {
			if b.Cell(t) != Empty {
				return true
			}
			return false
		},
	)
}

func (b *Board) adjacentCorporationTiles(t tileset.Position) []tileset.Position {
	return b.adjacentCellsWithFilter(
		t,
		func(t tileset.Position) bool {
			if b.Cell(t) != Empty && b.Cell(t) != OrphanTile {
				return true
			}
			return false
		},
	)
}

func (b *Board) adjacentOrphanTiles(t tileset.Position) []tileset.Position {
	return b.adjacentCellsWithFilter(
		t,
		func(t tileset.Position) bool {
			if b.Cell(t) == OrphanTile {
				return true
			}
			return false
		},
	)
}

func adjacentLetter(letter string, delta int) string {
	for i, currentLetter := range letters {
		if currentLetter == letter {
			return letters[i+delta]
		}
	}
	return ""
}

func previousLetter(letter string) string {
	return adjacentLetter(letter, -1)
}

func nextLetter(letter string) string {
	return adjacentLetter(letter, +1)
}

func (b *Board) SetCells(cp *corporation.Corporation, cells []tileset.Position) {
	for _, cell := range cells {
		b.grid[cell.Number][cell.Letter] = cp.Id()
	}
}
