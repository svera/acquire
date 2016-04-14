// Package mocks holds all mock implementations of the interfaces used by the
// Acquire library
package mocks

import (
	"github.com/svera/acquire/interfaces"
)

// Board is a structure that implements the Board interface for testing
type Board struct {
	FakeCellOwner              interfaces.Owner
	FakeFoundCorporation       bool
	FakeFoundCorporationTiles  []interfaces.Tile
	FakeMergeCorporations      bool
	FakeMergeCorporationsCorps map[string][]interfaces.Corporation
	FakeGrowCorporation        bool
	FakeGrowCorporationTiles   []interfaces.Tile
	FakeGrowCorporationCorp    interfaces.Corporation
	FakeAdjacentCells          []interfaces.Owner
	FakeAdjacentCorporations   []interfaces.Corporation
	TimesCalled                map[string]int
}

// Cell mocks the Cell method defined in the Board interface
func (b *Board) Cell(number int, letter string) interfaces.Owner {
	return b.FakeCellOwner
}

// TileFoundCorporation mocks the TileFoundCorporation method defined in the Board interface
func (b *Board) TileFoundCorporation(t interfaces.Tile) (bool, []interfaces.Tile) {
	return b.FakeFoundCorporation, b.FakeFoundCorporationTiles
}

// TileMergeCorporations mocks the TileMergeCorporations method defined in the Board interface
func (b *Board) TileMergeCorporations(t interfaces.Tile) (bool, map[string][]interfaces.Corporation) {
	return b.FakeMergeCorporations, b.FakeMergeCorporationsCorps
}

// TileGrowCorporation mocks the TileGrowCorporation method defined in the Board interface
func (b *Board) TileGrowCorporation(t interfaces.Tile) (bool, []interfaces.Tile, interfaces.Corporation) {
	return b.FakeGrowCorporation, b.FakeGrowCorporationTiles, b.FakeGrowCorporationCorp
}

// PutTile mocks the PutTile method defined in the Board interface
func (b *Board) PutTile(t interfaces.Tile) interfaces.Board {
	_ = t
	return b
}

// AdjacentCells mocks the AdjacentCells method defined in the Board interface
func (b *Board) AdjacentCells(number int, letter string) []interfaces.Owner {
	_, _ = number, letter
	return b.FakeAdjacentCells
}

// SetOwner mocks the SetOwner method defined in the Board interface
func (b *Board) SetOwner(cp interfaces.Corporation, tiles []interfaces.Tile) interfaces.Board {
	_, _ = cp, tiles
	b.TimesCalled["SetOwner"]++
	return b
}

// ChangeOwner mocks the ChangeOwner method defined in the Board interface
func (b *Board) ChangeOwner(oldOwner interfaces.Corporation, newOwner interfaces.Corporation) interfaces.Board {
	_, _ = oldOwner, newOwner
	b.TimesCalled["ChangeOwner"]++
	return b
}

// AdjacentCorporations mocks the AdjacentCorporations method defined in the Board interface
func (b *Board) AdjacentCorporations(number int, letter string) []interfaces.Corporation {
	_, _ = number, letter
	return b.FakeAdjacentCorporations
}
