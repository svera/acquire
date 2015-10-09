package game

import (
	"testing"
)

func TestGetStockPrice(t *testing.T) {
	var corporations = new([3]*Corporation)
	corporations[0], _ = NewCorporation("class0", 0)
	corporations[1], _ = NewCorporation("class1", 1)
	corporations[2], _ = NewCorporation("class2", 2)

	corporations[0].size = 8
	corporations[1].size = 8
	corporations[2].size = 8

	var expectedStockPrices = new([3]uint)
	expectedStockPrices[0] = 600
	expectedStockPrices[1] = 700
	expectedStockPrices[2] = 800

	for class, corporation := range corporations {
		if corporation.GetStockPrice() != expectedStockPrices[class] {
			t.Errorf(
				"Class %d corporation with a size of 8 must have a stock price of %d, got %d",
				class,
				expectedStockPrices[class],
				corporation.GetStockPrice(),
			)
		}
	}
}
