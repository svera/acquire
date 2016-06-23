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
