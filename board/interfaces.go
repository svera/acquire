package board

import (
	"github.com/svera/acquire/corporation"
)

type Interface interface {
	Cell(t Coordinates) int
	TileFoundCorporation(t Coordinates) (bool, []Coordinates)
	TileMergeCorporations(t Coordinates) (bool, []int)
	TileGrowCorporation(t Coordinates) (bool, []Coordinates, int)
	PutTile(t Coordinates)
	AdjacentCells(t Coordinates) []Container
	SetTiles(cp corporation.Interface, cells []Coordinates)
}

type Container interface {
	ContentType() string
}
