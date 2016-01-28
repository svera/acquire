// Package tile contains the struct and related methods which models a tile.
// Every tile stores its coordinates and owner
// (any implementation of the Owner interface, which can be Empty, Unincorporated or a
// Corporation implementation)
package tile

// Tile stores position and owner of a tile
type Tile struct {
	number int
	letter string
}

// New initialises and returns a Tile instance
func New(number int, letter string) *Tile {
	return &Tile{number, letter}
}

// Number returns tile position number
func (t *Tile) Number() int {
	return t.number
}

// Letter returns tile position letter
func (t *Tile) Letter() string {
	return t.letter
}

// Type returns owner interface Type method value
func (t *Tile) Type() string {
	return "unincorporated"
}
