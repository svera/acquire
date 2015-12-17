package player

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
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
			shares: map[string]int{},
		},
	}
}

func (p *Stub) SetOwner(tiles []tile.Interface) {
	p.tiles = tiles
}

// SetShares sets the shares the player has of a certain corporation
func (p *Stub) SetShares(corp corporation.Interface, amount int) Interface {
	p.shares[corp.Name()] = amount
	return p
}

// SetCash sets the amount of cash of the player
func (p *Stub) SetCash(amount int) Interface {
	p.cash = amount
	return p
}
