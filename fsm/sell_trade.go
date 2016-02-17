package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// SellTrade is a struct representing a finite state machine's state
type SellTrade struct {
	ErrorState
}

// Name returns state's name
func (s *SellTrade) Name() string {
	return interfaces.SellTradeStateName
}

// ToBuyStock returns a BuyStock instance because it's an allowed state transition
func (s *SellTrade) ToBuyStock() interfaces.State {
	return &BuyStock{}
}
