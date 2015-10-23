package fsm

import (
	"errors"
)

type EndGame struct{}

func (s *EndGame) ToPlayTile() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *EndGame) ToFoundCorp() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *EndGame) ToUntieMerge() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *EndGame) ToSellTrade() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *EndGame) ToBuyStock() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *EndGame) ToEndGame() (State, error) {
	return s, errors.New("Already in state EndGame")
}
