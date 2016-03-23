// Package bots implements different types of AI for playing Acquire games
package bots

import (
	"github.com/svera/acquire/interfaces"
	"strings"
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
	}
	return msg
}

func (r *Random) playTile() string {
	return r.status.Hand[0].Coords
}

func (r *Random) foundCorporation() string {
	for _, corp := range r.status.Corps {
		if corp.Size == 0 {
			return strings.ToLower(corp.Name)
		}
	}
	return ""
}

func (r *Random) buyStock() map[string]int {
	return map[string]int{
		"sackson": 0,
	}
}
