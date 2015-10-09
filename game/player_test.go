package game

import "testing"

func TestGetTile(t *testing.T) {
	player := NewPlayer("Test")
	tile := Tile{number: 2, letter: "C"}
	player.GetTile(tile)
	if len(player.tiles) != 1 {
		t.Errorf("Player must have exactly 1 tile, got %d", len(player.tiles))
	}

	tiles := []Tile{
		{number: 7, letter: "C"},
		{number: 5, letter: "A"},
		{number: 8, letter: "E"},
		{number: 3, letter: "D"},
		{number: 1, letter: "B"},
		{number: 4, letter: "I"},
	}
	player.tiles = tiles
	player.GetTile(tile)
	if len(player.tiles) > 6 {
		t.Errorf("Player can not have more than 6 tiles, got %d", len(player.tiles))
	}
}

func TestBuyStock(t *testing.T) {
	player := NewPlayer("Test")
	corporation, _ := NewCorporation("Test", 0)
	var buys []Buy
	var expectedAvailableStock uint = 23
	var expectedPlayerStock uint = 2
	buys = append(buys, Buy{corporation: corporation, amount: 2})
	player.BuyStocks(buys)

	if corporation.GetStock() != expectedAvailableStock {
		t.Errorf("Corporation stock shares has not decreased, must be %d, got %d", expectedAvailableStock, corporation.stock)
	}
	if player.shares[corporation.getId()] != expectedPlayerStock {
		t.Errorf("Player stock shares has not increased, must be %d, got %d", expectedPlayerStock, player.shares)
	}
}

func TestBuyStockWithNotEnoughCash(t *testing.T) {
	player := NewPlayer("Test")
	player.cash = 100
	corporation, _ := NewCorporation("Test", 0)
	corporation.size = 2
	var buys []Buy
	buys = append(buys, Buy{corporation: corporation, amount: 2})
	err := player.BuyStocks(buys)
	if err == nil {
		t.Errorf("Trying to buy stock shares withouth enough money must throw error")
	}
}
