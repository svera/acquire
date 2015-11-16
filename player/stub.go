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
			shares: [7]int{},
		},
	}
}

func (p *Stub) SetShares(c corporation.Interface, amount int) {
	p.shares[c.Id()] = amount
}

func (p *Stub) SetCash(amount int) {
	p.cash = amount
}

func (p *Stub) SetTiles(tiles []*tile.Orphan) {
	p.tiles = tiles
}
