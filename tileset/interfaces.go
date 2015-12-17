package tileset

import (
	"github.com/svera/acquire/tile"
)

// Interface declares all methods to be implemented by a tileset implementation
type Interface interface {
	Draw() (tile.Interface, error)
}
