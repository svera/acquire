package fsm

import "testing"

func TestEndGameToPlayTile(t *testing.T) {
	state := &EndGame{}
	if state.ToPlayTile().Name() != "EndGame" {
		t.Errorf("Transition from EndGame to PlayTile must not be possible")
	}
}

func TestEndGameToSellTrade(t *testing.T) {
	state := &EndGame{}
	if state.ToSellTrade().Name() != "EndGame" {
		t.Errorf("Transition from EndGame to SellTrade must not be possible")
	}
}

func TestEndGameToFoundCorp(t *testing.T) {
	state := &EndGame{}
	if state.ToFoundCorp().Name() != "EndGame" {
		t.Errorf("Transition from EndGame to FoundCorp must not be possible")
	}
}

func TestEndGameToUntieMerge(t *testing.T) {
	state := &EndGame{}
	if state.ToUntieMerge().Name() != "EndGame" {
		t.Errorf("Transition from EndGame to UntieMerge must not be possible")
	}
}

func TestEndGameToBuyStock(t *testing.T) {
	state := &EndGame{}
	if state.ToBuyStock().Name() != "EndGame" {
		t.Errorf("Transition from EndGame to BuyStock must not be possible")
	}
}
