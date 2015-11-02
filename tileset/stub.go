package tileset

type Stub struct {
	Tileset
}

func NewStub() *Stub {
	stub := Stub{}
	stub.Tileset = *New()
	return &stub
}

func (t *Stub) DiscardTile(tile Position) {
	for i, currentTile := range t.tiles {
		if currentTile.Number == tile.Number && currentTile.Letter == tile.Letter {
			t.tiles = append(t.tiles[:i], t.tiles[i+1:]...)
			break
		}
	}
}
