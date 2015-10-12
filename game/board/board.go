package board

import (
	"github.com/svera/acquire/game/tileset"
)

const boardCellEmpty = 0
const boardCellUsed = 9

var letters [9]string

type Board struct {
	grid *[13]map[string]byte
}

func init() {
	letters = [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
}

func New() *Board {
	board := Board{
		grid: new([13]map[string]byte),
	}

	for number := 1; number < 13; number++ {
		board.grid[number] = make(map[string]byte)
		for _, letter := range letters {
			board.grid[number][letter] = boardCellEmpty
		}
	}

	return &board
}

// Checks of the passed tile founds a new corporation, returns a slice of tiles composing this corporation
func (b *Board) CorporationFounded(t tileset.Tile) []tileset.Tile {
	var newCorporationTiles []tileset.Tile
	adjacent := b.adjacentTiles(t)
	for _, tile := range adjacent {
		if b.grid[tile.Number][tile.Letter] == boardCellEmpty {
			newCorporationTiles = append(newCorporationTiles, tile)
		}
	}
	return newCorporationTiles
}

// Placeholder function, pending implementation
func (b *Board) areCorporationsMerged() bool {
	return true
}

// Placeholder function, pending implementation
func (b *Board) isTilePlayable() bool {
	return true
}

func (b *Board) PutTile(t tileset.Tile) {
	b.grid[t.Number][t.Letter] = boardCellUsed
}

func (b *Board) adjacentTiles(t tileset.Tile) []tileset.Tile {
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

func getLetterByDelta(letter string, delta int) string {
	for i, currentLetter := range letters {
		if currentLetter == letter {
			return letters[i+delta]
		}
	}
	return ""
}

func previousLetter(letter string) string {
	return getLetterByDelta(letter, -1)
}

func nextLetter(letter string) string {
	return getLetterByDelta(letter, +1)
}