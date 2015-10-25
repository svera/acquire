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

type BaseState struct{}

func (s *BaseState) ToPlayTile() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *BaseState) ToFoundCorp() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *BaseState) ToUntieMerge() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *BaseState) ToSellTrade() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *BaseState) ToBuyStock() (State, error) {
	return s, errors.New("State transition not allowed")
}

func (s *BaseState) ToEndGame() (State, error) {
	return s, errors.New("State transition not allowed")
}
