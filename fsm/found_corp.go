package fsm

type FoundCorp struct {
	BaseState
}

func (s *FoundCorp) ToBuyStock() (State, error) {
	return &BuyStock{}, nil
}

func (s *FoundCorp) ToEndGame() (State, error) {
	return &EndGame{}, nil
}