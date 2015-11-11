package tile

import (
	"github.com/svera/acquire/board"
)

type Orphan struct {
	board.Coordinates
}

func New(number int, letter string) *Orphan {
	return &Orphan{board.Coordinates{number, letter}}
}

func (t *Orphan) ContentType() string {
	return "orphan"
}
