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
