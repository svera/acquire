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

func (g *Game) PayBonusesForDefunctCorporation(c *corporation.Corporation) {
	stockHolders := g.GetMainStockHolders(c)
	numberMajorityHolders := len(stockHolders["majority"])
	numberMinorityHolders := len(stockHolders["minority"])

	for _, majorityStockHolder := range stockHolders["majority"] {
		majorityStockHolder.ReceiveBonus(c.MajorityBonus() / uint(numberMajorityHolders))
	}
	for _, minorityStockHolder := range stockHolders["minority"] {
		minorityStockHolder.ReceiveBonus(c.MinorityBonus() / uint(numberMinorityHolders))
	}
}

// Taken from the game rules:
// "If only one player owns stock in the defunct corporation, that player gets both bonuses. If there's
// a tie for majority stockholder, add the majority and minority bonuses and divide evenly (the minority
// shareholder gets no bonus. If there's a tie for minority stockholder, split the minority bonus among
// the tied players"
func (g *Game) GetMainStockHolders(corporation *corporation.Corporation) map[string][]*player.Player {
	mainStockHolders := map[string][]*player.Player{"majority": {}, "minority": {}}
	stockHolders := g.getStockHolders(corporation)

	if len(stockHolders) == 1 {
		return map[string][]*player.Player{
			"majority": {stockHolders[0]},
			"minority": {stockHolders[0]},
		}
	}

	mainStockHolders["majority"] = stockHoldersWithSameAmount(0, stockHolders, corporation)
	if len(mainStockHolders["majority"]) > 1 {
		return mainStockHolders
	}
	mainStockHolders["minority"] = stockHoldersWithSameAmount(1, stockHolders, corporation)
	return mainStockHolders
}

// Loop stockHolders from start to get all stock holders with the same amount of shares for
// the passed corporation
func stockHoldersWithSameAmount(start int, stockHolders []*player.Player, corporation *corporation.Corporation) []*player.Player {
	group := []*player.Player{stockHolders[start]}

	i := start + 1
	for i < len(stockHolders) && stockHolders[start].Shares(corporation) == stockHolders[i].Shares(corporation) {
		group = append(group, stockHolders[i])
		i++
	}
	return group
}

// Get players who have stock of the passed corporation, ordered descendently by number of stock shares
// of that corporation
func (g *Game) getStockHolders(corporation *corporation.Corporation) []*player.Player {
	var stockHolders []*player.Player
	sharesDesc := func(p1, p2 *player.Player) bool {
		return p1.Shares(corporation) > p2.Shares(corporation)
	}

	for _, player := range g.players {
		if player.Shares(corporation) > 0 {
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
