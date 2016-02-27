package mocks

import (
  "github.com/svera/acquire/interfaces"
)

type Corporation struct {
	FakeSize          int
	FakeStock         int
	FakeStockPrice    int
	FakeMajorityBonus int
	FakeMinorityBonus int
	FakeIsSafe        bool
	FakeIsActive      bool
	FakeName          string
	FakeClass         int
	TimesCalled       map[string]int
}

func (c *Corporation) Grow(number int) interfaces.Corporation {
	c.FakeSize += number
	c.TimesCalled["Grow"]++
	return c
}

func (c *Corporation) Reset() interfaces.Corporation {
	c.FakeSize = 0
	c.TimesCalled["Reset"]++
	return c
}

func (c *Corporation) Stock() int {
	return c.FakeStock
}

func (c *Corporation) StockPrice() int {
	return c.FakeStockPrice
}

func (c *Corporation) MajorityBonus() int {
	return c.FakeMajorityBonus
}

func (c *Corporation) MinorityBonus() int {
	return c.FakeMinorityBonus
}

func (c *Corporation) IsSafe() bool {
	return c.FakeIsSafe
}

func (c *Corporation) IsActive() bool {
	return c.FakeIsActive
}

func (c *Corporation) Name() string {
	return c.FakeName
}

func (c *Corporation) AddStock(amount int) interfaces.Corporation {
	_ = amount
	return c
}

func (c *Corporation) RemoveStock(amount int) interfaces.Corporation {
	c.FakeStock -= amount
	return c
}

func (c *Corporation) Type() string {
	return "corporation"
}

func (c *Corporation) Size() int {
	return c.FakeSize
}

func (c *Corporation) Class() int {
	return c.FakeClass
}
