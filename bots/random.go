// Package bots implements different types of AI for playing Acquire games
package bots

import (
	"github.com/svera/acquire/interfaces"
	"math/rand"
	"strings"
	"time"
)

// Random is a struct which implements a very stupid AI, which basically
// chooses all its decisions randomly (So not that much an AI but an AS)
type Random struct {
	*base
}

func NewRandom() *Random {
	return &Random{
		&base{},
	}
}

func (r *Random) Play() interface{} {
	var msg Message
	switch r.status.State {
	case interfaces.PlayTileStateName:
		msg = Message{
			Type:   PlayTileResponseType,
			Params: r.playTile(),
		}
	case interfaces.FoundCorpStateName:
		msg = Message{
			Type:   NewCorpResponseType,
			Params: r.foundCorporation(),
		}
	case interfaces.BuyStockStateName:
		msg = Message{
			Type:   BuyResponseType,
			Params: r.buyStock(),
		}
	case interfaces.SellTradeStateName:
		msg = Message{
			Type:   SellTradeResponseType,
			Params: r.sellTrade(),
		}
	case interfaces.UntieMergeStateName:
		msg = Message{
			Type:   UntieMergeResponseType,
			Params: r.untieMerge(),
		}
	}
	return msg
}

func (r *Random) playTile() PlayTileResponseParams {
	source := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(source)
	tileNumber := rn.Intn(len(r.status.Hand))

	return PlayTileResponseParams{
		Tile: r.status.Hand[tileNumber].Coords,
	}
}

func (r *Random) foundCorporation() NewCorpResponseParams {
	response := NewCorpResponseParams{}
	for _, corp := range r.status.Corps {
		if corp.Size == 0 {
			response.Corporation = strings.ToLower(corp.Name)
		}
	}
	return response
}

// buyStock buys stock from a random active corporation
func (r *Random) buyStock() BuyResponseParams {
	source := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(source)
	buy := 0
	var corpNumber int
	var corp CorpData

	for {
		corpNumber = rn.Intn(len(r.status.Corps))
		corp = r.status.Corps[corpNumber]
		if corp.Size > 0 {
			break
		}
	}
	if corp.RemainingShares > 3 && corp.Size > 0 && r.hasEnoughCash(3, corp.Price) {
		buy = 3
	} else if corp.Size > 0 && r.hasEnoughCash(corp.RemainingShares, corp.Price) {
		buy = corp.RemainingShares
	}
	return BuyResponseParams{
		Corporations: map[string]int{
			strings.ToLower(corp.Name): buy,
		},
	}
}

func (r *Random) hasEnoughCash(amount int, price int) bool {
	return amount*price < r.status.PlayerInfo.Cash
}

func (r *Random) sellTrade() SellTradeResponseParams {
	var sellTrade SellTradeResponseParams
	sellTradeCorporations := map[string]SellTrade{}

	for i, corp := range r.status.Corps {
		if corp.Defunct && r.status.PlayerInfo.OwnedShares[i] > 0 {
			sellTradeCorporations[strings.ToLower(corp.Name)] = SellTrade{
				Sell: r.status.PlayerInfo.OwnedShares[i],
			}
		}
	}
	sellTrade.Corporations = sellTradeCorporations
	return sellTrade
}

func (r *Random) untieMerge() UntieMergeResponseParams {
	return UntieMergeResponseParams{
		Corporation: r.status.TiedCorps[0],
	}
}
