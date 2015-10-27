package player

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tileset"
)

type Interface interface {
	Shares(c *corporation.Corporation) int
	ReceiveBonus(amount int)
	Buy(corp *corporation.Corporation, amount int)
	PickTile(t tileset.Position) error
	Tiles() []tileset.Position
	UseTile(t tileset.Position) error
	Cash() int
}
