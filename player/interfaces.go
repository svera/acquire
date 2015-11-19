package player

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

type Interface interface {
	Shares(c corporation.Interface) int
	ReceiveBonus(amount int)
	Buy(corp corporation.Interface, amount int)
	PickTile(t tile.Interface) error
	Tiles() []tile.Interface
	DiscardTile(t tile.Interface) error
	Cash() int
	GetFounderStockShare(corp corporation.Interface)
}

type ShareInterface interface {
	Shares(c corporation.Interface) int
	ReceiveBonus(amount int)
}
