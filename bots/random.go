// Package bots implements different types of AI for playing Acquire games
package bots

import (
	"github.com/svera/acquire/interfaces"
)

// Random is a struct which implements a very stupid AI, which basically
// chooses all its decisions randomly (So not that much an AI but an AS)
type Random struct {
	cash   int
	tiles  []interfaces.Tile
	shares map[interfaces.Corporation]int
	board  interfaces.Board
}

func NewRandom() *Random {
	return &Random{
		tiles:  []interfaces.Tile{},
		shares: map[interfaces.Corporation]int{},
	}
}

func (r *Random) SetCash(cash int) {
	r.cash = cash
}

func (r *Random) SetTiles(tiles []interfaces.Tile) {
	r.tiles = tiles
}

func (r *Random) PlayTile() interfaces.Tile {
	return r.tiles[0]
}
