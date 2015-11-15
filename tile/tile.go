package tile

type Orphan struct {
	Number int
	Letter string
}

func New(number int, letter string) *Orphan {
	return &Orphan{number, letter}
}

func (t *Orphan) ContentType() string {
	return "orphan"
}
