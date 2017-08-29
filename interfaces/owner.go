// Package interfaces holds the declaration of all the interfaces used through the Acquire library.
package interfaces

// Owner interface declares all methods to be implemented by an owner implementation.
// Owner acts as a "marker" for board, because each board cell can contain an
// Owner instance, that is, an instance of Tile, Corporation or Empty structs.
type Owner interface {
	Type() string
}

// Constants which define identificators for all cell owner types.
const (
	EmptyOwner          = "empty"
	UnincorporatedOwner = "unincorporated"
	CorporationOwner    = "corporation"
)
