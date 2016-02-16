package fsm

import (
	"github.com/svera/acquire"
)

// ErrorState is a struct representing a finite state machine's state
type ErrorState struct{}

// Name returns state's name
func (s *ErrorState) Name() string {
	return ErrorStateName
}

// ToPlayTile is not a valid state transition
func (s *ErrorState) ToPlayTile() acquire.State {
	return s
}

// ToFoundCorp is not a valid state transition
func (s *ErrorState) ToFoundCorp() acquire.State {
	return s
}

// ToUntieMerge is not a valid state transition
func (s *ErrorState) ToUntieMerge() acquire.State {
	return s
}

// ToSellTrade is not a valid state transition
func (s *ErrorState) ToSellTrade() acquire.State {
	return s
}

// ToBuyStock is not a valid state transition
func (s *ErrorState) ToBuyStock() acquire.State {
	return s
}

// ToEndGame is not a valid state transition
func (s *ErrorState) ToEndGame() acquire.State {
	return s
}
