package mocks

import (
	"github.com/svera/acquire/interfaces"
)

// Player is a structure that implements the Player interface for testing
type Player struct {
	FakeShares  map[interfaces.Corporation]int
	FakeTiles   []interfaces.Tile
	FakeHasTile bool
	FakeCash    int
	TimesCalled map[string]int
	FakeActive  bool
}

// Shares mocks the Shares method defined in the Player interface
func (p *Player) Shares(c interfaces.Corporation) int {
	return p.FakeShares[c]
}

// AddShares mocks the AddShares method defined in the Player interface
func (p *Player) AddShares(c interfaces.Corporation, amount int) interfaces.Player {
	p.FakeShares[c] += amount
	return p
}

// RemoveShares mocks the RemoveShares method defined in the Player interface
func (p *Player) RemoveShares(c interfaces.Corporation, amount int) interfaces.Player {
	p.FakeShares[c] -= amount
	return p
}

// PickTile mocks the PickTile method defined in the Player interface
func (p *Player) PickTile(t interfaces.Tile) interfaces.Player {
	p.FakeTiles = append(p.FakeTiles, t)
	p.TimesCalled["PickTile"]++
	return p
}

// Tiles mocks the Tiles method defined in the Player interface
func (p *Player) Tiles() []interfaces.Tile {
	return p.FakeTiles
}

// DiscardTile mocks the DiscarTile method defined in the Player interface
func (p *Player) DiscardTile(t interfaces.Tile) interfaces.Player {
	p.TimesCalled["DiscardTile"]++
	return p
}

// HasTile mocks the HasTile method defined in the Player interface
func (p *Player) HasTile(t interfaces.Tile) bool {
	return p.FakeHasTile
}

// Cash mocks the Cash method defined in the Player interface
func (p *Player) Cash() int {
	return p.FakeCash
}

// AddCash mocks the AddCash method defined in the Player interface
func (p *Player) AddCash(amount int) interfaces.Player {
	p.FakeCash += amount
	return p
}

// RemoveCash mocks the RemoveCash method defined in the Player interface
func (p *Player) RemoveCash(amount int) interfaces.Player {
	p.FakeCash -= amount
	return p
}

// Active mocks the Active method defined in the Player interface
func (p *Player) Active() bool {
	return p.FakeActive
}

// Deactivate mocks the Deactivate method defined in the Player interface
func (p *Player) Deactivate() interfaces.Player {
	p.FakeActive = false
	p.TimesCalled["Deactivate"]++
	return p
}
