package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

type Interface interface {
	Cell(number int, letter string) tile.Interface
	TileFoundCorporation(t tile.Interface) (bool, []tile.Interface)
	TileMergeCorporations(t tile.Interface) (bool, []corporation.Interface)
	TileGrowCorporation(t tile.Interface) (bool, []tile.Interface, corporation.Interface)
	PutTile(t tile.Interface)
	AdjacentCells(t tile.Interface) []tile.Interface
	SetTiles(cp corporation.Interface, tiles []tile.Interface)
}
