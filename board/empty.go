package board

import (
	"github.com/svera/acquire/interfaces"
)

// Empty is a struct that represents an empty cell on board
type Empty struct{}

// Type returns tile type
func (e Empty) Type() string {
	return interfaces.EmptyOwner
}
