package interfaces

// StateMachine is an interface that defines the needed state transitions to be inplemented for this FSM
type StateMachine interface {
	CurrentStateName() string
	ToPlayTile()
	ToFoundCorp()
	ToUntieMerge()
	ToSellTrade()
	ToBuyStock()
	ToEndGame()
	ToInsufficientPlayers()
}
