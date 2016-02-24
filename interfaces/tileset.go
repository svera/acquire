package interfaces

// Tileset interface declares all methods to be implemented by a tileset implementation
type Tileset interface {
	Draw() (Tile, error)
}

type TilesetMock struct {
	FakeTile  Tile
	FakeError error
}

func (t *TilesetMock) Draw() (Tile, error) {
	return t.FakeTile, t.FakeError
}
