package bots

const (
	PlayTileResponseType   = "playTile"
	NewCorpResponseType    = "newCorp"
	BuyResponseType        = "buy"
	SellTradeResponseType  = "sellTrade"
	UntieMergeResponseType = "untieMerge"
	EndGameResponseType    = "end"
)

type Message struct {
	Type   string
	Params interface{}
}

type PlayTileResponseParams struct {
	Tile string
}

type NewCorpResponseParams struct {
	Corporation string
}

type BuyResponseParams struct {
	Corporations map[string]int
}

type SellTradeResponseParams struct {
	Corporations map[string]SellTrade
}

type SellTrade struct {
	Sell  int
	Trade int
}

type UntieMergeResponseParams struct {
	Corporation string
}
