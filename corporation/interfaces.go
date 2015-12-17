package corporation

// Interface declares all methods to be implemented by a corporation implementation
type Interface interface {
	Grow(number int) Interface
	Reset() Interface
	Stock() int
	AddStock(amount int) Interface
	RemoveStock(amount int) Interface
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
