package acquire

import (
	"errors"
	"github.com/svera/acquire/interfaces"
	"sort"
)

func (g *Game) startMerge(tl interfaces.Tile, mergeCorps map[string][]interfaces.Corporation) {
	g.board.PutTile(tl)
	g.mergeCorps = mergeCorps
	if g.isMergeTied() {
		g.state = g.state.ToUntieMerge()
	} else {
		for _, corp := range mergeCorps["defunct"] {
			g.payBonuses(corp)
		}
		g.sellTradePlayers = g.setSellTradePlayers(mergeCorps["defunct"])
		g.frozenPlayer = g.currentPlayerNumber
		g.setCurrentPlayer(g.nextSellTradePlayer())
		g.state = g.state.ToSellTrade()
	}
}

// Calculates and returns bonus amounts to be paid to owners of stock of a
// corporation
func (g *Game) payBonuses(corp interfaces.Corporation) {
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

// Taken from the game rules:
// "If only one player owns stock in the defunct corporation, that player gets both bonuses. If there's
// a tie for majority stockholder, add the majority and minority bonuses and divide evenly (the minority
// shareholder gets no bonus. If there's a tie for minority stockholder, split the minority bonus among
// the tied players"
func (g *Game) getMainStockHolders(corp interfaces.Corporation) map[string][]interfaces.Player {
	mainStockHolders := map[string][]interfaces.Player{"majority": {}, "minority": {}}
	stockHolders := g.getStockHolders(corp)

	if len(stockHolders) == 1 {
		return map[string][]interfaces.Player{
			"majority": {stockHolders[0]},
			"minority": {stockHolders[0]},
		}
	}

	if len(stockHolders) > 0 {
		mainStockHolders["majority"] = stockHoldersWithSameAmount(0, stockHolders, corp)
		if len(mainStockHolders["majority"]) > 1 {
			return mainStockHolders
		}
		mainStockHolders["minority"] = stockHoldersWithSameAmount(1, stockHolders, corp)
	}
	return mainStockHolders
}

// Returns players who are shareholders of at least one of the passed companies
// starting from the current one in play (mergemaker)
func (g *Game) setSellTradePlayers(corporations []interfaces.Corporation) []int {
	shareholders := []int{}
	index := g.currentPlayerNumber
	for _ = range g.players {
		for _, corp := range g.corporations {
			if g.players[index].Shares(corp) > 0 {
				shareholders = append(shareholders, index)
				break
			}
		}
		index++
		if index == len(g.players) {
			index = 0
		}
	}
	return shareholders
}

// Get players who have stock of the passed corporation, ordered descendently by number of stock shares
// of that corporation
func (g *Game) getStockHolders(corp interfaces.Corporation) []interfaces.Player {
	var stockHolders sortablePlayers

	for _, pl := range g.players {
		if pl.Shares(corp) > 0 {
			stockHolders.players = append(stockHolders.players, pl)
		}
	}
	stockHolders.corp = corp
	if len(stockHolders.players) > 0 {
		sort.Sort(sort.Reverse(stockHolders))
	}
	return stockHolders.players
}

// Loop stockHolders from start to get all stock holders with the same amount of shares for
// the passed corporation
func stockHoldersWithSameAmount(start int, stockHolders []interfaces.Player, corp interfaces.Corporation) []interfaces.Player {
	group := []interfaces.Player{stockHolders[start]}

	i := start + 1
	for i < len(stockHolders) && stockHolders[start].Shares(corp) == stockHolders[i].Shares(corp) {
		group = append(group, stockHolders[i])
		i++
	}
	return group
}

// Checks if two ore more corps are tied for be the acquirer in a merge
// whichs needs the merger player to decide which one would get that role.
func (g *Game) isMergeTied() bool {
	if len(g.mergeCorps["acquirer"]) > 1 {
		return true
	}
	return false
}

// TiedCorps returns all corporations that are tied in a merge
func (g *Game) TiedCorps() []interfaces.Corporation {
	corps := []interfaces.Corporation{}
	if g.isMergeTied() {
		corps = g.mergeCorps["acquirer"]
	}
	return corps
}

// UntieMerge resolves a tied merge selecting which corporation will be the acquirer,
// marking the rest as defunct
func (g *Game) UntieMerge(acquirer interfaces.Corporation) error {
	if g.state.Name() != interfaces.UntieMergeStateName {
		return errors.New(ActionNotAllowed)
	}
	for i, corp := range g.mergeCorps["acquirer"] {
		if corp == acquirer {
			g.mergeCorps["defunct"] = append(
				g.mergeCorps["defunct"],
				append(g.mergeCorps["acquirer"][:i], g.mergeCorps["acquirer"][i+1:]...)...,
			)
			g.mergeCorps["acquirer"] = []interfaces.Corporation{corp}
			g.startMerge(g.lastPlayedTile, g.mergeCorps)
			return nil
		}
	}

	return errors.New(NotAnAcquirerCorporation)
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
	g.board.SetOwner(acquirer, []interfaces.Tile{g.lastPlayedTile})
	acquirer.Grow(1)
	g.mergeCorps = map[string][]interfaces.Corporation{}
}
