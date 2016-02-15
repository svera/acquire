package acquire

// Owner interface declares all methods to be implemented by an owner implementation.
// Owner acts as a "marker" for board, because each board cell can contain an
// Owner instance, that is, an instance of Tile, Corporation or Empty structs.
type Owner interface {
	Type() string
}
