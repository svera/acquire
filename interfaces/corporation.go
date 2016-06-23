package interfaces

// Prices is a struct that stores the different prices and bonuses of a corporation
type Prices struct {
	Price         int
	MajorityBonus int
	MinorityBonus int
}

// Corporation declares all methods to be implemented by a corporation implementation
type Corporation interface {
	Grow(number int)
	Reset()
	Stock() int
	AddStock(amount int)
	RemoveStock(amount int)
	StockPrice() int
	MajorityBonus() int
	MinorityBonus() int
	IsSafe() bool
	IsActive() bool
	Size() int
	Type() string
	SetPricesChart(prices map[int]Prices)
}
