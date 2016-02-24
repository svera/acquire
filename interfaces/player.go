package interfaces

// Player is an interface that declares all methods to be implemented by a player implementation
type Player interface {
	Shares(corp Corporation) int
	AddShares(corp Corporation, amount int) Player
	RemoveShares(corp Corporation, amount int) Player
	PickTile(t Tile) Player
	Tiles() []Tile
	DiscardTile(t Tile) Player
	HasTile(t Tile) bool
	Cash() int
	AddCash(amount int) Player
	RemoveCash(amount int) Player
}

type PlayerMock struct {
	FakeShares  map[Corporation]int
	FakeTiles   []Tile
	FakeHasTile bool
	FakeCash    int
}

func (p *PlayerMock) Shares(c Corporation) int {
	return p.FakeShares[c]
}

func (p *PlayerMock) AddShares(c Corporation, amount int) Player {
	_, _ = c, amount
	return p
}

func (p *PlayerMock) RemoveShares(c Corporation, amount int) Player {
	_, _ = c, amount
	return p
}

func (p *PlayerMock) PickTile(t Tile) Player {
	_ = t
	return p
}

func (p *PlayerMock) Tiles() []Tile {
	return p.FakeTiles
}

func (p *PlayerMock) DiscardTile(t Tile) Player {
	_ = t
	return p
}

func (p *PlayerMock) HasTile(t Tile) bool {
	return p.FakeHasTile
}

func (p *PlayerMock) Cash() int {
	return p.FakeCash
}

func (p *PlayerMock) AddCash(amount int) Player {
	_ = amount
	return p
}

func (p *PlayerMock) RemoveCash(amount int) Player {
	_ = amount
	return p
}
