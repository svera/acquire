package player

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/interfaces"
	"github.com/svera/acquire/tile"
	"testing"
)

func TestPickTile(t *testing.T) {
	player := New()
	tl := tile.New(2, "C")
	player.PickTile(tl)
	if len(player.tiles) != 1 {
		t.Errorf("Player must have exactly 1 tile, got %d", len(player.tiles))
	}
}

func TestUseTile(t *testing.T) {
	player := New()

	player.tiles = []interfaces.Tile{
		tile.New(7, "C"),
		tile.New(5, "A"),
		tile.New(8, "E"),
		tile.New(3, "D"),
		tile.New(1, "B"),
		tile.New(4, "I"),
	}

	tl := tile.New(5, "A")
	player.DiscardTile(tl)
	if len(player.tiles) != 5 {
		t.Errorf("Players must have 5 tiles after using one, got %d", len(player.tiles))
	}
	if tl.Number() != 5 || tl.Letter() != "A" {
		t.Errorf("DiscardTile() must return tile 5A")
	}
}

func TestShares(t *testing.T) {
	corp, _ := corporation.New("Test corp", 0)
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
	corp, _ := corporation.New("Test corp", 0)
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
