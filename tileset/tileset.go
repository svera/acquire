// Package tileset holds the Tileset struct and related methods
package tileset

import (
	"errors"
	"math/rand"
	"time"

	"github.com/svera/acquire/interfaces"
	"github.com/svera/acquire/tile"
)

const (
	// NoTilesAvailable is an error returned when a player tries to get a tile from an empty tileset
	NoTilesAvailable = "no_tiles_available"
)

// Tileset stores all tiles used in game
type Tileset struct {
	tiles []interfaces.Tile
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
func (t *Tileset) Draw() (interfaces.Tile, error) {
	source := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(source)
	remainingTiles := len(t.tiles)
	if remainingTiles == 0 {
		return &tile.Tile{}, errors.New(NoTilesAvailable)
	}

	var pos int
	if remainingTiles > 1 {
		pos = rn.Intn(remainingTiles - 1)
	} else if remainingTiles == 1 {
		pos = 0
	}
	tile := t.tiles[pos]
	t.tiles = append(t.tiles[:pos], t.tiles[pos+1:]...)
	return tile, nil
}

// DiscardTile removes passed tile from the tileset
func (t *Tileset) DiscardTile(tl interfaces.Tile) {
	for i, currentTile := range t.tiles {
		if currentTile.Number() == tl.Number() && currentTile.Letter() == tl.Letter() {
			t.tiles = append(t.tiles[:i], t.tiles[i+1:]...)
			break
		}
	}
}

// Add appends the passed tiles to the tileset
func (t *Tileset) Add(tiles []interfaces.Tile) interfaces.Tileset {
	t.tiles = append(t.tiles, tiles...)
	return t
}
