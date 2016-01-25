package fsm

type FoundCorp struct {
	ErrorState
}

func (s *FoundCorp) Name() string {
	return FoundCorpStateName
}

func (s *FoundCorp) ToBuyStock() State {
	return &BuyStock{}
}
