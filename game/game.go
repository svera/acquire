package game

import (
	"errors"
	"github.com/svera/acquire/game/board"
	"github.com/svera/acquire/game/corporation"
	"github.com/svera/acquire/game/player"
	"github.com/svera/acquire/game/tileset"
)

const totalCorporations = 7

type Game struct {
	board         *board.Board
	status        []string
	players       []*player.Player
	corporations  [7]*corporation.Corporation
	tileset       *tileset.Tileset
	currentPlayer uint
}

func New(
	board *board.Board, players []*player.Player, corporations [7]*corporation.Corporation, tileset *tileset.Tileset) (*Game, error) {
	if len(players) < 3 || len(players) > 6 {
		return nil, errors.New("Number of players must be between 3 and 6")
	}

	game := Game{
		board:         board,
		players:       players,
		corporations:  corporations,
		tileset:       tileset,
		currentPlayer: 0,
	}
	for _, player := range game.players {
		game.giveInitialTileset(player)
	}
	for i, corporation := range game.corporations {
		corporation.SetId(uint(i))
	}
	return &game, nil
}

func (g *Game) giveInitialTileset(player *player.Player) {
	for i := 0; i < 6; i++ {
		tile, _ := g.tileset.Draw()
		player.GetTile(tile)
	}
}

// Check if game end conditions are reached
func (g *Game) AreEndConditionsReached() bool {
	active := g.getActiveCorporations()
	if len(active) == 0 {
		return false
	}
	for _, corporation := range active {
		if corporation.Size() >= 41 {
			return true
		}
		if !corporation.IsSafe() {
			return false
		}
	}
	return true
}

// Return all corporations on the board
func (g *Game) getActiveCorporations() []*corporation.Corporation {
	active := []*corporation.Corporation{}
	for _, corporation := range g.corporations {
		if corporation.IsActive() {
			active = append(active, corporation)
		}
	}
	return active
}

// Placeholder function, pending implementation
func (g *Game) GetMainStockHolders(corporation *corporation.Corporation) map[string][]*player.Player {
	mainStockHolders := map[string][]*player.Player{"primary": {}, "secondary": {}}
	stockHolders := g.getStockHolders(corporation)

	if len(stockHolders) == 1 {
		return map[string][]*player.Player{"primary": {stockHolders[0]}, "secondary": {stockHolders[0]}}
	}
	sharesDesc := func(p1, p2 *player.Player) bool {
		return p1.Shares(corporation) > p2.Shares(corporation)
	}
	player.By(sharesDesc).Sort(stockHolders)
	if stockHolders[0].Shares(corporation) == stockHolders[1].Shares(corporation) {
		mainStockHolders["primary"] = append(mainStockHolders["primary"], stockHolders[0])
		mainStockHolders["primary"] = append(mainStockHolders["primary"], stockHolders[1])
		i := 2
		for stockHolders[0] == stockHolders[i] && i < len(stockHolders) {
			mainStockHolders["primary"] = append(mainStockHolders["primary"], stockHolders[i])
			i++
		}
		return mainStockHolders
	}
	return mainStockHolders
}

// Get players who have stock of the passed corporation
func (g *Game) getStockHolders(corporation *corporation.Corporation) []*player.Player {
	var stockHolders []*player.Player

	for _, player := range g.players {
		if player.Shares(corporation) > 0 {
			stockHolders = append(stockHolders, player)
		}
	}
	return stockHolders
}

// Returns true if a tile is permanently unplayable, that is,
// that putting it on the board would merge two or more safe corporations
func (g *Game) isTileUnplayable(tile tileset.Position) bool {
	adjacents := g.board.AdjacentCells(tile)
	for _, adjacent := range adjacents {
		safeNeighbours := 0
		boardCell := g.board.Cell(adjacent)
		if boardCell != board.CellEmpty && boardCell != board.CellOrphanTile {
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
		if boardCell == board.CellOrphanTile {
			return true
		}
	}
	return false
}
