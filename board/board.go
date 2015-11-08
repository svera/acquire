// Model for the game board, storing its state and implementing related
// actions.
package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tileset"
)

var letters [9]string = [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

type Board struct {
	grid *[13]map[string]Container
}

func New() *Board {
	board := Board{
		grid: new([13]map[string]Container),
	}

	emptyCell := &Empty{}
	for number := 1; number < 13; number++ {
		board.grid[number] = make(map[string]Container)
		for _, letter := range letters {
			board.grid[number][letter] = emptyCell
		}
	}

	return &board
}

// Returns a board cell content
func (b *Board) Cell(t tileset.Position) Container {
	return b.grid[t.Number][t.Letter]
}

// Checks if the passed tile founds a new corporation, returns a slice of tiles
// composing this corporation
func (b *Board) TileFoundCorporation(t tileset.Position) (bool, []tileset.Position) {
	var newCorporationTiles []tileset.Position
	adjacent := b.adjacentNonCorporationTiles(t)
	if len(adjacent) == 4 {
		for _, adjacentCell := range adjacent {
			if b.Cell(adjacentCell).ContentType() == "orphan" {
				newCorporationTiles = append(newCorporationTiles, adjacentCell)
			}
		}
		if len(newCorporationTiles) > 0 {
			newCorporationTiles = append(newCorporationTiles, t)
			return true, newCorporationTiles
		}
	}
	return false, newCorporationTiles
}

// Checks if the passed tile merges two or more corporations, returns a slice of
// corporation IDs to be merged
func (b *Board) TileMergeCorporations(t tileset.Position) (bool, []corporation.Interface) {
	var corporations []corporation.Interface
	adjacent := b.adjacentCorporationTiles(t)
	for _, adjacentCell := range adjacent {
		corporations = append(corporations, b.Cell(adjacentCell))
	}
	if len(corporations) > 1 {
		return true, corporations
	}
	return false, []corporation.Interface{}
}

// Check if the passed tile grows a corporation
// Returns true if that's the case, the tiles to append to the corporation and
// the ID of the corporation which grows
func (b *Board) TileGrowCorporation(t tileset.Position) (bool, []tileset.Position, corporation.Interface) {
	tilesToAppend := []tileset.Position{{Number: t.Number, Letter: t.Letter}}
	nullCorporation := corporation.Interface{}
	corporationToGrow := nullCorporation
	adjacent := b.adjacentTiles(t)
	for _, adjacentCell := range adjacent {
		if b.Cell(adjacentCell).ContentType() != "orphan" {
			if corporationToGrow != nullCorporation {
				return false, []tileset.Position{}, nullCorporation
			}
			corporationToGrow = b.Cell(adjacentCell)
		} else {
			tilesToAppend = append(tilesToAppend, adjacentCell)
		}
	}
	if corporationToGrow == nullCorporation {
		return false, []tileset.Position{}, nullCorporation
	}
	return true, tilesToAppend, corporationToGrow
}

// Puts the passed tile on the board
func (b *Board) PutTile(t board.Container) {
	b.grid[t.Number][t.Letter] = t
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
	if t.Number < 12 {
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
			if b.Cell(t).ContentType() != "empty" {
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
			if b.Cell(t).ContentType() == "corporation" {
				return true
			}
			return false
		},
	)
}

func (b *Board) adjacentNonCorporationTiles(t tileset.Position) []tileset.Position {
	return b.adjacentCellsWithFilter(
		t,
		func(t tileset.Position) bool {
			if b.Cell(t).ContentType() == "orphan" || b.Cell(t).ContentType() == "empty" {
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

func (b *Board) SetTiles(cp corporation.Interface, cells []tileset.Position) {
	for _, cell := range cells {
		b.grid[cell.Number][cell.Letter] = cp.Id()
	}
}
