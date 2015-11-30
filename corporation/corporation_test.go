package corporation

import (
	"testing"
)

func TestStockPrice(t *testing.T) {
	var corporations = new([4]*Corporation)
	corporations[0], _ = New("class0", 0)
	corporations[1], _ = New("class1", 1)
	corporations[2], _ = New("class2", 2)
	corporations[3], _ = New("class0 big", 0)

	corporations[0].size = 2
	corporations[1].size = 2
	corporations[2].size = 2
	corporations[3].size = 42

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
	corp, _ := New("Test", 0)
	expectedSize := 8
	corp.size = expectedSize
	if size := corp.Size(); size != expectedSize {
		t.Errorf("Expected a corporation size of %d, got %d", expectedSize, size)
	}
}

func TestGrow(t *testing.T) {
	corp, _ := New("Test", 0)
	expectedSize := 2
	corp.Grow(2)
	if corp.size != expectedSize {
		t.Errorf("Tiles not added to corporation")
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

func TestAddStock(t *testing.T) {
	corp, _ := New("Test", 0)
	expectedStock := 45
	corp.AddStock(20)
	if corp.stock != expectedStock {
		t.Errorf("Corporation stock not added, expected %d, got %d", expectedStock, corp.stock)
	}
}

func TestRemoveStock(t *testing.T) {
	corp, _ := New("Test", 0)
	expectedStock := 5
	corp.RemoveStock(20)
	if corp.stock != expectedStock {
		t.Errorf("Corporation stock not removed, expected %d, got %d", expectedStock, corp.stock)
	}
}

func TestMajorityBonus(t *testing.T) {
	corp, _ := New("Test", 0)
	corp.size = 2
	expectedMajorityBonus := 2000
	if bonus := corp.MajorityBonus(); bonus != expectedMajorityBonus {
		t.Errorf("Expected majority bonus of %d, got %d", expectedMajorityBonus, bonus)
	}

	corp.size = 42
	expectedMajorityBonus = 10000
	if bonus := corp.MajorityBonus(); bonus != expectedMajorityBonus {
		t.Errorf("Expected majority bonus of %d, got %d", expectedMajorityBonus, bonus)
	}
}

func TestMinorityBonus(t *testing.T) {
	corp, _ := New("Test", 0)
	corp.size = 2
	expectedMinorityBonus := 1000
	if bonus := corp.MinorityBonus(); bonus != expectedMinorityBonus {
		t.Errorf("Expected minority bonus of %d, got %d", expectedMinorityBonus, bonus)
	}

	corp.size = 42
	expectedMinorityBonus = 5000
	if bonus := corp.MinorityBonus(); bonus != expectedMinorityBonus {
		t.Errorf("Expected minority bonus of %d, got %d", expectedMinorityBonus, bonus)
	}
}

func TestIsSafe(t *testing.T) {
	corp, _ := New("Test", 0)
	corp.size = 2
	if corp.IsSafe() {
		t.Errorf("Unsafe corporation regarded as safe")
	}
	corp.size = 11
	if !corp.IsSafe() {
		t.Errorf("Safe corporation regarded as unsafe")
	}
}

func TestIsActive(t *testing.T) {
	corp, _ := New("Test", 0)
	if corp.IsActive() {
		t.Errorf("Inactive corporation regarded as active")
	}
	corp.size = 2
	if !corp.IsActive() {
		t.Errorf("Active corporation regarded as inactive")
	}
}

func TestName(t *testing.T) {
	corp, _ := New("Test", 0)
	if corp.Name() != "Test" {
		t.Errorf("Expected corporation name 'Test'")
	}
}
