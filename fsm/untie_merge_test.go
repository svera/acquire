package fsm

import "testing"

func TestUntieMergeToPlayTile(t *testing.T) {
	state := &UntieMerge{}
	if state.ToPlayTile().Name() != "ErrorState" {
		t.Errorf("Transition from UntieMerge to PlayTile must return not be valid")
	}
}

func TestUntieMergeToSellTrade(t *testing.T) {
	state := &UntieMerge{}
	if state.ToSellTrade().Name() != "SellTrade" {
		t.Errorf("Transition from UntieMerge to SellTrade must be valid")
	}
}

func TestUntieMergeToFoundCorp(t *testing.T) {
	state := &UntieMerge{}
	if state.ToFoundCorp().Name() != "ErrorState" {
		t.Errorf("Transition from UntieMerge to FoundCorp must return not be valid")
	}
}

func TestUntieMergeToBuyStock(t *testing.T) {
	state := &UntieMerge{}
	if state.ToBuyStock().Name() != "ErrorState" {
		t.Errorf("Transition from UntieMerge to BuyStock must return not be valid")
	}
}
