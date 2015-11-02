package corporation

import (
	"github.com/svera/acquire/tileset"
)

type Interface interface {
	Id() int
	AddTiles(tiles []tileset.Position)
	AddTile(tile tileset.Position)
	Stock() int
	SetStock(stock int)
	StockPrice() int
	MajorityBonus() int
	MinorityBonus() int
	IsSafe() bool
	IsActive() bool
	Name() string
	Size() int
}
