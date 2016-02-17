package fsm

import (
	"github.com/svera/acquire/interfaces"
	"testing"
)

func TestFoundCorpToPlayTile(t *testing.T) {
	state := &FoundCorp{}
	if state.ToPlayTile().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from FoundCorp to PlayTile must not be valid")
	}
}

func TestFoundCorpToSellTrade(t *testing.T) {
	state := &FoundCorp{}
	if state.ToSellTrade().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from FoundCorp to SellTrade must not be valid")
	}
}

func TestFoundCorpToUntieMerge(t *testing.T) {
	state := &FoundCorp{}
	if state.ToUntieMerge().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from FoundCorp to UntieMerge must not be valid")
	}
}

func TestFoundCorpToBuyStock(t *testing.T) {
	state := &FoundCorp{}
	if state.ToBuyStock().Name() != interfaces.BuyStockStateName {
		t.Errorf("Transition from FoundCorp to BuyStock must be valid")
	}
}
