package fsm

import "testing"

func TestBuyStockToPlayTile(t *testing.T) {
	state := &BuyStock{}

	if state.ToPlayTile().Name() != PlayTileStateName {
		t.Errorf("Transition from BuyStock to PlayTile must be valid")
	}
}

func TestBuyStockToSellTrade(t *testing.T) {
	state := &BuyStock{}
	if state.ToSellTrade().Name() != ErrorStateName {
		t.Errorf("Transition from BuyStock to SellTrade must not be possible")
	}
}

func TestBuyStockToFoundCorp(t *testing.T) {
	state := &BuyStock{}
	if state.ToFoundCorp().Name() != ErrorStateName {
		t.Errorf("Transition from BuyStock to FoundCorp must not be possible")
	}
}

func TestBuyStockToUntieMerge(t *testing.T) {
	state := &BuyStock{}
	if state.ToUntieMerge().Name() != ErrorStateName {
		t.Errorf("Transition from BuyStock to UntieMerge must not be possible")
	}
}

func TestBuyStockToEndGame(t *testing.T) {
	state := &BuyStock{}

	if state.ToEndGame().Name() != EndGameStateName {
		t.Errorf("Transition from BuyStock to EndGame must be valid")
	}
}
