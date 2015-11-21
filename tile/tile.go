// Model for the tile. Every tile stores its coordinates and which is its owner
// (any implementation of the Owner interface, which can be Empty, Orphan or a
// Corporation implementation)
package tile

type Tile struct {
	number  int
	letter  string
	owner Owner
}

func New(number int, letter string, owner Owner) *Tile {
	return &Tile{number, letter, owner}
}

// Returns tile position number
func (t *Tile) Number() int {
	return t.number
}

// Returns tile position letter
func (t *Tile) Letter() string {
	return t.letter
}

// Sets an owner for the tile
func (t *Tile) SetOwner(owner Owner) {
	t.owner = owner
}

// Returns tile owner
func (t *Tile) Owner() Owner {
	return t.owner
}

type Empty struct{}

func (e Empty) Type() string {
	return "empty"
}

type Orphan struct{}

func (o Orphan) Type() string {
	return "orphan"
}
