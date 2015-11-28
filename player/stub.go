package player

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

type Stub struct {
	Player
}

func NewStub(name string) *Stub {
	return &Stub{
		Player{
			name:   name,
			cash:   6000,
			shares: map[string]int{},
		},
	}
}

func (p *Stub) SetTiles(tiles []tile.Interface) {
	p.tiles = tiles
}

func (p *Stub) SetShares(corp corporation.Interface, amount int) Interface {
	p.shares[corp.Name()] = amount
	return p
}

func (p *Stub) SetCash(amount int) Interface {
	p.cash = amount
	return p
}
