// Player model, which manages player status in game.
package player

import (
	"errors"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

const (
	TileNotOnHand = "tile_not_on_hand"
	TooManyTiles  = "too_many_tiles"
)

type Player struct {
	name   string
	cash   int
	tiles  []tile.Interface
	shares map[string]int
}

func New(name string) *Player {
	return &Player{
		name:   name,
		cash:   6000,
		shares: map[string]int{},
	}
}

// Returns the number of shares for the passed corporation owned by the player
func (p *Player) Shares(corp corporation.Interface) int {
	return p.shares[corp.Name()]
}

// Buys stock from corporation
func (p *Player) Buy(corp corporation.Interface, amount int) {
	corp.SetStock(corp.Stock() - amount)
	p.shares[corp.Name()] = amount
	p.cash -= corp.StockPrice() * amount
}

// Receive a free stock share from a rencently found corporation, if it has
// remaining shares available
func (p *Player) GetFounderStockShare(corp corporation.Interface) {
	if corp.Stock() > 0 {
		corp.SetStock(corp.Stock() - 1)
		p.shares[corp.Name()] += 1
	}
}

// Adds a new tile to the players' tileset
func (p *Player) PickTile(tile tile.Interface) error {
	if len(p.tiles) >= 6 {
		return errors.New(TooManyTiles)
	}
	p.tiles = append(p.tiles, tile)
	return nil
}

// Return player's tiles
func (p *Player) Tiles() []tile.Interface {
	return p.tiles
}

func (p *Player) ReceiveBonus(amount int) {
	p.cash += amount
}

// Discard passed tile from player's hand
func (p *Player) DiscardTile(tile tile.Interface) error {
	for i, currentTile := range p.tiles {
		if currentTile.Number() == tile.Number() && currentTile.Letter() == tile.Letter() {
			p.tiles = append(p.tiles[:i], p.tiles[i+1:]...)
			return nil
		}
	}
	return errors.New(TileNotOnHand)
}

// Returns player cash
func (p *Player) Cash() int {
	return p.cash
}
