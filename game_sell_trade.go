package game

import (
	"errors"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/player"
)

// TODO
func (g *Game) SellTrade(pl player.Interface, sell map[corporation.Interface]int, trade map[corporation.Interface]int) error {
	if err := g.checkSellTrade(pl, sell, trade); err != nil {
		return err
	}
	for corp, amount := range sell {
		g.sell(pl, corp, amount)
	}

	for corp, amount := range trade {
		g.trade(pl, corp, amount)
	}
	return nil
}

func (g *Game) sell(pl player.Interface, corp corporation.Interface, amount int) {
	corp.AddStock(amount)
	pl.
		RemoveShares(corp, amount).
		AddCash(corp.StockPrice() * amount)
}

func (g *Game) trade(pl player.Interface, corp corporation.Interface, amount int) {
	acquirer := g.mergeCorps["acquirer"][0]
	amountSharesAcquiringCorp := amount / 2
	corp.AddStock(amount)
	acquirer.RemoveStock(amountSharesAcquiringCorp)
	pl.
		RemoveShares(corp, amount).
		AddShares(acquirer, amountSharesAcquiringCorp)
}

func (g *Game) checkSellTrade(pl player.Interface, sell map[corporation.Interface]int, trade map[corporation.Interface]int) error {
	if g.state.Name() != "SellTrade" {
		return errors.New(ActionNotAllowed)
	}
	for corp, amount := range sell {
		if amount > 0 && pl.Shares(corp) == 0 {
			return errors.New(NoCorporationSharesOwned)
		}
		if pl.Shares(corp) < amount {
			return errors.New(NotEnoughCorporationSharesOwned)
		}
	}
	for corp, amount := range trade {
		if amount > 0 && pl.Shares(corp) == 0 {
			return errors.New(NoCorporationSharesOwned)
		}
		if corp.Stock() < (amount / 2) {
			return errors.New(NotEnoughStockShares)
		}
		if pl.Shares(corp) < amount {
			return errors.New(NotEnoughCorporationSharesOwned)
		}
	}
	return nil
}
