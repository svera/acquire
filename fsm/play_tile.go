package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// PlayTile is a struct representing a finite state machine's state
type PlayTile struct {
	ErrorState
}

// Name returns state's name
func (s *PlayTile) Name() string {
	return interfaces.PlayTileStateName
}

// ToSellTrade returns a SellTrade instance because it's an allowed state transition
func (s *PlayTile) ToSellTrade() interfaces.State {
	return &SellTrade{}
}

// ToFoundCorp returns a FoundCorp instance because it's an allowed state transition
func (s *PlayTile) ToFoundCorp() interfaces.State {
	return &FoundCorp{}
}

// ToUntieMerge returns a UntieMerge instance because it's an allowed state transition
func (s *PlayTile) ToUntieMerge() interfaces.State {
	return &UntieMerge{}
}

// ToBuyStock returns a BuyStock instance because it's an allowed state transition
func (s *PlayTile) ToBuyStock() interfaces.State {
	return &BuyStock{}
}
