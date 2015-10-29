package player

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tileset"
)

type Interface interface {
	Shares(c corporation.Interface) int
	ReceiveBonus(amount int)
	Buy(corp corporation.Interface, amount int)
	PickTile(t tileset.Position) error
	Tiles() []tileset.Position
	UseTile(t tileset.Position) error
	Cash() int
}

type ShareInterface interface {
	Shares(c corporation.Interface) int
	ReceiveBonus(amount int)
}
