package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// InsufficientPlayers is a struct representing a finite state machine's state
type InsufficientPlayers struct {
	ErrorState
}

// Name returns state's name
func (s *InsufficientPlayers) Name() string {
	return interfaces.InsufficientPlayersStateName
}
