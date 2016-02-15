package fsm

import (
	"github.com/svera/acquire"
)

// FoundCorp is a struct representing a finite state machine's state
type FoundCorp struct {
	ErrorState
}

// Name returns state's name
func (s *FoundCorp) Name() string {
	return FoundCorpStateName
}

// ToBuyStock returns a BuyStock instance because it's an allowed state transition
func (s *FoundCorp) ToBuyStock() acquire.State {
	return &BuyStock{}
}
