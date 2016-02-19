package fsm

import (
	"github.com/svera/acquire/interfaces"
	"testing"
)

func TestBuyStockToPlayTile(t *testing.T) {
	state := &BuyStock{}

	if state.ToPlayTile().Name() != interfaces.PlayTileStateName {
		t.Errorf("Transition from BuyStock to PlayTile must be valid")
	}
}

func TestBuyStockToSellTrade(t *testing.T) {
	state := &BuyStock{}
	if state.ToSellTrade().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from BuyStock to SellTrade must not be possible")
	}
}

func TestBuyStockToFoundCorp(t *testing.T) {
	state := &BuyStock{}
	if state.ToFoundCorp().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from BuyStock to FoundCorp must not be possible")
	}
}

func TestBuyStockToUntieMerge(t *testing.T) {
	state := &BuyStock{}
	if state.ToUntieMerge().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from BuyStock to UntieMerge must not be possible")
	}
}

func TestBuyStockToEndGame(t *testing.T) {
	state := &BuyStock{}

	if state.ToEndGame().Name() != interfaces.EndGameStateName {
		t.Errorf("Transition from BuyStock to EndGame must be valid")
	}
}
