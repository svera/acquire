package fsm

type UntieMerge struct {
	ErrorState
}

func (s *UntieMerge) Name() string {
	return "UntieMerge"
}

func (s *UntieMerge) ToSellTrade() State {
	return &SellTrade{}
}

func (s *UntieMerge) ToEndGame() State {
	return &EndGame{}
}
