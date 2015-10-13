package corporation

import (
	"testing"
)

func TestGetStockPrice(t *testing.T) {
	var corporations = new([3]*Corporation)
	corporations[0], _ = New("class0", 0)
	corporations[1], _ = New("class1", 1)
	corporations[2], _ = New("class2", 2)

	corporations[0].Size = func() uint {
		return 2
	}
	corporations[1].Size = func() uint {
		return 2
	}
	corporations[2].Size = func() uint {
		return 2
	}

	var expectedStockPrices = new([3]uint)
	expectedStockPrices[0] = 200
	expectedStockPrices[1] = 300
	expectedStockPrices[2] = 400

	for class, corporation := range corporations {
		if corporation.GetStockPrice() != expectedStockPrices[class] {
			t.Errorf(
				"Class %d corporation with a size of 2 must have a stock price of %d, got %d",
				class,
				expectedStockPrices[class],
				corporation.GetStockPrice(),
			)
		}
	}
}
