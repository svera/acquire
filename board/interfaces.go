package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tileset"
)

type Interface interface {
	Cell(t board.Coordinates) int
	TileFoundCorporation(t board.Coordinates) (bool, []board.Coordinates)
	TileMergeCorporations(t board.Coordinates) (bool, []int)
	TileGrowCorporation(t board.Coordinates) (bool, []board.Coordinates, int)
	PutTile(t board.Coordinates)
	AdjacentCells(t board.Coordinates) []board.Coordinates
	SetTiles(cp corporation.Interface, cells []board.Coordinates)
}

type Container interface {
	ContentType() string
}
