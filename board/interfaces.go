package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

type Interface interface {
	Cell(t Coordinates) int
	TileFoundCorporation(t *tile.Orphan) (bool, []Coordinates)
	TileMergeCorporations(t *tile.Orphan) (bool, []Container)
	TileGrowCorporation(t *tile.Orphan) (bool, []Container, int)
	PutTile(t *tile.Orphan)
	AdjacentCells(t *tile.Orphan) []Container
	SetTiles(cp corporation.Interface, cells []Coordinates)
}

type Container interface {
	Number() int
	Letter() string
	ContentType() string
}
