package mocks

import (
	"github.com/svera/acquire/interfaces"
)

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
	TimesCalled                map[string]int
}

func (b *Board) Cell(number int, letter string) interfaces.Owner {
	return b.FakeCellOwner
}

func (b *Board) TileFoundCorporation(t interfaces.Tile) (bool, []interfaces.Tile) {
	return b.FakeFoundCorporation, b.FakeFoundCorporationTiles
}

func (b *Board) TileMergeCorporations(t interfaces.Tile) (bool, map[string][]interfaces.Corporation) {
	return b.FakeMergeCorporations, b.FakeMergeCorporationsCorps
}

func (b *Board) TileGrowCorporation(t interfaces.Tile) (bool, []interfaces.Tile, interfaces.Corporation) {
	return b.FakeGrowCorporation, b.FakeGrowCorporationTiles, b.FakeGrowCorporationCorp
}

func (b *Board) PutTile(t interfaces.Tile) interfaces.Board {
	_ = t
	return b
}

func (b *Board) AdjacentCells(number int, letter string) []interfaces.Owner {
	_, _ = number, letter
	return b.FakeAdjacentCells
}

func (b *Board) SetOwner(cp interfaces.Corporation, tiles []interfaces.Tile) interfaces.Board {
	_, _ = cp, tiles
	b.TimesCalled["SetOwner"]++
	return b
}

func (b *Board) ChangeOwner(oldOwner interfaces.Corporation, newOwner interfaces.Corporation) interfaces.Board {
	_, _ = oldOwner, newOwner
	b.TimesCalled["ChangeOwner"]++
	return b
}
