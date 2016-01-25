package fsm

import "testing"

func TestUntieMergeToPlayTile(t *testing.T) {
	state := &UntieMerge{}
	if state.ToPlayTile().Name() != ErrorStateName {
		t.Errorf("Transition from UntieMerge to PlayTile must return not be valid")
	}
}

func TestUntieMergeToSellTrade(t *testing.T) {
	state := &UntieMerge{}
	if state.ToSellTrade().Name() != SellTradeStateName {
		t.Errorf("Transition from UntieMerge to SellTrade must be valid")
	}
}

func TestUntieMergeToFoundCorp(t *testing.T) {
	state := &UntieMerge{}
	if state.ToFoundCorp().Name() != ErrorStateName {
		t.Errorf("Transition from UntieMerge to FoundCorp must return not be valid")
	}
}

func TestUntieMergeToBuyStock(t *testing.T) {
	state := &UntieMerge{}
	if state.ToBuyStock().Name() != ErrorStateName {
		t.Errorf("Transition from UntieMerge to BuyStock must return not be valid")
	}
}
