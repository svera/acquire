package fsm

import "testing"

func TestPlayTileToSellTrade(t *testing.T) {
	state := &PlayTile{}
	if state.ToSellTrade().Name() != SellTradeStateName {
		t.Errorf("Transition from PlayTile to SellTrade must be valid")
	}
}

func TestPlayTileToFoundCorp(t *testing.T) {
	state := &PlayTile{}
	if state.ToFoundCorp().Name() != FoundCorpStateName {
		t.Errorf("Transition from PlayTile to FoundCorp must be valid")
	}
}

func TestPlayTileToUntieMerge(t *testing.T) {
	state := &PlayTile{}
	if state.ToUntieMerge().Name() != UntieMergeStateName {
		t.Errorf("Transition from PlayTile to UntieMerge must be valid")
	}
}

func TestPlayTileToBuyStock(t *testing.T) {
	state := &PlayTile{}
	if state.ToBuyStock().Name() != BuyStockStateName {
		t.Errorf("Transition from PlayTile to BuyStock must be valid")
	}
}
