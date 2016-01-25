package fsm

type PlayTile struct {
	ErrorState
}

func (s *PlayTile) Name() string {
	return PlayTileStateName
}

func (s *PlayTile) ToSellTrade() State {
	return &SellTrade{}
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
