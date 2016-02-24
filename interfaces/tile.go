package interfaces

// Tile interface declares all methods to be implemented by a tile implementation
type Tile interface {
	Number() int
	Letter() string
	Type() string
}

type TileMock struct {
	FakeNumber int
	FakeLetter string
}

func (t *TileMock) Number() int {
	return t.FakeNumber
}
func (t *TileMock) Letter() string {
	return t.FakeLetter
}

func (t *TileMock) Type() string {
	return "unincorporated"
}
