// Package player containst the struct Player and attached methods which manages player status in game.
package player

import (
	"github.com/svera/acquire"
)

// Player stores the status of a player
type Player struct {
	cash   int
	tiles  []acquire.Tile
	shares map[acquire.Corporation]int
}

// New initialises and returns a Player instance
func New() *Player {
	return &Player{
		cash:   6000,
		shares: map[acquire.Corporation]int{},
	}
}

// Shares returns the number of shares for the passed corporation owned by the player
func (p *Player) Shares(corp acquire.Corporation) int {
	return p.shares[corp]
}

// AddShares adds new stock shares of the passed corporation to the player
func (p *Player) AddShares(corp acquire.Corporation, amount int) acquire.Player {
	p.shares[corp] += amount
	return p
}

// RemoveShares removes stock shares of the passed corporation from the player
func (p *Player) RemoveShares(corp acquire.Corporation, amount int) acquire.Player {
	p.shares[corp] -= amount
	return p
}

// PickTile adds a new tile to the players' tileset
func (p *Player) PickTile(tile acquire.Tile) acquire.Player {
	p.tiles = append(p.tiles, tile)
	return p
}

// Tiles returns player's tiles
func (p *Player) Tiles() []acquire.Tile {
	return p.tiles
}

// DiscardTile discards the passed tile from player's hand
func (p *Player) DiscardTile(tile acquire.Tile) acquire.Player {
	for i, currentTile := range p.tiles {
		if currentTile.Number() == tile.Number() && currentTile.Letter() == tile.Letter() {
			p.tiles = append(p.tiles[:i], p.tiles[i+1:]...)
			return p
		}
	}
	return p
}

// HasTile checks if the passed tile is in player's hand
func (p *Player) HasTile(tile acquire.Tile) bool {
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
func (p *Player) AddCash(amount int) acquire.Player {
	p.cash += amount
	return p
}

// RemoveCash removes cash from player
func (p *Player) RemoveCash(amount int) acquire.Player {
	p.cash -= amount
	return p
}

func (p *Player) OrderByCashDesc(players []acquire.Player) []acquire.Player {
	var classification []acquire.Player

	cashDesc := func(pl1, pl2 acquire.Player) bool {
		return pl1.Cash() > pl2.Cash()
	}

	for _, pl := range players {
		classification = append(classification, pl)
	}
	By(cashDesc).Sort(classification)
	return classification
}
