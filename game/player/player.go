package player

import (
	"errors"
	c "github.com/svera/acquire/game/corporation"
	"github.com/svera/acquire/game/tileset"
)

type Buy struct {
	corporation *c.Corporation
	amount      uint
}

type Player struct {
	name   string
	cash   uint
	tiles  []tileset.Position
	shares [7]uint
}

func New(name string) *Player {
	return &Player{
		name:   name,
		cash:   6000,
		shares: [7]uint{},
	}
}

// Buys stock from corporations
func (p *Player) BuyStocks(buys []Buy) error {
	err := p.checkBuy(buys)

	if err != nil {
		return err
	}

	for _, buy := range buys {
		buy.corporation.SetStock(buy.corporation.Stock() - buy.amount)
		p.shares[buy.corporation.Id()] = buy.amount
		p.cash -= buy.corporation.GetStockPrice() * buy.amount
	}
	return nil
}

func (p *Player) checkBuy(buys []Buy) error {
	var totalStock, totalPrice uint = 0, 0
	for _, buy := range buys {
		if buy.corporation.Size() == 0 {
			return errors.New("Player cannot buy shares of a corporation not on board")
		}
		if buy.amount > buy.corporation.Stock() {
			return errors.New("Player cannot buy more shares than the available stock")
		}
		totalStock += buy.amount
		totalPrice += buy.corporation.GetStockPrice() * buy.amount
	}
	if totalStock > 3 {
		return errors.New("Player cannot buy more than 3 stock shares per turn")
	}

	if totalPrice > p.cash {
		return errors.New("Player doesn't have enough cash to buy those stock shares")
	}
	return nil
}

// Adds a new tile to the players' tileset
func (p *Player) GetTile(t tileset.Position) error {
	if len(p.tiles) >= 6 {
		return errors.New("Player cannot have more than 6 tiles")
	}
	p.tiles = append(p.tiles, t)
	return nil
}

func (p *Player) Tiles() []tileset.Position {
	return p.tiles
}

func (p *Player) Shares(corporationId uint) uint {
	return p.shares[corporationId]
}
