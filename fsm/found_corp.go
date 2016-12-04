package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// foundCorp is a struct representing a finite state machine's state
type foundCorp struct {
	errorState
}

// Name returns state's name
func (s *foundCorp) Name() string {
	return interfaces.FoundCorpStateName
}

// ToBuyStock returns a BuyStock instance because it's an allowed state transition
func (s *foundCorp) ToBuyStock() interfaces.State {
	return &buyStock{}
}

// ToInsufficientPlayers returns an InsufficientPlayers instance because it's an allowed state transition
func (s *foundCorp) ToInsufficientPlayers() interfaces.State {
	return &insufficientPlayers{}
}

// ToFoundCorp returns self because machine is already in that state
// Although it may make no sense, this is to avoid the machine to go to errorState
func (s *foundCorp) ToFoundCorp() interfaces.State {
	return s
}
