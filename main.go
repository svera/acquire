package main

import (
	"fmt"
	"github.com/svera/acquire/game"
	"github.com/svera/acquire/game/board"
	"github.com/svera/acquire/game/corporation"
	"github.com/svera/acquire/game/player"
	"github.com/svera/acquire/game/tileset"
)

func main() {
	var players []*player.Player
	players = append(players, player.New("Fulanito"))
	players = append(players, player.New("Menganito"))
	players = append(players, player.New("Zutanito"))

	var corporations [7]*corporation.Corporation
	corporations[0], _ = corporation.New("A", 0)
	corporations[1], _ = corporation.New("B", 0)
	corporations[2], _ = corporation.New("C", 1)
	corporations[3], _ = corporation.New("D", 1)
	corporations[4], _ = corporation.New("E", 1)
	corporations[5], _ = corporation.New("F", 2)
	corporations[6], _ = corporation.New("G", 2)

	board := board.New()
	tileset := tileset.New()

	game, _ := game.New(board, players, corporations, tileset)
	fmt.Print(game)
}
