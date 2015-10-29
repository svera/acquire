package fsm

import "testing"

func TestUntieMergeToPlayTile(t *testing.T) {
	state := &UntieMerge{}
	if _, err := state.ToPlayTile(); err == nil {
		t.Errorf("Transition from UntieMerge to PlayTile must return an error, returned nil")
	}
}

func TestUntieMergeToSellTrade(t *testing.T) {
	var state State
	state = &UntieMerge{}
	state, _ = state.ToSellTrade()

	if _, ok := state.(*SellTrade); !ok {
		t.Errorf("Transition from UntieMerge to SellTrade must be valid")
	}
}

func TestUntieMergeToFoundCorp(t *testing.T) {
	state := &UntieMerge{}
	if _, err := state.ToFoundCorp(); err == nil {
		t.Errorf("Transition from UntieMerge to FoundCorp must return an error, returned nil")
	}
}

func TestUntieMergeToUntieMerge(t *testing.T) {
	state := &UntieMerge{}
	if _, err := state.ToUntieMerge(); err == nil {
		t.Errorf("Transition from UntieMerge to itself must return an error, returned nil")
	}
}

func TestUntieMergeToBuyStock(t *testing.T) {
	state := &UntieMerge{}
	if _, err := state.ToBuyStock(); err == nil {
		t.Errorf("Transition from UntieMerge to BuyStock must return an error, returned nil")
	}
}

func TestUntieMergeToEndGame(t *testing.T) {
	var state State
	state = &UntieMerge{}
	state, _ = state.ToEndGame()

	if _, ok := state.(*EndGame); !ok {
		t.Errorf("Transition from UntieMerge to EndGame must be valid")
	}
}
