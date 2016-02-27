package mocks

type Tile struct {
	FakeNumber int
	FakeLetter string
}

func (t *Tile) Number() int {
	return t.FakeNumber
}
func (t *Tile) Letter() string {
	return t.FakeLetter
}

func (t *Tile) Type() string {
	return "unincorporated"
}
