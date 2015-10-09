package game

import (
	"errors"
)

type Buy struct {
	corporation *Corporation
	amount      uint
}

type Player struct {
	name   string
	cash   uint
	tiles  []Tile
	shares [7]uint
}

func NewPlayer(name string) *Player {
	return &Player{
		name:   name,
		cash:   6000,
		shares: [7]uint{},
	}
}

// Placeholder function, pending implementation
func (p *Player) BuyStocks(buys []Buy) error {
	err := p.checkBuy(buys)

	if err != nil {
		return err
	}

	for _, buy := range buys {
		buy.corporation.stock -= buy.amount
		p.shares[buy.corporation.getId()] = buy.amount
		p.cash -= buy.corporation.GetStockPrice() * buy.amount
	}
	return nil
}

func (p *Player) checkBuy(buys []Buy) error {
	var totalStock, totalPrice uint = 0, 0
	for _, buy := range buys {
		if buy.amount > buy.corporation.stock {
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
func (p *Player) GetTile(tile Tile) error {
	if len(p.tiles) >= 6 {
		return errors.New("Player cannot have more than 6 tiles")
	}
	p.tiles = append(p.tiles, tile)
	return nil
}
