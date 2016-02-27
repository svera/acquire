package mocks

import (
	"github.com/svera/acquire/interfaces"
)

type State struct {
	FakeStateName string
	TimesCalled   map[string]int
}

func (s *State) Name() string {
	return s.FakeStateName
}

func (s *State) ToPlayTile() interfaces.State {
	s.TimesCalled["ToPlayTile"]++
	return s
}

func (s *State) ToFoundCorp() interfaces.State {
	s.TimesCalled["ToFoundCorp"]++
	return s
}

func (s *State) ToUntieMerge() interfaces.State {
	s.TimesCalled["ToUntieMerge"]++
	return s
}

func (s *State) ToSellTrade() interfaces.State {
	s.TimesCalled["ToSellTrade"]++
	return s
}

func (s *State) ToBuyStock() interfaces.State {
	s.TimesCalled["ToBuyStock"]++
	return s
}

func (s *State) ToEndGame() interfaces.State {
	s.TimesCalled["ToEndGame"]++
	return s
}
