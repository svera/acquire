package player

import (
	"errors"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tileset"
)

const (
	TileNotOnHand = "tile_not_on_hand"
	TooManyTiles  = "too_many_tiles"
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

func (p *Player) Shares(corp corporation.Interface) int {
	return p.shares[corp.Id()]
}

// Buys stock from corporation
func (p *Player) Buy(corp corporation.Interface, amount int) {
	corp.SetStock(corp.Stock() - amount)
	p.shares[corp.Id()] = amount
	p.cash -= corp.StockPrice() * amount
}

// Adds a new tile to the players' tileset
func (p *Player) PickTile(tile tileset.Position) error {
	if len(p.tiles) >= 6 {
		return errors.New(TooManyTiles)
	}
	p.tiles = append(p.tiles, tile)
	return nil
}

func (p *Player) Tiles() []tileset.Position {
	return p.tiles
}

func (p *Player) ReceiveBonus(amount int) {
	p.cash += amount
}

func (p *Player) DiscardTile(tile tileset.Position) error {
	for i, currentTile := range p.tiles {
		if currentTile.Number == tile.Number && currentTile.Letter == tile.Letter {
			p.tiles = append(p.tiles[:i], p.tiles[i+1:]...)
			return nil
		}
	}
	return errors.New(TileNotOnHand)
}

func (p *Player) Cash() int {
	return p.cash
}
