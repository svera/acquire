package fsm

import "testing"

func TestSellTradeToPlayTile(t *testing.T) {
	state := &SellTrade{}
	if state.ToPlayTile().Name() != "ErrorState" {
		t.Errorf("Transition from SellTrade to PlayTile must not be valid")
	}
}

func TestSellTradeToFoundCorp(t *testing.T) {
	state := &SellTrade{}
	if state.ToFoundCorp().Name() != "ErrorState" {
		t.Errorf("Transition from SellTrade to FoundCorp must not be valid")
	}
}

func TestSellTradeToUntieMerge(t *testing.T) {
	state := &SellTrade{}
	if state.ToUntieMerge().Name() != "ErrorState" {
		t.Errorf("Transition from SellTrade to UntieMerge must not be valid")
	}
}

func TestSellTradeToBuyStock(t *testing.T) {
	state := &SellTrade{}
	if state.ToBuyStock().Name() != "BuyStock" {
		t.Errorf("Transition from SellTrade to BuyStock must be valid")
	}
}

func TestSellTradeToEndGame(t *testing.T) {
	state := &SellTrade{}
	if state.ToEndGame().Name() != "EndGame" {
		t.Errorf("Transition from SellTrade to EndGame must be valid")
	}
}
