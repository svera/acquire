package game

import (
	"errors"
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

	endGameCorporationSize = 41
)

type Game struct {
	board            board.Interface
	state            fsm.State
	players          []player.Interface
	corporations     [7]corporation.Interface
	tileset          tileset.Interface
	currentPlayer    int
	newCorpTiles     []tile.Interface
	mergeCorps       map[string][]corporation.Interface
	sellTradePlayers []player.Interface
	frozenPlayer     player.Interface
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
		if corp.Size() >= endGameCorporationSize {
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
		if adjacent.Owner().Type() == "unincorporated" {
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
			g.sellTradePlayers = g.shareholders(mergeCorps["defunct"])
			g.frozenPlayer = g.CurrentPlayer()
			g.setCurrentPlayer(g.nextSellTradePlayer())
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

// Returns players who are shareholders of at least one of the passed companies
func (g *Game) shareholders(corporations []corporation.Interface) []player.Interface {
	shareholders := []player.Interface{}
	for _, pl := range g.players {
		for _, corp := range g.corporations {
			if pl.Shares(corp) > 0 {
				shareholders = append(shareholders, pl)
				break
			}
		}
	}
	return shareholders
}

// Sets player currently in turn
func (g *Game) setCurrentPlayer(number int) *Game {
	g.currentPlayer = number
	return g
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
