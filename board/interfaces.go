package board

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/owner"
	"github.com/svera/acquire/tile"
)

// Interface declares all methods to be implemented by a board implementation
type Interface interface {
	Cell(number int, letter string) owner.Interface
	TileFoundCorporation(t tile.Interface) (bool, []tile.Interface)
	TileMergeCorporations(t tile.Interface) (bool, map[string][]corporation.Interface)
	TileGrowCorporation(t tile.Interface) (bool, []tile.Interface, corporation.Interface)
	PutTile(t tile.Interface) Interface
	AdjacentCells(number int, letter string) []owner.Interface
	SetOwner(cp corporation.Interface, tiles []tile.Interface) Interface
	ChangeOwner(oldOwner corporation.Interface, newOwner corporation.Interface) Interface
}
