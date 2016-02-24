package interfaces

// Board is an interface that declares all methods to be implemented by a board implementation
type Board interface {
	Cell(number int, letter string) Owner
	TileFoundCorporation(t Tile) (bool, []Tile)
	TileMergeCorporations(t Tile) (bool, map[string][]Corporation)
	TileGrowCorporation(t Tile) (bool, []Tile, Corporation)
	PutTile(t Tile) Board
	AdjacentCells(number int, letter string) []Owner
	SetOwner(cp Corporation, tiles []Tile) Board
	ChangeOwner(oldOwner Corporation, newOwner Corporation) Board
}

type BoardMock struct {
	FakeCellOwner              Owner
	FakeFoundCorporation       bool
	FakeFoundCorporationTiles  []Tile
	FakeMergeCorporations      bool
	FakeMergeCorporationsCorps map[string][]Corporation
	FakeGrowCorporation        bool
	FakeGrowCorporationTiles   []Tile
	FakeGrowCorporationCorp    Corporation
	FakeAdjacentCells          []Owner
}

func (b *BoardMock) Cell(number int, letter string) Owner {
	return b.FakeCellOwner
}

func (b *BoardMock) TileFoundCorporation(t Tile) (bool, []Tile) {
	return b.FakeFoundCorporation, b.FakeFoundCorporationTiles
}

func (b *BoardMock) TileMergeCorporations(t Tile) (bool, map[string][]Corporation) {
	return b.FakeMergeCorporations, b.FakeMergeCorporationsCorps
}

func (b *BoardMock) TileGrowCorporation(t Tile) (bool, []Tile, Corporation) {
	return b.FakeGrowCorporation, b.FakeGrowCorporationTiles, b.FakeGrowCorporationCorp
}

func (b *BoardMock) PutTile(t Tile) Board {
	_ = t
	return b
}

func (b *BoardMock) AdjacentCells(number int, letter string) []Owner {
	_, _ = number, letter
	return b.FakeAdjacentCells
}

func (b *BoardMock) SetOwner(cp Corporation, tiles []Tile) Board {
	_, _ = cp, tiles
	return b
}

func (b *BoardMock) ChangeOwner(oldOwner Corporation, newOwner Corporation) Board {
	_, _ = oldOwner, newOwner
	return b
}
