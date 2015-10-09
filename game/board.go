package game

const boardCellEmpty = 0
const boardCellUsed = 9

type Board struct {
	grid *[13]map[string]byte
}

func NewBoard() *Board {
	board := Board{
		grid: new([13]map[string]byte),
	}
	letters := [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	for number := 1; number < 13; number++ {
		board.grid[number] = make(map[string]byte)
		for _, letter := range letters {
			board.grid[number][letter] = boardCellEmpty
		}
	}

	return &board
}

// Placeholder function, pending implementation
func (b *Board) isCorporationFounded() bool {
	return true
}

// Placeholder function, pending implementation
func (b *Board) areCorporationsMerged() bool {
	return true
}

// Placeholder function, pending implementation
func (b *Board) isTilePlayable() bool {
	return true
}

func (b *Board) PutTile(tile Tile) {
	b.grid[tile.number][tile.letter] = boardCellUsed
}
