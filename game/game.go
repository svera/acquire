package game

import (
	"errors"
	"github.com/svera/acquire/game/board"
	"github.com/svera/acquire/game/corporation"
	"github.com/svera/acquire/game/player"
	"github.com/svera/acquire/game/tileset"
)

type Game struct {
	board         *board.Board
	status        []string
	players       []*player.Player
	corporations  [7]*corporation.Corporation
	tileset       *tileset.Tileset
	currentPlayer uint
}

func New(board *board.Board, players []*player.Player, corporations [7]*corporation.Corporation, tileset *tileset.Tileset) (*Game, error) {
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
		corporation.SetId(uint(i) + 1)
	}
	return &game, nil
}

func (g *Game) giveInitialTileset(player *player.Player) {
	for i := 0; i < 6; i++ {
		player.GetTile(g.tileset.Draw())
	}
}

// Check if game end conditions are reached
func (g *Game) AreEndConditionsReached() bool {
	for _, corporation := range g.corporations {
		if corporation.Size() >= 41 {
			return true
		}
		if corporation.IsSafe() == false {
			return false
		}
	}
	return true
}

// Placeholder function, pending implementation
func (g *Game) GetMainStockHolders() bool {
	return true
}

// Placeholder function, pending implementation
func (g *Game) isTilePlayable() bool {
	return true
}
