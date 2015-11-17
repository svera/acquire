package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

type Interface interface {
	Cell(t Coordinates) int
	TileFoundCorporation(t *tile.Orphan) (bool, []Container)
	TileMergeCorporations(t *tile.Orphan) (bool, []corporation.Interface)
	TileGrowCorporation(t *tile.Orphan) (bool, []Container, corporation.Interface)
	PutTile(t *tile.Orphan)
	AdjacentCells(t *tile.Orphan) []Container
	SetTiles(tls []Container)
}

type Container interface {
	Number() int
	Letter() string
	ContentType() string
}
