package bots

// Message types returned by bots
const (
	PlayTileResponseType   = "playTile"
	NewCorpResponseType    = "newCorp"
	BuyResponseType        = "buy"
	SellTradeResponseType  = "sellTrade"
	UntieMergeResponseType = "untieMerge"
	EndGameResponseType    = "end"
)

// Message is a struct that acts as a container for all types of messages returned
// by bots. A message indicates which action a bot want to do.
type Message struct {
	Type   string
	Params interface{}
}

// PlayTileResponseParams is a struct with all data needed to inform about playing a tile by a bot.
type PlayTileResponseParams struct {
	Tile string
}

// NewCorpResponseParams is a struct with all data needed to inform about founding a new corporation by a bot.
type NewCorpResponseParams struct {
	CorporationIndex int
}

// BuyResponseParams is a struct with all data needed to inform about buying stocks from corporations by a bot.
type BuyResponseParams struct {
	CorporationsIndexes map[string]int
}

// SellTradeResponseParams is a struct with all data needed to inform about selling or trading stocks from corporations by a bot.
type SellTradeResponseParams struct {
	CorporationsIndexes map[string]SellTrade
}

// SellTrade is a struct used by the SellTradeResponseParams one
type SellTrade struct {
	Sell  int
	Trade int
}

// UntieMergeResponseParams is a struct with all data needed to inform about
// which corporation must be the acquirer in a tied merge by a bot.
type UntieMergeResponseParams struct {
	CorporationIndex int
}
