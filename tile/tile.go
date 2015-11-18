package tile

type TileContent interface {
	Type() string
}

type Tile struct {
	number  int
	letter  string
	content TileContent
}

func New(number int, letter string, content TileContent) *Tile {
	return &Tile{number, letter, content}
}

func (t *Tile) Number() int {
	return t.number
}

func (t *Tile) Letter() string {
	return t.letter
}

func (t *Tile) SetContent(content TileContent) {
	t.content = content
}

func (t *Tile) Content() TileContent {
	return t.content
}

type Empty struct{}

func (e Empty) Type() string {
	return "empty"
}

type Orphan struct{}

func (o Orphan) Type() string {
	return "orphan"
}
