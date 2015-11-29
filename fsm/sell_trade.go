package fsm

type SellTrade struct {
	ErrorState
}

func (s *SellTrade) Name() string {
	return "SellTrade"
}

func (s *SellTrade) ToBuyStock() State {
	return &BuyStock{}
}

func (s *SellTrade) ToEndGame() State {
	return &EndGame{}
}
