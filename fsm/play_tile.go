package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// playTile is a struct representing a finite state machine's state
type playTile struct {
	errorState
}

// Name returns state's name
func (s *playTile) Name() string {
	return interfaces.PlayTileStateName
}

// ToSellTrade returns a SellTrade instance because it's an allowed state transition
func (s *playTile) ToSellTrade() interfaces.State {
	return &sellTrade{}
}

// ToFoundCorp returns a FoundCorp instance because it's an allowed state transition
func (s *playTile) ToFoundCorp() interfaces.State {
	return &foundCorp{}
}

// ToUntieMerge returns a UntieMerge instance because it's an allowed state transition
func (s *playTile) ToUntieMerge() interfaces.State {
	return &untieMerge{}
}

// ToBuyStock returns a BuyStock instance because it's an allowed state transition
func (s *playTile) ToBuyStock() interfaces.State {
	return &buyStock{}
}

// ToInsufficientPlayers returns an InsufficientPlayers instance because it's an allowed state transition
func (s *playTile) ToInsufficientPlayers() interfaces.State {
	return &insufficientPlayers{}
}

// ToEndGame returns an EndGame instance because it's an allowed state transition
func (s *playTile) ToEndGame() interfaces.State {
	return &endGame{}
}

// ToPlayTile returns self because machine is already in that state
// Although it may make no sense, this is to avoid the machine to go to errorState
func (s *playTile) ToPlayTile() interfaces.State {
	return s
}
