package game

import (
	"github.com/svera/acquire/board"
	"github.com/svera/acquire/corporation"
	"github.com/svera/acquire/fsm"
	"github.com/svera/acquire/player"
	"github.com/svera/acquire/tileset"
	"testing"
)

func TestNewGame(t *testing.T) {
	players, corporations, board, tileset := setup()
	players = players[:1]

	_, err := New(board, players, corporations, tileset)

	if err == nil {
		t.Errorf("Game must not be created with less than 3 players, got %d", len(players))
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

	corporations[0].Size = func() int {
		return 41
	}

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (a corporation bigger than 40 tiles) but not detected")
	}

	corporations[0].Size = func() int {
		return 11
	}

	if !game.AreEndConditionsReached() {
		t.Errorf("End game conditions reached (all active corporations safe) but not detected")
	}

	corporations[0].Size = func() int {
		return 11
	}

	corporations[1].Size = func() int {
		return 2
	}

	if game.AreEndConditionsReached() {
		t.Errorf("End game conditions not reached but detected as it were")
	}

}

func TestGetMainStockHolders(t *testing.T) {
	_, corporations, board, tileset := setup()
	players := []player.Complete{
		player.NewStub("Test1"),
		player.NewStub("Test2"),
		player.NewStub("Test3"),
	}
	players[0].(*player.CompleteStub).SetShares(corporations[0], 8)

	game, _ := New(board, players, corporations, tileset)
	stockHolders := game.GetMainStockHolders(corporations[0])
	expectedStockHolders := map[string][]player.Sharer{
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

	players[1].(*player.CompleteStub).SetShares(corporations[0], 5)

	stockHolders = game.GetMainStockHolders(corporations[0])
	expectedStockHolders = map[string][]player.Sharer{
		"majority": {players[0]},
		"minority": {players[1]},
	}
	if !slicesSameContent(stockHolders["majority"], expectedStockHolders["majority"]) ||
		!slicesSameContent(stockHolders["minority"], expectedStockHolders["minority"]) {
		t.Errorf(
			"Wrong main stock holders",
		)
	}

	players[1].(*player.CompleteStub).SetShares(corporations[0], 8)
	players[2].(*player.CompleteStub).SetShares(corporations[0], 5)

	stockHolders = game.GetMainStockHolders(corporations[0])
	expectedStockHolders = map[string][]player.Sharer{
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

	players[1].(*player.CompleteStub).SetShares(corporations[0], 5)
	players[2].(*player.CompleteStub).SetShares(corporations[0], 5)

	stockHolders = game.GetMainStockHolders(corporations[0])
	expectedStockHolders = map[string][]player.Sharer{
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

func TestPlayTileGrowCorporation(t *testing.T) {
	players, corporations, bd, ts := setup()
	tileToPlay := tileset.Position{Number: 6, Letter: "E"}
	corpTiles := []tileset.Position{{Number: 7, Letter: "E"}, {Number: 8, Letter: "E"}}
	corporations[0].AddTiles(corpTiles)
	bd.SetTiles(corporations[0], corpTiles)
	bd.PutTile(tileset.Position{Number: 5, Letter: "E"})

	game, _ := New(bd, players, corporations, ts)
	playerTiles := players[0].Tiles()
	players[0].UseTile(playerTiles[0])
	players[0].PickTile(tileToPlay)
	game.PlayTile(tileToPlay)

	expectedCorpSize := 4

	if _, ok := game.state.(*fsm.FoundCorp); ok {
		t.Errorf("Game must be in state FoundCorp")
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
	game.buyStocks(buys)

	if corporations[0].Stock() != expectedAvailableStock {
		t.Errorf("Corporation stock shares have not decreased, must be %d, got %d", expectedAvailableStock, corporations[0].Stock())
	}
	if players[0].Shares(corporations[0]) != expectedPlayerStock {
		t.Errorf("Player stock shares have not increased, must be %d, got %d", expectedPlayerStock, players[0].Shares(corporations[0]))
	}
}

/*
func TestBuyStockWithNotEnoughCash(t *testing.T) {
	player := New("Test")
	player.cash = 100
	corporation, _ := corporation.New("Test", 0)
	corporation.AddTiles(
		[]tileset.Position{
			{Number: 1, Letter: "A"},
			{Number: 2, Letter: "A"},
		},
	)
	var buys []Buy
	buys = append(buys, Buy{corporation: corporation, amount: 2})
	err := player.BuyStocks(buys)
	if err == nil {
		t.Errorf("Trying to buy stock shares without enough money must throw error")
	}
}
*/
func setup() ([]player.Complete, [7]*corporation.Corporation, *board.Board, *tileset.Tileset) {
	var players []player.Complete
	players = append(players, player.New("Test1"))
	players = append(players, player.New("Test2"))
	players = append(players, player.New("Test3"))

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
	return players, corporations, board, tileset
}

func slicesSameContent(slice1 []player.Sharer, slice2 []player.Sharer) bool {
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
