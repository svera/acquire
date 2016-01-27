package fsm

// BuyStock is a struct representing a finite state machine's state
type BuyStock struct {
	ErrorState
}

// Name returns state's name
func (s *BuyStock) Name() string {
	return BuyStockStateName
}

// ToPlayTile returns a PlayTile instance because it's an allowed state transition
func (s *BuyStock) ToPlayTile() State {
	return &PlayTile{}
}

// ToEndGame returns an EndGame instance because it's an allowed state transition
func (s *BuyStock) ToEndGame() State {
	return &EndGame{}
}
