package player

import (
	"github.com/svera/acquire/interfaces"
)

// Stub is a struct to be used in player tests as a replacement of the original,
// that includes several convenience methods for testing
type Stub struct {
	Player
}

// NewStub initialises and returns a new instance of Stub
func NewStub(name string) *Stub {
	return &Stub{
		Player{
			name:   name,
			cash:   6000,
			shares: map[interfaces.Corporation]int{},
		},
	}
}

// SetTiles sets player tiles
func (p *Stub) SetTiles(tiles []interfaces.Tile) {
	p.tiles = tiles
}

// SetShares sets the shares the player has of a certain corporation
func (p *Stub) SetShares(corp interfaces.Corporation, amount int) interfaces.Player {
	p.shares[corp] = amount
	return p
}

// SetCash sets the amount of cash of the player
func (p *Stub) SetCash(amount int) interfaces.Player {
	p.cash = amount
	return p
}
