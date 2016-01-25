package fsm

type State interface {
	Name() string
	ToPlayTile() State
	ToFoundCorp() State
	ToUntieMerge() State
	ToSellTrade() State
	ToBuyStock() State
	ToEndGame() State
}

const (
	ErrorStateName      = "Error"
	EndGameStateName    = "EndGame"
	BuyStockStateName   = "BuyStock"
	FoundCorpStateName  = "FoundCorp"
	PlayTileStateName   = "PlayTile"
	SellTradeStateName  = "SellTrade"
	UntieMergeStateName = "UntieMerge"
)
