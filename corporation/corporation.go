// Package corporation contains the model Corporation and attahced methods which manages corporations in game
package corporation

import (
	"errors"
)

type prices struct {
	price         int
	majorityBonus int
	minorityBonus int
}

const (
	// WrongCorporationClass is an error returned when corporation class is not between 0 and 2
	WrongCorporationClass = "wrong_corporation_class"
)

// Corporation holds data related to corporations
type Corporation struct {
	name        string
	class       int
	stock       int
	pricesChart map[int]prices
	size        int
}

// New initialises and returns a new instance of Corporation
func New(name string, class int) (*Corporation, error) {
	if class < 0 || class > 2 {
		return nil, errors.New(WrongCorporationClass)
	}

	corporation := &Corporation{
		name:        name,
		stock:       25,
		class:       class,
		pricesChart: initPricesChart(class),
	}

	return corporation, nil
}

// Size returns corporation size on board
func (c *Corporation) Size() int {
	return c.size
}

// Grow increases corporation size in tiles
func (c *Corporation) Grow(number int) Interface {
	c.size += number
	return c
}

// Reset sets corporation size to 0 (not on board)
func (c *Corporation) Reset() Interface {
	c.size = 0
	return c
}

//Fills the prices chart array with the amounts corresponding to the corporation
//class
func initPricesChart(class int) map[int]prices {
	initialValues := new([3]prices)
	initialValues[0] = prices{price: 200, majorityBonus: 2000, minorityBonus: 1000}
	initialValues[1] = prices{price: 300, majorityBonus: 3000, minorityBonus: 1500}
	initialValues[2] = prices{price: 400, majorityBonus: 4000, minorityBonus: 2000}
	pricesChart := make(map[int]prices)

	pricesChart[2] = prices{price: initialValues[class].price, majorityBonus: initialValues[class].majorityBonus, minorityBonus: initialValues[class].minorityBonus}
	pricesChart[3] = prices{price: initialValues[class].price + 100, majorityBonus: initialValues[class].majorityBonus + 1000, minorityBonus: initialValues[class].minorityBonus + 500}
	pricesChart[4] = prices{price: initialValues[class].price + 200, majorityBonus: initialValues[class].majorityBonus + 2000, minorityBonus: initialValues[class].minorityBonus + 1000}
	pricesChart[5] = prices{price: initialValues[class].price + 300, majorityBonus: initialValues[class].majorityBonus + 3000, minorityBonus: initialValues[class].minorityBonus + 1500}
	var i int
	for i = 6; i < 11; i++ {
		pricesChart[i] = prices{price: initialValues[class].price + 400, majorityBonus: initialValues[class].majorityBonus + 4000, minorityBonus: initialValues[class].minorityBonus + 2000}
	}
	for i = 11; i < 21; i++ {
		pricesChart[i] = prices{price: initialValues[class].price + 500, majorityBonus: initialValues[class].majorityBonus + 5000, minorityBonus: initialValues[class].minorityBonus + 2500}
	}
	for i = 21; i < 31; i++ {
		pricesChart[i] = prices{price: initialValues[class].price + 600, majorityBonus: initialValues[class].majorityBonus + 6000, minorityBonus: initialValues[class].minorityBonus + 3000}
	}
	for i = 31; i < 41; i++ {
		pricesChart[i] = prices{price: initialValues[class].price + 700, majorityBonus: initialValues[class].majorityBonus + 7000, minorityBonus: initialValues[class].minorityBonus + 3500}
	}
	pricesChart[41] = prices{price: initialValues[class].price + 800, majorityBonus: initialValues[class].majorityBonus + 8000, minorityBonus: initialValues[class].minorityBonus + 4000}
	return pricesChart
}

// Stock returns corporation's amount of stock shares available
func (c *Corporation) Stock() int {
	return c.stock
}

// AddStock adds amount of stock shares to corporation stock
func (c *Corporation) AddStock(amount int) Interface {
	c.stock += amount
	return c
}

// RemoveStock removes the passed amount of stock shares from corporation stock
func (c *Corporation) RemoveStock(amount int) Interface {
	c.stock -= amount
	return c
}

// StockPrice returns company's current value per stock share
func (c *Corporation) StockPrice() int {
	if c.Size() > 41 {
		return c.pricesChart[41].price
	}
	return c.pricesChart[c.Size()].price
}

// MajorityBonus returns company's current majority bonus value per stock share
func (c *Corporation) MajorityBonus() int {
	if c.Size() > 41 {
		return c.pricesChart[41].majorityBonus
	}
	return c.pricesChart[c.Size()].majorityBonus
}

// MinorityBonus returns company's current minority bonus value per stock share
func (c *Corporation) MinorityBonus() int {
	if c.Size() > 41 {
		return c.pricesChart[41].minorityBonus
	}
	return c.pricesChart[c.Size()].minorityBonus
}

// IsSafe returns true if the corporation is considered safe, false otherwise
func (c *Corporation) IsSafe() bool {
	return c.Size() >= 11
}

// IsActive returns true if the corporation is on the board, false otherwise
func (c *Corporation) IsActive() bool {
	return c.Size() > 0
}

// Name returns corporation name
func (c *Corporation) Name() string {
	return c.name
}

// Class returns corporation class
func (c *Corporation) Class() int {
	return c.class
}

// Type returns type, to comply with owner interface
func (c *Corporation) Type() string {
	return "corporation"
}
