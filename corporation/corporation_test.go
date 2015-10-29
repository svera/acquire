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
