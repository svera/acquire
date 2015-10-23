package fsm

import (
	"errors"
)

type State interface {
	ToPlayTile() (State, error)
	ToFoundCorp() (State, error)
	ToUntieMerge() (State, error)
	ToSellTrade() (State, error)
	ToBuyStock() (State, error)
	ToEndGame() (State, error)
}

type PlayTile struct{}

func (s *PlayTile) ToPlayTile() (State, error) {
	return s, errors.New("Already in state PlayTile")
}

func (s *PlayTile) ToFoundCorp() (State, error) {
	return &FoundCorp{}, nil
}

func (s *PlayTile) ToUntieMerge() (State, error) {
	return &UntieMerge{}, nil
}

func (s *PlayTile) ToSellTrade() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *PlayTile) ToBuyStock() (State, error) {
	return &BuyStock{}, nil
}

func (s *PlayTile) ToEndGame() (State, error) {
	return &EndGame{}, nil
}
