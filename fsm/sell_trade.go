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

// ToSellTrade returns self because machine is already in that state
// Although it may make no sense, this is to avoid the machine to go to errorState
func (s *sellTrade) ToSellTrade() interfaces.State {
	return s
}
