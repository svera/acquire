package player

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tileset"
	"reflect"
	"testing"
)

func TestPickTile(t *testing.T) {
	player := New("Test")
	tl := tileset.Position{Number: 2, Letter: "C"}
	player.PickTile(tl)
	if len(player.tiles) != 1 {
		t.Errorf("Player must have exactly 1 tile, got %d", len(player.tiles))
	}

	player.tiles = []tileset.Position{
		{Number: 7, Letter: "C"},
		{Number: 5, Letter: "A"},
		{Number: 8, Letter: "E"},
		{Number: 3, Letter: "D"},
		{Number: 1, Letter: "B"},
		{Number: 4, Letter: "I"},
	}
	player.PickTile(tl)
	if len(player.tiles) > 6 {
		t.Errorf("Player can not have more than 6 tiles, got %d", len(player.tiles))
	}
}

func TestSort(t *testing.T) {
	players := []Interface{
		NewStub("Test1"),
		NewStub("Test2"),
		NewStub("Test3"),
		NewStub("Test4"),
	}

	corp, _ := corporation.New("Test", 0)

	players[0].(*Stub).SetShares(corp, 3)
	players[1].(*Stub).SetShares(corp, 1)
	players[2].(*Stub).SetShares(corp, 0)
	players[3].(*Stub).SetShares(corp, 2)

	shares := func(p1, p2 Interface) bool {
		return p1.Shares(corp) > p2.Shares(corp)
	}
	expectedSort := []Interface{
		players[0],
		players[3],
		players[1],
		players[2],
	}
	By(shares).Sort(players)
	if !reflect.DeepEqual(players, expectedSort) {
		t.Errorf("Players not sorted by corporation %s's shares amount", corp.Name())
	}
}

func TestUseTile(t *testing.T) {
	player := New("Test")

	player.tiles = []tileset.Position{
		{Number: 7, Letter: "C"},
		{Number: 5, Letter: "A"},
		{Number: 8, Letter: "E"},
		{Number: 3, Letter: "D"},
		{Number: 1, Letter: "B"},
		{Number: 4, Letter: "I"},
	}

	tile := tileset.Position{Number: 5, Letter: "A"}
	player.UseTile(tile)
	if len(player.tiles) != 5 {
		t.Errorf("Players must have 5 tiles after using one, got %d", len(player.tiles))
	}
	if tile.Number != 5 || tile.Letter != "A" {
		t.Errorf("UseTile() must return tile 5A")
	}
}
