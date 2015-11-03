package fsm

type BuyStock struct {
	BaseState
}

func (s *BuyStock) Name() string {
	return "BuyStock"
}

func (s *BuyStock) ToPlayTile() (State, error) {
	return &PlayTile{}, nil
}

func (s *BuyStock) ToEndGame() (State, error) {
	return &EndGame{}, nil
}
