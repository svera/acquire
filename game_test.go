package game

import (
	"github.com/svera/acquire/board"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/fsm"
	"github.com/svera/acquire/player"
	"github.com/svera/acquire/tile"
	"github.com/svera/acquire/tileset"
	"testing"
)

func TestNewGameWrongNumberPlayers(t *testing.T) {
	players, corporations, board, tileset := setup()
	players = players[:1]

	if _, err := New(board, players, corporations, tileset); err.Error() != WrongNumberPlayers {
		t.Errorf("Game must not be created with less than 3 players, got %d", len(players))
	}
}

func TestNewGameNotUniqueCorpNames(t *testing.T) {
	players, corporations, board, tileset := setup()

	corporations[0] = corporation.NewStub("A", 0)
	corporations[1] = corporation.NewStub("A", 0)

	if _, err := New(board, players, corporations, tileset); err.Error() != CorpNameNotUnique {
		t.Errorf("Corporations must have unique names, expecting %s error, got %s", CorpNameNotUnique, err.Error())
	}
}

func TestNewGameWrongNumberOfCorpsPerClass(t *testing.T) {
	players, corporations, board, tileset := setup()

	corporations[2] = corporation.NewStub("C", 0)

	if _, err := New(board, players, corporations, tileset); err.Error() != WrongNumberCorpsClass {
		t.Errorf("Game must catch wrong number of corporations per class")
	}
}

func TestNewGameInitsPlayersTilesets(t *testing.T) {
	players, corporations, board, tileset := setup()
	New(board, players, corporations, tileset)

	for i, player := range players {
		if len(player.Tiles()) != 6 {
			t.Errorf("Players must have 6 tiles at the beginning, player %d got %d", i, len(player.Tiles()))
		}
	}
}

func TestAreEndConditionsReached(t *testing.T) {
	players, corporations, board, tileset := setup()
	game, _ := New(board, players, corporations, tileset)

	if game.AreEndConditionsReached() {
		t.Errorf("End game conditions not reached (no active corporations) but detected as it were")
	}

	corporations[0].(*corporation.Stub).SetSize(41)

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (a corporation bigger than 40 tiles) but not detected")
	}

	corporations[0].(*corporation.Stub).SetSize(11)

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (all active corporations safe) but not detected")
	}

	corporations[0].(*corporation.Stub).SetSize(11)
	corporations[1].(*corporation.Stub).SetSize(2)

	if game.AreEndConditionsReached() {
		t.Errorf("End game conditions not reached but detected as it were")
	}

}

func TestGetMainStockHolders(t *testing.T) {
	players, corporations, board, tileset := setup()

	players[0].(*player.Stub).SetShares(corporations[0], 8)

	game, _ := New(board, players, corporations, tileset)
	stockHolders := game.GetMainStockHolders(corporations[0])
	expectedStockHolders := map[string][]player.ShareInterface{
		"majority": {players[0]},
		"minority": {players[0]},
	}
	if !slicesSameContent(stockHolders["majority"], expectedStockHolders["majority"]) ||
		!slicesSameContent(stockHolders["minority"], expectedStockHolders["minority"]) {
		t.Errorf(
			"If there's just one player with stock in a defunct corporation, " +
				"he/she must get both majority and minority bonuses",
		)
	}

	players[1].(*player.Stub).SetShares(corporations[0], 5)

	stockHolders = game.GetMainStockHolders(corporations[0])
	expectedStockHolders = map[string][]player.ShareInterface{
		"majority": {players[0]},
		"minority": {players[1]},
	}
	if !slicesSameContent(stockHolders["majority"], expectedStockHolders["majority"]) ||
		!slicesSameContent(stockHolders["minority"], expectedStockHolders["minority"]) {
		t.Errorf(
			"Wrong main stock holders",
		)
	}

	players[1].(*player.Stub).SetShares(corporations[0], 8)
	players[2].(*player.Stub).SetShares(corporations[0], 5)

	stockHolders = game.GetMainStockHolders(corporations[0])
	expectedStockHolders = map[string][]player.ShareInterface{
		"majority": {players[0], players[1]},
		"minority": {},
	}
	if !slicesSameContent(stockHolders["majority"], expectedStockHolders["majority"]) ||
		!slicesSameContent(stockHolders["minority"], expectedStockHolders["minority"]) {
		t.Errorf(
			"If there are two or more majority stock holders in a defunct corporation, " +
				"the majority bonus must be splitted between them (no minority bonus given)",
		)
	}

	players[1].(*player.Stub).SetShares(corporations[0], 5)
	players[2].(*player.Stub).SetShares(corporations[0], 5)

	stockHolders = game.GetMainStockHolders(corporations[0])
	expectedStockHolders = map[string][]player.ShareInterface{
		"majority": {players[0]},
		"minority": {players[1], players[2]},
	}
	if !slicesSameContent(stockHolders["majority"], expectedStockHolders["majority"]) ||
		!slicesSameContent(stockHolders["minority"], expectedStockHolders["minority"]) {
		t.Errorf(
			"If there are two or more minority stock holders in a defunct corporation, " +
				"the minority bonus must be splitted between them",
		)
	}

}

func TestPlayTileFoundCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	tileToPlay := tile.New(6, "E", tile.Orphan{})
	bd.PutTile(tile.New(5, "E", tile.Orphan{}))

	game, _ := New(bd, players, corporations, ts)
	playerTiles := players[0].Tiles()
	players[0].DiscardTile(playerTiles[0])
	players[0].PickTile(tileToPlay)
	game.PlayTile(tileToPlay)

	if game.state.Name() != "FoundCorp" {
		t.Errorf("Game must be in state FoundCorp, got %s", game.state.Name())
	}
}

func TestFoundCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	game, _ := New(bd, players, corporations, ts)
	if err := game.FoundCorporation(corporations[0]); err == nil {
		t.Errorf("Game in a state different than FoundCorp must not execute FoundCorporation()")
	}
	game.state = &fsm.FoundCorp{}
	newCorpTiles := []tile.Interface{
		tile.New(5, "E", tile.Orphan{}),
		tile.New(6, "E", tile.Orphan{}),
	}
	game.newCorpTiles = newCorpTiles
	game.FoundCorporation(corporations[0])
	if game.state.Name() != "BuyStock" {
		t.Errorf("Game must be in state BuyStock, got %s", game.state.Name())
	}
	if players[0].Shares(corporations[0]) != 1 {
		t.Errorf("Player must have 1 share of corporation stock, got %d", players[0].Shares(corporations[0]))
	}
	if corporations[0].Size() != 2 {
		t.Errorf("Corporation must have 2 tiles, got %d", corporations[0].Size())
	}
	if game.board.Cell(newCorpTiles[0].Number(), newCorpTiles[0].Letter()).Content() != corporations[0] || game.board.Cell(newCorpTiles[1].Number(), newCorpTiles[1].Letter()).Content() != corporations[0] {
		t.Errorf("Corporation tiles are not set on board")
	}
}

func TestPlayTileGrowCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	tileToPlay := tile.New(6, "E", tile.Orphan{})
	corpTiles := []tile.Interface{
		tile.New(7, "E", corporations[0]),
		tile.New(8, "E", corporations[0]),
	}
	corporations[0].Grow(len(corpTiles))
	bd.SetTiles(corporations[0], corpTiles)
	bd.PutTile(tile.New(5, "E", tile.Orphan{}))

	game, _ := New(bd, players, corporations, ts)
	playerTiles := players[0].Tiles()
	players[0].DiscardTile(playerTiles[0])
	players[0].PickTile(tileToPlay)
	game.PlayTile(tileToPlay)

	expectedCorpSize := 4

	if game.state.Name() != "BuyStock" {
		t.Errorf("Game must be in state BuyStock, got %s", game.state.Name())
	}
	if corporations[0].Size() != expectedCorpSize {
		t.Errorf("Corporation size must be %d, got %d", expectedCorpSize, corporations[0].Size())
	}
}

func TestBuyStock(t *testing.T) {
	players, corporations, bd, ts := setup()
	corporations[0].Grow(2)
	buys := map[int]int{0: 2}
	var expectedAvailableStock int = 23
	var expectedPlayerStock int = 2
	game, _ := New(bd, players, corporations, ts)
	game.state = &fsm.BuyStock{}
	game.BuyStock(buys)

	if corporations[0].Stock() != expectedAvailableStock {
		t.Errorf("Corporation stock shares have not decreased, must be %d, got %d", expectedAvailableStock, corporations[0].Stock())
	}
	if players[0].Shares(corporations[0]) != expectedPlayerStock {
		t.Errorf("Player stock shares have not increased, must be %d, got %d", expectedPlayerStock, players[0].Shares(corporations[0]))
	}
}

func TestBuyStockWithNotEnoughCash(t *testing.T) {
	players, corporations, bd, ts := setup()
	players[0].(*player.Stub).SetCash(100)

	corporations[0].Grow(2)

	buys := map[int]int{0: 2}
	game, _ := New(bd, players, corporations, ts)
	err := game.BuyStock(buys)
	if err == nil {
		t.Errorf("Trying to buy stock shares without enough money must throw error")
	}
}

// Testing that if player has an permanently unplayable tile, this is exchanged:
// In the following example, tile 6D is unplayable because it would merge safe
// corporations 0 and 1
//
//    5 6  7 8
// D [0]><[1]
func TestDrawTile(t *testing.T) {
	players, corporations, bd, ts := setup()
	corporations[0].(*corporation.Stub).SetSize(11)
	corporations[1].(*corporation.Stub).SetSize(15)
	unplayableTile := tile.New(6, "D", tile.Orphan{})
	bd.SetTiles(corporations[0], []tile.Interface{tile.New(5, "D", tile.Orphan{})})
	bd.SetTiles(corporations[1], []tile.Interface{tile.New(7, "D", tile.Orphan{})})

	game, _ := New(bd, players, corporations, ts)
	players[0].(*player.Stub).SetTiles([]tile.Interface{unplayableTile})
	game.tileset.(*tileset.Stub).DiscardTile(unplayableTile)
	game.state = &fsm.BuyStock{}
	game.drawTile()
	for _, tile := range players[0].Tiles() {
		if tile.Number() == unplayableTile.Number() && tile.Letter() == unplayableTile.Letter() {
			t.Errorf("Unplayable tile not discarded after drawing new tile, got %v", players[0].Tiles())
		}
	}
}

func setup() ([]player.Interface, [7]corporation.Interface, board.Interface, tileset.Interface) {
	var players []player.Interface
	players = append(players, player.NewStub("Test1"))
	players = append(players, player.NewStub("Test2"))
	players = append(players, player.NewStub("Test3"))

	var corporations [7]corporation.Interface
	corporations[0] = corporation.NewStub("A", 0)
	corporations[1] = corporation.NewStub("B", 0)
	corporations[2] = corporation.NewStub("C", 1)
	corporations[3] = corporation.NewStub("D", 1)
	corporations[4] = corporation.NewStub("E", 1)
	corporations[5] = corporation.NewStub("F", 2)
	corporations[6] = corporation.NewStub("G", 2)

	board := board.New()
	tileset := tileset.NewStub()
	return players, corporations, board, tileset
}

func slicesSameContent(slice1 []player.ShareInterface, slice2 []player.ShareInterface) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	var inSlice bool
	for _, val1 := range slice1 {
		inSlice = false
		for _, val2 := range slice2 {
			if val1 == val2 {
				inSlice = true
				break
			}
		}
		if !inSlice {
			return false
		}
	}
	return true
}
