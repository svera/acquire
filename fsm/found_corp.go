package fsm

import (
	"errors"
)

type FoundCorp struct{}

func (s *FoundCorp) ToPlayTile() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *FoundCorp) ToFoundCorp() (State, error) {
	return s, errors.New("Already in state FoundCorp")
}

func (s *FoundCorp) ToUntieMerge() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *FoundCorp) ToSellTrade() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *FoundCorp) ToBuyStock() (State, error) {
	return &SellTrade{}, nil
}

func (s *FoundCorp) ToEndGame() (State, error) {
	return &EndGame{}, nil
}
