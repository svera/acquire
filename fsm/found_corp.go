package fsm

type FoundCorp struct {
	BaseState
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
