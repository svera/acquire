package corporation

type Interface interface {
	Id() int
	Grow(number int)
	Stock() int
	SetStock(stock int)
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
