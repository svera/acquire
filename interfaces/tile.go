package interfaces

// Tile interface declares all methods to be implemented by a tile implementation
type Tile interface {
	Number() int
	Letter() string
	Type() string
}
