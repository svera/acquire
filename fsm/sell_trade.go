package fsm

type SellTrade struct {
	BaseState
}

func (s *SellTrade) Name() string {
	return "SellTrade"
}

func (s *SellTrade) ToBuyStock() (State, error) {
	return &BuyStock{}, nil
}

func (s *SellTrade) ToEndGame() (State, error) {
	return &EndGame{}, nil
}
