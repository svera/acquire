// Package interfaces holds the declaration of all the interfaces used throught the Acquire library.
package interfaces

// Owner interface declares all methods to be implemented by an owner implementation.
// Owner acts as a "marker" for board, because each board cell can contain an
// Owner instance, that is, an instance of Tile, Corporation or Empty structs.
type Owner interface {
	Type() string
}

const (
	EmptyOwner          = "empty"
	UnincorporatedOwner = "unincorporated"
	CorporationOwner    = "corporation"
)
