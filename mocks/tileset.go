package mocks

import (
	"github.com/svera/acquire/interfaces"
)

// Tileset is a structure that implements the Tileset interface for testing
type Tileset struct {
	FakeTile    interfaces.Tile
	FakeError   error
	TimesCalled map[string]int
}

// Draw mocks the Draw method defined in the Tileset interface
func (t *Tileset) Draw() (interfaces.Tile, error) {
	return t.FakeTile, t.FakeError
}

// Add mocks the Add method defined in the Tileset interface
func (t *Tileset) Add([]interfaces.Tile) interfaces.Tileset {
	t.TimesCalled["Add"]++
	return t
}
