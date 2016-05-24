package fsm

import (
	"github.com/svera/acquire/interfaces"
)

// BuyStock is a struct representing a finite state machine's state
type BuyStock struct {
	ErrorState
}

// Name returns state's name
func (s *BuyStock) Name() string {
	return interfaces.BuyStockStateName
}

// ToPlayTile returns a PlayTile instance because it's an allowed state transition
func (s *BuyStock) ToPlayTile() interfaces.State {
	return &PlayTile{}
}

// ToEndGame returns an EndGame instance because it's an allowed state transition
func (s *BuyStock) ToEndGame() interfaces.State {
	return &EndGame{}
}

// ToInsufficientPlayers returns an InsufficientPlayers instance because it's an allowed state transition
func (s *BuyStock) ToInsufficientPlayers() interfaces.State {
	return &InsufficientPlayers{}
}
