package player

import (
	"github.com/svera/acquire"
)

// Stub is a struct to be used in player tests as a replacement of the original,
// that includes several convenience methods for testing
type Stub struct {
	Player
}

// NewStub initialises and returns a new instance of Stub
func NewStub() *Stub {
	return &Stub{
		Player{
			cash:   6000,
			shares: map[acquire.Corporation]int{},
		},
	}
}

// SetTiles sets player tiles
func (p *Stub) SetTiles(tiles []acquire.Tile) {
	p.tiles = tiles
}

// SetShares sets the shares the player has of a certain corporation
func (p *Stub) SetShares(corp acquire.Corporation, amount int) acquire.Player {
	p.shares[corp] = amount
	return p
}

// SetCash sets the amount of cash of the player
func (p *Stub) SetCash(amount int) acquire.Player {
	p.cash = amount
	return p
}
