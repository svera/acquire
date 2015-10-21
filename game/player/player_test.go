package player

import (
	"github.com/svera/acquire/game/corporation"
	"github.com/svera/acquire/game/tileset"
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

func TestBuyStock(t *testing.T) {
	player := New("Test")
	corporation, _ := corporation.New("Test", 0)
	corporation.AddTiles(
		[]tileset.Position{
			{Number: 1, Letter: "A"},
			{Number: 2, Letter: "A"},
		},
	)
	var buys []Buy
	var expectedAvailableStock uint = 23
	var expectedPlayerStock uint = 2
	buys = append(buys, Buy{corporation: corporation, amount: 2})
	player.BuyStocks(buys)

	if corporation.Stock() != expectedAvailableStock {
		t.Errorf("Corporation stock shares have not decreased, must be %d, got %d", expectedAvailableStock, corporation.Stock())
	}
	if player.shares[corporation.Id()] != expectedPlayerStock {
		t.Errorf("Player stock shares have not increased, must be %d, got %d", expectedPlayerStock, player.shares[corporation.Id()])
	}
}

func TestBuyStockWithNotEnoughCash(t *testing.T) {
	player := New("Test")
	player.cash = 100
	corporation, _ := corporation.New("Test", 0)
	corporation.AddTiles(
		[]tileset.Position{
			{Number: 1, Letter: "A"},
			{Number: 2, Letter: "A"},
		},
	)
	var buys []Buy
	buys = append(buys, Buy{corporation: corporation, amount: 2})
	err := player.BuyStocks(buys)
	if err == nil {
		t.Errorf("Trying to buy stock shares without enough money must throw error")
	}
}

func TestSort(t *testing.T) {
	players := []*Player{
		New("Test1"),
		New("Test2"),
		New("Test3"),
		New("Test4"),
	}

	corporation, _ := corporation.New("Test", 0)

	players[0].shares[0] = 3
	players[1].shares[0] = 1
	players[2].shares[0] = 0
	players[3].shares[0] = 2

	shares := func(p1, p2 *Player) bool {
		return p1.Shares(corporation) > p2.Shares(corporation)
	}
	expectedSort := []*Player{
		players[0],
		players[3],
		players[1],
		players[2],
	}
	By(shares).Sort(players)
	if !reflect.DeepEqual(players, expectedSort) {
		t.Errorf("Players not sorted by corporation %s's shares amount", corporation.Name())
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

	tile := player.UseTile(tileset.Position{Number: 5, Letter: "A"})
	if len(player.tiles) != 5 {
		t.Errorf("Players must have 5 tiles after using one, got %d", len(player.tiles))
	}
	if tile.Number != 5 || tile.Letter != "A" {
		t.Errorf("UseTile() must return tile 5A")
	}
}
