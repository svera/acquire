package corporation

type Stub struct {
	Corporation
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
	c.size = size
}
