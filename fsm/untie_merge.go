package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// UntieMerge is a struct representing a finite state machine's state
type untieMerge struct {
	errorState
}

// Name returns state's name
func (s *untieMerge) Name() string {
	return interfaces.UntieMergeStateName
}

// ToSellTrade returns a SellTrade instance because it's an allowed state transition
func (s *untieMerge) ToSellTrade() interfaces.State {
	return &sellTrade{}
}

// ToInsufficientPlayers returns an InsufficientPlayers instance because it's an allowed state transition
func (s *untieMerge) ToInsufficientPlayers() interfaces.State {
	return &insufficientPlayers{}
}

// ToUntieMerge returns self because machine is already in that state
// Although it may make no sense, this is to avoid the machine to go to errorState
func (s *untieMerge) ToUntieMerge() interfaces.State {
	return s
}
