package fsm

// BuyStock is a finite state machine's state
type BuyStock struct {
	ErrorState
}

// Name returns state's name
func (s *BuyStock) Name() string {
	return BuyStockStateName
}

func (s *BuyStock) ToPlayTile() State {
	return &PlayTile{}
}

func (s *BuyStock) ToEndGame() State {
	return &EndGame{}
}
