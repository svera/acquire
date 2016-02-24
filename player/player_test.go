package player

import (
	"github.com/svera/acquire/interfaces"
	"testing"
)

func TestPickTile(t *testing.T) {
	player := New()
	tl := &interfaces.TileMock{FakeNumber: 2, FakeLetter: "C"}
	player.PickTile(tl)
	if len(player.tiles) != 1 {
		t.Errorf("Player must have exactly 1 tile, got %d", len(player.tiles))
	}
}

func TestUseTile(t *testing.T) {
	player := New()

	player.tiles = []interfaces.Tile{
		&interfaces.TileMock{FakeNumber: 7, FakeLetter: "c"},
		&interfaces.TileMock{FakeNumber: 5, FakeLetter: "A"},
		&interfaces.TileMock{FakeNumber: 8, FakeLetter: "E"},
		&interfaces.TileMock{FakeNumber: 3, FakeLetter: "D"},
		&interfaces.TileMock{FakeNumber: 1, FakeLetter: "B"},
		&interfaces.TileMock{FakeNumber: 4, FakeLetter: "I"},
	}

	tl := &interfaces.TileMock{FakeNumber: 5, FakeLetter: "A"}
	player.DiscardTile(tl)
	if len(player.tiles) != 5 {
		t.Errorf("Players must have 5 tiles after using one, got %d", len(player.tiles))
	}
	if tl.Number() != 5 || tl.Letter() != "A" {
		t.Errorf("DiscardTile() must return tile 5A")
	}
}

func TestShares(t *testing.T) {
	corp := &interfaces.CorporationMock{}
	expected := 5
	player := &Player{
		shares: map[interfaces.Corporation]int{
			corp: expected,
		},
	}
	if player.Shares(corp) != expected {
		t.Errorf("Shares() must return that the player has exactly %d stock shares in corporation %s, got %d", expected, corp.Name(), player.Shares(corp))
	}
}

func TestAddShares(t *testing.T) {
	corp := &interfaces.CorporationMock{}
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
		t.Errorf("AddShares() must add %d stock shares as owned by the player in corporation %s, got %d", add, corp.Name(), player.Shares(corp))
	}
}
