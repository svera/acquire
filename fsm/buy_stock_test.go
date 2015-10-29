package fsm

import "testing"

func TestBuyStockToPlayTile(t *testing.T) {
	var state State
	state = &BuyStock{}
	state, _ = state.ToPlayTile()

	if _, ok := state.(*PlayTile); !ok {
		t.Errorf("Transition from BuyStock to PlayTile must be valid")
	}
}

func TestBuyStockToSellTrade(t *testing.T) {
	state := &BuyStock{}
	if _, err := state.ToSellTrade(); err == nil {
		t.Errorf("Transition from BuyStock to SellTrade must return an error, returned nil")
	}
}

func TestBuyStockToFoundCorp(t *testing.T) {
	state := &BuyStock{}
	if _, err := state.ToFoundCorp(); err == nil {
		t.Errorf("Transition from BuyStock to SellTrade must return an error, returned nil")
	}
}

func TestBuyStockToUntieMerge(t *testing.T) {
	state := &BuyStock{}
	if _, err := state.ToUntieMerge(); err == nil {
		t.Errorf("Transition from BuyStock to SellTrade must return an error, returned nil")
	}
}

func TestBuyStockToBuyStock(t *testing.T) {
	state := &BuyStock{}
	if _, err := state.ToBuyStock(); err == nil {
		t.Errorf("Transition from BuyStock to itself must return an error, returned nil")
	}
}

func TestBuyStockToEndGame(t *testing.T) {
	var state State
	state = &BuyStock{}
	state, _ = state.ToEndGame()

	if _, ok := state.(*EndGame); !ok {
		t.Errorf("Transition from BuyStock to EndGame must be valid")
	}
}
