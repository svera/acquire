package corporation

type Interface interface {
	Grow(number int)
	Stock() int
	AddStock(amount int)
	RemoveStock(amount int)
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
