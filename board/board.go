// Package board holds the struct and related methods that model the game board
package board

import (
	"github.com/svera/acquire/interfaces"
	"sort"
)

var letters = [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

// Board maps tiles on every position on board
type Board struct {
	grid *[13]map[string]interfaces.Owner
}

type sortableCorporations []interfaces.Corporation

func (s sortableCorporations) Len() int           { return len(s) }
func (s sortableCorporations) Less(i, j int) bool { return s[i].Size() < s[j].Size() }
func (s sortableCorporations) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// New initialises and returns a Board instance
func New() *Board {
	brd := Board{
		grid: new([13]map[string]interfaces.Owner),
	}

	for number := 1; number < 13; number++ {
		brd.grid[number] = make(map[string]interfaces.Owner)
		for _, letter := range letters {
			brd.grid[number][letter] = Empty{}
		}
	}

	return &brd
}

// Cell returns a board cell content
func (b *Board) Cell(number int, letter string) interfaces.Owner {
	return b.grid[number][letter]
}

// TileFoundCorporation checks if the passed tile founds a new corporation, returns a slice of tiles
// composing this corporation
func (b *Board) TileFoundCorporation(t interfaces.Tile) (bool, []interfaces.Tile) {
	var newCorporationTiles []interfaces.Tile
	adjacent := b.adjacentTiles(t.Number(), t.Letter())
	for _, adjacentTile := range adjacent {
		if adjacentTile.Type() == "corporation" {
			return false, []interfaces.Tile{}
		}
		newCorporationTiles = append(newCorporationTiles, adjacentTile.(interfaces.Tile))
	}
	if len(newCorporationTiles) > 0 {
		newCorporationTiles = append(newCorporationTiles, t)
		return true, newCorporationTiles
	}
	return false, newCorporationTiles
}

// TileMergeCorporations checks if the passed tile merges two or more corporations, returns a map of
// corporations categorized between "acquirer" and "defunct"
func (b *Board) TileMergeCorporations(t interfaces.Tile) (bool, map[string][]interfaces.Corporation) {
	var corporations sortableCorporations
	var exist bool

	adjacent := b.adjacentCorporationTiles(t.Number(), t.Letter())
	for _, adjacentCell := range adjacent {
		exist = false
		corp, _ := adjacentCell.(interfaces.Corporation)
		for i := range corporations {
			if corp == corporations[i] {
				exist = true
			}
		}
		if !exist {
			corporations = append(corporations, corp)
		}
	}
	if len(corporations) > 1 {
		return true, categorizeMerge(corporations)
	}
	return false, map[string][]interfaces.Corporation{}
}

// Distributes the corporation in a merge between acquirers and defuncts
func categorizeMerge(corporations sortableCorporations) map[string][]interfaces.Corporation {
	sort.Sort(sort.Reverse(corporations))

	merge := map[string][]interfaces.Corporation{
		"acquirer": []interfaces.Corporation{corporations[0]},
		"defunct":  []interfaces.Corporation{},
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
func (b *Board) TileGrowCorporation(tl interfaces.Tile) (bool, []interfaces.Tile, interfaces.Corporation) {
	tilesToAppend := []interfaces.Tile{tl}
	var nullCorporation interfaces.Corporation
	corporationToGrow := nullCorporation
	adjacent := b.adjacentTiles(tl.Number(), tl.Letter())
	for _, adjacentCell := range adjacent {
		if adjacentCell.Type() != "unincorporated" {
			if corporationToGrow != nullCorporation {
				return false, []interfaces.Tile{}, nullCorporation
			}
			corporationToGrow = adjacentCell.(interfaces.Corporation)
		} else {
			tilesToAppend = append(tilesToAppend, adjacentCell.(interfaces.Tile))
		}
	}
	if corporationToGrow == nullCorporation {
		return false, []interfaces.Tile{}, nullCorporation
	}
	return true, tilesToAppend, corporationToGrow
}

// PutTile puts the passed tile on the board
func (b *Board) PutTile(t interfaces.Tile) interfaces.Board {
	b.grid[t.Number()][t.Letter()] = t
	return b
}

// AdjacentCells returns all cells adjacent to the passed one
func (b *Board) AdjacentCells(number int, letter string) []interfaces.Owner {
	var adjacent []interfaces.Owner

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

func (b *Board) adjacentCellsWithFilter(number int, letter string, filter func(interfaces.Owner) bool) []interfaces.Owner {
	var adjacentFilteredCells []interfaces.Owner
	adjacent := b.AdjacentCells(number, letter)

	for _, adjacentCell := range adjacent {
		if filter(adjacentCell) {
			adjacentFilteredCells = append(adjacentFilteredCells, adjacentCell)
		}
	}
	return adjacentFilteredCells
}

func (b *Board) adjacentTiles(number int, letter string) []interfaces.Owner {
	return b.adjacentCellsWithFilter(
		number,
		letter,
		func(o interfaces.Owner) bool {
			if o.Type() != "empty" {
				return true
			}
			return false
		},
	)
}

func (b *Board) adjacentCorporationTiles(number int, letter string) []interfaces.Owner {
	return b.adjacentCellsWithFilter(
		number,
		letter,
		func(o interfaces.Owner) bool {
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
func (b *Board) SetOwner(cp interfaces.Corporation, tiles []interfaces.Tile) interfaces.Board {
	for _, tl := range tiles {
		b.grid[tl.Number()][tl.Letter()] = cp
	}
	return b
}

// ChangeOwner changes ownership of tiles belonging to oldOrder to newOrder
func (b *Board) ChangeOwner(oldOwner interfaces.Corporation, newOwner interfaces.Corporation) interfaces.Board {
	for number := 1; number < 13; number++ {
		for _, letter := range letters {
			if b.grid[number][letter] == oldOwner {
				b.grid[number][letter] = newOwner
			}
		}
	}
	return b
}
