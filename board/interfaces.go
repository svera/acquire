package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

type Interface interface {
	Cell(number int, letter string) *tile.Tile
	TileFoundCorporation(t *tile.Tile) (bool, []*tile.Tile)
	TileMergeCorporations(t *tile.Tile) (bool, []corporation.Interface)
	TileGrowCorporation(t *tile.Tile) (bool, []*tile.Tile, corporation.Interface)
	PutTile(t *tile.Tile)
	AdjacentCells(t *tile.Tile) []*tile.Tile
	SetTiles(cp corporation.Interface, tiles []*tile.Tile)
}
