package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// ErrorState is a struct representing a finite state machine's state
// This state must only be reached when a unsupported transition is tried
// It is reached automatically, there is not a specific transition to this state
type errorState struct{}

// Name returns state's name
func (s *errorState) Name() string {
	return interfaces.ErrorStateName
}

// ToPlayTile is not a valid state transition
func (s *errorState) ToPlayTile() interfaces.State {
	return s
}

// ToFoundCorp is not a valid state transition
func (s *errorState) ToFoundCorp() interfaces.State {
	return s
}

// ToUntieMerge is not a valid state transition
func (s *errorState) ToUntieMerge() interfaces.State {
	return s
}

// ToSellTrade is not a valid state transition
func (s *errorState) ToSellTrade() interfaces.State {
	return s
}

// ToBuyStock is not a valid state transition
func (s *errorState) ToBuyStock() interfaces.State {
	return s
}

// ToEndGame is not a valid state transition
func (s *errorState) ToEndGame() interfaces.State {
	return s
}

// ToInsufficientPlayers is not a valid state transition
func (s *errorState) ToInsufficientPlayers() interfaces.State {
	return s
}
