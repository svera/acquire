package fsm

import "github.com/svera/acquire/interfaces"

// StateMachine is a struct that works as an state machine for an Acquire game,
// managing its state transitions
type StateMachine struct {
	currentState interfaces.State
}

// New returns a new StateMachine instance
func New() *StateMachine {
	return &StateMachine{
		currentState: &playTile{},
	}
}

// CurrentStateName returns the name of the current state of the machine
func (m *StateMachine) CurrentStateName() string {
	return m.currentState.Name()
}

// ToPlayTile changes the state of the machine to PlayTile
func (m *StateMachine) ToPlayTile() {
	m.currentState = m.currentState.ToPlayTile()
}

// ToFoundCorp changes the state of the machine to FoundCorp
func (m *StateMachine) ToFoundCorp() {
	m.currentState = m.currentState.ToFoundCorp()
}

// ToUntieMerge changes the state of the machine to UntieMerge
func (m *StateMachine) ToUntieMerge() {
	m.currentState = m.currentState.ToUntieMerge()
}

// ToSellTrade changes the state of the machine to SellTrade
func (m *StateMachine) ToSellTrade() {
	m.currentState = m.currentState.ToSellTrade()
}

// ToBuyStock changes the state of the machine to BuyStock
func (m *StateMachine) ToBuyStock() {
	m.currentState = m.currentState.ToBuyStock()
}

// ToEndGame changes the state of the machine to EndGame
func (m *StateMachine) ToEndGame() {
	m.currentState = m.currentState.ToEndGame()
}

// ToInsufficientPlayers changes the state of the machine to InsufficientPlayers
func (m *StateMachine) ToInsufficientPlayers() {
	m.currentState = m.currentState.ToInsufficientPlayers()
}
