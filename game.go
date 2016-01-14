// Package game manages the flow and status of the game
package acquire

import (
	"errors"
	"github.com/svera/acquire/board"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/fsm"
	"github.com/svera/acquire/player"
	"github.com/svera/acquire/tile"
	"github.com/svera/acquire/tileset"
)

const (
	// ActionNotAllowed is an error returned when action not allowed at current state
	ActionNotAllowed = "action_not_allowed"
	// StockSharesNotBuyable is an error returned when stock shares from a corporation not on board are not buyable
	StockSharesNotBuyable = "stock_shares_not_buyable"
	// NotEnoughStockShares is an error returned when not enough stock shares of a corporation to buy
	NotEnoughStockShares = "not_enough_stock_shares"
	// TileTemporaryUnplayable is an error returned when tile temporarily unplayable
	TileTemporaryUnplayable = "tile_temporary_unplayable"
	// TilePermanentlyUnplayable is an error returned when tile permanently unplayable
	TilePermanentlyUnplayable = "tile_permanently_unplayable"
	// NotEnoughCash is an error returned when player has not enough cash to buy stock shares
	NotEnoughCash = "not_enough_cash"
	// TooManyStockSharesToBuy is an error returned when player can not buy more than 3 stock shares per turn
	TooManyStockSharesToBuy = "too_many_stock_shares_to_buy"
	// CorpNamesNotUnique is an error returned when some corporation names are repeated
	CorpNamesNotUnique = "corp_names_not_unique"
	// WrongNumberCorpsClass is an error returned when corporations classes do not fit rules
	WrongNumberCorpsClass = "wrong_number_corps_class"
	// CorporationAlreadyOnBoard is an error returned when corporation is already on board and cannot be founded
	CorporationAlreadyOnBoard = "corporation_already_on_board"
	// WrongNumberPlayers is an error returned when there must be between 3 and 6 players
	WrongNumberPlayers = "wrong_number_players"
	// NoCorporationSharesOwned is an error returned when player does not own stock shares of a certain corporation
	NoCorporationSharesOwned = "no_corporation_shares_owned"
	// NotEnoughCorporationSharesOwned is an error returned when player does not own enough stock shares of a certain corporation
	NotEnoughCorporationSharesOwned = "not_enough_corporation_shares_owned"
	// TileNotOnHand is an error returned when player does not have tile on hand
	TileNotOnHand = "tile_not_on_hand"
	// NotAnAcquirerCorporation is an error returned when corporation is not the acquirer in a merge
	NotAnAcquirerCorporation = "not_an_acquirer_corporation"
	// TradeAmountNotEven is an error returned when number of stock shares is not even in a trade
	TradeAmountNotEven = "trade_amount_not_even"

	totalCorporations      = 7
	endGameCorporationSize = 41
)

// Game stores state of game elements and provides methods to control game flow
type Game struct {
	board               board.Interface
	state               fsm.State
	players             []player.Interface
	corporations        [7]corporation.Interface
	tileset             tileset.Interface
	currentPlayerNumber int
	newCorpTiles        []tile.Interface
	mergeCorps          map[string][]corporation.Interface
	sellTradePlayers    []int
	lastPlayedTile      tile.Interface
	turn                int
	endGameClaimed      bool
	// When in sell_trade state, the current player is stored here temporary as the turn
	// is passed to all defunct corporations stockholders
	frozenPlayer int
}

// New initialises a new Acquire game
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
		board:               board,
		players:             players,
		corporations:        corporations,
		tileset:             tileset,
		currentPlayerNumber: 0,
		turn:                1,
		state:               &fsm.PlayTile{},
		endGameClaimed:      false,
	}
	for _, pl := range gm.players {
		gm.giveInitialTileset(pl)
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

// AreEndConditionsReached check if game end conditions are reached
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

// Player returns player with passed number
func (g *Game) Player(playerNumber int) player.Interface {
	return g.players[playerNumber]
}

// CurrentPlayer returns player currently in play
func (g *Game) CurrentPlayer() player.Interface {
	return g.players[g.currentPlayerNumber]
}

// PlayTile puts the given tile on board and triggers related actions
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
	g.lastPlayedTile = tl

	if merge, mergeCorps := g.board.TileMergeCorporations(tl); merge {
		g.mergeCorps = mergeCorps
		if g.isMergeTied() {
			g.state = g.state.ToUntieMerge()
		} else {
			for _, corp := range mergeCorps["defunct"] {
				g.payBonuses(corp)
			}
			g.sellTradePlayers = g.stockholders(mergeCorps["defunct"])
			g.frozenPlayer = g.currentPlayerNumber
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
// starting from the current one in play (mergemaker)
func (g *Game) stockholders(corporations []corporation.Interface) []int {
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

// Sets player currently in play
func (g *Game) setCurrentPlayer(number int) *Game {
	g.currentPlayerNumber = number
	return g
}

// FoundCorporation founds a new corporation
func (g *Game) FoundCorporation(corp corporation.Interface) error {
	if g.state.Name() != "FoundCorp" {
		return errors.New(ActionNotAllowed)
	}
	if corp.IsActive() {
		return errors.New(CorporationAlreadyOnBoard)
	}
	g.board.SetOwner(corp, g.newCorpTiles)
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

// Makes a corporation grow with the passed tiles
func (g *Game) growCorporation(corp corporation.Interface, tiles []tile.Interface) {
	g.board.SetOwner(corp, tiles)
	corp.Grow(len(tiles))
}

// Increases the number which specifies the current player
func (g *Game) nextPlayer() {
	g.currentPlayerNumber++
	if g.currentPlayerNumber == len(g.players) {
		g.currentPlayerNumber = 0
		g.turn++
	}
}

// Turn returns the current turn number
func (g *Game) Turn() int {
	return g.turn
}

// claimEndGame allows the current player to claim end game
// This can be done at any time. After announcing that the game is over,
// the player may finish the turn.
func (g *Game) ClaimEndGame() *Game {
	if g.AreEndConditionsReached() {
		g.endGameClaimed = true
	}
	return g
}

// Classification returns the players list ordered by cash,
// which is the metric used to know game's final classification
func (g *Game) Classification() []player.Interface {
	var classification []player.Interface

	cashDesc := func(pl1, pl2 player.Interface) bool {
		return pl1.Cash() > pl2.Cash()
	}

	for _, pl := range g.players {
		classification = append(classification, pl)
	}
	player.By(cashDesc).Sort(classification)
	return classification
}
