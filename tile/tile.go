// Package tile contains the struct and related methods which models a tile.
// Every tile stores its coordinates and owner
// (any implementation of the Owner interface, which can be Empty, Unincorporated or a
// Corporation implementation)
package tile

// Tile stores position and owner of a tile
type Tile struct {
	number int
	letter string
	owner  Owner
}

// New initialises and returns a Tile instance
func New(number int, letter string, owner Owner) *Tile {
	return &Tile{number, letter, owner}
}

// Number returns tile position number
func (t *Tile) Number() int {
	return t.number
}

// Letter returns tile position letter
func (t *Tile) Letter() string {
	return t.letter
}

// SetOwner sets an owner for the tile
func (t *Tile) SetOwner(owner Owner) Interface {
	t.owner = owner
	return t
}

// Owner returns tile owner
func (t *Tile) Owner() Owner {
	return t.owner
}

// Empty is a struct that represents an empty cell on board
type Empty struct{}

// Type returns tile type
func (e Empty) Type() string {
	return "empty"
}

// Unincorporated is a struct that represents an unincorporated cell on board
type Unincorporated struct{}

// Type returns tile type
func (o Unincorporated) Type() string {
	return "unincorporated"
}
