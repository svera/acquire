package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// EndGame is a struct representing a finite state machine's state
type endGame struct {
	errorState
}

// Name returns state's name
func (s *endGame) Name() string {
	return interfaces.EndGameStateName
}
