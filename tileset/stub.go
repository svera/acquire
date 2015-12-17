package tileset

import (
	"github.com/svera/acquire/tile"
)

// Stub is a struct to be used in tileset tests as a replacement of the original,
// that includes several convenience methods for testing
type Stub struct {
	Tileset
}

// NewStub initialises and returns a new instance of Stub
func NewStub() *Stub {
	stub := Stub{}
	stub.Tileset = *New()
	return &stub
}

// DiscardTile removes passed tile from the tileset
func (t *Stub) DiscardTile(tl tile.Interface) {
	for i, currentTile := range t.tiles {
		if currentTile.Number() == tl.Number() && currentTile.Letter() == tl.Letter() {
			t.tiles = append(t.tiles[:i], t.tiles[i+1:]...)
			break
		}
	}
}
