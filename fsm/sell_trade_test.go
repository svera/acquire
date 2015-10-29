package fsm

import "testing"

func TestSellTradeToPlayTile(t *testing.T) {
	state := &SellTrade{}
	if _, err := state.ToPlayTile(); err == nil {
		t.Errorf("Transition from SellTrade to PlayTile must return an error, returned nil")
	}
}

func TestSellTradeToSellTrade(t *testing.T) {
	state := &SellTrade{}
	if _, err := state.ToSellTrade(); err == nil {
		t.Errorf("Transition from SellTrade to itself must return an error, returned nil")
	}
}

func TestSellTradeToFoundCorp(t *testing.T) {
	state := &SellTrade{}
	if _, err := state.ToFoundCorp(); err == nil {
		t.Errorf("Transition from SellTrade to FoundCorp must return an error, returned nil")
	}
}

func TestSellTradeToUntieMerge(t *testing.T) {
	state := &SellTrade{}
	if _, err := state.ToUntieMerge(); err == nil {
		t.Errorf("Transition from SellTrade to UntieMerge must return an error, returned nil")
	}
}

func TestSellTradeToBuyStock(t *testing.T) {
	var state State
	state = &SellTrade{}
	state, _ = state.ToBuyStock()

	if _, ok := state.(*BuyStock); !ok {
		t.Errorf("Transition from SellTrade to BuyStock must be valid")
	}
}

func TestSellTradeToEndGame(t *testing.T) {
	var state State
	state = &SellTrade{}
	state, _ = state.ToEndGame()

	if _, ok := state.(*EndGame); !ok {
		t.Errorf("Transition from SellTrade to EndGame must be valid")
	}
}
