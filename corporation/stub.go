package corporation

import (
	"github.com/svera/acquire/tileset"
)

type Stub struct {
	Corporation
	size int
}

func NewStub(name string, class int, id int) *Stub {
	return &Stub{
		Corporation: Corporation{
			id:          id,
			name:        name,
			class:       class,
			stock:       25,
			pricesChart: initPricesChart(class),
		},
	}
}

func (c *Stub) SetSize(size int) {
	c.tiles = make([]tileset.Position, size)
}
