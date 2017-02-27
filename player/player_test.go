package player

import (
	"testing"

	"github.com/svera/acquire/interfaces"
	"github.com/svera/acquire/mocks"
)

func TestPickTile(t *testing.T) {
	player := New(0)
	tl := &mocks.Tile{FakeNumber: 2, FakeLetter: "C"}
	player.PickTile(tl)
	if len(player.tiles) != 1 {
		t.Errorf("Player must have exactly 1 tile, got %d", len(player.tiles))
	}
}

func TestUseTile(t *testing.T) {
	player := New(0)

	player.tiles = []interfaces.Tile{
		&mocks.Tile{FakeNumber: 7, FakeLetter: "c"},
		&mocks.Tile{FakeNumber: 5, FakeLetter: "A"},
		&mocks.Tile{FakeNumber: 8, FakeLetter: "E"},
		&mocks.Tile{FakeNumber: 3, FakeLetter: "D"},
		&mocks.Tile{FakeNumber: 1, FakeLetter: "B"},
		&mocks.Tile{FakeNumber: 4, FakeLetter: "I"},
	}

	tl := &mocks.Tile{FakeNumber: 5, FakeLetter: "A"}
	player.DiscardTile(tl)
	if len(player.tiles) != 5 {
		t.Errorf("Players must have 5 tiles after using one, got %d", len(player.tiles))
	}
	if tl.Number() != 5 || tl.Letter() != "A" {
		t.Errorf("DiscardTile() must return tile 5A")
	}
}

func TestShares(t *testing.T) {
	corp := &mocks.Corporation{}
	expected := 5
	player := &Player{
		shares: map[interfaces.Corporation]int{
			corp: expected,
		},
	}
	if player.Shares(corp) != expected {
		t.Errorf("Shares() must return that the player has exactly %d stock shares in corporation %p, got %d", expected, corp, player.Shares(corp))
	}
}

func TestAddShares(t *testing.T) {
	corp := &mocks.Corporation{}
	original := 5
	add := 2
	expected := 7
	player := &Player{
		shares: map[interfaces.Corporation]int{
			corp: original,
		},
	}
	player.AddShares(corp, add)
	if player.Shares(corp) != expected {
		t.Errorf("AddShares() must add %d stock shares as owned by the player in corporation %p, got %d", add, corp, player.Shares(corp))
	}
}

func TestRemoveShares(t *testing.T) {
	corp := &mocks.Corporation{}
	original := 5
	remove := 2
	expected := 3
	player := &Player{
		shares: map[interfaces.Corporation]int{
			corp: original,
		},
	}
	player.RemoveShares(corp, remove)
	if player.Shares(corp) != expected {
		t.Errorf("RemoveShares() must remove %d stock shares as owned by the player in corporation %p, got %d", remove, corp, player.Shares(corp))
	}
}

func TestTiles(t *testing.T) {
	player := &Player{
		tiles: []interfaces.Tile{
			&mocks.Tile{},
			&mocks.Tile{},
		},
	}
	if len(player.Tiles()) != 2 {
		t.Errorf("Tiles() must return 2 tiles in player's hand, got %d", len(player.Tiles()))
	}
}

func TestHasTile(t *testing.T) {
	player := &Player{
		tiles: []interfaces.Tile{
			&mocks.Tile{FakeNumber: 5, FakeLetter: "C"},
			&mocks.Tile{FakeNumber: 4, FakeLetter: "A"},
		},
	}
	if player.HasTile(&mocks.Tile{FakeNumber: 5, FakeLetter: "C"}) != true {
		t.Errorf("HasTile() must return true for tile 5C")
	}
	if player.HasTile(&mocks.Tile{FakeNumber: 2, FakeLetter: "F"}) != false {
		t.Errorf("HasTile() must return false for tile 2F")
	}

}
