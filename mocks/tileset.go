package mocks

import (
	"github.com/svera/acquire/interfaces"
)

// Tileset is a structure that implements the Tileset interface for testing
type Tileset struct {
	FakeTile  interfaces.Tile
	FakeError error
}

// Draw mocks the Draw method defined in the Tileset interface
func (t *Tileset) Draw() (interfaces.Tile, error) {
	return t.FakeTile, t.FakeError
}
