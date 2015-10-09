package game

import "errors"

type Game struct {
	board         *Board
	status        []string
	players       []*Player
	corporations  [7]*Corporation
	tileset       *Tileset
	currentPlayer uint
}

func NewGame(board *Board, players []*Player, corporations [7]*Corporation, tileset *Tileset) (*Game, error) {
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
		corporation.setId(uint(i))
	}
	return &game, nil
}

func (g *Game) giveInitialTileset(player *Player) {
	for i := 0; i < 6; i++ {
		player.GetTile(g.tileset.Draw())
	}
}

// Placeholder function, pending implementation
func (g *Game) AreEndConditionsReached() bool {
	return true
}

// Placeholder function, pending implementation
func (g *Game) GetMainStockHolders() bool {
	return true
}
