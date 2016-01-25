package fsm

import "testing"

func TestSellTradeToPlayTile(t *testing.T) {
	state := &SellTrade{}
	if state.ToPlayTile().Name() != ErrorStateName {
		t.Errorf("Transition from SellTrade to PlayTile must not be valid")
	}
}

func TestSellTradeToFoundCorp(t *testing.T) {
	state := &SellTrade{}
	if state.ToFoundCorp().Name() != ErrorStateName {
		t.Errorf("Transition from SellTrade to FoundCorp must not be valid")
	}
}

func TestSellTradeToUntieMerge(t *testing.T) {
	state := &SellTrade{}
	if state.ToUntieMerge().Name() != ErrorStateName {
		t.Errorf("Transition from SellTrade to UntieMerge must not be valid")
	}
}

func TestSellTradeToBuyStock(t *testing.T) {
	state := &SellTrade{}
	if state.ToBuyStock().Name() != BuyStockStateName {
		t.Errorf("Transition from SellTrade to BuyStock must be valid")
	}
}
