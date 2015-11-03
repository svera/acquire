package fsm

import (
	"errors"
)

type State interface {
	Name() string
	ToPlayTile() (State, error)
	ToFoundCorp() (State, error)
	ToUntieMerge() (State, error)
	ToSellTrade() (State, error)
	ToBuyStock() (State, error)
	ToEndGame() (State, error)
}

const (
	StateTransitionNotAllowed = "State transition not allowed"
)

type BaseState struct{}

func (s *BaseState) Name() string {
	return "BaseState"
}

func (s *BaseState) ToPlayTile() (State, error) {
	return s, errors.New(StateTransitionNotAllowed)
}

func (s *BaseState) ToFoundCorp() (State, error) {
	return s, errors.New(StateTransitionNotAllowed)
}

func (s *BaseState) ToUntieMerge() (State, error) {
	return s, errors.New(StateTransitionNotAllowed)
}

func (s *BaseState) ToSellTrade() (State, error) {
	return s, errors.New(StateTransitionNotAllowed)
}

func (s *BaseState) ToBuyStock() (State, error) {
	return s, errors.New(StateTransitionNotAllowed)
}

func (s *BaseState) ToEndGame() (State, error) {
	return s, errors.New(StateTransitionNotAllowed)
}
