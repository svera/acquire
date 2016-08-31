// Package bots implements different types of AI for playing Acquire games
package bots

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/svera/acquire/interfaces"
)

const (
	endGameCorporationSize = 41
	safeCorporationSize    = 11
)

// Chaotic is a struct which implements a very stupid AI, which basically
// chooses all its decisions randomly (So not that much an AI but an AS)
type Chaotic struct {
	*base
}

// NewChaotic returns a new instance of the chaotic AI bot
func NewChaotic() *Chaotic {
	return &Chaotic{
		&base{},
	}
}

// Play analyses the current game status and returns a message with the
// next play movement by the bot AI
func (r *Chaotic) Play() interface{} {
	var msg Message

	if !r.status.IsLastRound && r.claimEndGame() {
		msg = Message{
			Type: EndGameResponseType,
		}
	} else {
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
	}

	return msg
}

func (r *Chaotic) playTile() PlayTileResponseParams {
	source := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(source)
	tileCoords := r.tileCoords()
	tileNumber := rn.Intn(len(tileCoords))

	return PlayTileResponseParams{
		Tile: tileCoords[tileNumber],
	}
}

// As the tiles in hand come as a map, we need to store its coordinates in an array
// before selecting a random one (only the playable ones)
func (r *Chaotic) tileCoords() []string {
	coords := make([]string, 0, len(r.status.Hand))
	for k, playable := range r.status.Hand {
		if playable {
			coords = append(coords, k)
		}
	}
	return coords
}

func (r *Chaotic) foundCorporation() NewCorpResponseParams {
	source := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(source)
	var corpNumber int
	response := NewCorpResponseParams{}
	for {
		corpNumber = rn.Intn(len(r.status.Corps))
		if r.status.Corps[corpNumber].Size == 0 {
			response.CorporationIndex = corpNumber
			break
		}
	}
	return response
}

// buyStock buys stock from a random active corporation
func (r *Chaotic) buyStock() BuyResponseParams {
	source := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(source)
	buy := 0
	var corpIndex int
	var corp CorpData

	for {
		corpIndex = rn.Intn(len(r.status.Corps))
		corp = r.status.Corps[corpIndex]
		if corp.Size > 0 {
			break
		}
	}
	if corp.RemainingShares > 3 && corp.Size > 0 && r.hasEnoughCash(3, corp.Price) {
		buy = 3
	} else if corp.Size > 0 && r.hasEnoughCash(corp.RemainingShares, corp.Price) {
		buy = corp.RemainingShares
	}
	index := strconv.Itoa(corpIndex)
	return BuyResponseParams{
		CorporationsIndexes: map[string]int{
			index: buy,
		},
	}
}

func (r *Chaotic) hasEnoughCash(amount int, price int) bool {
	return amount*price < r.status.PlayerInfo.Cash
}

func (r *Chaotic) sellTrade() SellTradeResponseParams {
	var sellTrade SellTradeResponseParams
	sellTradeCorporations := map[string]SellTrade{}

	for i, corp := range r.status.Corps {
		if corp.Defunct && r.status.PlayerInfo.OwnedShares[i] > 0 {
			index := strconv.Itoa(i)
			sellTradeCorporations[index] = SellTrade{
				Sell: r.status.PlayerInfo.OwnedShares[i],
			}
		}
	}
	sellTrade.CorporationsIndexes = sellTradeCorporations
	return sellTrade
}

func (r *Chaotic) untieMerge() UntieMergeResponseParams {
	var untieMerge UntieMergeResponseParams
	for i, corp := range r.status.Corps {
		if corp.Tied {
			untieMerge = UntieMergeResponseParams{
				CorporationIndex: i,
			}
			break
		}
	}
	return untieMerge
}

func (r *Chaotic) claimEndGame() bool {
	var active, safe int
	for _, corp := range r.status.Corps {
		if corp.Size >= endGameCorporationSize {
			return true
		}
		if corp.Size > 0 {
			active++
		}
		if corp.Size >= safeCorporationSize {
			safe++
		}
	}
	if active > 0 && active == safe {
		return true
	}
	return false
}
