package game

import (
	"errors"
	"github.com/svera/acquire/board"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/fsm"
	"github.com/svera/acquire/player"
	"github.com/svera/acquire/tileset"
)

const totalCorporations = 7
const (
	ActionNotAllowed          = "action_not_allowed"
	WrongNumberPlayers        = "wrong_number_players"
	StockSharesNotBuyable     = "stock_shares_not_buyable"
	NotEnoughStockShares      = "not_enough_stock_shares"
	TileTemporaryUnplayable   = "tile_temporary_unplayable"
	TilePermanentlyUnplayable = "tile_permanently_unplayable"
	NotEnoughCash             = "not_enough_cash"
	TooManyStockSharesToBuy   = "too_many_stock_shares_to_buy"
	CorpIdNotUnique           = "corp_id_not_unique"
	WrongNumberCorpsClass     = "wrong_number_corps_class"
)

type Game struct {
	board         board.Interface
	state         fsm.State
	players       []player.Interface
	corporations  [7]corporation.Interface
	tileset       tileset.Interface
	currentPlayer int
}

func New(
	board board.Interface, players []player.Interface, corporations [7]corporation.Interface, tileset tileset.Interface) (*Game, error) {
	if len(players) < 3 || len(players) > 6 {
		return nil, errors.New(WrongNumberPlayers)
	}
	if !areIdsUnique(corporations) {
		return nil, errors.New(CorpIdNotUnique)
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

// Check that the passed corporations have unique IDs
func areIdsUnique(corporations [7]corporation.Interface) bool {
	for i, corp1 := range corporations {
		if i < len(corporations)-1 {
			for _, corp2 := range corporations[i+1:] {
				if corp1.Id() == corp2.Id() {
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

// Calculates and returns bonus amounts to be paid to owners of stock of a
// defunct corporation
func (g *Game) PayBonusesForDefunctCorporation(corp corporation.Interface) {
	stockHolders := g.GetMainStockHolders(corp)
	numberMajorityHolders := len(stockHolders["majority"])
	numberMinorityHolders := len(stockHolders["minority"])

	for _, majorityStockHolder := range stockHolders["majority"] {
		majorityStockHolder.ReceiveBonus(corp.MajorityBonus() / int(numberMajorityHolders))
	}
	for _, minorityStockHolder := range stockHolders["minority"] {
		minorityStockHolder.ReceiveBonus(corp.MinorityBonus() / int(numberMinorityHolders))
	}
}

// Taken from the game rules:
// "If only one player owns stock in the defunct corporation, that player gets both bonuses. If there's
// a tie for majority stockholder, add the majority and minority bonuses and divide evenly (the minority
// shareholder gets no bonus. If there's a tie for minority stockholder, split the minority bonus among
// the tied players"
func (g *Game) GetMainStockHolders(corp corporation.Interface) map[string][]player.ShareInterface {
	mainStockHolders := map[string][]player.ShareInterface{"majority": {}, "minority": {}}
	stockHolders := g.getStockHolders(corp)

	if len(stockHolders) == 1 {
		return map[string][]player.ShareInterface{
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
func stockHoldersWithSameAmount(start int, stockHolders []player.ShareInterface, corp corporation.Interface) []player.ShareInterface {
	group := []player.ShareInterface{stockHolders[start]}

	i := start + 1
	for i < len(stockHolders) && stockHolders[start].Shares(corp) == stockHolders[i].Shares(corp) {
		group = append(group, stockHolders[i])
		i++
	}
	return group
}

// Get players who have stock of the passed corporation, ordered descendently by number of stock shares
// of that corporation
func (g *Game) getStockHolders(corp corporation.Interface) []player.ShareInterface {
	var stockHolders []player.ShareInterface
	sharesDesc := func(p1, p2 player.ShareInterface) bool {
		return p1.Shares(corp) > p2.Shares(corp)
	}

	for _, player := range g.players {
		if player.Shares(corp) > 0 {
			stockHolders = append(stockHolders, player)
		}
	}
	player.By(sharesDesc).Sort(stockHolders)
	return stockHolders
}

// Returns true if a tile is permanently unplayable, that is,
// that putting it on the board would merge two or more safe corporations
func (g *Game) isTileUnplayable(tile tileset.Position) bool {
	adjacents := g.board.AdjacentCells(tile)
	safeNeighbours := 0
	for _, adjacent := range adjacents {
		boardCell := g.board.Cell(adjacent)
		if boardCell != board.Empty && boardCell != board.OrphanTile {
			if g.corporations[boardCell].IsSafe() {
				safeNeighbours++
			}
		}
		if safeNeighbours == 2 {
			return true
		}
	}
	return false
}

// Returns true if a tile is temporarily unplayable, that is,
// that putting it on the board would create an 8th corporation
func (g *Game) isTileTemporaryUnplayable(tile tileset.Position) bool {
	if len(g.getActiveCorporations()) < totalCorporations {
		return false
	}
	adjacents := g.board.AdjacentCells(tile)
	for _, adjacent := range adjacents {
		boardCell := g.board.Cell(adjacent)
		if boardCell == board.OrphanTile {
			return true
		}
	}
	return false
}

// Returns player currently in turn
func (g *Game) CurrentPlayer() player.Interface {
	return g.players[g.currentPlayer]
}

func (g *Game) PlayTile(tile tileset.Position) error {
	if _, ok := g.state.(*fsm.PlayTile); !ok {
		return errors.New(ActionNotAllowed)
	}
	if g.isTileTemporaryUnplayable(tile) {
		return errors.New(TileTemporaryUnplayable)
	}
	if err := g.CurrentPlayer().DiscardTile(tile); err != nil {
		return err
	}
	/*
		if merge, tiles := g.board.TileMergeCorporations(tile); merge {
			// move state machine status
		} else if found, tiles := g.board.TileFoundCorporation(tile); found {

		} else */if grow, tiles, corporationId := g.board.TileGrowCorporation(tile); grow {
		g.growCorporation(g.corporations[corporationId], tiles)
	} else {
		g.board.PutTile(tile)
		g.state.ToBuyStock()
	}
	return nil
}

func (g *Game) growCorporation(corp corporation.Interface, tiles []tileset.Position) {
	g.board.SetTiles(corp, tiles)
	corp.AddTiles(tiles)
	g.state.ToBuyStock()
}

// Increases the number which specifies the current player
func (g *Game) nextPlayer() {
	g.currentPlayer++
	if g.currentPlayer == len(g.players) {
		g.currentPlayer = 0
	}
}

// Buys stock from corporations
func (g *Game) BuyStock(buys map[int]int) error {
	if _, ok := g.state.(*fsm.BuyStock); !ok {
		return errors.New(ActionNotAllowed)
	}

	if err := g.checkBuy(buys); err != nil {
		return err
	}

	for corporationId, amount := range buys {
		corp := g.corporations[corporationId]
		g.CurrentPlayer().Buy(corp, amount)
	}

	return g.drawTile()
}

func (g *Game) checkBuy(buys map[int]int) error {
	var totalStock, totalPrice int = 0, 0
	for corporationId, amount := range buys {
		corp := g.corporations[corporationId]
		if corp.Size() == 0 {
			return errors.New(StockSharesNotBuyable)
		}
		if amount > corp.Stock() {
			return errors.New(NotEnoughStockShares)
		}
		totalStock += amount
		totalPrice += corp.StockPrice() * amount
	}
	if totalStock > 3 {
		return errors.New(TooManyStockSharesToBuy)
	}

	if totalPrice > g.CurrentPlayer().Cash() {
		return errors.New(NotEnoughCash)
	}
	return nil
}

// A player takes a tile from the facedown cluster to replace
// the one he/she played. This is not done until the end of
// the turn.
func (g *Game) drawTile() error {
	var tile tileset.Position
	var err error
	if tile, err = g.tileset.Draw(); err != nil {
		return err
	}
	g.CurrentPlayer().PickTile(tile)

	if err = g.replaceUnplayableTiles(); err != nil {
		return err
	}

	g.state.ToPlayTile()
	g.nextPlayer()
	return nil
}

// if a player has any permanently
// unplayable tiles that player discard the unplayable tiles
// and draws an equal number of replacement tiles. This can
// only be done once per turn.
func (g *Game) replaceUnplayableTiles() error {
	for _, tile := range g.CurrentPlayer().Tiles() {
		if g.isTileUnplayable(tile) {

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
