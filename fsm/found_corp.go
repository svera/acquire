package fsm

type FoundCorp struct {
	ErrorState
}

func (s *FoundCorp) Name() string {
	return "FoundCorp"
}

func (s *FoundCorp) ToBuyStock() State {
	return &BuyStock{}
}

func (s *FoundCorp) ToEndGame() State {
	return &EndGame{}
}
