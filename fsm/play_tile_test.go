package fsm

import "testing"

func TestTransitionToPlayTile(t *testing.T) {
	state := &PlayTile{}
	if _, err := state.ToPlayTile(); err == nil {
		t.Errorf("Transition from PlayTile to itself must return an error, returned nil")
	}
}

func TestTransitionToSellTrade(t *testing.T) {
	state := &PlayTile{}
	if _, err := state.ToSellTrade(); err == nil {
		t.Errorf("Transition from PlayTile to SellTrade must return an error, returned nil")
	}
}

func TestTransitionToFoundCorp(t *testing.T) {
	var state State
	state = &PlayTile{}
	state, _ = state.ToFoundCorp()

	if _, ok := state.(*FoundCorp); !ok {
		t.Errorf("Transition from PlayTile to FoundCorp must be valid")
	}
}

func TestTransitionToUntieMerge(t *testing.T) {
	var state State
	state = &PlayTile{}
	state, _ = state.ToUntieMerge()

	if _, ok := state.(*UntieMerge); !ok {
		t.Errorf("Transition from PlayTile to UntieMerge must be valid")
	}
}

func TestTransitionToBuyStock(t *testing.T) {
	var state State
	state = &PlayTile{}
	state, _ = state.ToBuyStock()

	if _, ok := state.(*BuyStock); !ok {
		t.Errorf("Transition from PlayTile to BuyStock must be valid")
	}
}

func TestTransitionToEndGame(t *testing.T) {
	var state State
	state = &PlayTile{}
	state, _ = state.ToEndGame()

	if _, ok := state.(*EndGame); !ok {
		t.Errorf("Transition from PlayTile to EndGame must be valid")
	}
}
