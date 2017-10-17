// Package acquire manages the flow and status of an acquire game. It acts
// like a finite state machine (FSM), in which received inputs modify
// machine state
package acquire

import (
	"container/ring"
	"errors"
	"fmt"
	"sort"
	"time"

	"math/rand"

	"github.com/svera/acquire/board"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/fsm"
	"github.com/svera/acquire/interfaces"
	"github.com/svera/acquire/tileset"
)

const (
	totalCorporations      = 7
	endGameCorporationSize = 41
)

// Game stores state of game elements and provides methods to control game flow
type Game struct {
	board               interfaces.Board
	stateMachine        interfaces.StateMachine
	players             *ring.Ring
	corporations        [7]interfaces.Corporation
	tileset             interfaces.Tileset
	initialPlayerNumber int
	newCorpTiles        []interfaces.Tile
	mergeCorps          map[string][]interfaces.Corporation
	sellTradePlayers    []interfaces.Player
	lastPlayedTile      interfaces.Tile
	round               int
	isLastRound         bool
	// When in sell_trade state, the current player is stored here temporary as the turn
	// is passed to all defunct corporations stockholders
	frozenPlayer interfaces.Player
}

// New initialises a new Acquire game
func New(players map[int]interfaces.Player, optional Optional) (*Game, error) {
	var err error
	if len(players) < 3 || len(players) > 6 {
		return nil, errors.New(WrongNumberPlayers)
	}
	if optional, err = initOptionalParameters(optional); err == nil {
		gm := Game{
			board:        optional.Board,
			corporations: optional.Corporations,
			tileset:      optional.Tileset,
			round:        1,
			stateMachine: optional.StateMachine,
			isLastRound:  false,
		}

		gm.initPlayersRing(players)

		for i := range gm.corporations {
			gm.corporations[i].SetPricesChart(gm.setPricesChart(i))
		}

		gm.pickRandomPlayer()
		return &gm, nil
	}
	return nil, err
}

// initRing puts all players into a ring struct, so they can be looped in an
// order, circular fashion.
func (g *Game) initPlayersRing(players map[int]interfaces.Player) {
	var playerNumbers []int

	g.players = ring.New(len(players))

	for k := range players {
		playerNumbers = append(playerNumbers, k)
	}
	sort.Ints(playerNumbers)
	for _, n := range playerNumbers {
		g.giveInitialHand(players[n])
		g.players = g.players.Next()
		g.players.Value = players[n]
	}
}

func initOptionalParameters(optional Optional) (Optional, error) {
	if areCorporationsEmpty(optional.Corporations) {
		optional.Corporations = defaultCorporations()
	}
	if optional.Board == nil {
		optional.Board = board.New()
	}
	if optional.Tileset == nil {
		optional.Tileset = tileset.New()
	}
	if optional.StateMachine == nil {
		optional.StateMachine = fsm.New()
	}
	return optional, nil
}

func areCorporationsEmpty(corporations [7]interfaces.Corporation) bool {
	for i := range corporations {
		if corporations[i] == nil {
			return true
		}
	}
	return false
}

// Corporations returns an array with all seven corporations
func (g *Game) Corporations() [7]interfaces.Corporation {
	return g.corporations
}

// Initialises player hand of tiles
func (g *Game) giveInitialHand(plyr interfaces.Player) {
	for i := 0; i < 6; i++ {
		tile, _ := g.tileset.Draw()
		plyr.PickTile(tile)
	}
}

// AreEndConditionsReached checks if game end conditions are reached
func (g *Game) AreEndConditionsReached() bool {
	active := g.activeCorporations()
	safe := 0
	if len(active) == 0 {
		return false
	}
	for _, corp := range active {
		if corp.Size() >= endGameCorporationSize {
			return true
		}
		if corp.IsSafe() {
			safe++
		}
	}
	if safe == len(active) {
		return true
	}
	return false
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
// and will disappear from the board after that merge is complete, false otherwise
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
	if g.isTileTemporarilyUnplayable(tl) || g.isTilePermanentlyUnplayable(tl) {
		return false
	}
	return true
}

// Returns true if a tile is permanently unplayable, that is,
// that putting it on the board would merge two or more safe corporations
func (g *Game) isTilePermanentlyUnplayable(tl interfaces.Tile) bool {
	adjacents := g.board.AdjacentCorporations(tl.Number(), tl.Letter())
	safeNeighbours := 0
	for _, adjacent := range adjacents {
		if adjacent.IsSafe() {
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
func (g *Game) isTileTemporarilyUnplayable(tl interfaces.Tile) bool {
	if len(g.activeCorporations()) < totalCorporations {
		return false
	}
	adjacents := g.board.AdjacentCells(tl.Number(), tl.Letter())
	emptyAdjacentCells := 0
	for _, adjacent := range adjacents {
		if adjacent.Type() == interfaces.CorporationOwner {
			return false
		}
		if adjacent.Type() == interfaces.EmptyOwner {
			emptyAdjacentCells++
		}
	}
	if emptyAdjacentCells == len(adjacents) {
		return false
	}
	return true
}

// Player returns player with passed number
func (g *Game) Player(playerNumber int) interfaces.Player {
	cp := g.players.Next()
	for i := 0; i < cp.Len(); i++ {
		if cp.Value.(interfaces.Player).Number() == playerNumber {
			return cp.Value.(interfaces.Player)
		}
		cp = cp.Next()
	}
	return nil
}

// CurrentPlayer returns player currently in play
func (g *Game) CurrentPlayer() interfaces.Player {
	return g.players.Value.(interfaces.Player)
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
		g.stateMachine.ToFoundCorp()
		g.newCorpTiles = tiles
	} else if grow, tiles, corp := g.board.TileGrowCorporation(tl); grow {
		g.growCorporation(corp, tiles)
		g.stateMachine.ToBuyStock()
	} else {
		return g.putUnincorporatedTile(tl)
	}
	return nil
}

func (g *Game) checkTile(tl interfaces.Tile) error {
	if g.stateMachine.CurrentStateName() != interfaces.PlayTileStateName {
		return fmt.Errorf(ActionNotAllowed, "play_tile", g.stateMachine.CurrentStateName())
	}
	if g.isTileTemporarilyUnplayable(tl) {
		return errors.New(TileTemporarilyUnplayable)
	}
	if !g.CurrentPlayer().HasTile(tl) {
		return errors.New(TileNotOnHand)
	}
	return nil
}

func (g *Game) putUnincorporatedTile(tl interfaces.Tile) error {
	g.board.PutTile(tl)
	if g.existActiveCorporations() {
		g.stateMachine.ToBuyStock()
		return nil
	}
	return g.endTurn()
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
func (g *Game) setCurrentPlayer(pl interfaces.Player) *Game {
	for i := 0; i < g.players.Len(); i++ {
		if g.players.Value.(interfaces.Player) != pl {
			g.players = g.players.Next()
		}
	}
	return g
}

// FoundCorporation founds a new corporation
func (g *Game) FoundCorporation(corp interfaces.Corporation) error {
	if g.stateMachine.CurrentStateName() != interfaces.FoundCorpStateName {
		return fmt.Errorf(ActionNotAllowed, "found_corp", g.stateMachine.CurrentStateName())
	}
	if corp.IsActive() {
		return errors.New(CorporationAlreadyOnBoard)
	}
	g.board.SetOwner(corp, g.newCorpTiles)
	corp.Grow(len(g.newCorpTiles))
	g.newCorpTiles = []interfaces.Tile{}
	g.getFounderStockShare(g.CurrentPlayer(), corp)
	g.stateMachine.ToBuyStock()
	return nil
}

// Receive a free stock share from a recently founded corporation, if it has
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

// RemovePlayer gets the received player and removes him/her from the game
// because the player has left the game. All player's
// assets are returned to its respective origin sets. If the removed player was in
// his/her turn, turn passes to the next player.
func (g *Game) RemovePlayer(pl interfaces.Player) {
	pl.RemoveCash(pl.Cash())
	g.tileset.Add(pl.Tiles())
	for _, corp := range g.Corporations() {
		if pl.Shares(corp) > 0 {
			corp.AddStock(pl.Shares(corp))
			pl.RemoveShares(corp, pl.Shares(corp))
		}
	}

	// Is current player the one to remove? Then end his/her turn before removing
	if g.players.Value.(interfaces.Player) == pl {
		g.players = g.players.Prev()
		g.players.Unlink(1)
		g.nextPlayer()
	} else {
		search := g.players

		for i := 0; i < g.players.Len(); i++ {
			if search.Next().Value.(interfaces.Player) == pl {
				search.Unlink(1)
				break
			}
			search = search.Next()
		}
	}

	if g.players.Len() < 3 {
		g.stateMachine.ToInsufficientPlayers()
		return
	}
}

// Round returns the current round number
func (g *Game) Round() int {
	return g.round
}

// ClaimEndGame allows the current player to claim end game
// This can be done at any time. After announcing that the game is over,
// the player may finish his/her turn.
func (g *Game) ClaimEndGame() *Game {
	if g.AreEndConditionsReached() {
		g.isLastRound = true
	}
	return g
}

// Board returns game's board instance
func (g *Game) Board() interfaces.Board {
	return g.board
}

// GameStateName returns game's current state
func (g *Game) GameStateName() string {
	return g.stateMachine.CurrentStateName()
}

func (g *Game) endTurn() error {
	var err error

	if g.isLastRound {
		g.stateMachine.ToEndGame()
		return g.finish()
	}

	if err = g.drawTile(); err == nil || err.Error() == tileset.NoTilesAvailable {
		g.nextPlayer()
	}

	return err
}

// Passes the turn to the next player
func (g *Game) nextPlayer() {
	if g.stateMachine.CurrentStateName() != interfaces.PlayTileStateName {
		g.stateMachine.ToPlayTile()
	}
	if g.players.Value.(interfaces.Player).Number() > g.players.Next().Value.(interfaces.Player).Number() {
		g.round++
	}
	g.players = g.players.Next()

	if len(g.CurrentPlayer().Tiles()) == 0 {
		g.stateMachine.ToEndGame()
		g.finish()
	}

	if g.isHandUnplayable() {
		g.replaceWholeHand()
	}
}

// IsLastRound returns if the current round will be the last one or not
func (g *Game) IsLastRound() bool {
	return g.isLastRound
}

// A player takes a tile from the facedown cluster to replace
// the one he/she played. This is not done until the end of
// the turn.
func (g *Game) drawTile() error {
	var tile interfaces.Tile
	var err error
	if tile, err = g.tileset.Draw(); err != nil {
		return err
	}
	g.CurrentPlayer().PickTile(tile)

	if err = g.replaceUnplayableTiles(); err != nil {
		return err
	}

	return nil
}

// If a player has any permanently
// unplayable tiles that player discard the unplayable tiles
// and draws an equal number of replacement tiles. This can
// only be done once per turn.
func (g *Game) replaceUnplayableTiles() error {
	for _, tile := range g.CurrentPlayer().Tiles() {
		if g.isTilePermanentlyUnplayable(tile) {
			g.CurrentPlayer().DiscardTile(tile)
			if newTile, err := g.tileset.Draw(); err == nil {
				g.CurrentPlayer().PickTile(newTile)
			} else {
				return err
			}
		}
	}

	return nil
}

// This checks the (highly improbable, but possible) case in which all tiles in a
// hand are not playable and thus, the player cannot place a tile as it is forced
// to
func (g *Game) isHandUnplayable() bool {
	for _, tile := range g.CurrentPlayer().Tiles() {
		if g.IsTilePlayable(tile) {
			return false
		}
	}
	return true
}

// This must be only used if a player hand is COMPLETELY unplayable
func (g *Game) replaceWholeHand() error {
	for _, tile := range g.CurrentPlayer().Tiles() {
		g.CurrentPlayer().DiscardTile(tile)
		if newTile, err := g.tileset.Draw(); err == nil {
			g.CurrentPlayer().PickTile(newTile)
		} else {
			return err
		}
	}
	return nil
}

func defaultCorporations() [7]interfaces.Corporation {
	var corporations [7]interfaces.Corporation

	for i := 0; i < 7; i++ {
		corporations[i] = corporation.New()
	}
	return corporations
}

// Choose a random player who'll be the one to start playing
func (g *Game) pickRandomPlayer() {
	source := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(source)
	numberPlayers := g.players.Len()
	move := rn.Intn(numberPlayers - 1)
	g.players = g.players.Move(move)
	g.initialPlayerNumber = g.players.Value.(interfaces.Player).Number()
}
