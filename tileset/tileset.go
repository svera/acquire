// Package tileset holds the Tileset struct and related methods
package tileset

import (
	"errors"
	"github.com/svera/acquire/tile"
	"math/rand"
	"time"
)

const (
	// NoTilesAvailable is an error returned when a player tries to get a tile from an empty tileset
	NoTilesAvailable = "no_tiles_available"
)

// Tileset stores all tiles used in game
type Tileset struct {
	tiles []tile.Interface
}

// New initialises and returns a Tileset instance
func New() *Tileset {
	tileset := Tileset{}
	letters := [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	for number := 1; number < 13; number++ {
		for _, letter := range letters {
			tileset.tiles = append(tileset.tiles, tile.New(number, letter))
		}
	}

	return &tileset
}

// Draw extracts a random tile from the tileset and returns it
func (t *Tileset) Draw() (tile.Interface, error) {
	source := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(source)
	remainingTiles := len(t.tiles)
	if remainingTiles > 0 {
		pos := rn.Intn(remainingTiles - 1)
		tile := t.tiles[pos]
		t.tiles = append(t.tiles[:pos], t.tiles[pos+1:]...)
		return tile, nil
	}
	return &tile.Tile{}, errors.New(NoTilesAvailable)
}
