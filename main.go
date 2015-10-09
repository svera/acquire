package main

import (
	"fmt"
	"github.com/svera/acquire/game"
)

func main() {
	var players []*game.Player
	players = append(players, game.NewPlayer("Fulanito"))
	players = append(players, game.NewPlayer("Menganito"))
	players = append(players, game.NewPlayer("Zutanito"))

	var corporations [7]*game.Corporation
	corporations[0], _ = game.NewCorporation("A", 0)
	corporations[1], _ = game.NewCorporation("B", 0)
	corporations[2], _ = game.NewCorporation("C", 1)
	corporations[3], _ = game.NewCorporation("D", 1)
	corporations[4], _ = game.NewCorporation("E", 1)
	corporations[5], _ = game.NewCorporation("F", 2)
	corporations[6], _ = game.NewCorporation("G", 2)

	board := game.NewBoard()
	tileset := game.NewTileset()

	game, _ := game.NewGame(board, players, corporations, tileset)
	fmt.Print(game)
}
