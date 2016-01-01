package fsm

import "testing"

func TestFoundCorpToPlayTile(t *testing.T) {
	state := &FoundCorp{}
	if state.ToPlayTile().Name() != "ErrorState" {
		t.Errorf("Transition from FoundCorp to PlayTile must not be valid")
	}
}

func TestFoundCorpToSellTrade(t *testing.T) {
	state := &FoundCorp{}
	if state.ToSellTrade().Name() != "ErrorState" {
		t.Errorf("Transition from FoundCorp to SellTrade must not be valid")
	}
}

func TestFoundCorpToUntieMerge(t *testing.T) {
	state := &FoundCorp{}
	if state.ToUntieMerge().Name() != "ErrorState" {
		t.Errorf("Transition from FoundCorp to UntieMerge must not be valid")
	}
}

func TestFoundCorpToBuyStock(t *testing.T) {
	state := &FoundCorp{}
	if state.ToBuyStock().Name() != "BuyStock" {
		t.Errorf("Transition from FoundCorp to BuyStock must be valid")
	}
}
