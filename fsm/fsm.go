package fsm

type State interface {
	Name() string
	ToPlayTile() State
	ToFoundCorp() State
	ToUntieMerge() State
	ToSellTrade() State
	ToBuyStock() State
	ToEndGame() State
}

type BaseState struct{}

func (s *BaseState) Name() string {
	return "BaseState"
}

func (s *BaseState) ToPlayTile() State {
	return s
}

func (s *BaseState) ToFoundCorp() State {
	return s
}

func (s *BaseState) ToUntieMerge() State {
	return s
}

func (s *BaseState) ToSellTrade() State {
	return s
}

func (s *BaseState) ToBuyStock() State {
	return s
}

func (s *BaseState) ToEndGame() State {
	return s
}
