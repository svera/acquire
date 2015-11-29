package player

import (
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
	"reflect"
	"testing"
)

func TestPickTile(t *testing.T) {
	player := New("Test")
	tl := tile.New(2, "C", tile.Orphan{})
	player.PickTile(tl)
	if len(player.tiles) != 1 {
		t.Errorf("Player must have exactly 1 tile, got %d", len(player.tiles))
	}
}

func TestSort(t *testing.T) {
	players := []ShareInterface{
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

	shares := func(p1, p2 ShareInterface) bool {
		return p1.Shares(corp) > p2.Shares(corp)
	}
	expectedSort := []ShareInterface{
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

	player.tiles = []tile.Interface{
		tile.New(7, "C", tile.Orphan{}),
		tile.New(5, "A", tile.Orphan{}),
		tile.New(8, "E", tile.Orphan{}),
		tile.New(3, "D", tile.Orphan{}),
		tile.New(1, "B", tile.Orphan{}),
		tile.New(4, "I", tile.Orphan{}),
	}

	tl := tile.New(5, "A", tile.Orphan{})
	player.DiscardTile(tl)
	if len(player.tiles) != 5 {
		t.Errorf("Players must have 5 tiles after using one, got %d", len(player.tiles))
	}
	if tl.Number() != 5 || tl.Letter() != "A" {
		t.Errorf("DiscardTile() must return tile 5A")
	}
}
