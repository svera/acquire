package game

import (
	"errors"
	"github.com/svera/acquire/corporation"
)

func (g *Game) SellTrade(sell map[corporation.Interface]int, trade map[corporation.Interface]int) error {
	if err := g.checkSellTrade(sell, trade); err != nil {
		return err
	}
	for corp, amount := range sell {
		g.sell(corp, amount)
	}

	for corp, amount := range trade {
		g.trade(corp, amount)
	}

	if len(g.sellTradePlayers) == 0 {
		g.setCurrentPlayer(g.frozenPlayer)
		g.state = g.state.ToBuyStock()
	} else {
		g.setCurrentPlayer(g.nextSellTradePlayer())
	}
	return nil
}

func (g *Game) nextSellTradePlayer() int {
	pl := g.sellTradePlayers[len(g.sellTradePlayers)-1]
	g.sellTradePlayers = g.sellTradePlayers[:len(g.sellTradePlayers)-1]
	return pl
}

func (g *Game) sell(corp corporation.Interface, amount int) {
	corp.AddStock(amount)
	g.CurrentPlayer().
		RemoveShares(corp, amount).
		AddCash(corp.StockPrice() * amount)
}

func (g *Game) trade(corp corporation.Interface, amount int) {
	acquirer := g.mergeCorps["acquirer"][0]
	amountSharesAcquiringCorp := amount / 2
	corp.AddStock(amount)
	acquirer.RemoveStock(amountSharesAcquiringCorp)
	g.CurrentPlayer().
		RemoveShares(corp, amount).
		AddShares(acquirer, amountSharesAcquiringCorp)
}

func (g *Game) checkSellTrade(sell map[corporation.Interface]int, trade map[corporation.Interface]int) error {
	if g.state.Name() != "SellTrade" {
		return errors.New(ActionNotAllowed)
	}
	for corp, amount := range sell {
		if amount > 0 && g.CurrentPlayer().Shares(corp) == 0 {
			return errors.New(NoCorporationSharesOwned)
		}
		if g.CurrentPlayer().Shares(corp) < amount {
			return errors.New(NotEnoughCorporationSharesOwned)
		}
	}
	for corp, amount := range trade {
		if amount > 0 && g.CurrentPlayer().Shares(corp) == 0 {
			return errors.New(NoCorporationSharesOwned)
		}
		if corp.Stock() < (amount / 2) {
			return errors.New(NotEnoughStockShares)
		}
		if g.CurrentPlayer().Shares(corp) < amount {
			return errors.New(NotEnoughCorporationSharesOwned)
		}
	}
	return nil
}
