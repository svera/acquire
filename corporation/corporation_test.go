package corporation

import (
	"github.com/svera/acquire/tileset"
	"testing"
)

func TestStockPrice(t *testing.T) {
	var corporations = new([3]*Corporation)
	corporations[0], _ = New("class0", 0)
	corporations[1], _ = New("class1", 1)
	corporations[2], _ = New("class2", 2)

	corporations[0].tiles = []tileset.Position{
		{Number: 1, Letter: "A"},
		{Number: 2, Letter: "A"},
	}
	corporations[1].tiles = []tileset.Position{
		{Number: 1, Letter: "B"},
		{Number: 2, Letter: "B"},
	}
	corporations[2].tiles = []tileset.Position{
		{Number: 1, Letter: "C"},
		{Number: 2, Letter: "C"},
	}

	var expectedStockPrices = new([3]int)
	expectedStockPrices[0] = 200
	expectedStockPrices[1] = 300
	expectedStockPrices[2] = 400

	for class, corporation := range corporations {
		if corporation.StockPrice() != expectedStockPrices[class] {
			t.Errorf(
				"Class %d corporation with a size of 2 must have a stock price of %d, got %d",
				class,
				expectedStockPrices[class],
				corporation.StockPrice(),
			)
		}
	}
}

func TestSize(t *testing.T) {
	corp, _ := New("Test", 0)
	expectedSize := 8
	corp.tiles = make([]tileset.Position, expectedSize)
	if size := corp.Size(); size != expectedSize {
		t.Errorf("Expected a corporation size of %d, got %d", expectedSize, size)
	}
}

func TestAddTiles(t *testing.T) {
	corp, _ := New("Test", 0)
	expectedTiles := []tileset.Position{
		{Number: 1, Letter: "A"},
		{Number: 2, Letter: "A"},
	}
	corp.AddTiles(expectedTiles)
	if corp.tiles[0] != expectedTiles[0] || corp.tiles[1] != expectedTiles[1] {
		t.Errorf("Tiles not added to corporation")
	}
}

func TestAddTile(t *testing.T) {
	corp, _ := New("Test", 0)
	expectedTile := tileset.Position{Number: 1, Letter: "A"}

	corp.AddTile(expectedTile)
	if corp.tiles[0] != expectedTile {
		t.Errorf("Tile not added to corporation")
	}
}

func TestSetId(t *testing.T) {
	corp, _ := New("Test", 0)
	expectedId := 5
	corp.SetId(expectedId)
	if corp.id != expectedId {
		t.Errorf("Corporation ID not set")
	}
}

func TestId(t *testing.T) {
	corp, _ := New("Test", 0)
	expectedId := 5
	corp.id = expectedId
	if corp.Id() != expectedId {
		t.Errorf("Corporation ID not got")
	}
}

func TestStock(t *testing.T) {
	corp, _ := New("Test", 0)
	expectedStock := 20
	corp.stock = expectedStock
	if corp.Stock() != expectedStock {
		t.Errorf("Corporation stock not got")
	}
}

func TestSetStock(t *testing.T) {
	corp, _ := New("Test", 0)
	expectedStock := 20
	corp.SetStock(expectedStock)
	if corp.stock != expectedStock {
		t.Errorf("Corporation stock not set")
	}
}

func TestMajorityBonus(t *testing.T) {
	corp, _ := New("Test", 0)
	corp.tiles = make([]tileset.Position, 2)
	expectedMajorityBonus := 2000
	if bonus := corp.MajorityBonus(); bonus != expectedMajorityBonus {
		t.Errorf("Expected majority bonus of %d, got %d", expectedMajorityBonus, bonus)
	}
}
