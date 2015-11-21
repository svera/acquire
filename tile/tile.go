package tile

type Tile struct {
	number  int
	letter  string
	owner Owner
}

func New(number int, letter string, owner Owner) *Tile {
	return &Tile{number, letter, owner}
}

func (t *Tile) Number() int {
	return t.number
}

func (t *Tile) Letter() string {
	return t.letter
}

func (t *Tile) SetOwner(owner Owner) {
	t.owner = owner
}

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
