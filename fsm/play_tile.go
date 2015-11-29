package fsm

type PlayTile struct {
	BaseState
}

func (s *PlayTile) Name() string {
	return "PlayTile"
}

func (s *PlayTile) ToFoundCorp() State {
	return &FoundCorp{}
}

func (s *PlayTile) ToUntieMerge() State {
	return &UntieMerge{}
}

func (s *PlayTile) ToBuyStock() State {
	return &BuyStock{}
}

func (s *PlayTile) ToEndGame() State {
	return &EndGame{}
}
