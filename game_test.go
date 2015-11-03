package game

import (
	"github.com/svera/acquire/board"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/fsm"
	"github.com/svera/acquire/player"
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

func TestNewGameNotUniqueCorpIds(t *testing.T) {
	players, corporations, board, tileset := setup()

	corporations[0] = corporation.NewStub("A", 0, 0)
	corporations[1] = corporation.NewStub("B", 0, 0)

	if _, err := New(board, players, corporations, tileset); err.Error() != CorpIdNotUnique {
		t.Errorf("Corporations must have unique values, expecting %s error, got %s", CorpIdNotUnique, err.Error())
	}
}

func TestNewGameWrongNumberOfCorpsPerClass(t *testing.T) {
	players, corporations, board, tileset := setup()

	corporations[2] = corporation.NewStub("C", 0, 2)

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

/*
func TestPlayTileFoundCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	tileToPlay := tileset.Position{Number: 6, Letter: "E"}
	bd.PutTile(tileset.Position{Number: 5, Letter: "E"})

	game, _ := New(bd, players, corporations, ts)
	playerTiles := players[0].Tiles()
	players[0].DiscardTile(playerTiles[0])
	players[0].PickTile(tileToPlay)
	game.PlayTile(tileToPlay)

	if _, ok := game.state.(*fsm.FoundCorp); ok {
		t.Errorf("Game must be in state FoundCorp")
	}
	if corporations[0].Size() != expectedCorpSize {
		t.Errorf("Corporation size must be %d, got %d", expectedCorpSize, corporations[0].Size())
	}
}
*/
func TestPlayTileGrowCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	tileToPlay := tileset.Position{Number: 6, Letter: "E"}
	corpTiles := []tileset.Position{{Number: 7, Letter: "E"}, {Number: 8, Letter: "E"}}
	corporations[0].AddTiles(corpTiles)
	bd.SetTiles(corporations[0], corpTiles)
	bd.PutTile(tileset.Position{Number: 5, Letter: "E"})

	game, _ := New(bd, players, corporations, ts)
	playerTiles := players[0].Tiles()
	players[0].DiscardTile(playerTiles[0])
	players[0].PickTile(tileToPlay)
	game.PlayTile(tileToPlay)

	expectedCorpSize := 4

	if game.state.Type() != "BuyStock" {
		t.Errorf("Game must be in state BuyStock, got %s", game.state.Type())
	}
	if corporations[0].Size() != expectedCorpSize {
		t.Errorf("Corporation size must be %d, got %d", expectedCorpSize, corporations[0].Size())
	}
}

func TestBuyStock(t *testing.T) {
	players, corporations, bd, ts := setup()
	corporations[0].AddTiles(
		[]tileset.Position{
			{Number: 1, Letter: "A"},
			{Number: 2, Letter: "A"},
		},
	)
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

	corporations[0].AddTiles(
		[]tileset.Position{
			{Number: 1, Letter: "A"},
			{Number: 2, Letter: "A"},
		},
	)

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
	unplayableTile := tileset.Position{Number: 6, Letter: "D"}
	bd.SetTiles(corporations[0], []tileset.Position{{Number: 5, Letter: "D"}})
	bd.SetTiles(corporations[1], []tileset.Position{{Number: 7, Letter: "D"}})

	game, _ := New(bd, players, corporations, ts)
	players[0].(*player.Stub).SetTiles([]tileset.Position{unplayableTile})
	game.tileset.(*tileset.Stub).DiscardTile(unplayableTile)
	game.state = &fsm.BuyStock{}
	game.drawTile()
	for _, tile := range players[0].Tiles() {
		if tile.Number == unplayableTile.Number && tile.Letter == unplayableTile.Letter {
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
	corporations[0] = corporation.NewStub("A", 0, 0)
	corporations[1] = corporation.NewStub("B", 0, 1)
	corporations[2] = corporation.NewStub("C", 1, 2)
	corporations[3] = corporation.NewStub("D", 1, 3)
	corporations[4] = corporation.NewStub("E", 1, 4)
	corporations[5] = corporation.NewStub("F", 2, 5)
	corporations[6] = corporation.NewStub("G", 2, 6)

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
