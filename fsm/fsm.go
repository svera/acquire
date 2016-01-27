package fsm

// State is an interface that defines the needed state transitions to be inplemented for this FSM
type State interface {
	Name() string
	ToPlayTile() State
	ToFoundCorp() State
	ToUntieMerge() State
	ToSellTrade() State
	ToBuyStock() State
	ToEndGame() State
}

// State names
const (
	ErrorStateName      = "Error"
	EndGameStateName    = "EndGame"
	BuyStockStateName   = "BuyStock"
	FoundCorpStateName  = "FoundCorp"
	PlayTileStateName   = "PlayTile"
	SellTradeStateName  = "SellTrade"
	UntieMergeStateName = "UntieMerge"
)
