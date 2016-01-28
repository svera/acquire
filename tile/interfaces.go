package tile

// Interface declares all methods to be implemented by a tile implementation
type Interface interface {
	Number() int
	Letter() string
	Type() string
}
