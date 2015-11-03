package fsm

type UntieMerge struct {
	BaseState
}

func (s *UntieMerge) Name() string {
	return "UntieMerge"
}

func (s *UntieMerge) ToSellTrade() (State, error) {
	return &SellTrade{}, nil
}

func (s *UntieMerge) ToEndGame() (State, error) {
	return &EndGame{}, nil
}
