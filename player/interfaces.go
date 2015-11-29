package player

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

type Interface interface {
	Shares(c corporation.Interface) int
	AddShares(corp corporation.Interface, amount int) Interface
	RemoveShares(corp corporation.Interface, amount int) Interface
	PickTile(t tile.Interface) Interface
	Tiles() []tile.Interface
	DiscardTile(t tile.Interface) Interface
	HasTile(t tile.Interface) bool
	Cash() int
	AddCash(amount int) Interface
	RemoveCash(amount int) Interface
}

type ShareInterface interface {
	Shares(c corporation.Interface) int
	AddCash(amount int) Interface
}
