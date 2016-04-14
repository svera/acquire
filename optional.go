package acquire

import (
	"github.com/svera/acquire/interfaces"
)

// Optional is a struct which stores the fields that are optional when creating
// a new Game instance with the New() method.
type Optional struct {
	Board        interfaces.Board
	Corporations [7]interfaces.Corporation
	Tileset      interfaces.Tileset
	State        interfaces.State
}
