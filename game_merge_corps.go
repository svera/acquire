package acquire

import (
	"errors"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/player"
)

// Taken from the game rules:
// "If only one player owns stock in the defunct corporation, that player gets both bonuses. If there's
// a tie for majority stockholder, add the majority and minority bonuses and divide evenly (the minority
// shareholder gets no bonus. If there's a tie for minority stockholder, split the minority bonus among
// the tied players"
func (g *Game) getMainStockHolders(corp corporation.Interface) map[string][]player.Interface {
	mainStockHolders := map[string][]player.Interface{"majority": {}, "minority": {}}
	stockHolders := g.getStockHolders(corp)

	if len(stockHolders) == 1 {
		return map[string][]player.Interface{
			"majority": {stockHolders[0]},
			"minority": {stockHolders[0]},
		}
	}

	mainStockHolders["majority"] = stockHoldersWithSameAmount(0, stockHolders, corp)
	if len(mainStockHolders["majority"]) > 1 {
		return mainStockHolders
	}
	mainStockHolders["minority"] = stockHoldersWithSameAmount(1, stockHolders, corp)
	return mainStockHolders
}

// Loop stockHolders from start to get all stock holders with the same amount of shares for
// the passed corporation
func stockHoldersWithSameAmount(start int, stockHolders []player.Interface, corp corporation.Interface) []player.Interface {
	group := []player.Interface{stockHolders[start]}

	i := start + 1
	for i < len(stockHolders) && stockHolders[start].Shares(corp) == stockHolders[i].Shares(corp) {
		group = append(group, stockHolders[i])
		i++
	}
	return group
}

// Get players who have stock of the passed corporation, ordered descendently by number of stock shares
// of that corporation
func (g *Game) getStockHolders(corp corporation.Interface) []player.Interface {
	var stockHolders []player.Interface
	sharesDesc := func(pl1, pl2 player.Interface) bool {
		return pl1.Shares(corp) > pl2.Shares(corp)
	}

	for _, pl := range g.players {
		if pl.Shares(corp) > 0 {
			stockHolders = append(stockHolders, pl)
		}
	}
	player.By(sharesDesc).Sort(stockHolders)
	return stockHolders
}

// Checks if two ore more corps are tied for be the acquirer in a merge
// whichs needs the merger player to decide which one would get that role.
func (g *Game) isMergeTied() bool {
	if len(g.mergeCorps["acquirer"]) > 1 {
		return true
	}
	return false
}

// UntieMerge resolves a tied merge selecting which corporation will be the acquirer,
// marking the rest as defunct
func (g *Game) UntieMerge(acquirer corporation.Interface) error {
	if g.state.Name() != "UntieMerge" {
		return errors.New(ActionNotAllowed)
	}
	for i, corp := range g.mergeCorps["acquirer"] {
		if corp == acquirer {
			g.mergeCorps["defunct"] = append(
				g.mergeCorps["defunct"],
				append(g.mergeCorps["acquirer"][:i], g.mergeCorps["acquirer"][i+1:]...)...,
			)
			g.mergeCorps["acquirer"] = []corporation.Interface{corp}
			g.state = g.state.ToSellTrade()
			return nil
		}
	}

	return errors.New(NotAnAcquirerCorporation)
}

// Calculates and returns bonus amounts to be paid to owners of stock of a
// corporation
func (g *Game) payBonuses(corp corporation.Interface) {
	stockHolders := g.getMainStockHolders(corp)
	numberMajorityHolders := len(stockHolders["majority"])
	numberMinorityHolders := len(stockHolders["minority"])

	for _, majorityStockHolder := range stockHolders["majority"] {
		if numberMajorityHolders > 1 {
			majorityStockHolder.AddCash((corp.MajorityBonus() + corp.MinorityBonus()) / numberMajorityHolders)
		} else {
			majorityStockHolder.AddCash(corp.MajorityBonus() / numberMajorityHolders)
		}
	}
	for _, minorityStockHolder := range stockHolders["minority"] {
		minorityStockHolder.AddCash(corp.MinorityBonus() / numberMinorityHolders)
	}
}

// Adds tiles from the defunct corporations to the acquirer one
// and set acquirer as owner of those tiles on board. Finally, resets
// merge information
func (g *Game) completeMerge() {
	acquirer := g.mergeCorps["acquirer"][0]
	for _, defunct := range g.mergeCorps["defunct"] {
		acquirer.Grow(defunct.Size())
		defunct.Reset()
		g.board.ChangeOwner(defunct, acquirer)
	}
	g.lastPlayedTile.SetOwner(acquirer)
	g.board.PutTile(g.lastPlayedTile)
	acquirer.Grow(1)
	g.mergeCorps = map[string][]corporation.Interface{}
}
