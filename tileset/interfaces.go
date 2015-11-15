package tileset

import (
	"github.com/svera/acquire/tile"
)

type Interface interface {
	Draw() (tile.Orphan, error)
}
