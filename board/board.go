// Package board holds the struct and related methods that model the game board
package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

var letters = [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

// Board maps tiles on every position on board
type Board struct {
	grid *[13]map[string]tile.Interface
}

// New initialises and returns a Board instance
func New() *Board {
	brd := Board{
		grid: new([13]map[string]tile.Interface),
	}

	for number := 1; number < 13; number++ {
		brd.grid[number] = make(map[string]tile.Interface)
		for _, letter := range letters {
			brd.grid[number][letter] = tile.New(number, letter, tile.Empty{})
		}
	}

	return &brd
}

// Cell returns a board cell content
func (b *Board) Cell(number int, letter string) tile.Interface {
	return b.grid[number][letter]
}

// TileFoundCorporation checks if the passed tile founds a new corporation, returns a slice of tiles
// composing this corporation
func (b *Board) TileFoundCorporation(t tile.Interface) (bool, []tile.Interface) {
	var newCorporationTiles []tile.Interface
	adjacent := b.adjacentTiles(t)
	for _, adjacentTile := range adjacent {
		if adjacentTile.Owner().Type() == "corporation" {
			return false, []tile.Interface{}
		}
		newCorporationTiles = append(newCorporationTiles, adjacentTile)
	}
	if len(newCorporationTiles) > 0 {
		newCorporationTiles = append(newCorporationTiles, t)
		return true, newCorporationTiles
	}
	return false, newCorporationTiles
}

// TileMergeCorporations checks if the passed tile merges two or more corporations, returns a slice of
// corporation IDs to be merged
func (b *Board) TileMergeCorporations(t tile.Interface) (bool, map[string][]corporation.Interface) {
	var corporations []corporation.Interface

	adjacent := b.adjacentCorporationTiles(t)
	for _, adjacentCell := range adjacent {
		corp, _ := adjacentCell.Owner().(corporation.Interface)
		corporations = append(corporations, corp)
	}
	if len(corporations) > 1 {
		return true, categorizeMerge(corporations)
	}
	return false, map[string][]corporation.Interface{}
}

// Distributes the corporation in a merge between acquirers and defuncts
func categorizeMerge(corporations []corporation.Interface) map[string][]corporation.Interface {
	sizeDesc := func(corp1, corp2 corporation.Interface) bool {
		return corp1.Size() > corp2.Size()
	}
	corporation.By(sizeDesc).Sort(corporations)

	merge := map[string][]corporation.Interface{
		"acquirer": []corporation.Interface{corporations[0]},
		"defunct":  []corporation.Interface{},
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
func (b *Board) TileGrowCorporation(t tile.Interface) (bool, []tile.Interface, corporation.Interface) {
	tilesToAppend := []tile.Interface{t}
	var nullCorporation corporation.Interface
	corporationToGrow := nullCorporation
	adjacent := b.adjacentTiles(t)
	for _, adjacentCell := range adjacent {
		if adjacentCell.Owner().Type() != "unincorporated" {
			if corporationToGrow != nullCorporation {
				return false, []tile.Interface{}, nullCorporation
			}
			corporationToGrow = adjacentCell.Owner().(corporation.Interface)
		} else {
			tilesToAppend = append(tilesToAppend, adjacentCell)
		}
	}
	if corporationToGrow == nullCorporation {
		return false, []tile.Interface{}, nullCorporation
	}
	return true, tilesToAppend, corporationToGrow
}

// PutTile puts the passed tile on the board
func (b *Board) PutTile(t tile.Interface) Interface {
	b.grid[t.Number()][t.Letter()] = t
	return b
}

// AdjacentCells returns all cells adjacent to the passed one
func (b *Board) AdjacentCells(t tile.Interface) []tile.Interface {
	var adjacent []tile.Interface

	if t.Letter() > "A" {
		adjacent = append(adjacent, b.grid[t.Number()][previousLetter(t.Letter())])
	}
	if t.Letter() < "I" {
		adjacent = append(adjacent, b.grid[t.Number()][nextLetter(t.Letter())])
	}
	if t.Number() > 1 {
		adjacent = append(adjacent, b.grid[t.Number()-1][t.Letter()])
	}
	if t.Number() < 12 {
		adjacent = append(adjacent, b.grid[t.Number()+1][t.Letter()])
	}
	return adjacent
}

func (b *Board) adjacentCellsWithFilter(t tile.Interface, filter func(tile.Interface) bool) []tile.Interface {
	var adjacentFilteredCells []tile.Interface
	adjacent := b.AdjacentCells(t)

	for _, adjacentCell := range adjacent {
		if filter(adjacentCell) {
			adjacentFilteredCells = append(adjacentFilteredCells, adjacentCell)
		}
	}
	return adjacentFilteredCells
}

func (b *Board) adjacentTiles(t tile.Interface) []tile.Interface {
	return b.adjacentCellsWithFilter(
		t,
		func(t tile.Interface) bool {
			if t.Owner().Type() != "empty" {
				return true
			}
			return false
		},
	)
}

func (b *Board) adjacentCorporationTiles(t tile.Interface) []tile.Interface {
	return b.adjacentCellsWithFilter(
		t,
		func(t tile.Interface) bool {
			if t.Owner().Type() == "corporation" {
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
func (b *Board) SetOwner(cp corporation.Interface, tiles []tile.Interface) Interface {
	for _, tl := range tiles {
		b.grid[tl.Number()][tl.Letter()].SetOwner(cp)
	}
	return b
}

// ChangeOwner changes ownership of tiles belonging to oldOrder to newOrder
func (b *Board) ChangeOwner(oldOwner corporation.Interface, newOwner corporation.Interface) Interface {
	for number := 1; number < 13; number++ {
		for _, letter := range letters {
			if b.grid[number][letter].Owner() == oldOwner {
				b.grid[number][letter].SetOwner(newOwner)
			}
		}
	}
	return b
}
