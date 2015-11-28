package corporation

type Stub struct {
	Corporation
}

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

func (c *Stub) SetSize(size int) Interface {
	c.size = size
	return c
}

func (c *Stub) SetStock(amount int) Interface {
	c.stock = amount
	return c
}
