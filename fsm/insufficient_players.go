package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// InsufficientPlayers is a struct representing a finite state machine's state
type insufficientPlayers struct {
	errorState
}

// Name returns state's name
func (s *insufficientPlayers) Name() string {
	return interfaces.InsufficientPlayersStateName
}

// ToInsufficientPlayers returns self because machine is already in that state
// Although it may make no sense, this is to avoid the machine to go to errorState
func (s *insufficientPlayers) ToInsufficientPlayers() interfaces.State {
	return s
}
