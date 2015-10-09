package game

import "errors"

type prices struct {
	price         uint
	majorityBonus uint
	minorityBonus uint
}

type Corporation struct {
	id          uint
	name        string
	stock       uint
	pricesChart map[uint]prices
	size        uint
}

func NewCorporation(name string, class uint) (*Corporation, error) {
	if class < 0 || class > 2 {
		return nil, errors.New("Corporation can only be of class 0, 1 or 2")
	}
	return &Corporation{
		name:        name,
		stock:       25,
		pricesChart: initPricesChart(class),
	}, nil
}

func (c *Corporation) setId(id uint) {
	c.id = id
}

func (c *Corporation) getId() uint {
	return c.id
}

//Fill the prices chart array with the amounts corresponding to the corporation
//class
func initPricesChart(class uint) map[uint]prices {
	initialValues := new([3]prices)
	initialValues[0] = prices{price: 200, majorityBonus: 2000, minorityBonus: 1000}
	initialValues[1] = prices{price: 300, majorityBonus: 3000, minorityBonus: 1500}
	initialValues[2] = prices{price: 400, majorityBonus: 4000, minorityBonus: 2000}
	pricesChart := make(map[uint]prices)

	pricesChart[2] = prices{price: initialValues[class].price, majorityBonus: initialValues[class].majorityBonus, minorityBonus: initialValues[class].minorityBonus}
	pricesChart[3] = prices{price: initialValues[class].price + 100, majorityBonus: initialValues[class].majorityBonus + 1000, minorityBonus: initialValues[class].minorityBonus + 500}
	pricesChart[4] = prices{price: initialValues[class].price + 200, majorityBonus: initialValues[class].majorityBonus + 2000, minorityBonus: initialValues[class].minorityBonus + 1000}
	pricesChart[5] = prices{price: initialValues[class].price + 300, majorityBonus: initialValues[class].majorityBonus + 3000, minorityBonus: initialValues[class].minorityBonus + 1500}
	var i uint
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

func (c *Corporation) GetStock() uint {
	return c.stock
}

// Returns company's current value per stock share
func (c *Corporation) GetStockPrice() uint {
	if c.size == 0 {
		return 0
	}
	if c.size > 41 {
		return c.pricesChart[41].price
	}
	return c.pricesChart[c.size].price
}
