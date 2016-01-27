package fsm

// ErrorState is a struct representing a finite state machine's state
type ErrorState struct{}

// Name returns state's name
func (s *ErrorState) Name() string {
	return ErrorStateName
}

// ToPlayTile is not a valid state transition
func (s *ErrorState) ToPlayTile() State {
	return s
}

// ToFoundCorp is not a valid state transition
func (s *ErrorState) ToFoundCorp() State {
	return s
}

// ToUntieMerge is not a valid state transition
func (s *ErrorState) ToUntieMerge() State {
	return s
}

// ToSellTrade is not a valid state transition
func (s *ErrorState) ToSellTrade() State {
	return s
}

// ToBuyStock is not a valid state transition
func (s *ErrorState) ToBuyStock() State {
	return s
}

// ToEndGame is not a valid state transition
func (s *ErrorState) ToEndGame() State {
	return s
}
