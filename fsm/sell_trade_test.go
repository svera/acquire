package fsm

import (
	"testing"

	"github.com/svera/acquire/interfaces"
)

func TestSellTradeToPlayTile(t *testing.T) {
	state := &sellTrade{}
	if state.ToPlayTile().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from SellTrade to PlayTile must not be valid")
	}
}

func TestSellTradeToFoundCorp(t *testing.T) {
	state := &sellTrade{}
	if state.ToFoundCorp().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from SellTrade to FoundCorp must not be valid")
	}
}

func TestSellTradeToUntieMerge(t *testing.T) {
	state := &sellTrade{}
	if state.ToUntieMerge().Name() != interfaces.ErrorStateName {
		t.Errorf("Transition from SellTrade to UntieMerge must not be valid")
	}
}

func TestSellTradeToBuyStock(t *testing.T) {
	state := &sellTrade{}
	if state.ToBuyStock().Name() != interfaces.BuyStockStateName {
		t.Errorf("Transition from SellTrade to BuyStock must be valid")
	}
}

func TestSellTradeToInsufficientPlayers(t *testing.T) {
	state := &sellTrade{}

	if state.ToInsufficientPlayers().Name() != interfaces.InsufficientPlayersStateName {
		t.Errorf("Transition from SellTrade to InsufficientPlayers must be valid")
	}
}

func TestSellTradeToSellTrade(t *testing.T) {
	state := &sellTrade{}

	if state.ToSellTrade().Name() != interfaces.SellTradeStateName {
		t.Errorf("Transition from SellTrade to SellTrade must be valid")
	}
}
