// Package acquire manages the flow and status of an acquire game. It acts
// like a finite state machine (FSM), in which received inputs modify
// machine state
package acquire

import (
	"errors"
	"github.com/svera/acquire/interfaces"
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

type sortablePlayers struct {
	players []interfaces.Player
	corp    interfaces.Corporation
}

func (s sortablePlayers) Len() int { return len(s.players) }
func (s sortablePlayers) Less(i, j int) bool {
	return s.players[i].Shares(s.corp) < s.players[j].Shares(s.corp)
}
func (s sortablePlayers) Swap(i, j int) { s.players[i], s.players[j] = s.players[j], s.players[i] }

// Game stores state of game elements and provides methods to control game flow
type Game struct {
	board               interfaces.Board
	state               interfaces.State
	players             []interfaces.Player
	corporations        [7]interfaces.Corporation
	tileset             interfaces.Tileset
	currentPlayerNumber int
	newCorpTiles        []interfaces.Tile
	mergeCorps          map[string][]interfaces.Corporation
	sellTradePlayers    []int
	lastPlayedTile      interfaces.Tile
	turn                int
	lastTurn            bool
	// When in sell_trade state, the current player is stored here temporary as the turn
	// is passed to all defunct corporations stockholders
	frozenPlayer int
}

// New initialises a new Acquire game
func New(
	board interfaces.Board,
	players []interfaces.Player,
	corporations [7]interfaces.Corporation,
	tileset interfaces.Tileset,
	state interfaces.State,
) (*Game, error) {
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
		state:               state,
		lastTurn:            false,
	}
	for _, pl := range gm.players {
		gm.giveInitialTileset(pl)
	}

	return &gm, nil
}

// Check that the passed corporations have unique names
func areNamesUnique(corporations [7]interfaces.Corporation) bool {
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
func isNumberOfCorpsPerClassRight(corporations [7]interfaces.Corporation) bool {
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
func (g *Game) giveInitialTileset(plyr interfaces.Player) {
	for i := 0; i < 6; i++ {
		tile, _ := g.tileset.Draw()
		plyr.PickTile(tile)
	}
}

// AreEndConditionsReached checks if game end conditions are reached
func (g *Game) AreEndConditionsReached() bool {
	active := g.activeCorporations()
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

// ActiveCorporations returns all corporations on the board
func (g *Game) activeCorporations() []interfaces.Corporation {
	return g.findCorporationsByActiveState(true)
}

func (g *Game) findCorporationsByActiveState(value bool) []interfaces.Corporation {
	result := []interfaces.Corporation{}
	for _, corp := range g.corporations {
		if corp.IsActive() == value {
			result = append(result, corp)
		}
	}
	return result
}

// IsCorporationDefunct return true if the passed corporation is in a merge process
// and will dissapear from the board after that merge is complete, false otherwise
func (g *Game) IsCorporationDefunct(corp interfaces.Corporation) bool {
	for _, defunct := range g.mergeCorps["defunct"] {
		if corp == defunct {
			return true
		}
	}
	return false
}

// IsTilePlayable returns false if the passed tile is either
// temporary or permanently unplayable, true otherwise
func (g *Game) IsTilePlayable(tl interfaces.Tile) bool {
	if g.isTileTemporaryUnplayable(tl) || g.isTilePermanentlyUnplayable(tl) {
		return false
	}
	return true
}

// Returns true if a tile is permanently unplayable, that is,
// that putting it on the board would merge two or more safe corporations
func (g *Game) isTilePermanentlyUnplayable(tl interfaces.Tile) bool {
	adjacents := g.board.AdjacentCells(tl.Number(), tl.Letter())
	safeNeighbours := 0
	for _, adjacent := range adjacents {
		if adjacent.Type() == "corporation" && adjacent.(interfaces.Corporation).IsSafe() {
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
func (g *Game) isTileTemporaryUnplayable(tl interfaces.Tile) bool {
	if len(g.activeCorporations()) < totalCorporations {
		return false
	}
	adjacents := g.board.AdjacentCells(tl.Number(), tl.Letter())
	for _, adjacent := range adjacents {
		if adjacent.Type() == "unincorporated" {
			return true
		}
	}
	return false
}

// Player returns player with passed number
func (g *Game) Player(playerNumber int) interfaces.Player {
	return g.players[playerNumber]
}

// CurrentPlayer returns player currently in play
func (g *Game) CurrentPlayer() interfaces.Player {
	return g.players[g.currentPlayerNumber]
}

// CurrentPlayerNumber returns the number of the player currently in play
func (g *Game) CurrentPlayerNumber() int {
	return g.currentPlayerNumber
}

// PlayTile puts the given tile on board and triggers related actions
func (g *Game) PlayTile(tl interfaces.Tile) error {
	if err := g.checkTile(tl); err != nil {
		return err
	}

	g.CurrentPlayer().DiscardTile(tl)
	g.lastPlayedTile = tl

	if merge, mergeCorps := g.board.TileMergeCorporations(tl); merge {
		g.startMerge(tl, mergeCorps)
	} else if corpFounded, tiles := g.board.TileFoundCorporation(tl); corpFounded {
		g.board.PutTile(tl)
		g.state = g.state.ToFoundCorp()
		g.newCorpTiles = tiles
	} else if grow, tiles, corp := g.board.TileGrowCorporation(tl); grow {
		g.growCorporation(corp, tiles)
		g.state = g.state.ToBuyStock()
	} else {
		if err := g.putUnincorporatedTile(tl); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) checkTile(tl interfaces.Tile) error {
	if g.state.Name() != interfaces.PlayTileStateName {
		return errors.New(ActionNotAllowed)
	}
	if g.isTileTemporaryUnplayable(tl) {
		return errors.New(TileTemporaryUnplayable)
	}
	if !g.CurrentPlayer().HasTile(tl) {
		return errors.New(TileNotOnHand)
	}
	return nil
}

func (g *Game) putUnincorporatedTile(tl interfaces.Tile) error {
	g.board.PutTile(tl)
	if g.existActiveCorporations() {
		g.state = g.state.ToBuyStock()
	} else {
		if err := g.drawTile(); err != nil {
			return err
		}
		g.nextPlayer()
	}
	return nil
}

func (g *Game) existActiveCorporations() bool {
	for _, corp := range g.corporations {
		if corp.IsActive() {
			return true
		}
	}
	return false
}

// Sets player currently in play
func (g *Game) setCurrentPlayer(number int) *Game {
	g.currentPlayerNumber = number
	return g
}

// FoundCorporation founds a new corporation
func (g *Game) FoundCorporation(corp interfaces.Corporation) error {
	if g.state.Name() != interfaces.FoundCorpStateName {
		return errors.New(ActionNotAllowed)
	}
	if corp.IsActive() {
		return errors.New(CorporationAlreadyOnBoard)
	}
	g.board.SetOwner(corp, g.newCorpTiles)
	corp.Grow(len(g.newCorpTiles))
	g.newCorpTiles = []interfaces.Tile{}
	g.getFounderStockShare(g.CurrentPlayer(), corp)
	g.state = g.state.ToBuyStock()
	return nil
}

// Receive a free stock share from a rencently found corporation, if it has
// remaining shares available
// TODO this should trigger an event warning that no founder stock share will be given
// of the founded corporation has no stock shares left
func (g *Game) getFounderStockShare(pl interfaces.Player, corp interfaces.Corporation) {
	if corp.Stock() > 0 {
		corp.RemoveStock(1)
		pl.AddShares(corp, 1)
	}
}

// Makes a corporation grow with the passed tiles
func (g *Game) growCorporation(corp interfaces.Corporation, tiles []interfaces.Tile) {
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

// ClaimEndGame allows the current player to claim end game
// This can be done at any time. After announcing that the game is over,
// the player may finish the turn.
func (g *Game) ClaimEndGame() *Game {
	if g.AreEndConditionsReached() {
		g.lastTurn = true
	}
	return g
}

// LastTurn returns if the current turn will be the last one or not
func (g *Game) LastTurn() bool {
	return g.lastTurn
}

// Classification returns the players list ordered by cash,
// which is the metric used to know game's final classification
/*
func (g *Game) Classification() []interfaces.Player {
	var classification []interfaces.Player

	cashDesc := func(pl1, pl2 interfaces.Player) bool {
		return pl1.Cash() > pl2.Cash()
	}

	for _, pl := range g.players {
		classification = append(classification, pl)
	}
	player.By(cashDesc).Sort(classification)
	return classification
}
*/

// Board returns game's board instance
func (g *Game) Board() interfaces.Board {
	return g.board
}

// GameStateName returns game's current state
func (g *Game) GameStateName() string {
	return g.state.Name()
}
