package player

import (
	"errors"
	"github.com/svera/acquire/game/corporation"
	"github.com/svera/acquire/game/tileset"
)

type Buy struct {
	corporation *corporation.Corporation
	amount      uint
}

type Player struct {
	name   string
	cash   uint
	tiles  []tileset.Position
	shares [7]uint
	Shares func(c *corporation.Corporation) uint
}

func New(name string) *Player {
	player := &Player{
		name:   name,
		cash:   6000,
		shares: [7]uint{},
	}
	player.Shares = makeSharesFunc(player)
	return player
}

// In order to make testing easier, we implement Shares() with a closure
// that can be later overwritten by tests
func makeSharesFunc(p *Player) func(c *corporation.Corporation) uint {
	return func(c *corporation.Corporation) uint {
		return p.shares[c.Id()]
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
		p.cash -= buy.corporation.StockPrice() * buy.amount
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
		totalPrice += buy.corporation.StockPrice() * buy.amount
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
func (p *Player) PickTile(t tileset.Position) error {
	if len(p.tiles) >= 6 {
		return errors.New("Player cannot have more than 6 tiles")
	}
	p.tiles = append(p.tiles, t)
	return nil
}

func (p *Player) Tiles() []tileset.Position {
	return p.tiles
}

func (p *Player) ReceiveBonus(amount uint) {
	p.cash += amount
}

func (p *Player) UseTile(t tileset.Position) tileset.Position {
	for i, currentTile := range p.tiles {
		if currentTile.Number == t.Number && currentTile.Letter == t.Letter {
			p.tiles = append(p.tiles[:i], p.tiles[i+1:]...)
			return currentTile
		}
	}
	return tileset.Position{}
}
