// Package board holds the struct and related methods that model the game board
package board

import (
	"github.com/svera/acquire/acquire"
)

var letters = [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

// Board maps tiles on every position on board
type Board struct {
	grid *[13]map[string]acquire.Owner
}

// New initialises and returns a Board instance
func New() *Board {
	brd := Board{
		grid: new([13]map[string]acquire.Owner),
	}

	for number := 1; number < 13; number++ {
		brd.grid[number] = make(map[string]acquire.Owner)
		for _, letter := range letters {
			brd.grid[number][letter] = Empty{}
		}
	}

	return &brd
}

// Cell returns a board cell content
func (b *Board) Cell(number int, letter string) acquire.Owner {
	return b.grid[number][letter]
}

// TileFoundCorporation checks if the passed tile founds a new corporation, returns a slice of tiles
// composing this corporation
func (b *Board) TileFoundCorporation(t acquire.Tile) (bool, []acquire.Tile) {
	var newCorporationTiles []acquire.Tile
	adjacent := b.adjacentTiles(t.Number(), t.Letter())
	for _, adjacentTile := range adjacent {
		if adjacentTile.Type() == "corporation" {
			return false, []acquire.Tile{}
		}
		newCorporationTiles = append(newCorporationTiles, adjacentTile.(acquire.Tile))
	}
	if len(newCorporationTiles) > 0 {
		newCorporationTiles = append(newCorporationTiles, t)
		return true, newCorporationTiles
	}
	return false, newCorporationTiles
}

// TileMergeCorporations checks if the passed tile merges two or more corporations, returns a map of
// corporations categorized between "acquirer" and "defunct"
func (b *Board) TileMergeCorporations(t acquire.Tile) (bool, map[string][]acquire.Corporation) {
	var corporations []acquire.Corporation

	adjacent := b.adjacentCorporationTiles(t.Number(), t.Letter())
	for _, adjacentCell := range adjacent {
		corp, _ := adjacentCell.(acquire.Corporation)
		corporations = append(corporations, corp)
	}
	if len(corporations) > 1 {
		return true, categorizeMerge(corporations)
	}
	return false, map[string][]acquire.Corporation{}
}

// Distributes the corporation in a merge between acquirers and defuncts
func categorizeMerge(corporations []acquire.Corporation) map[string][]acquire.Corporation {
	sizeDesc := func(corp1, corp2 acquire.Corporation) bool {
		return corp1.Size() > corp2.Size()
	}
	CorporationBy(sizeDesc).Sort(corporations)

	merge := map[string][]acquire.Corporation{
		"acquirer": []acquire.Corporation{corporations[0]},
		"defunct":  []acquire.Corporation{},
	}

	for i := 1; i < len(corporations); i++ {
		if corporations[i].Size() == corporations[0].Size() {
			merge["acquirer"] = append(merge["acquirer"], corporations[i])
		} else {
			merge["defunct"] = append(merge["defunct"], corporations[i])
		}
	}

	return merge
}

// TileGrowCorporation checks if the passed tile grows a corporation.
// Returns true if that's the case, the tiles to append to the corporation and
// the ID of the corporation which grows
func (b *Board) TileGrowCorporation(tl acquire.Tile) (bool, []acquire.Tile, acquire.Corporation) {
	tilesToAppend := []acquire.Tile{tl}
	var nullCorporation acquire.Corporation
	corporationToGrow := nullCorporation
	adjacent := b.adjacentTiles(tl.Number(), tl.Letter())
	for _, adjacentCell := range adjacent {
		if adjacentCell.Type() != "unincorporated" {
			if corporationToGrow != nullCorporation {
				return false, []acquire.Tile{}, nullCorporation
			}
			corporationToGrow = adjacentCell.(acquire.Corporation)
		} else {
			tilesToAppend = append(tilesToAppend, adjacentCell.(acquire.Tile))
		}
	}
	if corporationToGrow == nullCorporation {
		return false, []acquire.Tile{}, nullCorporation
	}
	return true, tilesToAppend, corporationToGrow
}

// PutTile puts the passed tile on the board
func (b *Board) PutTile(t acquire.Tile) acquire.Board {
	b.grid[t.Number()][t.Letter()] = t
	return b
}

// AdjacentCells returns all cells adjacent to the passed one
func (b *Board) AdjacentCells(number int, letter string) []acquire.Owner {
	var adjacent []acquire.Owner

	if letter > "A" {
		adjacent = append(adjacent, b.grid[number][previousLetter(letter)])
	}
	if letter < "I" {
		adjacent = append(adjacent, b.grid[number][nextLetter(letter)])
	}
	if number > 1 {
		adjacent = append(adjacent, b.grid[number-1][letter])
	}
	if number < 12 {
		adjacent = append(adjacent, b.grid[number+1][letter])
	}
	return adjacent
}

func (b *Board) adjacentCellsWithFilter(number int, letter string, filter func(acquire.Owner) bool) []acquire.Owner {
	var adjacentFilteredCells []acquire.Owner
	adjacent := b.AdjacentCells(number, letter)

	for _, adjacentCell := range adjacent {
		if filter(adjacentCell) {
			adjacentFilteredCells = append(adjacentFilteredCells, adjacentCell)
		}
	}
	return adjacentFilteredCells
}

func (b *Board) adjacentTiles(number int, letter string) []acquire.Owner {
	return b.adjacentCellsWithFilter(
		number,
		letter,
		func(o acquire.Owner) bool {
			if o.Type() != "empty" {
				return true
			}
			return false
		},
	)
}

func (b *Board) adjacentCorporationTiles(number int, letter string) []acquire.Owner {
	return b.adjacentCellsWithFilter(
		number,
		letter,
		func(o acquire.Owner) bool {
			if o.Type() == "corporation" {
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

// SetOwner sets tiles on board as belonging to the passed corporation
func (b *Board) SetOwner(cp acquire.Corporation, tiles []acquire.Tile) acquire.Board {
	for _, tl := range tiles {
		b.grid[tl.Number()][tl.Letter()] = cp
	}
	return b
}

// ChangeOwner changes ownership of tiles belonging to oldOrder to newOrder
func (b *Board) ChangeOwner(oldOwner acquire.Corporation, newOwner acquire.Corporation) acquire.Board {
	for number := 1; number < 13; number++ {
		for _, letter := range letters {
			if b.grid[number][letter] == oldOwner {
				b.grid[number][letter] = newOwner
			}
		}
	}
	return b
}
