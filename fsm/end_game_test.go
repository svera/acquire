package fsm

import (
	"testing"

	"github.com/svera/acquire/interfaces"
)

func TestEndGameToPlayTile(t *testing.T) {
	state := &endGame{}
	if state.ToPlayTile().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from EndGame to PlayTile must not be possible")
	}
}

func TestEndGameToSellTrade(t *testing.T) {
	state := &endGame{}
	if state.ToSellTrade().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from EndGame to SellTrade must not be possible")
	}
}

func TestEndGameToFoundCorp(t *testing.T) {
	state := &endGame{}
	if state.ToFoundCorp().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from EndGame to FoundCorp must not be possible")
	}
}

func TestEndGameToUntieMerge(t *testing.T) {
	state := &endGame{}
	if state.ToUntieMerge().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from EndGame to UntieMerge must not be possible")
	}
}

func TestEndGameToBuyStock(t *testing.T) {
	state := &endGame{}
	if state.ToBuyStock().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from EndGame to BuyStock must not be possible")
	}
}

func TestEndGameToInsufficientPlayers(t *testing.T) {
	state := &endGame{}

	if state.ToInsufficientPlayers().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from EndGame to InsufficientPlayers must not be possible")
	}
}
