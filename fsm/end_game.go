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

// ToEndGame returns self because machine is already in that state
// Although it may make no sense, this is to avoid the machine to go to errorState
func (s *endGame) ToEndGame() interfaces.State {
	return s
}
