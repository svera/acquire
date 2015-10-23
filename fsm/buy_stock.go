package fsm

import (
	"errors"
)

type BuyStock struct{}

func (s *BuyStock) ToPlayTile() (State, error) {
	return &PlayTile{}, nil
}

func (s *BuyStock) ToFoundCorp() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *BuyStock) ToUntieMerge() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *BuyStock) ToSellTrade() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *BuyStock) ToBuyStock() (State, error) {
	return s, errors.New("Already in state BuyStock")
}

func (s *BuyStock) ToEndGame() (State, error) {
	return &EndGame{}, nil
}
