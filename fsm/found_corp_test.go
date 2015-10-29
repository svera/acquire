package fsm

import "testing"

func TestFoundCorpToPlayTile(t *testing.T) {
	state := &FoundCorp{}
	if _, err := state.ToPlayTile(); err == nil {
		t.Errorf("Transition from FoundCorp to PlayTile must return an error, returned nil")
	}
}

func TestFoundCorpToSellTrade(t *testing.T) {
	state := &FoundCorp{}
	if _, err := state.ToSellTrade(); err == nil {
		t.Errorf("Transition from FoundCorp to SellTrade must return an error, returned nil")
	}
}

func TestFoundCorpToFoundCorp(t *testing.T) {
	state := &FoundCorp{}
	if _, err := state.ToFoundCorp(); err == nil {
		t.Errorf("Transition from FoundCorp to itself must return an error, returned nil")
	}
}

func TestFoundCorpToUntieMerge(t *testing.T) {
	state := &FoundCorp{}
	if _, err := state.ToUntieMerge(); err == nil {
		t.Errorf("Transition from FoundCorp to UntieMerge must return an error, returned nil")
	}
}

func TestFoundCorpToBuyStock(t *testing.T) {
	var state State
	state = &FoundCorp{}
	state, _ = state.ToBuyStock()

	if _, ok := state.(*BuyStock); !ok {
		t.Errorf("Transition from FoundCorp to BuyStock must be valid")
	}
}

func TestFoundCorpToEndGame(t *testing.T) {
	var state State
	state = &FoundCorp{}
	state, _ = state.ToEndGame()

	if _, ok := state.(*EndGame); !ok {
		t.Errorf("Transition from FoundCorp to EndGame must be valid")
	}
}
