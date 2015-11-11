package corporation

import (
	"github.com/svera/acquire/tileset"
)

type Interface interface {
	Id() int
	AddTiles(tiles []board.Coordinates)
	AddTile(tile board.Coordinates)
	Stock() int
	SetStock(stock int)
	StockPrice() int
	MajorityBonus() int
	MinorityBonus() int
	IsSafe() bool
	IsActive() bool
	Name() string
	Size() int
	Class() int
	ContentType() string
}
