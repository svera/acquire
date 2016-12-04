package fsm

import (
	"testing"

	"github.com/svera/acquire/interfaces"
)

func TestInsufficientPlayersToPlayTile(t *testing.T) {
	state := &insufficientPlayers{}
	if state.ToPlayTile().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from InsufficientPlayers to PlayTile must not be possible")
	}
}

func TestInsufficientPlayersToSellTrade(t *testing.T) {
	state := &insufficientPlayers{}
	if state.ToSellTrade().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from InsufficientPlayers to SellTrade must not be possible")
	}
}

func TestInsufficientPlayersToFoundCorp(t *testing.T) {
	state := &insufficientPlayers{}
	if state.ToFoundCorp().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from InsufficientPlayers to FoundCorp must not be possible")
	}
}

func TestInsufficientPlayersToUntieMerge(t *testing.T) {
	state := &insufficientPlayers{}
	if state.ToUntieMerge().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from InsufficientPlayers to UntieMerge must not be possible")
	}
}

func TestInsufficientPlayersToBuyStock(t *testing.T) {
	state := &insufficientPlayers{}
	if state.ToBuyStock().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from InsufficientPlayers to BuyStock must not be possible")
	}
}

func TestInsufficientPlayersToEndGame(t *testing.T) {
	state := &insufficientPlayers{}

	if state.ToEndGame().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from InsufficientPlayers to EndGame must not be possible")
	}
}

func TestInsufficientPlayersToInsufficientPlayers(t *testing.T) {
	state := &insufficientPlayers{}

	if state.ToInsufficientPlayers().Name() != interfaces.InsufficientPlayersStateName {
		t.Errorf("Transition from InsufficientPlayers to InsufficientPlayers must be valid")
	}
}
