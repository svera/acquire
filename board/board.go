// Model for the game board, storing its state and implementing related
// actions.
package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

var letters [9]string = [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

type Board struct {
	grid *[13]map[string]Container
}

type Coordinates struct {
	Number int
	Letter string
}

func New() *Board {
	brd := Board{
		grid: new([13]map[string]Container),
	}

	emptyCell := &Empty{}
	for number := 1; number < 13; number++ {
		brd.grid[number] = make(map[string]Container)
		for _, letter := range letters {
			brd.grid[number][letter] = emptyCell
		}
	}

	return &brd
}

// Returns a board cell content
func (b *Board) Cell(t Coordinates) Container {
	return b.grid[t.Number][t.Letter]
}

// Checks if the passed tile founds a new corporation, returns a slice of tiles
// composing this corporation
func (b *Board) TileFoundCorporation(t Coordinates) (bool, []Coordinates) {
	var newCorporationTiles []Coordinates
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
func (b *Board) TileMergeCorporations(t Coordinates) (bool, []corporation.Interface) {
	var corporations []corporation.Interface
	adjacent := b.adjacentCorporationTiles(t)
	for _, adjacentCell := range adjacent {
		corporations = append(corporations, b.Cell(adjacentCell).(corporation.Interface))
	}
	if len(corporations) > 1 {
		return true, corporations
	}
	return false, []corporation.Interface{}
}

// Check if the passed tile grows a corporation
// Returns true if that's the case, the tiles to append to the corporation and
// the ID of the corporation which grows
func (b *Board) TileGrowCorporation(t Coordinates) (bool, []Coordinates, corporation.Interface) {
	tilesToAppend := []Coordinates{{Number: t.Number, Letter: t.Letter}}
	var nullCorporation corporation.Interface
	corporationToGrow := nullCorporation
	adjacent := b.adjacentTiles(t)
	for _, adjacentCell := range adjacent {
		if b.Cell(adjacentCell).ContentType() != "orphan" {
			if corporationToGrow != nullCorporation {
				return false, []Coordinates{}, nullCorporation
			}
			corporationToGrow = b.Cell(adjacentCell).(corporation.Interface)
		} else {
			tilesToAppend = append(tilesToAppend, adjacentCell)
		}
	}
	if corporationToGrow == nullCorporation {
		return false, []Coordinates{}, nullCorporation
	}
	return true, tilesToAppend, corporationToGrow
}

// Puts the passed tile on the board
func (b *Board) PutTile(t *tile.Orphan) {
	b.grid[t.Number][t.Letter] = t
}

// Returns all cells adjacent to the passed one
func (b *Board) AdjacentCells(t Coordinates) []Coordinates {
	var adjacent []Coordinates

	if t.Letter > "A" {
		adjacent = append(adjacent, Coordinates{Number: t.Number, Letter: previousLetter(t.Letter)})
	}
	if t.Letter < "I" {
		adjacent = append(adjacent, Coordinates{Number: t.Number, Letter: nextLetter(t.Letter)})
	}
	if t.Number > 1 {
		adjacent = append(adjacent, Coordinates{Number: t.Number - 1, Letter: t.Letter})
	}
	if t.Number < 12 {
		adjacent = append(adjacent, Coordinates{Number: t.Number + 1, Letter: t.Letter})
	}
	return adjacent
}

func (b *Board) adjacentCellsWithFilter(t Coordinates, filter func(Coordinates) bool) []Coordinates {
	var adjacentFilteredCells []Coordinates
	adjacent := b.AdjacentCells(t)

	for _, adjacentCell := range adjacent {
		if filter(adjacentCell) {
			adjacentFilteredCells = append(adjacentFilteredCells, adjacentCell)
		}
	}
	return adjacentFilteredCells
}

func (b *Board) adjacentTiles(t Coordinates) []Coordinates {
	return b.adjacentCellsWithFilter(
		t,
		func(t Coordinates) bool {
			if b.Cell(t).ContentType() != "empty" {
				return true
			}
			return false
		},
	)
}

func (b *Board) adjacentCorporationTiles(t Coordinates) []Coordinates {
	return b.adjacentCellsWithFilter(
		t,
		func(t Coordinates) bool {
			if b.Cell(t).ContentType() == "corporation" {
				return true
			}
			return false
		},
	)
}

func (b *Board) adjacentNonCorporationTiles(t Coordinates) []Coordinates {
	return b.adjacentCellsWithFilter(
		t,
		func(t Coordinates) bool {
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

func (b *Board) SetTiles(cp corporation.Interface, cells []Coordinates) {
	for _, cell := range cells {
		b.grid[cell.Number][cell.Letter] = cp
	}
}
