package acquire

import (
	"errors"

	"github.com/svera/acquire/interfaces"
)

// BuyStock buys stock from corporations
func (g *Game) BuyStock(buys map[interfaces.Corporation]int) error {
	if g.stateMachine.CurrentStateName() != interfaces.BuyStockStateName {
		return errors.New(ActionNotAllowed)
	}

	if err := g.checkBuy(buys); err != nil {
		return err
	}

	for corp, amount := range buys {
		g.buy(corp, amount)
	}

	return g.endTurn()
}

func (g *Game) buy(corp interfaces.Corporation, amount int) {
	corp.RemoveStock(amount)
	g.CurrentPlayer().
		AddShares(corp, amount).
		RemoveCash(corp.StockPrice() * amount)
}

func (g *Game) checkBuy(buys map[interfaces.Corporation]int) error {
	var totalStock, totalPrice int = 0, 0
	for corp, amount := range buys {
		if corp.Size() == 0 {
			return errors.New(StockSharesNotBuyable)
		}
		if amount > corp.Stock() {
			return errors.New(NotEnoughStockShares)
		}
		totalStock += amount
		totalPrice += corp.StockPrice() * amount
	}

	if totalStock > 3 {
		return errors.New(TooManyStockSharesToBuy)
	}

	if totalPrice > g.CurrentPlayer().Cash() {
		return errors.New(NotEnoughCash)
	}
	return nil
}
