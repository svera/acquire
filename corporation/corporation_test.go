package corporation

import (
	"testing"

	"github.com/svera/acquire/interfaces"
)

func TestSize(t *testing.T) {
	corp := New()
	expectedSize := 8
	corp.size = expectedSize
	if size := corp.Size(); size != expectedSize {
		t.Errorf("Expected a corporation size of %d, got %d", expectedSize, size)
	}
}

func TestGrow(t *testing.T) {
	corp := New()
	expectedSize := 2
	corp.Grow(2)
	if corp.size != expectedSize {
		t.Errorf("Tiles not added to corporation")
	}
}

func TestStock(t *testing.T) {
	corp := New()
	expectedStock := 20
	corp.stock = expectedStock
	if corp.Stock() != expectedStock {
		t.Errorf("Corporation stock not got")
	}
}

func TestAddStock(t *testing.T) {
	corp := New()
	expectedStock := 45
	corp.AddStock(20)
	if corp.stock != expectedStock {
		t.Errorf("Corporation stock not added, expected %d, got %d", expectedStock, corp.stock)
	}
}

func TestRemoveStock(t *testing.T) {
	corp := New()
	expectedStock := 5
	corp.RemoveStock(20)
	if corp.stock != expectedStock {
		t.Errorf("Corporation stock not removed, expected %d, got %d", expectedStock, corp.stock)
	}
}

func TestMajorityBonus(t *testing.T) {
	corp := New()
	prices := make(map[int]interfaces.Prices)
	prices[2] = interfaces.Prices{Price: 200, MajorityBonus: 2000, MinorityBonus: 1000}
	prices[41] = interfaces.Prices{Price: 1000, MajorityBonus: 10000, MinorityBonus: 5000}
	corp.SetPricesChart(prices)
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

func TestStockPrice(t *testing.T) {
	corp := New()
	prices := make(map[int]interfaces.Prices)
	prices[2] = interfaces.Prices{Price: 200, MajorityBonus: 2000, MinorityBonus: 1000}
	prices[41] = interfaces.Prices{Price: 1000, MajorityBonus: 10000, MinorityBonus: 5000}
	corp.SetPricesChart(prices)
	corp.size = 2
	expectedStockPrice := 200
	if stockPrice := corp.StockPrice(); stockPrice != expectedStockPrice {
		t.Errorf("Expected stock price of %d, got %d", expectedStockPrice, stockPrice)
	}

	corp.size = 42
	expectedStockPrice = 1000
	if stockPrice := corp.StockPrice(); stockPrice != expectedStockPrice {
		t.Errorf("Expected stock price of %d, got %d", expectedStockPrice, stockPrice)
	}
}

func TestMinorityBonus(t *testing.T) {
	corp := New()
	prices := make(map[int]interfaces.Prices)
	prices[2] = interfaces.Prices{Price: 200, MajorityBonus: 2000, MinorityBonus: 1000}
	prices[41] = interfaces.Prices{Price: 1000, MajorityBonus: 10000, MinorityBonus: 5000}
	corp.SetPricesChart(prices)
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
	corp := New()
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
	corp := New()
	if corp.IsActive() {
		t.Errorf("Inactive corporation regarded as active")
	}
	corp.size = 2
	if !corp.IsActive() {
		t.Errorf("Active corporation regarded as inactive")
	}
}
