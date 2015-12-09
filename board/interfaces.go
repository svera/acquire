package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

type Interface interface {
	Cell(number int, letter string) tile.Interface
	TileFoundCorporation(t tile.Interface) (bool, []tile.Interface)
	TileMergeCorporations(t tile.Interface) (bool, map[string][]corporation.Interface)
	TileGrowCorporation(t tile.Interface) (bool, []tile.Interface, corporation.Interface)
	PutTile(t tile.Interface) Interface
	AdjacentCells(t tile.Interface) []tile.Interface
	SetOwner(cp corporation.Interface, tiles []tile.Interface) Interface
	ChangeOwner(oldOwner corporation.Interface, newOwner corporation.Interface) Interface
}
