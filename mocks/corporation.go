package mocks

import (
	"github.com/svera/acquire/interfaces"
)

// Corporation is a structure that implements the Corporation interface for testing
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

// Grow mocks the Grow method defined in the Corporation interface
func (c *Corporation) Grow(number int) interfaces.Corporation {
	c.FakeSize += number
	c.TimesCalled["Grow"]++
	return c
}

// Reset mocks the Reset method defined in the Corporation interface
func (c *Corporation) Reset() interfaces.Corporation {
	c.FakeSize = 0
	c.TimesCalled["Reset"]++
	return c
}

// Stock mocks the Stock method defined in the Corporation interface
func (c *Corporation) Stock() int {
	return c.FakeStock
}

// StockPrice mocks the StockPrice method defined in the Corporation interface
func (c *Corporation) StockPrice() int {
	return c.FakeStockPrice
}

// MajorityBonus mocks the MajorityBonus method defined in the Corporation interface
func (c *Corporation) MajorityBonus() int {
	return c.FakeMajorityBonus
}

// MinorityBonus mocks the MinorityBonus method defined in the Corporation interface
func (c *Corporation) MinorityBonus() int {
	return c.FakeMinorityBonus
}

// IsSafe mocks the IsSafe method defined in the Corporation interface
func (c *Corporation) IsSafe() bool {
	return c.FakeIsSafe
}

// IsActive mocks the IsActive method defined in the Corporation interface
func (c *Corporation) IsActive() bool {
	return c.FakeIsActive
}

// Name mocks the Name method defined in the Corporation interface
func (c *Corporation) Name() string {
	return c.FakeName
}

// AddStock mocks the AddStock method defined in the Corporation interface
func (c *Corporation) AddStock(amount int) interfaces.Corporation {
	_ = amount
	return c
}

// RemoveStock mocks the RemoveStock method defined in the Corporation interface
func (c *Corporation) RemoveStock(amount int) interfaces.Corporation {
	c.FakeStock -= amount
	return c
}

// Type mocks the Type method defined in the Corporation interface
func (c *Corporation) Type() string {
	return "corporation"
}

// Size mocks the Size method defined in the Corporation interface
func (c *Corporation) Size() int {
	return c.FakeSize
}

// Class mocks the Class method defined in the Corporation interface
func (c *Corporation) Class() int {
	return c.FakeClass
}
