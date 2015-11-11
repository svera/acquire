package player

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tileset"
)

type Interface interface {
	Shares(c corporation.Interface) int
	ReceiveBonus(amount int)
	Buy(corp corporation.Interface, amount int)
	PickTile(t board.Coordinates) error
	Tiles() []board.Coordinates
	DiscardTile(t board.Coordinates) error
	Cash() int
	GetFounderStockShare(corp corporation.Interface)
}

type ShareInterface interface {
	Shares(c corporation.Interface) int
	ReceiveBonus(amount int)
}
