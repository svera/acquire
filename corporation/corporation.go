package corporation

import (
	"errors"
	"github.com/svera/acquire/tileset"
)

type prices struct {
	price         int
	majorityBonus int
	minorityBonus int
}

const (
	WrongCorporationClass = "wrong_corporation_class"
)

type Corporation struct {
	id          int
	name        string
	stock       int
	pricesChart map[int]prices
	tiles       []tileset.Position
}

func New(name string, class int) (*Corporation, error) {
	if class < 0 || class > 2 {
		return nil, errors.New(WrongCorporationClass)
	}
	corporation := &Corporation{
		name:        name,
		stock:       25,
		pricesChart: initPricesChart(class),
	}

	return corporation, nil
}

func (c *Corporation) Size() int {
	return len(c.tiles)
}

func (c *Corporation) SetId(id int) {
	c.id = id
}

func (c *Corporation) Id() int {
	return c.id
}

func (c *Corporation) AddTiles(tiles []tileset.Position) {
	c.tiles = append(c.tiles, tiles...)
}

func (c *Corporation) AddTile(tile tileset.Position) {
	c.tiles = append(c.tiles, tile)
}

//Fill the prices chart array with the amounts corresponding to the corporation
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

func (c *Corporation) Stock() int {
	return c.stock
}

func (c *Corporation) SetStock(stock int) {
	c.stock = stock
}

// Returns company's current value per stock share
func (c *Corporation) StockPrice() int {
	if c.Size() > 41 {
		return c.pricesChart[41].price
	}
	return c.pricesChart[c.Size()].price
}

// Returns company's current majority bonus value per stock share
func (c *Corporation) MajorityBonus() int {
	if c.Size() > 41 {
		return c.pricesChart[41].majorityBonus
	}
	return c.pricesChart[c.Size()].majorityBonus
}

// Returns company's current minority bonus value per stock share
func (c *Corporation) MinorityBonus() int {
	if c.Size() > 41 {
		return c.pricesChart[41].minorityBonus
	}
	return c.pricesChart[c.Size()].minorityBonus
}

// Returns true if the corporation is considered safe, false otherwise
func (c *Corporation) IsSafe() bool {
	return c.Size() >= 11
}

// Returns true if the corporation is on the board, false otherwise
func (c *Corporation) IsActive() bool {
	return c.Size() > 0
}

func (c *Corporation) Name() string {
	return c.name
}
