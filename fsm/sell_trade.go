package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// SellTrade is a struct representing a finite state machine's state
type sellTrade struct {
	errorState
}

// Name returns state's name
func (s *sellTrade) Name() string {
	return interfaces.SellTradeStateName
}

// ToBuyStock returns a BuyStock instance because it's an allowed state transition
func (s *sellTrade) ToBuyStock() interfaces.State {
	return &buyStock{}
}

// ToInsufficientPlayers returns an InsufficientPlayers instance because it's an allowed state transition
func (s *sellTrade) ToInsufficientPlayers() interfaces.State {
	return &insufficientPlayers{}
}
