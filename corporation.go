package acquire

// Corporation declares all methods to be implemented by a corporation implementation
type Corporation interface {
	Grow(number int) Corporation
	Reset() Corporation
	Stock() int
	AddStock(amount int) Corporation
	RemoveStock(amount int) Corporation
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
