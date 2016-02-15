package acquire

// Tileset interface declares all methods to be implemented by a tileset implementation
type Tileset interface {
	Draw() (Tile, error)
}
