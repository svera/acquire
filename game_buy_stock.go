package acquire

import (
	"errors"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/tile"
)

// BuyStock buys stock from corporations
func (g *Game) BuyStock(buys map[int]int) error {
	if g.state.Name() != "BuyStock" {
		return errors.New(ActionNotAllowed)
	}

	if err := g.checkBuy(buys); err != nil {
		return err
	}

	for corporationID, amount := range buys {
		corp := g.corporations[corporationID]
		g.buy(corp, amount)
	}

	if g.endGameClaimed {
		g.state = g.state.ToEndGame()
		return g.finish()
	} else {
		return g.drawTile()
	}
}

func (g *Game) buy(corp corporation.Interface, amount int) {
	corp.RemoveStock(amount)
	g.CurrentPlayer().
		AddShares(corp, amount).
		RemoveCash(corp.StockPrice() * amount)
}

func (g *Game) checkBuy(buys map[int]int) error {
	var totalStock, totalPrice int = 0, 0
	for corporationID, amount := range buys {
		corp := g.corporations[corporationID]
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

// A player takes a tile from the facedown cluster to replace
// the one he/she played. This is not done until the end of
// the turn.
func (g *Game) drawTile() error {
	var tile tile.Interface
	var err error
	if tile, err = g.tileset.Draw(); err != nil {
		return err
	}
	g.CurrentPlayer().PickTile(tile)

	if err = g.replaceUnplayableTiles(); err != nil {
		return err
	}

	g.state = g.state.ToPlayTile()
	g.nextPlayer()
	return nil
}

// if a player has any permanently
// unplayable tiles that player discard the unplayable tiles
// and draws an equal number of replacement tiles. This can
// only be done once per turn.
func (g *Game) replaceUnplayableTiles() error {
	for _, tile := range g.CurrentPlayer().Tiles() {
		if g.isTileUnplayable(tile) {

			g.CurrentPlayer().DiscardTile(tile)
			if newTile, err := g.tileset.Draw(); err == nil {
				g.CurrentPlayer().PickTile(newTile)
			} else {
				return err
			}
		}
	}

	return nil
}
