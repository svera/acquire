package player

import (
	"errors"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tileset"
)

type Player struct {
	name   string
	cash   int
	tiles  []tileset.Position
	shares [7]int
}

func New(name string) *Player {
	return &Player{
		name:   name,
		cash:   6000,
		shares: [7]int{},
	}
}

func (p *Player) Shares(c *corporation.Corporation) int {
	return p.shares[c.Id()]
}

// Buys stock from corporation
func (p *Player) Buy(corp *corporation.Corporation, amount int) {
	corp.SetStock(corp.Stock() - amount)
	p.shares[corp.Id()] = amount
	p.cash -= corp.StockPrice() * amount
}

// Adds a new tile to the players' tileset
func (p *Player) PickTile(t tileset.Position) error {
	if len(p.tiles) >= 6 {
		return errors.New("Player cannot have more than 6 tiles")
	}
	p.tiles = append(p.tiles, t)
	return nil
}

func (p *Player) Tiles() []tileset.Position {
	return p.tiles
}

func (p *Player) ReceiveBonus(amount int) {
	p.cash += amount
}

func (p *Player) UseTile(t tileset.Position) error {
	for i, currentTile := range p.tiles {
		if currentTile.Number == t.Number && currentTile.Letter == t.Letter {
			p.tiles = append(p.tiles[:i], p.tiles[i+1:]...)
			return nil
		}
	}
	return errors.New("Player doesn't have tile on hand")
}

func (p *Player) Cash() int {
	return p.cash
}
