package fsm

import (
	"testing"

	"github.com/svera/acquire/interfaces"
)

func TestUntieMergeToPlayTile(t *testing.T) {
	state := &untieMerge{}
	if state.ToPlayTile().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from UntieMerge to PlayTile must return not be valid")
	}
}

func TestUntieMergeToSellTrade(t *testing.T) {
	state := &untieMerge{}
	if state.ToSellTrade().Name() != interfaces.SellTradeStateName {
		t.Errorf("Transition from UntieMerge to SellTrade must be valid")
	}
}

func TestUntieMergeToFoundCorp(t *testing.T) {
	state := &untieMerge{}
	if state.ToFoundCorp().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from UntieMerge to FoundCorp must return not be valid")
	}
}

func TestUntieMergeToBuyStock(t *testing.T) {
	state := &untieMerge{}
	if state.ToBuyStock().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from UntieMerge to BuyStock must return not be valid")
	}
}

func TestUntieMergeToInsufficientPlayers(t *testing.T) {
	state := &untieMerge{}

	if state.ToInsufficientPlayers().Name() != interfaces.InsufficientPlayersStateName {
		t.Errorf("Transition from UntieMerge to InsufficientPlayers must be valid")
	}
}

func TestUntieMergeToUntieMerge(t *testing.T) {
	state := &untieMerge{}

	if state.ToUntieMerge().Name() != interfaces.UntieMergeStateName {
		t.Errorf("Transition from UntieMerge to UntieMerge must be valid")
	}
}
