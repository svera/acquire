package tileset

import (
	"github.com/svera/acquire/tile"
)

type Stub struct {
	Tileset
}

func NewStub() *Stub {
	stub := Stub{}
	stub.Tileset = *New()
	return &stub
}

func (t *Stub) DiscardTile(tl tile.Orphan) {
	for i, currentTile := range t.tiles {
		if currentTile.Number == tl.Number && currentTile.Letter == tl.Letter {
			t.tiles = append(t.tiles[:i], t.tiles[i+1:]...)
			break
		}
	}
}
