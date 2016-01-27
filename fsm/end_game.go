package fsm

// EndGame is a struct representing a finite state machine's state
type EndGame struct {
	ErrorState
}

// Name returns state's name
func (s *EndGame) Name() string {
	return EndGameStateName
}
