package fsm

import "testing"

func TestPlayTileToSellTrade(t *testing.T) {
	state := &PlayTile{}
	if state.ToSellTrade().Name() != "SellTrade" {
		t.Errorf("Transition from PlayTile to SellTrade must be valid")
	}
}

func TestPlayTileToFoundCorp(t *testing.T) {
	state := &PlayTile{}
	if state.ToFoundCorp().Name() != "FoundCorp" {
		t.Errorf("Transition from PlayTile to FoundCorp must be valid")
	}
}

func TestPlayTileToUntieMerge(t *testing.T) {
	state := &PlayTile{}
	if state.ToUntieMerge().Name() != "UntieMerge" {
		t.Errorf("Transition from PlayTile to UntieMerge must be valid")
	}
}

func TestPlayTileToBuyStock(t *testing.T) {
	state := &PlayTile{}
	if state.ToBuyStock().Name() != "BuyStock" {
		t.Errorf("Transition from PlayTile to BuyStock must be valid")
	}
}

func TestPlayTileToEndGame(t *testing.T) {
	state := &PlayTile{}
	if state.ToEndGame().Name() != "EndGame" {
		t.Errorf("Transition from PlayTile to EndGame must be valid")
	}
}
