package fsm

type ErrorState struct{}

func (s *ErrorState) Name() string {
	return ErrorStateName
}

func (s *ErrorState) ToPlayTile() State {
	return s
}

func (s *ErrorState) ToFoundCorp() State {
	return s
}

func (s *ErrorState) ToUntieMerge() State {
	return s
}

func (s *ErrorState) ToSellTrade() State {
	return s
}

func (s *ErrorState) ToBuyStock() State {
	return s
}

func (s *ErrorState) ToEndGame() State {
	return s
}
