package fsm

import "testing"

func TestBuyStockToPlayTile(t *testing.T) {
	state := &BuyStock{}

	if state.ToPlayTile().Name() != "PlayTile" {
		t.Errorf("Transition from BuyStock to PlayTile must be valid")
	}
}

func TestBuyStockToSellTrade(t *testing.T) {
	state := &BuyStock{}
	if state.ToSellTrade().Name() != "ErrorState" {
		t.Errorf("Transition from BuyStock to SellTrade must not be possible")
	}
}

func TestBuyStockToFoundCorp(t *testing.T) {
	state := &BuyStock{}
	if state.ToFoundCorp().Name() != "ErrorState" {
		t.Errorf("Transition from BuyStock to FoundCorp must not be possible")
	}
}

func TestBuyStockToUntieMerge(t *testing.T) {
	state := &BuyStock{}
	if state.ToUntieMerge().Name() != "ErrorState" {
		t.Errorf("Transition from BuyStock to UntieMerge must not be possible")
	}
}

func TestBuyStockToEndGame(t *testing.T) {
	state := &BuyStock{}

	if state.ToEndGame().Name() != "EndGame" {
		t.Errorf("Transition from BuyStock to EndGame must be valid")
	}
}
