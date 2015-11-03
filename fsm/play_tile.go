package fsm

type PlayTile struct {
	BaseState
}

func (s *PlayTile) Name() string {
	return "PlayTile"
}

func (s *PlayTile) ToFoundCorp() (State, error) {
	return &FoundCorp{}, nil
}

func (s *PlayTile) ToUntieMerge() (State, error) {
	return &UntieMerge{}, nil
}

func (s *PlayTile) ToBuyStock() (State, error) {
	return &BuyStock{}, nil
}

func (s *PlayTile) ToEndGame() (State, error) {
	return &EndGame{}, nil
}
