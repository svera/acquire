package game

import (
	"errors"
	"fmt"
	"github.com/svera/acquire/board"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/fsm"
	"github.com/svera/acquire/player"
	"github.com/svera/acquire/tile"
	"github.com/svera/acquire/tileset"
)

const totalCorporations = 7
const (
	ActionNotAllowed                = "action_not_allowed"
	StockSharesNotBuyable           = "stock_shares_not_buyable"
	NotEnoughStockShares            = "not_enough_stock_shares"
	TileTemporaryUnplayable         = "tile_temporary_unplayable"
	TilePermanentlyUnplayable       = "tile_permanently_unplayable"
	NotEnoughCash                   = "not_enough_cash"
	TooManyStockSharesToBuy         = "too_many_stock_shares_to_buy"
	CorpNamesNotUnique              = "corp_names_not_unique"
	WrongNumberCorpsClass           = "wrong_number_corps_class"
	CorporationAlreadyOnBoard       = "corporation_already_on_board"
	WrongNumberPlayers              = "wrong_number_players"
	NoCorporationSharesOwned        = "no_corporation_shares_owned"
	NotEnoughCorporationSharesOwned = "not_enough_corporation_shares_owned"
	TileNotOnHand                   = "tile_not_on_hand"
)

type Game struct {
	board         board.Interface
	state         fsm.State
	players       []player.Interface
	corporations  [7]corporation.Interface
	tileset       tileset.Interface
	currentPlayer int
	newCorpTiles  []tile.Interface
	mergeCorps    map[string][]corporation.Interface
}

func New(
	board board.Interface, players []player.Interface, corporations [7]corporation.Interface, tileset tileset.Interface) (*Game, error) {
	if len(players) < 3 || len(players) > 6 {
		return nil, errors.New(WrongNumberPlayers)
	}
	if !areNamesUnique(corporations) {
		return nil, errors.New(CorpNamesNotUnique)
	}
	if !isNumberOfCorpsPerClassRight(corporations) {
		return nil, errors.New(WrongNumberCorpsClass)
	}
	gm := Game{
		board:         board,
		players:       players,
		corporations:  corporations,
		tileset:       tileset,
		currentPlayer: 0,
		state:         &fsm.PlayTile{},
	}
	for _, plyr := range gm.players {
		gm.giveInitialTileset(plyr)
	}

	return &gm, nil
}

// Check that the passed corporations have unique names
func areNamesUnique(corporations [7]corporation.Interface) bool {
	for i, corp1 := range corporations {
		if i < len(corporations)-1 {
			for _, corp2 := range corporations[i+1:] {
				if corp1.Name() == corp2.Name() {
					return false
				}
			}
		}
	}
	return true
}

// Check that the number of corporations per class is right
func isNumberOfCorpsPerClassRight(corporations [7]corporation.Interface) bool {
	corpsPerClass := [3]int{0, 0, 0}
	for _, corp := range corporations {
		corpsPerClass[corp.Class()]++
	}
	if corpsPerClass[0] != 2 || corpsPerClass[1] != 3 || corpsPerClass[2] != 2 {
		return false
	}
	return true
}

// Initialises player hand of tiles
func (g *Game) giveInitialTileset(plyr player.Interface) {
	for i := 0; i < 6; i++ {
		tile, _ := g.tileset.Draw()
		plyr.PickTile(tile)
	}
}

// Check if game end conditions are reached
func (g *Game) AreEndConditionsReached() bool {
	active := g.getActiveCorporations()
	if len(active) == 0 {
		return false
	}
	for _, corp := range active {
		if corp.Size() >= 41 {
			return true
		}
		if !corp.IsSafe() {
			return false
		}
	}
	return true
}

// Returns all corporations on the board
func (g *Game) getActiveCorporations() []corporation.Interface {
	active := []corporation.Interface{}
	for _, corp := range g.corporations {
		if corp.IsActive() {
			active = append(active, corp)
		}
	}
	return active
}

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

// Returns true if a tile is permanently unplayable, that is,
// that putting it on the board would merge two or more safe corporations
func (g *Game) isTileUnplayable(tl tile.Interface) bool {
	adjacents := g.board.AdjacentCells(tl)
	safeNeighbours := 0
	for _, adjacent := range adjacents {
		if adjacent.Owner().Type() == "corporation" && adjacent.Owner().(corporation.Interface).IsSafe() {
			safeNeighbours++
		}
		if safeNeighbours == 2 {
			return true
		}
	}
	return false
}

// Returns true if a tile is temporarily unplayable, that is,
// that putting it on the board would create an 8th corporation
func (g *Game) isTileTemporaryUnplayable(tl tile.Interface) bool {
	if len(g.getActiveCorporations()) < totalCorporations {
		return false
	}
	adjacents := g.board.AdjacentCells(tl)
	for _, adjacent := range adjacents {
		if adjacent.Owner().Type() == "orphan" {
			return true
		}
	}
	return false
}

// Returns player currently in turn
func (g *Game) CurrentPlayer() player.Interface {
	return g.players[g.currentPlayer]
}

// Puts the given tile on board and triggers related actions
func (g *Game) PlayTile(tl tile.Interface) error {
	if g.state.Name() != "PlayTile" {
		return errors.New(ActionNotAllowed)
	}
	if g.isTileTemporaryUnplayable(tl) {
		return errors.New(TileTemporaryUnplayable)
	}
	if !g.CurrentPlayer().HasTile(tl) {
		return errors.New(TileNotOnHand)
	}

	g.CurrentPlayer().DiscardTile(tl)

	if merge, mergeCorps := g.board.TileMergeCorporations(tl); merge {
		g.mergeCorps = mergeCorps
		if g.isMergeTied() {
			g.state = g.state.ToUntieMerge()
		} else {
			g.payMergeBonuses()
			g.state = g.state.ToSellTrade()
		}
	} else if found, tiles := g.board.TileFoundCorporation(tl); found {
		g.state = g.state.ToFoundCorp()
		g.newCorpTiles = tiles
	} else if grow, tiles, corp := g.board.TileGrowCorporation(tl); grow {
		g.growCorporation(corp, tiles)
		g.state = g.state.ToBuyStock()
	} else {
		g.board.PutTile(tl)
		g.state = g.state.ToBuyStock()
	}
	return nil
}

// Checks if two ore more corps are tied for be the acquirer in a merge
// whichs needs the merger player to decide which one would get that role.
func (g *Game) isMergeTied() bool {
	if len(g.mergeCorps["acquirer"]) > 1 {
		return true
	}
	return false
}

// Calculates and returns bonus amounts to be paid to owners of stock of a
// defunct corporation
func (g *Game) payMergeBonuses() {
	for _, corp := range g.mergeCorps["defunct"] {
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
			fmt.Println(minorityStockHolder.Name())
			minorityStockHolder.AddCash(corp.MinorityBonus() / numberMinorityHolders)
		}
	}
}

// TODO
func (g *Game) SellTrade(pl player.Interface, sell map[corporation.Interface]int, trade map[corporation.Interface]int) error {
	if err := g.checkSellTrade(pl, sell, trade); err != nil {
		return err
	}
	for corp, amount := range sell {
		g.sell(pl, corp, amount)
	}

	for corp, amount := range trade {
		g.trade(pl, corp, amount)
	}
	return nil
}

func (g *Game) sell(pl player.Interface, corp corporation.Interface, amount int) {
	corp.AddStock(amount)
	pl.
		RemoveShares(corp, amount).
		AddCash(corp.StockPrice() * amount)
}

func (g *Game) trade(pl player.Interface, corp corporation.Interface, amount int) {
	acquirer := g.mergeCorps["acquirer"][0]
	amountSharesAcquiringCorp := amount / 2
	corp.AddStock(amount)
	acquirer.RemoveStock(amountSharesAcquiringCorp)
	pl.
		RemoveShares(corp, amount).
		AddShares(acquirer, amountSharesAcquiringCorp)
}

func (g *Game) checkSellTrade(pl player.Interface, sell map[corporation.Interface]int, trade map[corporation.Interface]int) error {
	if g.state.Name() != "SellTrade" {
		return errors.New(ActionNotAllowed)
	}
	for corp, amount := range sell {
		if amount > 0 && pl.Shares(corp) == 0 {
			return errors.New(NoCorporationSharesOwned)
		}
		if pl.Shares(corp) < amount {
			return errors.New(NotEnoughCorporationSharesOwned)
		}
	}
	for corp, amount := range trade {
		if amount > 0 && pl.Shares(corp) == 0 {
			return errors.New(NoCorporationSharesOwned)
		}
		if corp.Stock() < (amount / 2) {
			return errors.New(NotEnoughStockShares)
		}
		if pl.Shares(corp) < amount {
			return errors.New(NotEnoughCorporationSharesOwned)
		}
	}
	return nil
}

func (g *Game) FoundCorporation(corp corporation.Interface) error {
	if g.state.Name() != "FoundCorp" {
		return errors.New(ActionNotAllowed)
	}
	if corp.IsActive() {
		return errors.New(CorporationAlreadyOnBoard)
	}
	g.board.SetTiles(corp, g.newCorpTiles)
	corp.Grow(len(g.newCorpTiles))
	g.newCorpTiles = []tile.Interface{}
	g.getFounderStockShare(g.CurrentPlayer(), corp)
	g.state = g.state.ToBuyStock()
	return nil
}

// Receive a free stock share from a rencently found corporation, if it has
// remaining shares available
// TODO this should trigger an event warning that no founder stock share will be given
// of the founded corporation has no stock shares left
func (g *Game) getFounderStockShare(pl player.Interface, corp corporation.Interface) {
	if corp.Stock() > 0 {
		corp.RemoveStock(1)
		pl.AddShares(corp, 1)
	}
}

func (g *Game) growCorporation(corp corporation.Interface, tiles []tile.Interface) {
	g.board.SetTiles(corp, tiles)
	corp.Grow(len(tiles))
}

// Increases the number which specifies the current player
func (g *Game) nextPlayer() {
	g.currentPlayer++
	if g.currentPlayer == len(g.players) {
		g.currentPlayer = 0
	}
}
