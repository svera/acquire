package mocks

import (
	"github.com/svera/acquire/interfaces"
)

type Player struct {
	FakeShares  map[interfaces.Corporation]int
	FakeTiles   []interfaces.Tile
	FakeHasTile bool
	FakeCash    int
	TimesCalled map[string]int
}

func (p *Player) Shares(c interfaces.Corporation) int {
	return p.FakeShares[c]
}

func (p *Player) AddShares(c interfaces.Corporation, amount int) interfaces.Player {
	p.FakeShares[c] += amount
	return p
}

func (p *Player) RemoveShares(c interfaces.Corporation, amount int) interfaces.Player {
	p.FakeShares[c] -= amount
	return p
}

func (p *Player) PickTile(t interfaces.Tile) interfaces.Player {
	p.FakeTiles = append(p.FakeTiles, t)
	return p
}

func (p *Player) Tiles() []interfaces.Tile {
	return p.FakeTiles
}

func (p *Player) DiscardTile(t interfaces.Tile) interfaces.Player {
	p.TimesCalled["DiscardTile"]++
	return p
}

func (p *Player) HasTile(t interfaces.Tile) bool {
	return p.FakeHasTile
}

func (p *Player) Cash() int {
	return p.FakeCash
}

func (p *Player) AddCash(amount int) interfaces.Player {
	p.FakeCash += amount
	return p
}

func (p *Player) RemoveCash(amount int) interfaces.Player {
	p.FakeCash -= amount
	return p
}
