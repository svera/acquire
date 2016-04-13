package acquire

import (
	"github.com/svera/acquire/interfaces"
)

type Optional struct {
	Board        interfaces.Board
	Corporations [7]interfaces.Corporation
	Tileset      interfaces.Tileset
	State        interfaces.State
}
