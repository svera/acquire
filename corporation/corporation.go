// Package corporation contains the model Corporation and attahced methods which manages corporations in game
package corporation

import "github.com/svera/acquire/interfaces"

// Corporation holds data related to corporations
type Corporation struct {
	stock       int
	pricesChart map[int]interfaces.Prices
	size        int
}

// New initialises and returns a new instance of Corporation
func New() *Corporation {
	corporation := &Corporation{
		stock:       25,
		pricesChart: make(map[int]interfaces.Prices),
	}

	return corporation
}

// Size returns corporation size on board
func (c *Corporation) Size() int {
	return c.size
}

// Grow increases corporation size in tiles
func (c *Corporation) Grow(number int) {
	c.size += number
}

// Reset sets corporation size to 0 (not on board)
func (c *Corporation) Reset() {
	c.size = 0
}

// Stock returns corporation's amount of stock shares available
func (c *Corporation) Stock() int {
	return c.stock
}

// AddStock adds amount of stock shares to corporation stock
func (c *Corporation) AddStock(amount int) {
	c.stock += amount
}

// RemoveStock removes the passed amount of stock shares from corporation stock
func (c *Corporation) RemoveStock(amount int) {
	c.stock -= amount
}

// StockPrice returns company's current value per stock share
func (c *Corporation) StockPrice() int {
	if c.Size() > 41 {
		return c.pricesChart[41].Price
	}
	return c.pricesChart[c.Size()].Price
}

// MajorityBonus returns company's current majority bonus value per stock share
func (c *Corporation) MajorityBonus() int {
	if c.Size() > 41 {
		return c.pricesChart[41].MajorityBonus
	}
	return c.pricesChart[c.Size()].MajorityBonus
}

// MinorityBonus returns company's current minority bonus value per stock share
func (c *Corporation) MinorityBonus() int {
	if c.Size() > 41 {
		return c.pricesChart[41].MinorityBonus
	}
	return c.pricesChart[c.Size()].MinorityBonus
}

// IsSafe returns true if the corporation is considered safe, false otherwise
func (c *Corporation) IsSafe() bool {
	return c.Size() >= 11
}

// IsActive returns true if the corporation is on the board, false otherwise
func (c *Corporation) IsActive() bool {
	return c.Size() > 0
}

// Type returns type, to comply with owner interface
func (c *Corporation) Type() string {
	return interfaces.CorporationOwner
}

// SetPricesChart sets the prices and bonuses of the corporation
func (c *Corporation) SetPricesChart(prices map[int]interfaces.Prices) {
	c.pricesChart = prices
}
