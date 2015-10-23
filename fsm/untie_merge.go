package fsm

import (
	"errors"
)

type UntieMerge struct{}

func (s *UntieMerge) ToPlayTile() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *UntieMerge) ToFoundCorp() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *UntieMerge) ToUntieMerge() (State, error) {
	return s, errors.New("Already in state UntieMerge")
}

func (s *UntieMerge) ToSellTrade() (State, error) {
	return &SellTrade{}, nil
}

func (s *UntieMerge) ToBuyStock() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *UntieMerge) ToEndGame() (State, error) {
	return s, nil
}
