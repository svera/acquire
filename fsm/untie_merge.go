package fsm

// UntieMerge is a struct representing a finite state machine's state
type UntieMerge struct {
	ErrorState
}

// Name returns state's name
func (s *UntieMerge) Name() string {
	return UntieMergeStateName
}

// ToSellTrade returns a SellTrade instance because it's an allowed state transition
func (s *UntieMerge) ToSellTrade() State {
	return &SellTrade{}
}
