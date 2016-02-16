package corporation

import (
	"github.com/svera/acquire/interfaces"
)

// Stub is a struct to be used in corporation tests as a replacement of the original,
// that includes several convenience methods for testing
type Stub struct {
	Corporation
}

// NewStub initialises and returns a new instance of Stub
func NewStub(name string, class int) *Stub {
	return &Stub{
		Corporation: Corporation{
			name:        name,
			class:       class,
			stock:       25,
			pricesChart: initPricesChart(class),
		},
	}
}

// SetSize sets the size of the stub
func (c *Stub) SetSize(size int) interfaces.Corporation {
	c.size = size
	return c
}

// SetStock sets the stock amount of the stub
func (c *Stub) SetStock(amount int) interfaces.Corporation {
	c.stock = amount
	return c
}
