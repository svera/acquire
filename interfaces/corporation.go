package interfaces

// Corporation declares all methods to be implemented by a corporation implementation
type Corporation interface {
	Grow(number int) Corporation
	Reset() Corporation
	Stock() int
	AddStock(amount int) Corporation
	RemoveStock(amount int) Corporation
	StockPrice() int
	MajorityBonus() int
	MinorityBonus() int
	IsSafe() bool
	IsActive() bool
	Name() string
	Size() int
	Class() int
	Type() string
}

type CorporationMock struct {
	FakeSize          int
	FakeStock         int
	FakeStockPrice    int
	FakeMajorityBonus int
	FakeMinorityBonus int
	FakeIsSafe        bool
	FakeIsActive      bool
	FakeName          string
	FakeClass         int
}

func (c *CorporationMock) Grow(number int) Corporation {
	c.FakeSize += number
	return c
}

func (c *CorporationMock) Reset() Corporation {
	return c
}

func (c *CorporationMock) Stock() int {
	return c.FakeStock
}

func (c *CorporationMock) StockPrice() int {
	return c.FakeStockPrice
}

func (c *CorporationMock) MajorityBonus() int {
	return c.FakeMajorityBonus
}

func (c *CorporationMock) MinorityBonus() int {
	return c.FakeMinorityBonus
}

func (c *CorporationMock) IsSafe() bool {
	return c.FakeIsSafe
}

func (c *CorporationMock) IsActive() bool {
	return c.FakeIsActive
}

func (c *CorporationMock) Name() string {
	return c.FakeName
}

func (c *CorporationMock) AddStock(amount int) Corporation {
	_ = amount
	return c
}

func (c *CorporationMock) RemoveStock(amount int) Corporation {
	_ = amount
	return c
}

func (c *CorporationMock) Type() string {
	return "corporation"
}

func (c *CorporationMock) Size() int {
	return c.FakeSize
}

func (c *CorporationMock) Class() int {
	return c.FakeClass
}
