package corporation

import (
	"github.com/svera/acquire/tileset"
	"testing"
)

func TestStockPrice(t *testing.T) {
	var corporations = new([4]*Corporation)
	corporations[0], _ = New("class0", 0, 0)
	corporations[1], _ = New("class1", 1, 1)
	corporations[2], _ = New("class2", 2, 2)
	corporations[3], _ = New("class0 big", 0, 3)

	corporations[0].tiles = make([]tileset.Position, 2)
	corporations[1].tiles = make([]tileset.Position, 2)
	corporations[2].tiles = make([]tileset.Position, 2)
	corporations[3].tiles = make([]tileset.Position, 42)

	var expectedStockPrices = new([4]int)
	expectedStockPrices[0] = 200
	expectedStockPrices[1] = 300
	expectedStockPrices[2] = 400
	expectedStockPrices[3] = 1000

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
	corp, _ := New("Test", 0, 0)
	expectedSize := 8
	corp.tiles = make([]tileset.Position, expectedSize)
	if size := corp.Size(); size != expectedSize {
		t.Errorf("Expected a corporation size of %d, got %d", expectedSize, size)
	}
}

func TestAddTiles(t *testing.T) {
	corp, _ := New("Test", 0, 0)
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
	corp, _ := New("Test", 0, 0)
	expectedTile := tileset.Position{Number: 1, Letter: "A"}

	corp.AddTile(expectedTile)
	if corp.tiles[0] != expectedTile {
		t.Errorf("Tile not added to corporation")
	}
}

func TestId(t *testing.T) {
	corp, _ := New("Test", 0, 0)
	expectedId := 5
	corp.id = expectedId
	if corp.Id() != expectedId {
		t.Errorf("Corporation ID not got")
	}
}

func TestStock(t *testing.T) {
	corp, _ := New("Test", 0, 0)
	expectedStock := 20
	corp.stock = expectedStock
	if corp.Stock() != expectedStock {
		t.Errorf("Corporation stock not got")
	}
}

func TestSetStock(t *testing.T) {
	corp, _ := New("Test", 0, 0)
	expectedStock := 20
	corp.SetStock(expectedStock)
	if corp.stock != expectedStock {
		t.Errorf("Corporation stock not set")
	}
}

func TestMajorityBonus(t *testing.T) {
	corp, _ := New("Test", 0, 0)
	corp.tiles = make([]tileset.Position, 2)
	expectedMajorityBonus := 2000
	if bonus := corp.MajorityBonus(); bonus != expectedMajorityBonus {
		t.Errorf("Expected majority bonus of %d, got %d", expectedMajorityBonus, bonus)
	}

	corp.tiles = make([]tileset.Position, 42)
	expectedMajorityBonus = 10000
	if bonus := corp.MajorityBonus(); bonus != expectedMajorityBonus {
		t.Errorf("Expected majority bonus of %d, got %d", expectedMajorityBonus, bonus)
	}
}

func TestMinorityBonus(t *testing.T) {
	corp, _ := New("Test", 0, 0)
	corp.tiles = make([]tileset.Position, 2)
	expectedMinorityBonus := 1000
	if bonus := corp.MinorityBonus(); bonus != expectedMinorityBonus {
		t.Errorf("Expected minority bonus of %d, got %d", expectedMinorityBonus, bonus)
	}

	corp.tiles = make([]tileset.Position, 42)
	expectedMinorityBonus = 5000
	if bonus := corp.MinorityBonus(); bonus != expectedMinorityBonus {
		t.Errorf("Expected minority bonus of %d, got %d", expectedMinorityBonus, bonus)
	}
}

func TestIsSafe(t *testing.T) {
	corp, _ := New("Test", 0, 0)
	corp.tiles = make([]tileset.Position, 2)
	if corp.IsSafe() {
		t.Errorf("Unsafe corporation regarded as safe")
	}
	corp.tiles = make([]tileset.Position, 11)
	if !corp.IsSafe() {
		t.Errorf("Safe corporation regarded as unsafe")
	}
}

func TestIsActive(t *testing.T) {
	corp, _ := New("Test", 0, 0)
	if corp.IsActive() {
		t.Errorf("Inactive corporation regarded as active")
	}
	corp.tiles = make([]tileset.Position, 2)
	if !corp.IsActive() {
		t.Errorf("Active corporation regarded as inactive")
	}
}

func TestName(t *testing.T) {
	corp, _ := New("Test", 0, 0)
	if corp.Name() != "Test" {
		t.Errorf("Expected corporation name 'Test'")
	}
}
