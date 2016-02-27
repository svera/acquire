package mocks

import (
	"github.com/svera/acquire/interfaces"
)

type Tileset struct {
	FakeTile  interfaces.Tile
	FakeError error
}

func (t *Tileset) Draw() (interfaces.Tile, error) {
	return t.FakeTile, t.FakeError
}
