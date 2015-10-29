package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tileset"
)

type Interface interface {
	Cell(t tileset.Position) int
	TileFoundCorporation(t tileset.Position) (bool, []tileset.Position)
	TileMergeCorporations(t tileset.Position) (bool, []int)
	TileGrowCorporation(t tileset.Position) (bool, []tileset.Position, int)
	PutTile(t tileset.Position)
	AdjacentCells(t tileset.Position) []tileset.Position
	SetTiles(cp corporation.Interface, cells []tileset.Position)
}
