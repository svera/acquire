package fsm

import (
	"github.com/svera/acquire/interfaces"
	"testing"
)

func TestUntieMergeToPlayTile(t *testing.T) {
	state := &UntieMerge{}
	if state.ToPlayTile().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from UntieMerge to PlayTile must return not be valid")
	}
}

func TestUntieMergeToSellTrade(t *testing.T) {
	state := &UntieMerge{}
	if state.ToSellTrade().Name() != interfaces.SellTradeStateName {
		t.Errorf("Transition from UntieMerge to SellTrade must be valid")
	}
}

func TestUntieMergeToFoundCorp(t *testing.T) {
	state := &UntieMerge{}
	if state.ToFoundCorp().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from UntieMerge to FoundCorp must return not be valid")
	}
}

func TestUntieMergeToBuyStock(t *testing.T) {
	state := &UntieMerge{}
	if state.ToBuyStock().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from UntieMerge to BuyStock must return not be valid")
	}
}
