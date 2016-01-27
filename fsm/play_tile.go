package fsm

// PlayTile is a struct representing a finite state machine's state
type PlayTile struct {
	ErrorState
}

// Name returns state's name
func (s *PlayTile) Name() string {
	return PlayTileStateName
}

// ToSellTrade returns a SellTrade instance because it's an allowed state transition
func (s *PlayTile) ToSellTrade() State {
	return &SellTrade{}
}

// ToFoundCorp returns a FoundCorp instance because it's an allowed state transition
func (s *PlayTile) ToFoundCorp() State {
	return &FoundCorp{}
}

// ToUntieMerge returns a UntieMerge instance because it's an allowed state transition
func (s *PlayTile) ToUntieMerge() State {
	return &UntieMerge{}
}

// ToBuyStock returns a BuyStock instance because it's an allowed state transition
func (s *PlayTile) ToBuyStock() State {
	return &BuyStock{}
}
