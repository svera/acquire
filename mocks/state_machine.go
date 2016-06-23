package mocks

// StateMachine is a structure that implements the StateMachine interface for testing
type StateMachine struct {
	FakeStateName string
	TimesCalled   map[string]int
}

// CurrentStateName mocks the Name method defined in the StateMachine interface
func (s *StateMachine) CurrentStateName() string {
	return s.FakeStateName
}

// ToPlayTile mocks the ToPlayTile method defined in the StateMachine interface
func (s *StateMachine) ToPlayTile() {
	s.TimesCalled["ToPlayTile"]++
}

// ToFoundCorp mocks the ToFoundCorp method defined in the StateMachine interface
func (s *StateMachine) ToFoundCorp() {
	s.TimesCalled["ToFoundCorp"]++
}

// ToUntieMerge mocks the ToUntieMerge method defined in the StateMachine interface
func (s *StateMachine) ToUntieMerge() {
	s.TimesCalled["ToUntieMerge"]++
}

// ToSellTrade mocks the ToSellTrade method defined in the StateMachine interface
func (s *StateMachine) ToSellTrade() {
	s.TimesCalled["ToSellTrade"]++
}

// ToBuyStock mocks the ToBuyStock method defined in the StateMachine interface
func (s *StateMachine) ToBuyStock() {
	s.TimesCalled["ToBuyStock"]++
}

// ToEndGame mocks the ToEndGame method defined in the StateMachine interface
func (s *StateMachine) ToEndGame() {
	s.TimesCalled["ToEndGame"]++
}

// ToInsufficientPlayers mocks the ToInsufficientPlayers method defined in the StateMachine interface
func (s *StateMachine) ToInsufficientPlayers() {
	s.TimesCalled["ToInsufficientPlayers"]++
}
