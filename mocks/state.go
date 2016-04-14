package mocks

import (
	"github.com/svera/acquire/interfaces"
)

// State is a structure that implements the State interface for testing
type State struct {
	FakeStateName string
	TimesCalled   map[string]int
}

// Name mocks the Name method defined in the State interface
func (s *State) Name() string {
	return s.FakeStateName
}

// ToPlayTile mocks the ToPlayTile method defined in the State interface
func (s *State) ToPlayTile() interfaces.State {
	s.TimesCalled["ToPlayTile"]++
	return s
}

// ToFoundCorp mocks the ToFoundCorp method defined in the State interface
func (s *State) ToFoundCorp() interfaces.State {
	s.TimesCalled["ToFoundCorp"]++
	return s
}

// ToUntieMerge mocks the ToUntieMerge method defined in the State interface
func (s *State) ToUntieMerge() interfaces.State {
	s.TimesCalled["ToUntieMerge"]++
	return s
}

// ToSellTrade mocks the ToSellTrade method defined in the State interface
func (s *State) ToSellTrade() interfaces.State {
	s.TimesCalled["ToSellTrade"]++
	return s
}

// ToBuyStock mocks the ToBuyStock method defined in the State interface
func (s *State) ToBuyStock() interfaces.State {
	s.TimesCalled["ToBuyStock"]++
	return s
}

// ToEndGame mocks the ToEndGame method defined in the State interface
func (s *State) ToEndGame() interfaces.State {
	s.TimesCalled["ToEndGame"]++
	return s
}
