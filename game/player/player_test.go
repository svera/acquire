package player

import (
	"github.com/svera/acquire/game/corporation"
	"github.com/svera/acquire/game/tileset"
	"testing"
)

func TestGetTile(t *testing.T) {
	player := New("Test")
	tl := tileset.Position{Number: 2, Letter: "C"}
	player.GetTile(tl)
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
	player.GetTile(tl)
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
