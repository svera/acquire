// Package player containst the struct Player and attached methods which manages player status in game.
package player

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

// Player stores the status of a player
type Player struct {
	name   string
	cash   int
	tiles  []tile.Interface
	shares map[string]int
}

// New initialises and returns a Player instance
func New(name string) *Player {
	return &Player{
		name:   name,
		cash:   6000,
		shares: map[string]int{},
	}
}

// Shares returns the number of shares for the passed corporation owned by the player
func (p *Player) Shares(corp corporation.Interface) int {
	return p.shares[corp.Name()]
}

// AddShares adds new stock shares of the passed corporation to the player
func (p *Player) AddShares(corp corporation.Interface, amount int) Interface {
	p.shares[corp.Name()] += amount
	return p
}

// RemoveShares removes stock shares of the passed corporation from the player
func (p *Player) RemoveShares(corp corporation.Interface, amount int) Interface {
	p.shares[corp.Name()] -= amount
	return p
}

// PickTile adds a new tile to the players' tileset
func (p *Player) PickTile(tile tile.Interface) Interface {
	p.tiles = append(p.tiles, tile)
	return p
}

// Tiles returns player's tiles
func (p *Player) Tiles() []tile.Interface {
	return p.tiles
}

// DiscardTile discards the passed tile from player's hand
func (p *Player) DiscardTile(tile tile.Interface) Interface {
	for i, currentTile := range p.tiles {
		if currentTile.Number() == tile.Number() && currentTile.Letter() == tile.Letter() {
			p.tiles = append(p.tiles[:i], p.tiles[i+1:]...)
			return p
		}
	}
	return p
}

// HasTile checks if the passed tile is in player's hand
func (p *Player) HasTile(tile tile.Interface) bool {
	for _, currentTile := range p.tiles {
		if currentTile.Number() == tile.Number() && currentTile.Letter() == tile.Letter() {
			return true
		}
	}
	return false
}

// Cash returns player's cash
func (p *Player) Cash() int {
	return p.cash
}

// AddCash adds cash to player
func (p *Player) AddCash(amount int) Interface {
	p.cash += amount
	return p
}

// RemoveCash removes cash from player
func (p *Player) RemoveCash(amount int) Interface {
	p.cash -= amount
	return p
}

// Name returns player name
func (p *Player) Name() string {
	return p.name
}
