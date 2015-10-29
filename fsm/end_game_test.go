package fsm

import "testing"

func TestEndGameToPlayTile(t *testing.T) {
	state := &EndGame{}
	if _, err := state.ToPlayTile(); err == nil {
		t.Errorf("Transition from EndGame to PlayTile must return an error, returned nil")
	}
}

func TestEndGameToSellTrade(t *testing.T) {
	state := &EndGame{}
	if _, err := state.ToSellTrade(); err == nil {
		t.Errorf("Transition from EndGame to SellTrade must return an error, returned nil")
	}
}

func TestEndGameToFoundCorp(t *testing.T) {
	state := &EndGame{}
	if _, err := state.ToFoundCorp(); err == nil {
		t.Errorf("Transition from EndGame to itself must return an error, returned nil")
	}
}

func TestEndGameToUntieMerge(t *testing.T) {
	state := &EndGame{}
	if _, err := state.ToUntieMerge(); err == nil {
		t.Errorf("Transition from EndGame to UntieMerge must return an error, returned nil")
	}
}

func TestEndGameToBuyStock(t *testing.T) {
	state := &EndGame{}
	if _, err := state.ToBuyStock(); err == nil {
		t.Errorf("Transition from EndGame to BuyStock must return an error, returned nil")
	}
}

func TestEndGameToEndGame(t *testing.T) {
	state := &EndGame{}
	if _, err := state.ToEndGame(); err == nil {
		t.Errorf("Transition from EndGame to itself must return an error, returned nil")
	}
}
