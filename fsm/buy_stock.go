package fsm

type BuyStock struct {
	BaseState
}

func (s *BuyStock) Name() string {
	return "BuyStock"
}

func (s *BuyStock) ToPlayTile() State {
	return &PlayTile{}
}

func (s *BuyStock) ToEndGame() State {
	return &EndGame{}
}
