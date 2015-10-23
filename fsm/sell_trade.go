package fsm

import (
	"errors"
)

type SellTrade struct{}

func (s *SellTrade) ToPlayTile() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *SellTrade) ToFoundCorp() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *SellTrade) ToUntieMerge() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *SellTrade) ToSellTrade() (State, error) {
	return s, errors.New("Already in state SellTrade")
}

func (s *SellTrade) ToBuyStock() (State, error) {
	return &BuyStock{}, nil
}

func (s *SellTrade) ToEndGame() (State, error) {
	return &EndGame{}, nil
}
