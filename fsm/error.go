package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// ErrorState is a struct representing a finite state machine's state
type ErrorState struct{}

// Name returns state's name
func (s *ErrorState) Name() string {
	return interfaces.ErrorStateName
}

// ToPlayTile is not a valid state transition
func (s *ErrorState) ToPlayTile() interfaces.State {
	return s
}

// ToFoundCorp is not a valid state transition
func (s *ErrorState) ToFoundCorp() interfaces.State {
	return s
}

// ToUntieMerge is not a valid state transition
func (s *ErrorState) ToUntieMerge() interfaces.State {
	return s
}

// ToSellTrade is not a valid state transition
func (s *ErrorState) ToSellTrade() interfaces.State {
	return s
}

// ToBuyStock is not a valid state transition
func (s *ErrorState) ToBuyStock() interfaces.State {
	return s
}

// ToEndGame is not a valid state transition
func (s *ErrorState) ToEndGame() interfaces.State {
	return s
}
