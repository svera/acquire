package fsm

import (
	"testing"

	"github.com/svera/acquire/interfaces"
)

func TestPlayTileToSellTrade(t *testing.T) {
	state := &playTile{}
	if state.ToSellTrade().Name() != interfaces.SellTradeStateName {
		t.Errorf("Transition from PlayTile to SellTrade must be valid")
	}
}

func TestPlayTileToFoundCorp(t *testing.T) {
	state := &playTile{}
	if state.ToFoundCorp().Name() != interfaces.FoundCorpStateName {
		t.Errorf("Transition from PlayTile to FoundCorp must be valid")
	}
}

func TestPlayTileToUntieMerge(t *testing.T) {
	state := &playTile{}
	if state.ToUntieMerge().Name() != interfaces.UntieMergeStateName {
		t.Errorf("Transition from PlayTile to UntieMerge must be valid")
	}
}

func TestPlayTileToBuyStock(t *testing.T) {
	state := &playTile{}
	if state.ToBuyStock().Name() != interfaces.BuyStockStateName {
		t.Errorf("Transition from PlayTile to BuyStock must be valid")
	}
}

func TestPlayTileToInsufficientPlayers(t *testing.T) {
	state := &playTile{}

	if state.ToInsufficientPlayers().Name() != interfaces.InsufficientPlayersStateName {
		t.Errorf("Transition from PlayTile to InsufficientPlayers must be valid")
	}
}

func TestPlayTileToEndGame(t *testing.T) {
	state := &playTile{}

	if state.ToEndGame().Name() != interfaces.EndGameStateName {
		t.Errorf("Transition from PlayTile to EndGame must be valid")
	}
}

func TestPlayTileToPlayTile(t *testing.T) {
	state := &playTile{}

	if state.ToPlayTile().Name() != interfaces.PlayTileStateName {
		t.Errorf("Transition from PlayTile to PlayTile must be valid")
	}
}
