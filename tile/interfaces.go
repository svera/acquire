package tile

// Owner interface declares all methods to be implemented by an owner implementation
type Owner interface {
	Type() string
}

// Interface declares all methods to be implemented by a tile implementation
type Interface interface {
	Number() int
	Letter() string
	SetOwner(owner Owner) Interface
	Owner() Owner
}
