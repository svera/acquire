package mocks

import (
	"github.com/svera/acquire/interfaces"
)

// Tile is a structure that implements the Tile interface for testing
type Tile struct {
	FakeNumber int
	FakeLetter string
}

// Number mocks the Number method defined in the Tile interface
func (t *Tile) Number() int {
	return t.FakeNumber
}

// Letter mocks the Letter method defined in the Tile interface
func (t *Tile) Letter() string {
	return t.FakeLetter
}

// Type mocks the Type method defined in the Tile interface
func (t *Tile) Type() string {
	return interfaces.UnincorporatedOwner
}
